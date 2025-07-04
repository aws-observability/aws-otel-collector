// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListPipelinesResponse Represents the response payload containing a list of pipelines and associated metadata.
type ListPipelinesResponse struct {
	// The `schema` `data`.
	Data []ObservabilityPipelineData `json:"data"`
	// Metadata about the response.
	Meta *ListPipelinesResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListPipelinesResponse instantiates a new ListPipelinesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListPipelinesResponse(data []ObservabilityPipelineData) *ListPipelinesResponse {
	this := ListPipelinesResponse{}
	this.Data = data
	return &this
}

// NewListPipelinesResponseWithDefaults instantiates a new ListPipelinesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListPipelinesResponseWithDefaults() *ListPipelinesResponse {
	this := ListPipelinesResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ListPipelinesResponse) GetData() []ObservabilityPipelineData {
	if o == nil {
		var ret []ObservabilityPipelineData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListPipelinesResponse) GetDataOk() (*[]ObservabilityPipelineData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ListPipelinesResponse) SetData(v []ObservabilityPipelineData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListPipelinesResponse) GetMeta() ListPipelinesResponseMeta {
	if o == nil || o.Meta == nil {
		var ret ListPipelinesResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPipelinesResponse) GetMetaOk() (*ListPipelinesResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListPipelinesResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ListPipelinesResponseMeta and assigns it to the Meta field.
func (o *ListPipelinesResponse) SetMeta(v ListPipelinesResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListPipelinesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListPipelinesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]ObservabilityPipelineData `json:"data"`
		Meta *ListPipelinesResponseMeta   `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
