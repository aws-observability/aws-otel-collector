// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateStatusPageRequestDataAttributesType The type of the status page controlling how the status page is accessed.
type CreateStatusPageRequestDataAttributesType string

// List of CreateStatusPageRequestDataAttributesType.
const (
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESTYPE_PUBLIC   CreateStatusPageRequestDataAttributesType = "public"
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESTYPE_INTERNAL CreateStatusPageRequestDataAttributesType = "internal"
)

var allowedCreateStatusPageRequestDataAttributesTypeEnumValues = []CreateStatusPageRequestDataAttributesType{
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESTYPE_PUBLIC,
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESTYPE_INTERNAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateStatusPageRequestDataAttributesType) GetAllowedValues() []CreateStatusPageRequestDataAttributesType {
	return allowedCreateStatusPageRequestDataAttributesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateStatusPageRequestDataAttributesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateStatusPageRequestDataAttributesType(value)
	return nil
}

// NewCreateStatusPageRequestDataAttributesTypeFromValue returns a pointer to a valid CreateStatusPageRequestDataAttributesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateStatusPageRequestDataAttributesTypeFromValue(v string) (*CreateStatusPageRequestDataAttributesType, error) {
	ev := CreateStatusPageRequestDataAttributesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateStatusPageRequestDataAttributesType: valid values are %v", v, allowedCreateStatusPageRequestDataAttributesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateStatusPageRequestDataAttributesType) IsValid() bool {
	for _, existing := range allowedCreateStatusPageRequestDataAttributesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateStatusPageRequestDataAttributesType value.
func (v CreateStatusPageRequestDataAttributesType) Ptr() *CreateStatusPageRequestDataAttributesType {
	return &v
}
