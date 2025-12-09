// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsDataSource A data source that is powered by the Events Platform.
type EventsDataSource string

// List of EventsDataSource.
const (
	EVENTSDATASOURCE_LOGS EventsDataSource = "logs"
	EVENTSDATASOURCE_RUM  EventsDataSource = "rum"
	EVENTSDATASOURCE_DORA EventsDataSource = "dora"
)

var allowedEventsDataSourceEnumValues = []EventsDataSource{
	EVENTSDATASOURCE_LOGS,
	EVENTSDATASOURCE_RUM,
	EVENTSDATASOURCE_DORA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventsDataSource) GetAllowedValues() []EventsDataSource {
	return allowedEventsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventsDataSource(value)
	return nil
}

// NewEventsDataSourceFromValue returns a pointer to a valid EventsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventsDataSourceFromValue(v string) (*EventsDataSource, error) {
	ev := EventsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventsDataSource: valid values are %v", v, allowedEventsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventsDataSource) IsValid() bool {
	for _, existing := range allowedEventsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventsDataSource value.
func (v EventsDataSource) Ptr() *EventsDataSource {
	return &v
}
