// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NetworkHealthInsightFailureType Specific failure type within the insight category. For DNS insights: `timeout`, `nxdomain`,
// `servfail`, or `general_failure`. For TLS certificate insights: `expired` or `expiring_soon`.
// For security group insights: `denied`.
type NetworkHealthInsightFailureType string

// List of NetworkHealthInsightFailureType.
const (
	NETWORKHEALTHINSIGHTFAILURETYPE_TIMEOUT         NetworkHealthInsightFailureType = "timeout"
	NETWORKHEALTHINSIGHTFAILURETYPE_NXDOMAIN        NetworkHealthInsightFailureType = "nxdomain"
	NETWORKHEALTHINSIGHTFAILURETYPE_SERVFAIL        NetworkHealthInsightFailureType = "servfail"
	NETWORKHEALTHINSIGHTFAILURETYPE_GENERAL_FAILURE NetworkHealthInsightFailureType = "general_failure"
	NETWORKHEALTHINSIGHTFAILURETYPE_EXPIRED         NetworkHealthInsightFailureType = "expired"
	NETWORKHEALTHINSIGHTFAILURETYPE_EXPIRING_SOON   NetworkHealthInsightFailureType = "expiring_soon"
	NETWORKHEALTHINSIGHTFAILURETYPE_DENIED          NetworkHealthInsightFailureType = "denied"
)

var allowedNetworkHealthInsightFailureTypeEnumValues = []NetworkHealthInsightFailureType{
	NETWORKHEALTHINSIGHTFAILURETYPE_TIMEOUT,
	NETWORKHEALTHINSIGHTFAILURETYPE_NXDOMAIN,
	NETWORKHEALTHINSIGHTFAILURETYPE_SERVFAIL,
	NETWORKHEALTHINSIGHTFAILURETYPE_GENERAL_FAILURE,
	NETWORKHEALTHINSIGHTFAILURETYPE_EXPIRED,
	NETWORKHEALTHINSIGHTFAILURETYPE_EXPIRING_SOON,
	NETWORKHEALTHINSIGHTFAILURETYPE_DENIED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NetworkHealthInsightFailureType) GetAllowedValues() []NetworkHealthInsightFailureType {
	return allowedNetworkHealthInsightFailureTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NetworkHealthInsightFailureType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NetworkHealthInsightFailureType(value)
	return nil
}

// NewNetworkHealthInsightFailureTypeFromValue returns a pointer to a valid NetworkHealthInsightFailureType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNetworkHealthInsightFailureTypeFromValue(v string) (*NetworkHealthInsightFailureType, error) {
	ev := NetworkHealthInsightFailureType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkHealthInsightFailureType: valid values are %v", v, allowedNetworkHealthInsightFailureTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NetworkHealthInsightFailureType) IsValid() bool {
	for _, existing := range allowedNetworkHealthInsightFailureTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkHealthInsightFailureType value.
func (v NetworkHealthInsightFailureType) Ptr() *NetworkHealthInsightFailureType {
	return &v
}
