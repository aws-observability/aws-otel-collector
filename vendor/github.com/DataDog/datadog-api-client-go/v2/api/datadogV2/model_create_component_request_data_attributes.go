// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateComponentRequestDataAttributes The supported attributes for creating a component.
type CreateComponentRequestDataAttributes struct {
	// If creating a component of type `group`, the components to create within the group.
	Components []CreateComponentRequestDataAttributesComponentsItems `json:"components,omitempty"`
	// The name of the component.
	Name string `json:"name"`
	// The zero-indexed position of the component.
	Position int64 `json:"position"`
	// The type of the component.
	Type CreateComponentRequestDataAttributesType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateComponentRequestDataAttributes instantiates a new CreateComponentRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateComponentRequestDataAttributes(name string, position int64, typeVar CreateComponentRequestDataAttributesType) *CreateComponentRequestDataAttributes {
	this := CreateComponentRequestDataAttributes{}
	this.Name = name
	this.Position = position
	this.Type = typeVar
	return &this
}

// NewCreateComponentRequestDataAttributesWithDefaults instantiates a new CreateComponentRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateComponentRequestDataAttributesWithDefaults() *CreateComponentRequestDataAttributes {
	this := CreateComponentRequestDataAttributes{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *CreateComponentRequestDataAttributes) GetComponents() []CreateComponentRequestDataAttributesComponentsItems {
	if o == nil || o.Components == nil {
		var ret []CreateComponentRequestDataAttributesComponentsItems
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComponentRequestDataAttributes) GetComponentsOk() (*[]CreateComponentRequestDataAttributesComponentsItems, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *CreateComponentRequestDataAttributes) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []CreateComponentRequestDataAttributesComponentsItems and assigns it to the Components field.
func (o *CreateComponentRequestDataAttributes) SetComponents(v []CreateComponentRequestDataAttributesComponentsItems) {
	o.Components = v
}

// GetName returns the Name field value.
func (o *CreateComponentRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateComponentRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateComponentRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetPosition returns the Position field value.
func (o *CreateComponentRequestDataAttributes) GetPosition() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *CreateComponentRequestDataAttributes) GetPositionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value.
func (o *CreateComponentRequestDataAttributes) SetPosition(v int64) {
	o.Position = v
}

// GetType returns the Type field value.
func (o *CreateComponentRequestDataAttributes) GetType() CreateComponentRequestDataAttributesType {
	if o == nil {
		var ret CreateComponentRequestDataAttributesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateComponentRequestDataAttributes) GetTypeOk() (*CreateComponentRequestDataAttributesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateComponentRequestDataAttributes) SetType(v CreateComponentRequestDataAttributesType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateComponentRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	toSerialize["name"] = o.Name
	toSerialize["position"] = o.Position
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateComponentRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components []CreateComponentRequestDataAttributesComponentsItems `json:"components,omitempty"`
		Name       *string                                               `json:"name"`
		Position   *int64                                                `json:"position"`
		Type       *CreateComponentRequestDataAttributesType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Position == nil {
		return fmt.Errorf("required field position missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components", "name", "position", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Components = all.Components
	o.Name = *all.Name
	o.Position = *all.Position
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
