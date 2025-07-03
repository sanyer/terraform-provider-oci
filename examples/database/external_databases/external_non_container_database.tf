// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// Create an External Non-Container Database resource
resource "oci_database_external_non_container_database" "test_external_non_container_database" {
  compartment_id = var.compartment_ocid
  display_name = var.external_non_container_database_display_name

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}

  freeform_tags = {
    "Department" = "Finance"
  }
}

// Create a connector using credential name for the External Non-Container Database resource
resource "oci_database_external_database_connector" "test_external_non_container_database_connector" {
    connection_credentials {
        credential_type = var.credential_type
        credential_name = var.credential_name
        password = var.password
        role = var.role
        username = var.username
    }
    connection_string {
        hostname = var.hostname
        port = var.port
        protocol = var.protocol
        service = var.service
    }
    connector_agent_id = var.connector_agent_id
    display_name = var.external_database_connector_display_name
    external_database_id = oci_database_external_non_container_database.test_external_non_container_database.id
    connector_type = var.connector_type
    #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
    #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}

#    freeform_tags = {
#      "Department" = "Finance"
#    }
}

// Enable Database Management

resource "oci_database_external_non_container_database_management" "test_enable_external_non_container_database_management" {
    external_non_container_database_id = oci_database_external_non_container_database.test_external_non_container_database.id
    external_database_connector_id = oci_database_external_database_connector.test_external_non_container_database_connector.id
    license_model = var.license_model
    enable_management = true
}

// Enable Operations Insights
resource "oci_database_external_non_container_database_operations_insights_management" "test_enable_external_non_container_database_operations_insights_management" {
  external_non_container_database_id = oci_database_external_non_container_database.test_external_non_container_database.id
  external_database_connector_id = oci_database_external_database_connector.test_external_non_container_database_connector.id
  enable_operations_insights = true
}

//Commenting out this code block to unblock the failure in backward compatibility test
//Disable Database Management
/*resource "oci_database_external_non_container_database_management" "test_disable_external_non_container_database_management" {
    depends_on = [oci_database_external_non_container_database_management.test_enable_external_non_container_database_management]
    external_non_container_database_id = oci_database_external_non_container_database.test_external_non_container_database.id
    external_database_connector_id = oci_database_external_database_connector.test_external_non_container_database_connector.id
    enable_management = false
}*/


data "oci_database_external_non_container_database" "test_external_non_container_database" {
	#Required
	external_non_container_database_id = oci_database_external_non_container_database.test_external_non_container_database.id
}

data "oci_database_external_non_container_databases" "test_external_non_container_databases" {
	#Required
	compartment_id = var.compartment_ocid

	#Optional
	display_name = var.external_non_container_database_display_name
	state = var.external_database_state
}