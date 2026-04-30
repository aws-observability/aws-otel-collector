// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentDataAttributesRequest Attributes for creating an LLM Observability experiment.
type LLMObsExperimentDataAttributesRequest struct {
	// Configuration parameters for the experiment.
	Config map[string]interface{} `json:"config,omitempty"`
	// Identifier of the dataset used in this experiment.
	DatasetId string `json:"dataset_id"`
	// Version of the dataset to use. Defaults to the current version if not specified.
	DatasetVersion *int64 `json:"dataset_version,omitempty"`
	// Description of the experiment.
	Description *string `json:"description,omitempty"`
	// Whether to ensure the experiment name is unique. Defaults to `true`.
	EnsureUnique *bool `json:"ensure_unique,omitempty"`
	// Arbitrary metadata associated with the experiment.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Name of the experiment.
	Name string `json:"name"`
	// Identifier of the project this experiment belongs to.
	ProjectId string `json:"project_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentDataAttributesRequest instantiates a new LLMObsExperimentDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentDataAttributesRequest(datasetId string, name string, projectId string) *LLMObsExperimentDataAttributesRequest {
	this := LLMObsExperimentDataAttributesRequest{}
	this.DatasetId = datasetId
	this.Name = name
	this.ProjectId = projectId
	return &this
}

// NewLLMObsExperimentDataAttributesRequestWithDefaults instantiates a new LLMObsExperimentDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentDataAttributesRequestWithDefaults() *LLMObsExperimentDataAttributesRequest {
	this := LLMObsExperimentDataAttributesRequest{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *LLMObsExperimentDataAttributesRequest) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesRequest) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesRequest) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *LLMObsExperimentDataAttributesRequest) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetDatasetId returns the DatasetId field value.
func (o *LLMObsExperimentDataAttributesRequest) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesRequest) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value.
func (o *LLMObsExperimentDataAttributesRequest) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetDatasetVersion returns the DatasetVersion field value if set, zero value otherwise.
func (o *LLMObsExperimentDataAttributesRequest) GetDatasetVersion() int64 {
	if o == nil || o.DatasetVersion == nil {
		var ret int64
		return ret
	}
	return *o.DatasetVersion
}

// GetDatasetVersionOk returns a tuple with the DatasetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesRequest) GetDatasetVersionOk() (*int64, bool) {
	if o == nil || o.DatasetVersion == nil {
		return nil, false
	}
	return o.DatasetVersion, true
}

// HasDatasetVersion returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesRequest) HasDatasetVersion() bool {
	return o != nil && o.DatasetVersion != nil
}

// SetDatasetVersion gets a reference to the given int64 and assigns it to the DatasetVersion field.
func (o *LLMObsExperimentDataAttributesRequest) SetDatasetVersion(v int64) {
	o.DatasetVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsExperimentDataAttributesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsExperimentDataAttributesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnsureUnique returns the EnsureUnique field value if set, zero value otherwise.
func (o *LLMObsExperimentDataAttributesRequest) GetEnsureUnique() bool {
	if o == nil || o.EnsureUnique == nil {
		var ret bool
		return ret
	}
	return *o.EnsureUnique
}

// GetEnsureUniqueOk returns a tuple with the EnsureUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesRequest) GetEnsureUniqueOk() (*bool, bool) {
	if o == nil || o.EnsureUnique == nil {
		return nil, false
	}
	return o.EnsureUnique, true
}

// HasEnsureUnique returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesRequest) HasEnsureUnique() bool {
	return o != nil && o.EnsureUnique != nil
}

// SetEnsureUnique gets a reference to the given bool and assigns it to the EnsureUnique field.
func (o *LLMObsExperimentDataAttributesRequest) SetEnsureUnique(v bool) {
	o.EnsureUnique = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LLMObsExperimentDataAttributesRequest) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesRequest) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LLMObsExperimentDataAttributesRequest) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LLMObsExperimentDataAttributesRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value.
func (o *LLMObsExperimentDataAttributesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsExperimentDataAttributesRequest) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value.
func (o *LLMObsExperimentDataAttributesRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentDataAttributesRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *LLMObsExperimentDataAttributesRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	toSerialize["dataset_id"] = o.DatasetId
	if o.DatasetVersion != nil {
		toSerialize["dataset_version"] = o.DatasetVersion
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EnsureUnique != nil {
		toSerialize["ensure_unique"] = o.EnsureUnique
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	toSerialize["project_id"] = o.ProjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config         map[string]interface{} `json:"config,omitempty"`
		DatasetId      *string                `json:"dataset_id"`
		DatasetVersion *int64                 `json:"dataset_version,omitempty"`
		Description    *string                `json:"description,omitempty"`
		EnsureUnique   *bool                  `json:"ensure_unique,omitempty"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
		Name           *string                `json:"name"`
		ProjectId      *string                `json:"project_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatasetId == nil {
		return fmt.Errorf("required field dataset_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "dataset_id", "dataset_version", "description", "ensure_unique", "metadata", "name", "project_id"})
	} else {
		return err
	}
	o.Config = all.Config
	o.DatasetId = *all.DatasetId
	o.DatasetVersion = all.DatasetVersion
	o.Description = all.Description
	o.EnsureUnique = all.EnsureUnique
	o.Metadata = all.Metadata
	o.Name = *all.Name
	o.ProjectId = *all.ProjectId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
