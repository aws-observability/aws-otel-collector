// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InterfaceAttributesStatus The interface status
type InterfaceAttributesStatus string

// List of InterfaceAttributesStatus.
const (
	INTERFACEATTRIBUTESSTATUS_UP      InterfaceAttributesStatus = "up"
	INTERFACEATTRIBUTESSTATUS_DOWN    InterfaceAttributesStatus = "down"
	INTERFACEATTRIBUTESSTATUS_WARNING InterfaceAttributesStatus = "warning"
	INTERFACEATTRIBUTESSTATUS_OFF     InterfaceAttributesStatus = "off"
)

var allowedInterfaceAttributesStatusEnumValues = []InterfaceAttributesStatus{
	INTERFACEATTRIBUTESSTATUS_UP,
	INTERFACEATTRIBUTESSTATUS_DOWN,
	INTERFACEATTRIBUTESSTATUS_WARNING,
	INTERFACEATTRIBUTESSTATUS_OFF,
}

// GetAllowedValues reeturns the list of possible values.
func (v *InterfaceAttributesStatus) GetAllowedValues() []InterfaceAttributesStatus {
	return allowedInterfaceAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InterfaceAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InterfaceAttributesStatus(value)
	return nil
}

// NewInterfaceAttributesStatusFromValue returns a pointer to a valid InterfaceAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInterfaceAttributesStatusFromValue(v string) (*InterfaceAttributesStatus, error) {
	ev := InterfaceAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InterfaceAttributesStatus: valid values are %v", v, allowedInterfaceAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InterfaceAttributesStatus) IsValid() bool {
	for _, existing := range allowedInterfaceAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InterfaceAttributesStatus value.
func (v InterfaceAttributesStatus) Ptr() *InterfaceAttributesStatus {
	return &v
}
