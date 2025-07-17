// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertEventCustomAttributesPriority The priority of the alert.
type AlertEventCustomAttributesPriority string

// List of AlertEventCustomAttributesPriority.
const (
	ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_ONE   AlertEventCustomAttributesPriority = "1"
	ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_TWO   AlertEventCustomAttributesPriority = "2"
	ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_THREE AlertEventCustomAttributesPriority = "3"
	ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_FOUR  AlertEventCustomAttributesPriority = "4"
	ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_FIVE  AlertEventCustomAttributesPriority = "5"
)

var allowedAlertEventCustomAttributesPriorityEnumValues = []AlertEventCustomAttributesPriority{
	ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_ONE,
	ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_TWO,
	ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_THREE,
	ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_FOUR,
	ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_FIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertEventCustomAttributesPriority) GetAllowedValues() []AlertEventCustomAttributesPriority {
	return allowedAlertEventCustomAttributesPriorityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertEventCustomAttributesPriority) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertEventCustomAttributesPriority(value)
	return nil
}

// NewAlertEventCustomAttributesPriorityFromValue returns a pointer to a valid AlertEventCustomAttributesPriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertEventCustomAttributesPriorityFromValue(v string) (*AlertEventCustomAttributesPriority, error) {
	ev := AlertEventCustomAttributesPriority(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertEventCustomAttributesPriority: valid values are %v", v, allowedAlertEventCustomAttributesPriorityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertEventCustomAttributesPriority) IsValid() bool {
	for _, existing := range allowedAlertEventCustomAttributesPriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertEventCustomAttributesPriority value.
func (v AlertEventCustomAttributesPriority) Ptr() *AlertEventCustomAttributesPriority {
	return &v
}
