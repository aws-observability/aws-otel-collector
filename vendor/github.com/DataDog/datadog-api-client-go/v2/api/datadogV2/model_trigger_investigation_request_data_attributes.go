// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerInvestigationRequestDataAttributes Attributes for the trigger investigation request.
type TriggerInvestigationRequestDataAttributes struct {
	// The trigger definition for starting an investigation.
	Trigger TriggerAttributes `json:"trigger"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTriggerInvestigationRequestDataAttributes instantiates a new TriggerInvestigationRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTriggerInvestigationRequestDataAttributes(trigger TriggerAttributes) *TriggerInvestigationRequestDataAttributes {
	this := TriggerInvestigationRequestDataAttributes{}
	this.Trigger = trigger
	return &this
}

// NewTriggerInvestigationRequestDataAttributesWithDefaults instantiates a new TriggerInvestigationRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTriggerInvestigationRequestDataAttributesWithDefaults() *TriggerInvestigationRequestDataAttributes {
	this := TriggerInvestigationRequestDataAttributes{}
	return &this
}

// GetTrigger returns the Trigger field value.
func (o *TriggerInvestigationRequestDataAttributes) GetTrigger() TriggerAttributes {
	if o == nil {
		var ret TriggerAttributes
		return ret
	}
	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *TriggerInvestigationRequestDataAttributes) GetTriggerOk() (*TriggerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value.
func (o *TriggerInvestigationRequestDataAttributes) SetTrigger(v TriggerAttributes) {
	o.Trigger = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TriggerInvestigationRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["trigger"] = o.Trigger

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TriggerInvestigationRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Trigger *TriggerAttributes `json:"trigger"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Trigger == nil {
		return fmt.Errorf("required field trigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"trigger"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Trigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Trigger = *all.Trigger

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
