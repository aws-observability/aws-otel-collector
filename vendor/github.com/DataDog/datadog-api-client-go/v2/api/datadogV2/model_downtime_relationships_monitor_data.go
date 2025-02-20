// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeRelationshipsMonitorData Data for the monitor.
type DowntimeRelationshipsMonitorData struct {
	// Monitor ID of the downtime.
	Id *string `json:"id,omitempty"`
	// Monitor resource type.
	Type *DowntimeIncludedMonitorType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeRelationshipsMonitorData instantiates a new DowntimeRelationshipsMonitorData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeRelationshipsMonitorData() *DowntimeRelationshipsMonitorData {
	this := DowntimeRelationshipsMonitorData{}
	var typeVar DowntimeIncludedMonitorType = DOWNTIMEINCLUDEDMONITORTYPE_MONITORS
	this.Type = &typeVar
	return &this
}

// NewDowntimeRelationshipsMonitorDataWithDefaults instantiates a new DowntimeRelationshipsMonitorData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeRelationshipsMonitorDataWithDefaults() *DowntimeRelationshipsMonitorData {
	this := DowntimeRelationshipsMonitorData{}
	var typeVar DowntimeIncludedMonitorType = DOWNTIMEINCLUDEDMONITORTYPE_MONITORS
	this.Type = &typeVar
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DowntimeRelationshipsMonitorData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeRelationshipsMonitorData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DowntimeRelationshipsMonitorData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DowntimeRelationshipsMonitorData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DowntimeRelationshipsMonitorData) GetType() DowntimeIncludedMonitorType {
	if o == nil || o.Type == nil {
		var ret DowntimeIncludedMonitorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeRelationshipsMonitorData) GetTypeOk() (*DowntimeIncludedMonitorType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DowntimeRelationshipsMonitorData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given DowntimeIncludedMonitorType and assigns it to the Type field.
func (o *DowntimeRelationshipsMonitorData) SetType(v DowntimeIncludedMonitorType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeRelationshipsMonitorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeRelationshipsMonitorData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                      `json:"id,omitempty"`
		Type *DowntimeIncludedMonitorType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableDowntimeRelationshipsMonitorData handles when a null is used for DowntimeRelationshipsMonitorData.
type NullableDowntimeRelationshipsMonitorData struct {
	value *DowntimeRelationshipsMonitorData
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeRelationshipsMonitorData) Get() *DowntimeRelationshipsMonitorData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeRelationshipsMonitorData) Set(val *DowntimeRelationshipsMonitorData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeRelationshipsMonitorData) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDowntimeRelationshipsMonitorData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeRelationshipsMonitorData initializes the struct as if Set has been called.
func NewNullableDowntimeRelationshipsMonitorData(val *DowntimeRelationshipsMonitorData) *NullableDowntimeRelationshipsMonitorData {
	return &NullableDowntimeRelationshipsMonitorData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeRelationshipsMonitorData) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeRelationshipsMonitorData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
