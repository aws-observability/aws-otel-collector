// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentDataAttributesResponse Attributes of an LLM Observability experiment.
type LLMObsExperimentDataAttributesResponse struct {
	// Pre-computed aggregate metrics for this experiment run, including eval score distributions, token costs, and error rates.
	AggregateData map[string]interface{} `json:"aggregate_data,omitempty"`
	// User data for the author of an experiment. Only present when `include[user_data]` is `true`.
	Author *LLMObsExperimentUser `json:"author,omitempty"`
	// Configuration parameters for the experiment.
	Config map[string]interface{} `json:"config"`
	// Timestamp when the experiment was created.
	CreatedAt time.Time `json:"created_at"`
	// Identifier of the dataset used in this experiment.
	DatasetId string `json:"dataset_id"`
	// Name of the dataset used in this experiment.
	// Only present when `include[dataset_names]` is `true`.
	DatasetName datadog.NullableString `json:"dataset_name,omitempty"`
	// Version of the dataset used in this experiment.
	DatasetVersion *int64 `json:"dataset_version,omitempty"`
	// Timestamp when the experiment was soft-deleted, if applicable.
	DeletedAt datadog.NullableTime `json:"deleted_at,omitempty"`
	// Description of the experiment.
	Description datadog.NullableString `json:"description"`
	// Error message describing why the experiment failed, if applicable.
	Error datadog.NullableString `json:"error,omitempty"`
	// Logical name of the experiment, shared across all runs of the same pipeline.
	Experiment *string `json:"experiment,omitempty"`
	// Arbitrary metadata associated with the experiment.
	Metadata map[string]interface{} `json:"metadata"`
	// Name of the experiment.
	Name string `json:"name"`
	// Identifier of the parent (baseline) experiment this experiment was run against, if any.
	ParentExperimentId datadog.NullableString `json:"parent_experiment_id,omitempty"`
	// Identifier of the project this experiment belongs to.
	ProjectId string `json:"project_id"`
	// Expected number of runs for this experiment.
	RunCount *int32 `json:"run_count,omitempty"`
	// Execution status of an LLM Observability experiment.
	Status *LLMObsExperimentStatus `json:"status,omitempty"`
	// Timestamp when the experiment was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentDataAttributesResponse instantiates a new LLMObsExperimentDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentDataAttributesResponse(config map[string]interface{}, createdAt time.Time, datasetId string, description datadog.NullableString, metadata map[string]interface{}, name string, projectId string, updatedAt time.Time) *LLMObsExperimentDataAttributesResponse {
	this := LLMObsExperimentDataAttributesResponse{}
	this.Config = config
	this.CreatedAt = createdAt
	this.DatasetId = datasetId
	this.Description = description
	this.Metadata = metadata
	this.Name = name
	this.ProjectId = projectId
	this.UpdatedAt = updatedAt
	return &this
}

// NewLLMObsExperimentDataAttributesResponseWithDefaults instantiates a new LLMObsExperimentDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentDataAttributesResponseWithDefaults() *LLMObsExperimentDataAttributesResponse {
	this := LLMObsExperimentDataAttributesResponse{}
	return &this
}

// GetAggregateData returns the AggregateData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentDataAttributesResponse) GetAggregateData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AggregateData
}

// GetAggregateDataOk returns a tuple with the AggregateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetAggregateDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.AggregateData == nil {
		return nil, false
	}
	return &o.AggregateData, true
}

// HasAggregateData returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesResponse) HasAggregateData() bool {
	return o != nil && o.AggregateData != nil
}

// SetAggregateData gets a reference to the given map[string]interface{} and assigns it to the AggregateData field.
func (o *LLMObsExperimentDataAttributesResponse) SetAggregateData(v map[string]interface{}) {
	o.AggregateData = v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *LLMObsExperimentDataAttributesResponse) GetAuthor() LLMObsExperimentUser {
	if o == nil || o.Author == nil {
		var ret LLMObsExperimentUser
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesResponse) GetAuthorOk() (*LLMObsExperimentUser, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesResponse) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given LLMObsExperimentUser and assigns it to the Author field.
func (o *LLMObsExperimentDataAttributesResponse) SetAuthor(v LLMObsExperimentUser) {
	o.Author = &v
}

// GetConfig returns the Config field value.
// If the value is explicit nil, the zero value for map[string]interface{} will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value.
func (o *LLMObsExperimentDataAttributesResponse) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsExperimentDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsExperimentDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDatasetId returns the DatasetId field value.
func (o *LLMObsExperimentDataAttributesResponse) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesResponse) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value.
func (o *LLMObsExperimentDataAttributesResponse) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetDatasetName returns the DatasetName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentDataAttributesResponse) GetDatasetName() string {
	if o == nil || o.DatasetName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatasetName.Get()
}

