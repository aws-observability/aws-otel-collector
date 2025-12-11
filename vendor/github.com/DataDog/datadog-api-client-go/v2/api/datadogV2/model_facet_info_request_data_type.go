// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FacetInfoRequestDataType Users facet info request resource type.
type FacetInfoRequestDataType string

// List of FacetInfoRequestDataType.
const (
	FACETINFOREQUESTDATATYPE_USERS_FACET_INFO_REQUEST FacetInfoRequestDataType = "users_facet_info_request"
)

var allowedFacetInfoRequestDataTypeEnumValues = []FacetInfoRequestDataType{
	FACETINFOREQUESTDATATYPE_USERS_FACET_INFO_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FacetInfoRequestDataType) GetAllowedValues() []FacetInfoRequestDataType {
	return allowedFacetInfoRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FacetInfoRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FacetInfoRequestDataType(value)
	return nil
}

// NewFacetInfoRequestDataTypeFromValue returns a pointer to a valid FacetInfoRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFacetInfoRequestDataTypeFromValue(v string) (*FacetInfoRequestDataType, error) {
	ev := FacetInfoRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FacetInfoRequestDataType: valid values are %v", v, allowedFacetInfoRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FacetInfoRequestDataType) IsValid() bool {
	for _, existing := range allowedFacetInfoRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FacetInfoRequestDataType value.
func (v FacetInfoRequestDataType) Ptr() *FacetInfoRequestDataType {
	return &v
}
