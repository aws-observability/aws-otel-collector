// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SearchSLOTimeframe The SLO time window options.
type SearchSLOTimeframe string

// List of SearchSLOTimeframe.
const (
	SEARCHSLOTIMEFRAME_SEVEN_DAYS  SearchSLOTimeframe = "7d"
	SEARCHSLOTIMEFRAME_THIRTY_DAYS SearchSLOTimeframe = "30d"
	SEARCHSLOTIMEFRAME_NINETY_DAYS SearchSLOTimeframe = "90d"
)

var allowedSearchSLOTimeframeEnumValues = []SearchSLOTimeframe{
	SEARCHSLOTIMEFRAME_SEVEN_DAYS,
	SEARCHSLOTIMEFRAME_THIRTY_DAYS,
	SEARCHSLOTIMEFRAME_NINETY_DAYS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SearchSLOTimeframe) GetAllowedValues() []SearchSLOTimeframe {
	return allowedSearchSLOTimeframeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SearchSLOTimeframe) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SearchSLOTimeframe(value)
	return nil
}

// NewSearchSLOTimeframeFromValue returns a pointer to a valid SearchSLOTimeframe
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSearchSLOTimeframeFromValue(v string) (*SearchSLOTimeframe, error) {
	ev := SearchSLOTimeframe(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SearchSLOTimeframe: valid values are %v", v, allowedSearchSLOTimeframeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SearchSLOTimeframe) IsValid() bool {
	for _, existing := range allowedSearchSLOTimeframeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SearchSLOTimeframe value.
func (v SearchSLOTimeframe) Ptr() *SearchSLOTimeframe {
	return &v
}
