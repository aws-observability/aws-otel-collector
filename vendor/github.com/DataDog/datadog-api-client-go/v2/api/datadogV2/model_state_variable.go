// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StateVariable A variable, which can be set and read by other components in the app.
type StateVariable struct {
	// The ID of the state variable.
	Id uuid.UUID `json:"id"`
	// A unique identifier for this state variable. This name is also used to access the variable's value throughout the app.
	Name string `json:"name"`
	// The properties of the state variable.
	Properties StateVariableProperties `json:"properties"`
	// The state variable type.
	Type StateVariableType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStateVariable instantiates a new StateVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStateVariable(id uuid.UUID, name string, properties StateVariableProperties, typeVar StateVariableType) *StateVariable {
	this := StateVariable{}
	this.Id = id
	this.Name = name
	this.Properties = properties
	this.Type = typeVar
	return &this
}

// NewStateVariableWithDefaults instantiates a new StateVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStateVariableWithDefaults() *StateVariable {
	this := StateVariable{}
	var typeVar StateVariableType = STATEVARIABLETYPE_STATEVARIABLE
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *StateVariable) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StateVariable) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *StateVariable) SetId(v uuid.UUID) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *StateVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StateVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *StateVariable) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value.
func (o *StateVariable) GetProperties() StateVariableProperties {
	if o == nil {
		var ret StateVariableProperties
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *StateVariable) GetPropertiesOk() (*StateVariableProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value.
func (o *StateVariable) SetProperties(v StateVariableProperties) {
	o.Properties = v
}

// GetType returns the Type field value.
func (o *StateVariable) GetType() StateVariableType {
	if o == nil {
		var ret StateVariableType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StateVariable) GetTypeOk() (*StateVariableType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StateVariable) SetType(v StateVariableType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["properties"] = o.Properties
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id         *uuid.UUID               `json:"id"`
		Name       *string                  `json:"name"`
		Properties *StateVariableProperties `json:"properties"`
		Type       *StateVariableType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "properties", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
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
