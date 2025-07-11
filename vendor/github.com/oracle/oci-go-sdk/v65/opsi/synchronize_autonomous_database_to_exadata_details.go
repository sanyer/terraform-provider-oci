// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SynchronizeAutonomousDatabaseToExadataDetails The details of onboarded autonomous database need to synchroized with infracture information.
type SynchronizeAutonomousDatabaseToExadataDetails struct {

	// Source of the database entity. Currently only AUTONOMOUS_DATABASE source is supported.
	EntitySource DatabaseEntitySourceAllEnum `mandatory:"true" json:"entitySource"`
}

func (m SynchronizeAutonomousDatabaseToExadataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SynchronizeAutonomousDatabaseToExadataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseEntitySourceAllEnum(string(m.EntitySource)); !ok && m.EntitySource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntitySource: %s. Supported values are: %s.", m.EntitySource, strings.Join(GetDatabaseEntitySourceAllEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
