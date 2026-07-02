// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NetworkHealthInsightsType The resource type for network health insights. Always `network-health-insights`.
type NetworkHealthInsightsType string

// List of NetworkHealthInsightsType.
const (
	NETWORKHEALTHINSIGHTSTYPE_NETWORK_HEALTH_INSIGHTS NetworkHealthInsightsType = "network-health-insights"
)

var allowedNetworkHealthInsightsTypeEnumValues = []NetworkHealthInsightsType{
	NETWORKHEALTHINSIGHTSTYPE_NETWORK_HEALTH_INSIGHTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NetworkHealthInsightsType) GetAllowedValues() []NetworkHealthInsightsType {
	return allowedNetworkHealthInsightsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NetworkHealthInsightsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NetworkHealthInsightsType(value)
	return nil
}

// NewNetworkHealthInsightsTypeFromValue returns a pointer to a valid NetworkHealthInsightsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNetworkHealthInsightsTypeFromValue(v string) (*NetworkHealthInsightsType, error) {
	ev := NetworkHealthInsightsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkHealthInsightsType: valid values are %v", v, allowedNetworkHealthInsightsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NetworkHealthInsightsType) IsValid() bool {
	for _, existing := range allowedNetworkHealthInsightsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkHealthInsightsType value.
func (v NetworkHealthInsightsType) Ptr() *NetworkHealthInsightsType {
	return &v
}
