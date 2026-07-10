// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueWidgetComparisonType How the delta is expressed: `absolute` (raw difference), `relative` (percentage), or `both`.
type QueryValueWidgetComparisonType string

// List of QueryValueWidgetComparisonType.
const (
	QUERYVALUEWIDGETCOMPARISONTYPE_ABSOLUTE QueryValueWidgetComparisonType = "absolute"
	QUERYVALUEWIDGETCOMPARISONTYPE_RELATIVE QueryValueWidgetComparisonType = "relative"
	QUERYVALUEWIDGETCOMPARISONTYPE_BOTH     QueryValueWidgetComparisonType = "both"
)

var allowedQueryValueWidgetComparisonTypeEnumValues = []QueryValueWidgetComparisonType{
	QUERYVALUEWIDGETCOMPARISONTYPE_ABSOLUTE,
	QUERYVALUEWIDGETCOMPARISONTYPE_RELATIVE,
	QUERYVALUEWIDGETCOMPARISONTYPE_BOTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryValueWidgetComparisonType) GetAllowedValues() []QueryValueWidgetComparisonType {
	return allowedQueryValueWidgetComparisonTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryValueWidgetComparisonType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryValueWidgetComparisonType(value)
	return nil
}

// NewQueryValueWidgetComparisonTypeFromValue returns a pointer to a valid QueryValueWidgetComparisonType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryValueWidgetComparisonTypeFromValue(v string) (*QueryValueWidgetComparisonType, error) {
	ev := QueryValueWidgetComparisonType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryValueWidgetComparisonType: valid values are %v", v, allowedQueryValueWidgetComparisonTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryValueWidgetComparisonType) IsValid() bool {
	for _, existing := range allowedQueryValueWidgetComparisonTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryValueWidgetComparisonType value.
func (v QueryValueWidgetComparisonType) Ptr() *QueryValueWidgetComparisonType {
	return &v
}
