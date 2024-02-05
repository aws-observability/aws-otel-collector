// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamWidgetDefinitionType Type of the list stream widget.
type ListStreamWidgetDefinitionType string

// List of ListStreamWidgetDefinitionType.
const (
	LISTSTREAMWIDGETDEFINITIONTYPE_LIST_STREAM ListStreamWidgetDefinitionType = "list_stream"
)

var allowedListStreamWidgetDefinitionTypeEnumValues = []ListStreamWidgetDefinitionType{
	LISTSTREAMWIDGETDEFINITIONTYPE_LIST_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListStreamWidgetDefinitionType) GetAllowedValues() []ListStreamWidgetDefinitionType {
	return allowedListStreamWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListStreamWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamWidgetDefinitionType(value)
	return nil
}

// NewListStreamWidgetDefinitionTypeFromValue returns a pointer to a valid ListStreamWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListStreamWidgetDefinitionTypeFromValue(v string) (*ListStreamWidgetDefinitionType, error) {
	ev := ListStreamWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListStreamWidgetDefinitionType: valid values are %v", v, allowedListStreamWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListStreamWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedListStreamWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamWidgetDefinitionType value.
func (v ListStreamWidgetDefinitionType) Ptr() *ListStreamWidgetDefinitionType {
	return &v
}
