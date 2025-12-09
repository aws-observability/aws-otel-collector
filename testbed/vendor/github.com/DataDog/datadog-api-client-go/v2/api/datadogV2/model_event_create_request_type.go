// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventCreateRequestType Entity type.
type EventCreateRequestType string

// List of EventCreateRequestType.
const (
	EVENTCREATEREQUESTTYPE_EVENT EventCreateRequestType = "event"
)

var allowedEventCreateRequestTypeEnumValues = []EventCreateRequestType{
	EVENTCREATEREQUESTTYPE_EVENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventCreateRequestType) GetAllowedValues() []EventCreateRequestType {
	return allowedEventCreateRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventCreateRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventCreateRequestType(value)
	return nil
}

// NewEventCreateRequestTypeFromValue returns a pointer to a valid EventCreateRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventCreateRequestTypeFromValue(v string) (*EventCreateRequestType, error) {
	ev := EventCreateRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventCreateRequestType: valid values are %v", v, allowedEventCreateRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventCreateRequestType) IsValid() bool {
	for _, existing := range allowedEventCreateRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventCreateRequestType value.
func (v EventCreateRequestType) Ptr() *EventCreateRequestType {
	return &v
}
