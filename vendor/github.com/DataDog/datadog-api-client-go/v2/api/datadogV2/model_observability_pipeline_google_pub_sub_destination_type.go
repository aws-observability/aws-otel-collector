// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGooglePubSubDestinationType The destination type. The value should always be `google_pubsub`.
type ObservabilityPipelineGooglePubSubDestinationType string

// List of ObservabilityPipelineGooglePubSubDestinationType.
const (
	OBSERVABILITYPIPELINEGOOGLEPUBSUBDESTINATIONTYPE_GOOGLE_PUBSUB ObservabilityPipelineGooglePubSubDestinationType = "google_pubsub"
)

var allowedObservabilityPipelineGooglePubSubDestinationTypeEnumValues = []ObservabilityPipelineGooglePubSubDestinationType{
	OBSERVABILITYPIPELINEGOOGLEPUBSUBDESTINATIONTYPE_GOOGLE_PUBSUB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGooglePubSubDestinationType) GetAllowedValues() []ObservabilityPipelineGooglePubSubDestinationType {
	return allowedObservabilityPipelineGooglePubSubDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGooglePubSubDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGooglePubSubDestinationType(value)
	return nil
}

// NewObservabilityPipelineGooglePubSubDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineGooglePubSubDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGooglePubSubDestinationTypeFromValue(v string) (*ObservabilityPipelineGooglePubSubDestinationType, error) {
	ev := ObservabilityPipelineGooglePubSubDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGooglePubSubDestinationType: valid values are %v", v, allowedObservabilityPipelineGooglePubSubDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGooglePubSubDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGooglePubSubDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGooglePubSubDestinationType value.
func (v ObservabilityPipelineGooglePubSubDestinationType) Ptr() *ObservabilityPipelineGooglePubSubDestinationType {
	return &v
}
