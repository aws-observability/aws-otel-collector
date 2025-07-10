// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComponentGrid A grid component. The grid component is the root canvas for an app and contains all other components.
type ComponentGrid struct {
	// Events to listen for on the grid component.
	Events []AppBuilderEvent `json:"events,omitempty"`
	// The ID of the grid component. This property is deprecated; use `name` to identify individual components instead.
	Id *string `json:"id,omitempty"`
	// A unique identifier for this grid component. This name is also visible in the app editor.
	Name string `json:"name"`
	// Properties of a grid component.
	Properties ComponentGridProperties `json:"properties"`
	// The grid component type.
	Type ComponentGridType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentGrid instantiates a new ComponentGrid object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentGrid(name string, properties ComponentGridProperties, typeVar ComponentGridType) *ComponentGrid {
	this := ComponentGrid{}
	this.Name = name
	this.Properties = properties
	this.Type = typeVar
	return &this
}

// NewComponentGridWithDefaults instantiates a new ComponentGrid object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentGridWithDefaults() *ComponentGrid {
	this := ComponentGrid{}
	var typeVar ComponentGridType = COMPONENTGRIDTYPE_GRID
	this.Type = typeVar
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ComponentGrid) GetEvents() []AppBuilderEvent {
	if o == nil || o.Events == nil {
		var ret []AppBuilderEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentGrid) GetEventsOk() (*[]AppBuilderEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ComponentGrid) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []AppBuilderEvent and assigns it to the Events field.
func (o *ComponentGrid) SetEvents(v []AppBuilderEvent) {
	o.Events = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentGrid) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentGrid) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentGrid) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComponentGrid) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value.
func (o *ComponentGrid) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ComponentGrid) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ComponentGrid) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value.
func (o *ComponentGrid) GetProperties() ComponentGridProperties {
	if o == nil {
		var ret ComponentGridProperties
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ComponentGrid) GetPropertiesOk() (*ComponentGridProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value.
func (o *ComponentGrid) SetProperties(v ComponentGridProperties) {
	o.Properties = v
}

// GetType returns the Type field value.
func (o *ComponentGrid) GetType() ComponentGridType {
	if o == nil {
		var ret ComponentGridType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ComponentGrid) GetTypeOk() (*ComponentGridType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ComponentGrid) SetType(v ComponentGridType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentGrid) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["properties"] = o.Properties
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentGrid) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Events     []AppBuilderEvent        `json:"events,omitempty"`
		Id         *string                  `json:"id,omitempty"`
		Name       *string                  `json:"name"`
		Properties *ComponentGridProperties `json:"properties"`
		Type       *ComponentGridType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Properties == nil {
		return fmt.Errorf("required field properties missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"events", "id", "name", "properties", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Events = all.Events
	o.Id = all.Id
	o.Name = *all.Name
	if all.Properties.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Properties = *all.Properties
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
