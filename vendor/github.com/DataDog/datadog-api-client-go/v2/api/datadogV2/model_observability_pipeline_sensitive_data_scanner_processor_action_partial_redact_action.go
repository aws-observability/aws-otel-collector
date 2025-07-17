// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction Action type that redacts part of the sensitive data while preserving a configurable number of characters, typically used for masking purposes (e.g., show last 4 digits of a credit card).
type ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction string

// List of ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction.
const (
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORACTIONPARTIALREDACTACTION_PARTIAL_REDACT ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction = "partial_redact"
)

var allowedObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactActionEnumValues = []ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction{
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORACTIONPARTIALREDACTACTION_PARTIAL_REDACT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction) GetAllowedValues() []ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction {
	return allowedObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction(value)
	return nil
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactActionFromValue returns a pointer to a valid ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactActionFromValue(v string) (*ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction, error) {
	ev := ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction: valid values are %v", v, allowedObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction value.
func (v ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction) Ptr() *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAction {
	return &v
}
