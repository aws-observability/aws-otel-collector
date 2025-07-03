// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComponentProperties Properties of a UI component. Different component types can have their own additional unique properties. See the [components documentation](https://docs.datadoghq.com/service_management/app_builder/components/) for more detail on each component type and its properties.
type ComponentProperties struct {
	// The child components of the UI component.
	Children []Component `json:"children,omitempty"`
	// Whether the UI component is visible. If this is a string, it must be a valid JavaScript expression that evaluates to a boolean.
	IsVisible *ComponentPropertiesIsVisible `json:"isVisible,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentProperties instantiates a new ComponentProperties object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentProperties() *ComponentProperties {
	this := ComponentProperties{}
	return &this
}

// NewComponentPropertiesWithDefaults instantiates a new ComponentProperties object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentPropertiesWithDefaults() *ComponentProperties {
	this := ComponentProperties{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *ComponentProperties) GetChildren() []Component {
	if o == nil || o.Children == nil {
		var ret []Component
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentProperties) GetChildrenOk() (*[]Component, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return &o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ComponentProperties) HasChildren() bool {
	return o != nil && o.Children != nil
}

// SetChildren gets a reference to the given []Component and assigns it to the Children field.
func (o *ComponentProperties) SetChildren(v []Component) {
	o.Children = v
}

// GetIsVisible returns the IsVisible field value if set, zero value otherwise.
func (o *ComponentProperties) GetIsVisible() ComponentPropertiesIsVisible {
	if o == nil || o.IsVisible == nil {
		var ret ComponentPropertiesIsVisible
		return ret
	}
	return *o.IsVisible
}

// GetIsVisibleOk returns a tuple with the IsVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentProperties) GetIsVisibleOk() (*ComponentPropertiesIsVisible, bool) {
	if o == nil || o.IsVisible == nil {
		return nil, false
	}
	return o.IsVisible, true
}

// HasIsVisible returns a boolean if a field has been set.
func (o *ComponentProperties) HasIsVisible() bool {
	return o != nil && o.IsVisible != nil
}

// SetIsVisible gets a reference to the given ComponentPropertiesIsVisible and assigns it to the IsVisible field.
func (o *ComponentProperties) SetIsVisible(v ComponentPropertiesIsVisible) {
	o.IsVisible = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.IsVisible != nil {
		toSerialize["isVisible"] = o.IsVisible
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentProperties) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Children  []Component                   `json:"children,omitempty"`
		IsVisible *ComponentPropertiesIsVisible `json:"isVisible,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"children", "isVisible"})
	} else {
		return err
	}
	o.Children = all.Children
	o.IsVisible = all.IsVisible

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
