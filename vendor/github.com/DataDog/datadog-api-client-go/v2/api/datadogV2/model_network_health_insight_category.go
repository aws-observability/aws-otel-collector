// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NetworkHealthInsightCategory Category of network health insight. Indicates whether the insight relates to a DNS issue (`dns`),
// a TCP issue (`tcp`), a TLS certificate issue (`tls-cert`), or a security group denial (`security-group`).
type NetworkHealthInsightCategory string

// List of NetworkHealthInsightCategory.
const (
	NETWORKHEALTHINSIGHTCATEGORY_DNS            NetworkHealthInsightCategory = "dns"
	NETWORKHEALTHINSIGHTCATEGORY_TCP            NetworkHealthInsightCategory = "tcp"
	NETWORKHEALTHINSIGHTCATEGORY_TLS_CERT       NetworkHealthInsightCategory = "tls-cert"
	NETWORKHEALTHINSIGHTCATEGORY_SECURITY_GROUP NetworkHealthInsightCategory = "security-group"
)

var allowedNetworkHealthInsightCategoryEnumValues = []NetworkHealthInsightCategory{
	NETWORKHEALTHINSIGHTCATEGORY_DNS,
	NETWORKHEALTHINSIGHTCATEGORY_TCP,
	NETWORKHEALTHINSIGHTCATEGORY_TLS_CERT,
	NETWORKHEALTHINSIGHTCATEGORY_SECURITY_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NetworkHealthInsightCategory) GetAllowedValues() []NetworkHealthInsightCategory {
	return allowedNetworkHealthInsightCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NetworkHealthInsightCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NetworkHealthInsightCategory(value)
	return nil
}

// NewNetworkHealthInsightCategoryFromValue returns a pointer to a valid NetworkHealthInsightCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNetworkHealthInsightCategoryFromValue(v string) (*NetworkHealthInsightCategory, error) {
	ev := NetworkHealthInsightCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkHealthInsightCategory: valid values are %v", v, allowedNetworkHealthInsightCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NetworkHealthInsightCategory) IsValid() bool {
	for _, existing := range allowedNetworkHealthInsightCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkHealthInsightCategory value.
func (v NetworkHealthInsightCategory) Ptr() *NetworkHealthInsightCategory {
	return &v
}
