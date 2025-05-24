---
subcategory: "Data Encryption Workshop (DEW)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_kms_get_public_key"
description: |-
  Use this data source to get the information of the public key.
---

# huaweicloud_kms_get_public_key

Use this data source to get the information of the public key.

## Example Usage

```hcl
variable "key_id" {}

data "huaweicloud_kms_get_public_key" "test" {
    key_id = var.key_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `key_id` - (Required, String) Specifies the ID of the key which is 36 bytes and meeting regular matching as
  `^ [0-9a-z] {8}- [0-9a-z] {4}- [0-9a-z] {4}- [0-9a-z] {4}- [0-9a-z]{12}$`.

* `sequence` - (Optional, String) Specifies the request sequence number which is a `36` byte serial number.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `public_key` - The information of the public key.
