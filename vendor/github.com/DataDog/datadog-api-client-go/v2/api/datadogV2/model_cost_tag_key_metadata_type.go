// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagKeyMetadataType Type of the Cloud Cost Management tag key metadata resource.
type CostTagKeyMetadataType string

// List of CostTagKeyMetadataType.
const (
	COSTTAGKEYMETADATATYPE_COST_TAG_KEY_METADATA CostTagKeyMetadataType = "cost_tag_key_metadata"
)

var allowedCostTagKeyMetadataTypeEnumValues = []CostTagKeyMetadataType{
	COSTTAGKEYMETADATATYPE_COST_TAG_KEY_METADATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostTagKeyMetadataType) GetAllowedValues() []CostTagKeyMetadataType {
	return allowedCostTagKeyMetadataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostTagKeyMetadataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostTagKeyMetadataType(value)
	return nil
}

// NewCostTagKeyMetadataTypeFromValue returns a pointer to a valid CostTagKeyMetadataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostTagKeyMetadataTypeFromValue(v string) (*CostTagKeyMetadataType, error) {
	ev := CostTagKeyMetadataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostTagKeyMetadataType: valid values are %v", v, allowedCostTagKeyMetadataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostTagKeyMetadataType) IsValid() bool {
	for _, existing := range allowedCostTagKeyMetadataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostTagKeyMetadataType value.
func (v CostTagKeyMetadataType) Ptr() *CostTagKeyMetadataType {
	return &v
}
