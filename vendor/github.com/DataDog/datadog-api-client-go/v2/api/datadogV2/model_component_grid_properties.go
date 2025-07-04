// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComponentGridProperties Properties of a grid component.
type ComponentGridProperties struct {
	// The background color of the grid.
	BackgroundColor *string `json:"backgroundColor,omitempty"`
	// The child components of the grid.
	Children []Component `json:"children,omitempty"`
	// Whether the grid component and its children are visible. If a string, it must be a valid JavaScript expression that evaluates to a boolean.
	IsVisible *ComponentGridPropertiesIsVisible `json:"isVisible,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentGridProperties instantiates a new ComponentGridProperties object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentGridProperties() *ComponentGridProperties {
	this := ComponentGridProperties{}
	var backgroundColor string = "default"
	this.BackgroundColor = &backgroundColor
	return &this
}

// NewComponentGridPropertiesWithDefaults instantiates a new ComponentGridProperties object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentGridPropertiesWithDefaults() *ComponentGridProperties {
	this := ComponentGridProperties{}
	var backgroundColor string = "default"
	this.BackgroundColor = &backgroundColor
	return &this
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise.
func (o *ComponentGridProperties) GetBackgroundColor() string {
	if o == nil || o.BackgroundColor == nil {
		var ret string
		return ret
	}
	return *o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentGridProperties) GetBackgroundColorOk() (*string, bool) {
	if o == nil || o.BackgroundColor == nil {
		return nil, false
	}
	return o.BackgroundColor, true
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *ComponentGridProperties) HasBackgroundColor() bool {
	return o != nil && o.BackgroundColor != nil
}

// SetBackgroundColor gets a reference to the given string and assigns it to the BackgroundColor field.
func (o *ComponentGridProperties) SetBackgroundColor(v string) {
	o.BackgroundColor = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *ComponentGridProperties) GetChildren() []Component {
	if o == nil || o.Children == nil {
		var ret []Component
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentGridProperties) GetChildrenOk() (*[]Component, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return &o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ComponentGridProperties) HasChildren() bool {
	return o != nil && o.Children != nil
}

// SetChildren gets a reference to the given []Component and assigns it to the Children field.
func (o *ComponentGridProperties) SetChildren(v []Component) {
	o.Children = v
}

// GetIsVisible returns the IsVisible field value if set, zero value otherwise.
func (o *ComponentGridProperties) GetIsVisible() ComponentGridPropertiesIsVisible {
	if o == nil || o.IsVisible == nil {
		var ret ComponentGridPropertiesIsVisible
		return ret
	}
	return *o.IsVisible
}

// GetIsVisibleOk returns a tuple with the IsVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentGridProperties) GetIsVisibleOk() (*ComponentGridPropertiesIsVisible, bool) {
	if o == nil || o.IsVisible == nil {
		return nil, false
	}
	return o.IsVisible, true
}

// HasIsVisible returns a boolean if a field has been set.
func (o *ComponentGridProperties) HasIsVisible() bool {
	return o != nil && o.IsVisible != nil
}

// SetIsVisible gets a reference to the given ComponentGridPropertiesIsVisible and assigns it to the IsVisible field.
func (o *ComponentGridProperties) SetIsVisible(v ComponentGridPropertiesIsVisible) {
	o.IsVisible = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentGridProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BackgroundColor != nil {
		toSerialize["backgroundColor"] = o.BackgroundColor
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
func (o *ComponentGridProperties) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BackgroundColor *string                           `json:"backgroundColor,omitempty"`
		Children        []Component                       `json:"children,omitempty"`
		IsVisible       *ComponentGridPropertiesIsVisible `json:"isVisible,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"backgroundColor", "children", "isVisible"})
	} else {
		return err
	}
	o.BackgroundColor = all.BackgroundColor
	o.Children = all.Children
	o.IsVisible = all.IsVisible

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
