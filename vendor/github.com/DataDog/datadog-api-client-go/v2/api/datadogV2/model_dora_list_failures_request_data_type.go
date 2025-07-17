// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORAListFailuresRequestDataType The definition of `DORAListFailuresRequestDataType` object.
type DORAListFailuresRequestDataType string

// List of DORAListFailuresRequestDataType.
const (
	DORALISTFAILURESREQUESTDATATYPE_DORA_FAILURES_LIST_REQUEST DORAListFailuresRequestDataType = "dora_failures_list_request"
)

var allowedDORAListFailuresRequestDataTypeEnumValues = []DORAListFailuresRequestDataType{
	DORALISTFAILURESREQUESTDATATYPE_DORA_FAILURES_LIST_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DORAListFailuresRequestDataType) GetAllowedValues() []DORAListFailuresRequestDataType {
	return allowedDORAListFailuresRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DORAListFailuresRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DORAListFailuresRequestDataType(value)
	return nil
}

// NewDORAListFailuresRequestDataTypeFromValue returns a pointer to a valid DORAListFailuresRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDORAListFailuresRequestDataTypeFromValue(v string) (*DORAListFailuresRequestDataType, error) {
	ev := DORAListFailuresRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DORAListFailuresRequestDataType: valid values are %v", v, allowedDORAListFailuresRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DORAListFailuresRequestDataType) IsValid() bool {
	for _, existing := range allowedDORAListFailuresRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DORAListFailuresRequestDataType value.
func (v DORAListFailuresRequestDataType) Ptr() *DORAListFailuresRequestDataType {
	return &v
}
