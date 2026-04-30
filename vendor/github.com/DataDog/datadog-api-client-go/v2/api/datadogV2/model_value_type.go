// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ValueType The type of values for the feature flag variants.
type ValueType string

// List of ValueType.
const (
	VALUETYPE_BOOLEAN ValueType = "BOOLEAN"
	VALUETYPE_INTEGER ValueType = "INTEGER"
	VALUETYPE_NUMERIC ValueType = "NUMERIC"
	VALUETYPE_STRING  ValueType = "STRING"
	VALUETYPE_JSON    ValueType = "JSON"
)

var allowedValueTypeEnumValues = []ValueType{
	VALUETYPE_BOOLEAN,
	VALUETYPE_INTEGER,
	VALUETYPE_NUMERIC,
	VALUETYPE_STRING,
	VALUETYPE_JSON,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ValueType) GetAllowedValues() []ValueType {
	return allowedValueTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ValueType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ValueType(value)
	return nil
}

// NewValueTypeFromValue returns a pointer to a valid ValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewValueTypeFromValue(v string) (*ValueType, error) {
	ev := ValueType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ValueType: valid values are %v", v, allowedValueTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ValueType) IsValid() bool {
	for _, existing := range allowedValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ValueType value.
func (v ValueType) Ptr() *ValueType {
	return &v
}
