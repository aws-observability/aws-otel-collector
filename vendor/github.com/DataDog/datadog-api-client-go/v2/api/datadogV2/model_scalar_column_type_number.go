// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarColumnTypeNumber The type of column present for numbers.
type ScalarColumnTypeNumber string

// List of ScalarColumnTypeNumber.
const (
	SCALARCOLUMNTYPENUMBER_NUMBER ScalarColumnTypeNumber = "number"
)

var allowedScalarColumnTypeNumberEnumValues = []ScalarColumnTypeNumber{
	SCALARCOLUMNTYPENUMBER_NUMBER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScalarColumnTypeNumber) GetAllowedValues() []ScalarColumnTypeNumber {
	return allowedScalarColumnTypeNumberEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScalarColumnTypeNumber) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScalarColumnTypeNumber(value)
	return nil
}

// NewScalarColumnTypeNumberFromValue returns a pointer to a valid ScalarColumnTypeNumber
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScalarColumnTypeNumberFromValue(v string) (*ScalarColumnTypeNumber, error) {
	ev := ScalarColumnTypeNumber(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScalarColumnTypeNumber: valid values are %v", v, allowedScalarColumnTypeNumberEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScalarColumnTypeNumber) IsValid() bool {
	for _, existing := range allowedScalarColumnTypeNumberEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScalarColumnTypeNumber value.
func (v ScalarColumnTypeNumber) Ptr() *ScalarColumnTypeNumber {
	return &v
}
