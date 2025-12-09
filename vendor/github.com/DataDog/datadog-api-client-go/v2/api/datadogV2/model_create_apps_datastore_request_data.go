// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreRequestData Data wrapper containing the configuration needed to create a new datastore.
type CreateAppsDatastoreRequestData struct {
	// Configuration and metadata to create a new datastore.
	Attributes *CreateAppsDatastoreRequestDataAttributes `json:"attributes,omitempty"`
	// Optional ID for the new datastore. If not provided, one will be generated automatically.
	Id *string `json:"id,omitempty"`
	// The resource type for datastores.
	Type DatastoreDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateAppsDatastoreRequestData instantiates a new CreateAppsDatastoreRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateAppsDatastoreRequestData(typeVar DatastoreDataType) *CreateAppsDatastoreRequestData {
	this := CreateAppsDatastoreRequestData{}
	this.Type = typeVar
	return &this
}

// NewCreateAppsDatastoreRequestDataWithDefaults instantiates a new CreateAppsDatastoreRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateAppsDatastoreRequestDataWithDefaults() *CreateAppsDatastoreRequestData {
	this := CreateAppsDatastoreRequestData{}
	var typeVar DatastoreDataType = DATASTOREDATATYPE_DATASTORES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateAppsDatastoreRequestData) GetAttributes() CreateAppsDatastoreRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret CreateAppsDatastoreRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreRequestData) GetAttributesOk() (*CreateAppsDatastoreRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateAppsDatastoreRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CreateAppsDatastoreRequestDataAttributes and assigns it to the Attributes field.
func (o *CreateAppsDatastoreRequestData) SetAttributes(v CreateAppsDatastoreRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateAppsDatastoreRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateAppsDatastoreRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateAppsDatastoreRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *CreateAppsDatastoreRequestData) GetType() DatastoreDataType {
	if o == nil {
		var ret DatastoreDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreRequestData) GetTypeOk() (*DatastoreDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateAppsDatastoreRequestData) SetType(v DatastoreDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateAppsDatastoreRequestData) MarshalJSON() ([]byte, error) {
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
func (o *CreateAppsDatastoreRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CreateAppsDatastoreRequestDataAttributes `json:"attributes,omitempty"`
		Id         *string                                   `json:"id,omitempty"`
		Type       *DatastoreDataType                        `json:"type"`
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
