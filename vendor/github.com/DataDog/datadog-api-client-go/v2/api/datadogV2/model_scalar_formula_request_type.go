// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarFormulaRequestType The type of the resource. The value should always be scalar_request.
type ScalarFormulaRequestType string

// List of ScalarFormulaRequestType.
const (
	SCALARFORMULAREQUESTTYPE_SCALAR_REQUEST ScalarFormulaRequestType = "scalar_request"
)

var allowedScalarFormulaRequestTypeEnumValues = []ScalarFormulaRequestType{
	SCALARFORMULAREQUESTTYPE_SCALAR_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScalarFormulaRequestType) GetAllowedValues() []ScalarFormulaRequestType {
	return allowedScalarFormulaRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScalarFormulaRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScalarFormulaRequestType(value)
	return nil
}

// NewScalarFormulaRequestTypeFromValue returns a pointer to a valid ScalarFormulaRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScalarFormulaRequestTypeFromValue(v string) (*ScalarFormulaRequestType, error) {
	ev := ScalarFormulaRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScalarFormulaRequestType: valid values are %v", v, allowedScalarFormulaRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScalarFormulaRequestType) IsValid() bool {
	for _, existing := range allowedScalarFormulaRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScalarFormulaRequestType value.
func (v ScalarFormulaRequestType) Ptr() *ScalarFormulaRequestType {
	return &v
}
