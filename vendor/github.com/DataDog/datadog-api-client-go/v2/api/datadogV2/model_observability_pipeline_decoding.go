// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDecoding The decoding format used to interpret incoming logs.
type ObservabilityPipelineDecoding string

// List of ObservabilityPipelineDecoding.
const (
	OBSERVABILITYPIPELINEDECODING_DECODE_BYTES  ObservabilityPipelineDecoding = "bytes"
	OBSERVABILITYPIPELINEDECODING_DECODE_GELF   ObservabilityPipelineDecoding = "gelf"
	OBSERVABILITYPIPELINEDECODING_DECODE_JSON   ObservabilityPipelineDecoding = "json"
	OBSERVABILITYPIPELINEDECODING_DECODE_SYSLOG ObservabilityPipelineDecoding = "syslog"
)

var allowedObservabilityPipelineDecodingEnumValues = []ObservabilityPipelineDecoding{
	OBSERVABILITYPIPELINEDECODING_DECODE_BYTES,
	OBSERVABILITYPIPELINEDECODING_DECODE_GELF,
	OBSERVABILITYPIPELINEDECODING_DECODE_JSON,
	OBSERVABILITYPIPELINEDECODING_DECODE_SYSLOG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineDecoding) GetAllowedValues() []ObservabilityPipelineDecoding {
	return allowedObservabilityPipelineDecodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineDecoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineDecoding(value)
	return nil
}

// NewObservabilityPipelineDecodingFromValue returns a pointer to a valid ObservabilityPipelineDecoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineDecodingFromValue(v string) (*ObservabilityPipelineDecoding, error) {
	ev := ObservabilityPipelineDecoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineDecoding: valid values are %v", v, allowedObservabilityPipelineDecodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineDecoding) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineDecodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineDecoding value.
func (v ObservabilityPipelineDecoding) Ptr() *ObservabilityPipelineDecoding {
	return &v
}
