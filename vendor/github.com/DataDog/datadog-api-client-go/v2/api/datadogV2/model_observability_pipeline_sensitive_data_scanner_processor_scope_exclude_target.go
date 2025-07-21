// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget Excludes specific fields from processing.
type ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget string

// List of ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget.
const (
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORSCOPEEXCLUDETARGET_EXCLUDE ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget = "exclude"
)

var allowedObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTargetEnumValues = []ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget{
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORSCOPEEXCLUDETARGET_EXCLUDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget) GetAllowedValues() []ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget {
	return allowedObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTargetEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget(value)
	return nil
}

// NewObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTargetFromValue returns a pointer to a valid ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTargetFromValue(v string) (*ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget, error) {
	ev := ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget: valid values are %v", v, allowedObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTargetEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget value.
func (v ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget) Ptr() *ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeTarget {
	return &v
}
