// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleModifiedBy The user or Datadog system who last modified the rule.
type AutomationRuleModifiedBy struct {
	// The actor's identifier (a user UUID or a system identifier).
	Id string `json:"id"`
	// The name of the actor.
	Name string `json:"name"`
	// Whether the actor is a user or the Datadog system.
	Type AutomationRuleActorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutomationRuleModifiedBy instantiates a new AutomationRuleModifiedBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomationRuleModifiedBy(id string, name string, typeVar AutomationRuleActorType) *AutomationRuleModifiedBy {
	this := AutomationRuleModifiedBy{}
	this.Id = id
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewAutomationRuleModifiedByWithDefaults instantiates a new AutomationRuleModifiedBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutomationRuleModifiedByWithDefaults() *AutomationRuleModifiedBy {
	this := AutomationRuleModifiedBy{}
	return &this
}

// GetId returns the Id field value.
func (o *AutomationRuleModifiedBy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleModifiedBy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AutomationRuleModifiedBy) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *AutomationRuleModifiedBy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleModifiedBy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AutomationRuleModifiedBy) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *AutomationRuleModifiedBy) GetType() AutomationRuleActorType {
	if o == nil {
		var ret AutomationRuleActorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleModifiedBy) GetTypeOk() (*AutomationRuleActorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AutomationRuleModifiedBy) SetType(v AutomationRuleActorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutomationRuleModifiedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutomationRuleModifiedBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                  `json:"id"`
		Name *string                  `json:"name"`
		Type *AutomationRuleActorType `json:"type"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Name = *all.Name
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
