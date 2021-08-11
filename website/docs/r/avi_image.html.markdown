<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_image"
sidebar_current: "docs-avi-resource-image"
description: |-
  Creates and manages Avi Image.
---

# avi_image

The Image resource allows the creation and management of Avi Image

## Example Usage

```hcl
resource "avi_image" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the image. Field introduced in 18.2.6.
* `controller_info` - (Optional) Controller package details. Field introduced in 18.2.6.
* `controller_patch_name` - (Optional) Mandatory controller patch name that is applied along with this base image. Field introduced in 18.2.10.
* `controller_patch_uuid` - (Optional) It references the controller-patch associated with the uber image. Field introduced in 18.2.8.
* `migrations` - (Optional) This field describes the api migration related information. Field introduced in 18.2.6.
* `se_info` - (Optional) Se package details. Field introduced in 18.2.6.
* `se_patch_name` - (Optional) Mandatory serviceengine patch name that is applied along with this base image. Field introduced in 18.2.10.
* `se_patch_uuid` - (Optional) It references the service engine patch associated with the uber image. Field introduced in 18.2.8.
* `status` - (Optional) Status to check if the image is present. Enum options - SYSERR_SUCCESS, SYSERR_FAILURE, SYSERR_OUT_OF_MEMORY, SYSERR_NO_ENT, SYSERR_INVAL, SYSERR_ACCESS, SYSERR_FAULT, SYSERR_IO, SYSERR_TIMEOUT, SYSERR_NOT_SUPPORTED, SYSERR_NOT_READY, SYSERR_UPGRADE_IN_PROGRESS, SYSERR_WARM_START_IN_PROGRESS, SYSERR_TRY_AGAIN, SYSERR_NOT_UPGRADING, SYSERR_PENDING, SYSERR_EVENT_GEN_FAILURE, SYSERR_CONFIG_PARAM_MISSING, SYSERR_BAD_REQUEST, SYSERR_TEST1... Field introduced in 18.2.6.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.6.
* `type` - (Optional) Type of the image patch/system. Enum options - IMAGE_TYPE_PATCH, IMAGE_TYPE_SYSTEM. Field introduced in 18.2.6.
* `uber_bundle` - (Optional) Status to check if the image is an uber bundle. Field introduced in 18.2.8.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the image. Field introduced in 18.2.6.

