// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationHttpDestinationAuthBasicType Type of the basic access authentication.
type CustomDestinationHttpDestinationAuthBasicType string

// List of CustomDestinationHttpDestinationAuthBasicType.
const (
	CUSTOMDESTINATIONHTTPDESTINATIONAUTHBASICTYPE_BASIC CustomDestinationHttpDestinationAuthBasicType = "basic"
)

var allowedCustomDestinationHttpDestinationAuthBasicTypeEnumValues = []CustomDestinationHttpDestinationAuthBasicType{
	CUSTOMDESTINATIONHTTPDESTINATIONAUTHBASICTYPE_BASIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationHttpDestinationAuthBasicType) GetAllowedValues() []CustomDestinationHttpDestinationAuthBasicType {
	return allowedCustomDestinationHttpDestinationAuthBasicTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationHttpDestinationAuthBasicType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationHttpDestinationAuthBasicType(value)
	return nil
}

// NewCustomDestinationHttpDestinationAuthBasicTypeFromValue returns a pointer to a valid CustomDestinationHttpDestinationAuthBasicType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationHttpDestinationAuthBasicTypeFromValue(v string) (*CustomDestinationHttpDestinationAuthBasicType, error) {
	ev := CustomDestinationHttpDestinationAuthBasicType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationHttpDestinationAuthBasicType: valid values are %v", v, allowedCustomDestinationHttpDestinationAuthBasicTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationHttpDestinationAuthBasicType) IsValid() bool {
	for _, existing := range allowedCustomDestinationHttpDestinationAuthBasicTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationHttpDestinationAuthBasicType value.
func (v CustomDestinationHttpDestinationAuthBasicType) Ptr() *CustomDestinationHttpDestinationAuthBasicType {
	return &v
}
