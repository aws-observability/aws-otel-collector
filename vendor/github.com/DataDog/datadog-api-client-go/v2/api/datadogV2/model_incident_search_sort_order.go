// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSearchSortOrder The ways searched incidents can be sorted.
type IncidentSearchSortOrder string

// List of IncidentSearchSortOrder.
const (
	INCIDENTSEARCHSORTORDER_CREATED_ASCENDING  IncidentSearchSortOrder = "created"
	INCIDENTSEARCHSORTORDER_CREATED_DESCENDING IncidentSearchSortOrder = "-created"
)

var allowedIncidentSearchSortOrderEnumValues = []IncidentSearchSortOrder{
	INCIDENTSEARCHSORTORDER_CREATED_ASCENDING,
	INCIDENTSEARCHSORTORDER_CREATED_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentSearchSortOrder) GetAllowedValues() []IncidentSearchSortOrder {
	return allowedIncidentSearchSortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentSearchSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentSearchSortOrder(value)
	return nil
}

// NewIncidentSearchSortOrderFromValue returns a pointer to a valid IncidentSearchSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentSearchSortOrderFromValue(v string) (*IncidentSearchSortOrder, error) {
	ev := IncidentSearchSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentSearchSortOrder: valid values are %v", v, allowedIncidentSearchSortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentSearchSortOrder) IsValid() bool {
	for _, existing := range allowedIncidentSearchSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentSearchSortOrder value.
func (v IncidentSearchSortOrder) Ptr() *IncidentSearchSortOrder {
	return &v
}
