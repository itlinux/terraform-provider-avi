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

func ResourceServiceEngineGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active_standby": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"advertise_backend_networks": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"aggressive_failure_detection": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"algo": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PLACEMENT_ALGO_PACKED",
		},
		"archive_shm_limit": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  8,
		},
		"async_ssl": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"async_ssl_threads": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"auto_rebalance": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"auto_rebalance_capacity_per_se": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"auto_rebalance_criteria": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"auto_rebalance_interval": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"auto_redistribute_active_standby_load": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"buffer_se": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "/api/cloud?name=Default-Cloud",
		},
		"connection_memory_percentage": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  50,
		},
		"cpu_reserve": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"cpu_socket_affinity": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"custom_securitygroups_data": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"custom_securitygroups_mgmt": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"custom_tag": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomTagSchema(),
		},
		"dedicated_dispatcher_core": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"disk_per_se": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"distribute_load_active_standby": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_routing": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_vip_on_all_interfaces": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"enable_vmac": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"extra_config_multiplier": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "0.0",
		},
		"extra_shared_config_memory": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"floating_intf_ip": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema(),
		},
		"floating_intf_ip_se_2": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema(),
		},
		"ha_mode": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "HA_MODE_SHARED",
		},
		"hardwaresecuritymodulegroup_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"hm_on_standby": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"host_attribute_key": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"host_attribute_value": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"host_gateway_monitor": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"hypervisor": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"ignore_rtt_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5000,
		},
		"ingress_access_data": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SG_INGRESS_ACCESS_ALL",
		},
		"ingress_access_mgmt": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SG_INGRESS_ACCESS_ALL",
		},
		"instance_flavor": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"iptables": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIptableRuleSetSchema(),
		},
		"least_load_core_selection": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"log_disksz": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10000,
		},
		"max_cpu_usage": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  80,
		},
		"max_scaleout_per_vs": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"max_se": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"max_vs_per_se": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"mem_reserve": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"memory_per_se": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2048,
		},
		"mgmt_network_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"mgmt_subnet": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceIpAddrPrefixSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"min_cpu_usage": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"min_scaleout_per_vs": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"non_significant_log_throttle": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"num_flow_cores_sum_changes_to_ignore": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  8,
		},
		"openstack_availability_zones": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"openstack_mgmt_network_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"openstack_mgmt_network_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"os_reserved_memory": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"per_app": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"placement_mode": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PLACEMENT_MODE_AUTO",
		},
		"realtime_se_metrics": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceMetricsRealTimeUpdateSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"se_deprovision_delay": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  120,
		},
		"se_dos_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceDosThresholdProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"se_ipc_udp_port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1500,
		},
		"se_name_prefix": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "Avi",
		},
		"se_probe_port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  7,
		},
		"se_remote_punt_udp_port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1501,
		},
		"se_sb_dedicated_core": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"se_sb_threads": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"se_thread_multiplier": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"se_tunnel_mode": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_tunnel_udp_port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1550,
		},
		"se_udp_encap_ipc": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_vs_hb_max_pkts_in_batch": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  8,
		},
		"se_vs_hb_max_vs_in_pkt": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  256,
		},
		"service_ip_subnets": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPrefixSchema(),
		},
		"significant_log_throttle": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"udf_log_throttle": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vcenter_clusters": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceVcenterClustersSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"vcenter_datastore_mode": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "VCENTER_DATASTORE_ANY",
		},
		"vcenter_datastores": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVcenterDatastoreSchema(),
		},
		"vcenter_datastores_include": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"vcenter_folder": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "AviSeFolder",
		},
		"vcenter_hosts": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceVcenterHostsSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"vcpus_per_se": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"vs_host_redundancy": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"vs_scalein_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"vs_scalein_timeout_for_upgrade": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"vs_scaleout_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"vss_placement": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceVssPlacementSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"waf_mempool": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"waf_mempool_size": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  64,
		},
	}
}

func resourceAviServiceEngineGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServiceEngineGroupCreate,
		Read:   ResourceAviServiceEngineGroupRead,
		Update: resourceAviServiceEngineGroupUpdate,
		Delete: resourceAviServiceEngineGroupDelete,
		Schema: ResourceServiceEngineGroupSchema(),
	}
}

func ResourceAviServiceEngineGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineGroupSchema()
	log.Printf("[INFO] ResourceAviServiceEngineGroupRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/serviceenginegroup/" + uuid.(string)
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
	log.Printf("ResourceAviServiceEngineGroupRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviServiceEngineGroupRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviServiceEngineGroupRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviServiceEngineGroupRead Updated %v\n", d)
	return nil
}

func resourceAviServiceEngineGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "serviceenginegroup", s)
	log.Printf("[DEBUG] created object %v: %v", "serviceenginegroup", d)
	if err == nil {
		err = ResourceAviServiceEngineGroupRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "serviceenginegroup", d)
	return err
}

func resourceAviServiceEngineGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "serviceenginegroup", s)
	if err == nil {
		err = ResourceAviServiceEngineGroupRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "serviceenginegroup", d)
	return err
}

func resourceAviServiceEngineGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "serviceenginegroup"
	log.Println("[INFO] ResourceAviServiceEngineGroupRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviServiceEngineGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
