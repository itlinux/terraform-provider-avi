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

func ResourceSystemConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_auth_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAdminAuthConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"default_license_tier": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ENTERPRISE_18",
		},
		"dns_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceDNSConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"dns_virtualservice_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"docker_mode": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"email_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceEmailConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"global_tenant_config": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceTenantConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"linux_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceLinuxConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"mgmt_ip_access_control": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceMgmtIpAccessControlSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"ntp_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceNTPConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"portal_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourcePortalConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"proxy_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceProxyConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"snmp_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSnmpConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"ssh_ciphers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ssh_hmacs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSystemConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSystemConfigurationCreate,
		Read:   ResourceAviSystemConfigurationRead,
		Update: resourceAviSystemConfigurationUpdate,
		Delete: resourceAviSystemConfigurationDelete,
		Schema: ResourceSystemConfigurationSchema(),
	}
}

func ResourceAviSystemConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemConfigurationSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/systemconfiguration/" + uuid.(string)
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

func resourceAviSystemConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemConfigurationSchema()
	err := ApiCreateOrUpdate(d, meta, "systemconfiguration", s)
	if err == nil {
		err = ResourceAviSystemConfigurationRead(d, meta)
	}
	return err
}

func resourceAviSystemConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemConfigurationSchema()
	err := ApiCreateOrUpdate(d, meta, "systemconfiguration", s)
	if err == nil {
		err = ResourceAviSystemConfigurationRead(d, meta)
	}
	return err
}

func resourceAviSystemConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "systemconfiguration"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviSystemConfigurationDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
