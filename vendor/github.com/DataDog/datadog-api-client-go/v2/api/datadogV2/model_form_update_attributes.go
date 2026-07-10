// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormUpdateAttributes The fields to update on a form. At least one field must be provided.
type FormUpdateAttributes struct {
	// The datastore configuration for a form.
	DatastoreConfig *FormDatastoreConfigAttributes `json:"datastore_config,omitempty"`
	// The updated description of the form.
	Description *string `json:"description,omitempty"`
	// The updated name of the form.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormUpdateAttributes instantiates a new FormUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormUpdateAttributes() *FormUpdateAttributes {
	this := FormUpdateAttributes{}
	return &this
}

// NewFormUpdateAttributesWithDefaults instantiates a new FormUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormUpdateAttributesWithDefaults() *FormUpdateAttributes {
	this := FormUpdateAttributes{}
	return &this
}

// GetDatastoreConfig returns the DatastoreConfig field value if set, zero value otherwise.
func (o *FormUpdateAttributes) GetDatastoreConfig() FormDatastoreConfigAttributes {
	if o == nil || o.DatastoreConfig == nil {
		var ret FormDatastoreConfigAttributes
		return ret
	}
	return *o.DatastoreConfig
}

// GetDatastoreConfigOk returns a tuple with the DatastoreConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormUpdateAttributes) GetDatastoreConfigOk() (*FormDatastoreConfigAttributes, bool) {
	if o == nil || o.DatastoreConfig == nil {
		return nil, false
	}
	return o.DatastoreConfig, true
}

// HasDatastoreConfig returns a boolean if a field has been set.
func (o *FormUpdateAttributes) HasDatastoreConfig() bool {
	return o != nil && o.DatastoreConfig != nil
}

// SetDatastoreConfig gets a reference to the given FormDatastoreConfigAttributes and assigns it to the DatastoreConfig field.
func (o *FormUpdateAttributes) SetDatastoreConfig(v FormDatastoreConfigAttributes) {
	o.DatastoreConfig = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FormUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FormUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FormUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FormUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FormUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FormUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DatastoreConfig != nil {
		toSerialize["datastore_config"] = o.DatastoreConfig
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatastoreConfig *FormDatastoreConfigAttributes `json:"datastore_config,omitempty"`
		Description     *string                        `json:"description,omitempty"`
		Name            *string                        `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"datastore_config", "description", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DatastoreConfig != nil && all.DatastoreConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatastoreConfig = all.DatastoreConfig
	o.Description = all.Description
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
