// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VolumeCostEstimation Cost estimation for volume
type VolumeCostEstimation struct {

	// Gigabyte storage capacity
	CapacityGb *float32 `mandatory:"true" json:"capacityGb"`

	// Gigabyte storage capacity per month.
	TotalGbPerMonth *float32 `mandatory:"true" json:"totalGbPerMonth"`

	// Volume description
	Description *string `mandatory:"false" json:"description"`

	// Gigabyte storage capacity per month by subscription
	TotalGbPerMonthBySubscription *float32 `mandatory:"false" json:"totalGbPerMonthBySubscription"`
}

func (m VolumeCostEstimation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VolumeCostEstimation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
