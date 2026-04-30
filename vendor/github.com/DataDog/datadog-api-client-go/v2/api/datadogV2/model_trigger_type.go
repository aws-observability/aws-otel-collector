// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerType The type of trigger for the investigation.
type TriggerType string

// List of TriggerType.
const (
	TRIGGERTYPE_MONITOR_ALERT_TRIGGER TriggerType = "monitor_alert_trigger"
)

var allowedTriggerTypeEnumValues = []TriggerType{
	TRIGGERTYPE_MONITOR_ALERT_TRIGGER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TriggerType) GetAllowedValues() []TriggerType {
	return allowedTriggerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TriggerType(value)
	return nil
}

// NewTriggerTypeFromValue returns a pointer to a valid TriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTriggerTypeFromValue(v string) (*TriggerType, error) {
	ev := TriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TriggerType: valid values are %v", v, allowedTriggerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TriggerType) IsValid() bool {
	for _, existing := range allowedTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerType value.
func (v TriggerType) Ptr() *TriggerType {
	return &v
}
