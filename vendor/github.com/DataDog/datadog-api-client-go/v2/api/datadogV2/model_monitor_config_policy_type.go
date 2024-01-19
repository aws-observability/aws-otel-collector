// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorConfigPolicyType The monitor configuration policy type.
type MonitorConfigPolicyType string

// List of MonitorConfigPolicyType.
const (
	MONITORCONFIGPOLICYTYPE_TAG MonitorConfigPolicyType = "tag"
)

var allowedMonitorConfigPolicyTypeEnumValues = []MonitorConfigPolicyType{
	MONITORCONFIGPOLICYTYPE_TAG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorConfigPolicyType) GetAllowedValues() []MonitorConfigPolicyType {
	return allowedMonitorConfigPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorConfigPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorConfigPolicyType(value)
	return nil
}

// NewMonitorConfigPolicyTypeFromValue returns a pointer to a valid MonitorConfigPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorConfigPolicyTypeFromValue(v string) (*MonitorConfigPolicyType, error) {
	ev := MonitorConfigPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorConfigPolicyType: valid values are %v", v, allowedMonitorConfigPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorConfigPolicyType) IsValid() bool {
	for _, existing := range allowedMonitorConfigPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorConfigPolicyType value.
func (v MonitorConfigPolicyType) Ptr() *MonitorConfigPolicyType {
	return &v
}