// GetDatasetNameOk returns a tuple with the DatasetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetDatasetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatasetName.Get(), o.DatasetName.IsSet()
}

// HasDatasetName returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesResponse) HasDatasetName() bool {
	return o != nil && o.DatasetName.IsSet()
}

// SetDatasetName gets a reference to the given datadog.NullableString and assigns it to the DatasetName field.
func (o *LLMObsExperimentDataAttributesResponse) SetDatasetName(v string) {
	o.DatasetName.Set(&v)
}

// SetDatasetNameNil sets the value for DatasetName to be an explicit nil.
func (o *LLMObsExperimentDataAttributesResponse) SetDatasetNameNil() {
	o.DatasetName.Set(nil)
}

// UnsetDatasetName ensures that no value is present for DatasetName, not even an explicit nil.
func (o *LLMObsExperimentDataAttributesResponse) UnsetDatasetName() {
	o.DatasetName.Unset()
}

// GetDatasetVersion returns the DatasetVersion field value if set, zero value otherwise.
func (o *LLMObsExperimentDataAttributesResponse) GetDatasetVersion() int64 {
	if o == nil || o.DatasetVersion == nil {
		var ret int64
		return ret
	}
	return *o.DatasetVersion
}

// GetDatasetVersionOk returns a tuple with the DatasetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesResponse) GetDatasetVersionOk() (*int64, bool) {
	if o == nil || o.DatasetVersion == nil {
		return nil, false
	}
	return o.DatasetVersion, true
}

// HasDatasetVersion returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesResponse) HasDatasetVersion() bool {
	return o != nil && o.DatasetVersion != nil
}

// SetDatasetVersion gets a reference to the given int64 and assigns it to the DatasetVersion field.
func (o *LLMObsExperimentDataAttributesResponse) SetDatasetVersion(v int64) {
	o.DatasetVersion = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentDataAttributesResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesResponse) HasDeletedAt() bool {
	return o != nil && o.DeletedAt.IsSet()
}

// SetDeletedAt gets a reference to the given datadog.NullableTime and assigns it to the DeletedAt field.
func (o *LLMObsExperimentDataAttributesResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil.
func (o *LLMObsExperimentDataAttributesResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil.
func (o *LLMObsExperimentDataAttributesResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetDescription returns the Description field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value.
func (o *LLMObsExperimentDataAttributesResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentDataAttributesResponse) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesResponse) HasError() bool {
	return o != nil && o.Error.IsSet()
}

// SetError gets a reference to the given datadog.NullableString and assigns it to the Error field.
func (o *LLMObsExperimentDataAttributesResponse) SetError(v string) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil.
func (o *LLMObsExperimentDataAttributesResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil.
func (o *LLMObsExperimentDataAttributesResponse) UnsetError() {
	o.Error.Unset()
}

// GetExperiment returns the Experiment field value if set, zero value otherwise.
func (o *LLMObsExperimentDataAttributesResponse) GetExperiment() string {
	if o == nil || o.Experiment == nil {
		var ret string
		return ret
	}
	return *o.Experiment
}

// GetExperimentOk returns a tuple with the Experiment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesResponse) GetExperimentOk() (*string, bool) {
	if o == nil || o.Experiment == nil {
		return nil, false
	}
	return o.Experiment, true
}

// HasExperiment returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesResponse) HasExperiment() bool {
	return o != nil && o.Experiment != nil
}

// SetExperiment gets a reference to the given string and assigns it to the Experiment field.
func (o *LLMObsExperimentDataAttributesResponse) SetExperiment(v string) {
	o.Experiment = &v
}

// GetMetadata returns the Metadata field value.
// If the value is explicit nil, the zero value for map[string]interface{} will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value.
func (o *LLMObsExperimentDataAttributesResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value.
func (o *LLMObsExperimentDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsExperimentDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetParentExperimentId returns the ParentExperimentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentDataAttributesResponse) GetParentExperimentId() string {
	if o == nil || o.ParentExperimentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentExperimentId.Get()
}

// GetParentExperimentIdOk returns a tuple with the ParentExperimentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentDataAttributesResponse) GetParentExperimentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentExperimentId.Get(), o.ParentExperimentId.IsSet()
}

// HasParentExperimentId returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesResponse) HasParentExperimentId() bool {
	return o != nil && o.ParentExperimentId.IsSet()
}

// SetParentExperimentId gets a reference to the given datadog.NullableString and assigns it to the ParentExperimentId field.
func (o *LLMObsExperimentDataAttributesResponse) SetParentExperimentId(v string) {
	o.ParentExperimentId.Set(&v)
}

// SetParentExperimentIdNil sets the value for ParentExperimentId to be an explicit nil.
func (o *LLMObsExperimentDataAttributesResponse) SetParentExperimentIdNil() {
	o.ParentExperimentId.Set(nil)
}

