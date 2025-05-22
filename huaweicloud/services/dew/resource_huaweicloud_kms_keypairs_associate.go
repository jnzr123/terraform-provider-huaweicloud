package dew

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API DEW POST /v3/{project_id}/keypairs/associate
// @API DEW POST /v3/{project_id}/keypairs/disassociate
// @API DEW GET /v3/{project_id}/tasks/{task_id}
func ResourceKmsKeypairsAssociate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKmsKeypairsAssociateCreate,
		ReadContext:   resourceKmsKeypairsAssociateRead,
		UpdateContext: resourceKmsKeypairsAssociateUpdate,
		DeleteContext: resourceKmsKeypairsAssociateDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the region in which to query the resource.`,
			},
			"keypair_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the name of SSH keypair.`,
			},
			"server": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Specifies ID of the virtual machine.`,
						},
						"disable_password": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Specifies whether the password is disabled.`,
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `Specifies the SSH listening port.`,
						},
						"auth": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Specifies the value of an enumeration type.`,
									},
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Specifies the value of the key depending on the authentication type.`,
									},
								},
							},
							Description: `Specifies the authentication type.`,
						},
					},
				},
				Description: `Specifies the virtual machine information that requires associating keypair.`,
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The status of the keypair being processed.`,
			},
		},
	}
}

func buildCreateBodyOpts(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"keypair_name": d.Get("keypair_name"),
		"server":       buildCreateBodyParams(d.Get("server").([]interface{})),
	}
}

func buildDeleteBodyOpts(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"server": buildDeleteBodyParams(d.Get("server").([]interface{})),
	}
}

func buildAuthBodyParams(auths []interface{}) map[string]interface{} {
	if len(auths) == 0 || auths[0] == nil {
		return nil
	}

	auth := auths[0].(map[string]interface{})
	bodyParams := map[string]interface{}{
		"type": utils.ValueIgnoreEmpty(auth["type"]),
		"key":  utils.ValueIgnoreEmpty(auth["key"]),
	}

	return bodyParams
}

func buildDeleteBodyParams(servers []interface{}) map[string]interface{} {
	if (len(servers) == 0) || servers[0] == nil {
		return nil
	}
	bodyParams := make(map[string]interface{}, 0)
	for _, v := range servers {
		item := v.(map[string]interface{})
		bodyParams = map[string]interface{}{
			"id":   item["id"],
			"auth": buildAuthBodyParams(item["auth"].([]interface{})),
		}
	}
	return bodyParams
}

func buildCreateBodyParams(servers []interface{}) map[string]interface{} {
	if (len(servers) == 0) || servers[0] == nil {
		return nil
	}
	bodyParams := make(map[string]interface{}, 0)
	for _, v := range servers {
		item := v.(map[string]interface{})
		log.Printf("[DEBUG] get CreateBodyParams: %v", item)
		bodyParams = map[string]interface{}{
			"id":               item["id"],
			"disable_password": utils.ValueIgnoreEmpty(item["disable_password"]),
			"port":             utils.ValueIgnoreEmpty(item["port"]),
			"auth":             buildAuthBodyParams(item["auth"].([]interface{})),
		}
	}

	return bodyParams
}

func taskStateRefreshFunc(taskId string, client *golangsdk.ServiceClient) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getPath := client.Endpoint + "v3/{project_id}/tasks/{task_id}"
		getPath = strings.ReplaceAll(getPath, "{project_id}", client.ProjectID)
		getPath = strings.ReplaceAll(getPath, "{task_id}", taskId)
		getOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			MoreHeaders:      map[string]string{"Content-Type": "application/json"},
		}

		getResp, err := client.Request("GET", getPath, &getOpt)
		if err != nil {
			return getResp, "ERROR", fmt.Errorf("error retrieving associate key pairs task: %s", err)
		}

		getRespBody, err := utils.FlattenResponse(getResp)
		if err != nil {
			return getRespBody, "ERROR", err
		}

		status := utils.PathSearch("task_status", getRespBody, "").(string)
		if status == "SUCCESS_BIND" || status == "SUCCESS_RESET" || status == "SUCCESS_REPLACE" {
			return getRespBody, "COMPLETED", nil
		}

		if status == "FAILED_BIND" || status == "FAILED_RESET" || status == "FAILED_REPLACE" {
			return getRespBody, "COMPLETED", errors.New("the key pairs associate failed")
		}

		return getRespBody, "PENDING", nil
	}
}

func resourceKmsKeypairsAssociateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		region  = cfg.GetRegion(d)
		httpUrl = "v3/{project_id}/keypairs/associate"
		product = "kms"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return diag.Errorf("error creating KMS client: %s", err)
	}

	requestPath := client.Endpoint + httpUrl
	requestPath = strings.ReplaceAll(requestPath, "{project_id}", client.ProjectID)
	requestOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders:      map[string]string{"Content-Type": "application/json"},
		JSONBody:         utils.RemoveNil(buildCreateBodyOpts(d)),
	}
	resp, err := client.Request("POST", requestPath, &requestOpt)
	if err != nil {
		return diag.Errorf("error creating key pairs associate: %s", err)
	}

	respBody, err := utils.FlattenResponse(resp)
	if err != nil {
		return diag.FromErr(err)
	}

	taskId := utils.PathSearch("task_id", respBody, "").(string)
	if taskId == "" {
		return diag.Errorf("unable to find task ID from API response")
	}
	d.SetId(taskId)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      taskStateRefreshFunc(taskId, client),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        10 * time.Second,
		PollInterval: 20 * time.Second,
	}
	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceKmsKeypairsAssociateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceKmsKeypairsAssociateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		region  = cfg.GetRegion(d)
		httpUrl = "v3/{project_id}/keypairs/associate"
		product = "kms"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return diag.Errorf("error creating KMS client: %s", err)
	}

	requestPath := client.Endpoint + httpUrl
	requestPath = strings.ReplaceAll(requestPath, "{project_id}", client.ProjectID)
	requestOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders:      map[string]string{"Content-Type": "application/json"},
		JSONBody:         utils.RemoveNil(buildDeleteBodyOpts(d)),
	}
	resp, err := client.Request("POST", requestPath, &requestOpt)
	if err != nil {
		return diag.Errorf("error creating key pairs associate: %s", err)
	}

	respBody, err := utils.FlattenResponse(resp)
	if err != nil {
		return diag.FromErr(err)
	}

	taskId := utils.PathSearch("task_id", respBody, "").(string)
	if taskId == "" {
		return diag.Errorf("unable to find task ID from API response")
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      taskStateRefreshFunc(taskId, client),
		Timeout:      d.Timeout(schema.TimeoutUpdate),
		PollInterval: 20 * time.Second,
	}
	if _, err = stateConf.WaitForStateContext(ctx); err != nil {
		return diag.FromErr(err)
	}

	return nil

}

func resourceKmsKeypairsAssociateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		region  = cfg.GetRegion(d)
		httpUrl = "v3/{project_id}/keypairs/disassociate"
		product = "kms"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return diag.Errorf("error creating KMS client: %s", err)
	}

	requestPath := client.Endpoint + httpUrl
	requestPath = strings.ReplaceAll(requestPath, "{project_id}", client.ProjectID)
	requestOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders:      map[string]string{"Content-Type": "application/json"},
		JSONBody:         utils.RemoveNil(buildCreateBodyOpts(d)),
	}
	resp, err := client.Request("POST", requestPath, &requestOpt)
	if err != nil {
		return diag.Errorf("error creating keypairs associate: %s", err)
	}

	respBody, err := utils.FlattenResponse(resp)
	if err != nil {
		return diag.FromErr(err)
	}

	taskId := utils.PathSearch("task_id", respBody, "").(string)
	if taskId == "" {
		return diag.Errorf("unable to find task ID from API response.")
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      taskStateRefreshFunc(taskId, client),
		Timeout:      d.Timeout(schema.TimeoutUpdate),
		PollInterval: 20 * time.Second,
	}
	if _, err = stateConf.WaitForStateContext(ctx); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
