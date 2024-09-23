// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventStep Details of a CI step.
type CIAppPipelineEventStep struct {
	// Time when the step run finished. The time format must be RFC3339.
	End time.Time `json:"end"`
	// Contains information of the CI error.
	Error NullableCIAppCIError `json:"error,omitempty"`
	// If pipelines are triggered due to actions to a Git repository, then all payloads must contain this.
	// Note that either `tag` or `branch` has to be provided, but not both.
	Git NullableCIAppGitInfo `json:"git,omitempty"`
	// UUID for the step. It has to be unique within each pipeline execution.
	Id string `json:"id"`
	// The parent job UUID (if applicable).
	JobId datadog.NullableString `json:"job_id,omitempty"`
	// The parent job name (if applicable).
	JobName datadog.NullableString `json:"job_name,omitempty"`
	// Used to distinguish between pipelines, stages, jobs and steps.
	Level CIAppPipelineEventStepLevel `json:"level"`
	// A list of user-defined metrics. The metrics must follow the `key:value` pattern and the value must be numeric.
	Metrics datadog.NullableList[string] `json:"metrics,omitempty"`
	// The name for the step.
	Name string `json:"name"`
	// Contains information of the host running the pipeline, stage, job, or step.
	Node NullableCIAppHostInfo `json:"node,omitempty"`
	// A map of key-value parameters or environment variables that were defined for the pipeline.
	Parameters map[string]string `json:"parameters,omitempty"`
	// The parent pipeline name.
	PipelineName string `json:"pipeline_name"`
	// The parent pipeline UUID.
	PipelineUniqueId string `json:"pipeline_unique_id"`
	// The parent stage UUID (if applicable).
	StageId datadog.NullableString `json:"stage_id,omitempty"`
	// The parent stage name (if applicable).
	StageName datadog.NullableString `json:"stage_name,omitempty"`
	// Time when the step run started. The time format must be RFC3339.
	Start time.Time `json:"start"`
	// The final status of the step.
	Status CIAppPipelineEventStepStatus `json:"status"`
	// A list of user-defined tags. The tags must follow the `key:value` pattern.
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// The URL to look at the step in the CI provider UI.
	Url datadog.NullableString `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppPipelineEventStep instantiates a new CIAppPipelineEventStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppPipelineEventStep(end time.Time, id string, level CIAppPipelineEventStepLevel, name string, pipelineName string, pipelineUniqueId string, start time.Time, status CIAppPipelineEventStepStatus) *CIAppPipelineEventStep {
	this := CIAppPipelineEventStep{}
	this.End = end
	this.Id = id
	this.Level = level
	this.Name = name
	this.PipelineName = pipelineName
	this.PipelineUniqueId = pipelineUniqueId
	this.Start = start
	this.Status = status
	return &this
}

// NewCIAppPipelineEventStepWithDefaults instantiates a new CIAppPipelineEventStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppPipelineEventStepWithDefaults() *CIAppPipelineEventStep {
	this := CIAppPipelineEventStep{}
	var level CIAppPipelineEventStepLevel = CIAPPPIPELINEEVENTSTEPLEVEL_STEP
	this.Level = level
	return &this
}

// GetEnd returns the End field value.
func (o *CIAppPipelineEventStep) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStep) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *CIAppPipelineEventStep) SetEnd(v time.Time) {
	o.End = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetError() CIAppCIError {
	if o == nil || o.Error.Get() == nil {
		var ret CIAppCIError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetErrorOk() (*CIAppCIError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasError() bool {
	return o != nil && o.Error.IsSet()
}

// SetError gets a reference to the given NullableCIAppCIError and assigns it to the Error field.
func (o *CIAppPipelineEventStep) SetError(v CIAppCIError) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil.
func (o *CIAppPipelineEventStep) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil.
func (o *CIAppPipelineEventStep) UnsetError() {
	o.Error.Unset()
}

// GetGit returns the Git field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetGit() CIAppGitInfo {
	if o == nil || o.Git.Get() == nil {
		var ret CIAppGitInfo
		return ret
	}
	return *o.Git.Get()
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetGitOk() (*CIAppGitInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Git.Get(), o.Git.IsSet()
}

// HasGit returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasGit() bool {
	return o != nil && o.Git.IsSet()
}

// SetGit gets a reference to the given NullableCIAppGitInfo and assigns it to the Git field.
func (o *CIAppPipelineEventStep) SetGit(v CIAppGitInfo) {
	o.Git.Set(&v)
}

// SetGitNil sets the value for Git to be an explicit nil.
func (o *CIAppPipelineEventStep) SetGitNil() {
	o.Git.Set(nil)
}

// UnsetGit ensures that no value is present for Git, not even an explicit nil.
func (o *CIAppPipelineEventStep) UnsetGit() {
	o.Git.Unset()
}

// GetId returns the Id field value.
func (o *CIAppPipelineEventStep) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStep) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CIAppPipelineEventStep) SetId(v string) {
	o.Id = v
}

// GetJobId returns the JobId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetJobId() string {
	if o == nil || o.JobId.Get() == nil {
		var ret string
		return ret
	}
	return *o.JobId.Get()
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobId.Get(), o.JobId.IsSet()
}

// HasJobId returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasJobId() bool {
	return o != nil && o.JobId.IsSet()
}

// SetJobId gets a reference to the given datadog.NullableString and assigns it to the JobId field.
func (o *CIAppPipelineEventStep) SetJobId(v string) {
	o.JobId.Set(&v)
}

// SetJobIdNil sets the value for JobId to be an explicit nil.
func (o *CIAppPipelineEventStep) SetJobIdNil() {
	o.JobId.Set(nil)
}

// UnsetJobId ensures that no value is present for JobId, not even an explicit nil.
func (o *CIAppPipelineEventStep) UnsetJobId() {
	o.JobId.Unset()
}

// GetJobName returns the JobName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetJobName() string {
	if o == nil || o.JobName.Get() == nil {
		var ret string
		return ret
	}
	return *o.JobName.Get()
}

// GetJobNameOk returns a tuple with the JobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetJobNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobName.Get(), o.JobName.IsSet()
}

// HasJobName returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasJobName() bool {
	return o != nil && o.JobName.IsSet()
}

// SetJobName gets a reference to the given datadog.NullableString and assigns it to the JobName field.
func (o *CIAppPipelineEventStep) SetJobName(v string) {
	o.JobName.Set(&v)
}

// SetJobNameNil sets the value for JobName to be an explicit nil.
func (o *CIAppPipelineEventStep) SetJobNameNil() {
	o.JobName.Set(nil)
}

// UnsetJobName ensures that no value is present for JobName, not even an explicit nil.
func (o *CIAppPipelineEventStep) UnsetJobName() {
	o.JobName.Unset()
}

// GetLevel returns the Level field value.
func (o *CIAppPipelineEventStep) GetLevel() CIAppPipelineEventStepLevel {
	if o == nil {
		var ret CIAppPipelineEventStepLevel
		return ret
	}
	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStep) GetLevelOk() (*CIAppPipelineEventStepLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value.
func (o *CIAppPipelineEventStep) SetLevel(v CIAppPipelineEventStepLevel) {
	o.Level = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetMetrics() []string {
	if o == nil || o.Metrics.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Metrics.Get()
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetMetricsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics.Get(), o.Metrics.IsSet()
}

// HasMetrics returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasMetrics() bool {
	return o != nil && o.Metrics.IsSet()
}

// SetMetrics gets a reference to the given datadog.NullableList[string] and assigns it to the Metrics field.
func (o *CIAppPipelineEventStep) SetMetrics(v []string) {
	o.Metrics.Set(&v)
}

// SetMetricsNil sets the value for Metrics to be an explicit nil.
func (o *CIAppPipelineEventStep) SetMetricsNil() {
	o.Metrics.Set(nil)
}

// UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil.
func (o *CIAppPipelineEventStep) UnsetMetrics() {
	o.Metrics.Unset()
}

// GetName returns the Name field value.
func (o *CIAppPipelineEventStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CIAppPipelineEventStep) SetName(v string) {
	o.Name = v
}

// GetNode returns the Node field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetNode() CIAppHostInfo {
	if o == nil || o.Node.Get() == nil {
		var ret CIAppHostInfo
		return ret
	}
	return *o.Node.Get()
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetNodeOk() (*CIAppHostInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Node.Get(), o.Node.IsSet()
}

// HasNode returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasNode() bool {
	return o != nil && o.Node.IsSet()
}

// SetNode gets a reference to the given NullableCIAppHostInfo and assigns it to the Node field.
func (o *CIAppPipelineEventStep) SetNode(v CIAppHostInfo) {
	o.Node.Set(&v)
}

// SetNodeNil sets the value for Node to be an explicit nil.
func (o *CIAppPipelineEventStep) SetNodeNil() {
	o.Node.Set(nil)
}

// UnsetNode ensures that no value is present for Node, not even an explicit nil.
func (o *CIAppPipelineEventStep) UnsetNode() {
	o.Node.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *CIAppPipelineEventStep) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetPipelineName returns the PipelineName field value.
func (o *CIAppPipelineEventStep) GetPipelineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PipelineName
}

// GetPipelineNameOk returns a tuple with the PipelineName field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStep) GetPipelineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PipelineName, true
}

// SetPipelineName sets field value.
func (o *CIAppPipelineEventStep) SetPipelineName(v string) {
	o.PipelineName = v
}

// GetPipelineUniqueId returns the PipelineUniqueId field value.
func (o *CIAppPipelineEventStep) GetPipelineUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PipelineUniqueId
}

// GetPipelineUniqueIdOk returns a tuple with the PipelineUniqueId field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStep) GetPipelineUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PipelineUniqueId, true
}

// SetPipelineUniqueId sets field value.
func (o *CIAppPipelineEventStep) SetPipelineUniqueId(v string) {
	o.PipelineUniqueId = v
}

// GetStageId returns the StageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetStageId() string {
	if o == nil || o.StageId.Get() == nil {
		var ret string
		return ret
	}
	return *o.StageId.Get()
}

// GetStageIdOk returns a tuple with the StageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetStageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StageId.Get(), o.StageId.IsSet()
}

// HasStageId returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasStageId() bool {
	return o != nil && o.StageId.IsSet()
}

// SetStageId gets a reference to the given datadog.NullableString and assigns it to the StageId field.
func (o *CIAppPipelineEventStep) SetStageId(v string) {
	o.StageId.Set(&v)
}

// SetStageIdNil sets the value for StageId to be an explicit nil.
func (o *CIAppPipelineEventStep) SetStageIdNil() {
	o.StageId.Set(nil)
}

// UnsetStageId ensures that no value is present for StageId, not even an explicit nil.
func (o *CIAppPipelineEventStep) UnsetStageId() {
	o.StageId.Unset()
}

// GetStageName returns the StageName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetStageName() string {
	if o == nil || o.StageName.Get() == nil {
		var ret string
		return ret
	}
	return *o.StageName.Get()
}

// GetStageNameOk returns a tuple with the StageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetStageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StageName.Get(), o.StageName.IsSet()
}

// HasStageName returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasStageName() bool {
	return o != nil && o.StageName.IsSet()
}

// SetStageName gets a reference to the given datadog.NullableString and assigns it to the StageName field.
func (o *CIAppPipelineEventStep) SetStageName(v string) {
	o.StageName.Set(&v)
}

// SetStageNameNil sets the value for StageName to be an explicit nil.
func (o *CIAppPipelineEventStep) SetStageNameNil() {
	o.StageName.Set(nil)
}

// UnsetStageName ensures that no value is present for StageName, not even an explicit nil.
func (o *CIAppPipelineEventStep) UnsetStageName() {
	o.StageName.Unset()
}

// GetStart returns the Start field value.
func (o *CIAppPipelineEventStep) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStep) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *CIAppPipelineEventStep) SetStart(v time.Time) {
	o.Start = v
}

// GetStatus returns the Status field value.
func (o *CIAppPipelineEventStep) GetStatus() CIAppPipelineEventStepStatus {
	if o == nil {
		var ret CIAppPipelineEventStepStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventStep) GetStatusOk() (*CIAppPipelineEventStepStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CIAppPipelineEventStep) SetStatus(v CIAppPipelineEventStepStatus) {
	o.Status = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *CIAppPipelineEventStep) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *CIAppPipelineEventStep) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *CIAppPipelineEventStep) UnsetTags() {
	o.Tags.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppPipelineEventStep) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppPipelineEventStep) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *CIAppPipelineEventStep) HasUrl() bool {
	return o != nil && o.Url.IsSet()
}

// SetUrl gets a reference to the given datadog.NullableString and assigns it to the Url field.
func (o *CIAppPipelineEventStep) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil.
func (o *CIAppPipelineEventStep) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil.
func (o *CIAppPipelineEventStep) UnsetUrl() {
	o.Url.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppPipelineEventStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End.Nanosecond() == 0 {
		toSerialize["end"] = o.End.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["end"] = o.End.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.Git.IsSet() {
		toSerialize["git"] = o.Git.Get()
	}
	toSerialize["id"] = o.Id
	if o.JobId.IsSet() {
		toSerialize["job_id"] = o.JobId.Get()
	}
	if o.JobName.IsSet() {
		toSerialize["job_name"] = o.JobName.Get()
	}
	toSerialize["level"] = o.Level
	if o.Metrics.IsSet() {
		toSerialize["metrics"] = o.Metrics.Get()
	}
	toSerialize["name"] = o.Name
	if o.Node.IsSet() {
		toSerialize["node"] = o.Node.Get()
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	toSerialize["pipeline_name"] = o.PipelineName
	toSerialize["pipeline_unique_id"] = o.PipelineUniqueId
	if o.StageId.IsSet() {
		toSerialize["stage_id"] = o.StageId.Get()
	}
	if o.StageName.IsSet() {
		toSerialize["stage_name"] = o.StageName.Get()
	}
	if o.Start.Nanosecond() == 0 {
		toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppPipelineEventStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End              *time.Time                    `json:"end"`
		Error            NullableCIAppCIError          `json:"error,omitempty"`
		Git              NullableCIAppGitInfo          `json:"git,omitempty"`
		Id               *string                       `json:"id"`
		JobId            datadog.NullableString        `json:"job_id,omitempty"`
		JobName          datadog.NullableString        `json:"job_name,omitempty"`
		Level            *CIAppPipelineEventStepLevel  `json:"level"`
		Metrics          datadog.NullableList[string]  `json:"metrics,omitempty"`
		Name             *string                       `json:"name"`
		Node             NullableCIAppHostInfo         `json:"node,omitempty"`
		Parameters       map[string]string             `json:"parameters,omitempty"`
		PipelineName     *string                       `json:"pipeline_name"`
		PipelineUniqueId *string                       `json:"pipeline_unique_id"`
		StageId          datadog.NullableString        `json:"stage_id,omitempty"`
		StageName        datadog.NullableString        `json:"stage_name,omitempty"`
		Start            *time.Time                    `json:"start"`
		Status           *CIAppPipelineEventStepStatus `json:"status"`
		Tags             datadog.NullableList[string]  `json:"tags,omitempty"`
		Url              datadog.NullableString        `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Level == nil {
		return fmt.Errorf("required field level missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.PipelineName == nil {
		return fmt.Errorf("required field pipeline_name missing")
	}
	if all.PipelineUniqueId == nil {
		return fmt.Errorf("required field pipeline_unique_id missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "error", "git", "id", "job_id", "job_name", "level", "metrics", "name", "node", "parameters", "pipeline_name", "pipeline_unique_id", "stage_id", "stage_name", "start", "status", "tags", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	o.End = *all.End
	o.Error = all.Error
	o.Git = all.Git
	o.Id = *all.Id
	o.JobId = all.JobId
	o.JobName = all.JobName
	if !all.Level.IsValid() {
		hasInvalidField = true
	} else {
		o.Level = *all.Level
	}
	o.Metrics = all.Metrics
	o.Name = *all.Name
	o.Node = all.Node
	o.Parameters = all.Parameters
	o.PipelineName = *all.PipelineName
	o.PipelineUniqueId = *all.PipelineUniqueId
	o.StageId = all.StageId
	o.StageName = all.StageName
	o.Start = *all.Start
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Tags = all.Tags
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
