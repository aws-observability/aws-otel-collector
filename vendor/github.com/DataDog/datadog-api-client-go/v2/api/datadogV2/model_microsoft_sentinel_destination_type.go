// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftSentinelDestinationType The destination type. The value should always be `microsoft_sentinel`.
type MicrosoftSentinelDestinationType string

// List of MicrosoftSentinelDestinationType.
const (
	MICROSOFTSENTINELDESTINATIONTYPE_MICROSOFT_SENTINEL MicrosoftSentinelDestinationType = "microsoft_sentinel"
)

var allowedMicrosoftSentinelDestinationTypeEnumValues = []MicrosoftSentinelDestinationType{
	MICROSOFTSENTINELDESTINATIONTYPE_MICROSOFT_SENTINEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MicrosoftSentinelDestinationType) GetAllowedValues() []MicrosoftSentinelDestinationType {
	return allowedMicrosoftSentinelDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MicrosoftSentinelDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MicrosoftSentinelDestinationType(value)
	return nil
}

// NewMicrosoftSentinelDestinationTypeFromValue returns a pointer to a valid MicrosoftSentinelDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMicrosoftSentinelDestinationTypeFromValue(v string) (*MicrosoftSentinelDestinationType, error) {
	ev := MicrosoftSentinelDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MicrosoftSentinelDestinationType: valid values are %v", v, allowedMicrosoftSentinelDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MicrosoftSentinelDestinationType) IsValid() bool {
	for _, existing := range allowedMicrosoftSentinelDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MicrosoftSentinelDestinationType value.
func (v MicrosoftSentinelDestinationType) Ptr() *MicrosoftSentinelDestinationType {
	return &v
}
