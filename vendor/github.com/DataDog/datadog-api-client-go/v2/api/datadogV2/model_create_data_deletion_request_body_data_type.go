// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateDataDeletionRequestBodyDataType The deletion request type.
type CreateDataDeletionRequestBodyDataType string

// List of CreateDataDeletionRequestBodyDataType.
const (
	CREATEDATADELETIONREQUESTBODYDATATYPE_CREATE_DELETION_REQ CreateDataDeletionRequestBodyDataType = "create_deletion_req"
)

var allowedCreateDataDeletionRequestBodyDataTypeEnumValues = []CreateDataDeletionRequestBodyDataType{
	CREATEDATADELETIONREQUESTBODYDATATYPE_CREATE_DELETION_REQ,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateDataDeletionRequestBodyDataType) GetAllowedValues() []CreateDataDeletionRequestBodyDataType {
	return allowedCreateDataDeletionRequestBodyDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateDataDeletionRequestBodyDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateDataDeletionRequestBodyDataType(value)
	return nil
}

// NewCreateDataDeletionRequestBodyDataTypeFromValue returns a pointer to a valid CreateDataDeletionRequestBodyDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateDataDeletionRequestBodyDataTypeFromValue(v string) (*CreateDataDeletionRequestBodyDataType, error) {
	ev := CreateDataDeletionRequestBodyDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateDataDeletionRequestBodyDataType: valid values are %v", v, allowedCreateDataDeletionRequestBodyDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateDataDeletionRequestBodyDataType) IsValid() bool {
	for _, existing := range allowedCreateDataDeletionRequestBodyDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateDataDeletionRequestBodyDataType value.
func (v CreateDataDeletionRequestBodyDataType) Ptr() *CreateDataDeletionRequestBodyDataType {
	return &v
}
