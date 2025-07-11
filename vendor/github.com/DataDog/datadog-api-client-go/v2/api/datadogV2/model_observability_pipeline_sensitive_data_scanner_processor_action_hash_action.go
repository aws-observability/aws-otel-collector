// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction Action type that replaces the matched sensitive data with a hashed representation, preserving structure while securing content.
type ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction string

// List of ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction.
const (
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORACTIONHASHACTION_HASH ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction = "hash"
)

var allowedObservabilityPipelineSensitiveDataScannerProcessorActionHashActionEnumValues = []ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction{
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORACTIONHASHACTION_HASH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction) GetAllowedValues() []ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction {
	return allowedObservabilityPipelineSensitiveDataScannerProcessorActionHashActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction(value)
	return nil
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionHashActionFromValue returns a pointer to a valid ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionHashActionFromValue(v string) (*ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction, error) {
	ev := ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction: valid values are %v", v, allowedObservabilityPipelineSensitiveDataScannerProcessorActionHashActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSensitiveDataScannerProcessorActionHashActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction value.
func (v ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction) Ptr() *ObservabilityPipelineSensitiveDataScannerProcessorActionHashAction {
	return &v
}
