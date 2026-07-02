// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationSubscriptionsStatusFilter Filter that controls whether to return only active subscriptions or every subscription on record.
type SampleLogGenerationSubscriptionsStatusFilter string

// List of SampleLogGenerationSubscriptionsStatusFilter.
const (
	SAMPLELOGGENERATIONSUBSCRIPTIONSSTATUSFILTER_ACTIVE SampleLogGenerationSubscriptionsStatusFilter = "active"
	SAMPLELOGGENERATIONSUBSCRIPTIONSSTATUSFILTER_ALL    SampleLogGenerationSubscriptionsStatusFilter = "all"
)

var allowedSampleLogGenerationSubscriptionsStatusFilterEnumValues = []SampleLogGenerationSubscriptionsStatusFilter{
	SAMPLELOGGENERATIONSUBSCRIPTIONSSTATUSFILTER_ACTIVE,
	SAMPLELOGGENERATIONSUBSCRIPTIONSSTATUSFILTER_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SampleLogGenerationSubscriptionsStatusFilter) GetAllowedValues() []SampleLogGenerationSubscriptionsStatusFilter {
	return allowedSampleLogGenerationSubscriptionsStatusFilterEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SampleLogGenerationSubscriptionsStatusFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SampleLogGenerationSubscriptionsStatusFilter(value)
	return nil
}

// NewSampleLogGenerationSubscriptionsStatusFilterFromValue returns a pointer to a valid SampleLogGenerationSubscriptionsStatusFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSampleLogGenerationSubscriptionsStatusFilterFromValue(v string) (*SampleLogGenerationSubscriptionsStatusFilter, error) {
	ev := SampleLogGenerationSubscriptionsStatusFilter(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SampleLogGenerationSubscriptionsStatusFilter: valid values are %v", v, allowedSampleLogGenerationSubscriptionsStatusFilterEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SampleLogGenerationSubscriptionsStatusFilter) IsValid() bool {
	for _, existing := range allowedSampleLogGenerationSubscriptionsStatusFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SampleLogGenerationSubscriptionsStatusFilter value.
func (v SampleLogGenerationSubscriptionsStatusFilter) Ptr() *SampleLogGenerationSubscriptionsStatusFilter {
	return &v
}
