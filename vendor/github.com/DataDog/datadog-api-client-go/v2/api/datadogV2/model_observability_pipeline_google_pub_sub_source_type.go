// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGooglePubSubSourceType The source type. The value should always be `google_pubsub`.
type ObservabilityPipelineGooglePubSubSourceType string

// List of ObservabilityPipelineGooglePubSubSourceType.
const (
	OBSERVABILITYPIPELINEGOOGLEPUBSUBSOURCETYPE_GOOGLE_PUBSUB ObservabilityPipelineGooglePubSubSourceType = "google_pubsub"
)

var allowedObservabilityPipelineGooglePubSubSourceTypeEnumValues = []ObservabilityPipelineGooglePubSubSourceType{
	OBSERVABILITYPIPELINEGOOGLEPUBSUBSOURCETYPE_GOOGLE_PUBSUB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGooglePubSubSourceType) GetAllowedValues() []ObservabilityPipelineGooglePubSubSourceType {
	return allowedObservabilityPipelineGooglePubSubSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGooglePubSubSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGooglePubSubSourceType(value)
	return nil
}

// NewObservabilityPipelineGooglePubSubSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineGooglePubSubSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGooglePubSubSourceTypeFromValue(v string) (*ObservabilityPipelineGooglePubSubSourceType, error) {
	ev := ObservabilityPipelineGooglePubSubSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGooglePubSubSourceType: valid values are %v", v, allowedObservabilityPipelineGooglePubSubSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGooglePubSubSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGooglePubSubSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGooglePubSubSourceType value.
func (v ObservabilityPipelineGooglePubSubSourceType) Ptr() *ObservabilityPipelineGooglePubSubSourceType {
	return &v
}
