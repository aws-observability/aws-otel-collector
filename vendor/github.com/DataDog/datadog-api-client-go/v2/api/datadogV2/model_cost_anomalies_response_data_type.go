// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostAnomaliesResponseDataType Type of the cost anomalies collection resource. Must be `anomalies`.
type CostAnomaliesResponseDataType string

// List of CostAnomaliesResponseDataType.
const (
	COSTANOMALIESRESPONSEDATATYPE_ANOMALIES CostAnomaliesResponseDataType = "anomalies"
)

var allowedCostAnomaliesResponseDataTypeEnumValues = []CostAnomaliesResponseDataType{
	COSTANOMALIESRESPONSEDATATYPE_ANOMALIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostAnomaliesResponseDataType) GetAllowedValues() []CostAnomaliesResponseDataType {
	return allowedCostAnomaliesResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostAnomaliesResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostAnomaliesResponseDataType(value)
	return nil
}

// NewCostAnomaliesResponseDataTypeFromValue returns a pointer to a valid CostAnomaliesResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostAnomaliesResponseDataTypeFromValue(v string) (*CostAnomaliesResponseDataType, error) {
	ev := CostAnomaliesResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostAnomaliesResponseDataType: valid values are %v", v, allowedCostAnomaliesResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostAnomaliesResponseDataType) IsValid() bool {
	for _, existing := range allowedCostAnomaliesResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostAnomaliesResponseDataType value.
func (v CostAnomaliesResponseDataType) Ptr() *CostAnomaliesResponseDataType {
	return &v
}
