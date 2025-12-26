// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingDataType Security findings resource type.
type FindingDataType string

// List of FindingDataType.
const (
	FINDINGDATATYPE_FINDINGS FindingDataType = "findings"
)

var allowedFindingDataTypeEnumValues = []FindingDataType{
	FINDINGDATATYPE_FINDINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FindingDataType) GetAllowedValues() []FindingDataType {
	return allowedFindingDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FindingDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FindingDataType(value)
	return nil
}

// NewFindingDataTypeFromValue returns a pointer to a valid FindingDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFindingDataTypeFromValue(v string) (*FindingDataType, error) {
	ev := FindingDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FindingDataType: valid values are %v", v, allowedFindingDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FindingDataType) IsValid() bool {
	for _, existing := range allowedFindingDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FindingDataType value.
func (v FindingDataType) Ptr() *FindingDataType {
	return &v
}
