// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType Indicates a custom regular expression is used for matching.
type ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType string

// List of ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType.
const (
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORCUSTOMPATTERNTYPE_CUSTOM ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType = "custom"
)

var allowedObservabilityPipelineSensitiveDataScannerProcessorCustomPatternTypeEnumValues = []ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType{
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORCUSTOMPATTERNTYPE_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType) GetAllowedValues() []ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType {
	return allowedObservabilityPipelineSensitiveDataScannerProcessorCustomPatternTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType(value)
	return nil
}

// NewObservabilityPipelineSensitiveDataScannerProcessorCustomPatternTypeFromValue returns a pointer to a valid ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSensitiveDataScannerProcessorCustomPatternTypeFromValue(v string) (*ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType, error) {
	ev := ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType: valid values are %v", v, allowedObservabilityPipelineSensitiveDataScannerProcessorCustomPatternTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSensitiveDataScannerProcessorCustomPatternTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType value.
func (v ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType) Ptr() *ObservabilityPipelineSensitiveDataScannerProcessorCustomPatternType {
	return &v
}
