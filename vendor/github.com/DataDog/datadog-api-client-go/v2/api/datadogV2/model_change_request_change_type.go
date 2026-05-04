// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestChangeType The type of the change request.
type ChangeRequestChangeType string

// List of ChangeRequestChangeType.
const (
	CHANGEREQUESTCHANGETYPE_NORMAL    ChangeRequestChangeType = "NORMAL"
	CHANGEREQUESTCHANGETYPE_STANDARD  ChangeRequestChangeType = "STANDARD"
	CHANGEREQUESTCHANGETYPE_EMERGENCY ChangeRequestChangeType = "EMERGENCY"
)

var allowedChangeRequestChangeTypeEnumValues = []ChangeRequestChangeType{
	CHANGEREQUESTCHANGETYPE_NORMAL,
	CHANGEREQUESTCHANGETYPE_STANDARD,
	CHANGEREQUESTCHANGETYPE_EMERGENCY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeRequestChangeType) GetAllowedValues() []ChangeRequestChangeType {
	return allowedChangeRequestChangeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeRequestChangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeRequestChangeType(value)
	return nil
}

// NewChangeRequestChangeTypeFromValue returns a pointer to a valid ChangeRequestChangeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeRequestChangeTypeFromValue(v string) (*ChangeRequestChangeType, error) {
	ev := ChangeRequestChangeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeRequestChangeType: valid values are %v", v, allowedChangeRequestChangeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeRequestChangeType) IsValid() bool {
	for _, existing := range allowedChangeRequestChangeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeRequestChangeType value.
func (v ChangeRequestChangeType) Ptr() *ChangeRequestChangeType {
	return &v
}
