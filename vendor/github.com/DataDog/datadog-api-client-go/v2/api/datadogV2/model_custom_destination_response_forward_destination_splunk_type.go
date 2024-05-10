// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseForwardDestinationSplunkType Type of the Splunk HTTP Event Collector (HEC) destination.
type CustomDestinationResponseForwardDestinationSplunkType string

// List of CustomDestinationResponseForwardDestinationSplunkType.
const (
	CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONSPLUNKTYPE_SPLUNK_HEC CustomDestinationResponseForwardDestinationSplunkType = "splunk_hec"
)

var allowedCustomDestinationResponseForwardDestinationSplunkTypeEnumValues = []CustomDestinationResponseForwardDestinationSplunkType{
	CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONSPLUNKTYPE_SPLUNK_HEC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationResponseForwardDestinationSplunkType) GetAllowedValues() []CustomDestinationResponseForwardDestinationSplunkType {
	return allowedCustomDestinationResponseForwardDestinationSplunkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationResponseForwardDestinationSplunkType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationResponseForwardDestinationSplunkType(value)
	return nil
}

// NewCustomDestinationResponseForwardDestinationSplunkTypeFromValue returns a pointer to a valid CustomDestinationResponseForwardDestinationSplunkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationResponseForwardDestinationSplunkTypeFromValue(v string) (*CustomDestinationResponseForwardDestinationSplunkType, error) {
	ev := CustomDestinationResponseForwardDestinationSplunkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationResponseForwardDestinationSplunkType: valid values are %v", v, allowedCustomDestinationResponseForwardDestinationSplunkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationResponseForwardDestinationSplunkType) IsValid() bool {
	for _, existing := range allowedCustomDestinationResponseForwardDestinationSplunkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationResponseForwardDestinationSplunkType value.
func (v CustomDestinationResponseForwardDestinationSplunkType) Ptr() *CustomDestinationResponseForwardDestinationSplunkType {
	return &v
}
