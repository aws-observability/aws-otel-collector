// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentRequestAttributes Attributes to create a DORA deployment event.
type DORADeploymentRequestAttributes struct {
	// Environment name to where the service was deployed.
	Env *string `json:"env,omitempty"`
	// Unix timestamp when the deployment finished. It must be in nanoseconds, milliseconds, or seconds, and it should not be older than 1 hour.
	FinishedAt int64 `json:"finished_at"`
	// Git info for DORA Metrics events.
	Git *DORAGitInfo `json:"git,omitempty"`
	// Deployment ID.
	Id *string `json:"id,omitempty"`
	// Service name from a service available in the Service Catalog.
	Service string `json:"service"`
	// Unix timestamp when the deployment started. It must be in nanoseconds, milliseconds, or seconds.
	StartedAt int64 `json:"started_at"`
	// Version to correlate with [APM Deployment Tracking](https://docs.datadoghq.com/tracing/services/deployment_tracking/).
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORADeploymentRequestAttributes instantiates a new DORADeploymentRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORADeploymentRequestAttributes(finishedAt int64, service string, startedAt int64) *DORADeploymentRequestAttributes {
	this := DORADeploymentRequestAttributes{}
	this.FinishedAt = finishedAt
	this.Service = service
	this.StartedAt = startedAt
	return &this
}

// NewDORADeploymentRequestAttributesWithDefaults instantiates a new DORADeploymentRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORADeploymentRequestAttributesWithDefaults() *DORADeploymentRequestAttributes {
	this := DORADeploymentRequestAttributes{}
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *DORADeploymentRequestAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentRequestAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *DORADeploymentRequestAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *DORADeploymentRequestAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetFinishedAt returns the FinishedAt field value.
func (o *DORADeploymentRequestAttributes) GetFinishedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentRequestAttributes) GetFinishedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedAt, true
}

// SetFinishedAt sets field value.
func (o *DORADeploymentRequestAttributes) SetFinishedAt(v int64) {
	o.FinishedAt = v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *DORADeploymentRequestAttributes) GetGit() DORAGitInfo {
	if o == nil || o.Git == nil {
		var ret DORAGitInfo
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentRequestAttributes) GetGitOk() (*DORAGitInfo, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *DORADeploymentRequestAttributes) HasGit() bool {
	return o != nil && o.Git != nil
}

// SetGit gets a reference to the given DORAGitInfo and assigns it to the Git field.
func (o *DORADeploymentRequestAttributes) SetGit(v DORAGitInfo) {
	o.Git = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DORADeploymentRequestAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentRequestAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DORADeploymentRequestAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DORADeploymentRequestAttributes) SetId(v string) {
	o.Id = &v
}

// GetService returns the Service field value.
func (o *DORADeploymentRequestAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentRequestAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *DORADeploymentRequestAttributes) SetService(v string) {
	o.Service = v
}

// GetStartedAt returns the StartedAt field value.
func (o *DORADeploymentRequestAttributes) GetStartedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentRequestAttributes) GetStartedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *DORADeploymentRequestAttributes) SetStartedAt(v int64) {
	o.StartedAt = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DORADeploymentRequestAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentRequestAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DORADeploymentRequestAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DORADeploymentRequestAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORADeploymentRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	toSerialize["finished_at"] = o.FinishedAt
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["service"] = o.Service
	toSerialize["started_at"] = o.StartedAt
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORADeploymentRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Env        *string      `json:"env,omitempty"`
		FinishedAt *int64       `json:"finished_at"`
		Git        *DORAGitInfo `json:"git,omitempty"`
		Id         *string      `json:"id,omitempty"`
		Service    *string      `json:"service"`
		StartedAt  *int64       `json:"started_at"`
		Version    *string      `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FinishedAt == nil {
		return fmt.Errorf("required field finished_at missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.StartedAt == nil {
		return fmt.Errorf("required field started_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"env", "finished_at", "git", "id", "service", "started_at", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Env = all.Env
	o.FinishedAt = *all.FinishedAt
	if all.Git != nil && all.Git.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Git = all.Git
	o.Id = all.Id
	o.Service = *all.Service
	o.StartedAt = *all.StartedAt
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
