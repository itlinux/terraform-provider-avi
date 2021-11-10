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

func ResourceDnsPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"rule": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsRuleSchema(),
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

func resourceAviDnsPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviDnsPolicyCreate,
		Read:   ResourceAviDnsPolicyRead,
		Update: resourceAviDnsPolicyUpdate,
		Delete: resourceAviDnsPolicyDelete,
		Schema: ResourceDnsPolicySchema(),
	}
}

func ResourceAviDnsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/dnspolicy/" + uuid.(string)
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

func resourceAviDnsPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "dnspolicy", s)
	if err == nil {
		err = ResourceAviDnsPolicyRead(d, meta)
	}
	return err
}

func resourceAviDnsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "dnspolicy", s)
	if err == nil {
		err = ResourceAviDnsPolicyRead(d, meta)
	}
	return err
}

func resourceAviDnsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "dnspolicy"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviDnsPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
