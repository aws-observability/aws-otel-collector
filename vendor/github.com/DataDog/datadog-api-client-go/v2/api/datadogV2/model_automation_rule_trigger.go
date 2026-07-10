// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleTrigger Defines when the rule activates. Combines a trigger type (the case event to listen for) with optional trigger data (conditions that narrow when the trigger fires).
type AutomationRuleTrigger struct {
	// Additional configuration for the trigger, dependent on the trigger type. For `status_transitioned` triggers, specify `from_status_name` and `to_status_name`. For `attribute_value_changed` triggers, specify `field` and `change_type`.
	Data *AutomationRuleTriggerData `json:"data,omitempty"`
	// The case event that activates the automation rule.
	Type AutomationRuleTriggerType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutomationRuleTrigger instantiates a new AutomationRuleTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomationRuleTrigger(typeVar AutomationRuleTriggerType) *AutomationRuleTrigger {
	this := AutomationRuleTrigger{}
	this.Type = typeVar
	return &this
}

// NewAutomationRuleTriggerWithDefaults instantiates a new AutomationRuleTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutomationRuleTriggerWithDefaults() *AutomationRuleTrigger {
	this := AutomationRuleTrigger{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AutomationRuleTrigger) GetData() AutomationRuleTriggerData {
	if o == nil || o.Data == nil {
		var ret AutomationRuleTriggerData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleTrigger) GetDataOk() (*AutomationRuleTriggerData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AutomationRuleTrigger) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given AutomationRuleTriggerData and assigns it to the Data field.
func (o *AutomationRuleTrigger) SetData(v AutomationRuleTriggerData) {
	o.Data = &v
}

// GetType returns the Type field value.
func (o *AutomationRuleTrigger) GetType() AutomationRuleTriggerType {
	if o == nil {
		var ret AutomationRuleTriggerType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AutomationRuleTrigger) GetTypeOk() (*AutomationRuleTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AutomationRuleTrigger) SetType(v AutomationRuleTriggerType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutomationRuleTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutomationRuleTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *AutomationRuleTriggerData `json:"data,omitempty"`
		Type *AutomationRuleTriggerType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
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
