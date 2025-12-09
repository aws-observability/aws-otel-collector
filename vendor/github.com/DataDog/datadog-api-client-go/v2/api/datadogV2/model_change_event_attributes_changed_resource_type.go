// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventAttributesChangedResourceType The type of the changed resource.
type ChangeEventAttributesChangedResourceType string

// List of ChangeEventAttributesChangedResourceType.
const (
	CHANGEEVENTATTRIBUTESCHANGEDRESOURCETYPE_FEATURE_FLAG  ChangeEventAttributesChangedResourceType = "feature_flag"
	CHANGEEVENTATTRIBUTESCHANGEDRESOURCETYPE_CONFIGURATION ChangeEventAttributesChangedResourceType = "configuration"
)

var allowedChangeEventAttributesChangedResourceTypeEnumValues = []ChangeEventAttributesChangedResourceType{
	CHANGEEVENTATTRIBUTESCHANGEDRESOURCETYPE_FEATURE_FLAG,
	CHANGEEVENTATTRIBUTESCHANGEDRESOURCETYPE_CONFIGURATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeEventAttributesChangedResourceType) GetAllowedValues() []ChangeEventAttributesChangedResourceType {
	return allowedChangeEventAttributesChangedResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeEventAttributesChangedResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeEventAttributesChangedResourceType(value)
	return nil
}

// NewChangeEventAttributesChangedResourceTypeFromValue returns a pointer to a valid ChangeEventAttributesChangedResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeEventAttributesChangedResourceTypeFromValue(v string) (*ChangeEventAttributesChangedResourceType, error) {
	ev := ChangeEventAttributesChangedResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeEventAttributesChangedResourceType: valid values are %v", v, allowedChangeEventAttributesChangedResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeEventAttributesChangedResourceType) IsValid() bool {
	for _, existing := range allowedChangeEventAttributesChangedResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeEventAttributesChangedResourceType value.
func (v ChangeEventAttributesChangedResourceType) Ptr() *ChangeEventAttributesChangedResourceType {
	return &v
}
