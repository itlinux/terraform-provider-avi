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

func ResourceNetworkSecurityPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_config_cksum": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"rules": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceNetworkSecurityRuleSchema(),
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
	}
}

func resourceAviNetworkSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkSecurityPolicyCreate,
		Read:   ResourceAviNetworkSecurityPolicyRead,
		Update: resourceAviNetworkSecurityPolicyUpdate,
		Delete: resourceAviNetworkSecurityPolicyDelete,
		Schema: ResourceNetworkSecurityPolicySchema(),
	}
}

func ResourceAviNetworkSecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSecurityPolicySchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/networksecuritypolicy/" + uuid.(string)
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

func resourceAviNetworkSecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSecurityPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "networksecuritypolicy", s)
	if err == nil {
		err = ResourceAviNetworkSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviNetworkSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSecurityPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "networksecuritypolicy", s)
	if err == nil {
		err = ResourceAviNetworkSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviNetworkSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "networksecuritypolicy"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviNetworkSecurityPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
