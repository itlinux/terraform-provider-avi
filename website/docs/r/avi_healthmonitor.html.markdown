<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_healthmonitor"
sidebar_current: "docs-avi-resource-healthmonitor"
description: |-
  Creates and manages Avi HealthMonitor.
---

# avi_healthmonitor

The HealthMonitor resource allows the creation and management of Avi HealthMonitor

## Example Usage

```hcl
resource "avi_healthmonitor" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A user friendly name for this health monitor.
* `type` - (Required) Type of the health monitor. Enum options - HEALTH_MONITOR_PING, HEALTH_MONITOR_TCP, HEALTH_MONITOR_HTTP, HEALTH_MONITOR_HTTPS, HEALTH_MONITOR_EXTERNAL, HEALTH_MONITOR_UDP, HEALTH_MONITOR_DNS, HEALTH_MONITOR_GSLB, HEALTH_MONITOR_SIP, HEALTH_MONITOR_RADIUS.
* `allow_duplicate_monitors` - (Optional) By default, multiple instances of the same healthmonitor to the same server are suppressed intelligently. In rare cases, the monitor may have specific constructs that go beyond the server keys (ip, port, etc.) during which such suppression is not desired. Use this knob to allow duplicates. Field introduced in 18.2.8.
* `description` - (Optional) User defined description for the object.
* `disable_quickstart` - (Optional) During addition of a server or healthmonitors or during bootup, avi performs sequential health checks rather than waiting for send-interval to kick in, to mark the server up as soon as possible. This knob may be used to turn this feature off. Field introduced in 18.2.7.
* `dns_monitor` - (Optional) Dict settings for healthmonitor.
* `external_monitor` - (Optional) Dict settings for healthmonitor.
* `failed_checks` - (Optional) Number of continuous failed health checks before the server is marked down. Allowed values are 1-50.
* `http_monitor` - (Optional) Dict settings for healthmonitor.
* `https_monitor` - (Optional) Dict settings for healthmonitor.
* `is_federated` - (Optional) This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 17.1.3.
* `monitor_port` - (Optional) Use this port instead of the port defined for the server in the pool. If the monitor succeeds to this port, the load balanced traffic will still be sent to the port of the server defined within the pool. Allowed values are 1-65535. Special values are 0 - 'use server port'.
* `radius_monitor` - (Optional) Health monitor for radius. Field introduced in 18.2.3.
* `receive_timeout` - (Optional) A valid response from the server is expected within the receive timeout window. This timeout must be less than the send interval. If server status is regularly flapping up and down, consider increasing this value. Allowed values are 1-2400.
* `send_interval` - (Optional) Frequency, in seconds, that monitors are sent to a server. Allowed values are 1-3600.
* `sip_monitor` - (Optional) Health monitor for sip. Field introduced in 17.2.8, 18.1.3, 18.2.1.
* `successful_checks` - (Optional) Number of continuous successful health checks before server is marked up. Allowed values are 1-50.
* `tcp_monitor` - (Optional) Dict settings for healthmonitor.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `udp_monitor` - (Optional) Dict settings for healthmonitor.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the health monitor.

