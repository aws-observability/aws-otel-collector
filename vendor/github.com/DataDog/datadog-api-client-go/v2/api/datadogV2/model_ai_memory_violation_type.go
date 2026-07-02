// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiMemoryViolationType The type of AI memory violation result indicating whether it is a true positive or false positive.
type AiMemoryViolationType string

// List of AiMemoryViolationType.
const (
	AIMEMORYVIOLATIONTYPE_TP AiMemoryViolationType = "TP"
	AIMEMORYVIOLATIONTYPE_FP AiMemoryViolationType = "FP"
)

var allowedAiMemoryViolationTypeEnumValues = []AiMemoryViolationType{
	AIMEMORYVIOLATIONTYPE_TP,
	AIMEMORYVIOLATIONTYPE_FP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiMemoryViolationType) GetAllowedValues() []AiMemoryViolationType {
	return allowedAiMemoryViolationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiMemoryViolationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiMemoryViolationType(value)
	return nil
}

// NewAiMemoryViolationTypeFromValue returns a pointer to a valid AiMemoryViolationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiMemoryViolationTypeFromValue(v string) (*AiMemoryViolationType, error) {
	ev := AiMemoryViolationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiMemoryViolationType: valid values are %v", v, allowedAiMemoryViolationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiMemoryViolationType) IsValid() bool {
	for _, existing := range allowedAiMemoryViolationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiMemoryViolationType value.
func (v AiMemoryViolationType) Ptr() *AiMemoryViolationType {
	return &v
}
