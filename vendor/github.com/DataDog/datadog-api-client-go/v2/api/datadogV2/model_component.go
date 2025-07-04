// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Component [Definition of a UI component in the app](https://docs.datadoghq.com/service_management/app_builder/components/)
type Component struct {
	// Events to listen for on the UI component.
	Events []AppBuilderEvent `json:"events,omitempty"`
	// The ID of the UI component. This property is deprecated; use `name` to identify individual components instead.
	Id datadog.NullableString `json:"id,omitempty"`
	// A unique identifier for this UI component. This name is also visible in the app editor.
	Name string `json:"name"`
	// Properties of a UI component. Different component types can have their own additional unique properties. See the [components documentation](https://docs.datadoghq.com/service_management/app_builder/components/) for more detail on each component type and its properties.
	Properties ComponentProperties `json:"properties"`
	// The UI component type.
	Type ComponentType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponent instantiates a new Component object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponent(name string, properties ComponentProperties, typeVar ComponentType) *Component {
	this := Component{}
	this.Name = name
	this.Properties = properties
	this.Type = typeVar
	return &this
}

// NewComponentWithDefaults instantiates a new Component object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentWithDefaults() *Component {
	this := Component{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *Component) GetEvents() []AppBuilderEvent {
	if o == nil || o.Events == nil {
		var ret []AppBuilderEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetEventsOk() (*[]AppBuilderEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *Component) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []AppBuilderEvent and assigns it to the Events field.
func (o *Component) SetEvents(v []AppBuilderEvent) {
	o.Events = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Component) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Component) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Component) HasId() bool {
	return o != nil && o.Id.IsSet()
}

// SetId gets a reference to the given datadog.NullableString and assigns it to the Id field.
func (o *Component) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil.
func (o *Component) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil.
func (o *Component) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value.
func (o *Component) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Component) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Component) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value.
func (o *Component) GetProperties() ComponentProperties {
	if o == nil {
		var ret ComponentProperties
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *Component) GetPropertiesOk() (*ComponentProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value.
func (o *Component) SetProperties(v ComponentProperties) {
	o.Properties = v
}

// GetType returns the Type field value.
func (o *Component) GetType() ComponentType {
	if o == nil {
		var ret ComponentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Component) GetTypeOk() (*ComponentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Component) SetType(v ComponentType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Component) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
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
func (o *Component) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Events     []AppBuilderEvent      `json:"events,omitempty"`
		Id         datadog.NullableString `json:"id,omitempty"`
		Name       *string                `json:"name"`
		Properties *ComponentProperties   `json:"properties"`
		Type       *ComponentType         `json:"type"`
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
