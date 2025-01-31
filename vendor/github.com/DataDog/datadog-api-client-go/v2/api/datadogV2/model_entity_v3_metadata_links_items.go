// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3MetadataLinksItems The definition of Entity V3 Metadata Links Items object.
type EntityV3MetadataLinksItems struct {
	// Link name.
	Name string `json:"name"`
	// Link provider.
	Provider *string `json:"provider,omitempty"`
	// Link type.
	Type string `json:"type"`
	// Link URL.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3MetadataLinksItems instantiates a new EntityV3MetadataLinksItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3MetadataLinksItems(name string, typeVar string, url string) *EntityV3MetadataLinksItems {
	this := EntityV3MetadataLinksItems{}
	this.Name = name
	this.Type = typeVar
	this.Url = url
	return &this
}

// NewEntityV3MetadataLinksItemsWithDefaults instantiates a new EntityV3MetadataLinksItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3MetadataLinksItemsWithDefaults() *EntityV3MetadataLinksItems {
	this := EntityV3MetadataLinksItems{}
	var typeVar string = "other"
	this.Type = typeVar
	return &this
}

// GetName returns the Name field value.
func (o *EntityV3MetadataLinksItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntityV3MetadataLinksItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EntityV3MetadataLinksItems) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *EntityV3MetadataLinksItems) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3MetadataLinksItems) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *EntityV3MetadataLinksItems) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *EntityV3MetadataLinksItems) SetProvider(v string) {
	o.Provider = &v
}

// GetType returns the Type field value.
func (o *EntityV3MetadataLinksItems) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntityV3MetadataLinksItems) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EntityV3MetadataLinksItems) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value.
func (o *EntityV3MetadataLinksItems) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *EntityV3MetadataLinksItems) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *EntityV3MetadataLinksItems) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3MetadataLinksItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3MetadataLinksItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string `json:"name"`
		Provider *string `json:"provider,omitempty"`
		Type     *string `json:"type"`
		Url      *string `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	o.Name = *all.Name
	o.Provider = all.Provider
	o.Type = *all.Type
	o.Url = *all.Url

	return nil
}
