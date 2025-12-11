// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateUploadResponseDataType Upload resource type.
type CreateUploadResponseDataType string

// List of CreateUploadResponseDataType.
const (
	CREATEUPLOADRESPONSEDATATYPE_UPLOAD CreateUploadResponseDataType = "upload"
)

var allowedCreateUploadResponseDataTypeEnumValues = []CreateUploadResponseDataType{
	CREATEUPLOADRESPONSEDATATYPE_UPLOAD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateUploadResponseDataType) GetAllowedValues() []CreateUploadResponseDataType {
	return allowedCreateUploadResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateUploadResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateUploadResponseDataType(value)
	return nil
}

// NewCreateUploadResponseDataTypeFromValue returns a pointer to a valid CreateUploadResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateUploadResponseDataTypeFromValue(v string) (*CreateUploadResponseDataType, error) {
	ev := CreateUploadResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateUploadResponseDataType: valid values are %v", v, allowedCreateUploadResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateUploadResponseDataType) IsValid() bool {
	for _, existing := range allowedCreateUploadResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateUploadResponseDataType value.
func (v CreateUploadResponseDataType) Ptr() *CreateUploadResponseDataType {
	return &v
}
