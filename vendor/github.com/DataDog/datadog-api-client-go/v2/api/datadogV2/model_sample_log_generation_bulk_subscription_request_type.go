// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationBulkSubscriptionRequestType The type of the resource. The value should always be `bulk_subscription_requests`.
type SampleLogGenerationBulkSubscriptionRequestType string

// List of SampleLogGenerationBulkSubscriptionRequestType.
const (
	SAMPLELOGGENERATIONBULKSUBSCRIPTIONREQUESTTYPE_BULK_SUBSCRIPTION_REQUESTS SampleLogGenerationBulkSubscriptionRequestType = "bulk_subscription_requests"
)

var allowedSampleLogGenerationBulkSubscriptionRequestTypeEnumValues = []SampleLogGenerationBulkSubscriptionRequestType{
	SAMPLELOGGENERATIONBULKSUBSCRIPTIONREQUESTTYPE_BULK_SUBSCRIPTION_REQUESTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SampleLogGenerationBulkSubscriptionRequestType) GetAllowedValues() []SampleLogGenerationBulkSubscriptionRequestType {
	return allowedSampleLogGenerationBulkSubscriptionRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SampleLogGenerationBulkSubscriptionRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SampleLogGenerationBulkSubscriptionRequestType(value)
	return nil
}

// NewSampleLogGenerationBulkSubscriptionRequestTypeFromValue returns a pointer to a valid SampleLogGenerationBulkSubscriptionRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSampleLogGenerationBulkSubscriptionRequestTypeFromValue(v string) (*SampleLogGenerationBulkSubscriptionRequestType, error) {
	ev := SampleLogGenerationBulkSubscriptionRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SampleLogGenerationBulkSubscriptionRequestType: valid values are %v", v, allowedSampleLogGenerationBulkSubscriptionRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SampleLogGenerationBulkSubscriptionRequestType) IsValid() bool {
	for _, existing := range allowedSampleLogGenerationBulkSubscriptionRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SampleLogGenerationBulkSubscriptionRequestType value.
func (v SampleLogGenerationBulkSubscriptionRequestType) Ptr() *SampleLogGenerationBulkSubscriptionRequestType {
	return &v
}
