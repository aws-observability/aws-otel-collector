// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostOrchestratorType Type of the Cloud Cost Management orchestrator resource.
type CostOrchestratorType string

// List of CostOrchestratorType.
const (
	COSTORCHESTRATORTYPE_COST_ORCHESTRATOR CostOrchestratorType = "cost_orchestrator"
)

var allowedCostOrchestratorTypeEnumValues = []CostOrchestratorType{
	COSTORCHESTRATORTYPE_COST_ORCHESTRATOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostOrchestratorType) GetAllowedValues() []CostOrchestratorType {
	return allowedCostOrchestratorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostOrchestratorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostOrchestratorType(value)
	return nil
}

// NewCostOrchestratorTypeFromValue returns a pointer to a valid CostOrchestratorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostOrchestratorTypeFromValue(v string) (*CostOrchestratorType, error) {
	ev := CostOrchestratorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostOrchestratorType: valid values are %v", v, allowedCostOrchestratorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostOrchestratorType) IsValid() bool {
	for _, existing := range allowedCostOrchestratorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostOrchestratorType value.
func (v CostOrchestratorType) Ptr() *CostOrchestratorType {
	return &v
}
