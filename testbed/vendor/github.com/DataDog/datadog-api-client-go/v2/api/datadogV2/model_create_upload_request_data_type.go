// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateUploadRequestDataType Upload resource type.
type CreateUploadRequestDataType string

// List of CreateUploadRequestDataType.
const (
	CREATEUPLOADREQUESTDATATYPE_UPLOAD CreateUploadRequestDataType = "upload"
)

var allowedCreateUploadRequestDataTypeEnumValues = []CreateUploadRequestDataType{
	CREATEUPLOADREQUESTDATATYPE_UPLOAD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateUploadRequestDataType) GetAllowedValues() []CreateUploadRequestDataType {
	return allowedCreateUploadRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateUploadRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateUploadRequestDataType(value)
	return nil
}

// NewCreateUploadRequestDataTypeFromValue returns a pointer to a valid CreateUploadRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateUploadRequestDataTypeFromValue(v string) (*CreateUploadRequestDataType, error) {
	ev := CreateUploadRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateUploadRequestDataType: valid values are %v", v, allowedCreateUploadRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateUploadRequestDataType) IsValid() bool {
	for _, existing := range allowedCreateUploadRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateUploadRequestDataType value.
func (v CreateUploadRequestDataType) Ptr() *CreateUploadRequestDataType {
	return &v
}
