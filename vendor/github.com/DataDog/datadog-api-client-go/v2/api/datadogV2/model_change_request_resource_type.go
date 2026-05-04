// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestResourceType Change request resource type.
type ChangeRequestResourceType string

// List of ChangeRequestResourceType.
const (
	CHANGEREQUESTRESOURCETYPE_CHANGE_REQUEST ChangeRequestResourceType = "change_request"
)

var allowedChangeRequestResourceTypeEnumValues = []ChangeRequestResourceType{
	CHANGEREQUESTRESOURCETYPE_CHANGE_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeRequestResourceType) GetAllowedValues() []ChangeRequestResourceType {
	return allowedChangeRequestResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeRequestResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeRequestResourceType(value)
	return nil
}

// NewChangeRequestResourceTypeFromValue returns a pointer to a valid ChangeRequestResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeRequestResourceTypeFromValue(v string) (*ChangeRequestResourceType, error) {
	ev := ChangeRequestResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeRequestResourceType: valid values are %v", v, allowedChangeRequestResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeRequestResourceType) IsValid() bool {
	for _, existing := range allowedChangeRequestResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeRequestResourceType value.
func (v ChangeRequestResourceType) Ptr() *ChangeRequestResourceType {
	return &v
}
