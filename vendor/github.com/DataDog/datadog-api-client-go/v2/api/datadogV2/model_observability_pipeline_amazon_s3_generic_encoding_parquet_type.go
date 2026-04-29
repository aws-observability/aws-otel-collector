// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericEncodingParquetType The encoding type. Always `parquet`.
type ObservabilityPipelineAmazonS3GenericEncodingParquetType string

// List of ObservabilityPipelineAmazonS3GenericEncodingParquetType.
const (
	OBSERVABILITYPIPELINEAMAZONS3GENERICENCODINGPARQUETTYPE_PARQUET ObservabilityPipelineAmazonS3GenericEncodingParquetType = "parquet"
)

var allowedObservabilityPipelineAmazonS3GenericEncodingParquetTypeEnumValues = []ObservabilityPipelineAmazonS3GenericEncodingParquetType{
	OBSERVABILITYPIPELINEAMAZONS3GENERICENCODINGPARQUETTYPE_PARQUET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3GenericEncodingParquetType) GetAllowedValues() []ObservabilityPipelineAmazonS3GenericEncodingParquetType {
	return allowedObservabilityPipelineAmazonS3GenericEncodingParquetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3GenericEncodingParquetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3GenericEncodingParquetType(value)
	return nil
}

// NewObservabilityPipelineAmazonS3GenericEncodingParquetTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3GenericEncodingParquetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3GenericEncodingParquetTypeFromValue(v string) (*ObservabilityPipelineAmazonS3GenericEncodingParquetType, error) {
	ev := ObservabilityPipelineAmazonS3GenericEncodingParquetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3GenericEncodingParquetType: valid values are %v", v, allowedObservabilityPipelineAmazonS3GenericEncodingParquetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3GenericEncodingParquetType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3GenericEncodingParquetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3GenericEncodingParquetType value.
func (v ObservabilityPipelineAmazonS3GenericEncodingParquetType) Ptr() *ObservabilityPipelineAmazonS3GenericEncodingParquetType {
	return &v
}
