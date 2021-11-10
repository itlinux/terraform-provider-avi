/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConfigUserNotAuthrzByRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"roles": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "No Roles",
			},
			"tenants": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "No Tenants",
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceClusterNodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "node",
			},
			"public_ip_or_name": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vm_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_mor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSSLKeyRSAParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exponent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  65537,
			},
			"key_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_KEY_2048_BITS",
			},
		},
	}
}

func ResourceSeVipInterfaceListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_portchannel": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vip_intf_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip_intf_ip6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip_intf_mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceConfigUpdateDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"new_resource_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"old_resource_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceControllerLicenseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active_burst_resources": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBurstResourceSchema(),
			},
			"backend_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"burst_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"customer_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disable_enforcement": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"expired_burst_resources": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBurstResourceSchema(),
			},
			"license_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_tier": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"license_tiers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCumulativeLicenseSchema(),
			},
			"licenses": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSingleLicenseSchema(),
			},
			"max_apps": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_ses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_vses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_bandwidth_limits": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSEBandwidthLimitSchema(),
			},
			"sockets": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_on": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"throughput": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"valid_until": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHostUnavailEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reasons": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHTTPStatusRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"begin": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceDnsRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRuleActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRuleMatchTargetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDisableSeMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"migrate_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsMigrateParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAlertFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filter_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter_string": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"add_networks_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmAddNetworksEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"all_seupgrade_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAllSeUpgradeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"anomaly_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAnomalyEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"apic_agent_bd_vrf_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApicAgentBridgeDomainVrfChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"apic_agent_generic_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApicAgentGenericEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"apic_agent_vs_network_error": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApicAgentVsNetworkErrorSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"avg_uptime_change_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAvgUptimeChangeDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"aws_asg_notif_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAWSASGNotifDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"aws_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAWSSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"azure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"bind_vs_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmBindVsSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"bm_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceBMSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"bootup_fail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmSeBootupFailEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"burst_checkout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceBurstLicenseDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_cluster_vip_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudClusterVipSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_dns_update_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudDnsUpdateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_health_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudHealthSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudGenericSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_ip_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudIpChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_parkintf_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudVipParkingIntfSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_se_vm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudSeVmChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_sync_services_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudSyncServicesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_tenant_del_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudTenantsDeletedSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_vip_update_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudVipUpdateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_vnic_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudVnicChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_config_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterConfigFailedEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_leader_failover_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterLeaderFailoverEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_node_add_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeAddEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_node_db_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeDbFailedEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_node_remove_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeRemoveEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_node_shutdown_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeShutdownEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_node_started_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeStartedEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_service_critical_failure_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterServiceCriticalFailureEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_service_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterServiceFailedEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_service_restored_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterServiceRestoredEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_warm_reboot_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterWarmRebootEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cntlr_host_list_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraCntlrHostUnreachableListSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_action_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigActionDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_create_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigCreateDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_delete_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigDeleteDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_password_change_request_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserPasswordChangeRequestSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_update_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUpdateDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_user_authrz_rule_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserAuthrzByRuleSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_user_login_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_user_logout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserLogoutSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_user_not_authrz_rule_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserNotAuthrzByRuleSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"container_cloud_setup": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceContainerCloudSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"container_cloud_sevice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceContainerCloudServiceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cs_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudStackSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"delete_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmDeleteSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"disable_se_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDisableSeMigrateEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"disc_summary": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraDiscSummaryDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dns_sync_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDNSVsSyncInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"docker_ucp_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDockerUCPSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dos_attack_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDosAttackEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gcp_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGCPSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"glb_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gs_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbServiceStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"host_unavail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHostUnavailEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hs_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthScoreDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ip_fail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmSeIpFailEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"license_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLicenseDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"license_expiry_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLicenseExpiryDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"marathon_service_port_conflict_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMarathonServicePortConflictSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"memory_balancer_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMemoryBalancerInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mesos_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMesosSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"metric_threshold_up_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricThresoldUpDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"metrics_db_disk_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsDbDiskEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mgmt_nw_change_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraMgmtNwChangeDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"modify_networks_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmModifyNetworksEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"network_subnet_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkSubnetInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"nw_subnet_clash_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkSubnetClashSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"nw_summarized_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSummarizedInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackClusterSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_ip_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackIpChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_lbaudit_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackLbProvAuditCheckSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_lbplugin_op_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackLbPluginOpSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_se_vm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackSeVmChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_sync_services_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackSyncServicesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_vnic_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackVnicChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_deployment_failure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolDeploymentFailureInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_deployment_success_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolDeploymentSuccessInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_server_delete_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraPoolServerDeleteDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rebalance_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRebalanceMigrateEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rebalance_scalein_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRebalanceScaleinEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rebalance_scaleout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRebalanceScaleoutEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"reboot_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmRebootSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"scheduler_action_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSchedulerActionDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_bgp_peer_state_change_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeBgpPeerStateChangeDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMgrEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_dupip_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeDupipEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_gateway_heartbeat_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeGatewayHeartbeatFailedDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_gateway_heartbeat_success_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeGatewayHeartbeatSuccessDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_geo_db_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeGeoDbDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_hb_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHBEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_hm_gs_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventGSDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_hm_gsgroup_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventGslbPoolDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_hm_pool_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventPoolDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_hm_vs_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventVsDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_ip_added_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeIpAddedEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_ip_removed_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeIpRemovedEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_ipfailure_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeIpfailureEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_persistence_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSePersistenceEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_pool_lb_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSePoolLbEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_thresh_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeThreshEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_version_check_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVersionCheckFailedEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_vnic_down_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVnicDownEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_vnic_tx_queue_stall_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVnicTxQueueStallEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"semigrate_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMigrateEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_autoscale_failed_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleFailedInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_autoscalein_complete_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleInCompleteInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_autoscalein_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleInInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_autoscaleout_complete_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleOutCompleteInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_autoscaleout_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleOutInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"seupgrade_disrupted_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeVsDisruptedEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"seupgrade_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"seupgrade_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeMigrateEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"seupgrade_scalein_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeScaleinEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"seupgrade_scaleout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeScaleoutEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"spawn_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmSpawnSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_expire_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLExpireDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_export_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLExportDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_renew_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLRenewDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_renew_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLRenewFailedDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"switchover_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSwitchoverEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"switchover_fail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSwitchoverFailEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"system_upgrade_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSystemUpgradeDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"unbind_vs_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmUnbindVsSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vca_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVCASetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcenter_connectivity_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterConnectivityStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcenter_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterBadCredentialsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcenter_disc_failure": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterDiscoveryFailureSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcenter_obj_delete_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterObjDeleteDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip_dns_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDNSRegisterInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVmDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_awaitingse_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsAwaitingSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_error_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsErrorEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_fsm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsFsmEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_initialplacement_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsInitialPlacementEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsMigrateEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_pool_nw_fltr_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsPoolNwFilterEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_scalein_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleInEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_scaleout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleOutEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceGslbPoolMemberRuntimeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"datapath_status": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberDatapathStatusSchema(),
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ip_value_to_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oper_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"services": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServiceSchema(),
			},
			"site_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_pools": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbServiceSitePersistencePoolSchema(),
			},
			"vip_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vserver_l4_metrics": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVserverL4MetricsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vserver_l7_metrics": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVserverL7MetricsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceContainerCloudServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourcePathMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"match_str": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"string_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceFloatingIpSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceRmBindVsSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"standby": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDosThresholdProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"thresh_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDosThresholdSchema(),
			},
			"thresh_period": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceWafExcludeListEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"match_element": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVsInitialPlacementEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_assigned": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSwitchoverFailEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVIMgrVcenterRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"datacenter_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"disc_end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disc_start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovered_datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inventory_progress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inventory_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"num_clusters": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_dcs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_hosts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_nws": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vcenter_req_pending": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vms": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"progress": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			"vcenter_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vcenter_fullname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_template_se_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceLdapDirectorySettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_bind_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "(objectClass=*)",
			},
			"group_member_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "member",
			},
			"group_member_is_full_dn": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"group_search_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_search_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_LDAP_SCOPE_SUBTREE",
			},
			"ignore_referrals": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"user_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"user_id_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_search_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_search_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_LDAP_SCOPE_ONE",
			},
		},
	}
}

func ResourceSeHmEventGslbPoolMemberDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAppInfoSchema(),
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"failure_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"shm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeHmEventShmDetailsSchema(),
			},
			"ssl_error_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDnsRuleActionGslbSiteSelectionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fallback_site_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"site_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceRmModifyVnicSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceRateProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateLimiterActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"burst_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"explicit_tracking": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fine_grain": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"http_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}

func ResourceScaleStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_se_assigned": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_requested": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason_code_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scale_se": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_placement_resolution_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVipPlacementResolutionInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceClusterServiceRestoredEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGslbServiceStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"details": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gs_runtime": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbServiceRuntimeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceKeyValueSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGcpInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGslbDnsSeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fd_download": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDownloadStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"fd_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudVnicChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_addrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_VnicInfoSchema(),
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeHmEventVsDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ha_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip6_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHTTPResponsePolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPResponseRuleSchema(),
			},
		},
	}
}

func ResourceTCPProxyProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aggressive_congestion_avoidance": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"automatic": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"cc_algo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CC_ALGO_NEW_RENO",
			},
			"idle_connection_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},
			"idle_connection_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "KEEP_ALIVE",
			},
			"ignore_time_wait": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ip_dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_retransmissions": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8,
			},
			"max_segment_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_syn_retransmissions": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8,
			},
			"nagles_algorithm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"receive_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"time_wait_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"use_interface_mtu": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceDebugServiceEngineSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_shares": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugSeCpuSharesSchema(),
			},
			"flags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugSeDataplaneSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VM name unknown",
			},
			"seagent_debug": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugSeAgentSchema(),
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
		},
	}
}

func ResourceSwitchoverEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"to_se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMetricsMissingDataIntervalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceConfigInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queue": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVersionInfoSchema(),
			},
			"reader_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"writer_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceWafRuleMatchDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_internal": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"match_element": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCifSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cif": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDNSConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceOpenStackSeVmChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourcePortRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceDnsServiceApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aaaa_empty_response": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"authoritative_domain_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dns_over_tcp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"domain_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ecs_stripping_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"edns": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"edns_client_subnet_prefix_len": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"error_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_ERROR_RESPONSE_NONE",
			},
			"negative_caching_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"num_dns_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
		},
	}
}

func ResourceOpenStackLbPluginOpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"detail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"elapsed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"prov": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourcePoolDeploymentSuccessInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prev_in_service_pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"results": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePGDeploymentRuleResultSchema(),
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webhook_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webhook_result": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func ResourceMetricThresoldUpDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_value": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pool_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func ResourceSSLRenewFailedDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVlanRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceAnalyticsPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_insights": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NO_INSIGHTS",
			},
			"client_insights_sampling": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClientInsightsSamplingSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"client_log_filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceClientLogFilterSchema(),
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"full_client_logs": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFullClientLogsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"metrics_realtime_update": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRealTimeUpdateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"significant_log_throttle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"udf_log_throttle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceGslbDnsGeoUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"obj_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbObjectInfoSchema(),
			},
			"ops": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceVipRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ev": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ev_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsEvStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"first_se_assigned_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"first_time_placement": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fsm_state_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
			},
			"fsm_state_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VipFsmMap::Inactive",
			},
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"last_scale_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceScaleStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"marked_for_delete": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"migrate_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"migrate_request": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsMigrateParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"migrate_scalein_pending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"migrate_scaleout_pending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"num_additional_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"prev_metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"progress_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"requested_resource": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"scale_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceScaleStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"scalein_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"scalein_request": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleinParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"scaleout_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeListSchema(),
			},
			"servers_configured": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"supp_runtime_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"user_scaleout_pending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceUserActivitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"concurrent_sessions": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"failed_login_attempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"last_login_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_login_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_password_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logged_in": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"previous_password": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePGDeploymentRuleResultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_value": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourcePGDeploymentRuleSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceDNSRegisterInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsInfoSchema(),
			},
			"error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"total_records": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceCloudIpChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ip_mask": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
			"mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDnsOptRecordSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dnssec_ok": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsEdnsOptionSchema(),
			},
			"udp_payload_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceMetricsDimensionDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dimension": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"dimension_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceConfigUserPasswordChangeRequestSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAuthTacacsPlusAttributeValuePairSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mandatory": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceTagSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "USER_DEFINED",
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceWafRuleGroupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"exclude_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafExcludeListEntrySchema(),
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleSchema(),
			},
		},
	}
}

func ResourceIpamDnsInternalProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_service_domain": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsServiceDomainSchema(),
			},
			"dns_virtualservice_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"usable_network_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceApicAgentVsNetworkErrorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDnsRequestSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_records_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"answer_records_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authentic_data": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"checking_disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"client_location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"identifier": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nameserver_records_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"opcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"opt_record": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsOptRecordSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query_or_response": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"question_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"recursion_desired": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func ResourceWafRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"exclude_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafExcludeListEntrySchema(),
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGeoLocationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"latitude": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"longitude": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceNetworkSubnetClashSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_nw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDosRateLimitProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dos_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDosThresholdProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rl_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateLimiterProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceBgpPeerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"advertise_snat_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"advertise_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"advertisement_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"connect_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"ebgp_multihop": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"hold_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  180,
			},
			"keepalive_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"local_as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"md5_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"peer_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"peer_ip6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"remote_as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"shutdown": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"subnet6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceHdrMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hdr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"match_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceGslbPoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_ALGORITHM_ROUND_ROBIN",
			},
			"consistent_hash_mask": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceDnsQueryTypeMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"query_type": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSSHSeDeploymentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceStaticRouteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_gateway_monitor": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_hop": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"route_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDnsNsRdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip6_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"nsname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceApplicationLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adf": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"all_request_headers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"all_response_headers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_response_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"body_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NOT_UPDATED",
			},
			"cache_hit": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cacheable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"client_browser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_dest_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_insights": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_rtt": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_src_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"compression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"compression_percentage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"connection_error_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConnErrorInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"data_transfer_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"datascript_error_trace": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDataScriptErrorTraceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"datascript_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"headers_received_from_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"headers_sent_to_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_request_policy_rule_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_response_policy_rule_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_security_policy_rule_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"microservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"microservice_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_security_policy_rule_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistence_used": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"persistent_session_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirected_uri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"referer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"report_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"request_content_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_headers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request_served_locally_remote_site_down": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"request_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"response_content_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_headers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"response_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"response_time_first_byte": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"response_time_last_byte": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rewritten_uri_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rewritten_uri_query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_conn_src_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_conn_src_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_connection_reused": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_dest_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_response_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_response_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_response_time_first_byte": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_response_time_last_byte": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_rtt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_side_redirect_uri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_src_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_ssl_session_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_ssl_session_reused": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"service_engine": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"significance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"significant": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"significant_log": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"spdy_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_session_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"total_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udf": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"uri_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri_query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcpu_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"virtualservice": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vs_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"waf_log": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafLogSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"xff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSamlServiceProviderSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_entity_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_SAML_CLUSTER_VIP",
			},
			"sp_nodes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSamlServiceProviderNodeSchema(),
			},
			"tech_contact_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tech_contact_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVsMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceScaleStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_assigned": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSSLVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceServerIdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceMicroServiceContainerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"container_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"task_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceIpamDnsOpenstackProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"keystone_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_network_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCRLSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"body": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"common_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"distinguished_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_refreshed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"text": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceGslbPoolMemberDatapathStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"site_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeBootupPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"distribute_queues": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"distribute_vnics": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_backend_portend": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30720,
			},
			"docker_backend_portstart": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20480,
			},
			"fair_queueing_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"geo_db_granularity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"l7_conns_per_core": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16384,
			},
			"l7_resvd_listen_conns_per_core": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  256,
			},
			"log_agent_debug_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"log_agent_trace_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_dp_compression": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeBootupCompressionPropertiesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_dpdk_pmd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_emulated_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_ip_encap_ipc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_l3_encap_ipc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_log_buffer_app_blocking_dequeue": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_log_buffer_applog_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"se_log_buffer_chunk_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"se_log_buffer_conn_blocking_dequeue": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_log_buffer_connlog_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"se_log_buffer_events_blocking_dequeue": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_log_buffer_events_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  512,
			},
			"se_lro": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_pcap_pkt_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  256,
			},
			"se_pcap_pkt_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  65536,
			},
			"se_rum_sampling_nav_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_rum_sampling_nav_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_rum_sampling_res_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"se_rum_sampling_res_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"se_tx_batch_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"se_use_dpdk": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"ssl_sess_cache_per_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"ssl_sess_cache_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"tcp_syncache_hashsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8192,
			},
		},
	}
}

func ResourceCloudClusterVipSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMetricLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end_timestamp": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"report_timestamp": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"time_series": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsQueryResponseSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
	}
}

func ResourceSeUpgradeErrorsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_group_ha_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"task": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"to_se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceIpAddrPortSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceVIDCInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func ResourceHTTPResponseRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"hdr_action": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPHdrActionSchema(),
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"loc_hdr_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRewriteLocHdrActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceResponseMatchTargetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceEmailConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin@avicontroller.net",
			},
			"mail_server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "localhost",
			},
			"mail_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"smtp_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourcePortMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func ResourceGatewayMonitorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gateway_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gateway_monitor_fail_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"gateway_monitor_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"gateway_monitor_success_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceVsAwaitingSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"awaitingse_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_assigned": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGslbSiteRuntimeStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_file_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_file_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gap_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_geo_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ghm_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_glb_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gpki_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_cr_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_cr_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_del_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_del_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_upd_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_gs_upd_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_health_msgs_rxed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_health_msgs_txed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_bad_responses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_events_generated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_skip_outstanding_requests": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_of_timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceIpamDnsGCPProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_se_group_subnet": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"usable_network_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceAutoScaleMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_aws_autoscale_integration": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"intelligent_autoscale_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceRmAddVnicSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceClusterConfigFailedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeVnicTxQueueStallEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"linux_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"queue": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVIGuestvNicIPAddrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mask": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceDnsRuleMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsClientIpMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"geo_location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsGeoLocationMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsTransportProtocolMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query_name": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsQueryNameMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query_type": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsQueryTypeMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceDnsRuleActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRuleActionAllowDropSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslb_site_selection": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRuleActionGslbSiteSelectionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"response": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRuleActionResponseSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceDebugVrfSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flag": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMgmtIpAccessControlSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_access": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"shell_server_access": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"snmp_access": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssh_access": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceRancherConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_sync_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"container_port_match_http_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"coredump_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"disable_auto_backend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_frontend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_se_creation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_registry_se": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDockerRegistrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"east_west_placement_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"enable_event_subscription": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_container_port_as_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"feproxy_vips_enable_proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fleet_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_container_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"nuage_controller": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNuageSDNControllerSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rancher_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_deployment_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_CREATE_SSH",
			},
			"se_exclude_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_include_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_spawn_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"se_volume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi",
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"services_accessible_all_interfaces": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_container_ip_port": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_controller_image": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceDebugSeAgentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"log_level": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sub_module": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"trace_level": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGslbHealthMonitorProxySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"proxy_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_HEALTH_MONITOR_PROXY_PRIVATE_MEMBERS",
			},
			"site_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourcePerformanceLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_concurrent_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_throughput": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceDockerConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_sync_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"ca_tls_key_and_certificate_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_tls_key_and_certificate_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"container_port_match_http_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"coredump_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"disable_auto_backend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_frontend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_se_creation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_registry_se": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDockerRegistrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"east_west_placement_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"enable_event_subscription": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_container_port_as_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"feproxy_vips_enable_proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fleet_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_container_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"se_deployment_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_CREATE_SSH",
			},
			"se_exclude_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_include_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_spawn_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"se_volume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi",
			},
			"services_accessible_all_interfaces": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ucp_nodes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"use_container_ip_port": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_controller_image": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceDnsGeoLocationMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"geolocation_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"geolocation_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_edns_client_subnet_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceGslbServiceRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flr_state": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolRuntimeSchema(),
			},
			"ldr_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"send_event": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"send_status": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"services_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceEquivalentLabelsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"labels": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceURIParamSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tokens": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceURIParamTokenSchema(),
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGslbGeoDbFileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_GEODB_FILE_FORMAT_AVI",
			},
		},
	}
}

func ResourceChildProcessInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceSCPoolServerStateInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_server": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_states": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSCServerStateInfoSchema(),
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
		},
	}
}

func ResourceServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoscaling_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovered_networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"external_orchestration_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nw_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"prst_hdr_val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"resolve_server_by_dns": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewrite_host_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_node": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"static": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"verify_network": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vm_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceNetworkProfileUnionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tcp_fast_path_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTCPFastPathProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tcp_proxy_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTCPProxyProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"udp_fast_path_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceUDPFastPathProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceCloudTenantsDeletedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenants": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudTenantCleanupSchema(),
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceClientLogConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_significant_log_collection": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"filtered_log_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND",
			},
			"non_significant_log_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND",
			},
			"significant_log_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND",
			},
		},
	}
}

func ResourceSingleLicenseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"burst_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"created_on": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"customer_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enforced_params": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"last_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"license_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_tier": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"license_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_apps": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_ses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_bandwidth_limits": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSEBandwidthLimitSchema(),
			},
			"sockets": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_on": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"throughput": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tier_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"valid_until": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVssPlacementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core_nonaffinity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"num_subcores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
		},
	}
}

func ResourceVersionInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ops": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHealthScoreDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anomaly_penalty": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"anomaly_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"performance_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"performance_score": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"previous_value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources_penalty": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"resources_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_penalty": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"security_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sub_resource_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
	}
}

func ResourceCloudGenericSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSnmpTrapServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  162,
			},
			"user": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSnmpV3UserParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_VER2",
			},
		},
	}
}

func ResourceHTTPSwitchingActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"file": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPLocalFileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolServerSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceClusterNodeStartedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDnsARdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceStringMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"match_str": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"string_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceMetricStatisticsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_sample": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_ts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mean": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"min": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"min_ts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_samples": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sum": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"trend": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func ResourceHostAttributesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attr_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"attr_val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHTTPVersionMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"versions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceGslbPerDnsStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"geo_download": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDownloadStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslb_download": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDownloadStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"placement_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSubDomainPlacementRuntimeSchema(),
			},
			"se_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"valid_dns_vs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceDiscoveredNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceClusterNodeDbFailedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"failure_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDnsCnameRdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDebugVirtualServiceCaptureSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"num_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pkt_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
		},
	}
}

func ResourceOpenStackRoleMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avi_role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"os_role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourcevCloudAirConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_instance": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_mgmt_network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_orgnization": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vca_vdc": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHdrPersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prst_hdr_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCloudStackConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"api_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cntr_public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "KVM",
			},
			"mgmt_network_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_access_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceCloudVipParkingIntfSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"intf_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSEBandwidthLimitSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVsFsmEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_rt": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceRuntimeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceLdapUserBindSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dn_template": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "<user>",
			},
			"user_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"user_id_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceTenantConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_in_provider_context": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tenant_access_to_provider_se": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tenant_vrf": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceVIMgrClusterRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"datacenter_managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"datacenter_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func ResourceVIPGNameInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func ResourceVsScaleOutEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceScaleStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_assigned": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceNetworkRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"subnet_runtime": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSubnetRuntimeSchema(),
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
		},
	}
}

func ResourceSeGatewayHeartbeatSuccessDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gateway_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vrf_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceLinuxServerHostSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_attr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHostAttributesSchema(),
			},
			"host_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"node_availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourcevCenterConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_ip_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"management_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter_template_se_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCloudHealthSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"first_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_ok": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_fails": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCPUUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceRmRebootSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDebugVrfContextSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugVrfSchema(),
			},
		},
	}
}

func ResourceGslbSiteDnsVsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceVipSeAssignedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down_requested": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"connected": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scalein_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"snat_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"standby": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func ResourcePlacementNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourcePropertySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHSMSafenetLunaSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ha_group_num": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_ha": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"node_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHSMSafenetClientInfoSchema(),
			},
			"server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHSMSafenetLunaServerSchema(),
			},
			"server_pem": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_dedicated_network": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceMethodMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"methods": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceNuageSDNControllerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nuage_organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8443,
			},
			"nuage_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_vsd_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_policy_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAzureSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alb_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nic_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"vnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceHTTPStatusMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ranges": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPStatusRangeSchema(),
			},
			"status_codes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func ResourceVIMgrControllerRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
			"vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIControllerVnicInfoSchema(),
			},
		},
	}
}

func ResourceCdpLldpInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chassis": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmtaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"switch_info_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeUpgradeMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"migrate_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsMigrateParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMetricsMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_hw_training": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_grace_period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_first_n": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logging_freq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_cluster_map_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_metrics_db_writes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHTTPHdrValueSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"var": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVIMgrDCRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"cluster_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"host_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"interested_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrInterestedEntitySchema(),
			},
			"interested_nws": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrInterestedEntitySchema(),
			},
			"interested_vms": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrInterestedEntitySchema(),
			},
			"inventory_state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"nw_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pending_vcenter_reqs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sevm_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"vcenter_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceAzureConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zones": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cloud_credentials_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAzureNetworkInfoSchema(),
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_azure_dns": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_enhanced_ha": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_managed_disks": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceSubnetRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free_ip_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_alloced": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAllocInfoSchema(),
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"total_ip_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"used_ip_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceGslbThirdPartySiteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"hm_proxies": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbHealthMonitorProxySchema(),
			},
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbGeoLocationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"delete_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dns_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"event_cache": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceEventCacheSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"flr_state": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"ldr_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"site": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeSchema(),
			},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"third_party_sites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbThirdPartySiteRuntimeSchema(),
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPSecurityActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"file": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPLocalFileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSnmpConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"snmp_v3_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSnmpV3ConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sys_contact": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "support@avinetworks.com",
			},
			"sys_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_VER2",
			},
		},
	}
}

func ResourceGslbGeoDbEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbGeoDbFileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceIpAddrRangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"begin": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"end": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceVIControllerVnicInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"portgroup": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vnic_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIGuestvNicIPAddrSchema(),
			},
		},
	}
}

func ResourceConnErrorInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_syn_retransmit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_window_shrink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"out_of_orders": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"retransmits": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_num_window_shrink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_out_of_orders": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_retransmits": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_rx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_tx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_zero_window_size_events": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"zero_window_size_events": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceGslbServiceSitePersistencePoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_servers_up": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHTTPApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_dots_in_header_name": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cache_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHttpCacheConfigSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"client_body_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"client_header_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"client_max_body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"client_max_header_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  12,
			},
			"client_max_request_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  48,
			},
			"compression_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCompressionProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"connection_multiplexing_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"disable_keepalive_posts_msie6": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_fire_and_forget": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_request_body_buffering": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"hsts_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"hsts_max_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  365,
			},
			"http_to_https": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"httponly_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"keepalive_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"keepalive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"max_bad_rps_cip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_bad_rps_cip_uri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_bad_rps_uri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_rps_cip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_rps_cip_uri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_rps_unknown_cip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_rps_unknown_uri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_rps_uri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"pki_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"post_accept_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"secure_cookie_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_side_redirect_to_https": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"spdy_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"spdy_fwd_proxy_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssl_client_certificate_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLClientCertificateActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_client_certificate_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CLIENT_CERTIFICATE_NONE",
			},
			"ssl_everywhere_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_app_keepalive_timeout": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"websockets_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"x_forwarded_proto_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"xff_alternate_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "X-Forwarded-For",
			},
			"xff_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceTCPFastPathProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_syn_protection": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"session_idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
		},
	}
}

func ResourceSeVnicDownEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"linux_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDnsServiceDomainSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"num_dns_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"pass_through": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"record_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceControllerUpgradeStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_progress": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"rollback": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tasks_completed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceUpgradeTaskSchema(),
			},
		},
	}
}

func ResourceDebugVirtualServiceSeParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceOpenStackHypervisorPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image_properties": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePropertySchema(),
			},
		},
	}
}

func ResourceAuthMatchGroupMembershipSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"criteria": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceAvgUptimeChangeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_value": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func ResourceRebalanceScaleinEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scalein_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleinParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceClusterServiceCriticalFailureEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourcePoolServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceSamlServiceProviderNodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"signing_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"signing_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"single_signon_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceWafLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"latency_request_body_phase": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_request_header_phase": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_response_body_phase": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_response_header_phase": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rule_logs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleLogSchema(),
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAwsConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"asg_poll_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},
			"ebs_encryption": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     Resourceaws_encryptionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"free_elasticips": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"iam_assume_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "us-west-1",
			},
			"route53_integration": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"s3_encryption": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     Resourceaws_encryptionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"secret_access_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"use_iam_roles": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_sns_sqs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vpc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"zones": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAwsZoneConfigSchema(),
			},
		},
	}
}

func ResourceIpAddrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDnsInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RECORD_RESPONSE_CONSISTENT_HASH",
			},
			"cname": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsCnameRdataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_records_in_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RECORD_A",
			},
		},
	}
}

func ResourceAttackMetaDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_resp_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceWafConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allowed_http_versions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"allowed_methods": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"allowed_request_content_types": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"argument_separator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "&",
			},
			"buffer_response_body_for_inspection": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"client_file_upload_max_body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"client_nonfile_upload_max_body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"cookie_format_version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"regex_match_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"request_body_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:2,deny,status:403,log,auditlog",
			},
			"request_hdr_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:1,deny,status:403,log,auditlog",
			},
			"response_body_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:4,deny,status:403,log,auditlog",
			},
			"response_hdr_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:3,deny,status:403,log,auditlog",
			},
			"restricted_extensions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"restricted_headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"server_response_max_body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"static_extensions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceFeProxyRoutePublishConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FE_PROXY_ROUTE_PUBLISH_NONE",
			},
			"publisher_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHTTPServerReselectSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"num_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"retry_nonidempotent": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"svr_resp_code": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPReselectRespCodeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceSeListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down_requested": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"at_curr_ver": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"attach_ip_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Programming Network reachability to the Virtual Service IP in the Cloud",
			},
			"attach_ip_success": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"delete_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"download_selist_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"floating_intf_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"geo_download": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"geodb_download": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gslb_download": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_portchannel": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_standby": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2001,
			},
			"pending_download": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"scalein_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sec_idx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"snat_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0",
			},
			"vip6_subnet_mask": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"vip_intf_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip_intf_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVipInterfaceListSchema(),
			},
			"vip_intf_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_subnet_mask": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vnic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsSeVnicSchema(),
			},
		},
	}
}

func ResourceGslbDnsGsStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"num_partial_updates": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"partial_update_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSCServerStateInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceNetworkFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceIpamDnsCustomProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"custom_ipam_dns_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_params": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomParamsSchema(),
			},
			"usable_domains": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"usable_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
		},
	}
}

func ResourceIpCommunitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ip_begin": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ip_end": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceSSLClientRequestHeaderSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"request_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_header_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMetricsDataHeaderSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"derivation_data": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsDerivationDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dimension_data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsDimensionDataSchema(),
			},
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metrics_min_scale": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"metrics_sum_agg_invalid": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"missing_intervals": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsMissingDataIntervalSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"statistics": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricStatisticsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"units": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRIC_COUNT",
			},
		},
	}
}

func ResourceVlanInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dhcp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip6_autocfg_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_mgmt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vnic_networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICNetworkSchema(),
			},
			"vrf_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHealthMonitorSSLAttributesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pki_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_key_and_certificate_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceIpamDnsAwsProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iam_assume_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_access_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"usable_domains": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"usable_network_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"use_iam_roles": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vpc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"zones": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAwsZoneNetworkSchema(),
			},
		},
	}
}

func ResourceDebugSeDataplaneSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flag": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSeUpgradeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"num_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_grp_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceRmSeIpFailEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmAddVnicSchema(),
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGslbSiteRuntimeCfgSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fd_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gap_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"geo_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ghm_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"glb_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gpki_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gs_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mm_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sync_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteCfgSyncInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourcevNICNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ctlr_alloc": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSePoolLbEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"failure_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeBgpPeerStateChangeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"peer_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"peer_state": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vrf_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceContainerCloudSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_access": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"failed_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fleet_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"master_nodes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"missing_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"new_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_deploy_method_access": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMesosSeResourcesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"attribute_value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "2.0",
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
		},
	}
}

func ResourceJobEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"expires_at": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"owner": &schema.Schema{
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
		},
	}
}

func ResourceVsPoolNwFilterEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceRebalanceScaleoutEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scaleout_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleoutParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAutoScaleOpenStackSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"heat_scale_down_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"heat_scale_up_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeRateLimitersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arp_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"default_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"flow_probe_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  250,
			},
			"icmp_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"icmp_rsp_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  500,
			},
			"rst_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}

func ResourceAuthProfileHTTPClientParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cache_expiration_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"group_member_is_full_dn": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"request_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"require_user_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceAPICLifsRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_allocated": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cifs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCifSchema(),
			},
			"lif_label": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"multi_vrf": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transaction_uuid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDnsSrvRdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default.host",
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceClusterNodeShutdownEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMemberInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGslbServiceDownResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fallback_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_SERVICE_DOWN_RESPONSE_NONE",
			},
		},
	}
}

func ResourceMesosSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"failed_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fleet_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mesos_access": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mesos_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"missing_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"new_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_deploy_method_access": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAuthMappingRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"assign_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"assign_tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAuthMatchAttributeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"group_match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAuthMatchGroupMembershipSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"is_superuser": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"role_attribute_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tenant_attribute_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSeVersionCheckFailedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVsErrorEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_assigned": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVCASetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVIMgrNWRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_vrf_context": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_expand": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"datacenter_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dvs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"host_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"interested_nw": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ip_subnet": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrIPSubnetRuntimeSchema(),
			},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmtnw": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"num_ports": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"switch_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVlanRangeSchema(),
			},
			"vm_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vrf_context_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAWSSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceConfigUserAuthrzByRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"roles": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenants": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeAgentPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_echo_miss_aggressive_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"controller_echo_miss_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"controller_echo_rpc_aggressive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"controller_echo_rpc_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"controller_heartbeat_miss_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"controller_heartbeat_timeout_sec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  12,
			},
			"controller_registration_timeout_sec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"controller_rpc_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"cpustats_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"ctrl_reg_pending_max_wait_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  150,
			},
			"debug_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dp_aggressive_deq_interval_msec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"dp_aggressive_enq_interval_msec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"dp_batch_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"dp_deq_interval_msec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"dp_enq_interval_msec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"dp_max_wait_rsp_time_sec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"dp_reg_pending_max_wait_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  75,
			},
			"headless_timeout_sec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"ignore_docker_mac_change": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"sdb_flush_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"sdb_pipeline_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"sdb_scan_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"vnic_dhcp_ip_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"vnic_dhcp_ip_max_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"vnic_ip_delete_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"vnic_probe_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
		},
	}
}

func ResourceHTTPClientAuthenticationParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_uri_path": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceResponseMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cookie": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCookieMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hdrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHdrMatchSchema(),
			},
			"host_hdr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHostHdrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"loc_hdr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLocationHdrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"method": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMethodMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"path": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePathMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProtocolMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceQueryMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rsp_hdrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHdrMatchSchema(),
			},
			"status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPStatusMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"version": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPVersionMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceHTTPRewriteURLActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_hdr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"path": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamQuerySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceRmDeleteSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceLinuxConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"banner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"motd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceClusterServiceFailedEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVsApicExtensionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"txn_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsSeVnicSchema(),
			},
		},
	}
}

func ResourceDnsResourceRecordSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addr6_ip_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"addr_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dclass": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCustomTagSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tag_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tag_val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAlertRuleEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"event_details": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceEventDetailsFilterSchema(),
			},
			"event_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"not_cond": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceTCPApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"proxy_protocol_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"proxy_protocol_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "PROXY_PROTOCOL_VERSION_1",
			},
		},
	}
}

func ResourceConnectionLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adf": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"average_turntime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"client_dest_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_rtt": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"client_src_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"connection_ended": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"dns_etype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"dns_qtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_request": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsRequestSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dns_response": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsResponseSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslbpool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslbservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslbservice_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"microservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"microservice_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mss": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"network_security_policy_rule_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_syn_retransmit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_transaction": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_window_shrink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"out_of_orders": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"report_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"retransmits": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"rx_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"rx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_conn_src_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_conn_src_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_dest_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_num_window_shrink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_out_of_orders": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_retransmits": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_rtt": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_rx_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_rx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_src_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_total_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_total_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_tx_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_tx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_zero_window_size_events": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"service_engine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"significance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"significant": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"significant_log": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ssl_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_session_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"total_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"total_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"total_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"tx_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"tx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"udf": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"vcpu_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"virtualservice": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vs_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"zero_window_size_events": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceMarathonConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"framework_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"marathon_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"marathon_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http://leader.mesos:8080",
			},
			"marathon_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_port_range": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"public_port_range": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin",
			},
			"use_token_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vs_name_tag_framework": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceOshiftSharedVirtualServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"virtualservice_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVSDataScriptSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"evt": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSeHmEventPoolDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ha_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"percent_servers_up": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventServerDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"src_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceNTPServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceCompressionFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"devices_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ip_addr_prefixes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ip_addr_ranges": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
			"ip_addrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ip_addrs_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IS_IN",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"user_agent": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceFailActionBackupPoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDnsResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_records_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"answer_records_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authoritative_answer": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"nameserver_records_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"opcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"opt_record": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsOptRecordSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query_or_response": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"question_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"records": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsResourceRecordSchema(),
			},
			"recursion_available": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"recursion_desired": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"response_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"truncation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func ResourceVcenterDatastoreSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datastore_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceGslbThirdPartySiteRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"site_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceGslbPoolMemberSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hm_proxies": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbHealthMonitorProxySchema(),
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbGeoLocationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"public_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVirtualServiceResourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_exclusive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_standby_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scalein_primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"scalein_se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceOpenStackIpChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeUpgradeVsDisruptedEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDebugVsDataplaneSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flag": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceCloneServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceDNSVsSyncInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"total_records": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceSchedulerActionDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_uri": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"control_script_output": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"execution_datestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduler_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDNSConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"search_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
		},
	}
}

func ResourceAlertRuleMetricSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_threshold": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceAlertMetricThresholdSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceVsScaleInEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceScaleStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_assigned": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema(),
			},
			"se_requested": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVsScaleinParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"from_se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scalein_primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceReplaceStringVarSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMetricsDataSeriesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsDataSchema(),
			},
			"header": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceMetricsDataHeaderSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceDebugIpAddrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"prefixes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ranges": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
		},
	}
}

func ResourceCloudSyncServicesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceRebalanceMigrateEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"migrate_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsMigrateParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceServerAutoScaleInCompleteInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nscalein": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"scaled_in_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerIdSchema(),
			},
		},
	}
}

func ResourceVIMgrSEVMRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"azure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cloud_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"connection_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"creation_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"deletion_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"discovery_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovery_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"flavor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guest_nic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrGuestNicRuntimeSchema(),
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"init_vnics": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_discovery": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"powerstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_group_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"segroup_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_group_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"vcenter_datacenter_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_rm_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_se_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_template_vm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vcenter_vappname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vappvendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vm_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeUpgradeScaleoutEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scaleout_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleoutParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHTTPReselectRespCodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"codes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"ranges": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPStatusRangeSchema(),
			},
			"resp_code_block": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceServicePoolSelectorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"service_pool_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"service_port_range_end": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"service_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"override_application_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"override_network_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"port_range_end": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceServerAutoScaleOutCompleteInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"launch_config_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nscaleout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"scaled_out_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerIdSchema(),
			},
		},
	}
}

func ResourceAzureInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fault_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"update_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceRateLimiterActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPLocalFileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"redirect": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRedirectActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_LOCAL_RESPONSE_STATUS_CODE_429",
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "RL_ACTION_NONE",
			},
		},
	}
}

func ResourceHSMThalesRFSSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9004,
			},
		},
	}
}

func ResourceAutoScaleMesosSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"force": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceAllSeUpgradeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceAzureServicePrincipalCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authentication_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCompressionProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compressible_content_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"compression": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCompressionFilterSchema(),
			},
			"remove_accept_encoding_header": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAzureClusterInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_credential_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceUDPFastPathProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"per_pkt_loadbalance": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"session_idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"snat": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceGslbStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"details": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"gslb_runtime": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbRuntimeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"site": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"third_party_site": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbThirdPartySiteRuntimeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceHardwareSecurityModuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nethsm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHSMThalesNetHsmSchema(),
			},
			"rfs": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHSMThalesRFSSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sluna": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHSMSafenetLunaSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVinfraPoolServerDeleteDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceAlertSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_script_output": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_config_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"app_events": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceApplicationLogSchema(),
			},
			"conn_events": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConnectionLogSchema(),
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_pages": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"events": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceEventLogSchema(),
			},
			"last_throttle_timestamp": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"metric_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricLogSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"related_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"summary": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"throttle_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbSiteHealthStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_gsinfo": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberRuntimeInfoSchema(),
			},
			"datapath_gsinfo": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberRuntimeInfoSchema(),
			},
			"dns_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDnsInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gap_table": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"geo_table": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"ghm_table": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"glb_table": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"gs_table": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema(),
			},
			"sw_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func ResourceURIParamQuerySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"add_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"keep_query": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceVIMgrGuestNicRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avi_internal_network": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"del_pending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"guest_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrIPSubnetRuntimeSchema(),
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Unknown",
			},
			"mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_vnic": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"network_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSeResourcesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cores_per_socket": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disk": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"hyper_threading": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"sockets": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceCloudRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"network_sync_complete": &schema.Schema{
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
				Computed: true,
			},
		},
	}
}

func ResourceDebugControllerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filters": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDebugFilterUnionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"log_level": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sub_module": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trace_level": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSystemUpgradeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"upgrade_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSystemUpgradeStateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceCC_AgentPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"async_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"async_retries_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"poll_duration_target": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"poll_fast_target": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"poll_slow_target": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  240,
			},
			"vnic_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"vnic_retries_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
		},
	}
}

func ResourceProxyConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceIPPersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_persistent_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
		},
	}
}

func ResourceHTTPRequestPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPRequestRuleSchema(),
			},
		},
	}
}

func ResourceGslbClientIpAddrGroupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"prefixes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ranges": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_IP_PUBLIC",
			},
		},
	}
}

func ResourceDiskUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cntlr_disk_free": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOverallInfoSchema(),
			},
			"cntlr_disk_usage": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePartitionInfoSchema(),
			},
			"se_disk_free": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOverallInfoSchema(),
			},
			"se_disk_usage": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePartitionInfoSchema(),
			},
		},
	}
}

func ResourceApicAgentBridgeDomainVrfChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bridge_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"new_vrf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"old_vrf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceClusterWarmRebootEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ResourceSecureChannelMetadataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVsErrorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reason": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_group_ha_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceConfigUserLogoutSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceNetworkSecurityRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceNetworkSecurityMatchTargetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rl_param": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkSecurityPolicyActionRLParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceVSDataScriptsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"vs_datascript_set_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceCC_PropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rpc_poll_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"rpc_queue_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}

func ResourceQueryMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"match_str": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"string_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceFailActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_pool": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionBackupPoolSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"local_rsp": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionHTTPLocalResponseSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"redirect": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionHTTPRedirectSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAdminAuthConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_local_user_login": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"auth_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mapping_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAuthMappingRuleSchema(),
			},
		},
	}
}

func ResourceSeHmEventServerDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAppInfoSchema(),
			},
			"failure_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"shm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeHmEventShmDetailsSchema(),
			},
			"ssl_error_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVinfraCntlrHostUnreachableListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceNTPConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ntp_authentication_keys": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNTPAuthenticationKeySchema(),
			},
			"ntp_server_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ntp_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNTPServerSchema(),
			},
		},
	}
}

func ResourceControllerVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"patch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMetricsDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_response_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"blocking_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"browser_rendering_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"client_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"connection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"dns_lookup_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"dom_content_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"is_null": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"num_samples": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page_download_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"page_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"prediction_interval_high": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"prediction_interval_low": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"redirection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"rum_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"server_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"service_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"value_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value_str_desc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"waiting_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func ResourceServerAutoScaleOutInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alertconfig_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alertconfig_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"available_capacity": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"load": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"num_scaleout_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_servers_up": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS",
			},
		},
	}
}

func ResourceTacacsPlusAuthSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorization_attrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAuthTacacsPlusAttributeValuePairSchema(),
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  49,
			},
			"server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_TACACS_PLUS_SERVICE_LOGIN",
			},
		},
	}
}

func ResourceTimeStampSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secs": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"usecs": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceIpAddrPrefixSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mask": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceEventCacheSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_state": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"exceptions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceConfigUserLoginSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDnsClientIpMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"use_edns_client_subnet_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceHSMThalesNetHsmSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"esn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"keyhash": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"module_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"remote_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9004,
			},
		},
	}
}

func ResourceOpenStackConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_tenant": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"admin_tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowed_address_pairs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"anti_affinity": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"auth_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_drive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"contrail_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"contrail_plugin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"external_networks": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"free_floatingips": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "KVM",
			},
			"hypervisor_properties": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOpenStackHypervisorPropertiesSchema(),
			},
			"img_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OS_IMG_FMT_AUTO",
			},
			"import_keystone_tenants": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"insecure": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"intf_sec_ips": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"keystone_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"map_admin_to_cloudadmin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"mgmt_network_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_owner": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"neutron_rbac": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"nuage_organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8443,
			},
			"nuage_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_virtualip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"nuage_vsd_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_security": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"prov_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOpenStackRoleMappingSchema(),
			},
			"se_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_groups": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tenant_se": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"usable_network_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"use_admin_url": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"use_internal_endpoints": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_keystone_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"use_nuagevip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceAlertRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"conn_app_log_rule": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAlertFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"event_match_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metrics_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAlertRuleMetricSchema(),
			},
			"operator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OPERATOR_AND",
			},
			"sys_event_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAlertRuleEventSchema(),
			},
		},
	}
}

func ResourceClusterLeaderFailoverEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"leader_node": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"previous_leader_node": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceDnsEdnsOptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addr_family": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"scope_prefix_len": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source_prefix_len": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"subnet_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceCloudTenantCleanupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_ports": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_secgrp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_svrgrp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceMetricsDerivationDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"derivation_fn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"exclude_derived_metric": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"include_derivation_metrics": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"join_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_ids": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"result_has_additional_fields": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"skip_backend_derivation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceDnsTransportProtocolMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceIptableRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"dnat_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dst_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dst_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"input_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"output_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"src_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAuthMatchAttributeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"criteria": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"values": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDockerUCPSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"docker_ucp_access": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"failed_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fleet_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"missing_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"new_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_deploy_method_access": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ucp_nodes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGslbSiteCfgSyncInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"errored_objects": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVersionInfoSchema(),
			},
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sync_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceIpamDnsInfobloxProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_view": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"network_view": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"usable_domains": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"usable_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"wapi_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2.0",
			},
		},
	}
}

func ResourceHTTPRewriteLocHdrActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"keep_query": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceNsxConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avi_nsx_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_manager_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_manager_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_manager_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_poll_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
		},
	}
}

func ResourceBurstLicenseDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGslbDownloadStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceRmModifyNetworksEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmModifyVnicSchema(),
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceGslbObjectInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"object_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pb_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDnsRuleActionResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authoritative": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"rcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RCODE_NOERROR",
			},
			"truncation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func Resourceaws_encryptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"master_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AWS_ENCRYPTION_MODE_NONE",
			},
		},
	}
}

func ResourceIpamDnsAzureProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"azure_serviceprincipal": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureServicePrincipalCredentialsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"azure_userpass": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureUserPassCredentialsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"usable_domains": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"usable_network_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"use_enhanced_ha": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"virtual_network_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSnmpV3UserParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_passphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "avinetworks",
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_V3_AUTH_MD5",
			},
			"priv_passphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "avinetworks",
			},
			"priv_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_V3_PRIV_DES",
			},
		},
	}
}

func ResourceGslbDnsUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"clear_on_max_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"obj_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbObjectInfoSchema(),
			},
			"send_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSeRuntimePropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAppHdrSchema(),
			},
			"baremetal_dispatcher_handles_flows": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"connections_lossy_log_rate_limiter_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"connections_udfnf_log_rate_limiter_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"disable_flow_probes": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dos_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDosThresholdProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"downstream_send_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"dp_aggressive_hb_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"dp_aggressive_hb_timeout_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"dp_hb_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"dp_hb_timeout_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"dupip_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"dupip_timeout_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"enable_hsm_log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"feproxy_vips_enable_proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"flow_table_batch_push_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"global_mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"http_rum_console_log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"http_rum_min_content_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"lbaction_num_requests_to_dispatch": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"lbaction_rq_per_request_max_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  22,
			},
			"log_agent_compress_logs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"log_agent_conn_send_buffer_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16384,
			},
			"log_agent_export_msg_buffer_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  524288,
			},
			"log_agent_export_wait_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"log_agent_file_sz_appl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_file_sz_conn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_file_sz_debug": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_file_sz_event": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_log_storage_min_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"log_agent_max_active_adf_files_per_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"log_agent_max_concurrent_rsync": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"log_agent_max_logmessage_proto_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  65536,
			},
			"log_agent_max_storage_excess_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  110,
			},
			"log_agent_max_storage_ignore_percent": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "20.0",
			},
			"log_agent_min_storage_per_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"log_agent_pause_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"log_agent_sleep_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"log_agent_unknown_vs_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1800,
			},
			"log_message_max_file_list_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"mcache_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mcache_fetch_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mcache_max_cache_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mcache_store_in_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mcache_store_in_max_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mcache_store_in_min_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mcache_store_out_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mcache_store_se_max_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"ngx_free_connection_stack": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"persistence_mem_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"scaleout_udp_per_pkt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_auth_ldap_bind_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5000,
			},
			"se_auth_ldap_cache_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100000,
			},
			"se_auth_ldap_connect_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_auth_ldap_conns_per_server": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_auth_ldap_reconnect_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_auth_ldap_request_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_auth_ldap_servers_failover_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_dp_compression": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeRuntimeCompressionPropertiesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_dp_hm_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_dp_if_state_poll_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"se_dp_log_nf_enqueue_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  70,
			},
			"se_dp_log_udf_enqueue_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  90,
			},
			"se_dp_vnic_queue_stall_event_sleep": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_dp_vnic_queue_stall_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"se_dp_vnic_queue_stall_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_dp_vnic_restart_on_queue_stall_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"se_dp_vnic_stall_se_restart_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600,
			},
			"se_handle_interface_routes": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_hb_persist_fudge_bits": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"se_mac_error_threshold_to_disable_promiscious": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"se_memory_poison": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_metrics_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60000,
			},
			"se_metrics_rt_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_metrics_rt_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"se_packet_buffer_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"se_random_tcp_drops": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_rate_limiters": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeRateLimitersSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"service_ip_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"service_port_ranges": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"services_accessible_all_interfaces": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"spdy_fwd_proxy_parse_enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tcp_syncache_max_retransmit_default": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"upstream_connect_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"upstream_connpool_cache_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_conn_idle_thresh_tmo": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_conn_idle_tmo": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_conn_life_tmo": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_conn_max_reuse": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_core_max_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"upstream_connpool_server_max_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_strategy": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_keepalive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"upstream_read_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"upstream_send_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"user_defined_metric_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
		},
	}
}

func ResourceVipPlacementResolutionInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
			},
			"pool_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGslbPoolRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberRuntimeInfoSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVinfraVcenterObjDeleteDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"obj_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVsDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAwsZoneConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_network_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSCVsStateInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
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
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourcevNICSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"can_se_dp_takeover": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"del_pending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dhcp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6_autocfg_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_asm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_avi_internal_network": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_hsm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_mgmt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_portchannel": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"linux_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMemberInterfaceSchema(),
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"network_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pci_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vlan_interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVlanInterfaceSchema(),
			},
			"vnic_networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICNetworkSchema(),
			},
			"vrf_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vrf_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceIpAddrMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"prefixes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ranges": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
		},
	}
}

func ResourceLinuxServerConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceLinuxServerHostSchema(),
			},
			"se_inband_mgmt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_log_disk_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_log_disk_size_gb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"se_sys_disk_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_sys_disk_size_gb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"ssh_user_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeHmEventGslbPoolDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gsgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslb_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gsmember": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventGslbPoolMemberDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ha_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGslbSiteRuntimeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_leader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDnsInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"event_cache": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceEventCacheSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hs_state": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_of_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_NOT_A_MEMBER",
			},
			"rrtoken": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"site_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sw_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Not-Initialized",
			},
		},
	}
}

func ResourceSamlIdentityProviderSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metadata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSSLCertificateDescriptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"common_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"distinguished_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"locality": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHypervisor_PropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"htype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"max_ips_per_nic": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_nics": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceApplicationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
			"virtualservice_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceHTTPRequestRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"hdr_action": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPHdrActionSchema(),
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMatchTargetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"redirect_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRedirectActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rewrite_url_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRewriteURLActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"switching_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPSwitchingActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceWafDataFileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHttpCookiePersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_send_cookie": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cookie_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCookiePersistenceKeySchema(),
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceHSMSafenetLunaServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"partition_passwd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"partition_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceMemoryBalancerInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"child": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceChildProcessInfoSchema(),
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory_used": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"process": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMetricsRealTimeUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func ResourceSeMgrEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcp_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGcpInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"migrate_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceCloudDnsUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceBMSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceWafRuleLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"matches": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleMatchDataSchema(),
			},
			"msg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"phase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rule_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSeUpgradeStatusSummarySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_already_upgraded_at_start": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_disconnected_at_start": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_group_status": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeGroupStatusSchema(),
			},
			"se_ip_missing_at_start": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_poweredoff_at_start": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_completed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_errors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeUpgradeErrorsSchema(),
			},
			"se_upgrade_failed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_in_progress": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_not_started": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_retry_completed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_retry_failed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_upgrade_retry_in_progress": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_errors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsErrorSchema(),
			},
		},
	}
}

func ResourceVinfraVcenterConnectivityStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceEventLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"context": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"details_summary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"event_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"event_pages": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ignore_event_details_display": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"internal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EVENT_INTERNAL",
			},
			"is_security_event": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"module": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"obj_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"related_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"report_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourcePermissionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeIpAddedEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"linux_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mask": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCookieMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeGeoDbDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"geo_db_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"geo_db_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMemoryUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"total": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceLicenseDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expiry_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHealthMonitorDNSSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"qtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_QUERY_TYPE",
			},
			"query_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "RCODE_NO_ERROR",
			},
			"response_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeHBEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hb_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_ref1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMesosConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_vses_are_feproxy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"app_sync_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"container_port_match_http_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"coredump_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"disable_auto_backend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_frontend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_gs_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_se_creation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_registry_se": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDockerRegistrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"east_west_placement_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"enable_event_subscription": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_bridge_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cbr1",
			},
			"feproxy_container_port_as_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_route_publish": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFeProxyRoutePublishConfigSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"feproxy_vips_enable_proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fleet_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_container_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"marathon_configurations": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMarathonConfigurationSchema(),
			},
			"marathon_se_deployment": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMarathonSeDeploymentSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mesos_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http://leader.mesos:5050",
			},
			"node_availability_zone_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_controller": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNuageSDNControllerSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_deployment_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "MESOS_SE_CREATE_FLEET",
			},
			"se_exclude_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_include_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_resources": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosSeResourcesSchema(),
			},
			"se_spawn_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"se_volume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi/se",
			},
			"services_accessible_all_interfaces": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_bridge_ip_as_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_container_ip_port": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_controller_image": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_vips_for_east_west_services": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceErrorPageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"error_page_body_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPStatusMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceNetworkSecurityPolicyActionRLParamSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"burst_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"max_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceGslbIpAddrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceSSLCertificateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_signing_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"chain_verified": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"days_until_expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  365,
			},
			"expiry_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CERTIFICATE_GOOD",
			},
			"fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"issuer": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLCertificateDescriptionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"key_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLKeyParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"not_after": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"not_before": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"self_signed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLCertificateDescriptionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"subject_alt_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"text": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeThreshEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"curr_value": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceContentRewriteProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"req_match_replace_pair": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMatchReplacePairSchema(),
			},
			"request_rewrite_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"response_rewrite_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewritable_content_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsp_match_replace_pair": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMatchReplacePairSchema(),
			},
		},
	}
}

func ResourceHealthMonitorUdpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maintenance_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"udp_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"udp_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeRuntimeCompressionPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_low_rtt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"min_high_rtt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  200,
			},
			"min_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"mobile_str": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceOShiftK8SConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_sync_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"avi_bridge_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ca_tls_key_and_certificate_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_tls_key_and_certificate_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"container_port_match_http_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"coredump_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"default_service_as_east_west_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"disable_auto_backend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_frontend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_gs_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_se_creation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_registry_se": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDockerRegistrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"east_west_placement_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"enable_event_subscription": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_vips_enable_proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fleet_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_container_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"l4_health_monitoring": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"master_nodes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"node_availability_zone_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ns_exclude_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"ns_include_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"nuage_controller": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNuageSDNControllerSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sdn_overlay": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_deployment_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_CREATE_SSH",
			},
			"se_exclude_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_include_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_spawn_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"se_volume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi",
			},
			"secure_egress_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"service_account_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_port_match_http_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"shared_virtualservice_namespace": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_controller_image": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_scheduling_disabled_nodes": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_service_cluster_ip_as_ew_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSecureChannelMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_controller": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"local_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"marked_for_delete": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pub_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pub_key_pem": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SECURE_CHANNEL_NONE",
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceAbPoolSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func ResourceUpgradeTaskSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"task": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSSLRatingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compatibility_rating": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"performance_rating": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_score": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceIpAllocInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourcePGDeploymentRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "health.health_score_value",
			},
			"operator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CO_GE",
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func ResourceDnsQueryNameMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"query_domain_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"string_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDnsRecordSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DNS_RECORD_RESPONSE_ROUND_ROBIN",
			},
			"cname": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsCnameRdataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"delegated": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ip6_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsAAAARdataSchema(),
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsARdataSchema(),
			},
			"ns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsNsRdataSchema(),
			},
			"num_records_in_response": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"service_locator": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsSrvRdataSchema(),
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"wildcard_match": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSummarizedSubnetInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cidr_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceClientLogStreamingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  514,
			},
			"log_types_to_send": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_ALL",
			},
			"max_logs_per_second": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}

func ResourceServerAutoScaleFailedInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_scalein_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_servers_up": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHealthMonitorHttpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exact_http_request": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"http_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GET / HTTP/1.0",
			},
			"http_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_response_code": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"maintenance_code": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"maintenance_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_attributes": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorSSLAttributesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceGCPSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nhop_inst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nhop_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"route_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceOverallInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"available": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"free_percent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"used": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceSeIpRemovedEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"linux_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mask": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cookie": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCookieMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hdrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHdrMatchSchema(),
			},
			"host_hdr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHostHdrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"method": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMethodMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"path": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePathMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProtocolMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceQueryMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"version": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPVersionMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceSidebandProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"sideband_max_request_body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
		},
	}
}

func ResourceRmSpawnSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"network_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceRateLimiterProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip_connections_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"client_ip_failed_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"client_ip_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"client_ip_scanners_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"client_ip_to_uri_failed_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"client_ip_to_uri_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"http_header_rate_limits": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"uri_failed_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"uri_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"uri_scanners_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceVipSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_allocate_floating_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auto_allocate_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auto_allocate_ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "V4_ONLY",
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"avi_allocated_fip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"avi_allocated_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"discovered_networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"floating_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"floating_ip6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"floating_subnet6_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"floating_subnet_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ipam_network_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIPNetworkSubnetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"subnet6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"subnet6_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVIMgrIPSubnetRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fip_available": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fip_subnet_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"floatingip_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceFloatingIpSubnetSchema(),
			},
			"ip_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ref_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_ref_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourcePortalConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_basic_authentication": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_remote_cli_shell": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_clickjacking_protection": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_http": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_https": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"http_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"password_strength_check": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"redirect_to_https": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"sslkeyandcertificate_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sslprofile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_uuid_from_input": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceGslbSiteRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"clear_on_max_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"glb_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rxed_site_hs": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteHealthStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"send_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_cfg": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeCfgSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"site_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"site_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"view_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceCloudStackSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"api_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGslbGeoLocationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSSLExportDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDockerRegistrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"oshift_registry": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOshiftDockerRegistryMetaDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"private": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"registry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "avinetworks/se",
			},
			"se_repository_push": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSecureChannelTokenSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expiry_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSecureChannelMetadataSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"node_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prefix": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"static_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"static_ranges": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrRangeSchema(),
			},
		},
	}
}

func ResourceSnmpV3ConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"engine_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSnmpV3UserParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceAppHdrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hdr_match_case": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hdr_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hdr_string_op": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceBgpProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hold_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  180,
			},
			"ibgp": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"ip_communities": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpCommunitySchema(),
			},
			"keepalive_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"local_as": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"peers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBgpPeerSchema(),
			},
			"send_community": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"shutdown": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSeGroupStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disrupted_vs_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enqueue_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_with_no_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_with_vs_not_scaledout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_with_vs_scaledout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs_disrupted": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"progress": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"request_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"se_group_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"se_reboot_in_progress_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_with_no_vs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_with_vs_not_scaledout": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_with_vs_scaledout": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"thread": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_migrate_in_progress_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_scalein_in_progress_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_scaleout_in_progress_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"worker": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDosAttackEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"attack_count": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"ipgroup_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"meta_data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAttackMetaDataSchema(),
			},
			"src_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"urls": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceMicroServiceMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceLdapAuthSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"base_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bind_as_administrator": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"email_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "email",
			},
			"full_name_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "name",
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  389,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"settings": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLdapDirectorySettingsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"user_bind": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLdapUserBindSettingsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceSeMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVinfraVcenterDiscoveryFailureSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSeUpgradeScaleinEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"scalein_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleinParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAwsZoneNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"usable_network_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDnsAAAARdataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip6_address": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceVserverL7MetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apdexr": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_application_response_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_blocking_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_browser_rendering_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cache_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cache_hits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cacheable_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cacheable_hits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_txn_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_complete_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_connection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dns_lookup_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dom_content_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_error_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_errors_excluded": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_frustrated_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_page_download_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_page_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_post_compression_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_pre_compression_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_redirection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_reqs_per_session": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_1xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_2xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_3xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_4xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_4xx_avi_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_5xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_5xx_avi_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rum_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_satisfactory_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_server_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_service_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_auth_dsa": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_auth_ecdsa": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_auth_rsa": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ecdsa_non_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ecdsa_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_failed_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshake_network_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshake_protocol_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_new": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_non_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_reused": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_timedout": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_kx_dh": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_kx_ecdh": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_kx_rsa": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_rsa_non_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_rsa_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_ssl30": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_tls10": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_tls11": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_tls12": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tolerated_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_requests": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waiting_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_concurrent_sessions": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_ssl_open_sessions": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pct_cache_hits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_cacheable_hits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_response_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_ssl_failed_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_evaluated": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_flagged": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_matched": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_rejected": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"rum_apdexr": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"ssl_protocol_strength": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_application_response_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_blocking_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_browser_rendering_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_client_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dns_lookup_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dom_content_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_finished_sessions": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_client_txn_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_client_txn_latency_bucket1": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_client_txn_latency_bucket2": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_reqs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_num_page_load_time_bucket1": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_num_page_load_time_bucket2": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_num_rum_samples": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_client_txn_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_client_txn_latency_bucket1": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_client_txn_latency_bucket2": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_reqs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_page_download_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_page_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_client_txn_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_client_txn_latency_bucket1": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_client_txn_latency_bucket2": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_reqs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_redirection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_reqs_finished_sessions": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_1xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_2xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_3xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_4xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_5xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_rum_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_server_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_service_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_total_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waiting_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func ResourceAlertSyslogServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"syslog_server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"syslog_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  514,
			},
			"udp": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func ResourceServerAutoScaleInInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alertconfig_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alertconfig_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"available_capacity": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"load": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"num_scalein_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"num_servers_up": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS",
			},
			"scalein_server_candidates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerIdSchema(),
			},
		},
	}
}

func ResourceHealthMonitorTcpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maintenance_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_half_open": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tcp_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVserverL4MetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apdexc": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"apdexrtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_application_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_bytes_policy_drops": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_complete_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_connections_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_app_error": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_bad_rst_flood": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_conn": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_conn_ip_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_conn_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_fake_session": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_http_abort": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_http_error": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_http_timeout": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_malformed_flood": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_non_syn_flood": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_cookie_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_hdr_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_rl_drop_bad": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_scan_bad_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_scan_unknown_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_uri_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_uri_rl_drop_bad": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_rl_drop_bad": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_scan_bad_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_scan_unknown_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_rx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_slow_uri": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_small_window_stress": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_ssl_error": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_syn_flood": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_total_req": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_tx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_zero_window_stress": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_errored_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_l4_client_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_lossy_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_lossy_req": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_network_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_new_established_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_pkts_policy_drops": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_policy_drops": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_bytes_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_pkts": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_pkts_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_syns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tx_pkts": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_active_se": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_open_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_rx_bytes_absolute": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_rx_pkts_absolute": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_tx_bytes_absolute": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_tx_pkts_absolute": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pct_application_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_connection_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_connections_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_dos_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_dos_rx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_network_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_pkts_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_policy_drops": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_conn_duration": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_dropped_user_limit": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connections_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dup_ack_retransmits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_end_to_end_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_end_to_end_rtt_bucket1": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_end_to_end_rtt_bucket2": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_finished_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_lossy_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_lossy_req": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_out_of_orders": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_packet_dropped_user_bandwidth_limit": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_rtt_valid_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_sack_retransmits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_server_flow_control": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_timeout_retransmits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_zero_window_size_events": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func ResourceLicenseExpiryDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"burst_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expiry_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_tier": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"license_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_apps": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_ses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sockets": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"throughput": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceMatchReplacePairSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"replacement_string": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceReplaceStringVarSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceLogControllerMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_MGR_PORT_0",
			},
			"node_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prev_controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prev_metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_MGR_PORT_0",
			},
			"static_mapping": &schema.Schema{
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
				Computed: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceCloudSeVmChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCertificateAuthoritySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ca_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHSMSafenetClientInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chrystoki_conf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_priv_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_major_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"session_minor_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceAzureUserPassCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceNetworkSubnetInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"network_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"total": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"used": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceSeBootupCompressionPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"buf_num": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"buf_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"hash_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16384,
			},
			"level_aggressive": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"level_normal": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
		},
	}
}

func ResourceOpenStackSyncServicesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAlertMetricThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"comparator": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceConfigActionDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMarathonSeDeploymentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"docker_image": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "fedora",
			},
			"host_os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "COREOS",
			},
			"resource_roles": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"uris": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSeDupipEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"local_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceFullClientLogsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"throttle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}

func ResourceCloudInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cca_props": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_AgentPropertiesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"controller_props": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceControllerPropertiesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"flavor_props": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudFlavorSchema(),
			},
			"flavor_regex_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"htypes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDataScriptErrorTraceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"stack_trace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCloudMetaSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceStateCacheMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMesosAttributeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAppInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_hdr_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"app_hdr_value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDnsPoliciesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_policy_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceCloudConnectorDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disable_se_reboot": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceOpenStackLbProvAuditCheckSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"detail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"elapsed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVinfraVcenterBadCredentialsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceCloudFlavorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disk_gb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enhanced_nw": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"max_ip6s_per_nic": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_ips_per_nic": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_nics": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"meta": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudMetaSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"public": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ram_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourcePartitionInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceInternalGatewayMonitorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_gateway_monitor": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gateway_monitor_failure_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"gateway_monitor_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"gateway_monitor_success_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15,
			},
		},
	}
}

func ResourceDebugVirtualServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capture": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"capture_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDebugVirtualServiceCaptureSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"debug_hm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DEBUG_VS_HM_NONE",
			},
			"debug_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDebugIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"flags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugVsDataplaneSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"se_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDebugVirtualServiceSeParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
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
		},
	}
}

func ResourceVsEvStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHTTPPoliciesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"http_policy_set_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceURIParamTokenSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"str_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceCfgStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cfg_version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cfg_version_in_flight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS",
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceVsSeVnicSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"lif": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceOshiftDockerRegistryMetaDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"registry_namespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"registry_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "docker-registry",
			},
			"registry_vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceSePersistenceEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"patch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourcePoolGroupMemberSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deployment_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"priority_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}

func ResourceMarathonServicePortConflictSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"marathon_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceProtocolMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"protocols": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSSLRenewDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceApicAgentGenericEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"contract_graphs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lif_cif_attachment": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lifs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_engine_vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_network_attachment": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVsScaleoutParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_up": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"new_vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"to_host_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"to_new_se": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"to_se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHSMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_entity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_hs_db_writes": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func ResourceCumulativeLicenseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"burst_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_ses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_bandwidth_limits": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSEBandwidthLimitSchema(),
			},
			"sockets": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tier_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMetricsDbDiskEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_deleted_tables": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"metrics_free_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"metrics_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceVcenterHostsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"include": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceHTTPHdrDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPHdrValueSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceHealthMonitorExternalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"command_parameters": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"command_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"command_variables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceOperationalStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"reason": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason_code_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSeIpfailureEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSystemUpgradeStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceControllerUpgradeStateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rollback": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeStatusSummarySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"to_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceBurstResourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accounted_license_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_alert_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_tier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceRmUnbindVsSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVIMgrHostRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"cluster_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cntlr_accessible": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"connection_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "connected",
			},
			"cpu_hz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maintenance_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mem": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mgmt_portgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"network_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"num_cpu_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_cpu_packages": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_cpu_threads": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCdpLldpInfoSchema(),
			},
			"powerstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quarantine_start_ts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quarantined": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"quarantined_periods": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_fail_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_success_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			"vm_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceHTTPSecurityPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPSecurityRuleSchema(),
			},
		},
	}
}

func ResourceHTTPSecurityRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPSecurityActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMatchTargetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVcenterClustersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"include": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func ResourceSSLClientCertificateActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"close_connection": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSSLClientRequestHeaderSchema(),
			},
		},
	}
}

func ResourceSeUpgradeParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disruptive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"force": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"patch": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rollback": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"suspend_on_failure": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"test": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceCloudVipUpdateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDebugFilterUnionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alert_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAlertMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"autoscale_mgr_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAutoScaleMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cloud_connector_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudConnectorDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hs_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHSMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mesos_metrics_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMesosMetricsDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"metrics_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_mgr_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"state_cache_mgr_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStateCacheMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceSeHmEventShmDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"average_response_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"health_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"resp_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceDnsRuleActionAllowDropSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"reset_conn": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceGslbDnsInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dns_vs_states": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPerDnsStateSchema(),
			},
			"gs_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDnsGsStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"retry_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_table": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbDnsSeInfoSchema(),
			},
		},
	}
}

func ResourceVinfraMgmtNwChangeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"existing_nw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"new_nw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAlertMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alert_objid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cfg_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSummarizedInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSummarizedSubnetInfoSchema(),
			},
		},
	}
}

func ResourceVIMgrInterestedEntitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"interested_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceClusterNodeAddEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceClientInsightsSamplingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sample_uris": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"skip_uris": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceAWSASGNotifDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"asg_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_ip_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceMicroServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"containers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMicroServiceContainerSchema(),
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_list": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"orchestrator_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func ResourceHTTPHdrActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cookie": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPCookieDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hdr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPHdrDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceEventDetailsFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"comparator": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"event_details_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"event_details_value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceNetworkSecurityMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"microservice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMicroServiceMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceMetricsQueryResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"metric_entity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"series": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsDataSeriesSchema(),
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"stop": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHTTPRedirectActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"keep_query": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_REDIRECT_STATUS_CODE_302",
			},
		},
	}
}

func ResourceAPICConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_admin_tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "common",
			},
			"apic_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"apic_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_product": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ASP",
			},
			"apic_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Avi",
			},
			"avi_controller_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"avi_controller_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin",
			},
			"context_aware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SINGLE_CONTEXT",
			},
			"deployment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"minor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2",
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1.0",
			},
		},
	}
}

func ResourceCC_VnicInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS",
			},
			"status_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceRmSeBootupFailEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVirtualServiceRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_extension": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsApicExtensionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"datapath_debug": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDebugVirtualServiceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"east_west": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gslb_dns_geo_update": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDnsGeoUpdateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslb_dns_update": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDnsUpdateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ipam_dns_records": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsRecordSchema(),
			},
			"is_dns_vs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"lif": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"manual_placement": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"marked_for_delete": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_additional_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"one_plus_one_ha": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"prev_controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prev_metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redis_db": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"redis_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redis_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rules_configured": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"servers_configured": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tls_ticket_key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceTLSTicketSchema(),
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VS_TYPE_NORMAL",
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vh_child_vs_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vip_runtime": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipRuntimeSchema(),
			},
		},
	}
}

func ResourceSeHmEventGSDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gslb_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceTLSTicketSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aes_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hmac_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceIPNetworkSubnetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"subnet6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"subnet6_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceClusterNodeRemoveEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVIMgrVMRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"connection_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_vm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cpu_reservation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cpu_shares": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"creation_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"guest_nic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrGuestNicRuntimeSchema(),
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"init_vnics": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mem_shares": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory_reservation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"num_cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ovf_avisetype_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"powerstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
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
			"vcenter_datacenter_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_rm_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_se_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_template_vm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vcenter_vappname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vappvendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vm_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vnic_discovered": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vm_lb_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceClientLogFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uri": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceHTTPCookieDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceVinfraDiscSummaryDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_clusters": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_dcs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_hosts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_nws": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vms": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceHttpCacheConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"age_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"aggressive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"date_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"default_expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"heuristic_expire": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"max_cache_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_object_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4194304,
			},
			"mime_types_black_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mime_types_black_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mime_types_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mime_types_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"min_object_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"query_cacheable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"xcache_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func ResourceVsMigrateParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"new_vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"to_host_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"to_new_se": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"to_se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceOpenStackClusterSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"keystone_host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceOpenStackVnicChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_addrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceVinfraVmDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceDebugSeCpuSharesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"shares": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourceLocationHdrMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceSamlSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"idp": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSamlIdentityProviderSettingsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sp": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSamlServiceProviderSettingsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceSecureChannelAvailableLocalIPsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"free_controller_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"free_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceGslbSiteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_vses": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSiteDnsVsSchema(),
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"hm_proxies": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbHealthMonitorProxySchema(),
			},
			"ip_addresses": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbGeoLocationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"member_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_PASSIVE_MEMBER",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  443,
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ResourceIptableRuleSetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIptableRuleSchema(),
			},
			"table": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceAppCookiePersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"encryption_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prst_hdr_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
		},
	}
}

func ResourceSeGatewayHeartbeatFailedDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gateway_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vrf_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceDosThresholdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attack": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"max_value": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"min_value": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func ResourcePoolDeploymentFailureInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"curr_in_service_pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"results": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePGDeploymentRuleResultSchema(),
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webhook_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webhook_result": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func ResourceFailActionHTTPRedirectSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTPS",
			},
			"query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_REDIRECT_STATUS_CODE_302",
			},
		},
	}
}

func ResourceConfigCreateDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceSSLKeyParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ec_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLKeyECParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rsa_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLKeyRSAParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}

func ResourceFailActionHTTPLocalResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPLocalFileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FAIL_HTTP_STATUS_CODE_503",
			},
		},
	}
}

func ResourceCustomParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_dynamic": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_sensitive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceGslbSubDomainPlacementRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"placement_allowed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sub_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transition_ops": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_NONE",
			},
		},
	}
}

func ResourceHttpCookiePersistenceKeySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aes_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hmac_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAnomalyEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deviation": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"metric_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EXPONENTIAL_WEIGHTED_MOVING_AVG",
			},
			"models": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"node_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
	}
}

func ResourceNTPAuthenticationKeySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NTP_AUTH_ALGORITHM_MD5",
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceAlertObjectListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"objects": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func ResourceMesosMetricsDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mesos_master": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mesos_slave": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_entity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metrics_collection_frq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
		},
	}
}

func ResourceSSLKeyECParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"curve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_KEY_EC_CURVE_SECP256R1",
			},
		},
	}
}

func ResourceRmAddNetworksEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmAddVnicSchema(),
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceHostHdrMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceConfigDeleteDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAzureNetworkInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_network_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_network_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceHTTPLocalFileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"content_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"file_content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceSSLExpireDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"days_left": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
