// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateCloudDbSystemMacsConnectorDetails The details for updating the cloud Management Agent Cloud Service (MACS) (https://docs.oracle.com/iaas/management-agents/index.html)
// connector used to connect to a cloud DB system component.
type UpdateCloudDbSystemMacsConnectorDetails struct {
	ConnectionInfo CloudDbSystemConnectionInfo `mandatory:"false" json:"connectionInfo"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCloudDbSystemMacsConnectorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCloudDbSystemMacsConnectorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateCloudDbSystemMacsConnectorDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateCloudDbSystemMacsConnectorDetails UpdateCloudDbSystemMacsConnectorDetails
	s := struct {
		DiscriminatorParam string `json:"connectorType"`
		MarshalTypeUpdateCloudDbSystemMacsConnectorDetails
	}{
		"MACS",
		(MarshalTypeUpdateCloudDbSystemMacsConnectorDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateCloudDbSystemMacsConnectorDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectionInfo clouddbsystemconnectioninfo       `json:"connectionInfo"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ConnectionInfo.UnmarshalPolymorphicJSON(model.ConnectionInfo.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectionInfo = nn.(CloudDbSystemConnectionInfo)
	} else {
		m.ConnectionInfo = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
