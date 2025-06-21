package cbh

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDeleteIntance_basic(t *testing.T) {
	// lintignore:AT001
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			// Please prepare a ID of CBH instance in fault, then config it to the environment variable.
			acceptance.TestAccPreCheckCbhInstanceID(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteIntance_basic(),
			},
		},
	})
}

func testAccDeleteIntance_basic() string {
	return fmt.Sprintf(`
resource "huaweicloud_cbh_delete_instance" "test" {
  instance_id = "%[1]s"
}
`, acceptance.HW_CBH_INSTANCE_ID)
}
