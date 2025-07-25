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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HttpEndpointAuthScopeConfig Subset of AuthScopeConfig allowed for HTTP Endpoint Tool.
type HttpEndpointAuthScopeConfig interface {
}

type httpendpointauthscopeconfig struct {
	JsonData                        []byte
	HttpEndpointAuthScopeConfigType string `json:"httpEndpointAuthScopeConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *httpendpointauthscopeconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhttpendpointauthscopeconfig httpendpointauthscopeconfig
	s := struct {
		Model Unmarshalerhttpendpointauthscopeconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.HttpEndpointAuthScopeConfigType = s.Model.HttpEndpointAuthScopeConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *httpendpointauthscopeconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.HttpEndpointAuthScopeConfigType {
	case "HTTP_ENDPOINT_BEARER_AUTH_SCOPE_CONFIG":
		mm := HttpEndpointBearerAuthScopeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG":
		mm := HttpEndpointNoAuthScopeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG":
		mm := HttpEndpointBasicAuthScopeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_ENDPOINT_OCI_AUTH_SCOPE_CONFIG":
		mm := HttpEndpointOciAuthScopeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG":
		mm := HttpEndpointIdcsAuthScopeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG":
		mm := HttpEndpointApiKeyAuthScopeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for HttpEndpointAuthScopeConfig: %s.", m.HttpEndpointAuthScopeConfigType)
		return *m, nil
	}
}

func (m httpendpointauthscopeconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m httpendpointauthscopeconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum Enum with underlying type: string
type HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum string

// Set of constants representing the allowable values for HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum
const (
	HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeNoAuthScopeConfig     HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum = "HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG"
	HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeBasicAuthScopeConfig  HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum = "HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG"
	HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeBearerAuthScopeConfig HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum = "HTTP_ENDPOINT_BEARER_AUTH_SCOPE_CONFIG"
	HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeApiKeyAuthScopeConfig HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum = "HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG"
	HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeIdcsAuthScopeConfig   HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum = "HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG"
	HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeOciAuthScopeConfig    HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum = "HTTP_ENDPOINT_OCI_AUTH_SCOPE_CONFIG"
)

var mappingHttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum = map[string]HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum{
	"HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG":      HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeNoAuthScopeConfig,
	"HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG":   HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeBasicAuthScopeConfig,
	"HTTP_ENDPOINT_BEARER_AUTH_SCOPE_CONFIG":  HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeBearerAuthScopeConfig,
	"HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG": HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeApiKeyAuthScopeConfig,
	"HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG":    HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeIdcsAuthScopeConfig,
	"HTTP_ENDPOINT_OCI_AUTH_SCOPE_CONFIG":     HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeOciAuthScopeConfig,
}

var mappingHttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnumLowerCase = map[string]HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum{
	"http_endpoint_no_auth_scope_config":      HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeNoAuthScopeConfig,
	"http_endpoint_basic_auth_scope_config":   HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeBasicAuthScopeConfig,
	"http_endpoint_bearer_auth_scope_config":  HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeBearerAuthScopeConfig,
	"http_endpoint_api_key_auth_scope_config": HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeApiKeyAuthScopeConfig,
	"http_endpoint_idcs_auth_scope_config":    HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeIdcsAuthScopeConfig,
	"http_endpoint_oci_auth_scope_config":     HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeOciAuthScopeConfig,
}

// GetHttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnumValues Enumerates the set of values for HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum
func GetHttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnumValues() []HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum {
	values := make([]HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum, 0)
	for _, v := range mappingHttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnumStringValues Enumerates the set of values in String for HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum
func GetHttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnumStringValues() []string {
	return []string{
		"HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG",
		"HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG",
		"HTTP_ENDPOINT_BEARER_AUTH_SCOPE_CONFIG",
		"HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG",
		"HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG",
		"HTTP_ENDPOINT_OCI_AUTH_SCOPE_CONFIG",
	}
}

// GetMappingHttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum(val string) (HttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnum, bool) {
	enum, ok := mappingHttpEndpointAuthScopeConfigHttpEndpointAuthScopeConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
