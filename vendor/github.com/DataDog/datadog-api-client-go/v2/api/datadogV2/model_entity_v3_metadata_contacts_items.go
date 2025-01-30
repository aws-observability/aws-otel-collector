// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3MetadataContactsItems The definition of Entity V3 Metadata Contacts Items object.
type EntityV3MetadataContactsItems struct {
	// Contact value.
	Contact string `json:"contact"`
	// Contact name.
	Name *string `json:"name,omitempty"`
	// Contact type.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3MetadataContactsItems instantiates a new EntityV3MetadataContactsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3MetadataContactsItems(contact string, typeVar string) *EntityV3MetadataContactsItems {
	this := EntityV3MetadataContactsItems{}
	this.Contact = contact
	this.Type = typeVar
	return &this
}

// NewEntityV3MetadataContactsItemsWithDefaults instantiates a new EntityV3MetadataContactsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3MetadataContactsItemsWithDefaults() *EntityV3MetadataContactsItems {
	this := EntityV3MetadataContactsItems{}
	return &this
}

// GetContact returns the Contact field value.
func (o *EntityV3MetadataContactsItems) GetContact() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *EntityV3MetadataContactsItems) GetContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value.
func (o *EntityV3MetadataContactsItems) SetContact(v string) {
	o.Contact = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntityV3MetadataContactsItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3MetadataContactsItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntityV3MetadataContactsItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntityV3MetadataContactsItems) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value.
func (o *EntityV3MetadataContactsItems) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntityV3MetadataContactsItems) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EntityV3MetadataContactsItems) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3MetadataContactsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["contact"] = o.Contact
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3MetadataContactsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Contact *string `json:"contact"`
		Name    *string `json:"name,omitempty"`
		Type    *string `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Contact == nil {
		return fmt.Errorf("required field contact missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	o.Contact = *all.Contact
	o.Name = all.Name
	o.Type = *all.Type

	return nil
}
