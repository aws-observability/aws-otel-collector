// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpanType Resource type for an LLM Observability span.
type LLMObsSpanType string

// List of LLMObsSpanType.
const (
	LLMOBSSPANTYPE_SPAN LLMObsSpanType = "span"
)

var allowedLLMObsSpanTypeEnumValues = []LLMObsSpanType{
	LLMOBSSPANTYPE_SPAN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsSpanType) GetAllowedValues() []LLMObsSpanType {
	return allowedLLMObsSpanTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsSpanType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsSpanType(value)
	return nil
}

// NewLLMObsSpanTypeFromValue returns a pointer to a valid LLMObsSpanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsSpanTypeFromValue(v string) (*LLMObsSpanType, error) {
	ev := LLMObsSpanType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsSpanType: valid values are %v", v, allowedLLMObsSpanTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsSpanType) IsValid() bool {
	for _, existing := range allowedLLMObsSpanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsSpanType value.
func (v LLMObsSpanType) Ptr() *LLMObsSpanType {
	return &v
}
