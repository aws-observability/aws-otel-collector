// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppTestEventTypeName Type of the event.
type CIAppTestEventTypeName string

// List of CIAppTestEventTypeName.
const (
	CIAPPTESTEVENTTYPENAME_CITEST CIAppTestEventTypeName = "citest"
)

var allowedCIAppTestEventTypeNameEnumValues = []CIAppTestEventTypeName{
	CIAPPTESTEVENTTYPENAME_CITEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppTestEventTypeName) GetAllowedValues() []CIAppTestEventTypeName {
	return allowedCIAppTestEventTypeNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppTestEventTypeName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppTestEventTypeName(value)
	return nil
}

// NewCIAppTestEventTypeNameFromValue returns a pointer to a valid CIAppTestEventTypeName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppTestEventTypeNameFromValue(v string) (*CIAppTestEventTypeName, error) {
	ev := CIAppTestEventTypeName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppTestEventTypeName: valid values are %v", v, allowedCIAppTestEventTypeNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppTestEventTypeName) IsValid() bool {
	for _, existing := range allowedCIAppTestEventTypeNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppTestEventTypeName value.
func (v CIAppTestEventTypeName) Ptr() *CIAppTestEventTypeName {
	return &v
}
