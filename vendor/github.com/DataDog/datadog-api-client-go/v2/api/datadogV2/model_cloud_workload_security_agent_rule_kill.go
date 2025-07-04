// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleKill Kill system call applied on the container matching the rule
type CloudWorkloadSecurityAgentRuleKill struct {
	// Supported signals for the kill system call
	Signal *string `json:"signal,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleKill instantiates a new CloudWorkloadSecurityAgentRuleKill object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleKill() *CloudWorkloadSecurityAgentRuleKill {
	this := CloudWorkloadSecurityAgentRuleKill{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleKillWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleKill object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleKillWithDefaults() *CloudWorkloadSecurityAgentRuleKill {
	this := CloudWorkloadSecurityAgentRuleKill{}
	return &this
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleKill) GetSignal() string {
	if o == nil || o.Signal == nil {
		var ret string
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleKill) GetSignalOk() (*string, bool) {
	if o == nil || o.Signal == nil {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleKill) HasSignal() bool {
	return o != nil && o.Signal != nil
}

// SetSignal gets a reference to the given string and assigns it to the Signal field.
func (o *CloudWorkloadSecurityAgentRuleKill) SetSignal(v string) {
	o.Signal = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleKill) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Signal != nil {
		toSerialize["signal"] = o.Signal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentRuleKill) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Signal *string `json:"signal,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"signal"})
	} else {
		return err
	}
	o.Signal = all.Signal

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
