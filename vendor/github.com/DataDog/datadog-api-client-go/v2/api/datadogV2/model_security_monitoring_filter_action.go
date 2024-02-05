// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringFilterAction The type of filtering action.
type SecurityMonitoringFilterAction string

// List of SecurityMonitoringFilterAction.
const (
	SECURITYMONITORINGFILTERACTION_REQUIRE  SecurityMonitoringFilterAction = "require"
	SECURITYMONITORINGFILTERACTION_SUPPRESS SecurityMonitoringFilterAction = "suppress"
)

var allowedSecurityMonitoringFilterActionEnumValues = []SecurityMonitoringFilterAction{
	SECURITYMONITORINGFILTERACTION_REQUIRE,
	SECURITYMONITORINGFILTERACTION_SUPPRESS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringFilterAction) GetAllowedValues() []SecurityMonitoringFilterAction {
	return allowedSecurityMonitoringFilterActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringFilterAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringFilterAction(value)
	return nil
}

// NewSecurityMonitoringFilterActionFromValue returns a pointer to a valid SecurityMonitoringFilterAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringFilterActionFromValue(v string) (*SecurityMonitoringFilterAction, error) {
	ev := SecurityMonitoringFilterAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringFilterAction: valid values are %v", v, allowedSecurityMonitoringFilterActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringFilterAction) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringFilterActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringFilterAction value.
func (v SecurityMonitoringFilterAction) Ptr() *SecurityMonitoringFilterAction {
	return &v
}
