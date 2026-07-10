// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PageAnnotationsType Page annotations resource type.
type PageAnnotationsType string

// List of PageAnnotationsType.
const (
	PAGEANNOTATIONSTYPE_PAGE_ANNOTATIONS PageAnnotationsType = "page_annotations"
)

var allowedPageAnnotationsTypeEnumValues = []PageAnnotationsType{
	PAGEANNOTATIONSTYPE_PAGE_ANNOTATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PageAnnotationsType) GetAllowedValues() []PageAnnotationsType {
	return allowedPageAnnotationsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PageAnnotationsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PageAnnotationsType(value)
	return nil
}

// NewPageAnnotationsTypeFromValue returns a pointer to a valid PageAnnotationsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPageAnnotationsTypeFromValue(v string) (*PageAnnotationsType, error) {
	ev := PageAnnotationsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PageAnnotationsType: valid values are %v", v, allowedPageAnnotationsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PageAnnotationsType) IsValid() bool {
	for _, existing := range allowedPageAnnotationsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PageAnnotationsType value.
func (v PageAnnotationsType) Ptr() *PageAnnotationsType {
	return &v
}
