// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventPayloadIntegrationId Integration ID sourced from integration manifests.
type EventPayloadIntegrationId string

// List of EventPayloadIntegrationId.
const (
	EVENTPAYLOADINTEGRATIONID_CUSTOM_EVENTS EventPayloadIntegrationId = "custom-events"
)

var allowedEventPayloadIntegrationIdEnumValues = []EventPayloadIntegrationId{
	EVENTPAYLOADINTEGRATIONID_CUSTOM_EVENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventPayloadIntegrationId) GetAllowedValues() []EventPayloadIntegrationId {
	return allowedEventPayloadIntegrationIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventPayloadIntegrationId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventPayloadIntegrationId(value)
	return nil
}

// NewEventPayloadIntegrationIdFromValue returns a pointer to a valid EventPayloadIntegrationId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventPayloadIntegrationIdFromValue(v string) (*EventPayloadIntegrationId, error) {
	ev := EventPayloadIntegrationId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventPayloadIntegrationId: valid values are %v", v, allowedEventPayloadIntegrationIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventPayloadIntegrationId) IsValid() bool {
	for _, existing := range allowedEventPayloadIntegrationIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventPayloadIntegrationId value.
func (v EventPayloadIntegrationId) Ptr() *EventPayloadIntegrationId {
	return &v
}
