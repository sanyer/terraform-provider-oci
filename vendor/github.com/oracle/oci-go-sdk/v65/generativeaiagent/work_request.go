// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents API to create and manage agents, knowledge bases, data sources, endpoints, data ingestion jobs, and work requests.
// For creating and managing client chat sessions see the /EN/generative-ai-agents-client/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequest An asynchronous work request. Work requests help you monitor long-running operations. When you start a long-running operation,
// the service creates a work request. A work request is an activity log that lets you track each step in the operation's
// progress. Each work request has an OCID that lets you interact with it programmatically and use it for automation.
type WorkRequest struct {

	// The asynchronous operation tracked by this work request.
	OperationType OperationTypeEnum `mandatory:"true" json:"operationType"`

	// The status of the work request.
	Status OperationStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources that are affected by the work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Shows the progress of the operation tracked by the work request, as a percentage of the total work
	// that must be performed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the work request was created, in the format defined by
	// RFC 3339 (https://tools.ietf.org/html/rfc3339).
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the work request was started, in the format defined by
	// RFC 3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the work request was finished, in the format defined by
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The date and time the work request was updated, in the format defined by
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOperationStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOperationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
