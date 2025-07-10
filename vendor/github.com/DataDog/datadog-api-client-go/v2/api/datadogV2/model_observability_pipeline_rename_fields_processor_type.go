// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineRenameFieldsProcessorType The processor type. The value should always be `rename_fields`.
type ObservabilityPipelineRenameFieldsProcessorType string

// List of ObservabilityPipelineRenameFieldsProcessorType.
const (
	OBSERVABILITYPIPELINERENAMEFIELDSPROCESSORTYPE_RENAME_FIELDS ObservabilityPipelineRenameFieldsProcessorType = "rename_fields"
)

var allowedObservabilityPipelineRenameFieldsProcessorTypeEnumValues = []ObservabilityPipelineRenameFieldsProcessorType{
	OBSERVABILITYPIPELINERENAMEFIELDSPROCESSORTYPE_RENAME_FIELDS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineRenameFieldsProcessorType) GetAllowedValues() []ObservabilityPipelineRenameFieldsProcessorType {
	return allowedObservabilityPipelineRenameFieldsProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineRenameFieldsProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineRenameFieldsProcessorType(value)
	return nil
}

// NewObservabilityPipelineRenameFieldsProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineRenameFieldsProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineRenameFieldsProcessorTypeFromValue(v string) (*ObservabilityPipelineRenameFieldsProcessorType, error) {
	ev := ObservabilityPipelineRenameFieldsProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineRenameFieldsProcessorType: valid values are %v", v, allowedObservabilityPipelineRenameFieldsProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineRenameFieldsProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineRenameFieldsProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineRenameFieldsProcessorType value.
func (v ObservabilityPipelineRenameFieldsProcessorType) Ptr() *ObservabilityPipelineRenameFieldsProcessorType {
	return &v
}
