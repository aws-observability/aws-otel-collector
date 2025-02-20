// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot2Link Service's external links.
type ServiceDefinitionV2Dot2Link struct {
	// Link name.
	Name string `json:"name"`
	// Link provider.
	Provider *string `json:"provider,omitempty"`
	// Link type. Datadog recognizes the following types: `runbook`, `doc`, `repo`, `dashboard`, and `other`.
	Type string `json:"type"`
	// Link URL.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceDefinitionV2Dot2Link instantiates a new ServiceDefinitionV2Dot2Link object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionV2Dot2Link(name string, typeVar string, url string) *ServiceDefinitionV2Dot2Link {
	this := ServiceDefinitionV2Dot2Link{}
	this.Name = name
	this.Type = typeVar
	this.Url = url
	return &this
}

// NewServiceDefinitionV2Dot2LinkWithDefaults instantiates a new ServiceDefinitionV2Dot2Link object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionV2Dot2LinkWithDefaults() *ServiceDefinitionV2Dot2Link {
	this := ServiceDefinitionV2Dot2Link{}
	return &this
}

// GetName returns the Name field value.
func (o *ServiceDefinitionV2Dot2Link) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2Link) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ServiceDefinitionV2Dot2Link) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2Link) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2Link) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2Link) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ServiceDefinitionV2Dot2Link) SetProvider(v string) {
	o.Provider = &v
}

// GetType returns the Type field value.
func (o *ServiceDefinitionV2Dot2Link) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2Link) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ServiceDefinitionV2Dot2Link) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value.
func (o *ServiceDefinitionV2Dot2Link) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2Link) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *ServiceDefinitionV2Dot2Link) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionV2Dot2Link) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionV2Dot2Link) UnmarshalJSON(bytes []byte) (err error) {
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "provider", "type", "url"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Provider = all.Provider
	o.Type = *all.Type
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
