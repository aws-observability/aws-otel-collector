// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AllocationType The type of targeting rule (called allocation in the API model).
type AllocationType string

// List of AllocationType.
const (
	ALLOCATIONTYPE_FEATURE_GATE AllocationType = "FEATURE_GATE"
	ALLOCATIONTYPE_CANARY       AllocationType = "CANARY"
)

var allowedAllocationTypeEnumValues = []AllocationType{
	ALLOCATIONTYPE_FEATURE_GATE,
	ALLOCATIONTYPE_CANARY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AllocationType) GetAllowedValues() []AllocationType {
	return allowedAllocationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AllocationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AllocationType(value)
	return nil
}

// NewAllocationTypeFromValue returns a pointer to a valid AllocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAllocationTypeFromValue(v string) (*AllocationType, error) {
	ev := AllocationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AllocationType: valid values are %v", v, allowedAllocationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AllocationType) IsValid() bool {
	for _, existing := range allowedAllocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AllocationType value.
func (v AllocationType) Ptr() *AllocationType {
	return &v
}
