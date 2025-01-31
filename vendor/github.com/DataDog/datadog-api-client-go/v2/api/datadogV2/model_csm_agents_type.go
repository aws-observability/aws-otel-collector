// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CSMAgentsType The type of the resource. The value should always be `datadog_agent`.
type CSMAgentsType string

// List of CSMAgentsType.
const (
	CSMAGENTSTYPE_DATADOG_AGENT CSMAgentsType = "datadog_agent"
)

var allowedCSMAgentsTypeEnumValues = []CSMAgentsType{
	CSMAGENTSTYPE_DATADOG_AGENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CSMAgentsType) GetAllowedValues() []CSMAgentsType {
	return allowedCSMAgentsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CSMAgentsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CSMAgentsType(value)
	return nil
}

// NewCSMAgentsTypeFromValue returns a pointer to a valid CSMAgentsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCSMAgentsTypeFromValue(v string) (*CSMAgentsType, error) {
	ev := CSMAgentsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CSMAgentsType: valid values are %v", v, allowedCSMAgentsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CSMAgentsType) IsValid() bool {
	for _, existing := range allowedCSMAgentsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CSMAgentsType value.
func (v CSMAgentsType) Ptr() *CSMAgentsType {
	return &v
}
