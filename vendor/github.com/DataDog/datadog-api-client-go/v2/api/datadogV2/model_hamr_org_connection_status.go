// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HamrOrgConnectionStatus Status of the HAMR connection:
// - 0: UNSPECIFIED - Connection status not specified
// - 1: ONBOARDING - Initial setup of HAMR connection
// - 2: PASSIVE - Secondary organization in passive standby mode
// - 3: FAILOVER - Liminal status between PASSIVE and ACTIVE
// - 4: ACTIVE - Organization is an active failover
// - 5: RECOVERY - Recovery operation in progress
type HamrOrgConnectionStatus int32

// List of HamrOrgConnectionStatus.
const (
	HAMRORGCONNECTIONSTATUS_UNSPECIFIED HamrOrgConnectionStatus = 0
	HAMRORGCONNECTIONSTATUS_ONBOARDING  HamrOrgConnectionStatus = 1
	HAMRORGCONNECTIONSTATUS_PASSIVE     HamrOrgConnectionStatus = 2
	HAMRORGCONNECTIONSTATUS_FAILOVER    HamrOrgConnectionStatus = 3
	HAMRORGCONNECTIONSTATUS_ACTIVE      HamrOrgConnectionStatus = 4
	HAMRORGCONNECTIONSTATUS_RECOVERY    HamrOrgConnectionStatus = 5
)

var allowedHamrOrgConnectionStatusEnumValues = []HamrOrgConnectionStatus{
	HAMRORGCONNECTIONSTATUS_UNSPECIFIED,
	HAMRORGCONNECTIONSTATUS_ONBOARDING,
	HAMRORGCONNECTIONSTATUS_PASSIVE,
	HAMRORGCONNECTIONSTATUS_FAILOVER,
	HAMRORGCONNECTIONSTATUS_ACTIVE,
	HAMRORGCONNECTIONSTATUS_RECOVERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HamrOrgConnectionStatus) GetAllowedValues() []HamrOrgConnectionStatus {
	return allowedHamrOrgConnectionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HamrOrgConnectionStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HamrOrgConnectionStatus(value)
	return nil
}

// NewHamrOrgConnectionStatusFromValue returns a pointer to a valid HamrOrgConnectionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHamrOrgConnectionStatusFromValue(v int32) (*HamrOrgConnectionStatus, error) {
	ev := HamrOrgConnectionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HamrOrgConnectionStatus: valid values are %v", v, allowedHamrOrgConnectionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HamrOrgConnectionStatus) IsValid() bool {
	for _, existing := range allowedHamrOrgConnectionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HamrOrgConnectionStatus value.
func (v HamrOrgConnectionStatus) Ptr() *HamrOrgConnectionStatus {
	return &v
}
