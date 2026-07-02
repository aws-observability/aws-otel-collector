// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisEditType The type of code edit to apply when fixing a violation.
type AnalysisEditType string

// List of AnalysisEditType.
const (
	ANALYSISEDITTYPE_ADD    AnalysisEditType = "ADD"
	ANALYSISEDITTYPE_UPDATE AnalysisEditType = "UPDATE"
	ANALYSISEDITTYPE_REMOVE AnalysisEditType = "REMOVE"
)

var allowedAnalysisEditTypeEnumValues = []AnalysisEditType{
	ANALYSISEDITTYPE_ADD,
	ANALYSISEDITTYPE_UPDATE,
	ANALYSISEDITTYPE_REMOVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnalysisEditType) GetAllowedValues() []AnalysisEditType {
	return allowedAnalysisEditTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnalysisEditType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnalysisEditType(value)
	return nil
}

// NewAnalysisEditTypeFromValue returns a pointer to a valid AnalysisEditType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnalysisEditTypeFromValue(v string) (*AnalysisEditType, error) {
	ev := AnalysisEditType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnalysisEditType: valid values are %v", v, allowedAnalysisEditTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnalysisEditType) IsValid() bool {
	for _, existing := range allowedAnalysisEditTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnalysisEditType value.
func (v AnalysisEditType) Ptr() *AnalysisEditType {
	return &v
}
