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

func ResourceBackupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"backup_config_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"file_name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"local_file_url": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"remote_file_url": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"scheduler_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"timestamp": &schema.Schema{
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

func resourceAviBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviBackupCreate,
		Read:   ResourceAviBackupRead,
		Update: resourceAviBackupUpdate,
		Delete: resourceAviBackupDelete,
		Schema: ResourceBackupSchema(),
	}
}

func ResourceAviBackupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupSchema()
	log.Printf("[INFO] ResourceAviBackupRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/backup/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	// no need to set the ID
	log.Printf("ResourceAviBackupRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviBackupRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviBackupRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviBackupRead Updated %v\n", d)
	return nil
}

func resourceAviBackupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupSchema()
	err := ApiCreateOrUpdate(d, meta, "backup", s)
	log.Printf("[DEBUG] created object %v: %v", "backup", d)
	if err == nil {
		err = ResourceAviBackupRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "backup", d)
	return err
}

func resourceAviBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupSchema()
	err := ApiCreateOrUpdate(d, meta, "backup", s)
	if err == nil {
		err = ResourceAviBackupRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "backup", d)
	return err
}

func resourceAviBackupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "backup"
	log.Println("[INFO] ResourceAviBackupRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviBackupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
