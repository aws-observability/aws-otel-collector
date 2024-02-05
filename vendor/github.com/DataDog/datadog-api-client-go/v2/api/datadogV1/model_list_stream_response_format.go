// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamResponseFormat Widget response format.
type ListStreamResponseFormat string

// List of ListStreamResponseFormat.
const (
	LISTSTREAMRESPONSEFORMAT_EVENT_LIST ListStreamResponseFormat = "event_list"
)

var allowedListStreamResponseFormatEnumValues = []ListStreamResponseFormat{
	LISTSTREAMRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListStreamResponseFormat) GetAllowedValues() []ListStreamResponseFormat {
	return allowedListStreamResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListStreamResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamResponseFormat(value)
	return nil
}

// NewListStreamResponseFormatFromValue returns a pointer to a valid ListStreamResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListStreamResponseFormatFromValue(v string) (*ListStreamResponseFormat, error) {
	ev := ListStreamResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListStreamResponseFormat: valid values are %v", v, allowedListStreamResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListStreamResponseFormat) IsValid() bool {
	for _, existing := range allowedListStreamResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamResponseFormat value.
func (v ListStreamResponseFormat) Ptr() *ListStreamResponseFormat {
	return &v
}
