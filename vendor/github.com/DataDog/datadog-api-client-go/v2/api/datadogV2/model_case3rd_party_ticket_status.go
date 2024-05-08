// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Case3rdPartyTicketStatus Case status
type Case3rdPartyTicketStatus string

// List of Case3rdPartyTicketStatus.
const (
	CASE3RDPARTYTICKETSTATUS_IN_PROGRESS Case3rdPartyTicketStatus = "IN_PROGRESS"
	CASE3RDPARTYTICKETSTATUS_COMPLETED   Case3rdPartyTicketStatus = "COMPLETED"
	CASE3RDPARTYTICKETSTATUS_FAILED      Case3rdPartyTicketStatus = "FAILED"
)

var allowedCase3rdPartyTicketStatusEnumValues = []Case3rdPartyTicketStatus{
	CASE3RDPARTYTICKETSTATUS_IN_PROGRESS,
	CASE3RDPARTYTICKETSTATUS_COMPLETED,
	CASE3RDPARTYTICKETSTATUS_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Case3rdPartyTicketStatus) GetAllowedValues() []Case3rdPartyTicketStatus {
	return allowedCase3rdPartyTicketStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Case3rdPartyTicketStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Case3rdPartyTicketStatus(value)
	return nil
}

// NewCase3rdPartyTicketStatusFromValue returns a pointer to a valid Case3rdPartyTicketStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCase3rdPartyTicketStatusFromValue(v string) (*Case3rdPartyTicketStatus, error) {
	ev := Case3rdPartyTicketStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Case3rdPartyTicketStatus: valid values are %v", v, allowedCase3rdPartyTicketStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Case3rdPartyTicketStatus) IsValid() bool {
	for _, existing := range allowedCase3rdPartyTicketStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Case3rdPartyTicketStatus value.
func (v Case3rdPartyTicketStatus) Ptr() *Case3rdPartyTicketStatus {
	return &v
}
