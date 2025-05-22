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
			acceptance.TestAccPreCheckKmsKeyPair(t)
			acceptance.TestAccPreCheckECSAccount(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccKeypairsAssociate_basic(),
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
  keypair_name = "%[1]s"

  server{
    id   = "%[2]s"
	port = %[3]s

	auth {
	  type = "password"
	  key  = "%[4]s"
	}
  }
}
`, acceptance.HW_KMS_KEYPAIRS_NAME_1, acceptance.HW_ECS_ID, acceptance.HW_KMS_KEYPAIRS_SSH_PORT, acceptance.HW_ECS_ROOT_PWD)
}

func testAccKeypairsAssociate_replace() string {
	return fmt.Sprintf(`
resource "huaweicloud_kms_keypairs_associate" "test1" {
  keypair_name = "%[1]s"
  
  server {
    id   = "%[2]s"
	port = %[3]s

	auth {
	  key = "keypair"
	  value = "%[4]s"
	}
  }
}
`, acceptance.HW_KMS_KEYPAIRS_NAME_2, acceptance.HW_ECS_ID, acceptance.HW_KMS_KEYPAIRS_SSH_PORT, acceptance.HW_KMS_KEYPAIRS_KEY_1)
}

func testAccKeypairsAssociate_disassociate() string {
	return fmt.Sprintf(`
resource "huaweicloud_kms_keypairs_associate" "test2" {
  server{
    id   = "%[1]s"
	port = %[2]s

	auth {
	  type = "keypair"
	  key  = "%[3]s"
	}
  }
}
`, acceptance.HW_ECS_ID, acceptance.HW_KMS_KEYPAIRS_SSH_PORT, acceptance.HW_KMS_KEYPAIRS_KEY_2)
}
