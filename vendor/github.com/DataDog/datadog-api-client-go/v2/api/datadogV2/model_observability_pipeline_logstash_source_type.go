// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineLogstashSourceType The source type. The value should always be `logstash`.
type ObservabilityPipelineLogstashSourceType string

// List of ObservabilityPipelineLogstashSourceType.
const (
	OBSERVABILITYPIPELINELOGSTASHSOURCETYPE_LOGSTASH ObservabilityPipelineLogstashSourceType = "logstash"
)

var allowedObservabilityPipelineLogstashSourceTypeEnumValues = []ObservabilityPipelineLogstashSourceType{
	OBSERVABILITYPIPELINELOGSTASHSOURCETYPE_LOGSTASH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineLogstashSourceType) GetAllowedValues() []ObservabilityPipelineLogstashSourceType {
	return allowedObservabilityPipelineLogstashSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineLogstashSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineLogstashSourceType(value)
	return nil
}

// NewObservabilityPipelineLogstashSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineLogstashSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineLogstashSourceTypeFromValue(v string) (*ObservabilityPipelineLogstashSourceType, error) {
	ev := ObservabilityPipelineLogstashSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineLogstashSourceType: valid values are %v", v, allowedObservabilityPipelineLogstashSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineLogstashSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineLogstashSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineLogstashSourceType value.
func (v ObservabilityPipelineLogstashSourceType) Ptr() *ObservabilityPipelineLogstashSourceType {
	return &v
}
