// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreatePageResponseDataType The type of resource used when creating an On-Call Page.
type CreatePageResponseDataType string

// List of CreatePageResponseDataType.
const (
	CREATEPAGERESPONSEDATATYPE_PAGES CreatePageResponseDataType = "pages"
)

var allowedCreatePageResponseDataTypeEnumValues = []CreatePageResponseDataType{
	CREATEPAGERESPONSEDATATYPE_PAGES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreatePageResponseDataType) GetAllowedValues() []CreatePageResponseDataType {
	return allowedCreatePageResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreatePageResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreatePageResponseDataType(value)
	return nil
}

// NewCreatePageResponseDataTypeFromValue returns a pointer to a valid CreatePageResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreatePageResponseDataTypeFromValue(v string) (*CreatePageResponseDataType, error) {
	ev := CreatePageResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreatePageResponseDataType: valid values are %v", v, allowedCreatePageResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreatePageResponseDataType) IsValid() bool {
	for _, existing := range allowedCreatePageResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreatePageResponseDataType value.
func (v CreatePageResponseDataType) Ptr() *CreatePageResponseDataType {
	return &v
}
