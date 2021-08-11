<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_authprofile"
sidebar_current: "docs-avi-resource-authprofile"
description: |-
  Creates and manages Avi AuthProfile.
---

# avi_authprofile

The AuthProfile resource allows the creation and management of Avi AuthProfile

## Example Usage

```hcl
resource "avi_authprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the auth profile.
* `type` - (Required) Type of the auth profile. Enum options - AUTH_PROFILE_LDAP, AUTH_PROFILE_TACACS_PLUS, AUTH_PROFILE_SAML, AUTH_PROFILE_PINGACCESS.
* `description` - (Optional) User defined description for the object.
* `http` - (Optional) Http user authentication params.
* `ldap` - (Optional) Ldap server and directory settings.
* `pa_agent_ref` - (Optional) Pingaccessagent uuid. It is a reference to an object of type pingaccessagent. Field introduced in 18.2.3.
* `saml` - (Optional) Saml settings. Field introduced in 17.2.3.
* `tacacs_plus` - (Optional) Tacacs+ settings.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the auth profile.

