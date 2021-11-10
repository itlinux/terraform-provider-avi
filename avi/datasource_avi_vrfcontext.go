/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviVrfContext() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVrfContextRead,
		Schema: map[string]*schema.Schema{
			"bgp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceBgpProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"debugvrfcontext": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDebugVrfContextSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_mon": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGatewayMonitorSchema(),
			},
			"internal_gateway_monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceInternalGatewayMonitorSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"static_routes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceStaticRouteSchema(),
			},
			"system_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
