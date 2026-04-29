// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultCI CI information associated with the test result.
type SyntheticsTestResultCI struct {
	// Details of the CI pipeline.
	Pipeline *SyntheticsTestResultCIPipeline `json:"pipeline,omitempty"`
	// Details of the CI provider.
	Provider *SyntheticsTestResultCIProvider `json:"provider,omitempty"`
	// Details of the CI stage.
	Stage *SyntheticsTestResultCIStage `json:"stage,omitempty"`
	// Path of the workspace that ran the CI job.
	WorkspacePath *string `json:"workspace_path,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultCI instantiates a new SyntheticsTestResultCI object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultCI() *SyntheticsTestResultCI {
	this := SyntheticsTestResultCI{}
	return &this
}

// NewSyntheticsTestResultCIWithDefaults instantiates a new SyntheticsTestResultCI object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultCIWithDefaults() *SyntheticsTestResultCI {
	this := SyntheticsTestResultCI{}
	return &this
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *SyntheticsTestResultCI) GetPipeline() SyntheticsTestResultCIPipeline {
	if o == nil || o.Pipeline == nil {
		var ret SyntheticsTestResultCIPipeline
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCI) GetPipelineOk() (*SyntheticsTestResultCIPipeline, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *SyntheticsTestResultCI) HasPipeline() bool {
	return o != nil && o.Pipeline != nil
}

// SetPipeline gets a reference to the given SyntheticsTestResultCIPipeline and assigns it to the Pipeline field.
func (o *SyntheticsTestResultCI) SetPipeline(v SyntheticsTestResultCIPipeline) {
	o.Pipeline = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *SyntheticsTestResultCI) GetProvider() SyntheticsTestResultCIProvider {
	if o == nil || o.Provider == nil {
		var ret SyntheticsTestResultCIProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCI) GetProviderOk() (*SyntheticsTestResultCIProvider, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *SyntheticsTestResultCI) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given SyntheticsTestResultCIProvider and assigns it to the Provider field.
func (o *SyntheticsTestResultCI) SetProvider(v SyntheticsTestResultCIProvider) {
	o.Provider = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *SyntheticsTestResultCI) GetStage() SyntheticsTestResultCIStage {
	if o == nil || o.Stage == nil {
		var ret SyntheticsTestResultCIStage
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCI) GetStageOk() (*SyntheticsTestResultCIStage, bool) {
	if o == nil || o.Stage == nil {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *SyntheticsTestResultCI) HasStage() bool {
	return o != nil && o.Stage != nil
}

// SetStage gets a reference to the given SyntheticsTestResultCIStage and assigns it to the Stage field.
func (o *SyntheticsTestResultCI) SetStage(v SyntheticsTestResultCIStage) {
	o.Stage = &v
}

// GetWorkspacePath returns the WorkspacePath field value if set, zero value otherwise.
func (o *SyntheticsTestResultCI) GetWorkspacePath() string {
	if o == nil || o.WorkspacePath == nil {
		var ret string
		return ret
	}
	return *o.WorkspacePath
}

// GetWorkspacePathOk returns a tuple with the WorkspacePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCI) GetWorkspacePathOk() (*string, bool) {
	if o == nil || o.WorkspacePath == nil {
		return nil, false
	}
	return o.WorkspacePath, true
}

// HasWorkspacePath returns a boolean if a field has been set.
func (o *SyntheticsTestResultCI) HasWorkspacePath() bool {
	return o != nil && o.WorkspacePath != nil
}

// SetWorkspacePath gets a reference to the given string and assigns it to the WorkspacePath field.
func (o *SyntheticsTestResultCI) SetWorkspacePath(v string) {
	o.WorkspacePath = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultCI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Stage != nil {
		toSerialize["stage"] = o.Stage
	}
	if o.WorkspacePath != nil {
		toSerialize["workspace_path"] = o.WorkspacePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultCI) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pipeline      *SyntheticsTestResultCIPipeline `json:"pipeline,omitempty"`
		Provider      *SyntheticsTestResultCIProvider `json:"provider,omitempty"`
		Stage         *SyntheticsTestResultCIStage    `json:"stage,omitempty"`
		WorkspacePath *string                         `json:"workspace_path,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"pipeline", "provider", "stage", "workspace_path"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Pipeline != nil && all.Pipeline.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pipeline = all.Pipeline
	if all.Provider != nil && all.Provider.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Provider = all.Provider
	if all.Stage != nil && all.Stage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Stage = all.Stage
	o.WorkspacePath = all.WorkspacePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
