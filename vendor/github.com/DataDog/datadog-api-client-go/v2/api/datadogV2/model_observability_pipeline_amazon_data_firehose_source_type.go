// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonDataFirehoseSourceType The source type. The value should always be `amazon_data_firehose`.
type ObservabilityPipelineAmazonDataFirehoseSourceType string

// List of ObservabilityPipelineAmazonDataFirehoseSourceType.
const (
	OBSERVABILITYPIPELINEAMAZONDATAFIREHOSESOURCETYPE_AMAZON_DATA_FIREHOSE ObservabilityPipelineAmazonDataFirehoseSourceType = "amazon_data_firehose"
)

var allowedObservabilityPipelineAmazonDataFirehoseSourceTypeEnumValues = []ObservabilityPipelineAmazonDataFirehoseSourceType{
	OBSERVABILITYPIPELINEAMAZONDATAFIREHOSESOURCETYPE_AMAZON_DATA_FIREHOSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonDataFirehoseSourceType) GetAllowedValues() []ObservabilityPipelineAmazonDataFirehoseSourceType {
	return allowedObservabilityPipelineAmazonDataFirehoseSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonDataFirehoseSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonDataFirehoseSourceType(value)
	return nil
}

// NewObservabilityPipelineAmazonDataFirehoseSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonDataFirehoseSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonDataFirehoseSourceTypeFromValue(v string) (*ObservabilityPipelineAmazonDataFirehoseSourceType, error) {
	ev := ObservabilityPipelineAmazonDataFirehoseSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonDataFirehoseSourceType: valid values are %v", v, allowedObservabilityPipelineAmazonDataFirehoseSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonDataFirehoseSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonDataFirehoseSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonDataFirehoseSourceType value.
func (v ObservabilityPipelineAmazonDataFirehoseSourceType) Ptr() *ObservabilityPipelineAmazonDataFirehoseSourceType {
	return &v
}
