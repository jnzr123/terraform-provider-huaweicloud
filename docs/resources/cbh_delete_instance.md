---
subcategory: "Cloud Bastion Host (CBH)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_cbh_delete_instance"
description: |-
  Use this resource to delete the CBH instance in fault within HuaweiCloud.
---

# huaweicloud_cbh_delete_instance

Use this resource to delete the CBH instance in fault within HuaweiCloud.

-> This resource is only a one-time action resource for deleting CBH instance in fault. Deleting this resource
  will not recover the instance, but will only remove the resource information from the tfstate file.

## Example Usage

```hcl
variable "instance_id" {}

resource "huaweicloud_cbh_delete_instance" "test" {
  instance_id = var.instance_id
} 
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource.
  If omitted, the provider-level region will be used. Changing this parameter will create a new resource.

* `instance_id` - (Required, String, NonUpdatable) Specifies the ID of instance.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.
