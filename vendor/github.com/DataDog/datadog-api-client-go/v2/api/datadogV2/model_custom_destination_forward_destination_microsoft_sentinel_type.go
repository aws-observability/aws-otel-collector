// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationForwardDestinationMicrosoftSentinelType Type of the Microsoft Sentinel destination.
type CustomDestinationForwardDestinationMicrosoftSentinelType string

// List of CustomDestinationForwardDestinationMicrosoftSentinelType.
const (
	CUSTOMDESTINATIONFORWARDDESTINATIONMICROSOFTSENTINELTYPE_MICROSOFT_SENTINEL CustomDestinationForwardDestinationMicrosoftSentinelType = "microsoft_sentinel"
)

var allowedCustomDestinationForwardDestinationMicrosoftSentinelTypeEnumValues = []CustomDestinationForwardDestinationMicrosoftSentinelType{
	CUSTOMDESTINATIONFORWARDDESTINATIONMICROSOFTSENTINELTYPE_MICROSOFT_SENTINEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationForwardDestinationMicrosoftSentinelType) GetAllowedValues() []CustomDestinationForwardDestinationMicrosoftSentinelType {
	return allowedCustomDestinationForwardDestinationMicrosoftSentinelTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationForwardDestinationMicrosoftSentinelType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationForwardDestinationMicrosoftSentinelType(value)
	return nil
}

// NewCustomDestinationForwardDestinationMicrosoftSentinelTypeFromValue returns a pointer to a valid CustomDestinationForwardDestinationMicrosoftSentinelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationForwardDestinationMicrosoftSentinelTypeFromValue(v string) (*CustomDestinationForwardDestinationMicrosoftSentinelType, error) {
	ev := CustomDestinationForwardDestinationMicrosoftSentinelType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationForwardDestinationMicrosoftSentinelType: valid values are %v", v, allowedCustomDestinationForwardDestinationMicrosoftSentinelTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationForwardDestinationMicrosoftSentinelType) IsValid() bool {
	for _, existing := range allowedCustomDestinationForwardDestinationMicrosoftSentinelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationForwardDestinationMicrosoftSentinelType value.
func (v CustomDestinationForwardDestinationMicrosoftSentinelType) Ptr() *CustomDestinationForwardDestinationMicrosoftSentinelType {
	return &v
}
