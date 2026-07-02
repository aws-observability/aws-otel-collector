// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabRunAttributes Attributes of a Model Lab run.
type ModelLabRunAttributes struct {
	// The date and time the run completed.
	CompletedAt datadog.NullableTime `json:"completed_at,omitempty"`
	// The date and time the run was created.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the run was soft-deleted.
	DeletedAt datadog.NullableTime `json:"deleted_at,omitempty"`
	// Whether a descendant run matched the applied filters.
	DescendantMatch bool `json:"descendant_match"`
	// A description of the run.
	Description string `json:"description"`
	// The duration of the run in seconds.
	Duration datadog.NullableFloat64 `json:"duration,omitempty"`
	// An optional external URL associated with the run.
	ExternalUrl datadog.NullableString `json:"external_url,omitempty"`
	// Whether the run has child runs.
	HasChildren bool `json:"has_children"`
	// Whether the run is pinned by the current user.
	IsPinned bool `json:"is_pinned"`
	// Summary statistics for metrics recorded during the run.
	MetricSummaries []ModelLabMetricSummary `json:"metric_summaries"`
	// The MLflow artifact storage location for this run.
	MlflowArtifactLocation string `json:"mlflow_artifact_location"`
	// The name of the run.
	Name string `json:"name"`
	// The UUID of the run owner.
	OwnerId datadog.NullableString `json:"owner_id,omitempty"`
	// The list of parameters used for the run.
	Params datadog.NullableList[ModelLabRunParam] `json:"params"`
	// The ID of the project this run belongs to.
	ProjectId int64 `json:"project_id"`
	// The date and time the run started.
	StartedAt time.Time `json:"started_at"`
	// The status of a Model Lab run.
	Status ModelLabRunStatus `json:"status"`
	// The list of tags associated with the run.
	Tags []ModelLabTag `json:"tags"`
	// The date and time the run was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabRunAttributes instantiates a new ModelLabRunAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabRunAttributes(createdAt time.Time, descendantMatch bool, description string, hasChildren bool, isPinned bool, metricSummaries []ModelLabMetricSummary, mlflowArtifactLocation string, name string, params datadog.NullableList[ModelLabRunParam], projectId int64, startedAt time.Time, status ModelLabRunStatus, tags []ModelLabTag, updatedAt time.Time) *ModelLabRunAttributes {
	this := ModelLabRunAttributes{}
	this.CreatedAt = createdAt
	this.DescendantMatch = descendantMatch
	this.Description = description
	this.HasChildren = hasChildren
	this.IsPinned = isPinned
	this.MetricSummaries = metricSummaries
	this.MlflowArtifactLocation = mlflowArtifactLocation
	this.Name = name
	this.Params = params
	this.ProjectId = projectId
	this.StartedAt = startedAt
	this.Status = status
	this.Tags = tags
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelLabRunAttributesWithDefaults instantiates a new ModelLabRunAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabRunAttributesWithDefaults() *ModelLabRunAttributes {
	this := ModelLabRunAttributes{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabRunAttributes) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabRunAttributes) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *ModelLabRunAttributes) HasCompletedAt() bool {
	return o != nil && o.CompletedAt.IsSet()
}

// SetCompletedAt gets a reference to the given datadog.NullableTime and assigns it to the CompletedAt field.
func (o *ModelLabRunAttributes) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}

// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil.
func (o *ModelLabRunAttributes) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil.
func (o *ModelLabRunAttributes) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ModelLabRunAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ModelLabRunAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabRunAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabRunAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ModelLabRunAttributes) HasDeletedAt() bool {
	return o != nil && o.DeletedAt.IsSet()
}

// SetDeletedAt gets a reference to the given datadog.NullableTime and assigns it to the DeletedAt field.
func (o *ModelLabRunAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil.
func (o *ModelLabRunAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil.
func (o *ModelLabRunAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetDescendantMatch returns the DescendantMatch field value.
func (o *ModelLabRunAttributes) GetDescendantMatch() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DescendantMatch
}

// GetDescendantMatchOk returns a tuple with the DescendantMatch field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetDescendantMatchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DescendantMatch, true
}

// SetDescendantMatch sets field value.
func (o *ModelLabRunAttributes) SetDescendantMatch(v bool) {
	o.DescendantMatch = v
}

// GetDescription returns the Description field value.
func (o *ModelLabRunAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ModelLabRunAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabRunAttributes) GetDuration() float64 {
	if o == nil || o.Duration.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabRunAttributes) GetDurationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *ModelLabRunAttributes) HasDuration() bool {
	return o != nil && o.Duration.IsSet()
}

