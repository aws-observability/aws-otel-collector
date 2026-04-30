// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPageDataType Status pages resource type.
type StatusPageDataType string

// List of StatusPageDataType.
const (
	STATUSPAGEDATATYPE_STATUS_PAGES StatusPageDataType = "status_pages"
)

var allowedStatusPageDataTypeEnumValues = []StatusPageDataType{
	STATUSPAGEDATATYPE_STATUS_PAGES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatusPageDataType) GetAllowedValues() []StatusPageDataType {
	return allowedStatusPageDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatusPageDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatusPageDataType(value)
	return nil
}

// NewStatusPageDataTypeFromValue returns a pointer to a valid StatusPageDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatusPageDataTypeFromValue(v string) (*StatusPageDataType, error) {
	ev := StatusPageDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatusPageDataType: valid values are %v", v, allowedStatusPageDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatusPageDataType) IsValid() bool {
	for _, existing := range allowedStatusPageDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusPageDataType value.
func (v StatusPageDataType) Ptr() *StatusPageDataType {
	return &v
}
