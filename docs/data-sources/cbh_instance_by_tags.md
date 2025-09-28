---
subcategory: "Cloud Bastion Host (CBH)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_cbh_instance_by_tags"
description: ""
---

# huaweicloud_cbh_instance_by_tags

Use this data source to filter CBH instances by tags.

## Example Usage

```hcl
variable "key" {}
variable "value" {}
variable "enterprise_project_id" {}

data "huaweicloud_cbh_instance_by_tags" "test" {
  tags {
    key    = var.key
    values = [var.value]
  }
  
  sys_tags {
    key    = "_sys_enterprise_project_id"
    values = [var.enterprise_project_id]
  }
  
  matches {
    key   = "resource_name"
    value = var.resource_name
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the data source.
  If omitted, the provider-level region will be used.

* `limit` - (Optional, String) Specifies the number of records to query. The default value is `1000`, 
  and the maximum value is `1000`. The value cannot be negative, and the minimum value is `1`.

* `offset` - (Optional, String) Specifies the index position. The offset starts from the first data record. 
  The default value is `0` (offset 0 data records, indicating query from the first data record). 
  The value must be a number and cannot be negative.

* `without_any_tag` - (Optional, Bool) Specifies whether to query all resources without tags. 
  When this field is **true**, the `tags`, `tags_any`, `not_tags`, and `not_tags_any` fields are ignored.

* `tags` - (Optional, List) Specifies the included tags. A maximum of `50` keys are supported, 
  and each key supports a maximum of `10` values. The [tags](#cbh_tags) structure is documented below.

* `tags_any` - (Optional, List) Specifies the included any tags. A maximum of `50` keys are supported, 
  and each key supports a maximum of `10` values. The [tags_any](#cbh_tags) structure is documented below.

* `not_tags` - (Optional, List) Specifies the excluded tags. A maximum of `50` keys are supported, 
  and each key supports a maximum of `10` values. The [not_tags](#cbh_tags) structure is documented below.

* `not_tags_any` - (Optional, List) Specifies the excluded any tags. A maximum of `50` keys are supported, 
  and each key supports a maximum of `10` values. The [not_tags_any](#cbh_tags) structure is documented below.

* `sys_tags` - (Optional, List) Specifies the system tags. Only **op_service** permission can use this field 
  to filter resource instances. The [sys_tags](#cbh_tags) structure is documented below.

* `matches` - (Optional, List) Specifies the search fields. The [matches](#cbh_matches) structure is documented below.

<a name="cbh_tags"></a>
The `tags`, `tags_any`, `not_tags`, `not_tags_any`, and `sys_tags` block supports:

* `key` - (Required, String) Specifies the tag key. The key cannot be empty and can contain a maximum of `128` characters. 
  It can contain UTF-8 letters (including Chinese), digits, spaces, and the following characters: `_`, `.`, `:`, `=`, `+`, `-`, `@`.

* `values` - (Required, List) Specifies the tag values. Each value can contain a maximum of `255` characters. 
  It can contain UTF-8 letters (including Chinese), digits, spaces, and the following characters: `_`, `.`, `:`, `/`, `=`, `+`, `-`, `@`.

<a name="cbh_matches"></a>
The `matches` block supports:

* `key` - (Required, String) Specifies the key to match, such as **resource_name**. The key is a fixed dictionary value 
  and cannot contain duplicate keys or unsupported keys.

* `value` - (Required, String) Specifies the value to match. Each value can contain a maximum of `255` unicode characters.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `total_count` - The total number of records.

* `resources` - The list of CBH instances that match the filter criteria.
  The [resources](#cbh_resources) structure is documented below.

<a name="cbh_resources"></a>
The `resources` block supports:

* `resource_id` - The instance ID.

* `resource_name` - The resource name.

* `resource_detail` - The CBH instance details.
  The [resource_detail](#cbh_resource_detail) structure is documented below.

* `tags` - The tags.
  The [tags](#cbh_resource_tags) structure is documented below.

* `sys_tags` - The system tags.
  The [sys_tags](#cbh_resource_tags) structure is documented below.

<a name="cbh_resource_detail"></a>
The `resource_detail` block supports:

* `name` - The CBH instance name.

* `server_id` - The CBH server ID.

* `instance_id` - The CBH instance ID.

* `alter_permit` - The CBH instance expansion capability. The value can be **true** or **false**.

* `enterprise_project_id` - The enterprise project ID.

* `period_num` - The CBH instance subscription period.

* `start_time` - The CBH instance start time, in timestamp format.

* `end_time` - The CBH instance end time, in timestamp format.

* `created_time` - The CBH instance creation time, in UTC format.

* `upgrade_time` - The CBH instance scheduled upgrade time, in timestamp format.

* `update` - The CBH instance upgrade capability. The value can be **OLD**, **NEW**, **CROSS_OS**, or **ROLLBACK**.

* `bastion_version` - The current version of the CBH instance.

* `az_info` - The availability zone information.
  The [az_info](#cbh_az_info) structure is documented below.

* `status_info` - The status information.
  The [status_info](#cbh_status_info) structure is documented below.

* `resource_info` - The resource information.
  The [resource_info](#cbh_resource_info) structure is documented below.

* `network` - The network information.
  The [network](#cbh_network) structure is documented below.

* `ha_info` - The high availability information.
  The [ha_info](#cbh_ha_info) structure is documented below.

<a name="cbh_az_info"></a>
The `az_info` block supports:

* `region` - The region ID where the CBH instance is located.

* `zone` - The availability zone ID where the CBH instance is located.

* `availability_zone_display` - The Chinese name of the availability zone where the CBH instance is located.

* `slave_zone` - The availability zone where the CBH standby instance is located.

* `slave_zone_display` - The Chinese name of the availability zone where the CBH standby instance is located.

<a name="cbh_status_info"></a>
The `status_info` block supports:

* `status` - The CBH instance status. The value can be **SHUTOFF**, **ACTIVE**, **DELETING**, **BUILD**, 
  **DELETED**, **ERROR**, **HAWAIT**, **FROZEN**, **UPGRADING**, **UNPAID**, **RESIZE**, **DILATATION**, or **HA**.

* `task_status` - The current task status of the CBH instance. The value can be **powering-on**, **powering-off**, 
  **rebooting**, **delete_wait**, **frozen**, **NO_TASK**, **unfrozen**, **alter**, **updating**, 
  **configuring-ha**, **data-migrating**, **rollback**, or **traffic-switchover**.

* `create_instance_status` - The status information during the CBH instance creation process. The value can be 
  **Waiting for payment**, **creating-network**, **creating-server**, **tranfering-horizontal-network**, 
  **adding-policy-route**, **configing-dns**, **starting-cbs-service**, **setting-init-conf**, or **buying-EIP**.

* `instance_status` - The CBH instance status. The value can be **building**, **deleting**, **deleted**, 
  **unpaid**, **upgrading**, **resizing**, **abnormal**, **error**, or **ok**.

* `instance_description` - The CBH instance information description.

* `fail_reason` - The reason for CBH instance creation failure.

<a name="cbh_resource_info"></a>
The `resource_info` block supports:

* `specification` - The CBH instance specification.

* `order_id` - The order ID.

* `resource_id` - The resource ID of the CBH instance, displayed in UUID format.

* `data_disk_size` - The data disk size of the CBH instance, in TB.

* `disk_resource_id` - The data disk resource IDs of the CBH instance.

<a name="cbh_network"></a>
The `network` block supports:

* `vip` - The floating IP of the CBH instance.

* `web_port` - The port number for accessing the CBH instance web interface.

* `public_ip` - The elastic public IP of the CBH instance.

* `public_id` - The ID of the elastic IP bound to the CBH instance, displayed in UUID format.

* `private_ip` - The private IP of the CBH instance.

* `vpc_id` - The VPC ID where the CBH instance is located.

* `subnet_id` - The subnet ID where the CBH instance is located.

* `security_group_id` - The security group ID to which the CBH instance belongs.

<a name="cbh_ha_info"></a>
The `ha_info` block supports:

* `ha_id` - The high availability ID.

* `instance_type` - The instance type. The value can be **master** or **slave**.

<a name="cbh_resource_tags"></a>
The `tags` and `sys_tags` block supports:

* `key` - The tag key. The maximum length is `128` characters.

* `value` - The tag value. Each value can contain a maximum of `255` characters.