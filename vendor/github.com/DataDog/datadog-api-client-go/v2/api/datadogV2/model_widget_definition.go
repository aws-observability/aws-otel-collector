// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetDefinition The definition of a widget, including its type and configuration.
type WidgetDefinition struct {
	// The display title of the widget.
	Title string `json:"title"`
	// Widget types that are allowed to be stored as individual records.
	// This is not a complete list of dashboard and notebook widget types.
	Type WidgetType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetDefinition instantiates a new WidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetDefinition(title string, typeVar WidgetType) *WidgetDefinition {
	this := WidgetDefinition{}
	this.Title = title
	this.Type = typeVar
	return &this
}

// NewWidgetDefinitionWithDefaults instantiates a new WidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetDefinitionWithDefaults() *WidgetDefinition {
	this := WidgetDefinition{}
	return &this
}

// GetTitle returns the Title field value.
func (o *WidgetDefinition) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *WidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *WidgetDefinition) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value.
func (o *WidgetDefinition) GetType() WidgetType {
	if o == nil {
		var ret WidgetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WidgetDefinition) GetTypeOk() (*WidgetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WidgetDefinition) SetType(v WidgetType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Title *string     `json:"title"`
		Type  *WidgetType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Title = *all.Title
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
