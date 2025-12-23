// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRuleOptionsFaultyDeploymentDetection Faulty deployment detection options for deployment rules.
type DeploymentRuleOptionsFaultyDeploymentDetection struct {
	// The duration for faulty deployment detection.
	Duration *int64 `json:"duration,omitempty"`
	// Resources to exclude from faulty deployment detection.
	ExcludedResources []string `json:"excluded_resources,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDeploymentRuleOptionsFaultyDeploymentDetection instantiates a new DeploymentRuleOptionsFaultyDeploymentDetection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentRuleOptionsFaultyDeploymentDetection() *DeploymentRuleOptionsFaultyDeploymentDetection {
	this := DeploymentRuleOptionsFaultyDeploymentDetection{}
	return &this
}

// NewDeploymentRuleOptionsFaultyDeploymentDetectionWithDefaults instantiates a new DeploymentRuleOptionsFaultyDeploymentDetection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentRuleOptionsFaultyDeploymentDetectionWithDefaults() *DeploymentRuleOptionsFaultyDeploymentDetection {
	this := DeploymentRuleOptionsFaultyDeploymentDetection{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) SetDuration(v int64) {
	o.Duration = &v
}

// GetExcludedResources returns the ExcludedResources field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) GetExcludedResources() []string {
	if o == nil || o.ExcludedResources == nil {
		var ret []string
		return ret
	}
	return o.ExcludedResources
}

// GetExcludedResourcesOk returns a tuple with the ExcludedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) GetExcludedResourcesOk() (*[]string, bool) {
	if o == nil || o.ExcludedResources == nil {
		return nil, false
	}
	return &o.ExcludedResources, true
}

// HasExcludedResources returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) HasExcludedResources() bool {
	return o != nil && o.ExcludedResources != nil
}

// SetExcludedResources gets a reference to the given []string and assigns it to the ExcludedResources field.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) SetExcludedResources(v []string) {
	o.ExcludedResources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentRuleOptionsFaultyDeploymentDetection) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration          *int64   `json:"duration,omitempty"`
		ExcludedResources []string `json:"excluded_resources,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Duration = all.Duration
	o.ExcludedResources = all.ExcludedResources

	return nil
}