// SetDuration gets a reference to the given datadog.NullableFloat64 and assigns it to the Duration field.
func (o *ModelLabRunAttributes) SetDuration(v float64) {
	o.Duration.Set(&v)
}

// SetDurationNil sets the value for Duration to be an explicit nil.
func (o *ModelLabRunAttributes) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil.
func (o *ModelLabRunAttributes) UnsetDuration() {
	o.Duration.Unset()
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabRunAttributes) GetExternalUrl() string {
	if o == nil || o.ExternalUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalUrl.Get()
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabRunAttributes) GetExternalUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalUrl.Get(), o.ExternalUrl.IsSet()
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *ModelLabRunAttributes) HasExternalUrl() bool {
	return o != nil && o.ExternalUrl.IsSet()
}

// SetExternalUrl gets a reference to the given datadog.NullableString and assigns it to the ExternalUrl field.
func (o *ModelLabRunAttributes) SetExternalUrl(v string) {
	o.ExternalUrl.Set(&v)
}

// SetExternalUrlNil sets the value for ExternalUrl to be an explicit nil.
func (o *ModelLabRunAttributes) SetExternalUrlNil() {
	o.ExternalUrl.Set(nil)
}

// UnsetExternalUrl ensures that no value is present for ExternalUrl, not even an explicit nil.
func (o *ModelLabRunAttributes) UnsetExternalUrl() {
	o.ExternalUrl.Unset()
}

// GetHasChildren returns the HasChildren field value.
func (o *ModelLabRunAttributes) GetHasChildren() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasChildren
}

// GetHasChildrenOk returns a tuple with the HasChildren field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetHasChildrenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasChildren, true
}

// SetHasChildren sets field value.
func (o *ModelLabRunAttributes) SetHasChildren(v bool) {
	o.HasChildren = v
}

// GetIsPinned returns the IsPinned field value.
func (o *ModelLabRunAttributes) GetIsPinned() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPinned
}

// GetIsPinnedOk returns a tuple with the IsPinned field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetIsPinnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPinned, true
}

// SetIsPinned sets field value.
func (o *ModelLabRunAttributes) SetIsPinned(v bool) {
	o.IsPinned = v
}

// GetMetricSummaries returns the MetricSummaries field value.
func (o *ModelLabRunAttributes) GetMetricSummaries() []ModelLabMetricSummary {
	if o == nil {
		var ret []ModelLabMetricSummary
		return ret
	}
	return o.MetricSummaries
}

// GetMetricSummariesOk returns a tuple with the MetricSummaries field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetMetricSummariesOk() (*[]ModelLabMetricSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricSummaries, true
}

// SetMetricSummaries sets field value.
func (o *ModelLabRunAttributes) SetMetricSummaries(v []ModelLabMetricSummary) {
	o.MetricSummaries = v
}

// GetMlflowArtifactLocation returns the MlflowArtifactLocation field value.
func (o *ModelLabRunAttributes) GetMlflowArtifactLocation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MlflowArtifactLocation
}

// GetMlflowArtifactLocationOk returns a tuple with the MlflowArtifactLocation field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetMlflowArtifactLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MlflowArtifactLocation, true
}

// SetMlflowArtifactLocation sets field value.
func (o *ModelLabRunAttributes) SetMlflowArtifactLocation(v string) {
	o.MlflowArtifactLocation = v
}

// GetName returns the Name field value.
func (o *ModelLabRunAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ModelLabRunAttributes) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabRunAttributes) GetOwnerId() string {
	if o == nil || o.OwnerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabRunAttributes) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ModelLabRunAttributes) HasOwnerId() bool {
	return o != nil && o.OwnerId.IsSet()
}

// SetOwnerId gets a reference to the given datadog.NullableString and assigns it to the OwnerId field.
func (o *ModelLabRunAttributes) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}

// SetOwnerIdNil sets the value for OwnerId to be an explicit nil.
func (o *ModelLabRunAttributes) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil.
func (o *ModelLabRunAttributes) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetParams returns the Params field value.
// If the value is explicit nil, the zero value for []ModelLabRunParam will be returned.
func (o *ModelLabRunAttributes) GetParams() []ModelLabRunParam {
	if o == nil {
		var ret []ModelLabRunParam
		return ret
	}
	return *o.Params.Get()
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabRunAttributes) GetParamsOk() (*[]ModelLabRunParam, bool) {
	if o == nil {
		return nil, false
	}
	return o.Params.Get(), o.Params.IsSet()
}

// SetParams sets field value.
func (o *ModelLabRunAttributes) SetParams(v []ModelLabRunParam) {
	o.Params.Set(&v)
}

// GetProjectId returns the ProjectId field value.
func (o *ModelLabRunAttributes) GetProjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetProjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *ModelLabRunAttributes) SetProjectId(v int64) {
	o.ProjectId = v
}

