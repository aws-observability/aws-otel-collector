// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiMemoryViolationResultDataType AI memory violation result resource type.
type AiMemoryViolationResultDataType string

// List of AiMemoryViolationResultDataType.
const (
	AIMEMORYVIOLATIONRESULTDATATYPE_AI_MEMORY_VIOLATION_RESULT AiMemoryViolationResultDataType = "ai_memory_violation_result"
)

var allowedAiMemoryViolationResultDataTypeEnumValues = []AiMemoryViolationResultDataType{
	AIMEMORYVIOLATIONRESULTDATATYPE_AI_MEMORY_VIOLATION_RESULT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiMemoryViolationResultDataType) GetAllowedValues() []AiMemoryViolationResultDataType {
	return allowedAiMemoryViolationResultDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiMemoryViolationResultDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiMemoryViolationResultDataType(value)
	return nil
}

// NewAiMemoryViolationResultDataTypeFromValue returns a pointer to a valid AiMemoryViolationResultDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiMemoryViolationResultDataTypeFromValue(v string) (*AiMemoryViolationResultDataType, error) {
	ev := AiMemoryViolationResultDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiMemoryViolationResultDataType: valid values are %v", v, allowedAiMemoryViolationResultDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiMemoryViolationResultDataType) IsValid() bool {
	for _, existing := range allowedAiMemoryViolationResultDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiMemoryViolationResultDataType value.
func (v AiMemoryViolationResultDataType) Ptr() *AiMemoryViolationResultDataType {
	return &v
}
