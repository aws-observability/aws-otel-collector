// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleAction The action the rule can perform if triggered
type CloudWorkloadSecurityAgentRuleAction struct {
	// SECL expression used to target the container to apply the action on
	Filter *string `json:"filter,omitempty"`
	// Kill system call applied on the container matching the rule
	Kill *CloudWorkloadSecurityAgentRuleKill `json:"kill,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleAction instantiates a new CloudWorkloadSecurityAgentRuleAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleAction() *CloudWorkloadSecurityAgentRuleAction {
	this := CloudWorkloadSecurityAgentRuleAction{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleActionWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleActionWithDefaults() *CloudWorkloadSecurityAgentRuleAction {
	this := CloudWorkloadSecurityAgentRuleAction{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetFilter(v string) {
	o.Filter = &v
}

// GetKill returns the Kill field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetKill() CloudWorkloadSecurityAgentRuleKill {
	if o == nil || o.Kill == nil {
		var ret CloudWorkloadSecurityAgentRuleKill
		return ret
	}
	return *o.Kill
}

// GetKillOk returns a tuple with the Kill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetKillOk() (*CloudWorkloadSecurityAgentRuleKill, bool) {
	if o == nil || o.Kill == nil {
		return nil, false
	}
	return o.Kill, true
}

// HasKill returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasKill() bool {
	return o != nil && o.Kill != nil
}

// SetKill gets a reference to the given CloudWorkloadSecurityAgentRuleKill and assigns it to the Kill field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetKill(v CloudWorkloadSecurityAgentRuleKill) {
	o.Kill = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Kill != nil {
		toSerialize["kill"] = o.Kill
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentRuleAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter *string                             `json:"filter,omitempty"`
		Kill   *CloudWorkloadSecurityAgentRuleKill `json:"kill,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "kill"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Filter = all.Filter
	if all.Kill != nil && all.Kill.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Kill = all.Kill

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