// GetStartedAt returns the StartedAt field value.
func (o *ModelLabRunAttributes) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *ModelLabRunAttributes) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStatus returns the Status field value.
func (o *ModelLabRunAttributes) GetStatus() ModelLabRunStatus {
	if o == nil {
		var ret ModelLabRunStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetStatusOk() (*ModelLabRunStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ModelLabRunAttributes) SetStatus(v ModelLabRunStatus) {
	o.Status = v
}

// GetTags returns the Tags field value.
func (o *ModelLabRunAttributes) GetTags() []ModelLabTag {
	if o == nil {
		var ret []ModelLabTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetTagsOk() (*[]ModelLabTag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *ModelLabRunAttributes) SetTags(v []ModelLabTag) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ModelLabRunAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ModelLabRunAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabRunAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	toSerialize["descendant_match"] = o.DescendantMatch
	toSerialize["description"] = o.Description
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.ExternalUrl.IsSet() {
		toSerialize["external_url"] = o.ExternalUrl.Get()
	}
	toSerialize["has_children"] = o.HasChildren
	toSerialize["is_pinned"] = o.IsPinned
	toSerialize["metric_summaries"] = o.MetricSummaries
	toSerialize["mlflow_artifact_location"] = o.MlflowArtifactLocation
	toSerialize["name"] = o.Name
	if o.OwnerId.IsSet() {
		toSerialize["owner_id"] = o.OwnerId.Get()
	}
	toSerialize["params"] = o.Params.Get()
	toSerialize["project_id"] = o.ProjectId
	if o.StartedAt.Nanosecond() == 0 {
		toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status
	toSerialize["tags"] = o.Tags
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabRunAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletedAt            datadog.NullableTime                   `json:"completed_at,omitempty"`
		CreatedAt              *time.Time                             `json:"created_at"`
		DeletedAt              datadog.NullableTime                   `json:"deleted_at,omitempty"`
		DescendantMatch        *bool                                  `json:"descendant_match"`
		Description            *string                                `json:"description"`
		Duration               datadog.NullableFloat64                `json:"duration,omitempty"`
		ExternalUrl            datadog.NullableString                 `json:"external_url,omitempty"`
		HasChildren            *bool                                  `json:"has_children"`
		IsPinned               *bool                                  `json:"is_pinned"`
		MetricSummaries        *[]ModelLabMetricSummary               `json:"metric_summaries"`
		MlflowArtifactLocation *string                                `json:"mlflow_artifact_location"`
		Name                   *string                                `json:"name"`
		OwnerId                datadog.NullableString                 `json:"owner_id,omitempty"`
		Params                 datadog.NullableList[ModelLabRunParam] `json:"params"`
		ProjectId              *int64                                 `json:"project_id"`
		StartedAt              *time.Time                             `json:"started_at"`
		Status                 *ModelLabRunStatus                     `json:"status"`
		Tags                   *[]ModelLabTag                         `json:"tags"`
		UpdatedAt              *time.Time                             `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DescendantMatch == nil {
		return fmt.Errorf("required field descendant_match missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.HasChildren == nil {
		return fmt.Errorf("required field has_children missing")
	}
	if all.IsPinned == nil {
		return fmt.Errorf("required field is_pinned missing")
	}
	if all.MetricSummaries == nil {
		return fmt.Errorf("required field metric_summaries missing")
	}
	if all.MlflowArtifactLocation == nil {
		return fmt.Errorf("required field mlflow_artifact_location missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if !all.Params.IsSet() {
		return fmt.Errorf("required field params missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.StartedAt == nil {
		return fmt.Errorf("required field started_at missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed_at", "created_at", "deleted_at", "descendant_match", "description", "duration", "external_url", "has_children", "is_pinned", "metric_summaries", "mlflow_artifact_location", "name", "owner_id", "params", "project_id", "started_at", "status", "tags", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CompletedAt = all.CompletedAt
	o.CreatedAt = *all.CreatedAt
	o.DeletedAt = all.DeletedAt
	o.DescendantMatch = *all.DescendantMatch
	o.Description = *all.Description
	o.Duration = all.Duration
	o.ExternalUrl = all.ExternalUrl
	o.HasChildren = *all.HasChildren
	o.IsPinned = *all.IsPinned
	o.MetricSummaries = *all.MetricSummaries
	o.MlflowArtifactLocation = *all.MlflowArtifactLocation
	o.Name = *all.Name
	o.OwnerId = all.OwnerId
	o.Params = all.Params
	o.ProjectId = *all.ProjectId
	o.StartedAt = *all.StartedAt
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Tags = *all.Tags
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
