// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppCreatePipelineEventRequestDataType Type of the event.
type CIAppCreatePipelineEventRequestDataType string

// List of CIAppCreatePipelineEventRequestDataType.
const (
	CIAPPCREATEPIPELINEEVENTREQUESTDATATYPE_CIPIPELINE_RESOURCE_REQUEST CIAppCreatePipelineEventRequestDataType = "cipipeline_resource_request"
)

var allowedCIAppCreatePipelineEventRequestDataTypeEnumValues = []CIAppCreatePipelineEventRequestDataType{
	CIAPPCREATEPIPELINEEVENTREQUESTDATATYPE_CIPIPELINE_RESOURCE_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppCreatePipelineEventRequestDataType) GetAllowedValues() []CIAppCreatePipelineEventRequestDataType {
	return allowedCIAppCreatePipelineEventRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppCreatePipelineEventRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppCreatePipelineEventRequestDataType(value)
	return nil
}

// NewCIAppCreatePipelineEventRequestDataTypeFromValue returns a pointer to a valid CIAppCreatePipelineEventRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppCreatePipelineEventRequestDataTypeFromValue(v string) (*CIAppCreatePipelineEventRequestDataType, error) {
	ev := CIAppCreatePipelineEventRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppCreatePipelineEventRequestDataType: valid values are %v", v, allowedCIAppCreatePipelineEventRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppCreatePipelineEventRequestDataType) IsValid() bool {
	for _, existing := range allowedCIAppCreatePipelineEventRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppCreatePipelineEventRequestDataType value.
func (v CIAppCreatePipelineEventRequestDataType) Ptr() *CIAppCreatePipelineEventRequestDataType {
	return &v
}
