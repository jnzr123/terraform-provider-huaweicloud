package dew

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccKeypairsAssociate_basic(t *testing.T) {
	// lintignore:AT001
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			// Please prepare a valid KMS keypairs ID and ECS ID, then config it to the environment variable.
			acceptance.TestAccPreCheckKmsKeyPairsName(t)
			acceptance.TestAccPreCheckECSID(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccKeypairsAssociate_basic(),
			},
			{
				Config: testAccKeypairsAssociate_reset(),
			},
			{
				Config: testAccKeypairsAssociate_replace(),
			},
			{
				Config: testAccKeypairsAssociate_disassociate(),
			},
		},
	})
}

func testAccKeypairsAssociate_basic() string {
	return fmt.Sprintf(`
resource "huaweicloud_kms_keypairs_associate" "test0" {
  keypair_name      = "%[1]s"

  server{
    id = "%[2]s"
  }
}
`, acceptance.HW_KMS_KEYPAIRS_NAME, acceptance.HW_ECS_ID)
}

func testAccKeypairsAssociate_reset() string {
	return fmt.Sprintf(`
resource "huaweicloud_kms_keypairs_associate" "test1" {
  keypair_name      = "%[1]s"
  
  server{
    id = "%[2]s"
  }
}
`, acceptance.HW_KMS_KEYPAIRS_NAME, acceptance.HW_ECS_ID)
}

func testAccKeypairsAssociate_replace() string {
	return fmt.Sprintf(`
resource "huaweicloud_kms_keypairs_associate" "test2" {
  keypair_name      = "%[1]s"
  
  server {
    id = "%[2]s"
	auth {
	  key = "password"
	  value = "1q@W3e4r5t"
	}
  }
}
`, acceptance.HW_KMS_KEYPAIRS_NAME, acceptance.HW_ECS_ID)
}

func testAccKeypairsAssociate_disassociate() string {
	return fmt.Sprintf(`
resource "huaweicloud_kms_keypairs_associate" "test3" {
  keypair_name      = "%[1]s"
  
  server{
    id = "%[2]s"
  }
}
`, acceptance.HW_KMS_KEYPAIRS_NAME, acceptance.HW_ECS_ID)
}
