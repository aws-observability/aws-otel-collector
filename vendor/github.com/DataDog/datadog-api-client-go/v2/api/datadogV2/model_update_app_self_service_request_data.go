// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppSelfServiceRequestData Data for updating an app's self-service status.
type UpdateAppSelfServiceRequestData struct {
	// Attributes for updating an app's self-service status.
	Attributes *UpdateAppSelfServiceRequestDataAttributes `json:"attributes,omitempty"`
	// The self-service resource type.
	Type *AppSelfServiceType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateAppSelfServiceRequestData instantiates a new UpdateAppSelfServiceRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateAppSelfServiceRequestData() *UpdateAppSelfServiceRequestData {
	this := UpdateAppSelfServiceRequestData{}
	var typeVar AppSelfServiceType = APPSELFSERVICETYPE_SELFSERVICE
	this.Type = &typeVar
	return &this
}

// NewUpdateAppSelfServiceRequestDataWithDefaults instantiates a new UpdateAppSelfServiceRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateAppSelfServiceRequestDataWithDefaults() *UpdateAppSelfServiceRequestData {
	this := UpdateAppSelfServiceRequestData{}
	var typeVar AppSelfServiceType = APPSELFSERVICETYPE_SELFSERVICE
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateAppSelfServiceRequestData) GetAttributes() UpdateAppSelfServiceRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret UpdateAppSelfServiceRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppSelfServiceRequestData) GetAttributesOk() (*UpdateAppSelfServiceRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateAppSelfServiceRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UpdateAppSelfServiceRequestDataAttributes and assigns it to the Attributes field.
func (o *UpdateAppSelfServiceRequestData) SetAttributes(v UpdateAppSelfServiceRequestDataAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateAppSelfServiceRequestData) GetType() AppSelfServiceType {
	if o == nil || o.Type == nil {
		var ret AppSelfServiceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppSelfServiceRequestData) GetTypeOk() (*AppSelfServiceType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateAppSelfServiceRequestData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given AppSelfServiceType and assigns it to the Type field.
func (o *UpdateAppSelfServiceRequestData) SetType(v AppSelfServiceType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateAppSelfServiceRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
func (o *UpdateAppSelfServiceRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UpdateAppSelfServiceRequestDataAttributes `json:"attributes,omitempty"`
		Type       *AppSelfServiceType                        `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
