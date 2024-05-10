// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldType The incident user defined fields type.
type IncidentUserDefinedFieldType string

// List of IncidentUserDefinedFieldType.
const (
	INCIDENTUSERDEFINEDFIELDTYPE_USER_DEFINED_FIELD IncidentUserDefinedFieldType = "user_defined_field"
)

var allowedIncidentUserDefinedFieldTypeEnumValues = []IncidentUserDefinedFieldType{
	INCIDENTUSERDEFINEDFIELDTYPE_USER_DEFINED_FIELD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentUserDefinedFieldType) GetAllowedValues() []IncidentUserDefinedFieldType {
	return allowedIncidentUserDefinedFieldTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentUserDefinedFieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentUserDefinedFieldType(value)
	return nil
}

// NewIncidentUserDefinedFieldTypeFromValue returns a pointer to a valid IncidentUserDefinedFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentUserDefinedFieldTypeFromValue(v string) (*IncidentUserDefinedFieldType, error) {
	ev := IncidentUserDefinedFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentUserDefinedFieldType: valid values are %v", v, allowedIncidentUserDefinedFieldTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentUserDefinedFieldType) IsValid() bool {
	for _, existing := range allowedIncidentUserDefinedFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentUserDefinedFieldType value.
func (v IncidentUserDefinedFieldType) Ptr() *IncidentUserDefinedFieldType {
	return &v
}
