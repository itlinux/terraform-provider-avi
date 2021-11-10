/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceCloudSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apic_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAPICConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"apic_mode": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"aws_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAwsConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"azure_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAzureConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"cloudstack_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceCloudStackConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"custom_tags": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomTagSchema(),
		},
		"dhcp_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"dns_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"docker_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceDockerConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"east_west_dns_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"east_west_ipam_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"enable_vip_static_routes": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ip6_autocfg_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ipam_provider_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"license_tier": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"license_type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"linuxserver_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceLinuxServerConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"mesos_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceMesosConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"mtu": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1500,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"nsx_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceNsxConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"obj_name_prefix": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"openstack_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceOpenStackConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"oshiftk8s_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceOShiftK8SConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"prefer_static_routes": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"proxy_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceProxyConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"rancher_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceRancherConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"state_based_dns_registration": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vca_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourcevCloudAirConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"vcenter_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourcevCenterConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"vtype": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func resourceAviCloud() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCloudCreate,
		Read:   ResourceAviCloudRead,
		Update: resourceAviCloudUpdate,
		Delete: resourceAviCloudDelete,
		Schema: ResourceCloudSchema(),
	}
}

func ResourceAviCloudRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/cloud/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
}

func resourceAviCloudCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	err := ApiCreateOrUpdate(d, meta, "cloud", s)
	if err == nil {
		err = ResourceAviCloudRead(d, meta)
	}
	return err
}

func resourceAviCloudUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCloudSchema()
	err := ApiCreateOrUpdate(d, meta, "cloud", s)
	if err == nil {
		err = ResourceAviCloudRead(d, meta)
	}
	return err
}

func resourceAviCloudDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "cloud"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviCloudDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
