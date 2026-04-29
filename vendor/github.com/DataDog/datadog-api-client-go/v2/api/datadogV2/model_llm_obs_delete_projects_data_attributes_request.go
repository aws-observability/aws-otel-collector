// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDeleteProjectsDataAttributesRequest Attributes for deleting LLM Observability projects.
type LLMObsDeleteProjectsDataAttributesRequest struct {
	// List of project IDs to delete.
	ProjectIds []string `json:"project_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDeleteProjectsDataAttributesRequest instantiates a new LLMObsDeleteProjectsDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDeleteProjectsDataAttributesRequest(projectIds []string) *LLMObsDeleteProjectsDataAttributesRequest {
	this := LLMObsDeleteProjectsDataAttributesRequest{}
	this.ProjectIds = projectIds
	return &this
}

// NewLLMObsDeleteProjectsDataAttributesRequestWithDefaults instantiates a new LLMObsDeleteProjectsDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDeleteProjectsDataAttributesRequestWithDefaults() *LLMObsDeleteProjectsDataAttributesRequest {
	this := LLMObsDeleteProjectsDataAttributesRequest{}
	return &this
}

// GetProjectIds returns the ProjectIds field value.
func (o *LLMObsDeleteProjectsDataAttributesRequest) GetProjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeleteProjectsDataAttributesRequest) GetProjectIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectIds, true
}

// SetProjectIds sets field value.
func (o *LLMObsDeleteProjectsDataAttributesRequest) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDeleteProjectsDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["project_ids"] = o.ProjectIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDeleteProjectsDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProjectIds *[]string `json:"project_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProjectIds == nil {
		return fmt.Errorf("required field project_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"project_ids"})
	} else {
		return err
	}
	o.ProjectIds = *all.ProjectIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
