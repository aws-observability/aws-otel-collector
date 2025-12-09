// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackInnerWidgets Powerpack group widget definition of individual widgets.
type PowerpackInnerWidgets struct {
	// Information about widget.
	Definition map[string]interface{} `json:"definition"`
	// Powerpack inner widget layout.
	Layout *PowerpackInnerWidgetLayout `json:"layout,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPowerpackInnerWidgets instantiates a new PowerpackInnerWidgets object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackInnerWidgets(definition map[string]interface{}) *PowerpackInnerWidgets {
	this := PowerpackInnerWidgets{}
	this.Definition = definition
	return &this
}

// NewPowerpackInnerWidgetsWithDefaults instantiates a new PowerpackInnerWidgets object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackInnerWidgetsWithDefaults() *PowerpackInnerWidgets {
	this := PowerpackInnerWidgets{}
	return &this
}

// GetDefinition returns the Definition field value.
func (o *PowerpackInnerWidgets) GetDefinition() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *PowerpackInnerWidgets) GetDefinitionOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *PowerpackInnerWidgets) SetDefinition(v map[string]interface{}) {
	o.Definition = v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *PowerpackInnerWidgets) GetLayout() PowerpackInnerWidgetLayout {
	if o == nil || o.Layout == nil {
		var ret PowerpackInnerWidgetLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackInnerWidgets) GetLayoutOk() (*PowerpackInnerWidgetLayout, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *PowerpackInnerWidgets) HasLayout() bool {
	return o != nil && o.Layout != nil
}

// SetLayout gets a reference to the given PowerpackInnerWidgetLayout and assigns it to the Layout field.
func (o *PowerpackInnerWidgets) SetLayout(v PowerpackInnerWidgetLayout) {
	o.Layout = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackInnerWidgets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["definition"] = o.Definition
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackInnerWidgets) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition *map[string]interface{}     `json:"definition"`
		Layout     *PowerpackInnerWidgetLayout `json:"layout,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"definition", "layout"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Definition = *all.Definition
	if all.Layout != nil && all.Layout.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Layout = all.Layout

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
