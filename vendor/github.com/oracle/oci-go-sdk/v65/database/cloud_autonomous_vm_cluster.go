// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudAutonomousVmCluster Details of the cloud Autonomous VM cluster.
type CloudAutonomousVmCluster struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud Autonomous VM cluster.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the availability domain that the cloud Autonomous VM cluster is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the cloud Autonomous VM Cluster is associated with.
	// **Subnet Restrictions:**
	// - For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The current state of the cloud Autonomous VM cluster.
	LifecycleState CloudAutonomousVmClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the cloud Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
	CloudExadataInfrastructureId *string `mandatory:"true" json:"cloudExadataInfrastructureId"`

	// User defined description of the cloud Autonomous VM cluster.
	Description *string `mandatory:"false" json:"description"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance update history. This value is updated when a maintenance update starts.
	LastUpdateHistoryEntryId *string `mandatory:"false" json:"lastUpdateHistoryEntryId"`

	// The date and time that the cloud Autonomous VM cluster was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last date and time that the cloud Autonomous VM cluster was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time zone of the Cloud Autonomous VM Cluster.
	ClusterTimeZone *string `mandatory:"false" json:"clusterTimeZone"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The hostname for the cloud Autonomous VM cluster.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The domain name for the cloud Autonomous VM cluster.
	Domain *string `mandatory:"false" json:"domain"`

	// The model name of the Exadata hardware running the cloud Autonomous VM cluster.
	Shape *string `mandatory:"false" json:"shape"`

	// The number of database servers in the cloud VM cluster.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The total data storage allocated, in terabytes (TB).
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The total data storage allocated, in gigabytes (GB).
	DataStorageSizeInGBs *float64 `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The number of CPU cores on the cloud Autonomous VM cluster.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The number of CPU cores on the cloud Autonomous VM cluster. Only 1 decimal place is allowed for the fractional part.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The compute model of the Cloud Autonomous VM Cluster. ECPU compute model is the recommended model and OCPU compute model is legacy. See Compute Models in Autonomous Database on Dedicated Exadata #Infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbak) for more details.
	ComputeModel CloudAutonomousVmClusterComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

	// Enable mutual TLS(mTLS) authentication for database at time of provisioning a VMCluster. This is applicable to database TLS Certificates only. Default is TLS
	IsMtlsEnabledVmCluster *bool `mandatory:"false" json:"isMtlsEnabledVmCluster"`

	// The number of CPU cores enabled per VM cluster node.
	CpuCoreCountPerNode *int `mandatory:"false" json:"cpuCoreCountPerNode"`

	// The memory allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the
	// Autonomous Exadata Infrastructure level. When provisioning an Autonomous Database Serverless  (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, dataStorageSizeInTBs, adminPassword, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	LicenseModel CloudAutonomousVmClusterLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The SCAN Listenenr TLS port. Default is 2484.
	ScanListenerPortTls *int `mandatory:"false" json:"scanListenerPortTls"`

	// The SCAN Listener Non TLS port. Default is 1521.
	ScanListenerPortNonTls *int `mandatory:"false" json:"scanListenerPortNonTls"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Security Attributes for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// The date and time of Database SSL certificate expiration.
	TimeDatabaseSslCertificateExpires *common.SDKTime `mandatory:"false" json:"timeDatabaseSslCertificateExpires"`

	// The date and time of ORDS certificate expiration.
	TimeOrdsCertificateExpires *common.SDKTime `mandatory:"false" json:"timeOrdsCertificateExpires"`

	// CPU cores available for allocation to Autonomous Databases.
	AvailableCpus *float32 `mandatory:"false" json:"availableCpus"`

	// CPUs that continue to be included in the count of CPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available CPUs at its parent Autonomous VM Cluster level by restarting the Autonomous Container Database.
	ReclaimableCpus *float32 `mandatory:"false" json:"reclaimableCpus"`

	// The number of Autonomous Container Databases that can be created with the currently available local storage.
	AvailableContainerDatabases *int `mandatory:"false" json:"availableContainerDatabases"`

	// The total number of Autonomous Container Databases that can be created with the allocated local storage.
	TotalContainerDatabases *int `mandatory:"false" json:"totalContainerDatabases"`

	// The data disk group size available for Autonomous Databases, in TBs.
	AvailableAutonomousDataStorageSizeInTBs *float64 `mandatory:"false" json:"availableAutonomousDataStorageSizeInTBs"`

	// The data disk group size allocated for Autonomous Databases, in TBs.
	AutonomousDataStorageSizeInTBs *float64 `mandatory:"false" json:"autonomousDataStorageSizeInTBs"`

	// The local node storage allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The amount of memory (in GBs) enabled per OCPU or ECPU.
	MemoryPerOracleComputeUnitInGBs *int `mandatory:"false" json:"memoryPerOracleComputeUnitInGBs"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Db servers.
	DbServers []string `mandatory:"false" json:"dbServers"`

	// The percentage of total number of CPUs used in an Autonomous VM Cluster.
	CpuPercentage *float32 `mandatory:"false" json:"cpuPercentage"`

	// The percentage of the data storage used for the Autonomous Databases in an Autonomous VM Cluster.
	AutonomousDataStoragePercentage *float32 `mandatory:"false" json:"autonomousDataStoragePercentage"`

	// The number of CPUs provisioned in an Autonomous VM Cluster.
	ProvisionedCpus *float32 `mandatory:"false" json:"provisionedCpus"`

	// The total number of CPUs in an Autonomous VM Cluster.
	TotalCpus *float32 `mandatory:"false" json:"totalCpus"`

	// The total data disk group size for Autonomous Databases, in TBs.
	TotalAutonomousDataStorageInTBs *float32 `mandatory:"false" json:"totalAutonomousDataStorageInTBs"`

	// The number of CPUs reserved in an Autonomous VM Cluster.
	ReservedCpus *float32 `mandatory:"false" json:"reservedCpus"`

	// The number of provisionable Autonomous Container Databases in an Autonomous VM Cluster.
	ProvisionableAutonomousContainerDatabases *int `mandatory:"false" json:"provisionableAutonomousContainerDatabases"`

	// The number of provisioned Autonomous Container Databases in an Autonomous VM Cluster.
	ProvisionedAutonomousContainerDatabases *int `mandatory:"false" json:"provisionedAutonomousContainerDatabases"`

	// The number of non-provisionable Autonomous Container Databases in an Autonomous VM Cluster.
	NonProvisionableAutonomousContainerDatabases *int `mandatory:"false" json:"nonProvisionableAutonomousContainerDatabases"`

	// The lowest value to which exadataStorage (in TBs) can be scaled down.
	ExadataStorageInTBsLowestScaledValue *float64 `mandatory:"false" json:"exadataStorageInTBsLowestScaledValue"`

	// The lowest value to which ocpus can be scaled down.
	OcpusLowestScaledValue *int `mandatory:"false" json:"ocpusLowestScaledValue"`

	// The lowest value to which maximum number of ACDs can be scaled down.
	MaxAcdsLowestScaledValue *int `mandatory:"false" json:"maxAcdsLowestScaledValue"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`
}

func (m CloudAutonomousVmCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudAutonomousVmCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudAutonomousVmClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudAutonomousVmClusterLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCloudAutonomousVmClusterComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetCloudAutonomousVmClusterComputeModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudAutonomousVmClusterLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCloudAutonomousVmClusterLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudAutonomousVmClusterLifecycleStateEnum Enum with underlying type: string
type CloudAutonomousVmClusterLifecycleStateEnum string

// Set of constants representing the allowable values for CloudAutonomousVmClusterLifecycleStateEnum
const (
	CloudAutonomousVmClusterLifecycleStateProvisioning          CloudAutonomousVmClusterLifecycleStateEnum = "PROVISIONING"
	CloudAutonomousVmClusterLifecycleStateAvailable             CloudAutonomousVmClusterLifecycleStateEnum = "AVAILABLE"
	CloudAutonomousVmClusterLifecycleStateUpdating              CloudAutonomousVmClusterLifecycleStateEnum = "UPDATING"
	CloudAutonomousVmClusterLifecycleStateTerminating           CloudAutonomousVmClusterLifecycleStateEnum = "TERMINATING"
	CloudAutonomousVmClusterLifecycleStateTerminated            CloudAutonomousVmClusterLifecycleStateEnum = "TERMINATED"
	CloudAutonomousVmClusterLifecycleStateFailed                CloudAutonomousVmClusterLifecycleStateEnum = "FAILED"
	CloudAutonomousVmClusterLifecycleStateMaintenanceInProgress CloudAutonomousVmClusterLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingCloudAutonomousVmClusterLifecycleStateEnum = map[string]CloudAutonomousVmClusterLifecycleStateEnum{
	"PROVISIONING":            CloudAutonomousVmClusterLifecycleStateProvisioning,
	"AVAILABLE":               CloudAutonomousVmClusterLifecycleStateAvailable,
	"UPDATING":                CloudAutonomousVmClusterLifecycleStateUpdating,
	"TERMINATING":             CloudAutonomousVmClusterLifecycleStateTerminating,
	"TERMINATED":              CloudAutonomousVmClusterLifecycleStateTerminated,
	"FAILED":                  CloudAutonomousVmClusterLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": CloudAutonomousVmClusterLifecycleStateMaintenanceInProgress,
}

var mappingCloudAutonomousVmClusterLifecycleStateEnumLowerCase = map[string]CloudAutonomousVmClusterLifecycleStateEnum{
	"provisioning":            CloudAutonomousVmClusterLifecycleStateProvisioning,
	"available":               CloudAutonomousVmClusterLifecycleStateAvailable,
	"updating":                CloudAutonomousVmClusterLifecycleStateUpdating,
	"terminating":             CloudAutonomousVmClusterLifecycleStateTerminating,
	"terminated":              CloudAutonomousVmClusterLifecycleStateTerminated,
	"failed":                  CloudAutonomousVmClusterLifecycleStateFailed,
	"maintenance_in_progress": CloudAutonomousVmClusterLifecycleStateMaintenanceInProgress,
}

// GetCloudAutonomousVmClusterLifecycleStateEnumValues Enumerates the set of values for CloudAutonomousVmClusterLifecycleStateEnum
func GetCloudAutonomousVmClusterLifecycleStateEnumValues() []CloudAutonomousVmClusterLifecycleStateEnum {
	values := make([]CloudAutonomousVmClusterLifecycleStateEnum, 0)
	for _, v := range mappingCloudAutonomousVmClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudAutonomousVmClusterLifecycleStateEnumStringValues Enumerates the set of values in String for CloudAutonomousVmClusterLifecycleStateEnum
func GetCloudAutonomousVmClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"MAINTENANCE_IN_PROGRESS",
	}
}

// GetMappingCloudAutonomousVmClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudAutonomousVmClusterLifecycleStateEnum(val string) (CloudAutonomousVmClusterLifecycleStateEnum, bool) {
	enum, ok := mappingCloudAutonomousVmClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudAutonomousVmClusterComputeModelEnum Enum with underlying type: string
type CloudAutonomousVmClusterComputeModelEnum string

// Set of constants representing the allowable values for CloudAutonomousVmClusterComputeModelEnum
const (
	CloudAutonomousVmClusterComputeModelEcpu CloudAutonomousVmClusterComputeModelEnum = "ECPU"
	CloudAutonomousVmClusterComputeModelOcpu CloudAutonomousVmClusterComputeModelEnum = "OCPU"
)

var mappingCloudAutonomousVmClusterComputeModelEnum = map[string]CloudAutonomousVmClusterComputeModelEnum{
	"ECPU": CloudAutonomousVmClusterComputeModelEcpu,
	"OCPU": CloudAutonomousVmClusterComputeModelOcpu,
}

var mappingCloudAutonomousVmClusterComputeModelEnumLowerCase = map[string]CloudAutonomousVmClusterComputeModelEnum{
	"ecpu": CloudAutonomousVmClusterComputeModelEcpu,
	"ocpu": CloudAutonomousVmClusterComputeModelOcpu,
}

// GetCloudAutonomousVmClusterComputeModelEnumValues Enumerates the set of values for CloudAutonomousVmClusterComputeModelEnum
func GetCloudAutonomousVmClusterComputeModelEnumValues() []CloudAutonomousVmClusterComputeModelEnum {
	values := make([]CloudAutonomousVmClusterComputeModelEnum, 0)
	for _, v := range mappingCloudAutonomousVmClusterComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudAutonomousVmClusterComputeModelEnumStringValues Enumerates the set of values in String for CloudAutonomousVmClusterComputeModelEnum
func GetCloudAutonomousVmClusterComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingCloudAutonomousVmClusterComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudAutonomousVmClusterComputeModelEnum(val string) (CloudAutonomousVmClusterComputeModelEnum, bool) {
	enum, ok := mappingCloudAutonomousVmClusterComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CloudAutonomousVmClusterLicenseModelEnum Enum with underlying type: string
type CloudAutonomousVmClusterLicenseModelEnum string

// Set of constants representing the allowable values for CloudAutonomousVmClusterLicenseModelEnum
const (
	CloudAutonomousVmClusterLicenseModelLicenseIncluded     CloudAutonomousVmClusterLicenseModelEnum = "LICENSE_INCLUDED"
	CloudAutonomousVmClusterLicenseModelBringYourOwnLicense CloudAutonomousVmClusterLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCloudAutonomousVmClusterLicenseModelEnum = map[string]CloudAutonomousVmClusterLicenseModelEnum{
	"LICENSE_INCLUDED":       CloudAutonomousVmClusterLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CloudAutonomousVmClusterLicenseModelBringYourOwnLicense,
}

var mappingCloudAutonomousVmClusterLicenseModelEnumLowerCase = map[string]CloudAutonomousVmClusterLicenseModelEnum{
	"license_included":       CloudAutonomousVmClusterLicenseModelLicenseIncluded,
	"bring_your_own_license": CloudAutonomousVmClusterLicenseModelBringYourOwnLicense,
}

// GetCloudAutonomousVmClusterLicenseModelEnumValues Enumerates the set of values for CloudAutonomousVmClusterLicenseModelEnum
func GetCloudAutonomousVmClusterLicenseModelEnumValues() []CloudAutonomousVmClusterLicenseModelEnum {
	values := make([]CloudAutonomousVmClusterLicenseModelEnum, 0)
	for _, v := range mappingCloudAutonomousVmClusterLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudAutonomousVmClusterLicenseModelEnumStringValues Enumerates the set of values in String for CloudAutonomousVmClusterLicenseModelEnum
func GetCloudAutonomousVmClusterLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCloudAutonomousVmClusterLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudAutonomousVmClusterLicenseModelEnum(val string) (CloudAutonomousVmClusterLicenseModelEnum, bool) {
	enum, ok := mappingCloudAutonomousVmClusterLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
