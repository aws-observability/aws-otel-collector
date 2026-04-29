// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateComponentRequestDataAttributesType The type of the component.
type CreateComponentRequestDataAttributesType string

// List of CreateComponentRequestDataAttributesType.
const (
	CREATECOMPONENTREQUESTDATAATTRIBUTESTYPE_COMPONENT CreateComponentRequestDataAttributesType = "component"
	CREATECOMPONENTREQUESTDATAATTRIBUTESTYPE_GROUP     CreateComponentRequestDataAttributesType = "group"
)

var allowedCreateComponentRequestDataAttributesTypeEnumValues = []CreateComponentRequestDataAttributesType{
	CREATECOMPONENTREQUESTDATAATTRIBUTESTYPE_COMPONENT,
	CREATECOMPONENTREQUESTDATAATTRIBUTESTYPE_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateComponentRequestDataAttributesType) GetAllowedValues() []CreateComponentRequestDataAttributesType {
	return allowedCreateComponentRequestDataAttributesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateComponentRequestDataAttributesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateComponentRequestDataAttributesType(value)
	return nil
}

// NewCreateComponentRequestDataAttributesTypeFromValue returns a pointer to a valid CreateComponentRequestDataAttributesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateComponentRequestDataAttributesTypeFromValue(v string) (*CreateComponentRequestDataAttributesType, error) {
	ev := CreateComponentRequestDataAttributesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateComponentRequestDataAttributesType: valid values are %v", v, allowedCreateComponentRequestDataAttributesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateComponentRequestDataAttributesType) IsValid() bool {
	for _, existing := range allowedCreateComponentRequestDataAttributesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateComponentRequestDataAttributesType value.
func (v CreateComponentRequestDataAttributesType) Ptr() *CreateComponentRequestDataAttributesType {
	return &v
}
