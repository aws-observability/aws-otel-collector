// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssetEntityType The JSON:API type.
type AssetEntityType string

// List of AssetEntityType.
const (
	ASSETENTITYTYPE_ASSETS AssetEntityType = "assets"
)

var allowedAssetEntityTypeEnumValues = []AssetEntityType{
	ASSETENTITYTYPE_ASSETS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AssetEntityType) GetAllowedValues() []AssetEntityType {
	return allowedAssetEntityTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AssetEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AssetEntityType(value)
	return nil
}

// NewAssetEntityTypeFromValue returns a pointer to a valid AssetEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAssetEntityTypeFromValue(v string) (*AssetEntityType, error) {
	ev := AssetEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AssetEntityType: valid values are %v", v, allowedAssetEntityTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AssetEntityType) IsValid() bool {
	for _, existing := range allowedAssetEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssetEntityType value.
func (v AssetEntityType) Ptr() *AssetEntityType {
	return &v
}
