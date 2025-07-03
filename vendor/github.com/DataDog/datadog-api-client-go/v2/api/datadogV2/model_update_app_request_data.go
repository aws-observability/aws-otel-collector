// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppRequestData The data object containing the new app definition. Any fields not included in the request remain unchanged.
type UpdateAppRequestData struct {
	// App definition attributes to be updated, such as name, description, and components.
	Attributes *UpdateAppRequestDataAttributes `json:"attributes,omitempty"`
	// The ID of the app to update. The app ID must match the ID in the URL path.
	Id *uuid.UUID `json:"id,omitempty"`
	// The app definition type.
	Type AppDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateAppRequestData instantiates a new UpdateAppRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateAppRequestData(typeVar AppDefinitionType) *UpdateAppRequestData {
	this := UpdateAppRequestData{}
	this.Type = typeVar
	return &this
}

// NewUpdateAppRequestDataWithDefaults instantiates a new UpdateAppRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateAppRequestDataWithDefaults() *UpdateAppRequestData {
	this := UpdateAppRequestData{}
	var typeVar AppDefinitionType = APPDEFINITIONTYPE_APPDEFINITIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateAppRequestData) GetAttributes() UpdateAppRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret UpdateAppRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppRequestData) GetAttributesOk() (*UpdateAppRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateAppRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UpdateAppRequestDataAttributes and assigns it to the Attributes field.
func (o *UpdateAppRequestData) SetAttributes(v UpdateAppRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateAppRequestData) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppRequestData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateAppRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *UpdateAppRequestData) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *UpdateAppRequestData) GetType() AppDefinitionType {
	if o == nil {
		var ret AppDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateAppRequestData) GetTypeOk() (*AppDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UpdateAppRequestData) SetType(v AppDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateAppRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateAppRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UpdateAppRequestDataAttributes `json:"attributes,omitempty"`
		Id         *uuid.UUID                      `json:"id,omitempty"`
		Type       *AppDefinitionType              `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
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
