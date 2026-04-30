// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericEncodingJsonType The encoding type. Always `json`.
type ObservabilityPipelineAmazonS3GenericEncodingJsonType string

// List of ObservabilityPipelineAmazonS3GenericEncodingJsonType.
const (
	OBSERVABILITYPIPELINEAMAZONS3GENERICENCODINGJSONTYPE_JSON ObservabilityPipelineAmazonS3GenericEncodingJsonType = "json"
)

var allowedObservabilityPipelineAmazonS3GenericEncodingJsonTypeEnumValues = []ObservabilityPipelineAmazonS3GenericEncodingJsonType{
	OBSERVABILITYPIPELINEAMAZONS3GENERICENCODINGJSONTYPE_JSON,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3GenericEncodingJsonType) GetAllowedValues() []ObservabilityPipelineAmazonS3GenericEncodingJsonType {
	return allowedObservabilityPipelineAmazonS3GenericEncodingJsonTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3GenericEncodingJsonType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3GenericEncodingJsonType(value)
	return nil
}

// NewObservabilityPipelineAmazonS3GenericEncodingJsonTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3GenericEncodingJsonType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3GenericEncodingJsonTypeFromValue(v string) (*ObservabilityPipelineAmazonS3GenericEncodingJsonType, error) {
	ev := ObservabilityPipelineAmazonS3GenericEncodingJsonType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3GenericEncodingJsonType: valid values are %v", v, allowedObservabilityPipelineAmazonS3GenericEncodingJsonTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3GenericEncodingJsonType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3GenericEncodingJsonTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3GenericEncodingJsonType value.
func (v ObservabilityPipelineAmazonS3GenericEncodingJsonType) Ptr() *ObservabilityPipelineAmazonS3GenericEncodingJsonType {
	return &v
}
