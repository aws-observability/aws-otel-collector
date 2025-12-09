// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection Indicates whether to redact characters from the first or last part of the matched value.
type ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection string

// List of ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection.
const (
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORACTIONPARTIALREDACTOPTIONSDIRECTION_FIRST ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection = "first"
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORACTIONPARTIALREDACTOPTIONSDIRECTION_LAST  ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection = "last"
)

var allowedObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirectionEnumValues = []ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection{
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORACTIONPARTIALREDACTOPTIONSDIRECTION_FIRST,
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORACTIONPARTIALREDACTOPTIONSDIRECTION_LAST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection) GetAllowedValues() []ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection {
	return allowedObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection(value)
	return nil
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirectionFromValue returns a pointer to a valid ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirectionFromValue(v string) (*ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection, error) {
	ev := ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection: valid values are %v", v, allowedObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection value.
func (v ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection) Ptr() *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection {
	return &v
}
