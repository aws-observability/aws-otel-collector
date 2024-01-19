// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TargetFormatType If the `target_type` of the remapper is `attribute`, try to cast the value to a new specific type.
// If the cast is not possible, the original type is kept. `string`, `integer`, or `double` are the possible types.
// If the `target_type` is `tag`, this parameter may not be specified.
type TargetFormatType string

// List of TargetFormatType.
const (
	TARGETFORMATTYPE_AUTO    TargetFormatType = "auto"
	TARGETFORMATTYPE_STRING  TargetFormatType = "string"
	TARGETFORMATTYPE_INTEGER TargetFormatType = "integer"
	TARGETFORMATTYPE_DOUBLE  TargetFormatType = "double"
)

var allowedTargetFormatTypeEnumValues = []TargetFormatType{
	TARGETFORMATTYPE_AUTO,
	TARGETFORMATTYPE_STRING,
	TARGETFORMATTYPE_INTEGER,
	TARGETFORMATTYPE_DOUBLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TargetFormatType) GetAllowedValues() []TargetFormatType {
	return allowedTargetFormatTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TargetFormatType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TargetFormatType(value)
	return nil
}

// NewTargetFormatTypeFromValue returns a pointer to a valid TargetFormatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTargetFormatTypeFromValue(v string) (*TargetFormatType, error) {
	ev := TargetFormatType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TargetFormatType: valid values are %v", v, allowedTargetFormatTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TargetFormatType) IsValid() bool {
	for _, existing := range allowedTargetFormatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TargetFormatType value.
func (v TargetFormatType) Ptr() *TargetFormatType {
	return &v
}
