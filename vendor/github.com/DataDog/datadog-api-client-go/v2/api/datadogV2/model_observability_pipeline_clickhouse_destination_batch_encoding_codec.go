// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineClickhouseDestinationBatchEncodingCodec The codec used for batch encoding. Only `arrow_stream` is supported.
type ObservabilityPipelineClickhouseDestinationBatchEncodingCodec string

// List of ObservabilityPipelineClickhouseDestinationBatchEncodingCodec.
const (
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONBATCHENCODINGCODEC_ARROW_STREAM ObservabilityPipelineClickhouseDestinationBatchEncodingCodec = "arrow_stream"
)

var allowedObservabilityPipelineClickhouseDestinationBatchEncodingCodecEnumValues = []ObservabilityPipelineClickhouseDestinationBatchEncodingCodec{
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONBATCHENCODINGCODEC_ARROW_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineClickhouseDestinationBatchEncodingCodec) GetAllowedValues() []ObservabilityPipelineClickhouseDestinationBatchEncodingCodec {
	return allowedObservabilityPipelineClickhouseDestinationBatchEncodingCodecEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineClickhouseDestinationBatchEncodingCodec) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineClickhouseDestinationBatchEncodingCodec(value)
	return nil
}

// NewObservabilityPipelineClickhouseDestinationBatchEncodingCodecFromValue returns a pointer to a valid ObservabilityPipelineClickhouseDestinationBatchEncodingCodec
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineClickhouseDestinationBatchEncodingCodecFromValue(v string) (*ObservabilityPipelineClickhouseDestinationBatchEncodingCodec, error) {
	ev := ObservabilityPipelineClickhouseDestinationBatchEncodingCodec(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineClickhouseDestinationBatchEncodingCodec: valid values are %v", v, allowedObservabilityPipelineClickhouseDestinationBatchEncodingCodecEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineClickhouseDestinationBatchEncodingCodec) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineClickhouseDestinationBatchEncodingCodecEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineClickhouseDestinationBatchEncodingCodec value.
func (v ObservabilityPipelineClickhouseDestinationBatchEncodingCodec) Ptr() *ObservabilityPipelineClickhouseDestinationBatchEncodingCodec {
	return &v
}
