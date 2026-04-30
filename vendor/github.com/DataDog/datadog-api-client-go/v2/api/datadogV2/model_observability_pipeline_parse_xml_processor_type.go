// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineParseXMLProcessorType The processor type. The value should always be `parse_xml`.
type ObservabilityPipelineParseXMLProcessorType string

// List of ObservabilityPipelineParseXMLProcessorType.
const (
	OBSERVABILITYPIPELINEPARSEXMLPROCESSORTYPE_PARSE_XML ObservabilityPipelineParseXMLProcessorType = "parse_xml"
)

var allowedObservabilityPipelineParseXMLProcessorTypeEnumValues = []ObservabilityPipelineParseXMLProcessorType{
	OBSERVABILITYPIPELINEPARSEXMLPROCESSORTYPE_PARSE_XML,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineParseXMLProcessorType) GetAllowedValues() []ObservabilityPipelineParseXMLProcessorType {
	return allowedObservabilityPipelineParseXMLProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineParseXMLProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineParseXMLProcessorType(value)
	return nil
}

// NewObservabilityPipelineParseXMLProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineParseXMLProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineParseXMLProcessorTypeFromValue(v string) (*ObservabilityPipelineParseXMLProcessorType, error) {
	ev := ObservabilityPipelineParseXMLProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineParseXMLProcessorType: valid values are %v", v, allowedObservabilityPipelineParseXMLProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineParseXMLProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineParseXMLProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineParseXMLProcessorType value.
func (v ObservabilityPipelineParseXMLProcessorType) Ptr() *ObservabilityPipelineParseXMLProcessorType {
	return &v
}
