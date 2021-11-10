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

func ResourceRoleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"privileges": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePermissionSchema(),
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

func resourceAviRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviRoleCreate,
		Read:   ResourceAviRoleRead,
		Update: resourceAviRoleUpdate,
		Delete: resourceAviRoleDelete,
		Schema: ResourceRoleSchema(),
	}
}

func ResourceAviRoleRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/role/" + uuid.(string)
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

func resourceAviRoleCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	err := ApiCreateOrUpdate(d, meta, "role", s)
	if err == nil {
		err = ResourceAviRoleRead(d, meta)
	}
	return err
}

func resourceAviRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	err := ApiCreateOrUpdate(d, meta, "role", s)
	if err == nil {
		err = ResourceAviRoleRead(d, meta)
	}
	return err
}

func resourceAviRoleDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "role"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviRoleDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
