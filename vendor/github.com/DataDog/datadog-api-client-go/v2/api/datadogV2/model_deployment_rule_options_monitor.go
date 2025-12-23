// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRuleOptionsMonitor Monitor options for deployment rules.
type DeploymentRuleOptionsMonitor struct {
	// Seconds the monitor needs to stay in OK status for the rule to pass.
	Duration *int64 `json:"duration,omitempty"`
	// Monitors that match this query are evaluated.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDeploymentRuleOptionsMonitor instantiates a new DeploymentRuleOptionsMonitor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentRuleOptionsMonitor(query string) *DeploymentRuleOptionsMonitor {
	this := DeploymentRuleOptionsMonitor{}
	this.Query = query
	return &this
}

// NewDeploymentRuleOptionsMonitorWithDefaults instantiates a new DeploymentRuleOptionsMonitor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentRuleOptionsMonitorWithDefaults() *DeploymentRuleOptionsMonitor {
	this := DeploymentRuleOptionsMonitor{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsMonitor) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitor) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsMonitor) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DeploymentRuleOptionsMonitor) SetDuration(v int64) {
	o.Duration = &v
}

// GetQuery returns the Query field value.
func (o *DeploymentRuleOptionsMonitor) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitor) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *DeploymentRuleOptionsMonitor) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentRuleOptionsMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["query"] = o.Query
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentRuleOptionsMonitor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration *int64  `json:"duration,omitempty"`
		Query    *string `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	o.Duration = all.Duration
	o.Query = *all.Query

	return nil
}
