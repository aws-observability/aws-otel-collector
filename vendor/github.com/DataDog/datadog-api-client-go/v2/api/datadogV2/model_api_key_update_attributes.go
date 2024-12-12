// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APIKeyUpdateAttributes Attributes used to update an API Key.
type APIKeyUpdateAttributes struct {
	// The APIKeyUpdateAttributes category.
	Category *string `json:"category,omitempty"`
	// Name of the API key.
	Name string `json:"name"`
	// The APIKeyUpdateAttributes remote_config_read_enabled.
	RemoteConfigReadEnabled *bool `json:"remote_config_read_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAPIKeyUpdateAttributes instantiates a new APIKeyUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAPIKeyUpdateAttributes(name string) *APIKeyUpdateAttributes {
	this := APIKeyUpdateAttributes{}
	this.Name = name
	return &this
}

// NewAPIKeyUpdateAttributesWithDefaults instantiates a new APIKeyUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAPIKeyUpdateAttributesWithDefaults() *APIKeyUpdateAttributes {
	this := APIKeyUpdateAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *APIKeyUpdateAttributes) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeyUpdateAttributes) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *APIKeyUpdateAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *APIKeyUpdateAttributes) SetCategory(v string) {
	o.Category = &v
}

// GetName returns the Name field value.
func (o *APIKeyUpdateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *APIKeyUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *APIKeyUpdateAttributes) SetName(v string) {
	o.Name = v
}

// GetRemoteConfigReadEnabled returns the RemoteConfigReadEnabled field value if set, zero value otherwise.
func (o *APIKeyUpdateAttributes) GetRemoteConfigReadEnabled() bool {
	if o == nil || o.RemoteConfigReadEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RemoteConfigReadEnabled
}

// GetRemoteConfigReadEnabledOk returns a tuple with the RemoteConfigReadEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeyUpdateAttributes) GetRemoteConfigReadEnabledOk() (*bool, bool) {
	if o == nil || o.RemoteConfigReadEnabled == nil {
		return nil, false
	}
	return o.RemoteConfigReadEnabled, true
}

// HasRemoteConfigReadEnabled returns a boolean if a field has been set.
func (o *APIKeyUpdateAttributes) HasRemoteConfigReadEnabled() bool {
	return o != nil && o.RemoteConfigReadEnabled != nil
}

// SetRemoteConfigReadEnabled gets a reference to the given bool and assigns it to the RemoteConfigReadEnabled field.
func (o *APIKeyUpdateAttributes) SetRemoteConfigReadEnabled(v bool) {
	o.RemoteConfigReadEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o APIKeyUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	toSerialize["name"] = o.Name
	if o.RemoteConfigReadEnabled != nil {
		toSerialize["remote_config_read_enabled"] = o.RemoteConfigReadEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *APIKeyUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category                *string `json:"category,omitempty"`
		Name                    *string `json:"name"`
		RemoteConfigReadEnabled *bool   `json:"remote_config_read_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "name", "remote_config_read_enabled"})
	} else {
		return err
	}
	o.Category = all.Category
	o.Name = *all.Name
	o.RemoteConfigReadEnabled = all.RemoteConfigReadEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
