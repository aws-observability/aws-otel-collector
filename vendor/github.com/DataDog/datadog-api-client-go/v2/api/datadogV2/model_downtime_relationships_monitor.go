// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeRelationshipsMonitor The monitor identified by the downtime.
type DowntimeRelationshipsMonitor struct {
	// Data for the monitor.
	Data NullableDowntimeRelationshipsMonitorData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeRelationshipsMonitor instantiates a new DowntimeRelationshipsMonitor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeRelationshipsMonitor() *DowntimeRelationshipsMonitor {
	this := DowntimeRelationshipsMonitor{}
	return &this
}

// NewDowntimeRelationshipsMonitorWithDefaults instantiates a new DowntimeRelationshipsMonitor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeRelationshipsMonitorWithDefaults() *DowntimeRelationshipsMonitor {
	this := DowntimeRelationshipsMonitor{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntimeRelationshipsMonitor) GetData() DowntimeRelationshipsMonitorData {
	if o == nil || o.Data.Get() == nil {
		var ret DowntimeRelationshipsMonitorData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DowntimeRelationshipsMonitor) GetDataOk() (*DowntimeRelationshipsMonitorData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *DowntimeRelationshipsMonitor) HasData() bool {
	return o != nil && o.Data.IsSet()
}

// SetData gets a reference to the given NullableDowntimeRelationshipsMonitorData and assigns it to the Data field.
func (o *DowntimeRelationshipsMonitor) SetData(v DowntimeRelationshipsMonitorData) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil.
func (o *DowntimeRelationshipsMonitor) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil.
func (o *DowntimeRelationshipsMonitor) UnsetData() {
	o.Data.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeRelationshipsMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeRelationshipsMonitor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableDowntimeRelationshipsMonitorData `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
