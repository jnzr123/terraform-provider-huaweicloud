---
subcategory: "Data Encryption Workshop (DEW)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_kms_key"
description: |-
  Manages a KMS keypair associate resource within HuaweiCloud.
---

# huaweicloud_kms_keypairs_associate

Manages a KMS keypair associate resource within HuaweiCloud.

-> There are several prerequisites for using this resource:  
  1.The image ECS used must be the puclic image provided by HuaweiCloud.  
  2.The operation of keypair is achieved by modifying the server's `/root/.ssh/authorized_keys` file.
  Please ensure that the file has not been modified before using the keypair, otherwise the operation will fail.  
  3.The SSH port of the ECS security group needs to be opened for network segment `100.125.0.0/16` in advance.  
  4.There can be at most `10` ECS associating keypair simultaneously.

## Example Usage

```hcl
variable "keypair_name" {}
variable "server_id" {}

resource "huaweicloud_kms_keypairs_associate" "test" {
  keypair_name    = var.keypair_name

  server {
    id = var.server_id
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used. Changing this parameter will create a new resource.

* `keypair_name` - (Required, String) Specifies the name of SSH keypair.

* `server` - (Required, List, NonUpdatable) Specifies the virtual machine information that requires associating keypair.
  These public images which are not supported to associate keypair are as follows:  
  `CoreOS`, `OpenEuler`, `FreeBSD（Other）`, `Kylin V10 64bit`, `UnionTech OS Server 20`,  
  `Euler 64bit`, `CentOS Stream 8 64bit`.
  The [server](#server) structure is documented below.

<a name="server"></a>
The `server` block supports:

* `id` - (Required, String, NonUpdatable) Specifies ID of the virtual machine which need to associate
  (replace or reset) the SSH keypair.

* `disable_password` - (Optional, Bool, NonUpdatable) Specifies whether the password is disabled.  
  The valid values are as follows:  
  `true`: Indicates disable SSH login for virtual machines.  
  `false`: Indicates enable SSH login for virtual machines.

* `port` - (Optional, Int, NonUpdatable) Specifies the SSH listening port. The default value is 22.

* `auth` - (Optional, List) Specifies the authentication information. This parameter is **required** for **replacement**
  and not required for reset.
  The [auth](#auth) structure is documented below.

<a name="auth"></a>
The `auth` block supports:

* `type` - (Optional, String) Specifies the value of an authentication type. The valid values are `password` or `keypair`.

* `key` - (Optional, String) Specifies the value of the key depending on the `type`.
  When the `type` is set to `password`, it represents the password.  
  When the `type` is set to `keypair`, it represents the private key.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID which is also the task ID.

* `status` - The status of the keypair being processed.The details are as follows:
  `READY_BIND`, `RUNNING_BIND`, `FAILED_BIND`, `READY_RESET`, `RUNNING_RESET`, `FAILED_RESET`, `SUCCESS_RESET`, `READY_REPLACE`,
  `RUNNING_REPLACE`, `FAILED_RESET`, `SUCCESS_RESET`, `READY_UNBIND`, `RUNNING_UNBIND`, `FAILED_UNBIND`, `SUCCESS_UNBIND`.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Defaults to `10` minutes.
* `update` - Defaults to `10` minutes.
* `delete` - Defaults to `10` minutes.
