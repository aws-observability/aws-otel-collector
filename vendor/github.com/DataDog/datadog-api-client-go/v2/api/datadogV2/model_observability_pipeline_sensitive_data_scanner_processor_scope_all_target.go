// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget Applies the rule to all fields.
type ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget string

// List of ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget.
const (
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORSCOPEALLTARGET_ALL ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget = "all"
)

var allowedObservabilityPipelineSensitiveDataScannerProcessorScopeAllTargetEnumValues = []ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget{
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORSCOPEALLTARGET_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget) GetAllowedValues() []ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget {
	return allowedObservabilityPipelineSensitiveDataScannerProcessorScopeAllTargetEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget(value)
	return nil
}

// NewObservabilityPipelineSensitiveDataScannerProcessorScopeAllTargetFromValue returns a pointer to a valid ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSensitiveDataScannerProcessorScopeAllTargetFromValue(v string) (*ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget, error) {
	ev := ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget: valid values are %v", v, allowedObservabilityPipelineSensitiveDataScannerProcessorScopeAllTargetEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSensitiveDataScannerProcessorScopeAllTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget value.
func (v ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget) Ptr() *ObservabilityPipelineSensitiveDataScannerProcessorScopeAllTarget {
	return &v
}
