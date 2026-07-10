// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmFacetInfoType The JSON:API type for facet info resources. The value should always be `facet_info`.
type CsmFacetInfoType string

// List of CsmFacetInfoType.
const (
	CSMFACETINFOTYPE_FACET_INFO CsmFacetInfoType = "facet_info"
)

var allowedCsmFacetInfoTypeEnumValues = []CsmFacetInfoType{
	CSMFACETINFOTYPE_FACET_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CsmFacetInfoType) GetAllowedValues() []CsmFacetInfoType {
	return allowedCsmFacetInfoTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CsmFacetInfoType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CsmFacetInfoType(value)
	return nil
}

// NewCsmFacetInfoTypeFromValue returns a pointer to a valid CsmFacetInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCsmFacetInfoTypeFromValue(v string) (*CsmFacetInfoType, error) {
	ev := CsmFacetInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CsmFacetInfoType: valid values are %v", v, allowedCsmFacetInfoTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CsmFacetInfoType) IsValid() bool {
	for _, existing := range allowedCsmFacetInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CsmFacetInfoType value.
func (v CsmFacetInfoType) Ptr() *CsmFacetInfoType {
	return &v
}
