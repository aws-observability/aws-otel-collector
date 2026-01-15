// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateNotificationChannelData Data for creating an on-call notification channel
type CreateNotificationChannelData struct {
	// Attributes for creating an on-call notification channel.
	Attributes *CreateNotificationChannelAttributes `json:"attributes,omitempty"`
	// Indicates that the resource is of type 'notification_channels'.
	Type NotificationChannelType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateNotificationChannelData instantiates a new CreateNotificationChannelData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateNotificationChannelData(typeVar NotificationChannelType) *CreateNotificationChannelData {
	this := CreateNotificationChannelData{}
	this.Type = typeVar
	return &this
}

// NewCreateNotificationChannelDataWithDefaults instantiates a new CreateNotificationChannelData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateNotificationChannelDataWithDefaults() *CreateNotificationChannelData {
	this := CreateNotificationChannelData{}
	var typeVar NotificationChannelType = NOTIFICATIONCHANNELTYPE_NOTIFICATION_CHANNELS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateNotificationChannelData) GetAttributes() CreateNotificationChannelAttributes {
	if o == nil || o.Attributes == nil {
		var ret CreateNotificationChannelAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationChannelData) GetAttributesOk() (*CreateNotificationChannelAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateNotificationChannelData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CreateNotificationChannelAttributes and assigns it to the Attributes field.
func (o *CreateNotificationChannelData) SetAttributes(v CreateNotificationChannelAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *CreateNotificationChannelData) GetType() NotificationChannelType {
	if o == nil {
		var ret NotificationChannelType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationChannelData) GetTypeOk() (*NotificationChannelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateNotificationChannelData) SetType(v NotificationChannelType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateNotificationChannelData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateNotificationChannelData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CreateNotificationChannelAttributes `json:"attributes,omitempty"`
		Type       *NotificationChannelType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
