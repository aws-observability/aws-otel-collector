// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionAggregateQueryJoinType Join type for aggregate query join conditions.
type MonitorFormulaAndFunctionAggregateQueryJoinType string

// List of MonitorFormulaAndFunctionAggregateQueryJoinType.
const (
	MONITORFORMULAANDFUNCTIONAGGREGATEQUERYJOINTYPE_INNER MonitorFormulaAndFunctionAggregateQueryJoinType = "inner"
	MONITORFORMULAANDFUNCTIONAGGREGATEQUERYJOINTYPE_LEFT  MonitorFormulaAndFunctionAggregateQueryJoinType = "left"
)

var allowedMonitorFormulaAndFunctionAggregateQueryJoinTypeEnumValues = []MonitorFormulaAndFunctionAggregateQueryJoinType{
	MONITORFORMULAANDFUNCTIONAGGREGATEQUERYJOINTYPE_INNER,
	MONITORFORMULAANDFUNCTIONAGGREGATEQUERYJOINTYPE_LEFT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionAggregateQueryJoinType) GetAllowedValues() []MonitorFormulaAndFunctionAggregateQueryJoinType {
	return allowedMonitorFormulaAndFunctionAggregateQueryJoinTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionAggregateQueryJoinType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionAggregateQueryJoinType(value)
	return nil
}

// NewMonitorFormulaAndFunctionAggregateQueryJoinTypeFromValue returns a pointer to a valid MonitorFormulaAndFunctionAggregateQueryJoinType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionAggregateQueryJoinTypeFromValue(v string) (*MonitorFormulaAndFunctionAggregateQueryJoinType, error) {
	ev := MonitorFormulaAndFunctionAggregateQueryJoinType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionAggregateQueryJoinType: valid values are %v", v, allowedMonitorFormulaAndFunctionAggregateQueryJoinTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionAggregateQueryJoinType) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionAggregateQueryJoinTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionAggregateQueryJoinType value.
func (v MonitorFormulaAndFunctionAggregateQueryJoinType) Ptr() *MonitorFormulaAndFunctionAggregateQueryJoinType {
	return &v
}
