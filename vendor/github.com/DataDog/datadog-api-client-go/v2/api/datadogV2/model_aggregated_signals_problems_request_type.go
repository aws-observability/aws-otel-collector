// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedSignalsProblemsRequestType The JSON:API type for aggregated signals and problems requests.
type AggregatedSignalsProblemsRequestType string

// List of AggregatedSignalsProblemsRequestType.
const (
	AGGREGATEDSIGNALSPROBLEMSREQUESTTYPE_AGGREGATED_SIGNALS_PROBLEMS AggregatedSignalsProblemsRequestType = "aggregated_signals_problems"
)

var allowedAggregatedSignalsProblemsRequestTypeEnumValues = []AggregatedSignalsProblemsRequestType{
	AGGREGATEDSIGNALSPROBLEMSREQUESTTYPE_AGGREGATED_SIGNALS_PROBLEMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AggregatedSignalsProblemsRequestType) GetAllowedValues() []AggregatedSignalsProblemsRequestType {
	return allowedAggregatedSignalsProblemsRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregatedSignalsProblemsRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregatedSignalsProblemsRequestType(value)
	return nil
}

// NewAggregatedSignalsProblemsRequestTypeFromValue returns a pointer to a valid AggregatedSignalsProblemsRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregatedSignalsProblemsRequestTypeFromValue(v string) (*AggregatedSignalsProblemsRequestType, error) {
	ev := AggregatedSignalsProblemsRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregatedSignalsProblemsRequestType: valid values are %v", v, allowedAggregatedSignalsProblemsRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregatedSignalsProblemsRequestType) IsValid() bool {
	for _, existing := range allowedAggregatedSignalsProblemsRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregatedSignalsProblemsRequestType value.
func (v AggregatedSignalsProblemsRequestType) Ptr() *AggregatedSignalsProblemsRequestType {
	return &v
}
