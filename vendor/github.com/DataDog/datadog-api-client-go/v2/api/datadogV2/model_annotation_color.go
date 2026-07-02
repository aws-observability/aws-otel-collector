// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnnotationColor Color used to render the annotation in the UI.
type AnnotationColor string

// List of AnnotationColor.
const (
	ANNOTATIONCOLOR_GRAY   AnnotationColor = "gray"
	ANNOTATIONCOLOR_BLUE   AnnotationColor = "blue"
	ANNOTATIONCOLOR_PURPLE AnnotationColor = "purple"
	ANNOTATIONCOLOR_GREEN  AnnotationColor = "green"
	ANNOTATIONCOLOR_YELLOW AnnotationColor = "yellow"
	ANNOTATIONCOLOR_RED    AnnotationColor = "red"
)

var allowedAnnotationColorEnumValues = []AnnotationColor{
	ANNOTATIONCOLOR_GRAY,
	ANNOTATIONCOLOR_BLUE,
	ANNOTATIONCOLOR_PURPLE,
	ANNOTATIONCOLOR_GREEN,
	ANNOTATIONCOLOR_YELLOW,
	ANNOTATIONCOLOR_RED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnnotationColor) GetAllowedValues() []AnnotationColor {
	return allowedAnnotationColorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnnotationColor) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnnotationColor(value)
	return nil
}

// NewAnnotationColorFromValue returns a pointer to a valid AnnotationColor
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnnotationColorFromValue(v string) (*AnnotationColor, error) {
	ev := AnnotationColor(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnnotationColor: valid values are %v", v, allowedAnnotationColorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnnotationColor) IsValid() bool {
	for _, existing := range allowedAnnotationColorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnnotationColor value.
func (v AnnotationColor) Ptr() *AnnotationColor {
	return &v
}
