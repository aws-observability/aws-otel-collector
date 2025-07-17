// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventCustomAttributesChangedResourceType The type of the resource that was changed.
type ChangeEventCustomAttributesChangedResourceType string

// List of ChangeEventCustomAttributesChangedResourceType.
const (
	CHANGEEVENTCUSTOMATTRIBUTESCHANGEDRESOURCETYPE_FEATURE_FLAG  ChangeEventCustomAttributesChangedResourceType = "feature_flag"
	CHANGEEVENTCUSTOMATTRIBUTESCHANGEDRESOURCETYPE_CONFIGURATION ChangeEventCustomAttributesChangedResourceType = "configuration"
)

var allowedChangeEventCustomAttributesChangedResourceTypeEnumValues = []ChangeEventCustomAttributesChangedResourceType{
	CHANGEEVENTCUSTOMATTRIBUTESCHANGEDRESOURCETYPE_FEATURE_FLAG,
	CHANGEEVENTCUSTOMATTRIBUTESCHANGEDRESOURCETYPE_CONFIGURATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeEventCustomAttributesChangedResourceType) GetAllowedValues() []ChangeEventCustomAttributesChangedResourceType {
	return allowedChangeEventCustomAttributesChangedResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeEventCustomAttributesChangedResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeEventCustomAttributesChangedResourceType(value)
	return nil
}

// NewChangeEventCustomAttributesChangedResourceTypeFromValue returns a pointer to a valid ChangeEventCustomAttributesChangedResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeEventCustomAttributesChangedResourceTypeFromValue(v string) (*ChangeEventCustomAttributesChangedResourceType, error) {
	ev := ChangeEventCustomAttributesChangedResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeEventCustomAttributesChangedResourceType: valid values are %v", v, allowedChangeEventCustomAttributesChangedResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeEventCustomAttributesChangedResourceType) IsValid() bool {
	for _, existing := range allowedChangeEventCustomAttributesChangedResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeEventCustomAttributesChangedResourceType value.
func (v ChangeEventCustomAttributesChangedResourceType) Ptr() *ChangeEventCustomAttributesChangedResourceType {
	return &v
}
