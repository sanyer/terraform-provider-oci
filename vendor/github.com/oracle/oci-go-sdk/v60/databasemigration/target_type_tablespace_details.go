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

// TargetTypeTablespaceDetails Migration tablespace settings.
type TargetTypeTablespaceDetails interface {
}

type targettypetablespacedetails struct {
	JsonData   []byte
	TargetType string `json:"targetType"`
}

// UnmarshalJSON unmarshals json
func (m *targettypetablespacedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargettypetablespacedetails targettypetablespacedetails
	s := struct {
		Model Unmarshalertargettypetablespacedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TargetType = s.Model.TargetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targettypetablespacedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TargetType {
	case "ADB_D_AUTOCREATE":
		mm := AdbDedicatedAutoCreateTablespaceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ADB_S_REMAP":
		mm := AdbServerlesTablespaceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ADB_D_REMAP":
		mm := AdbDedicatedRemapTargetTablespaceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NON_ADB_REMAP":
		mm := NonAdbRemapTablespaceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NON_ADB_AUTOCREATE":
		mm := NonAdbAutoCreateTablespaceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m targettypetablespacedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m targettypetablespacedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
