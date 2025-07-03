// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_vcn" "test_vcn" {
  compartment_id = var.compartment_ocid
  cidr_block     = "10.1.0.0/16"
  display_name   = "TestVcn"
  dns_label      = "examplevcn"
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.compartment_ocid
  ad_number      = 1
}

resource "oci_core_security_list" "exadata_shapes_security_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "ExadataSecurityList"

  ingress_security_rules {
    source   = "10.1.22.0/24"
    protocol = "6"
  }

  ingress_security_rules {
    source   = "10.1.22.0/24"
    protocol = "1"
  }

  egress_security_rules {
    destination = "10.1.22.0/24"
    protocol    = "6"
  }

  egress_security_rules {
    destination = "10.1.22.0/24"
    protocol    = "1"
  }
}

resource "oci_core_subnet" "exadata_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.22.0/24"
  display_name        = "TestExadataSubnet"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  route_table_id      = oci_core_vcn.test_vcn.default_route_table_id
  dhcp_options_id     = oci_core_vcn.test_vcn.default_dhcp_options_id
  security_list_ids   = [oci_core_vcn.test_vcn.default_security_list_id, oci_core_security_list.exadata_shapes_security_list.id]
  dns_label           = "subnetexadata"
}

resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TFATPD"
  shape               = var.cloud_exadata_infrastructure_shape

  #Optional
  compute_count = var.cloud_exadata_infrastructure_compute_count
  storage_count = var.cloud_exadata_infrastructure_storage_count
}

resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure_primary" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "ATPDPRIMARY"
  shape               = var.cloud_exadata_infrastructure_shape

  #Optional
  compute_count = var.cloud_exadata_infrastructure_compute_count
  storage_count = var.cloud_exadata_infrastructure_storage_count
}

resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure_standby" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "ATPDSTANDBY"
  shape               = var.cloud_exadata_infrastructure_shape

  #Optional
  compute_count = var.cloud_exadata_infrastructure_compute_count
  storage_count = var.cloud_exadata_infrastructure_storage_count
}


resource "oci_database_cloud_autonomous_vm_cluster" "test_cloud_autonomous_vm_cluster" {
  cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
  compartment_id                  = var.compartment_ocid
  display_name                    = "TestCloudAutonomousVmCluster"
  freeform_tags                   = var.autonomous_database_freeform_tags
  license_model                   = "LICENSE_INCLUDED"
  subnet_id                       = oci_core_subnet.exadata_subnet.id
#  #Optional
#  autonomous_data_storage_size_in_tbs   = 2
#  memory_per_oracle_compute_unit_in_gbs = 5
#  cpu_core_count_per_node               = 40
#  total_container_databases             = 1
  compute_model                   = "ECPU"

  //To ignore changes to autonomous_data_storage_size_in_tbs and db_servers
  lifecycle {
    ignore_changes = [
      autonomous_data_storage_size_in_tbs,
      db_servers,
    ]
  }
}

resource "oci_database_cloud_autonomous_vm_cluster" "test_cloud_autonomous_vm_cluster_primary" {
  cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure_primary.id
  compartment_id                  = var.compartment_ocid
  display_name                    = "TestCloudAutonomousVmClusterPrimary"
  freeform_tags                   = var.autonomous_database_freeform_tags
  license_model                   = "LICENSE_INCLUDED"
  subnet_id                       = oci_core_subnet.exadata_subnet.id
  
  #Optional
#  autonomous_data_storage_size_in_tbs   = 2
#  memory_per_oracle_compute_unit_in_gbs = 5
#  cpu_core_count_per_node               = 40
#  total_container_databases             = 1
  compute_model                   = "ECPU"

  //To ignore changes to autonomous_data_storage_size_in_tbs and db_servers
  lifecycle {
    ignore_changes = [
      autonomous_data_storage_size_in_tbs,
      db_servers,
    ]
  }
}

resource "oci_database_cloud_autonomous_vm_cluster" "test_cloud_autonomous_vm_cluster_standby" {
  cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure_standby.id
  compartment_id                  = var.compartment_ocid
  display_name                    = "TestCloudAutonomousVmClusterStandby"
  freeform_tags                   = var.autonomous_database_freeform_tags
  license_model                   = "LICENSE_INCLUDED"
  subnet_id                       = oci_core_subnet.exadata_subnet.id
#  #Optional
#  autonomous_data_storage_size_in_tbs   = 2
#  memory_per_oracle_compute_unit_in_gbs = 5
#  cpu_core_count_per_node               = 40
#  total_container_databases             = 1
  compute_model                   = "ECPU"

  //To ignore changes to autonomous_data_storage_size_in_tbs and db_servers
  lifecycle {
    ignore_changes = [
      autonomous_data_storage_size_in_tbs,
      db_servers,
    ]
  }
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

data "oci_database_cloud_autonomous_vm_cluster" "test_cloud_autonomous_vm_cluster" {
  cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
}

data "oci_database_cloud_autonomous_vm_clusters" "test_cloud_autonomous_vm_clusters" {
    compartment_id = var.compartment_ocid
}