// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorAlertTriggerAttributes Attributes for a monitor alert trigger.
type MonitorAlertTriggerAttributes struct {
	// The event ID associated with the monitor alert.
	EventId string `json:"event_id"`
	// The timestamp of the event in Unix milliseconds.
	EventTs int64 `json:"event_ts"`
	// The monitor ID that triggered the alert.
	MonitorId int64 `json:"monitor_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorAlertTriggerAttributes instantiates a new MonitorAlertTriggerAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorAlertTriggerAttributes(eventId string, eventTs int64, monitorId int64) *MonitorAlertTriggerAttributes {
	this := MonitorAlertTriggerAttributes{}
	this.EventId = eventId
	this.EventTs = eventTs
	this.MonitorId = monitorId
	return &this
}

// NewMonitorAlertTriggerAttributesWithDefaults instantiates a new MonitorAlertTriggerAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorAlertTriggerAttributesWithDefaults() *MonitorAlertTriggerAttributes {
	this := MonitorAlertTriggerAttributes{}
	return &this
}

// GetEventId returns the EventId field value.
func (o *MonitorAlertTriggerAttributes) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *MonitorAlertTriggerAttributes) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value.
func (o *MonitorAlertTriggerAttributes) SetEventId(v string) {
	o.EventId = v
}

// GetEventTs returns the EventTs field value.
func (o *MonitorAlertTriggerAttributes) GetEventTs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EventTs
}

// GetEventTsOk returns a tuple with the EventTs field value
// and a boolean to check if the value has been set.
func (o *MonitorAlertTriggerAttributes) GetEventTsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTs, true
}

// SetEventTs sets field value.
func (o *MonitorAlertTriggerAttributes) SetEventTs(v int64) {
	o.EventTs = v
}

// GetMonitorId returns the MonitorId field value.
func (o *MonitorAlertTriggerAttributes) GetMonitorId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value
// and a boolean to check if the value has been set.
func (o *MonitorAlertTriggerAttributes) GetMonitorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorId, true
}

// SetMonitorId sets field value.
func (o *MonitorAlertTriggerAttributes) SetMonitorId(v int64) {
	o.MonitorId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorAlertTriggerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["event_id"] = o.EventId
	toSerialize["event_ts"] = o.EventTs
	toSerialize["monitor_id"] = o.MonitorId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorAlertTriggerAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventId   *string `json:"event_id"`
		EventTs   *int64  `json:"event_ts"`
		MonitorId *int64  `json:"monitor_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EventId == nil {
		return fmt.Errorf("required field event_id missing")
	}
	if all.EventTs == nil {
		return fmt.Errorf("required field event_ts missing")
	}
	if all.MonitorId == nil {
		return fmt.Errorf("required field monitor_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_id", "event_ts", "monitor_id"})
	} else {
		return err
	}
	o.EventId = *all.EventId
	o.EventTs = *all.EventTs
	o.MonitorId = *all.MonitorId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
