// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataAttributesRulesItemsIfTagExists The behavior when the tag already exists.
type DataAttributesRulesItemsIfTagExists string

// List of DataAttributesRulesItemsIfTagExists.
const (
	DATAATTRIBUTESRULESITEMSIFTAGEXISTS_APPEND       DataAttributesRulesItemsIfTagExists = "append"
	DATAATTRIBUTESRULESITEMSIFTAGEXISTS_DO_NOT_APPLY DataAttributesRulesItemsIfTagExists = "do_not_apply"
	DATAATTRIBUTESRULESITEMSIFTAGEXISTS_REPLACE      DataAttributesRulesItemsIfTagExists = "replace"
)

var allowedDataAttributesRulesItemsIfTagExistsEnumValues = []DataAttributesRulesItemsIfTagExists{
	DATAATTRIBUTESRULESITEMSIFTAGEXISTS_APPEND,
	DATAATTRIBUTESRULESITEMSIFTAGEXISTS_DO_NOT_APPLY,
	DATAATTRIBUTESRULESITEMSIFTAGEXISTS_REPLACE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DataAttributesRulesItemsIfTagExists) GetAllowedValues() []DataAttributesRulesItemsIfTagExists {
	return allowedDataAttributesRulesItemsIfTagExistsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataAttributesRulesItemsIfTagExists) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataAttributesRulesItemsIfTagExists(value)
	return nil
}

// NewDataAttributesRulesItemsIfTagExistsFromValue returns a pointer to a valid DataAttributesRulesItemsIfTagExists
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataAttributesRulesItemsIfTagExistsFromValue(v string) (*DataAttributesRulesItemsIfTagExists, error) {
	ev := DataAttributesRulesItemsIfTagExists(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataAttributesRulesItemsIfTagExists: valid values are %v", v, allowedDataAttributesRulesItemsIfTagExistsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataAttributesRulesItemsIfTagExists) IsValid() bool {
	for _, existing := range allowedDataAttributesRulesItemsIfTagExistsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataAttributesRulesItemsIfTagExists value.
func (v DataAttributesRulesItemsIfTagExists) Ptr() *DataAttributesRulesItemsIfTagExists {
	return &v
}
