// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightResourceType JSON:API resource type for a governance insight.
type GovernanceInsightResourceType string

// List of GovernanceInsightResourceType.
const (
	GOVERNANCEINSIGHTRESOURCETYPE_INSIGHT GovernanceInsightResourceType = "insight"
)

var allowedGovernanceInsightResourceTypeEnumValues = []GovernanceInsightResourceType{
	GOVERNANCEINSIGHTRESOURCETYPE_INSIGHT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GovernanceInsightResourceType) GetAllowedValues() []GovernanceInsightResourceType {
	return allowedGovernanceInsightResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GovernanceInsightResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GovernanceInsightResourceType(value)
	return nil
}

// NewGovernanceInsightResourceTypeFromValue returns a pointer to a valid GovernanceInsightResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGovernanceInsightResourceTypeFromValue(v string) (*GovernanceInsightResourceType, error) {
	ev := GovernanceInsightResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GovernanceInsightResourceType: valid values are %v", v, allowedGovernanceInsightResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GovernanceInsightResourceType) IsValid() bool {
	for _, existing := range allowedGovernanceInsightResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GovernanceInsightResourceType value.
func (v GovernanceInsightResourceType) Ptr() *GovernanceInsightResourceType {
	return &v
}
