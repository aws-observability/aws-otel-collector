// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertEventCustomAttributesStatus The status of the alert.
type AlertEventCustomAttributesStatus string

// List of AlertEventCustomAttributesStatus.
const (
	ALERTEVENTCUSTOMATTRIBUTESSTATUS_WARN  AlertEventCustomAttributesStatus = "warn"
	ALERTEVENTCUSTOMATTRIBUTESSTATUS_ERROR AlertEventCustomAttributesStatus = "error"
	ALERTEVENTCUSTOMATTRIBUTESSTATUS_OK    AlertEventCustomAttributesStatus = "ok"
)

var allowedAlertEventCustomAttributesStatusEnumValues = []AlertEventCustomAttributesStatus{
	ALERTEVENTCUSTOMATTRIBUTESSTATUS_WARN,
	ALERTEVENTCUSTOMATTRIBUTESSTATUS_ERROR,
	ALERTEVENTCUSTOMATTRIBUTESSTATUS_OK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertEventCustomAttributesStatus) GetAllowedValues() []AlertEventCustomAttributesStatus {
	return allowedAlertEventCustomAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertEventCustomAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertEventCustomAttributesStatus(value)
	return nil
}

// NewAlertEventCustomAttributesStatusFromValue returns a pointer to a valid AlertEventCustomAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertEventCustomAttributesStatusFromValue(v string) (*AlertEventCustomAttributesStatus, error) {
	ev := AlertEventCustomAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertEventCustomAttributesStatus: valid values are %v", v, allowedAlertEventCustomAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertEventCustomAttributesStatus) IsValid() bool {
	for _, existing := range allowedAlertEventCustomAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertEventCustomAttributesStatus value.
func (v AlertEventCustomAttributesStatus) Ptr() *AlertEventCustomAttributesStatus {
	return &v
}
