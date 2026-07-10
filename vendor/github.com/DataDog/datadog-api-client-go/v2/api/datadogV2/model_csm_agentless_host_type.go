// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmAgentlessHostType The JSON:API type for agentless host resources. The value should always be `agentless_host`.
type CsmAgentlessHostType string

// List of CsmAgentlessHostType.
const (
	CSMAGENTLESSHOSTTYPE_AGENTLESS_HOST CsmAgentlessHostType = "agentless_host"
)

var allowedCsmAgentlessHostTypeEnumValues = []CsmAgentlessHostType{
	CSMAGENTLESSHOSTTYPE_AGENTLESS_HOST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CsmAgentlessHostType) GetAllowedValues() []CsmAgentlessHostType {
	return allowedCsmAgentlessHostTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CsmAgentlessHostType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CsmAgentlessHostType(value)
	return nil
}

// NewCsmAgentlessHostTypeFromValue returns a pointer to a valid CsmAgentlessHostType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCsmAgentlessHostTypeFromValue(v string) (*CsmAgentlessHostType, error) {
	ev := CsmAgentlessHostType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CsmAgentlessHostType: valid values are %v", v, allowedCsmAgentlessHostTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CsmAgentlessHostType) IsValid() bool {
	for _, existing := range allowedCsmAgentlessHostTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CsmAgentlessHostType value.
func (v CsmAgentlessHostType) Ptr() *CsmAgentlessHostType {
	return &v
}
