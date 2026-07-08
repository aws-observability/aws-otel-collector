// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationSubscriptionRequestType The type of the resource. The value should always be `subscription_requests`.
type SampleLogGenerationSubscriptionRequestType string

// List of SampleLogGenerationSubscriptionRequestType.
const (
	SAMPLELOGGENERATIONSUBSCRIPTIONREQUESTTYPE_SUBSCRIPTION_REQUESTS SampleLogGenerationSubscriptionRequestType = "subscription_requests"
)

var allowedSampleLogGenerationSubscriptionRequestTypeEnumValues = []SampleLogGenerationSubscriptionRequestType{
	SAMPLELOGGENERATIONSUBSCRIPTIONREQUESTTYPE_SUBSCRIPTION_REQUESTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SampleLogGenerationSubscriptionRequestType) GetAllowedValues() []SampleLogGenerationSubscriptionRequestType {
	return allowedSampleLogGenerationSubscriptionRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SampleLogGenerationSubscriptionRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SampleLogGenerationSubscriptionRequestType(value)
	return nil
}

// NewSampleLogGenerationSubscriptionRequestTypeFromValue returns a pointer to a valid SampleLogGenerationSubscriptionRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSampleLogGenerationSubscriptionRequestTypeFromValue(v string) (*SampleLogGenerationSubscriptionRequestType, error) {
	ev := SampleLogGenerationSubscriptionRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SampleLogGenerationSubscriptionRequestType: valid values are %v", v, allowedSampleLogGenerationSubscriptionRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SampleLogGenerationSubscriptionRequestType) IsValid() bool {
	for _, existing := range allowedSampleLogGenerationSubscriptionRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SampleLogGenerationSubscriptionRequestType value.
func (v SampleLogGenerationSubscriptionRequestType) Ptr() *SampleLogGenerationSubscriptionRequestType {
	return &v
}
