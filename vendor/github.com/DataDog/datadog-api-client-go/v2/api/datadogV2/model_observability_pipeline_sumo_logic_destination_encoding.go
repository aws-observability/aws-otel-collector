// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSumoLogicDestinationEncoding The output encoding format.
type ObservabilityPipelineSumoLogicDestinationEncoding string

// List of ObservabilityPipelineSumoLogicDestinationEncoding.
const (
	OBSERVABILITYPIPELINESUMOLOGICDESTINATIONENCODING_JSON        ObservabilityPipelineSumoLogicDestinationEncoding = "json"
	OBSERVABILITYPIPELINESUMOLOGICDESTINATIONENCODING_RAW_MESSAGE ObservabilityPipelineSumoLogicDestinationEncoding = "raw_message"
	OBSERVABILITYPIPELINESUMOLOGICDESTINATIONENCODING_LOGFMT      ObservabilityPipelineSumoLogicDestinationEncoding = "logfmt"
)

var allowedObservabilityPipelineSumoLogicDestinationEncodingEnumValues = []ObservabilityPipelineSumoLogicDestinationEncoding{
	OBSERVABILITYPIPELINESUMOLOGICDESTINATIONENCODING_JSON,
	OBSERVABILITYPIPELINESUMOLOGICDESTINATIONENCODING_RAW_MESSAGE,
	OBSERVABILITYPIPELINESUMOLOGICDESTINATIONENCODING_LOGFMT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSumoLogicDestinationEncoding) GetAllowedValues() []ObservabilityPipelineSumoLogicDestinationEncoding {
	return allowedObservabilityPipelineSumoLogicDestinationEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSumoLogicDestinationEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSumoLogicDestinationEncoding(value)
	return nil
}

// NewObservabilityPipelineSumoLogicDestinationEncodingFromValue returns a pointer to a valid ObservabilityPipelineSumoLogicDestinationEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSumoLogicDestinationEncodingFromValue(v string) (*ObservabilityPipelineSumoLogicDestinationEncoding, error) {
	ev := ObservabilityPipelineSumoLogicDestinationEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSumoLogicDestinationEncoding: valid values are %v", v, allowedObservabilityPipelineSumoLogicDestinationEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSumoLogicDestinationEncoding) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSumoLogicDestinationEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSumoLogicDestinationEncoding value.
func (v ObservabilityPipelineSumoLogicDestinationEncoding) Ptr() *ObservabilityPipelineSumoLogicDestinationEncoding {
	return &v
}
