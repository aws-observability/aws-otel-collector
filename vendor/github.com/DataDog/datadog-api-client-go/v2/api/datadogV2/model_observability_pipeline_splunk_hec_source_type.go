// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecSourceType The source type. Always `splunk_hec`.
type ObservabilityPipelineSplunkHecSourceType string

// List of ObservabilityPipelineSplunkHecSourceType.
const (
	OBSERVABILITYPIPELINESPLUNKHECSOURCETYPE_SPLUNK_HEC ObservabilityPipelineSplunkHecSourceType = "splunk_hec"
)

var allowedObservabilityPipelineSplunkHecSourceTypeEnumValues = []ObservabilityPipelineSplunkHecSourceType{
	OBSERVABILITYPIPELINESPLUNKHECSOURCETYPE_SPLUNK_HEC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSplunkHecSourceType) GetAllowedValues() []ObservabilityPipelineSplunkHecSourceType {
	return allowedObservabilityPipelineSplunkHecSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSplunkHecSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSplunkHecSourceType(value)
	return nil
}

// NewObservabilityPipelineSplunkHecSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineSplunkHecSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSplunkHecSourceTypeFromValue(v string) (*ObservabilityPipelineSplunkHecSourceType, error) {
	ev := ObservabilityPipelineSplunkHecSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSplunkHecSourceType: valid values are %v", v, allowedObservabilityPipelineSplunkHecSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSplunkHecSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSplunkHecSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSplunkHecSourceType value.
func (v ObservabilityPipelineSplunkHecSourceType) Ptr() *ObservabilityPipelineSplunkHecSourceType {
	return &v
}
