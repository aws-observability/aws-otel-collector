// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerSource The type of security issues on which the rule applies. Notification rules based on security signals need to use the trigger source "security_signals",
// while notification rules based on security vulnerabilities need to use the trigger source "security_findings".
type TriggerSource string

// List of TriggerSource.
const (
	TRIGGERSOURCE_SECURITY_FINDINGS TriggerSource = "security_findings"
	TRIGGERSOURCE_SECURITY_SIGNALS  TriggerSource = "security_signals"
)

var allowedTriggerSourceEnumValues = []TriggerSource{
	TRIGGERSOURCE_SECURITY_FINDINGS,
	TRIGGERSOURCE_SECURITY_SIGNALS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TriggerSource) GetAllowedValues() []TriggerSource {
	return allowedTriggerSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TriggerSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TriggerSource(value)
	return nil
}

// NewTriggerSourceFromValue returns a pointer to a valid TriggerSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTriggerSourceFromValue(v string) (*TriggerSource, error) {
	ev := TriggerSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TriggerSource: valid values are %v", v, allowedTriggerSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TriggerSource) IsValid() bool {
	for _, existing := range allowedTriggerSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerSource value.
func (v TriggerSource) Ptr() *TriggerSource {
	return &v
}
