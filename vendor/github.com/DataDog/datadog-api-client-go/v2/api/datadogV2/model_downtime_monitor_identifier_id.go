// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeMonitorIdentifierId Object of the monitor identifier.
type DowntimeMonitorIdentifierId struct {
	// ID of the monitor to prevent notifications.
	MonitorId int64 `json:"monitor_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeMonitorIdentifierId instantiates a new DowntimeMonitorIdentifierId object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeMonitorIdentifierId(monitorId int64) *DowntimeMonitorIdentifierId {
	this := DowntimeMonitorIdentifierId{}
	this.MonitorId = monitorId
	return &this
}

// NewDowntimeMonitorIdentifierIdWithDefaults instantiates a new DowntimeMonitorIdentifierId object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeMonitorIdentifierIdWithDefaults() *DowntimeMonitorIdentifierId {
	this := DowntimeMonitorIdentifierId{}
	return &this
}

// GetMonitorId returns the MonitorId field value.
func (o *DowntimeMonitorIdentifierId) GetMonitorId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value
// and a boolean to check if the value has been set.
func (o *DowntimeMonitorIdentifierId) GetMonitorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorId, true
}

// SetMonitorId sets field value.
func (o *DowntimeMonitorIdentifierId) SetMonitorId(v int64) {
	o.MonitorId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeMonitorIdentifierId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["monitor_id"] = o.MonitorId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeMonitorIdentifierId) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MonitorId *int64 `json:"monitor_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MonitorId == nil {
		return fmt.Errorf("required field monitor_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"monitor_id"})
	} else {
		return err
	}
	o.MonitorId = *all.MonitorId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
