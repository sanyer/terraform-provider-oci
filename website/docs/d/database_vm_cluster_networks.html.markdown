---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_networks"
sidebar_current: "docs-oci-datasource-database-vm_cluster_networks"
description: |-
  Provides the list of Vm Cluster Networks in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_networks
This data source provides the list of Vm Cluster Networks in Oracle Cloud Infrastructure Database service.

Gets a list of the VM cluster networks in the specified compartment. Applies to Exadata Cloud@Customer instances only.


## Example Usage

```hcl
data "oci_database_vm_cluster_networks" "test_vm_cluster_networks" {
	#Required
	compartment_id = var.compartment_id
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

	#Optional
	display_name = var.vm_cluster_network_display_name
	state = var.vm_cluster_network_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `exadata_infrastructure_id` - (Required) The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `vm_cluster_networks` - The list of vm_cluster_networks.

### VmClusterNetwork Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the VM cluster network. The name does not need to be unique.
* `dns` - The list of DNS server IP addresses. Maximum of 3 allowed.
* `dr_scans` - The SCAN details for DR network
	* `hostname` - The Disaster recovery SCAN hostname.
	* `ips` - The list of Disaster recovery SCAN IP addresses. Three addresses should be provided.
	* `scan_listener_port_tcp` - The Disaster recovery SCAN TCPIP port. Default is 1521.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `ntp` - The list of NTP server IP addresses. Maximum of 3 allowed.
* `scans` - The SCAN details.
	* `hostname` - The SCAN hostname.
	* `ips` - The list of SCAN IP addresses. Three addresses should be provided.
	* `port` - **Deprecated.** This field is deprecated. You may use 'scanListenerPortTcp' to specify the port. The SCAN TCPIP port. Default is 1521. 
	* `scan_listener_port_tcp` - The SCAN TCPIP port. Default is 1521.
	* `scan_listener_port_tcp_ssl` - The SCAN TCPIP SSL port. Default is 2484.
* `state` - The current state of the VM cluster network. CREATING - The resource is being created REQUIRES_VALIDATION - The resource is created and may not be usable until it is validated. VALIDATING - The resource is being validated and not available to use. VALIDATED - The resource is validated and is available for consumption by VM cluster. VALIDATION_FAILED - The resource validation has failed and might require user input to be corrected. UPDATING - The resource is being updated and not available to use. ALLOCATED - The resource is is currently being used by VM cluster. TERMINATING - The resource is being deleted and not available to use. TERMINATED - The resource is deleted and unavailable. FAILED - The resource is in a failed state due to validation or other errors. NEEDS_ATTENTION - The resource is in needs attention state as some of it's child nodes are not validated and unusable by VM cluster. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time when the VM cluster network was created.
* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated VM Cluster.
* `vm_networks` - Details of the client and backup networks.
	* `domain_name` - The network domain name.
	* `gateway` - The network gateway.
	* `netmask` - The network netmask.
	* `network_type` - The network type.
	* `nodes` - The list of node details.
		* `db_server_id` - The Db server associated with the node.
		* `hostname` - The node host name.
		* `ip` - The node IP address.
		* `state` - The current state of the VM cluster network nodes. CREATING - The resource is being created REQUIRES_VALIDATION - The resource is created and may not be usable until it is validated. VALIDATING - The resource is being validated and not available to use. VALIDATED - The resource is validated and is available for consumption by VM cluster. VALIDATION_FAILED - The resource validation has failed and might require user input to be corrected. UPDATING - The resource is being updated and not available to use. ALLOCATED - The resource is currently being used by VM cluster. TERMINATING - The resource is being deleted and not available to use. TERMINATED - The resource is deleted and unavailable. FAILED - The resource is in a failed state due to validation or other errors. 
		* `vip` - The node virtual IP (VIP) address.
		* `vip_hostname` - The node virtual IP (VIP) host name.
	* `vlan_id` - The network VLAN ID.

