// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansListRequestType The type of resource. The value should always be search_request.
type SpansListRequestType string

// List of SpansListRequestType.
const (
	SPANSLISTREQUESTTYPE_SEARCH_REQUEST SpansListRequestType = "search_request"
)

var allowedSpansListRequestTypeEnumValues = []SpansListRequestType{
	SPANSLISTREQUESTTYPE_SEARCH_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansListRequestType) GetAllowedValues() []SpansListRequestType {
	return allowedSpansListRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansListRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansListRequestType(value)
	return nil
}

// NewSpansListRequestTypeFromValue returns a pointer to a valid SpansListRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansListRequestTypeFromValue(v string) (*SpansListRequestType, error) {
	ev := SpansListRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansListRequestType: valid values are %v", v, allowedSpansListRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansListRequestType) IsValid() bool {
	for _, existing := range allowedSpansListRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansListRequestType value.
func (v SpansListRequestType) Ptr() *SpansListRequestType {
	return &v
}
