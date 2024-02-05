// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORAIncidentRequestAttributes Attributes to create a DORA incident event.
type DORAIncidentRequestAttributes struct {
	// Unix timestamp in nanoseconds when the incident finished. It should not be older than 3 hours.
	FinishedAt *int64 `json:"finished_at,omitempty"`
	// Git info for DORA Metrics events.
	Git *DORAGitInfo `json:"git,omitempty"`
	// Incident ID
	Id *string `json:"id,omitempty"`
	// Incident name.
	Name *string `json:"name,omitempty"`
	// Service name from a service available in the Service Catalog.
	Service string `json:"service"`
	// Incident severity.
	Severity *string `json:"severity,omitempty"`
	// Unix timestamp in nanoseconds when the incident started.
	StartedAt int64 `json:"started_at"`
	// Version to correlate with [APM Deployment Tracking](https://docs.datadoghq.com/tracing/services/deployment_tracking/).
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDORAIncidentRequestAttributes instantiates a new DORAIncidentRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORAIncidentRequestAttributes(service string, startedAt int64) *DORAIncidentRequestAttributes {
	this := DORAIncidentRequestAttributes{}
	this.Service = service
	this.StartedAt = startedAt
	return &this
}

// NewDORAIncidentRequestAttributesWithDefaults instantiates a new DORAIncidentRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORAIncidentRequestAttributesWithDefaults() *DORAIncidentRequestAttributes {
	this := DORAIncidentRequestAttributes{}
	return &this
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DORAIncidentRequestAttributes) GetFinishedAt() int64 {
	if o == nil || o.FinishedAt == nil {
		var ret int64
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentRequestAttributes) GetFinishedAtOk() (*int64, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DORAIncidentRequestAttributes) HasFinishedAt() bool {
	return o != nil && o.FinishedAt != nil
}

// SetFinishedAt gets a reference to the given int64 and assigns it to the FinishedAt field.
func (o *DORAIncidentRequestAttributes) SetFinishedAt(v int64) {
	o.FinishedAt = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *DORAIncidentRequestAttributes) GetGit() DORAGitInfo {
	if o == nil || o.Git == nil {
		var ret DORAGitInfo
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentRequestAttributes) GetGitOk() (*DORAGitInfo, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *DORAIncidentRequestAttributes) HasGit() bool {
	return o != nil && o.Git != nil
}

// SetGit gets a reference to the given DORAGitInfo and assigns it to the Git field.
func (o *DORAIncidentRequestAttributes) SetGit(v DORAGitInfo) {
	o.Git = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DORAIncidentRequestAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentRequestAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DORAIncidentRequestAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DORAIncidentRequestAttributes) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DORAIncidentRequestAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DORAIncidentRequestAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DORAIncidentRequestAttributes) SetName(v string) {
	o.Name = &v
}

// GetService returns the Service field value.
func (o *DORAIncidentRequestAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *DORAIncidentRequestAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *DORAIncidentRequestAttributes) SetService(v string) {
	o.Service = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *DORAIncidentRequestAttributes) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentRequestAttributes) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *DORAIncidentRequestAttributes) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *DORAIncidentRequestAttributes) SetSeverity(v string) {
	o.Severity = &v
}

// GetStartedAt returns the StartedAt field value.
func (o *DORAIncidentRequestAttributes) GetStartedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *DORAIncidentRequestAttributes) GetStartedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *DORAIncidentRequestAttributes) SetStartedAt(v int64) {
	o.StartedAt = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DORAIncidentRequestAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAIncidentRequestAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DORAIncidentRequestAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DORAIncidentRequestAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORAIncidentRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FinishedAt != nil {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["service"] = o.Service
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
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
func (o *DORAIncidentRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FinishedAt *int64       `json:"finished_at,omitempty"`
		Git        *DORAGitInfo `json:"git,omitempty"`
		Id         *string      `json:"id,omitempty"`
		Name       *string      `json:"name,omitempty"`
		Service    *string      `json:"service"`
		Severity   *string      `json:"severity,omitempty"`
		StartedAt  *int64       `json:"started_at"`
		Version    *string      `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.StartedAt == nil {
		return fmt.Errorf("required field started_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"finished_at", "git", "id", "name", "service", "severity", "started_at", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FinishedAt = all.FinishedAt
	if all.Git != nil && all.Git.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Git = all.Git
	o.Id = all.Id
	o.Name = all.Name
	o.Service = *all.Service
	o.Severity = all.Severity
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
