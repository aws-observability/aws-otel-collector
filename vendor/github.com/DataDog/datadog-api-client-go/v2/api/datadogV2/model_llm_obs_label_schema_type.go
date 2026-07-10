// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsLabelSchemaType Type of a label in an annotation queue label schema.
type LLMObsLabelSchemaType string

// List of LLMObsLabelSchemaType.
const (
	LLMOBSLABELSCHEMATYPE_SCORE       LLMObsLabelSchemaType = "score"
	LLMOBSLABELSCHEMATYPE_CATEGORICAL LLMObsLabelSchemaType = "categorical"
	LLMOBSLABELSCHEMATYPE_BOOLEAN     LLMObsLabelSchemaType = "boolean"
	LLMOBSLABELSCHEMATYPE_TEXT        LLMObsLabelSchemaType = "text"
)

var allowedLLMObsLabelSchemaTypeEnumValues = []LLMObsLabelSchemaType{
	LLMOBSLABELSCHEMATYPE_SCORE,
	LLMOBSLABELSCHEMATYPE_CATEGORICAL,
	LLMOBSLABELSCHEMATYPE_BOOLEAN,
	LLMOBSLABELSCHEMATYPE_TEXT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsLabelSchemaType) GetAllowedValues() []LLMObsLabelSchemaType {
	return allowedLLMObsLabelSchemaTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsLabelSchemaType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsLabelSchemaType(value)
	return nil
}

// NewLLMObsLabelSchemaTypeFromValue returns a pointer to a valid LLMObsLabelSchemaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsLabelSchemaTypeFromValue(v string) (*LLMObsLabelSchemaType, error) {
	ev := LLMObsLabelSchemaType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsLabelSchemaType: valid values are %v", v, allowedLLMObsLabelSchemaTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsLabelSchemaType) IsValid() bool {
	for _, existing := range allowedLLMObsLabelSchemaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsLabelSchemaType value.
func (v LLMObsLabelSchemaType) Ptr() *LLMObsLabelSchemaType {
	return &v
}
