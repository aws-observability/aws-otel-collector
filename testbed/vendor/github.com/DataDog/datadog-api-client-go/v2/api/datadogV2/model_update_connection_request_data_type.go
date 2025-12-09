// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateConnectionRequestDataType Connection id resource type.
type UpdateConnectionRequestDataType string

// List of UpdateConnectionRequestDataType.
const (
	UPDATECONNECTIONREQUESTDATATYPE_CONNECTION_ID UpdateConnectionRequestDataType = "connection_id"
)

var allowedUpdateConnectionRequestDataTypeEnumValues = []UpdateConnectionRequestDataType{
	UPDATECONNECTIONREQUESTDATATYPE_CONNECTION_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateConnectionRequestDataType) GetAllowedValues() []UpdateConnectionRequestDataType {
	return allowedUpdateConnectionRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateConnectionRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateConnectionRequestDataType(value)
	return nil
}

// NewUpdateConnectionRequestDataTypeFromValue returns a pointer to a valid UpdateConnectionRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateConnectionRequestDataTypeFromValue(v string) (*UpdateConnectionRequestDataType, error) {
	ev := UpdateConnectionRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateConnectionRequestDataType: valid values are %v", v, allowedUpdateConnectionRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateConnectionRequestDataType) IsValid() bool {
	for _, existing := range allowedUpdateConnectionRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateConnectionRequestDataType value.
func (v UpdateConnectionRequestDataType) Ptr() *UpdateConnectionRequestDataType {
	return &v
}
