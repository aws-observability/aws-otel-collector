// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LicensesListResponseDataType The type identifier for license list responses.
type LicensesListResponseDataType string

// List of LicensesListResponseDataType.
const (
	LICENSESLISTRESPONSEDATATYPE_LICENSEREQUEST LicensesListResponseDataType = "licenserequest"
)

var allowedLicensesListResponseDataTypeEnumValues = []LicensesListResponseDataType{
	LICENSESLISTRESPONSEDATATYPE_LICENSEREQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LicensesListResponseDataType) GetAllowedValues() []LicensesListResponseDataType {
	return allowedLicensesListResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LicensesListResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LicensesListResponseDataType(value)
	return nil
}

// NewLicensesListResponseDataTypeFromValue returns a pointer to a valid LicensesListResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLicensesListResponseDataTypeFromValue(v string) (*LicensesListResponseDataType, error) {
	ev := LicensesListResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LicensesListResponseDataType: valid values are %v", v, allowedLicensesListResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LicensesListResponseDataType) IsValid() bool {
	for _, existing := range allowedLicensesListResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicensesListResponseDataType value.
func (v LicensesListResponseDataType) Ptr() *LicensesListResponseDataType {
	return &v
}
