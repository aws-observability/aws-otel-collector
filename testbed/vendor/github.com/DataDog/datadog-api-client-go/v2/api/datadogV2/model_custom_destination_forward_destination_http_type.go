// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationForwardDestinationHttpType Type of the HTTP destination.
type CustomDestinationForwardDestinationHttpType string

// List of CustomDestinationForwardDestinationHttpType.
const (
	CUSTOMDESTINATIONFORWARDDESTINATIONHTTPTYPE_HTTP CustomDestinationForwardDestinationHttpType = "http"
)

var allowedCustomDestinationForwardDestinationHttpTypeEnumValues = []CustomDestinationForwardDestinationHttpType{
	CUSTOMDESTINATIONFORWARDDESTINATIONHTTPTYPE_HTTP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationForwardDestinationHttpType) GetAllowedValues() []CustomDestinationForwardDestinationHttpType {
	return allowedCustomDestinationForwardDestinationHttpTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationForwardDestinationHttpType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationForwardDestinationHttpType(value)
	return nil
}

// NewCustomDestinationForwardDestinationHttpTypeFromValue returns a pointer to a valid CustomDestinationForwardDestinationHttpType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationForwardDestinationHttpTypeFromValue(v string) (*CustomDestinationForwardDestinationHttpType, error) {
	ev := CustomDestinationForwardDestinationHttpType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationForwardDestinationHttpType: valid values are %v", v, allowedCustomDestinationForwardDestinationHttpTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationForwardDestinationHttpType) IsValid() bool {
	for _, existing := range allowedCustomDestinationForwardDestinationHttpTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationForwardDestinationHttpType value.
func (v CustomDestinationForwardDestinationHttpType) Ptr() *CustomDestinationForwardDestinationHttpType {
	return &v
}
