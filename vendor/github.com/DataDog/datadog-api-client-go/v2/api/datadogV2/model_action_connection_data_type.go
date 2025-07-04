// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionConnectionDataType The definition of `ActionConnectionDataType` object.
type ActionConnectionDataType string

// List of ActionConnectionDataType.
const (
	ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION ActionConnectionDataType = "action_connection"
)

var allowedActionConnectionDataTypeEnumValues = []ActionConnectionDataType{
	ACTIONCONNECTIONDATATYPE_ACTION_CONNECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ActionConnectionDataType) GetAllowedValues() []ActionConnectionDataType {
	return allowedActionConnectionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ActionConnectionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ActionConnectionDataType(value)
	return nil
}

// NewActionConnectionDataTypeFromValue returns a pointer to a valid ActionConnectionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewActionConnectionDataTypeFromValue(v string) (*ActionConnectionDataType, error) {
	ev := ActionConnectionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ActionConnectionDataType: valid values are %v", v, allowedActionConnectionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ActionConnectionDataType) IsValid() bool {
	for _, existing := range allowedActionConnectionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActionConnectionDataType value.
func (v ActionConnectionDataType) Ptr() *ActionConnectionDataType {
	return &v
}
