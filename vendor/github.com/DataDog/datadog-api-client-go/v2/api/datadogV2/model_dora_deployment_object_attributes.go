// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentObjectAttributes The attributes of the deployment event.
type DORADeploymentObjectAttributes struct {
	// A list of user-defined tags. The tags must follow the `key:value` pattern. Up to 100 may be added per event.
	CustomTags datadog.NullableList[string] `json:"custom_tags,omitempty"`
	// Environment name to where the service was deployed.
	Env *string `json:"env,omitempty"`
	// Unix timestamp when the deployment finished.
	FinishedAt int64 `json:"finished_at"`
	// Git info returned by DORA Metrics events.
	Git *DORAGitInfoResponse `json:"git,omitempty"`
	// Service name.
	Service string `json:"service"`
	// Unix timestamp when the deployment started.
	StartedAt int64 `json:"started_at"`
	// Name of the team owning the deployed service.
	Team *string `json:"team,omitempty"`
	// Version to correlate with APM Deployment Tracking.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORADeploymentObjectAttributes instantiates a new DORADeploymentObjectAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORADeploymentObjectAttributes(finishedAt int64, service string, startedAt int64) *DORADeploymentObjectAttributes {
	this := DORADeploymentObjectAttributes{}
	this.FinishedAt = finishedAt
	this.Service = service
	this.StartedAt = startedAt
	return &this
}

// NewDORADeploymentObjectAttributesWithDefaults instantiates a new DORADeploymentObjectAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORADeploymentObjectAttributesWithDefaults() *DORADeploymentObjectAttributes {
	this := DORADeploymentObjectAttributes{}
	return &this
}

// GetCustomTags returns the CustomTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DORADeploymentObjectAttributes) GetCustomTags() []string {
	if o == nil || o.CustomTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.CustomTags.Get()
}

// GetCustomTagsOk returns a tuple with the CustomTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DORADeploymentObjectAttributes) GetCustomTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomTags.Get(), o.CustomTags.IsSet()
}

// HasCustomTags returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasCustomTags() bool {
	return o != nil && o.CustomTags.IsSet()
}

// SetCustomTags gets a reference to the given datadog.NullableList[string] and assigns it to the CustomTags field.
func (o *DORADeploymentObjectAttributes) SetCustomTags(v []string) {
	o.CustomTags.Set(&v)
}

// SetCustomTagsNil sets the value for CustomTags to be an explicit nil.
func (o *DORADeploymentObjectAttributes) SetCustomTagsNil() {
	o.CustomTags.Set(nil)
}

// UnsetCustomTags ensures that no value is present for CustomTags, not even an explicit nil.
func (o *DORADeploymentObjectAttributes) UnsetCustomTags() {
	o.CustomTags.Unset()
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *DORADeploymentObjectAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetFinishedAt returns the FinishedAt field value.
func (o *DORADeploymentObjectAttributes) GetFinishedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetFinishedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedAt, true
}

// SetFinishedAt sets field value.
func (o *DORADeploymentObjectAttributes) SetFinishedAt(v int64) {
	o.FinishedAt = v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetGit() DORAGitInfoResponse {
	if o == nil || o.Git == nil {
		var ret DORAGitInfoResponse
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetGitOk() (*DORAGitInfoResponse, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasGit() bool {
	return o != nil && o.Git != nil
}

// SetGit gets a reference to the given DORAGitInfoResponse and assigns it to the Git field.
func (o *DORADeploymentObjectAttributes) SetGit(v DORAGitInfoResponse) {
	o.Git = &v
}

// GetService returns the Service field value.
func (o *DORADeploymentObjectAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *DORADeploymentObjectAttributes) SetService(v string) {
	o.Service = v
}

// GetStartedAt returns the StartedAt field value.
func (o *DORADeploymentObjectAttributes) GetStartedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetStartedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *DORADeploymentObjectAttributes) SetStartedAt(v int64) {
	o.StartedAt = v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetTeam() string {
	if o == nil || o.Team == nil {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetTeamOk() (*string, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *DORADeploymentObjectAttributes) SetTeam(v string) {
	o.Team = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DORADeploymentObjectAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObjectAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DORADeploymentObjectAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DORADeploymentObjectAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORADeploymentObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomTags.IsSet() {
		toSerialize["custom_tags"] = o.CustomTags.Get()
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	toSerialize["finished_at"] = o.FinishedAt
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	toSerialize["service"] = o.Service
	toSerialize["started_at"] = o.StartedAt
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORADeploymentObjectAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomTags datadog.NullableList[string] `json:"custom_tags,omitempty"`
		Env        *string                      `json:"env,omitempty"`
		FinishedAt *int64                       `json:"finished_at"`
		Git        *DORAGitInfoResponse         `json:"git,omitempty"`
		Service    *string                      `json:"service"`
		StartedAt  *int64                       `json:"started_at"`
		Team       *string                      `json:"team,omitempty"`
		Version    *string                      `json:"version,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_tags", "env", "finished_at", "git", "service", "started_at", "team", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CustomTags = all.CustomTags
	o.Env = all.Env
	o.FinishedAt = *all.FinishedAt
	if all.Git != nil && all.Git.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Git = all.Git
	o.Service = *all.Service
	o.StartedAt = *all.StartedAt
	o.Team = all.Team
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
