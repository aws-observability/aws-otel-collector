// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsScalarColumnType The column type. "group" for dimension columns, "number" for metric columns.
type CommitmentsScalarColumnType string

// List of CommitmentsScalarColumnType.
const (
	COMMITMENTSSCALARCOLUMNTYPE_GROUP  CommitmentsScalarColumnType = "group"
	COMMITMENTSSCALARCOLUMNTYPE_NUMBER CommitmentsScalarColumnType = "number"
)

var allowedCommitmentsScalarColumnTypeEnumValues = []CommitmentsScalarColumnType{
	COMMITMENTSSCALARCOLUMNTYPE_GROUP,
	COMMITMENTSSCALARCOLUMNTYPE_NUMBER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitmentsScalarColumnType) GetAllowedValues() []CommitmentsScalarColumnType {
	return allowedCommitmentsScalarColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitmentsScalarColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitmentsScalarColumnType(value)
	return nil
}

// NewCommitmentsScalarColumnTypeFromValue returns a pointer to a valid CommitmentsScalarColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitmentsScalarColumnTypeFromValue(v string) (*CommitmentsScalarColumnType, error) {
	ev := CommitmentsScalarColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitmentsScalarColumnType: valid values are %v", v, allowedCommitmentsScalarColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitmentsScalarColumnType) IsValid() bool {
	for _, existing := range allowedCommitmentsScalarColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitmentsScalarColumnType value.
func (v CommitmentsScalarColumnType) Ptr() *CommitmentsScalarColumnType {
	return &v
}
