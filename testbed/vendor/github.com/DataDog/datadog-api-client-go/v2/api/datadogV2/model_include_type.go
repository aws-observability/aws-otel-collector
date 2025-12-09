// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncludeType Supported include types.
type IncludeType string

// List of IncludeType.
const (
	INCLUDETYPE_SCHEMA     IncludeType = "schema"
	INCLUDETYPE_RAW_SCHEMA IncludeType = "raw_schema"
	INCLUDETYPE_ONCALL     IncludeType = "oncall"
	INCLUDETYPE_INCIDENT   IncludeType = "incident"
	INCLUDETYPE_RELATION   IncludeType = "relation"
)

var allowedIncludeTypeEnumValues = []IncludeType{
	INCLUDETYPE_SCHEMA,
	INCLUDETYPE_RAW_SCHEMA,
	INCLUDETYPE_ONCALL,
	INCLUDETYPE_INCIDENT,
	INCLUDETYPE_RELATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncludeType) GetAllowedValues() []IncludeType {
	return allowedIncludeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncludeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncludeType(value)
	return nil
}

// NewIncludeTypeFromValue returns a pointer to a valid IncludeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncludeTypeFromValue(v string) (*IncludeType, error) {
	ev := IncludeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncludeType: valid values are %v", v, allowedIncludeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncludeType) IsValid() bool {
	for _, existing := range allowedIncludeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncludeType value.
func (v IncludeType) Ptr() *IncludeType {
	return &v
}
