// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebIntegrationAccountType Account resource type.
type WebIntegrationAccountType string

// List of WebIntegrationAccountType.
const (
	WEBINTEGRATIONACCOUNTTYPE_ACCOUNT WebIntegrationAccountType = "Account"
)

var allowedWebIntegrationAccountTypeEnumValues = []WebIntegrationAccountType{
	WEBINTEGRATIONACCOUNTTYPE_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WebIntegrationAccountType) GetAllowedValues() []WebIntegrationAccountType {
	return allowedWebIntegrationAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WebIntegrationAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WebIntegrationAccountType(value)
	return nil
}

// NewWebIntegrationAccountTypeFromValue returns a pointer to a valid WebIntegrationAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWebIntegrationAccountTypeFromValue(v string) (*WebIntegrationAccountType, error) {
	ev := WebIntegrationAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WebIntegrationAccountType: valid values are %v", v, allowedWebIntegrationAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WebIntegrationAccountType) IsValid() bool {
	for _, existing := range allowedWebIntegrationAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebIntegrationAccountType value.
func (v WebIntegrationAccountType) Ptr() *WebIntegrationAccountType {
	return &v
}
