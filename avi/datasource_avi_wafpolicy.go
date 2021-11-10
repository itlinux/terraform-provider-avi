/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviWafPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafPolicyRead,
		Schema: map[string]*schema.Schema{
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"crs_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_MODE_DETECTION_ONLY"},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"paranoia_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_PARANOIA_LEVEL_LOW"},
			"post_crs_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"pre_crs_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"waf_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
