// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SingleAggregatedDnsResponseDataType Aggregated DNS resource type.
type SingleAggregatedDnsResponseDataType string

// List of SingleAggregatedDnsResponseDataType.
const (
	SINGLEAGGREGATEDDNSRESPONSEDATATYPE_AGGREGATED_DNS SingleAggregatedDnsResponseDataType = "aggregated_dns"
)

var allowedSingleAggregatedDnsResponseDataTypeEnumValues = []SingleAggregatedDnsResponseDataType{
	SINGLEAGGREGATEDDNSRESPONSEDATATYPE_AGGREGATED_DNS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SingleAggregatedDnsResponseDataType) GetAllowedValues() []SingleAggregatedDnsResponseDataType {
	return allowedSingleAggregatedDnsResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SingleAggregatedDnsResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SingleAggregatedDnsResponseDataType(value)
	return nil
}

// NewSingleAggregatedDnsResponseDataTypeFromValue returns a pointer to a valid SingleAggregatedDnsResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSingleAggregatedDnsResponseDataTypeFromValue(v string) (*SingleAggregatedDnsResponseDataType, error) {
	ev := SingleAggregatedDnsResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SingleAggregatedDnsResponseDataType: valid values are %v", v, allowedSingleAggregatedDnsResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SingleAggregatedDnsResponseDataType) IsValid() bool {
	for _, existing := range allowedSingleAggregatedDnsResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SingleAggregatedDnsResponseDataType value.
func (v SingleAggregatedDnsResponseDataType) Ptr() *SingleAggregatedDnsResponseDataType {
	return &v
}
