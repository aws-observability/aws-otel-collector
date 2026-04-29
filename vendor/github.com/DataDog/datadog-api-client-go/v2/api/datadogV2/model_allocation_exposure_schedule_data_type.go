// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AllocationExposureScheduleDataType The resource type for progressive rollout schedules.
type AllocationExposureScheduleDataType string

// List of AllocationExposureScheduleDataType.
const (
	ALLOCATIONEXPOSURESCHEDULEDATATYPE_ALLOCATION_EXPOSURE_SCHEDULES AllocationExposureScheduleDataType = "allocation_exposure_schedules"
)

var allowedAllocationExposureScheduleDataTypeEnumValues = []AllocationExposureScheduleDataType{
	ALLOCATIONEXPOSURESCHEDULEDATATYPE_ALLOCATION_EXPOSURE_SCHEDULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AllocationExposureScheduleDataType) GetAllowedValues() []AllocationExposureScheduleDataType {
	return allowedAllocationExposureScheduleDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AllocationExposureScheduleDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AllocationExposureScheduleDataType(value)
	return nil
}

// NewAllocationExposureScheduleDataTypeFromValue returns a pointer to a valid AllocationExposureScheduleDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAllocationExposureScheduleDataTypeFromValue(v string) (*AllocationExposureScheduleDataType, error) {
	ev := AllocationExposureScheduleDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AllocationExposureScheduleDataType: valid values are %v", v, allowedAllocationExposureScheduleDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AllocationExposureScheduleDataType) IsValid() bool {
	for _, existing := range allowedAllocationExposureScheduleDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AllocationExposureScheduleDataType value.
func (v AllocationExposureScheduleDataType) Ptr() *AllocationExposureScheduleDataType {
	return &v
}
