// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnnotationType Annotation resource type.
type AnnotationType string

// List of AnnotationType.
const (
	ANNOTATIONTYPE_ANNOTATION AnnotationType = "annotation"
)

var allowedAnnotationTypeEnumValues = []AnnotationType{
	ANNOTATIONTYPE_ANNOTATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnnotationType) GetAllowedValues() []AnnotationType {
	return allowedAnnotationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnnotationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnnotationType(value)
	return nil
}

// NewAnnotationTypeFromValue returns a pointer to a valid AnnotationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnnotationTypeFromValue(v string) (*AnnotationType, error) {
	ev := AnnotationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnnotationType: valid values are %v", v, allowedAnnotationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnnotationType) IsValid() bool {
	for _, existing := range allowedAnnotationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnnotationType value.
func (v AnnotationType) Ptr() *AnnotationType {
	return &v
}
