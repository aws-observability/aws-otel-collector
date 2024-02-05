// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateRequestType The type of resource. The value should always be aggregate_request.
type SpansAggregateRequestType string

// List of SpansAggregateRequestType.
const (
	SPANSAGGREGATEREQUESTTYPE_AGGREGATE_REQUEST SpansAggregateRequestType = "aggregate_request"
)

var allowedSpansAggregateRequestTypeEnumValues = []SpansAggregateRequestType{
	SPANSAGGREGATEREQUESTTYPE_AGGREGATE_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansAggregateRequestType) GetAllowedValues() []SpansAggregateRequestType {
	return allowedSpansAggregateRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansAggregateRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansAggregateRequestType(value)
	return nil
}

// NewSpansAggregateRequestTypeFromValue returns a pointer to a valid SpansAggregateRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansAggregateRequestTypeFromValue(v string) (*SpansAggregateRequestType, error) {
	ev := SpansAggregateRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansAggregateRequestType: valid values are %v", v, allowedSpansAggregateRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansAggregateRequestType) IsValid() bool {
	for _, existing := range allowedSpansAggregateRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansAggregateRequestType value.
func (v SpansAggregateRequestType) Ptr() *SpansAggregateRequestType {
	return &v
}
