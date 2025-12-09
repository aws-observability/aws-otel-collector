// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryCostUpsertRequestDataType Upsert arbitrary rule resource type.
type ArbitraryCostUpsertRequestDataType string

// List of ArbitraryCostUpsertRequestDataType.
const (
	ARBITRARYCOSTUPSERTREQUESTDATATYPE_UPSERT_ARBITRARY_RULE ArbitraryCostUpsertRequestDataType = "upsert_arbitrary_rule"
)

var allowedArbitraryCostUpsertRequestDataTypeEnumValues = []ArbitraryCostUpsertRequestDataType{
	ARBITRARYCOSTUPSERTREQUESTDATATYPE_UPSERT_ARBITRARY_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ArbitraryCostUpsertRequestDataType) GetAllowedValues() []ArbitraryCostUpsertRequestDataType {
	return allowedArbitraryCostUpsertRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ArbitraryCostUpsertRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ArbitraryCostUpsertRequestDataType(value)
	return nil
}

// NewArbitraryCostUpsertRequestDataTypeFromValue returns a pointer to a valid ArbitraryCostUpsertRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewArbitraryCostUpsertRequestDataTypeFromValue(v string) (*ArbitraryCostUpsertRequestDataType, error) {
	ev := ArbitraryCostUpsertRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ArbitraryCostUpsertRequestDataType: valid values are %v", v, allowedArbitraryCostUpsertRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ArbitraryCostUpsertRequestDataType) IsValid() bool {
	for _, existing := range allowedArbitraryCostUpsertRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ArbitraryCostUpsertRequestDataType value.
func (v ArbitraryCostUpsertRequestDataType) Ptr() *ArbitraryCostUpsertRequestDataType {
	return &v
}
