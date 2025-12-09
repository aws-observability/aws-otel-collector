// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationHttpDestinationAuthCustomHeaderType Type of the custom header access authentication.
type CustomDestinationHttpDestinationAuthCustomHeaderType string

// List of CustomDestinationHttpDestinationAuthCustomHeaderType.
const (
	CUSTOMDESTINATIONHTTPDESTINATIONAUTHCUSTOMHEADERTYPE_CUSTOM_HEADER CustomDestinationHttpDestinationAuthCustomHeaderType = "custom_header"
)

var allowedCustomDestinationHttpDestinationAuthCustomHeaderTypeEnumValues = []CustomDestinationHttpDestinationAuthCustomHeaderType{
	CUSTOMDESTINATIONHTTPDESTINATIONAUTHCUSTOMHEADERTYPE_CUSTOM_HEADER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationHttpDestinationAuthCustomHeaderType) GetAllowedValues() []CustomDestinationHttpDestinationAuthCustomHeaderType {
	return allowedCustomDestinationHttpDestinationAuthCustomHeaderTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationHttpDestinationAuthCustomHeaderType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationHttpDestinationAuthCustomHeaderType(value)
	return nil
}

// NewCustomDestinationHttpDestinationAuthCustomHeaderTypeFromValue returns a pointer to a valid CustomDestinationHttpDestinationAuthCustomHeaderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationHttpDestinationAuthCustomHeaderTypeFromValue(v string) (*CustomDestinationHttpDestinationAuthCustomHeaderType, error) {
	ev := CustomDestinationHttpDestinationAuthCustomHeaderType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationHttpDestinationAuthCustomHeaderType: valid values are %v", v, allowedCustomDestinationHttpDestinationAuthCustomHeaderTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationHttpDestinationAuthCustomHeaderType) IsValid() bool {
	for _, existing := range allowedCustomDestinationHttpDestinationAuthCustomHeaderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationHttpDestinationAuthCustomHeaderType value.
func (v CustomDestinationHttpDestinationAuthCustomHeaderType) Ptr() *CustomDestinationHttpDestinationAuthCustomHeaderType {
	return &v
}
