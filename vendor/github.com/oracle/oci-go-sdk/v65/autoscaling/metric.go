// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// Use the Autoscaling API to dynamically scale compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.oracle.com/iaas/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Compute (https://docs.oracle.com/iaas/Content/Compute/home.htm).
//

package autoscaling

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Metric Metric and threshold details for triggering an autoscaling action based on CPU or memory utilization.
type Metric struct {
	Threshold *Threshold `mandatory:"true" json:"threshold"`

	// The period of time that the condition defined in the alarm must persist before the alarm state
	// changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the
	// alarm must persist in breaching the condition for five minutes before the alarm updates its
	// state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five
	// minutes before the alarm updates its state to "OK."
	// The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H`
	// for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M.
	PendingDuration *string `mandatory:"false" json:"pendingDuration"`

	MetricType MetricMetricTypeEnum `mandatory:"true" json:"metricType"`
}

// GetPendingDuration returns PendingDuration
func (m Metric) GetPendingDuration() *string {
	return m.PendingDuration
}

func (m Metric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Metric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMetricMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetMetricMetricTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Metric) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMetric Metric
	s := struct {
		DiscriminatorParam string `json:"metricSource"`
		MarshalTypeMetric
	}{
		"COMPUTE_AGENT",
		(MarshalTypeMetric)(m),
	}

	return json.Marshal(&s)
}

// MetricMetricTypeEnum Enum with underlying type: string
type MetricMetricTypeEnum string

// Set of constants representing the allowable values for MetricMetricTypeEnum
const (
	MetricMetricTypeCpuUtilization    MetricMetricTypeEnum = "CPU_UTILIZATION"
	MetricMetricTypeMemoryUtilization MetricMetricTypeEnum = "MEMORY_UTILIZATION"
)

var mappingMetricMetricTypeEnum = map[string]MetricMetricTypeEnum{
	"CPU_UTILIZATION":    MetricMetricTypeCpuUtilization,
	"MEMORY_UTILIZATION": MetricMetricTypeMemoryUtilization,
}

var mappingMetricMetricTypeEnumLowerCase = map[string]MetricMetricTypeEnum{
	"cpu_utilization":    MetricMetricTypeCpuUtilization,
	"memory_utilization": MetricMetricTypeMemoryUtilization,
}

// GetMetricMetricTypeEnumValues Enumerates the set of values for MetricMetricTypeEnum
func GetMetricMetricTypeEnumValues() []MetricMetricTypeEnum {
	values := make([]MetricMetricTypeEnum, 0)
	for _, v := range mappingMetricMetricTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricMetricTypeEnumStringValues Enumerates the set of values in String for MetricMetricTypeEnum
func GetMetricMetricTypeEnumStringValues() []string {
	return []string{
		"CPU_UTILIZATION",
		"MEMORY_UTILIZATION",
	}
}

// GetMappingMetricMetricTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricMetricTypeEnum(val string) (MetricMetricTypeEnum, bool) {
	enum, ok := mappingMetricMetricTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
