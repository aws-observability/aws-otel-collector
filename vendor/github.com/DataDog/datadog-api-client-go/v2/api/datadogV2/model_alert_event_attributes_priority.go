// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertEventAttributesPriority The priority of the alert.
type AlertEventAttributesPriority string

// List of AlertEventAttributesPriority.
const (
	ALERTEVENTATTRIBUTESPRIORITY_PRIORITY_ONE   AlertEventAttributesPriority = "1"
	ALERTEVENTATTRIBUTESPRIORITY_PRIORITY_TWO   AlertEventAttributesPriority = "2"
	ALERTEVENTATTRIBUTESPRIORITY_PRIORITY_THREE AlertEventAttributesPriority = "3"
	ALERTEVENTATTRIBUTESPRIORITY_PRIORITY_FOUR  AlertEventAttributesPriority = "4"
	ALERTEVENTATTRIBUTESPRIORITY_PRIORITY_FIVE  AlertEventAttributesPriority = "5"
)

var allowedAlertEventAttributesPriorityEnumValues = []AlertEventAttributesPriority{
	ALERTEVENTATTRIBUTESPRIORITY_PRIORITY_ONE,
	ALERTEVENTATTRIBUTESPRIORITY_PRIORITY_TWO,
	ALERTEVENTATTRIBUTESPRIORITY_PRIORITY_THREE,
	ALERTEVENTATTRIBUTESPRIORITY_PRIORITY_FOUR,
	ALERTEVENTATTRIBUTESPRIORITY_PRIORITY_FIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertEventAttributesPriority) GetAllowedValues() []AlertEventAttributesPriority {
	return allowedAlertEventAttributesPriorityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertEventAttributesPriority) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertEventAttributesPriority(value)
	return nil
}

// NewAlertEventAttributesPriorityFromValue returns a pointer to a valid AlertEventAttributesPriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertEventAttributesPriorityFromValue(v string) (*AlertEventAttributesPriority, error) {
	ev := AlertEventAttributesPriority(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertEventAttributesPriority: valid values are %v", v, allowedAlertEventAttributesPriorityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertEventAttributesPriority) IsValid() bool {
	for _, existing := range allowedAlertEventAttributesPriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertEventAttributesPriority value.
func (v AlertEventAttributesPriority) Ptr() *AlertEventAttributesPriority {
	return &v
}
