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

func ResourceStringGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"kv": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviStringGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviStringGroupCreate,
		Read:   ResourceAviStringGroupRead,
		Update: resourceAviStringGroupUpdate,
		Delete: resourceAviStringGroupDelete,
		Schema: ResourceStringGroupSchema(),
	}
}

func ResourceAviStringGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStringGroupSchema()
	err := ApiRead(d, meta, "stringgroup", s)
	return err
}

func resourceAviStringGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStringGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "stringgroup", s)
	if err == nil {
		err = ResourceAviStringGroupRead(d, meta)
	}
	return err
}

func resourceAviStringGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStringGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "stringgroup", s)
	if err == nil {
		err = ResourceAviStringGroupRead(d, meta)
	}
	return err
}

func resourceAviStringGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "stringgroup"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviStringGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
