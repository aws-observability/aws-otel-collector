// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagDescriptionSource Origin of the description. `human` indicates the description was written by a user, `ai_generated` was produced by AI, and `datadog` is a default supplied by Datadog.
type CostTagDescriptionSource string

// List of CostTagDescriptionSource.
const (
	COSTTAGDESCRIPTIONSOURCE_HUMAN        CostTagDescriptionSource = "human"
	COSTTAGDESCRIPTIONSOURCE_AI_GENERATED CostTagDescriptionSource = "ai_generated"
	COSTTAGDESCRIPTIONSOURCE_DATADOG      CostTagDescriptionSource = "datadog"
)

var allowedCostTagDescriptionSourceEnumValues = []CostTagDescriptionSource{
	COSTTAGDESCRIPTIONSOURCE_HUMAN,
	COSTTAGDESCRIPTIONSOURCE_AI_GENERATED,
	COSTTAGDESCRIPTIONSOURCE_DATADOG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostTagDescriptionSource) GetAllowedValues() []CostTagDescriptionSource {
	return allowedCostTagDescriptionSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostTagDescriptionSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostTagDescriptionSource(value)
	return nil
}

// NewCostTagDescriptionSourceFromValue returns a pointer to a valid CostTagDescriptionSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostTagDescriptionSourceFromValue(v string) (*CostTagDescriptionSource, error) {
	ev := CostTagDescriptionSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostTagDescriptionSource: valid values are %v", v, allowedCostTagDescriptionSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostTagDescriptionSource) IsValid() bool {
	for _, existing := range allowedCostTagDescriptionSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostTagDescriptionSource value.
func (v CostTagDescriptionSource) Ptr() *CostTagDescriptionSource {
	return &v
}
