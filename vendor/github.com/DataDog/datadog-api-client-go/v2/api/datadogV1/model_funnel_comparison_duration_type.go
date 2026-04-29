// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelComparisonDurationType Type of comparison duration.
type FunnelComparisonDurationType string

// List of FunnelComparisonDurationType.
const (
	FUNNELCOMPARISONDURATIONTYPE_PREVIOUS_TIMEFRAME FunnelComparisonDurationType = "previous_timeframe"
	FUNNELCOMPARISONDURATIONTYPE_CUSTOM_TIMEFRAME   FunnelComparisonDurationType = "custom_timeframe"
	FUNNELCOMPARISONDURATIONTYPE_PREVIOUS_DAY       FunnelComparisonDurationType = "previous_day"
	FUNNELCOMPARISONDURATIONTYPE_PREVIOUS_WEEK      FunnelComparisonDurationType = "previous_week"
	FUNNELCOMPARISONDURATIONTYPE_PREVIOUS_MONTH     FunnelComparisonDurationType = "previous_month"
)

var allowedFunnelComparisonDurationTypeEnumValues = []FunnelComparisonDurationType{
	FUNNELCOMPARISONDURATIONTYPE_PREVIOUS_TIMEFRAME,
	FUNNELCOMPARISONDURATIONTYPE_CUSTOM_TIMEFRAME,
	FUNNELCOMPARISONDURATIONTYPE_PREVIOUS_DAY,
	FUNNELCOMPARISONDURATIONTYPE_PREVIOUS_WEEK,
	FUNNELCOMPARISONDURATIONTYPE_PREVIOUS_MONTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FunnelComparisonDurationType) GetAllowedValues() []FunnelComparisonDurationType {
	return allowedFunnelComparisonDurationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FunnelComparisonDurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelComparisonDurationType(value)
	return nil
}

// NewFunnelComparisonDurationTypeFromValue returns a pointer to a valid FunnelComparisonDurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFunnelComparisonDurationTypeFromValue(v string) (*FunnelComparisonDurationType, error) {
	ev := FunnelComparisonDurationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FunnelComparisonDurationType: valid values are %v", v, allowedFunnelComparisonDurationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FunnelComparisonDurationType) IsValid() bool {
	for _, existing := range allowedFunnelComparisonDurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelComparisonDurationType value.
func (v FunnelComparisonDurationType) Ptr() *FunnelComparisonDurationType {
	return &v
}
