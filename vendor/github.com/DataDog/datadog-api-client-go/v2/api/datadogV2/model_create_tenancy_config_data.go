// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTenancyConfigData The data object for creating a new OCI tenancy integration configuration, including the tenancy ID, type, and configuration attributes.
type CreateTenancyConfigData struct {
	// Attributes for creating a new OCI tenancy integration configuration, including credentials, region settings, and collection options.
	Attributes *CreateTenancyConfigDataAttributes `json:"attributes,omitempty"`
	// The OCID of the OCI tenancy to configure.
	Id string `json:"id"`
	// OCI tenancy resource type.
	Type UpdateTenancyConfigDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateTenancyConfigData instantiates a new CreateTenancyConfigData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTenancyConfigData(id string, typeVar UpdateTenancyConfigDataType) *CreateTenancyConfigData {
	this := CreateTenancyConfigData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCreateTenancyConfigDataWithDefaults instantiates a new CreateTenancyConfigData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTenancyConfigDataWithDefaults() *CreateTenancyConfigData {
	this := CreateTenancyConfigData{}
	var typeVar UpdateTenancyConfigDataType = UPDATETENANCYCONFIGDATATYPE_OCI_TENANCY
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateTenancyConfigData) GetAttributes() CreateTenancyConfigDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret CreateTenancyConfigDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigData) GetAttributesOk() (*CreateTenancyConfigDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateTenancyConfigData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CreateTenancyConfigDataAttributes and assigns it to the Attributes field.
func (o *CreateTenancyConfigData) SetAttributes(v CreateTenancyConfigDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *CreateTenancyConfigData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CreateTenancyConfigData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CreateTenancyConfigData) GetType() UpdateTenancyConfigDataType {
	if o == nil {
		var ret UpdateTenancyConfigDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigData) GetTypeOk() (*UpdateTenancyConfigDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateTenancyConfigData) SetType(v UpdateTenancyConfigDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTenancyConfigData) MarshalJSON() ([]byte, error) {
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
func (o *CreateTenancyConfigData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CreateTenancyConfigDataAttributes `json:"attributes,omitempty"`
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
