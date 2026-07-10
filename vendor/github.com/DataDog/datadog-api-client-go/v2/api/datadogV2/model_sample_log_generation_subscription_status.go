// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationSubscriptionStatus The status of the subscription.
type SampleLogGenerationSubscriptionStatus string

// List of SampleLogGenerationSubscriptionStatus.
const (
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_SUBSCRIBED             SampleLogGenerationSubscriptionStatus = "subscribed"
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_RENEWED                SampleLogGenerationSubscriptionStatus = "renewed"
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_UNSUBSCRIBED           SampleLogGenerationSubscriptionStatus = "unsubscribed"
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_NO_ACTIVE_SUBSCRIPTION SampleLogGenerationSubscriptionStatus = "no_active_subscription"
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_NOT_AVAILABLE          SampleLogGenerationSubscriptionStatus = "not_available"
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_ACTIVE                 SampleLogGenerationSubscriptionStatus = "active"
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_EXPIRED                SampleLogGenerationSubscriptionStatus = "expired"
)

var allowedSampleLogGenerationSubscriptionStatusEnumValues = []SampleLogGenerationSubscriptionStatus{
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_SUBSCRIBED,
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_RENEWED,
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_UNSUBSCRIBED,
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_NO_ACTIVE_SUBSCRIPTION,
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_NOT_AVAILABLE,
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_ACTIVE,
	SAMPLELOGGENERATIONSUBSCRIPTIONSTATUS_EXPIRED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SampleLogGenerationSubscriptionStatus) GetAllowedValues() []SampleLogGenerationSubscriptionStatus {
	return allowedSampleLogGenerationSubscriptionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SampleLogGenerationSubscriptionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SampleLogGenerationSubscriptionStatus(value)
	return nil
}

// NewSampleLogGenerationSubscriptionStatusFromValue returns a pointer to a valid SampleLogGenerationSubscriptionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSampleLogGenerationSubscriptionStatusFromValue(v string) (*SampleLogGenerationSubscriptionStatus, error) {
	ev := SampleLogGenerationSubscriptionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SampleLogGenerationSubscriptionStatus: valid values are %v", v, allowedSampleLogGenerationSubscriptionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SampleLogGenerationSubscriptionStatus) IsValid() bool {
	for _, existing := range allowedSampleLogGenerationSubscriptionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SampleLogGenerationSubscriptionStatus value.
func (v SampleLogGenerationSubscriptionStatus) Ptr() *SampleLogGenerationSubscriptionStatus {
	return &v
}
