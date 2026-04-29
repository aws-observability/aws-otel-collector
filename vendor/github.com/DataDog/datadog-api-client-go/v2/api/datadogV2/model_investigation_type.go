// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InvestigationType The resource type for investigations.
type InvestigationType string

// List of InvestigationType.
const (
	INVESTIGATIONTYPE_INVESTIGATION InvestigationType = "investigation"
)

var allowedInvestigationTypeEnumValues = []InvestigationType{
	INVESTIGATIONTYPE_INVESTIGATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *InvestigationType) GetAllowedValues() []InvestigationType {
	return allowedInvestigationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InvestigationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InvestigationType(value)
	return nil
}

// NewInvestigationTypeFromValue returns a pointer to a valid InvestigationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInvestigationTypeFromValue(v string) (*InvestigationType, error) {
	ev := InvestigationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InvestigationType: valid values are %v", v, allowedInvestigationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InvestigationType) IsValid() bool {
	for _, existing := range allowedInvestigationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InvestigationType value.
func (v InvestigationType) Ptr() *InvestigationType {
	return &v
}
