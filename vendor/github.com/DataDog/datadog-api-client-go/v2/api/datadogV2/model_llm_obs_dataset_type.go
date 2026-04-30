// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetType Resource type of an LLM Observability dataset.
type LLMObsDatasetType string

// List of LLMObsDatasetType.
const (
	LLMOBSDATASETTYPE_DATASETS LLMObsDatasetType = "datasets"
)

var allowedLLMObsDatasetTypeEnumValues = []LLMObsDatasetType{
	LLMOBSDATASETTYPE_DATASETS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsDatasetType) GetAllowedValues() []LLMObsDatasetType {
	return allowedLLMObsDatasetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsDatasetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsDatasetType(value)
	return nil
}

// NewLLMObsDatasetTypeFromValue returns a pointer to a valid LLMObsDatasetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsDatasetTypeFromValue(v string) (*LLMObsDatasetType, error) {
	ev := LLMObsDatasetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsDatasetType: valid values are %v", v, allowedLLMObsDatasetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsDatasetType) IsValid() bool {
	for _, existing := range allowedLLMObsDatasetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsDatasetType value.
func (v LLMObsDatasetType) Ptr() *LLMObsDatasetType {
	return &v
}
