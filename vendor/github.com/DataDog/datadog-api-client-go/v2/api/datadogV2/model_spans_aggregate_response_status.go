// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateResponseStatus The status of the response.
type SpansAggregateResponseStatus string

// List of SpansAggregateResponseStatus.
const (
	SPANSAGGREGATERESPONSESTATUS_DONE    SpansAggregateResponseStatus = "done"
	SPANSAGGREGATERESPONSESTATUS_TIMEOUT SpansAggregateResponseStatus = "timeout"
)

var allowedSpansAggregateResponseStatusEnumValues = []SpansAggregateResponseStatus{
	SPANSAGGREGATERESPONSESTATUS_DONE,
	SPANSAGGREGATERESPONSESTATUS_TIMEOUT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansAggregateResponseStatus) GetAllowedValues() []SpansAggregateResponseStatus {
	return allowedSpansAggregateResponseStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansAggregateResponseStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansAggregateResponseStatus(value)
	return nil
}

// NewSpansAggregateResponseStatusFromValue returns a pointer to a valid SpansAggregateResponseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansAggregateResponseStatusFromValue(v string) (*SpansAggregateResponseStatus, error) {
	ev := SpansAggregateResponseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansAggregateResponseStatus: valid values are %v", v, allowedSpansAggregateResponseStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansAggregateResponseStatus) IsValid() bool {
	for _, existing := range allowedSpansAggregateResponseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansAggregateResponseStatus value.
func (v SpansAggregateResponseStatus) Ptr() *SpansAggregateResponseStatus {
	return &v
}
