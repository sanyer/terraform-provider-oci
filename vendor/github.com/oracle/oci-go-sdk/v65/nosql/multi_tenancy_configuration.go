// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MultiTenancyConfiguration The service-level configuration for a multi-tenancy environment.
type MultiTenancyConfiguration struct {
}

func (m MultiTenancyConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MultiTenancyConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MultiTenancyConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMultiTenancyConfiguration MultiTenancyConfiguration
	s := struct {
		DiscriminatorParam string `json:"environment"`
		MarshalTypeMultiTenancyConfiguration
	}{
		"MULTI_TENANCY",
		(MarshalTypeMultiTenancyConfiguration)(m),
	}

	return json.Marshal(&s)
}
