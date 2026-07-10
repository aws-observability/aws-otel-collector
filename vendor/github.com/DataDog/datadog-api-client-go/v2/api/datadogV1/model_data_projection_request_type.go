// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataProjectionRequestType Type of a data projection request.
type DataProjectionRequestType string

// List of DataProjectionRequestType.
const (
	DATAPROJECTIONREQUESTTYPE_DATA_PROJECTION DataProjectionRequestType = "data_projection"
)

var allowedDataProjectionRequestTypeEnumValues = []DataProjectionRequestType{
	DATAPROJECTIONREQUESTTYPE_DATA_PROJECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DataProjectionRequestType) GetAllowedValues() []DataProjectionRequestType {
	return allowedDataProjectionRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataProjectionRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataProjectionRequestType(value)
	return nil
}

// NewDataProjectionRequestTypeFromValue returns a pointer to a valid DataProjectionRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataProjectionRequestTypeFromValue(v string) (*DataProjectionRequestType, error) {
	ev := DataProjectionRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataProjectionRequestType: valid values are %v", v, allowedDataProjectionRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataProjectionRequestType) IsValid() bool {
	for _, existing := range allowedDataProjectionRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataProjectionRequestType value.
func (v DataProjectionRequestType) Ptr() *DataProjectionRequestType {
	return &v
}
