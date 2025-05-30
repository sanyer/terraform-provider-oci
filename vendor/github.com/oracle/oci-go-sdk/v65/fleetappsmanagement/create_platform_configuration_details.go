// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePlatformConfigurationDetails The information about new PlatformConfiguration.
type CreatePlatformConfigurationDetails struct {

	// Compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	ConfigCategoryDetails ConfigCategoryDetails `mandatory:"true" json:"configCategoryDetails"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`
}

func (m CreatePlatformConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePlatformConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreatePlatformConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description           *string               `json:"description"`
		CompartmentId         *string               `json:"compartmentId"`
		DisplayName           *string               `json:"displayName"`
		ConfigCategoryDetails configcategorydetails `json:"configCategoryDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	nn, e = model.ConfigCategoryDetails.UnmarshalPolymorphicJSON(model.ConfigCategoryDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConfigCategoryDetails = nn.(ConfigCategoryDetails)
	} else {
		m.ConfigCategoryDetails = nil
	}

	return
}
