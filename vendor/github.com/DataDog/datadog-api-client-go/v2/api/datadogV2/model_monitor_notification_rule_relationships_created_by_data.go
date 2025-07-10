// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleRelationshipsCreatedByData Data for the user who created the monitor notification rule.
type MonitorNotificationRuleRelationshipsCreatedByData struct {
	// User ID of the monitor notification rule creator.
	Id *string `json:"id,omitempty"`
	// Users resource type.
	Type *UsersType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleRelationshipsCreatedByData instantiates a new MonitorNotificationRuleRelationshipsCreatedByData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleRelationshipsCreatedByData() *MonitorNotificationRuleRelationshipsCreatedByData {
	this := MonitorNotificationRuleRelationshipsCreatedByData{}
	var typeVar UsersType = USERSTYPE_USERS
	this.Type = &typeVar
	return &this
}

// NewMonitorNotificationRuleRelationshipsCreatedByDataWithDefaults instantiates a new MonitorNotificationRuleRelationshipsCreatedByData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleRelationshipsCreatedByDataWithDefaults() *MonitorNotificationRuleRelationshipsCreatedByData {
	this := MonitorNotificationRuleRelationshipsCreatedByData{}
	var typeVar UsersType = USERSTYPE_USERS
	this.Type = &typeVar
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MonitorNotificationRuleRelationshipsCreatedByData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleRelationshipsCreatedByData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MonitorNotificationRuleRelationshipsCreatedByData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MonitorNotificationRuleRelationshipsCreatedByData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitorNotificationRuleRelationshipsCreatedByData) GetType() UsersType {
	if o == nil || o.Type == nil {
		var ret UsersType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleRelationshipsCreatedByData) GetTypeOk() (*UsersType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitorNotificationRuleRelationshipsCreatedByData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given UsersType and assigns it to the Type field.
func (o *MonitorNotificationRuleRelationshipsCreatedByData) SetType(v UsersType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleRelationshipsCreatedByData) MarshalJSON() ([]byte, error) {
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
func (o *MonitorNotificationRuleRelationshipsCreatedByData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string    `json:"id,omitempty"`
		Type *UsersType `json:"type,omitempty"`
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

// NullableMonitorNotificationRuleRelationshipsCreatedByData handles when a null is used for MonitorNotificationRuleRelationshipsCreatedByData.
type NullableMonitorNotificationRuleRelationshipsCreatedByData struct {
	value *MonitorNotificationRuleRelationshipsCreatedByData
	isSet bool
}

// Get returns the associated value.
func (v NullableMonitorNotificationRuleRelationshipsCreatedByData) Get() *MonitorNotificationRuleRelationshipsCreatedByData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableMonitorNotificationRuleRelationshipsCreatedByData) Set(val *MonitorNotificationRuleRelationshipsCreatedByData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableMonitorNotificationRuleRelationshipsCreatedByData) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableMonitorNotificationRuleRelationshipsCreatedByData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableMonitorNotificationRuleRelationshipsCreatedByData initializes the struct as if Set has been called.
func NewNullableMonitorNotificationRuleRelationshipsCreatedByData(val *MonitorNotificationRuleRelationshipsCreatedByData) *NullableMonitorNotificationRuleRelationshipsCreatedByData {
	return &NullableMonitorNotificationRuleRelationshipsCreatedByData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableMonitorNotificationRuleRelationshipsCreatedByData) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableMonitorNotificationRuleRelationshipsCreatedByData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
