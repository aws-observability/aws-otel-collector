// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssetType The asset type
type AssetType string

// List of AssetType.
const (
	ASSETTYPE_REPOSITORY AssetType = "Repository"
	ASSETTYPE_SERVICE    AssetType = "Service"
	ASSETTYPE_HOST       AssetType = "Host"
	ASSETTYPE_HOSTIMAGE  AssetType = "HostImage"
	ASSETTYPE_IMAGE      AssetType = "Image"
)

var allowedAssetTypeEnumValues = []AssetType{
	ASSETTYPE_REPOSITORY,
	ASSETTYPE_SERVICE,
	ASSETTYPE_HOST,
	ASSETTYPE_HOSTIMAGE,
	ASSETTYPE_IMAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AssetType) GetAllowedValues() []AssetType {
	return allowedAssetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AssetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AssetType(value)
	return nil
}

// NewAssetTypeFromValue returns a pointer to a valid AssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAssetTypeFromValue(v string) (*AssetType, error) {
	ev := AssetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AssetType: valid values are %v", v, allowedAssetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AssetType) IsValid() bool {
	for _, existing := range allowedAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssetType value.
func (v AssetType) Ptr() *AssetType {
	return &v
}
