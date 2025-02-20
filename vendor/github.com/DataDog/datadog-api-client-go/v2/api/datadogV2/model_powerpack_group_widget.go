// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackGroupWidget Powerpack group widget definition object.
type PowerpackGroupWidget struct {
	// Powerpack group widget object.
	Definition PowerpackGroupWidgetDefinition `json:"definition"`
	// Powerpack group widget layout.
	Layout *PowerpackGroupWidgetLayout `json:"layout,omitempty"`
	// The available timeframes depend on the widget you are using.
	LiveSpan *WidgetLiveSpan `json:"live_span,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPowerpackGroupWidget instantiates a new PowerpackGroupWidget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackGroupWidget(definition PowerpackGroupWidgetDefinition) *PowerpackGroupWidget {
	this := PowerpackGroupWidget{}
	this.Definition = definition
	return &this
}

// NewPowerpackGroupWidgetWithDefaults instantiates a new PowerpackGroupWidget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackGroupWidgetWithDefaults() *PowerpackGroupWidget {
	this := PowerpackGroupWidget{}
	return &this
}

// GetDefinition returns the Definition field value.
func (o *PowerpackGroupWidget) GetDefinition() PowerpackGroupWidgetDefinition {
	if o == nil {
		var ret PowerpackGroupWidgetDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *PowerpackGroupWidget) GetDefinitionOk() (*PowerpackGroupWidgetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *PowerpackGroupWidget) SetDefinition(v PowerpackGroupWidgetDefinition) {
	o.Definition = v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *PowerpackGroupWidget) GetLayout() PowerpackGroupWidgetLayout {
	if o == nil || o.Layout == nil {
		var ret PowerpackGroupWidgetLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackGroupWidget) GetLayoutOk() (*PowerpackGroupWidgetLayout, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *PowerpackGroupWidget) HasLayout() bool {
	return o != nil && o.Layout != nil
}

// SetLayout gets a reference to the given PowerpackGroupWidgetLayout and assigns it to the Layout field.
func (o *PowerpackGroupWidget) SetLayout(v PowerpackGroupWidgetLayout) {
	o.Layout = &v
}

// GetLiveSpan returns the LiveSpan field value if set, zero value otherwise.
func (o *PowerpackGroupWidget) GetLiveSpan() WidgetLiveSpan {
	if o == nil || o.LiveSpan == nil {
		var ret WidgetLiveSpan
		return ret
	}
	return *o.LiveSpan
}

// GetLiveSpanOk returns a tuple with the LiveSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackGroupWidget) GetLiveSpanOk() (*WidgetLiveSpan, bool) {
	if o == nil || o.LiveSpan == nil {
		return nil, false
	}
	return o.LiveSpan, true
}

// HasLiveSpan returns a boolean if a field has been set.
func (o *PowerpackGroupWidget) HasLiveSpan() bool {
	return o != nil && o.LiveSpan != nil
}

// SetLiveSpan gets a reference to the given WidgetLiveSpan and assigns it to the LiveSpan field.
func (o *PowerpackGroupWidget) SetLiveSpan(v WidgetLiveSpan) {
	o.LiveSpan = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackGroupWidget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["definition"] = o.Definition
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if o.LiveSpan != nil {
		toSerialize["live_span"] = o.LiveSpan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackGroupWidget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition *PowerpackGroupWidgetDefinition `json:"definition"`
		Layout     *PowerpackGroupWidgetLayout     `json:"layout,omitempty"`
		LiveSpan   *WidgetLiveSpan                 `json:"live_span,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"definition", "layout", "live_span"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition
	if all.Layout != nil && all.Layout.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Layout = all.Layout
	if all.LiveSpan != nil && !all.LiveSpan.IsValid() {
		hasInvalidField = true
	} else {
		o.LiveSpan = all.LiveSpan
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
