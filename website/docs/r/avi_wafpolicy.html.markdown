<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_wafpolicy"
sidebar_current: "docs-avi-resource-wafpolicy"
description: |-
  Creates and manages Avi WafPolicy.
---

# avi_wafpolicy

The WafPolicy resource allows the creation and management of Avi WafPolicy

## Example Usage

```hcl
resource "avi_wafpolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Field introduced in 17.2.1.
* `waf_profile_ref` - (Required) Waf profile for waf policy. It is a reference to an object of type wafprofile. Field introduced in 17.2.1.
* `allow_mode_delegation` - (Optional) Allow rules to overwrite the policy mode. This must be set if the policy mode is set to enforcement. Field introduced in 18.1.5, 18.2.1.
* `created_by` - (Optional) Creator name. Field introduced in 17.2.4.
* `crs_groups` - (Optional) Waf rules are categorized in to groups based on their characterization. These groups are system created with crs groups. Field introduced in 17.2.1.
* `description` - (Optional) Field introduced in 17.2.1.
* `enable_app_learning` - (Optional) Enable application learning for this waf policy. Field introduced in 18.2.3.
* `failure_mode` - (Optional) Waf policy failure mode. This can be 'open' or 'closed'. Enum options - WAF_FAILURE_MODE_OPEN, WAF_FAILURE_MODE_CLOSED. Field introduced in 18.1.2.
* `mode` - (Optional) Waf policy mode. This can be detection or enforcement. It can be overwritten by rules if allow_mode_delegation is set. Enum options - WAF_MODE_DETECTION_ONLY, WAF_MODE_ENFORCEMENT. Field introduced in 17.2.1.
* `paranoia_level` - (Optional) Waf ruleset paranoia  mode. This is used to select rules based on the paranoia-level tag. Enum options - WAF_PARANOIA_LEVEL_LOW, WAF_PARANOIA_LEVEL_MEDIUM, WAF_PARANOIA_LEVEL_HIGH, WAF_PARANOIA_LEVEL_EXTREME. Field introduced in 17.2.1.
* `positive_security_model` - (Optional) The positive security model. This is used to describe how the request or parts of the request should look like. It is executed in the request body phase of avi waf. Field introduced in 18.2.3.
* `post_crs_groups` - (Optional) Waf rules are categorized in to groups based on their characterization. These groups are created by the user and will be enforced after the crs groups. Field introduced in 17.2.1.
* `pre_crs_groups` - (Optional) Waf rules are categorized in to groups based on their characterization. These groups are created by the user and will be  enforced before the crs groups. Field introduced in 17.2.1.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 17.2.1.
* `waf_crs_ref` - (Optional) Waf core ruleset used for the crs part of this policy. It is a reference to an object of type wafcrs. Field introduced in 18.1.1.
* `whitelist` - (Optional) A set of rules which describe conditions under which the request will bypass the waf. This will be executed in the request header phase before any other waf related code. Field introduced in 18.2.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 17.2.1.

