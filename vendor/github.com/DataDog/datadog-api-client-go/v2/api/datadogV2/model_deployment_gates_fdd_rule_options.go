// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesFDDRuleOptions Options for a `faulty_deployment_detection` rule.
type DeploymentGatesFDDRuleOptions struct {
	// Evaluation window in seconds. Maximum 7200 (2 hours).
	Duration *int64 `json:"duration,omitempty"`
	// APM resource names to exclude from analysis.
	ExcludedResources []string `json:"excluded_resources,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGatesFDDRuleOptions instantiates a new DeploymentGatesFDDRuleOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGatesFDDRuleOptions() *DeploymentGatesFDDRuleOptions {
	this := DeploymentGatesFDDRuleOptions{}
	return &this
}

// NewDeploymentGatesFDDRuleOptionsWithDefaults instantiates a new DeploymentGatesFDDRuleOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGatesFDDRuleOptionsWithDefaults() *DeploymentGatesFDDRuleOptions {
	this := DeploymentGatesFDDRuleOptions{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DeploymentGatesFDDRuleOptions) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesFDDRuleOptions) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DeploymentGatesFDDRuleOptions) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DeploymentGatesFDDRuleOptions) SetDuration(v int64) {
	o.Duration = &v
}

// GetExcludedResources returns the ExcludedResources field value if set, zero value otherwise.
func (o *DeploymentGatesFDDRuleOptions) GetExcludedResources() []string {
	if o == nil || o.ExcludedResources == nil {
		var ret []string
		return ret
	}
	return o.ExcludedResources
}

// GetExcludedResourcesOk returns a tuple with the ExcludedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesFDDRuleOptions) GetExcludedResourcesOk() (*[]string, bool) {
	if o == nil || o.ExcludedResources == nil {
		return nil, false
	}
	return &o.ExcludedResources, true
}

// HasExcludedResources returns a boolean if a field has been set.
func (o *DeploymentGatesFDDRuleOptions) HasExcludedResources() bool {
	return o != nil && o.ExcludedResources != nil
}

// SetExcludedResources gets a reference to the given []string and assigns it to the ExcludedResources field.
func (o *DeploymentGatesFDDRuleOptions) SetExcludedResources(v []string) {
	o.ExcludedResources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGatesFDDRuleOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.ExcludedResources != nil {
		toSerialize["excluded_resources"] = o.ExcludedResources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGatesFDDRuleOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration          *int64   `json:"duration,omitempty"`
		ExcludedResources []string `json:"excluded_resources,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration", "excluded_resources"})
	} else {
		return err
	}
	o.Duration = all.Duration
	o.ExcludedResources = all.ExcludedResources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
