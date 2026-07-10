// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationSubscriptionResourceType The type of the resource. The value should always be `subscriptions`.
type SampleLogGenerationSubscriptionResourceType string

// List of SampleLogGenerationSubscriptionResourceType.
const (
	SAMPLELOGGENERATIONSUBSCRIPTIONRESOURCETYPE_SUBSCRIPTIONS SampleLogGenerationSubscriptionResourceType = "subscriptions"
)

var allowedSampleLogGenerationSubscriptionResourceTypeEnumValues = []SampleLogGenerationSubscriptionResourceType{
	SAMPLELOGGENERATIONSUBSCRIPTIONRESOURCETYPE_SUBSCRIPTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SampleLogGenerationSubscriptionResourceType) GetAllowedValues() []SampleLogGenerationSubscriptionResourceType {
	return allowedSampleLogGenerationSubscriptionResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SampleLogGenerationSubscriptionResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SampleLogGenerationSubscriptionResourceType(value)
	return nil
}

// NewSampleLogGenerationSubscriptionResourceTypeFromValue returns a pointer to a valid SampleLogGenerationSubscriptionResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSampleLogGenerationSubscriptionResourceTypeFromValue(v string) (*SampleLogGenerationSubscriptionResourceType, error) {
	ev := SampleLogGenerationSubscriptionResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SampleLogGenerationSubscriptionResourceType: valid values are %v", v, allowedSampleLogGenerationSubscriptionResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SampleLogGenerationSubscriptionResourceType) IsValid() bool {
	for _, existing := range allowedSampleLogGenerationSubscriptionResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SampleLogGenerationSubscriptionResourceType value.
func (v SampleLogGenerationSubscriptionResourceType) Ptr() *SampleLogGenerationSubscriptionResourceType {
	return &v
}
