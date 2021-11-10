/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviCustomIpamDnsProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCustomIpamDnsProfileRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"script_params": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomParamsSchema(),
			},
			"script_uri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
