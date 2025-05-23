// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Channel Properties of a Channel.
type Channel interface {

	// Unique immutable identifier that was assigned when the Channel was created.
	GetId() *string

	// The Channel's name. The name can contain only letters, numbers, periods, and underscores. The name must begin with a letter.
	GetName() *string

	// The category of the Channel.
	GetCategory() ChannelCategoryEnum

	// The Channel's current state.
	GetLifecycleState() LifecycleStateEnum

	// When the resource was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	GetTimeCreated() *common.SDKTime

	// When the resource was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	GetTimeUpdated() *common.SDKTime

	// A short description of the Channel.
	GetDescription() *string

	// The number of milliseconds before a session expires.
	GetSessionExpiryDurationInMilliseconds() *int64

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type channel struct {
	JsonData                            []byte
	Description                         *string                           `mandatory:"false" json:"description"`
	SessionExpiryDurationInMilliseconds *int64                            `mandatory:"false" json:"sessionExpiryDurationInMilliseconds"`
	FreeformTags                        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Id                                  *string                           `mandatory:"true" json:"id"`
	Name                                *string                           `mandatory:"true" json:"name"`
	Category                            ChannelCategoryEnum               `mandatory:"true" json:"category"`
	LifecycleState                      LifecycleStateEnum                `mandatory:"true" json:"lifecycleState"`
	TimeCreated                         *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated                         *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	Type                                string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *channel) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerchannel channel
	s := struct {
		Model Unmarshalerchannel
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.Name = s.Model.Name
	m.Category = s.Model.Category
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Description = s.Model.Description
	m.SessionExpiryDurationInMilliseconds = s.Model.SessionExpiryDurationInMilliseconds
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *channel) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "APPEVENT":
		mm := AppEventChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FACEBOOK":
		mm := FacebookChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WEB":
		mm := WebChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MSTEAMS":
		mm := MsTeamsChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TWILIO":
		mm := TwilioChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OSVC":
		mm := OsvcChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ANDROID":
		mm := AndroidChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OSS":
		mm := OssChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WEBHOOK":
		mm := WebhookChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPLICATION":
		mm := ApplicationChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SLACK":
		mm := SlackChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CORTANA":
		mm := CortanaChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SERVICECLOUD":
		mm := ServiceCloudChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IOS":
		mm := IosChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEST":
		mm := TestChannel{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Channel: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m channel) GetDescription() *string {
	return m.Description
}

// GetSessionExpiryDurationInMilliseconds returns SessionExpiryDurationInMilliseconds
func (m channel) GetSessionExpiryDurationInMilliseconds() *int64 {
	return m.SessionExpiryDurationInMilliseconds
}

// GetFreeformTags returns FreeformTags
func (m channel) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m channel) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetId returns Id
func (m channel) GetId() *string {
	return m.Id
}

// GetName returns Name
func (m channel) GetName() *string {
	return m.Name
}

// GetCategory returns Category
func (m channel) GetCategory() ChannelCategoryEnum {
	return m.Category
}

// GetLifecycleState returns LifecycleState
func (m channel) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m channel) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m channel) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m channel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m channel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingChannelCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetChannelCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
