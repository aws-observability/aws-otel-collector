// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedLongTasksRequestType The JSON:API type for aggregated long tasks requests.
type AggregatedLongTasksRequestType string

// List of AggregatedLongTasksRequestType.
const (
	AGGREGATEDLONGTASKSREQUESTTYPE_AGGREGATED_LONG_TASKS AggregatedLongTasksRequestType = "aggregated_long_tasks"
)

var allowedAggregatedLongTasksRequestTypeEnumValues = []AggregatedLongTasksRequestType{
	AGGREGATEDLONGTASKSREQUESTTYPE_AGGREGATED_LONG_TASKS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AggregatedLongTasksRequestType) GetAllowedValues() []AggregatedLongTasksRequestType {
	return allowedAggregatedLongTasksRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregatedLongTasksRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregatedLongTasksRequestType(value)
	return nil
}

// NewAggregatedLongTasksRequestTypeFromValue returns a pointer to a valid AggregatedLongTasksRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregatedLongTasksRequestTypeFromValue(v string) (*AggregatedLongTasksRequestType, error) {
	ev := AggregatedLongTasksRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregatedLongTasksRequestType: valid values are %v", v, allowedAggregatedLongTasksRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregatedLongTasksRequestType) IsValid() bool {
	for _, existing := range allowedAggregatedLongTasksRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregatedLongTasksRequestType value.
func (v AggregatedLongTasksRequestType) Ptr() *AggregatedLongTasksRequestType {
	return &v
}
