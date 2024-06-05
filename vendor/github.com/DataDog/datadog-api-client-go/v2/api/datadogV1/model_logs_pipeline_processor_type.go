// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsPipelineProcessorType Type of logs pipeline processor.
type LogsPipelineProcessorType string

// List of LogsPipelineProcessorType.
const (
	LOGSPIPELINEPROCESSORTYPE_PIPELINE LogsPipelineProcessorType = "pipeline"
)

var allowedLogsPipelineProcessorTypeEnumValues = []LogsPipelineProcessorType{
	LOGSPIPELINEPROCESSORTYPE_PIPELINE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsPipelineProcessorType) GetAllowedValues() []LogsPipelineProcessorType {
	return allowedLogsPipelineProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsPipelineProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsPipelineProcessorType(value)
	return nil
}

// NewLogsPipelineProcessorTypeFromValue returns a pointer to a valid LogsPipelineProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsPipelineProcessorTypeFromValue(v string) (*LogsPipelineProcessorType, error) {
	ev := LogsPipelineProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsPipelineProcessorType: valid values are %v", v, allowedLogsPipelineProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsPipelineProcessorType) IsValid() bool {
	for _, existing := range allowedLogsPipelineProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsPipelineProcessorType value.
func (v LogsPipelineProcessorType) Ptr() *LogsPipelineProcessorType {
	return &v
}
