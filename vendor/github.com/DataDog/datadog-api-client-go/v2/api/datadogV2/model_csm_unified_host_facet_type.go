// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmUnifiedHostFacetType The JSON:API type for unified host facet resources. The value should always be `unified_host_facet`.
type CsmUnifiedHostFacetType string

// List of CsmUnifiedHostFacetType.
const (
	CSMUNIFIEDHOSTFACETTYPE_UNIFIED_HOST_FACET CsmUnifiedHostFacetType = "unified_host_facet"
)

var allowedCsmUnifiedHostFacetTypeEnumValues = []CsmUnifiedHostFacetType{
	CSMUNIFIEDHOSTFACETTYPE_UNIFIED_HOST_FACET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CsmUnifiedHostFacetType) GetAllowedValues() []CsmUnifiedHostFacetType {
	return allowedCsmUnifiedHostFacetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CsmUnifiedHostFacetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CsmUnifiedHostFacetType(value)
	return nil
}

// NewCsmUnifiedHostFacetTypeFromValue returns a pointer to a valid CsmUnifiedHostFacetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCsmUnifiedHostFacetTypeFromValue(v string) (*CsmUnifiedHostFacetType, error) {
	ev := CsmUnifiedHostFacetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CsmUnifiedHostFacetType: valid values are %v", v, allowedCsmUnifiedHostFacetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CsmUnifiedHostFacetType) IsValid() bool {
	for _, existing := range allowedCsmUnifiedHostFacetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CsmUnifiedHostFacetType value.
func (v CsmUnifiedHostFacetType) Ptr() *CsmUnifiedHostFacetType {
	return &v
}
