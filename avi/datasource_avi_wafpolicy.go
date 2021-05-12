// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviWafPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafPolicyRead,
		Schema: map[string]*schema.Schema{
			"allow_mode_delegation": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crs_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_app_learning": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"failure_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"paranoia_level": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"positive_security_model": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceWafPositiveSecurityModelSchema(),
			},
			"post_crs_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"pre_crs_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"waf_crs_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"waf_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"whitelist": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceWafPolicyWhitelistSchema(),
			},
		},
	}
}
