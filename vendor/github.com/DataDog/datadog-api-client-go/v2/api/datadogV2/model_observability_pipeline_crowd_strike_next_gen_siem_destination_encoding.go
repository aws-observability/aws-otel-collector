// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding Encoding format for log events.
type ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding string

// List of ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding.
const (
	OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONENCODING_JSON        ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding = "json"
	OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONENCODING_RAW_MESSAGE ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding = "raw_message"
)

var allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncodingEnumValues = []ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding{
	OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONENCODING_JSON,
	OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONENCODING_RAW_MESSAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding) GetAllowedValues() []ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding {
	return allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding(value)
	return nil
}

// NewObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncodingFromValue returns a pointer to a valid ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncodingFromValue(v string) (*ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding, error) {
	ev := ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding: valid values are %v", v, allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding value.
func (v ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding) Ptr() *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding {
	return &v
}
