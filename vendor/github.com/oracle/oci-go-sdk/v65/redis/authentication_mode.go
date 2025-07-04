// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuthenticationMode These are the Authentication details of an OCI cache user.
type AuthenticationMode interface {
}

type authenticationmode struct {
	JsonData           []byte
	AuthenticationType string `json:"authenticationType"`
}

// UnmarshalJSON unmarshals json
func (m *authenticationmode) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerauthenticationmode authenticationmode
	s := struct {
		Model Unmarshalerauthenticationmode
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AuthenticationType = s.Model.AuthenticationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *authenticationmode) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AuthenticationType {
	case "IAM":
		mm := IamAuthenticationMode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PASSWORD":
		mm := PasswordAuthenticationMode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AuthenticationMode: %s.", m.AuthenticationType)
		return *m, nil
	}
}

func (m authenticationmode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m authenticationmode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthenticationModeAuthenticationTypeEnum Enum with underlying type: string
type AuthenticationModeAuthenticationTypeEnum string

// Set of constants representing the allowable values for AuthenticationModeAuthenticationTypeEnum
const (
	AuthenticationModeAuthenticationTypeIam      AuthenticationModeAuthenticationTypeEnum = "IAM"
	AuthenticationModeAuthenticationTypePassword AuthenticationModeAuthenticationTypeEnum = "PASSWORD"
)

var mappingAuthenticationModeAuthenticationTypeEnum = map[string]AuthenticationModeAuthenticationTypeEnum{
	"IAM":      AuthenticationModeAuthenticationTypeIam,
	"PASSWORD": AuthenticationModeAuthenticationTypePassword,
}

var mappingAuthenticationModeAuthenticationTypeEnumLowerCase = map[string]AuthenticationModeAuthenticationTypeEnum{
	"iam":      AuthenticationModeAuthenticationTypeIam,
	"password": AuthenticationModeAuthenticationTypePassword,
}

// GetAuthenticationModeAuthenticationTypeEnumValues Enumerates the set of values for AuthenticationModeAuthenticationTypeEnum
func GetAuthenticationModeAuthenticationTypeEnumValues() []AuthenticationModeAuthenticationTypeEnum {
	values := make([]AuthenticationModeAuthenticationTypeEnum, 0)
	for _, v := range mappingAuthenticationModeAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationModeAuthenticationTypeEnumStringValues Enumerates the set of values in String for AuthenticationModeAuthenticationTypeEnum
func GetAuthenticationModeAuthenticationTypeEnumStringValues() []string {
	return []string{
		"IAM",
		"PASSWORD",
	}
}

// GetMappingAuthenticationModeAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationModeAuthenticationTypeEnum(val string) (AuthenticationModeAuthenticationTypeEnum, bool) {
	enum, ok := mappingAuthenticationModeAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
