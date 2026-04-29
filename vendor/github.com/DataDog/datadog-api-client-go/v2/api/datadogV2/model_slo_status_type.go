// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloStatusType The type of the SLO status resource.
type SloStatusType string

// List of SloStatusType.
const (
	SLOSTATUSTYPE_SLO_STATUS SloStatusType = "slo_status"
)

var allowedSloStatusTypeEnumValues = []SloStatusType{
	SLOSTATUSTYPE_SLO_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SloStatusType) GetAllowedValues() []SloStatusType {
	return allowedSloStatusTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SloStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SloStatusType(value)
	return nil
}

// NewSloStatusTypeFromValue returns a pointer to a valid SloStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSloStatusTypeFromValue(v string) (*SloStatusType, error) {
	ev := SloStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SloStatusType: valid values are %v", v, allowedSloStatusTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SloStatusType) IsValid() bool {
	for _, existing := range allowedSloStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SloStatusType value.
func (v SloStatusType) Ptr() *SloStatusType {
	return &v
}
