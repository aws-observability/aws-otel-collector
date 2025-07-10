// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorType The processor type. The value should always be `sensitive_data_scanner`.
type ObservabilityPipelineSensitiveDataScannerProcessorType string

// List of ObservabilityPipelineSensitiveDataScannerProcessorType.
const (
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORTYPE_SENSITIVE_DATA_SCANNER ObservabilityPipelineSensitiveDataScannerProcessorType = "sensitive_data_scanner"
)

var allowedObservabilityPipelineSensitiveDataScannerProcessorTypeEnumValues = []ObservabilityPipelineSensitiveDataScannerProcessorType{
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORTYPE_SENSITIVE_DATA_SCANNER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorType) GetAllowedValues() []ObservabilityPipelineSensitiveDataScannerProcessorType {
	return allowedObservabilityPipelineSensitiveDataScannerProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSensitiveDataScannerProcessorType(value)
	return nil
}

// NewObservabilityPipelineSensitiveDataScannerProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineSensitiveDataScannerProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSensitiveDataScannerProcessorTypeFromValue(v string) (*ObservabilityPipelineSensitiveDataScannerProcessorType, error) {
	ev := ObservabilityPipelineSensitiveDataScannerProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSensitiveDataScannerProcessorType: valid values are %v", v, allowedObservabilityPipelineSensitiveDataScannerProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSensitiveDataScannerProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSensitiveDataScannerProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSensitiveDataScannerProcessorType value.
func (v ObservabilityPipelineSensitiveDataScannerProcessorType) Ptr() *ObservabilityPipelineSensitiveDataScannerProcessorType {
	return &v
}
