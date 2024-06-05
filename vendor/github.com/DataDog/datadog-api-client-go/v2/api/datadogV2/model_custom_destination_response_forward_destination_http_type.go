// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseForwardDestinationHttpType Type of the HTTP destination.
type CustomDestinationResponseForwardDestinationHttpType string

// List of CustomDestinationResponseForwardDestinationHttpType.
const (
	CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONHTTPTYPE_HTTP CustomDestinationResponseForwardDestinationHttpType = "http"
)

var allowedCustomDestinationResponseForwardDestinationHttpTypeEnumValues = []CustomDestinationResponseForwardDestinationHttpType{
	CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONHTTPTYPE_HTTP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationResponseForwardDestinationHttpType) GetAllowedValues() []CustomDestinationResponseForwardDestinationHttpType {
	return allowedCustomDestinationResponseForwardDestinationHttpTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationResponseForwardDestinationHttpType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationResponseForwardDestinationHttpType(value)
	return nil
}

// NewCustomDestinationResponseForwardDestinationHttpTypeFromValue returns a pointer to a valid CustomDestinationResponseForwardDestinationHttpType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationResponseForwardDestinationHttpTypeFromValue(v string) (*CustomDestinationResponseForwardDestinationHttpType, error) {
	ev := CustomDestinationResponseForwardDestinationHttpType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationResponseForwardDestinationHttpType: valid values are %v", v, allowedCustomDestinationResponseForwardDestinationHttpTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationResponseForwardDestinationHttpType) IsValid() bool {
	for _, existing := range allowedCustomDestinationResponseForwardDestinationHttpTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationResponseForwardDestinationHttpType value.
func (v CustomDestinationResponseForwardDestinationHttpType) Ptr() *CustomDestinationResponseForwardDestinationHttpType {
	return &v
}
