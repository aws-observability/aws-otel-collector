// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NoteWidgetDefinitionType Type of the note widget.
type NoteWidgetDefinitionType string

// List of NoteWidgetDefinitionType.
const (
	NOTEWIDGETDEFINITIONTYPE_NOTE NoteWidgetDefinitionType = "note"
)

var allowedNoteWidgetDefinitionTypeEnumValues = []NoteWidgetDefinitionType{
	NOTEWIDGETDEFINITIONTYPE_NOTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NoteWidgetDefinitionType) GetAllowedValues() []NoteWidgetDefinitionType {
	return allowedNoteWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NoteWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NoteWidgetDefinitionType(value)
	return nil
}

// NewNoteWidgetDefinitionTypeFromValue returns a pointer to a valid NoteWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNoteWidgetDefinitionTypeFromValue(v string) (*NoteWidgetDefinitionType, error) {
	ev := NoteWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NoteWidgetDefinitionType: valid values are %v", v, allowedNoteWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NoteWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedNoteWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NoteWidgetDefinitionType value.
func (v NoteWidgetDefinitionType) Ptr() *NoteWidgetDefinitionType {
	return &v
}
