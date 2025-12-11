// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSchemaCategoryMapperType Type of logs schema category mapper.
type LogsSchemaCategoryMapperType string

// List of LogsSchemaCategoryMapperType.
const (
	LOGSSCHEMACATEGORYMAPPERTYPE_SCHEMA_CATEGORY_MAPPER LogsSchemaCategoryMapperType = "schema-category-mapper"
)

var allowedLogsSchemaCategoryMapperTypeEnumValues = []LogsSchemaCategoryMapperType{
	LOGSSCHEMACATEGORYMAPPERTYPE_SCHEMA_CATEGORY_MAPPER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsSchemaCategoryMapperType) GetAllowedValues() []LogsSchemaCategoryMapperType {
	return allowedLogsSchemaCategoryMapperTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsSchemaCategoryMapperType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsSchemaCategoryMapperType(value)
	return nil
}

// NewLogsSchemaCategoryMapperTypeFromValue returns a pointer to a valid LogsSchemaCategoryMapperType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsSchemaCategoryMapperTypeFromValue(v string) (*LogsSchemaCategoryMapperType, error) {
	ev := LogsSchemaCategoryMapperType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsSchemaCategoryMapperType: valid values are %v", v, allowedLogsSchemaCategoryMapperTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsSchemaCategoryMapperType) IsValid() bool {
	for _, existing := range allowedLogsSchemaCategoryMapperTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsSchemaCategoryMapperType value.
func (v LogsSchemaCategoryMapperType) Ptr() *LogsSchemaCategoryMapperType {
	return &v
}
