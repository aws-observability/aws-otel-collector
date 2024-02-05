// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOListWidgetRequestType Widget request type.
type SLOListWidgetRequestType string

// List of SLOListWidgetRequestType.
const (
	SLOLISTWIDGETREQUESTTYPE_SLO_LIST SLOListWidgetRequestType = "slo_list"
)

var allowedSLOListWidgetRequestTypeEnumValues = []SLOListWidgetRequestType{
	SLOLISTWIDGETREQUESTTYPE_SLO_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOListWidgetRequestType) GetAllowedValues() []SLOListWidgetRequestType {
	return allowedSLOListWidgetRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOListWidgetRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOListWidgetRequestType(value)
	return nil
}

// NewSLOListWidgetRequestTypeFromValue returns a pointer to a valid SLOListWidgetRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOListWidgetRequestTypeFromValue(v string) (*SLOListWidgetRequestType, error) {
	ev := SLOListWidgetRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOListWidgetRequestType: valid values are %v", v, allowedSLOListWidgetRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOListWidgetRequestType) IsValid() bool {
	for _, existing := range allowedSLOListWidgetRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOListWidgetRequestType value.
func (v SLOListWidgetRequestType) Ptr() *SLOListWidgetRequestType {
	return &v
}
