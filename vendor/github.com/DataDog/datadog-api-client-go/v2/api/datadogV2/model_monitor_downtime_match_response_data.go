// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorDowntimeMatchResponseData A downtime match.
type MonitorDowntimeMatchResponseData struct {
	// Downtime match details.
	Attributes *MonitorDowntimeMatchResponseAttributes `json:"attributes,omitempty"`
	// The downtime ID.
	Id datadog.NullableString `json:"id,omitempty"`
	// Monitor Downtime Match resource type.
	Type *MonitorDowntimeMatchResourceType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorDowntimeMatchResponseData instantiates a new MonitorDowntimeMatchResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorDowntimeMatchResponseData() *MonitorDowntimeMatchResponseData {
	this := MonitorDowntimeMatchResponseData{}
	var typeVar MonitorDowntimeMatchResourceType = MONITORDOWNTIMEMATCHRESOURCETYPE_DOWNTIME_MATCH
	this.Type = &typeVar
	return &this
}

// NewMonitorDowntimeMatchResponseDataWithDefaults instantiates a new MonitorDowntimeMatchResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorDowntimeMatchResponseDataWithDefaults() *MonitorDowntimeMatchResponseData {
	this := MonitorDowntimeMatchResponseData{}
	var typeVar MonitorDowntimeMatchResourceType = MONITORDOWNTIMEMATCHRESOURCETYPE_DOWNTIME_MATCH
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MonitorDowntimeMatchResponseData) GetAttributes() MonitorDowntimeMatchResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret MonitorDowntimeMatchResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDowntimeMatchResponseData) GetAttributesOk() (*MonitorDowntimeMatchResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given MonitorDowntimeMatchResponseAttributes and assigns it to the Attributes field.
func (o *MonitorDowntimeMatchResponseData) SetAttributes(v MonitorDowntimeMatchResponseAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDowntimeMatchResponseData) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDowntimeMatchResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponseData) HasId() bool {
	return o != nil && o.Id.IsSet()
}

// SetId gets a reference to the given datadog.NullableString and assigns it to the Id field.
func (o *MonitorDowntimeMatchResponseData) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil.
func (o *MonitorDowntimeMatchResponseData) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil.
func (o *MonitorDowntimeMatchResponseData) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitorDowntimeMatchResponseData) GetType() MonitorDowntimeMatchResourceType {
	if o == nil || o.Type == nil {
		var ret MonitorDowntimeMatchResourceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDowntimeMatchResponseData) GetTypeOk() (*MonitorDowntimeMatchResourceType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponseData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given MonitorDowntimeMatchResourceType and assigns it to the Type field.
func (o *MonitorDowntimeMatchResponseData) SetType(v MonitorDowntimeMatchResourceType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorDowntimeMatchResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
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
func (o *MonitorDowntimeMatchResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *MonitorDowntimeMatchResponseAttributes `json:"attributes,omitempty"`
		Id         datadog.NullableString                  `json:"id,omitempty"`
		Type       *MonitorDowntimeMatchResourceType       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
