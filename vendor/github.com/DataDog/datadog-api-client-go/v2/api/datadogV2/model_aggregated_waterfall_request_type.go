// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedWaterfallRequestType The JSON:API type for aggregated waterfall requests.
type AggregatedWaterfallRequestType string

// List of AggregatedWaterfallRequestType.
const (
	AGGREGATEDWATERFALLREQUESTTYPE_AGGREGATED_WATERFALL AggregatedWaterfallRequestType = "aggregated_waterfall"
)

var allowedAggregatedWaterfallRequestTypeEnumValues = []AggregatedWaterfallRequestType{
	AGGREGATEDWATERFALLREQUESTTYPE_AGGREGATED_WATERFALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AggregatedWaterfallRequestType) GetAllowedValues() []AggregatedWaterfallRequestType {
	return allowedAggregatedWaterfallRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregatedWaterfallRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregatedWaterfallRequestType(value)
	return nil
}

// NewAggregatedWaterfallRequestTypeFromValue returns a pointer to a valid AggregatedWaterfallRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregatedWaterfallRequestTypeFromValue(v string) (*AggregatedWaterfallRequestType, error) {
	ev := AggregatedWaterfallRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregatedWaterfallRequestType: valid values are %v", v, allowedAggregatedWaterfallRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregatedWaterfallRequestType) IsValid() bool {
	for _, existing := range allowedAggregatedWaterfallRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregatedWaterfallRequestType value.
func (v AggregatedWaterfallRequestType) Ptr() *AggregatedWaterfallRequestType {
	return &v
}
