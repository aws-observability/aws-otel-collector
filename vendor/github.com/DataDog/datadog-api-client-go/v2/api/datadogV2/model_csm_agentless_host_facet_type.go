// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmAgentlessHostFacetType The JSON:API type for agentless host facet resources. The value should always be `agentless_host_facet`.
type CsmAgentlessHostFacetType string

// List of CsmAgentlessHostFacetType.
const (
	CSMAGENTLESSHOSTFACETTYPE_AGENTLESS_HOST_FACET CsmAgentlessHostFacetType = "agentless_host_facet"
)

var allowedCsmAgentlessHostFacetTypeEnumValues = []CsmAgentlessHostFacetType{
	CSMAGENTLESSHOSTFACETTYPE_AGENTLESS_HOST_FACET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CsmAgentlessHostFacetType) GetAllowedValues() []CsmAgentlessHostFacetType {
	return allowedCsmAgentlessHostFacetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CsmAgentlessHostFacetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CsmAgentlessHostFacetType(value)
	return nil
}

// NewCsmAgentlessHostFacetTypeFromValue returns a pointer to a valid CsmAgentlessHostFacetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCsmAgentlessHostFacetTypeFromValue(v string) (*CsmAgentlessHostFacetType, error) {
	ev := CsmAgentlessHostFacetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CsmAgentlessHostFacetType: valid values are %v", v, allowedCsmAgentlessHostFacetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CsmAgentlessHostFacetType) IsValid() bool {
	for _, existing := range allowedCsmAgentlessHostFacetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CsmAgentlessHostFacetType value.
func (v CsmAgentlessHostFacetType) Ptr() *CsmAgentlessHostFacetType {
	return &v
}
