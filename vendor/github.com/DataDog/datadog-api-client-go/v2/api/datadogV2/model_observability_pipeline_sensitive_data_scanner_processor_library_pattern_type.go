// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType Indicates that a predefined library pattern is used.
type ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType string

// List of ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType.
const (
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORLIBRARYPATTERNTYPE_LIBRARY ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType = "library"
)

var allowedObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternTypeEnumValues = []ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType{
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORLIBRARYPATTERNTYPE_LIBRARY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType) GetAllowedValues() []ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType {
	return allowedObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType(value)
	return nil
}

// NewObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternTypeFromValue returns a pointer to a valid ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternTypeFromValue(v string) (*ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType, error) {
	ev := ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType: valid values are %v", v, allowedObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType value.
func (v ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType) Ptr() *ObservabilityPipelineSensitiveDataScannerProcessorLibraryPatternType {
	return &v
}
