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
	// Configuration parameters for the experiment.
	Config map[string]interface{} `json:"config"`
	// Timestamp when the experiment was created.
	CreatedAt time.Time `json:"created_at"`
	// Identifier of the dataset used in this experiment.
	DatasetId string `json:"dataset_id"`
	// Description of the experiment.
	Description datadog.NullableString `json:"description"`
	// Arbitrary metadata associated with the experiment.
	Metadata map[string]interface{} `json:"metadata"`
	// Name of the experiment.
	Name string `json:"name"`
	// Identifier of the project this experiment belongs to.
	ProjectId string `json:"project_id"`
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
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["dataset_id"] = o.DatasetId
	toSerialize["description"] = o.Description.Get()
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	toSerialize["project_id"] = o.ProjectId
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
		Config      map[string]interface{} `json:"config"`
		CreatedAt   *time.Time             `json:"created_at"`
		DatasetId   *string                `json:"dataset_id"`
		Description datadog.NullableString `json:"description"`
		Metadata    map[string]interface{} `json:"metadata"`
		Name        *string                `json:"name"`
		ProjectId   *string                `json:"project_id"`
		UpdatedAt   *time.Time             `json:"updated_at"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "created_at", "dataset_id", "description", "metadata", "name", "project_id", "updated_at"})
	} else {
		return err
	}
	o.Config = all.Config
	o.CreatedAt = *all.CreatedAt
	o.DatasetId = *all.DatasetId
	o.Description = all.Description
	o.Metadata = all.Metadata
	o.Name = *all.Name
	o.ProjectId = *all.ProjectId
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
