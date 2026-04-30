// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORAIncidentObjectAttributes The attributes of the incident event.
type DORAIncidentObjectAttributes struct {
	// A list of user-defined tags. The tags must follow the `key:value` pattern. Up to 100 may be added per event.
	CustomTags datadog.NullableList[string] `json:"custom_tags,omitempty"`
	// Environment name that was impacted by the incident.
	Env *string `json:"env,omitempty"`
	// Unix timestamp when the incident finished.
	FinishedAt *int64 `json:"finished_at,omitempty"`
	// Git info for DORA Metrics events.
	Git *DORAGitInfo `json:"git,omitempty"`
	// Incident name.
	Name *string `json:"name,omitempty"`
	// Service names impacted by the incident.
	Services []string `json:"services,omitempty"`
	// Incident severity.
	Severity *string `json:"severity,omitempty"`
	// Unix timestamp when the incident started.
	StartedAt int64 `json:"started_at"`
	// Name of the team owning the services impacted.
	Team *string `json:"team,omitempty"`
	// Version to correlate with APM Deployment Tracking.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORAIncidentObjectAttributes instantiates a new DORAIncidentObjectAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORAIncidentObjectAttributes(startedAt int64) *DORAIncidentObjectAttributes {
	this := DORAIncidentObjectAttributes{}
	this.StartedAt = startedAt
	return &this
}

// NewDORAIncidentObjectAttributesWithDefaults instantiates a new DORAIncidentObjectAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORAIncidentObjectAttributesWithDefaults() *DORAIncidentObjectAttributes {
	this := DORAIncidentObjectAttributes{}
	return &this
}

// GetCustomTags returns the CustomTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DORAIncidentObjectAttributes) GetCustomTags() []string {
	if o == nil || o.CustomTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.CustomTags.Get()
}

// GetCustomTagsOk returns a tuple with the CustomTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DORAIncidentObjectAttributes) GetCustomTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomTags.Get(), o.CustomTags.IsSet()
}

// HasCustomTags returns a boolean if a field has been set.
func (o *DORAIncidentObjectAttributes) HasCustomTags() bool {
	return o != nil && o.CustomTags.IsSet()
}

// SetCustomTags gets a reference to the given datadog.NullableList[string] and assigns it to the CustomTags field.
func (o *DORAIncidentObjectAttributes) SetCustomTags(v []string) {
	o.CustomTags.Set(&v)
}

// SetCustomTagsNil sets the value for CustomTags to be an explicit nil.
func (o *DORAIncidentObjectAttributes) SetCustomTagsNil() {
	o.CustomTags.Set(nil)
}

// UnsetCustomTags ensures that no value is present for CustomTags, not even an explicit nil.
func (o *DORAIncidentObjectAttributes) UnsetCustomTags() {
	o.CustomTags.Unset()
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *DORAIncidentObjectAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentObjectAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *DORAIncidentObjectAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *DORAIncidentObjectAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DORAIncidentObjectAttributes) GetFinishedAt() int64 {
	if o == nil || o.FinishedAt == nil {
		var ret int64
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentObjectAttributes) GetFinishedAtOk() (*int64, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DORAIncidentObjectAttributes) HasFinishedAt() bool {
	return o != nil && o.FinishedAt != nil
}

// SetFinishedAt gets a reference to the given int64 and assigns it to the FinishedAt field.
func (o *DORAIncidentObjectAttributes) SetFinishedAt(v int64) {
	o.FinishedAt = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *DORAIncidentObjectAttributes) GetGit() DORAGitInfo {
	if o == nil || o.Git == nil {
		var ret DORAGitInfo
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentObjectAttributes) GetGitOk() (*DORAGitInfo, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *DORAIncidentObjectAttributes) HasGit() bool {
	return o != nil && o.Git != nil
}

// SetGit gets a reference to the given DORAGitInfo and assigns it to the Git field.
func (o *DORAIncidentObjectAttributes) SetGit(v DORAGitInfo) {
	o.Git = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DORAIncidentObjectAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentObjectAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DORAIncidentObjectAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DORAIncidentObjectAttributes) SetName(v string) {
	o.Name = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *DORAIncidentObjectAttributes) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentObjectAttributes) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *DORAIncidentObjectAttributes) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *DORAIncidentObjectAttributes) SetServices(v []string) {
	o.Services = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *DORAIncidentObjectAttributes) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentObjectAttributes) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *DORAIncidentObjectAttributes) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *DORAIncidentObjectAttributes) SetSeverity(v string) {
	o.Severity = &v
}

// GetStartedAt returns the StartedAt field value.
func (o *DORAIncidentObjectAttributes) GetStartedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *DORAIncidentObjectAttributes) GetStartedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *DORAIncidentObjectAttributes) SetStartedAt(v int64) {
	o.StartedAt = v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *DORAIncidentObjectAttributes) GetTeam() string {
	if o == nil || o.Team == nil {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentObjectAttributes) GetTeamOk() (*string, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *DORAIncidentObjectAttributes) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *DORAIncidentObjectAttributes) SetTeam(v string) {
	o.Team = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DORAIncidentObjectAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentObjectAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DORAIncidentObjectAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DORAIncidentObjectAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORAIncidentObjectAttributes) MarshalJSON() ([]byte, error) {
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
	if o.FinishedAt != nil {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
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
func (o *DORAIncidentObjectAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomTags datadog.NullableList[string] `json:"custom_tags,omitempty"`
		Env        *string                      `json:"env,omitempty"`
		FinishedAt *int64                       `json:"finished_at,omitempty"`
		Git        *DORAGitInfo                 `json:"git,omitempty"`
		Name       *string                      `json:"name,omitempty"`
		Services   []string                     `json:"services,omitempty"`
		Severity   *string                      `json:"severity,omitempty"`
		StartedAt  *int64                       `json:"started_at"`
		Team       *string                      `json:"team,omitempty"`
		Version    *string                      `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.StartedAt == nil {
		return fmt.Errorf("required field started_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_tags", "env", "finished_at", "git", "name", "services", "severity", "started_at", "team", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CustomTags = all.CustomTags
	o.Env = all.Env
	o.FinishedAt = all.FinishedAt
	if all.Git != nil && all.Git.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Git = all.Git
	o.Name = all.Name
	o.Services = all.Services
	o.Severity = all.Severity
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
