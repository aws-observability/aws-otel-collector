// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDeleteDatasetsDataAttributesRequest Attributes for deleting LLM Observability datasets.
type LLMObsDeleteDatasetsDataAttributesRequest struct {
	// List of dataset IDs to delete.
	DatasetIds []string `json:"dataset_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDeleteDatasetsDataAttributesRequest instantiates a new LLMObsDeleteDatasetsDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDeleteDatasetsDataAttributesRequest(datasetIds []string) *LLMObsDeleteDatasetsDataAttributesRequest {
	this := LLMObsDeleteDatasetsDataAttributesRequest{}
	this.DatasetIds = datasetIds
	return &this
}

// NewLLMObsDeleteDatasetsDataAttributesRequestWithDefaults instantiates a new LLMObsDeleteDatasetsDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDeleteDatasetsDataAttributesRequestWithDefaults() *LLMObsDeleteDatasetsDataAttributesRequest {
	this := LLMObsDeleteDatasetsDataAttributesRequest{}
	return &this
}

// GetDatasetIds returns the DatasetIds field value.
func (o *LLMObsDeleteDatasetsDataAttributesRequest) GetDatasetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DatasetIds
}

// GetDatasetIdsOk returns a tuple with the DatasetIds field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeleteDatasetsDataAttributesRequest) GetDatasetIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetIds, true
}

// SetDatasetIds sets field value.
func (o *LLMObsDeleteDatasetsDataAttributesRequest) SetDatasetIds(v []string) {
	o.DatasetIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDeleteDatasetsDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dataset_ids"] = o.DatasetIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDeleteDatasetsDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetIds *[]string `json:"dataset_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatasetIds == nil {
		return fmt.Errorf("required field dataset_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dataset_ids"})
	} else {
		return err
	}
	o.DatasetIds = *all.DatasetIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
