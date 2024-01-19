// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsPatchTestOperationName The operation to perform
type SyntheticsPatchTestOperationName string

// List of SyntheticsPatchTestOperationName.
const (
	SYNTHETICSPATCHTESTOPERATIONNAME_ADD     SyntheticsPatchTestOperationName = "add"
	SYNTHETICSPATCHTESTOPERATIONNAME_REMOVE  SyntheticsPatchTestOperationName = "remove"
	SYNTHETICSPATCHTESTOPERATIONNAME_REPLACE SyntheticsPatchTestOperationName = "replace"
	SYNTHETICSPATCHTESTOPERATIONNAME_MOVE    SyntheticsPatchTestOperationName = "move"
	SYNTHETICSPATCHTESTOPERATIONNAME_COPY    SyntheticsPatchTestOperationName = "copy"
	SYNTHETICSPATCHTESTOPERATIONNAME_TEST    SyntheticsPatchTestOperationName = "test"
)

var allowedSyntheticsPatchTestOperationNameEnumValues = []SyntheticsPatchTestOperationName{
	SYNTHETICSPATCHTESTOPERATIONNAME_ADD,
	SYNTHETICSPATCHTESTOPERATIONNAME_REMOVE,
	SYNTHETICSPATCHTESTOPERATIONNAME_REPLACE,
	SYNTHETICSPATCHTESTOPERATIONNAME_MOVE,
	SYNTHETICSPATCHTESTOPERATIONNAME_COPY,
	SYNTHETICSPATCHTESTOPERATIONNAME_TEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsPatchTestOperationName) GetAllowedValues() []SyntheticsPatchTestOperationName {
	return allowedSyntheticsPatchTestOperationNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsPatchTestOperationName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsPatchTestOperationName(value)
	return nil
}

// NewSyntheticsPatchTestOperationNameFromValue returns a pointer to a valid SyntheticsPatchTestOperationName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsPatchTestOperationNameFromValue(v string) (*SyntheticsPatchTestOperationName, error) {
	ev := SyntheticsPatchTestOperationName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsPatchTestOperationName: valid values are %v", v, allowedSyntheticsPatchTestOperationNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsPatchTestOperationName) IsValid() bool {
	for _, existing := range allowedSyntheticsPatchTestOperationNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsPatchTestOperationName value.
func (v SyntheticsPatchTestOperationName) Ptr() *SyntheticsPatchTestOperationName {
	return &v
}
