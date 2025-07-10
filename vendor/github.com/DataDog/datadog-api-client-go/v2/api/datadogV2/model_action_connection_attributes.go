// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionConnectionAttributes The definition of `ActionConnectionAttributes` object.
type ActionConnectionAttributes struct {
	// The definition of `ActionConnectionIntegration` object.
	Integration ActionConnectionIntegration `json:"integration"`
	// Name of the connection
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionConnectionAttributes instantiates a new ActionConnectionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionConnectionAttributes(integration ActionConnectionIntegration, name string) *ActionConnectionAttributes {
	this := ActionConnectionAttributes{}
	this.Integration = integration
	this.Name = name
	return &this
}

// NewActionConnectionAttributesWithDefaults instantiates a new ActionConnectionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionConnectionAttributesWithDefaults() *ActionConnectionAttributes {
	this := ActionConnectionAttributes{}
	return &this
}

// GetIntegration returns the Integration field value.
func (o *ActionConnectionAttributes) GetIntegration() ActionConnectionIntegration {
	if o == nil {
		var ret ActionConnectionIntegration
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *ActionConnectionAttributes) GetIntegrationOk() (*ActionConnectionIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *ActionConnectionAttributes) SetIntegration(v ActionConnectionIntegration) {
	o.Integration = v
}

// GetName returns the Name field value.
func (o *ActionConnectionAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActionConnectionAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ActionConnectionAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionConnectionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["integration"] = o.Integration
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionConnectionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Integration *ActionConnectionIntegration `json:"integration"`
		Name        *string                      `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Integration == nil {
		return fmt.Errorf("required field integration missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"integration", "name"})
	} else {
		return err
	}
	o.Integration = *all.Integration
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
