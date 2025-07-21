// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventCustomAttributesAuthorType Author's type.
type ChangeEventCustomAttributesAuthorType string

// List of ChangeEventCustomAttributesAuthorType.
const (
	CHANGEEVENTCUSTOMATTRIBUTESAUTHORTYPE_USER       ChangeEventCustomAttributesAuthorType = "user"
	CHANGEEVENTCUSTOMATTRIBUTESAUTHORTYPE_SYSTEM     ChangeEventCustomAttributesAuthorType = "system"
	CHANGEEVENTCUSTOMATTRIBUTESAUTHORTYPE_API        ChangeEventCustomAttributesAuthorType = "api"
	CHANGEEVENTCUSTOMATTRIBUTESAUTHORTYPE_AUTOMATION ChangeEventCustomAttributesAuthorType = "automation"
)

var allowedChangeEventCustomAttributesAuthorTypeEnumValues = []ChangeEventCustomAttributesAuthorType{
	CHANGEEVENTCUSTOMATTRIBUTESAUTHORTYPE_USER,
	CHANGEEVENTCUSTOMATTRIBUTESAUTHORTYPE_SYSTEM,
	CHANGEEVENTCUSTOMATTRIBUTESAUTHORTYPE_API,
	CHANGEEVENTCUSTOMATTRIBUTESAUTHORTYPE_AUTOMATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeEventCustomAttributesAuthorType) GetAllowedValues() []ChangeEventCustomAttributesAuthorType {
	return allowedChangeEventCustomAttributesAuthorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeEventCustomAttributesAuthorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeEventCustomAttributesAuthorType(value)
	return nil
}

// NewChangeEventCustomAttributesAuthorTypeFromValue returns a pointer to a valid ChangeEventCustomAttributesAuthorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeEventCustomAttributesAuthorTypeFromValue(v string) (*ChangeEventCustomAttributesAuthorType, error) {
	ev := ChangeEventCustomAttributesAuthorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeEventCustomAttributesAuthorType: valid values are %v", v, allowedChangeEventCustomAttributesAuthorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeEventCustomAttributesAuthorType) IsValid() bool {
	for _, existing := range allowedChangeEventCustomAttributesAuthorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeEventCustomAttributesAuthorType value.
func (v ChangeEventCustomAttributesAuthorType) Ptr() *ChangeEventCustomAttributesAuthorType {
	return &v
}
