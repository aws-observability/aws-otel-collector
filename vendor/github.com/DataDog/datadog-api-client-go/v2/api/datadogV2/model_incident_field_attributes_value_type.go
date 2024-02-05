// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentFieldAttributesValueType Type of the multiple value field definitions.
type IncidentFieldAttributesValueType string

// List of IncidentFieldAttributesValueType.
const (
	INCIDENTFIELDATTRIBUTESVALUETYPE_MULTISELECT  IncidentFieldAttributesValueType = "multiselect"
	INCIDENTFIELDATTRIBUTESVALUETYPE_TEXTARRAY    IncidentFieldAttributesValueType = "textarray"
	INCIDENTFIELDATTRIBUTESVALUETYPE_METRICTAG    IncidentFieldAttributesValueType = "metrictag"
	INCIDENTFIELDATTRIBUTESVALUETYPE_AUTOCOMPLETE IncidentFieldAttributesValueType = "autocomplete"
)

var allowedIncidentFieldAttributesValueTypeEnumValues = []IncidentFieldAttributesValueType{
	INCIDENTFIELDATTRIBUTESVALUETYPE_MULTISELECT,
	INCIDENTFIELDATTRIBUTESVALUETYPE_TEXTARRAY,
	INCIDENTFIELDATTRIBUTESVALUETYPE_METRICTAG,
	INCIDENTFIELDATTRIBUTESVALUETYPE_AUTOCOMPLETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentFieldAttributesValueType) GetAllowedValues() []IncidentFieldAttributesValueType {
	return allowedIncidentFieldAttributesValueTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentFieldAttributesValueType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentFieldAttributesValueType(value)
	return nil
}

// NewIncidentFieldAttributesValueTypeFromValue returns a pointer to a valid IncidentFieldAttributesValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentFieldAttributesValueTypeFromValue(v string) (*IncidentFieldAttributesValueType, error) {
	ev := IncidentFieldAttributesValueType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentFieldAttributesValueType: valid values are %v", v, allowedIncidentFieldAttributesValueTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentFieldAttributesValueType) IsValid() bool {
	for _, existing := range allowedIncidentFieldAttributesValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentFieldAttributesValueType value.
func (v IncidentFieldAttributesValueType) Ptr() *IncidentFieldAttributesValueType {
	return &v
}
