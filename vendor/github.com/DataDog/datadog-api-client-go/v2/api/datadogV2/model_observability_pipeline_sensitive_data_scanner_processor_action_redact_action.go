// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction Action type that completely replaces the matched sensitive data with a fixed replacement string to remove all visibility.
type ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction string

// List of ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction.
const (
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORACTIONREDACTACTION_REDACT ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction = "redact"
)

var allowedObservabilityPipelineSensitiveDataScannerProcessorActionRedactActionEnumValues = []ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction{
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORACTIONREDACTACTION_REDACT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction) GetAllowedValues() []ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction {
	return allowedObservabilityPipelineSensitiveDataScannerProcessorActionRedactActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction(value)
	return nil
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionRedactActionFromValue returns a pointer to a valid ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionRedactActionFromValue(v string) (*ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction, error) {
	ev := ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction: valid values are %v", v, allowedObservabilityPipelineSensitiveDataScannerProcessorActionRedactActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSensitiveDataScannerProcessorActionRedactActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction value.
func (v ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction) Ptr() *ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAction {
	return &v
}
