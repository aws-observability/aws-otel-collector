// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateTenancyConfigData The data object for updating an existing OCI tenancy integration configuration, including the tenancy ID, type, and updated attributes.
type UpdateTenancyConfigData struct {
	// Attributes for updating an existing OCI tenancy integration configuration, including optional credentials, region settings, and collection options.
	Attributes *UpdateTenancyConfigDataAttributes `json:"attributes,omitempty"`
	// The OCID of the OCI tenancy to update.
	Id string `json:"id"`
	// OCI tenancy resource type.
	Type UpdateTenancyConfigDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateTenancyConfigData instantiates a new UpdateTenancyConfigData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateTenancyConfigData(id string, typeVar UpdateTenancyConfigDataType) *UpdateTenancyConfigData {
	this := UpdateTenancyConfigData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewUpdateTenancyConfigDataWithDefaults instantiates a new UpdateTenancyConfigData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateTenancyConfigDataWithDefaults() *UpdateTenancyConfigData {
	this := UpdateTenancyConfigData{}
	var typeVar UpdateTenancyConfigDataType = UPDATETENANCYCONFIGDATATYPE_OCI_TENANCY
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateTenancyConfigData) GetAttributes() UpdateTenancyConfigDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret UpdateTenancyConfigDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigData) GetAttributesOk() (*UpdateTenancyConfigDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateTenancyConfigData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UpdateTenancyConfigDataAttributes and assigns it to the Attributes field.
func (o *UpdateTenancyConfigData) SetAttributes(v UpdateTenancyConfigDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *UpdateTenancyConfigData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *UpdateTenancyConfigData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *UpdateTenancyConfigData) GetType() UpdateTenancyConfigDataType {
	if o == nil {
		var ret UpdateTenancyConfigDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateTenancyConfigData) GetTypeOk() (*UpdateTenancyConfigDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UpdateTenancyConfigData) SetType(v UpdateTenancyConfigDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateTenancyConfigData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateTenancyConfigData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UpdateTenancyConfigDataAttributes `json:"attributes,omitempty"`
		Id         *string                            `json:"id"`
		Type       *UpdateTenancyConfigDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
	o.Id = *all.Id
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