// UnsetParentExperimentId ensures that no value is present for ParentExperimentId, not even an explicit nil.
func (o *LLMObsExperimentDataAttributesResponse) UnsetParentExperimentId() {
	o.ParentExperimentId.Unset()
}

// GetProjectId returns the ProjectId field value.
func (o *LLMObsExperimentDataAttributesResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *LLMObsExperimentDataAttributesResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// GetRunCount returns the RunCount field value if set, zero value otherwise.
func (o *LLMObsExperimentDataAttributesResponse) GetRunCount() int32 {
	if o == nil || o.RunCount == nil {
		var ret int32
		return ret
	}
	return *o.RunCount
}

// GetRunCountOk returns a tuple with the RunCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesResponse) GetRunCountOk() (*int32, bool) {
	if o == nil || o.RunCount == nil {
		return nil, false
	}
	return o.RunCount, true
}

// HasRunCount returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesResponse) HasRunCount() bool {
	return o != nil && o.RunCount != nil
}

// SetRunCount gets a reference to the given int32 and assigns it to the RunCount field.
func (o *LLMObsExperimentDataAttributesResponse) SetRunCount(v int32) {
	o.RunCount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LLMObsExperimentDataAttributesResponse) GetStatus() LLMObsExperimentStatus {
	if o == nil || o.Status == nil {
		var ret LLMObsExperimentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesResponse) GetStatusOk() (*LLMObsExperimentStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given LLMObsExperimentStatus and assigns it to the Status field.
func (o *LLMObsExperimentDataAttributesResponse) SetStatus(v LLMObsExperimentStatus) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *LLMObsExperimentDataAttributesResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *LLMObsExperimentDataAttributesResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AggregateData != nil {
		toSerialize["aggregate_data"] = o.AggregateData
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["dataset_id"] = o.DatasetId
	if o.DatasetName.IsSet() {
		toSerialize["dataset_name"] = o.DatasetName.Get()
	}
	if o.DatasetVersion != nil {
		toSerialize["dataset_version"] = o.DatasetVersion
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	toSerialize["description"] = o.Description.Get()
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.Experiment != nil {
		toSerialize["experiment"] = o.Experiment
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if o.ParentExperimentId.IsSet() {
		toSerialize["parent_experiment_id"] = o.ParentExperimentId.Get()
	}
	toSerialize["project_id"] = o.ProjectId
	if o.RunCount != nil {
		toSerialize["run_count"] = o.RunCount
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
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
func (o *LLMObsExperimentDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregateData      map[string]interface{}  `json:"aggregate_data,omitempty"`
		Author             *LLMObsExperimentUser   `json:"author,omitempty"`
		Config             map[string]interface{}  `json:"config"`
		CreatedAt          *time.Time              `json:"created_at"`
		DatasetId          *string                 `json:"dataset_id"`
		DatasetName        datadog.NullableString  `json:"dataset_name,omitempty"`
		DatasetVersion     *int64                  `json:"dataset_version,omitempty"`
		DeletedAt          datadog.NullableTime    `json:"deleted_at,omitempty"`
		Description        datadog.NullableString  `json:"description"`
		Error              datadog.NullableString  `json:"error,omitempty"`
		Experiment         *string                 `json:"experiment,omitempty"`
		Metadata           map[string]interface{}  `json:"metadata"`
		Name               *string                 `json:"name"`
		ParentExperimentId datadog.NullableString  `json:"parent_experiment_id,omitempty"`
		ProjectId          *string                 `json:"project_id"`
		RunCount           *int32                  `json:"run_count,omitempty"`
		Status             *LLMObsExperimentStatus `json:"status,omitempty"`
		UpdatedAt          *time.Time              `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Config == nil {
		return fmt.Errorf("required field config missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DatasetId == nil {
		return fmt.Errorf("required field dataset_id missing")
	}
	if !all.Description.IsSet() {
		return fmt.Errorf("required field description missing")
	}
	if all.Metadata == nil {
		return fmt.Errorf("required field metadata missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregate_data", "author", "config", "created_at", "dataset_id", "dataset_name", "dataset_version", "deleted_at", "description", "error", "experiment", "metadata", "name", "parent_experiment_id", "project_id", "run_count", "status", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AggregateData = all.AggregateData
	if all.Author != nil && all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = all.Author
	o.Config = all.Config
	o.CreatedAt = *all.CreatedAt
	o.DatasetId = *all.DatasetId
	o.DatasetName = all.DatasetName
	o.DatasetVersion = all.DatasetVersion
	o.DeletedAt = all.DeletedAt
	o.Description = all.Description
	o.Error = all.Error
	o.Experiment = all.Experiment
	o.Metadata = all.Metadata
	o.Name = *all.Name
	o.ParentExperimentId = all.ParentExperimentId
	o.ProjectId = *all.ProjectId
	o.RunCount = all.RunCount
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
