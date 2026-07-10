// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryRuleStatusResponseDataType Custom allocation rule status resource type.
type ArbitraryRuleStatusResponseDataType string

// List of ArbitraryRuleStatusResponseDataType.
const (
	ARBITRARYRULESTATUSRESPONSEDATATYPE_ARBITRARY_RULE_STATUS ArbitraryRuleStatusResponseDataType = "arbitrary_rule_status"
)

var allowedArbitraryRuleStatusResponseDataTypeEnumValues = []ArbitraryRuleStatusResponseDataType{
	ARBITRARYRULESTATUSRESPONSEDATATYPE_ARBITRARY_RULE_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ArbitraryRuleStatusResponseDataType) GetAllowedValues() []ArbitraryRuleStatusResponseDataType {
	return allowedArbitraryRuleStatusResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ArbitraryRuleStatusResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ArbitraryRuleStatusResponseDataType(value)
	return nil
}

// NewArbitraryRuleStatusResponseDataTypeFromValue returns a pointer to a valid ArbitraryRuleStatusResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewArbitraryRuleStatusResponseDataTypeFromValue(v string) (*ArbitraryRuleStatusResponseDataType, error) {
	ev := ArbitraryRuleStatusResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ArbitraryRuleStatusResponseDataType: valid values are %v", v, allowedArbitraryRuleStatusResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ArbitraryRuleStatusResponseDataType) IsValid() bool {
	for _, existing := range allowedArbitraryRuleStatusResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ArbitraryRuleStatusResponseDataType value.
func (v ArbitraryRuleStatusResponseDataType) Ptr() *ArbitraryRuleStatusResponseDataType {
	return &v
}
