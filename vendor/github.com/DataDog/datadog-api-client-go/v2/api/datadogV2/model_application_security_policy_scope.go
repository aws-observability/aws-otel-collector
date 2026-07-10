// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityPolicyScope The scope of the WAF policy.
type ApplicationSecurityPolicyScope struct {
	// The environment scope for the WAF policy.
	Env string `json:"env"`
	// The service scope for the WAF policy.
	Service string `json:"service"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityPolicyScope instantiates a new ApplicationSecurityPolicyScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityPolicyScope(env string, service string) *ApplicationSecurityPolicyScope {
	this := ApplicationSecurityPolicyScope{}
	this.Env = env
	this.Service = service
	return &this
}

// NewApplicationSecurityPolicyScopeWithDefaults instantiates a new ApplicationSecurityPolicyScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityPolicyScopeWithDefaults() *ApplicationSecurityPolicyScope {
	this := ApplicationSecurityPolicyScope{}
	return &this
}

// GetEnv returns the Env field value.
func (o *ApplicationSecurityPolicyScope) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyScope) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *ApplicationSecurityPolicyScope) SetEnv(v string) {
	o.Env = v
}

// GetService returns the Service field value.
func (o *ApplicationSecurityPolicyScope) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyScope) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *ApplicationSecurityPolicyScope) SetService(v string) {
	o.Service = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityPolicyScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["env"] = o.Env
	toSerialize["service"] = o.Service

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityPolicyScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Env     *string `json:"env"`
		Service *string `json:"service"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"env", "service"})
	} else {
		return err
	}
	o.Env = *all.Env
	o.Service = *all.Service

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
