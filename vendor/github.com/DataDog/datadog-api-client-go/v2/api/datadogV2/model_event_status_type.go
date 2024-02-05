// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventStatusType If an alert event is enabled, its status is one of the following:
// `failure`, `error`, `warning`, `info`, `success`, `user_update`,
// `recommendation`, or `snapshot`.
type EventStatusType string

// List of EventStatusType.
const (
	EVENTSTATUSTYPE_FAILURE        EventStatusType = "failure"
	EVENTSTATUSTYPE_ERROR          EventStatusType = "error"
	EVENTSTATUSTYPE_WARNING        EventStatusType = "warning"
	EVENTSTATUSTYPE_INFO           EventStatusType = "info"
	EVENTSTATUSTYPE_SUCCESS        EventStatusType = "success"
	EVENTSTATUSTYPE_USER_UPDATE    EventStatusType = "user_update"
	EVENTSTATUSTYPE_RECOMMENDATION EventStatusType = "recommendation"
	EVENTSTATUSTYPE_SNAPSHOT       EventStatusType = "snapshot"
)

var allowedEventStatusTypeEnumValues = []EventStatusType{
	EVENTSTATUSTYPE_FAILURE,
	EVENTSTATUSTYPE_ERROR,
	EVENTSTATUSTYPE_WARNING,
	EVENTSTATUSTYPE_INFO,
	EVENTSTATUSTYPE_SUCCESS,
	EVENTSTATUSTYPE_USER_UPDATE,
	EVENTSTATUSTYPE_RECOMMENDATION,
	EVENTSTATUSTYPE_SNAPSHOT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventStatusType) GetAllowedValues() []EventStatusType {
	return allowedEventStatusTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventStatusType(value)
	return nil
}

// NewEventStatusTypeFromValue returns a pointer to a valid EventStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventStatusTypeFromValue(v string) (*EventStatusType, error) {
	ev := EventStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventStatusType: valid values are %v", v, allowedEventStatusTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventStatusType) IsValid() bool {
	for _, existing := range allowedEventStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventStatusType value.
func (v EventStatusType) Ptr() *EventStatusType {
	return &v
}
