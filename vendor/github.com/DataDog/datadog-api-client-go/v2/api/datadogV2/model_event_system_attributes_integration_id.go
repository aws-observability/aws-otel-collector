// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventSystemAttributesIntegrationId Integration ID sourced from integration manifests.
type EventSystemAttributesIntegrationId string

// List of EventSystemAttributesIntegrationId.
const (
	EVENTSYSTEMATTRIBUTESINTEGRATIONID_CUSTOM_EVENTS EventSystemAttributesIntegrationId = "custom-events"
)

var allowedEventSystemAttributesIntegrationIdEnumValues = []EventSystemAttributesIntegrationId{
	EVENTSYSTEMATTRIBUTESINTEGRATIONID_CUSTOM_EVENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventSystemAttributesIntegrationId) GetAllowedValues() []EventSystemAttributesIntegrationId {
	return allowedEventSystemAttributesIntegrationIdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventSystemAttributesIntegrationId) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventSystemAttributesIntegrationId(value)
	return nil
}

// NewEventSystemAttributesIntegrationIdFromValue returns a pointer to a valid EventSystemAttributesIntegrationId
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventSystemAttributesIntegrationIdFromValue(v string) (*EventSystemAttributesIntegrationId, error) {
	ev := EventSystemAttributesIntegrationId(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventSystemAttributesIntegrationId: valid values are %v", v, allowedEventSystemAttributesIntegrationIdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventSystemAttributesIntegrationId) IsValid() bool {
	for _, existing := range allowedEventSystemAttributesIntegrationIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventSystemAttributesIntegrationId value.
func (v EventSystemAttributesIntegrationId) Ptr() *EventSystemAttributesIntegrationId {
	return &v
}
