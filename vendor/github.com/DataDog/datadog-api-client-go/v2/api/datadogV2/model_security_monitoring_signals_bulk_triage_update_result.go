// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalsBulkTriageUpdateResult The result payload of a bulk signal triage update.
type SecurityMonitoringSignalsBulkTriageUpdateResult struct {
	// The number of signals updated.
	Count int64 `json:"count"`
	// The list of updated signals.
	Events []SecurityMonitoringSignalsBulkTriageEvent `json:"events"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalsBulkTriageUpdateResult instantiates a new SecurityMonitoringSignalsBulkTriageUpdateResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalsBulkTriageUpdateResult(count int64, events []SecurityMonitoringSignalsBulkTriageEvent) *SecurityMonitoringSignalsBulkTriageUpdateResult {
	this := SecurityMonitoringSignalsBulkTriageUpdateResult{}
	this.Count = count
	this.Events = events
	return &this
}

// NewSecurityMonitoringSignalsBulkTriageUpdateResultWithDefaults instantiates a new SecurityMonitoringSignalsBulkTriageUpdateResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalsBulkTriageUpdateResultWithDefaults() *SecurityMonitoringSignalsBulkTriageUpdateResult {
	this := SecurityMonitoringSignalsBulkTriageUpdateResult{}
	return &this
}

// GetCount returns the Count field value.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResult) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResult) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResult) SetCount(v int64) {
	o.Count = v
}

// GetEvents returns the Events field value.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResult) GetEvents() []SecurityMonitoringSignalsBulkTriageEvent {
	if o == nil {
		var ret []SecurityMonitoringSignalsBulkTriageEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResult) GetEventsOk() (*[]SecurityMonitoringSignalsBulkTriageEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResult) SetEvents(v []SecurityMonitoringSignalsBulkTriageEvent) {
	o.Events = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalsBulkTriageUpdateResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	toSerialize["events"] = o.Events

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count  *int64                                      `json:"count"`
		Events *[]SecurityMonitoringSignalsBulkTriageEvent `json:"events"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.Events == nil {
		return fmt.Errorf("required field events missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "events"})
	} else {
		return err
	}
	o.Count = *all.Count
	o.Events = *all.Events

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
