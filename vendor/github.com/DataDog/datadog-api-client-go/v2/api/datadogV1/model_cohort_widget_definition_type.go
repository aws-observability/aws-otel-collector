// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CohortWidgetDefinitionType Type of the Cohort widget.
type CohortWidgetDefinitionType string

// List of CohortWidgetDefinitionType.
const (
	COHORTWIDGETDEFINITIONTYPE_COHORT CohortWidgetDefinitionType = "cohort"
)

var allowedCohortWidgetDefinitionTypeEnumValues = []CohortWidgetDefinitionType{
	COHORTWIDGETDEFINITIONTYPE_COHORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CohortWidgetDefinitionType) GetAllowedValues() []CohortWidgetDefinitionType {
	return allowedCohortWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CohortWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CohortWidgetDefinitionType(value)
	return nil
}

// NewCohortWidgetDefinitionTypeFromValue returns a pointer to a valid CohortWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCohortWidgetDefinitionTypeFromValue(v string) (*CohortWidgetDefinitionType, error) {
	ev := CohortWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CohortWidgetDefinitionType: valid values are %v", v, allowedCohortWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CohortWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedCohortWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CohortWidgetDefinitionType value.
func (v CohortWidgetDefinitionType) Ptr() *CohortWidgetDefinitionType {
	return &v
}
