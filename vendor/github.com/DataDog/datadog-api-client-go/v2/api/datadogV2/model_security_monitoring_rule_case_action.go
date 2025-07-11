// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleCaseAction Action to perform when a signal is triggered. Only available for Application Security rule type.
type SecurityMonitoringRuleCaseAction struct {
	// Options for the rule action
	Options *SecurityMonitoringRuleCaseActionOptions `json:"options,omitempty"`
	// The action type.
	Type *SecurityMonitoringRuleCaseActionType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleCaseAction instantiates a new SecurityMonitoringRuleCaseAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleCaseAction() *SecurityMonitoringRuleCaseAction {
	this := SecurityMonitoringRuleCaseAction{}
	return &this
}

// NewSecurityMonitoringRuleCaseActionWithDefaults instantiates a new SecurityMonitoringRuleCaseAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleCaseActionWithDefaults() *SecurityMonitoringRuleCaseAction {
	this := SecurityMonitoringRuleCaseAction{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCaseAction) GetOptions() SecurityMonitoringRuleCaseActionOptions {
	if o == nil || o.Options == nil {
		var ret SecurityMonitoringRuleCaseActionOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCaseAction) GetOptionsOk() (*SecurityMonitoringRuleCaseActionOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCaseAction) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given SecurityMonitoringRuleCaseActionOptions and assigns it to the Options field.
func (o *SecurityMonitoringRuleCaseAction) SetOptions(v SecurityMonitoringRuleCaseActionOptions) {
	o.Options = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCaseAction) GetType() SecurityMonitoringRuleCaseActionType {
	if o == nil || o.Type == nil {
		var ret SecurityMonitoringRuleCaseActionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCaseAction) GetTypeOk() (*SecurityMonitoringRuleCaseActionType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCaseAction) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SecurityMonitoringRuleCaseActionType and assigns it to the Type field.
func (o *SecurityMonitoringRuleCaseAction) SetType(v SecurityMonitoringRuleCaseActionType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleCaseAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleCaseAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Options *SecurityMonitoringRuleCaseActionOptions `json:"options,omitempty"`
		Type    *SecurityMonitoringRuleCaseActionType    `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"options", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
