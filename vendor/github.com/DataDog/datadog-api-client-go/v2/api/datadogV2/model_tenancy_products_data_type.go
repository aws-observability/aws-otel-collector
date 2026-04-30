// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TenancyProductsDataType OCI tenancy product resource type.
type TenancyProductsDataType string

// List of TenancyProductsDataType.
const (
	TENANCYPRODUCTSDATATYPE_OCI_TENANCY_PRODUCT TenancyProductsDataType = "oci_tenancy_product"
)

var allowedTenancyProductsDataTypeEnumValues = []TenancyProductsDataType{
	TENANCYPRODUCTSDATATYPE_OCI_TENANCY_PRODUCT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TenancyProductsDataType) GetAllowedValues() []TenancyProductsDataType {
	return allowedTenancyProductsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TenancyProductsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TenancyProductsDataType(value)
	return nil
}

// NewTenancyProductsDataTypeFromValue returns a pointer to a valid TenancyProductsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTenancyProductsDataTypeFromValue(v string) (*TenancyProductsDataType, error) {
	ev := TenancyProductsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TenancyProductsDataType: valid values are %v", v, allowedTenancyProductsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TenancyProductsDataType) IsValid() bool {
	for _, existing := range allowedTenancyProductsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TenancyProductsDataType value.
func (v TenancyProductsDataType) Ptr() *TenancyProductsDataType {
	return &v
}
