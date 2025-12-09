// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomAttributeConfigResourceType Custom attributes config JSON:API resource type
type CustomAttributeConfigResourceType string

// List of CustomAttributeConfigResourceType.
const (
	CUSTOMATTRIBUTECONFIGRESOURCETYPE_CUSTOM_ATTRIBUTE CustomAttributeConfigResourceType = "custom_attribute"
)

var allowedCustomAttributeConfigResourceTypeEnumValues = []CustomAttributeConfigResourceType{
	CUSTOMATTRIBUTECONFIGRESOURCETYPE_CUSTOM_ATTRIBUTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomAttributeConfigResourceType) GetAllowedValues() []CustomAttributeConfigResourceType {
	return allowedCustomAttributeConfigResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomAttributeConfigResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomAttributeConfigResourceType(value)
	return nil
}

// NewCustomAttributeConfigResourceTypeFromValue returns a pointer to a valid CustomAttributeConfigResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomAttributeConfigResourceTypeFromValue(v string) (*CustomAttributeConfigResourceType, error) {
	ev := CustomAttributeConfigResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomAttributeConfigResourceType: valid values are %v", v, allowedCustomAttributeConfigResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomAttributeConfigResourceType) IsValid() bool {
	for _, existing := range allowedCustomAttributeConfigResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomAttributeConfigResourceType value.
func (v CustomAttributeConfigResourceType) Ptr() *CustomAttributeConfigResourceType {
	return &v
}
