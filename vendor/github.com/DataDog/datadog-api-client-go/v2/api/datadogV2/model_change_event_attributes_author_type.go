// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventAttributesAuthorType The type of the author.
type ChangeEventAttributesAuthorType string

// List of ChangeEventAttributesAuthorType.
const (
	CHANGEEVENTATTRIBUTESAUTHORTYPE_USER       ChangeEventAttributesAuthorType = "user"
	CHANGEEVENTATTRIBUTESAUTHORTYPE_SYSTEM     ChangeEventAttributesAuthorType = "system"
	CHANGEEVENTATTRIBUTESAUTHORTYPE_API        ChangeEventAttributesAuthorType = "api"
	CHANGEEVENTATTRIBUTESAUTHORTYPE_AUTOMATION ChangeEventAttributesAuthorType = "automation"
)

var allowedChangeEventAttributesAuthorTypeEnumValues = []ChangeEventAttributesAuthorType{
	CHANGEEVENTATTRIBUTESAUTHORTYPE_USER,
	CHANGEEVENTATTRIBUTESAUTHORTYPE_SYSTEM,
	CHANGEEVENTATTRIBUTESAUTHORTYPE_API,
	CHANGEEVENTATTRIBUTESAUTHORTYPE_AUTOMATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeEventAttributesAuthorType) GetAllowedValues() []ChangeEventAttributesAuthorType {
	return allowedChangeEventAttributesAuthorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeEventAttributesAuthorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeEventAttributesAuthorType(value)
	return nil
}

// NewChangeEventAttributesAuthorTypeFromValue returns a pointer to a valid ChangeEventAttributesAuthorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeEventAttributesAuthorTypeFromValue(v string) (*ChangeEventAttributesAuthorType, error) {
	ev := ChangeEventAttributesAuthorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeEventAttributesAuthorType: valid values are %v", v, allowedChangeEventAttributesAuthorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeEventAttributesAuthorType) IsValid() bool {
	for _, existing := range allowedChangeEventAttributesAuthorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeEventAttributesAuthorType value.
func (v ChangeEventAttributesAuthorType) Ptr() *ChangeEventAttributesAuthorType {
	return &v
}
