---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_blob_mount"
sidebar_current: "docs-oci-datasource-dbmulticloud-oracle_db_azure_blob_mount"
description: |-
  Provides details about a specific Oracle Db Azure Blob Mount in Oracle Cloud Infrastructure Dbmulticloud service
---

# Data Source: oci_dbmulticloud_oracle_db_azure_blob_mount
This data source provides details about a specific Oracle Db Azure Blob Mount resource in Oracle Cloud Infrastructure Dbmulticloud service.

Get Oracle DB Azure Blob Mount Details form a particular Container Resource ID.


## Example Usage

```hcl
data "oci_dbmulticloud_oracle_db_azure_blob_mount" "test_oracle_db_azure_blob_mount" {
	#Required
	oracle_db_azure_blob_mount_id = oci_dbmulticloud_oracle_db_azure_blob_mount.test_oracle_db_azure_blob_mount.id
}
```

## Argument Reference

The following arguments are supported:

* `oracle_db_azure_blob_mount_id` - (Required) The ID of the Azure Container Resource.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains Oracle DB Azure Blob Mount resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Oracle DB Azure Blob Mount name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID for the new Oracle DB Azure Blob Mount resource.
* `last_modification` - Description of the latest modification of the Oracle DB Azure Blob Mount Resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `mount_path` - Azure Container mount path.
* `oracle_db_azure_blob_container_id` - The OCID of the Oracle DB Azure Blob Container Resource.
* `oracle_db_azure_connector_id` - The OCID of the Oracle DB Azure Connector Resource.
* `state` - The current lifecycle state of the Azure Arc Agent Resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB Azure Blob Mount was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Oracle DB Azure Blob Mount was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

