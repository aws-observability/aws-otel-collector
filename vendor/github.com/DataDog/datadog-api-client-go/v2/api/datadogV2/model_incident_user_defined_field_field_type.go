// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldFieldType The data type of the field. 1=dropdown, 2=multiselect, 3=textbox, 4=textarray, 5=metrictag, 6=autocomplete, 7=number, 8=datetime.
type IncidentUserDefinedFieldFieldType int32

// List of IncidentUserDefinedFieldFieldType.
const (
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_DROPDOWN     IncidentUserDefinedFieldFieldType = 1
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_MULTISELECT  IncidentUserDefinedFieldFieldType = 2
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_TEXTBOX      IncidentUserDefinedFieldFieldType = 3
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_TEXTARRAY    IncidentUserDefinedFieldFieldType = 4
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_METRICTAG    IncidentUserDefinedFieldFieldType = 5
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_AUTOCOMPLETE IncidentUserDefinedFieldFieldType = 6
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_NUMBER       IncidentUserDefinedFieldFieldType = 7
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_DATETIME     IncidentUserDefinedFieldFieldType = 8
)

var allowedIncidentUserDefinedFieldFieldTypeEnumValues = []IncidentUserDefinedFieldFieldType{
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_DROPDOWN,
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_MULTISELECT,
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_TEXTBOX,
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_TEXTARRAY,
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_METRICTAG,
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_AUTOCOMPLETE,
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_NUMBER,
	INCIDENTUSERDEFINEDFIELDFIELDTYPE_DATETIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentUserDefinedFieldFieldType) GetAllowedValues() []IncidentUserDefinedFieldFieldType {
	return allowedIncidentUserDefinedFieldFieldTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentUserDefinedFieldFieldType) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentUserDefinedFieldFieldType(value)
	return nil
}

// NewIncidentUserDefinedFieldFieldTypeFromValue returns a pointer to a valid IncidentUserDefinedFieldFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentUserDefinedFieldFieldTypeFromValue(v int32) (*IncidentUserDefinedFieldFieldType, error) {
	ev := IncidentUserDefinedFieldFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentUserDefinedFieldFieldType: valid values are %v", v, allowedIncidentUserDefinedFieldFieldTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentUserDefinedFieldFieldType) IsValid() bool {
	for _, existing := range allowedIncidentUserDefinedFieldFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentUserDefinedFieldFieldType value.
func (v IncidentUserDefinedFieldFieldType) Ptr() *IncidentUserDefinedFieldFieldType {
	return &v
}
