// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// CreateAdbDedicatedAutoCreateTablespaceDetails Migration tablespace settings valid for ADB-D target type using auto create feature.
type CreateAdbDedicatedAutoCreateTablespaceDetails struct {

	// True to auto-create tablespace in the target Database.
	IsAutoCreate *bool `mandatory:"false" json:"isAutoCreate"`

	// True set tablespace to big file.
	IsBigFile *bool `mandatory:"false" json:"isBigFile"`

	// Size of extend in MB. Can only be specified if 'isBigFile' property is set to true.
	ExtendSizeInMBs *int `mandatory:"false" json:"extendSizeInMBs"`
}

func (m CreateAdbDedicatedAutoCreateTablespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAdbDedicatedAutoCreateTablespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAdbDedicatedAutoCreateTablespaceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAdbDedicatedAutoCreateTablespaceDetails CreateAdbDedicatedAutoCreateTablespaceDetails
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeCreateAdbDedicatedAutoCreateTablespaceDetails
	}{
		"ADB_D_AUTOCREATE",
		(MarshalTypeCreateAdbDedicatedAutoCreateTablespaceDetails)(m),
	}

	return json.Marshal(&s)
}
