// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDeleteExperimentsDataAttributesRequest Attributes for deleting LLM Observability experiments.
type LLMObsDeleteExperimentsDataAttributesRequest struct {
	// List of experiment IDs to delete.
	ExperimentIds []string `json:"experiment_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDeleteExperimentsDataAttributesRequest instantiates a new LLMObsDeleteExperimentsDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDeleteExperimentsDataAttributesRequest(experimentIds []string) *LLMObsDeleteExperimentsDataAttributesRequest {
	this := LLMObsDeleteExperimentsDataAttributesRequest{}
	this.ExperimentIds = experimentIds
	return &this
}

// NewLLMObsDeleteExperimentsDataAttributesRequestWithDefaults instantiates a new LLMObsDeleteExperimentsDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDeleteExperimentsDataAttributesRequestWithDefaults() *LLMObsDeleteExperimentsDataAttributesRequest {
	this := LLMObsDeleteExperimentsDataAttributesRequest{}
	return &this
}

// GetExperimentIds returns the ExperimentIds field value.
func (o *LLMObsDeleteExperimentsDataAttributesRequest) GetExperimentIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExperimentIds
}

// GetExperimentIdsOk returns a tuple with the ExperimentIds field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeleteExperimentsDataAttributesRequest) GetExperimentIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExperimentIds, true
}

// SetExperimentIds sets field value.
func (o *LLMObsDeleteExperimentsDataAttributesRequest) SetExperimentIds(v []string) {
	o.ExperimentIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDeleteExperimentsDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["experiment_ids"] = o.ExperimentIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDeleteExperimentsDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExperimentIds *[]string `json:"experiment_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExperimentIds == nil {
		return fmt.Errorf("required field experiment_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"experiment_ids"})
	} else {
		return err
	}
	o.ExperimentIds = *all.ExperimentIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
