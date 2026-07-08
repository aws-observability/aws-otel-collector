// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetExportFormat Supported export format for an LLM Observability dataset.
type LLMObsDatasetExportFormat string

// List of LLMObsDatasetExportFormat.
const (
	LLMOBSDATASETEXPORTFORMAT_CSV LLMObsDatasetExportFormat = "csv"
)

var allowedLLMObsDatasetExportFormatEnumValues = []LLMObsDatasetExportFormat{
	LLMOBSDATASETEXPORTFORMAT_CSV,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsDatasetExportFormat) GetAllowedValues() []LLMObsDatasetExportFormat {
	return allowedLLMObsDatasetExportFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsDatasetExportFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsDatasetExportFormat(value)
	return nil
}

// NewLLMObsDatasetExportFormatFromValue returns a pointer to a valid LLMObsDatasetExportFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsDatasetExportFormatFromValue(v string) (*LLMObsDatasetExportFormat, error) {
	ev := LLMObsDatasetExportFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsDatasetExportFormat: valid values are %v", v, allowedLLMObsDatasetExportFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsDatasetExportFormat) IsValid() bool {
	for _, existing := range allowedLLMObsDatasetExportFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsDatasetExportFormat value.
func (v LLMObsDatasetExportFormat) Ptr() *LLMObsDatasetExportFormat {
	return &v
}
