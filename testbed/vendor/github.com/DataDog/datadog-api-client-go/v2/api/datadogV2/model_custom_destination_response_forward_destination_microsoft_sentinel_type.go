// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseForwardDestinationMicrosoftSentinelType Type of the Microsoft Sentinel destination.
type CustomDestinationResponseForwardDestinationMicrosoftSentinelType string

// List of CustomDestinationResponseForwardDestinationMicrosoftSentinelType.
const (
	CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONMICROSOFTSENTINELTYPE_MICROSOFT_SENTINEL CustomDestinationResponseForwardDestinationMicrosoftSentinelType = "microsoft_sentinel"
)

var allowedCustomDestinationResponseForwardDestinationMicrosoftSentinelTypeEnumValues = []CustomDestinationResponseForwardDestinationMicrosoftSentinelType{
	CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONMICROSOFTSENTINELTYPE_MICROSOFT_SENTINEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationResponseForwardDestinationMicrosoftSentinelType) GetAllowedValues() []CustomDestinationResponseForwardDestinationMicrosoftSentinelType {
	return allowedCustomDestinationResponseForwardDestinationMicrosoftSentinelTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationResponseForwardDestinationMicrosoftSentinelType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationResponseForwardDestinationMicrosoftSentinelType(value)
	return nil
}

// NewCustomDestinationResponseForwardDestinationMicrosoftSentinelTypeFromValue returns a pointer to a valid CustomDestinationResponseForwardDestinationMicrosoftSentinelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationResponseForwardDestinationMicrosoftSentinelTypeFromValue(v string) (*CustomDestinationResponseForwardDestinationMicrosoftSentinelType, error) {
	ev := CustomDestinationResponseForwardDestinationMicrosoftSentinelType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationResponseForwardDestinationMicrosoftSentinelType: valid values are %v", v, allowedCustomDestinationResponseForwardDestinationMicrosoftSentinelTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationResponseForwardDestinationMicrosoftSentinelType) IsValid() bool {
	for _, existing := range allowedCustomDestinationResponseForwardDestinationMicrosoftSentinelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationResponseForwardDestinationMicrosoftSentinelType value.
func (v CustomDestinationResponseForwardDestinationMicrosoftSentinelType) Ptr() *CustomDestinationResponseForwardDestinationMicrosoftSentinelType {
	return &v
}
