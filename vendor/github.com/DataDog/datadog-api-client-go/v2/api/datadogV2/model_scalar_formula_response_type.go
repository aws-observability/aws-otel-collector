// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarFormulaResponseType The type of the resource. The value should always be scalar_response.
type ScalarFormulaResponseType string

// List of ScalarFormulaResponseType.
const (
	SCALARFORMULARESPONSETYPE_SCALAR_RESPONSE ScalarFormulaResponseType = "scalar_response"
)

var allowedScalarFormulaResponseTypeEnumValues = []ScalarFormulaResponseType{
	SCALARFORMULARESPONSETYPE_SCALAR_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScalarFormulaResponseType) GetAllowedValues() []ScalarFormulaResponseType {
	return allowedScalarFormulaResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScalarFormulaResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScalarFormulaResponseType(value)
	return nil
}

// NewScalarFormulaResponseTypeFromValue returns a pointer to a valid ScalarFormulaResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScalarFormulaResponseTypeFromValue(v string) (*ScalarFormulaResponseType, error) {
	ev := ScalarFormulaResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScalarFormulaResponseType: valid values are %v", v, allowedScalarFormulaResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScalarFormulaResponseType) IsValid() bool {
	for _, existing := range allowedScalarFormulaResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScalarFormulaResponseType value.
func (v ScalarFormulaResponseType) Ptr() *ScalarFormulaResponseType {
	return &v
}
