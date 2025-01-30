// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3MetadataAdditionalOwnersItems The definition of Entity V3 Metadata Additional Owners Items object.
type EntityV3MetadataAdditionalOwnersItems struct {
	// Team name.
	Name string `json:"name"`
	// Team type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityV3MetadataAdditionalOwnersItems instantiates a new EntityV3MetadataAdditionalOwnersItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3MetadataAdditionalOwnersItems(name string) *EntityV3MetadataAdditionalOwnersItems {
	this := EntityV3MetadataAdditionalOwnersItems{}
	this.Name = name
	return &this
}

// NewEntityV3MetadataAdditionalOwnersItemsWithDefaults instantiates a new EntityV3MetadataAdditionalOwnersItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3MetadataAdditionalOwnersItemsWithDefaults() *EntityV3MetadataAdditionalOwnersItems {
	this := EntityV3MetadataAdditionalOwnersItems{}
	return &this
}

// GetName returns the Name field value.
func (o *EntityV3MetadataAdditionalOwnersItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntityV3MetadataAdditionalOwnersItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EntityV3MetadataAdditionalOwnersItems) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntityV3MetadataAdditionalOwnersItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3MetadataAdditionalOwnersItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntityV3MetadataAdditionalOwnersItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntityV3MetadataAdditionalOwnersItems) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3MetadataAdditionalOwnersItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3MetadataAdditionalOwnersItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name"`
		Type *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "type"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
