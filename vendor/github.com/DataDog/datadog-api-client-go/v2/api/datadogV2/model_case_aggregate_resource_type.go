// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseAggregateResourceType JSON:API resource type for case aggregation requests.
type CaseAggregateResourceType string

// List of CaseAggregateResourceType.
const (
	CASEAGGREGATERESOURCETYPE_AGGREGATE CaseAggregateResourceType = "aggregate"
)

var allowedCaseAggregateResourceTypeEnumValues = []CaseAggregateResourceType{
	CASEAGGREGATERESOURCETYPE_AGGREGATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseAggregateResourceType) GetAllowedValues() []CaseAggregateResourceType {
	return allowedCaseAggregateResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseAggregateResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseAggregateResourceType(value)
	return nil
}

// NewCaseAggregateResourceTypeFromValue returns a pointer to a valid CaseAggregateResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseAggregateResourceTypeFromValue(v string) (*CaseAggregateResourceType, error) {
	ev := CaseAggregateResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseAggregateResourceType: valid values are %v", v, allowedCaseAggregateResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseAggregateResourceType) IsValid() bool {
	for _, existing := range allowedCaseAggregateResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseAggregateResourceType value.
func (v CaseAggregateResourceType) Ptr() *CaseAggregateResourceType {
	return &v
}
