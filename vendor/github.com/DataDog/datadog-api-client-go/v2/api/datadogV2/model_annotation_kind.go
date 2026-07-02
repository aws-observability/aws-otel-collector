// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnnotationKind Kind of annotation. `pointInTime` annotations mark a single moment in time,
// while `timeRegion` annotations span a window of time and require an `end_time`.
type AnnotationKind string

// List of AnnotationKind.
const (
	ANNOTATIONKIND_POINT_IN_TIME AnnotationKind = "pointInTime"
	ANNOTATIONKIND_TIME_REGION   AnnotationKind = "timeRegion"
)

var allowedAnnotationKindEnumValues = []AnnotationKind{
	ANNOTATIONKIND_POINT_IN_TIME,
	ANNOTATIONKIND_TIME_REGION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnnotationKind) GetAllowedValues() []AnnotationKind {
	return allowedAnnotationKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnnotationKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnnotationKind(value)
	return nil
}

// NewAnnotationKindFromValue returns a pointer to a valid AnnotationKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnnotationKindFromValue(v string) (*AnnotationKind, error) {
	ev := AnnotationKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnnotationKind: valid values are %v", v, allowedAnnotationKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnnotationKind) IsValid() bool {
	for _, existing := range allowedAnnotationKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnnotationKind value.
func (v AnnotationKind) Ptr() *AnnotationKind {
	return &v
}
