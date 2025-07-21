// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget Applies the rule only to included fields.
type ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget string

// List of ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget.
const (
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORSCOPEINCLUDETARGET_INCLUDE ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget = "include"
)

var allowedObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTargetEnumValues = []ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget{
	OBSERVABILITYPIPELINESENSITIVEDATASCANNERPROCESSORSCOPEINCLUDETARGET_INCLUDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget) GetAllowedValues() []ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget {
	return allowedObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTargetEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget(value)
	return nil
}

// NewObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTargetFromValue returns a pointer to a valid ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTargetFromValue(v string) (*ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget, error) {
	ev := ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget: valid values are %v", v, allowedObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTargetEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget value.
func (v ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget) Ptr() *ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeTarget {
	return &v
}
