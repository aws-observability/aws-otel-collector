// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertEventAttributesStatus The status of the alert.
type AlertEventAttributesStatus string

// List of AlertEventAttributesStatus.
const (
	ALERTEVENTATTRIBUTESSTATUS_WARN  AlertEventAttributesStatus = "warn"
	ALERTEVENTATTRIBUTESSTATUS_ERROR AlertEventAttributesStatus = "error"
	ALERTEVENTATTRIBUTESSTATUS_OK    AlertEventAttributesStatus = "ok"
)

var allowedAlertEventAttributesStatusEnumValues = []AlertEventAttributesStatus{
	ALERTEVENTATTRIBUTESSTATUS_WARN,
	ALERTEVENTATTRIBUTESSTATUS_ERROR,
	ALERTEVENTATTRIBUTESSTATUS_OK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertEventAttributesStatus) GetAllowedValues() []AlertEventAttributesStatus {
	return allowedAlertEventAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertEventAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertEventAttributesStatus(value)
	return nil
}

// NewAlertEventAttributesStatusFromValue returns a pointer to a valid AlertEventAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertEventAttributesStatusFromValue(v string) (*AlertEventAttributesStatus, error) {
	ev := AlertEventAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertEventAttributesStatus: valid values are %v", v, allowedAlertEventAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertEventAttributesStatus) IsValid() bool {
	for _, existing := range allowedAlertEventAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertEventAttributesStatus value.
func (v AlertEventAttributesStatus) Ptr() *AlertEventAttributesStatus {
	return &v
}
