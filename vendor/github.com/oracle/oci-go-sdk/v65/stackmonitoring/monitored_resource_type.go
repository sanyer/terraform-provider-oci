// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoredResourceType The response object for create monitored resource type and get monitored resource type operations.
type MonitoredResourceType struct {

	// Monitored resource type identifier OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// A unique monitored resource type name. The name must be unique across tenancy.
	// Name can not be changed.
	Name *string `mandatory:"true" json:"name"`

	// Compartment Identifier OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Monitored resource type display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A friendly description.
	Description *string `mandatory:"false" json:"description"`

	// Metric namespace for resource type.
	MetricNamespace *string `mandatory:"false" json:"metricNamespace"`

	// Tenancy Identifier OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	AvailabilityMetricsConfig *AvailabilityMetricsDetails `mandatory:"false" json:"availabilityMetricsConfig"`

	HandlerConfig *AgentExtensionHandlerConfiguration `mandatory:"false" json:"handlerConfig"`

	// If boolean flag is true, then the resource type cannot be modified or deleted.
	IsSystemDefined *bool `mandatory:"false" json:"isSystemDefined"`

	// Lifecycle state of the monitored resource type.
	LifecycleState ResourceTypeLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Source type to indicate if the resource is stack monitoring discovered, OCI native resource, etc.
	SourceType SourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// Resource Category to indicate the kind of resource type.
	ResourceCategory ResourceCategoryEnum `mandatory:"false" json:"resourceCategory,omitempty"`

	// The date and time when the monitored resource type was created, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the monitored resource was updated, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Key/Value pair for additional namespaces used by stack monitoring services for SYSTEM (SMB) resource types.
	AdditionalNamespaceMap map[string]string `mandatory:"false" json:"additionalNamespaceMap"`

	Metadata ResourceTypeMetadataDetails `mandatory:"false" json:"metadata"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MonitoredResourceType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoredResourceType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceTypeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResourceTypeLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceCategoryEnum(string(m.ResourceCategory)); !ok && m.ResourceCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceCategory: %s. Supported values are: %s.", m.ResourceCategory, strings.Join(GetResourceCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MonitoredResourceType) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName               *string                             `json:"displayName"`
		Description               *string                             `json:"description"`
		MetricNamespace           *string                             `json:"metricNamespace"`
		TenancyId                 *string                             `json:"tenancyId"`
		AvailabilityMetricsConfig *AvailabilityMetricsDetails         `json:"availabilityMetricsConfig"`
		HandlerConfig             *AgentExtensionHandlerConfiguration `json:"handlerConfig"`
		IsSystemDefined           *bool                               `json:"isSystemDefined"`
		LifecycleState            ResourceTypeLifecycleStateEnum      `json:"lifecycleState"`
		SourceType                SourceTypeEnum                      `json:"sourceType"`
		ResourceCategory          ResourceCategoryEnum                `json:"resourceCategory"`
		TimeCreated               *common.SDKTime                     `json:"timeCreated"`
		TimeUpdated               *common.SDKTime                     `json:"timeUpdated"`
		AdditionalNamespaceMap    map[string]string                   `json:"additionalNamespaceMap"`
		Metadata                  resourcetypemetadatadetails         `json:"metadata"`
		FreeformTags              map[string]string                   `json:"freeformTags"`
		DefinedTags               map[string]map[string]interface{}   `json:"definedTags"`
		SystemTags                map[string]map[string]interface{}   `json:"systemTags"`
		Id                        *string                             `json:"id"`
		Name                      *string                             `json:"name"`
		CompartmentId             *string                             `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.MetricNamespace = model.MetricNamespace

	m.TenancyId = model.TenancyId

	m.AvailabilityMetricsConfig = model.AvailabilityMetricsConfig

	m.HandlerConfig = model.HandlerConfig

	m.IsSystemDefined = model.IsSystemDefined

	m.LifecycleState = model.LifecycleState

	m.SourceType = model.SourceType

	m.ResourceCategory = model.ResourceCategory

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.AdditionalNamespaceMap = model.AdditionalNamespaceMap

	nn, e = model.Metadata.UnmarshalPolymorphicJSON(model.Metadata.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Metadata = nn.(ResourceTypeMetadataDetails)
	} else {
		m.Metadata = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	return
}
