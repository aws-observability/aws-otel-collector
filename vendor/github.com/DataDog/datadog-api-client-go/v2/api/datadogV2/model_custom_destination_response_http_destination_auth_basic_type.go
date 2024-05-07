// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseHttpDestinationAuthBasicType Type of the basic access authentication.
type CustomDestinationResponseHttpDestinationAuthBasicType string

// List of CustomDestinationResponseHttpDestinationAuthBasicType.
const (
	CUSTOMDESTINATIONRESPONSEHTTPDESTINATIONAUTHBASICTYPE_BASIC CustomDestinationResponseHttpDestinationAuthBasicType = "basic"
)

var allowedCustomDestinationResponseHttpDestinationAuthBasicTypeEnumValues = []CustomDestinationResponseHttpDestinationAuthBasicType{
	CUSTOMDESTINATIONRESPONSEHTTPDESTINATIONAUTHBASICTYPE_BASIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationResponseHttpDestinationAuthBasicType) GetAllowedValues() []CustomDestinationResponseHttpDestinationAuthBasicType {
	return allowedCustomDestinationResponseHttpDestinationAuthBasicTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationResponseHttpDestinationAuthBasicType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationResponseHttpDestinationAuthBasicType(value)
	return nil
}

// NewCustomDestinationResponseHttpDestinationAuthBasicTypeFromValue returns a pointer to a valid CustomDestinationResponseHttpDestinationAuthBasicType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationResponseHttpDestinationAuthBasicTypeFromValue(v string) (*CustomDestinationResponseHttpDestinationAuthBasicType, error) {
	ev := CustomDestinationResponseHttpDestinationAuthBasicType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationResponseHttpDestinationAuthBasicType: valid values are %v", v, allowedCustomDestinationResponseHttpDestinationAuthBasicTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationResponseHttpDestinationAuthBasicType) IsValid() bool {
	for _, existing := range allowedCustomDestinationResponseHttpDestinationAuthBasicTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationResponseHttpDestinationAuthBasicType value.
func (v CustomDestinationResponseHttpDestinationAuthBasicType) Ptr() *CustomDestinationResponseHttpDestinationAuthBasicType {
	return &v
}
