// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationForwardDestinationSplunkType Type of the Splunk HTTP Event Collector (HEC) destination.
type CustomDestinationForwardDestinationSplunkType string

// List of CustomDestinationForwardDestinationSplunkType.
const (
	CUSTOMDESTINATIONFORWARDDESTINATIONSPLUNKTYPE_SPLUNK_HEC CustomDestinationForwardDestinationSplunkType = "splunk_hec"
)

var allowedCustomDestinationForwardDestinationSplunkTypeEnumValues = []CustomDestinationForwardDestinationSplunkType{
	CUSTOMDESTINATIONFORWARDDESTINATIONSPLUNKTYPE_SPLUNK_HEC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationForwardDestinationSplunkType) GetAllowedValues() []CustomDestinationForwardDestinationSplunkType {
	return allowedCustomDestinationForwardDestinationSplunkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationForwardDestinationSplunkType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationForwardDestinationSplunkType(value)
	return nil
}

// NewCustomDestinationForwardDestinationSplunkTypeFromValue returns a pointer to a valid CustomDestinationForwardDestinationSplunkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationForwardDestinationSplunkTypeFromValue(v string) (*CustomDestinationForwardDestinationSplunkType, error) {
	ev := CustomDestinationForwardDestinationSplunkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationForwardDestinationSplunkType: valid values are %v", v, allowedCustomDestinationForwardDestinationSplunkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationForwardDestinationSplunkType) IsValid() bool {
	for _, existing := range allowedCustomDestinationForwardDestinationSplunkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationForwardDestinationSplunkType value.
func (v CustomDestinationForwardDestinationSplunkType) Ptr() *CustomDestinationForwardDestinationSplunkType {
	return &v
}
