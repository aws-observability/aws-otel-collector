// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetUpdateDataAttributesRequest Attributes for updating an LLM Observability dataset.
type LLMObsDatasetUpdateDataAttributesRequest struct {
	// Updated description of the dataset.
	Description *string `json:"description,omitempty"`
	// Updated metadata associated with the dataset.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Updated name of the dataset.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetUpdateDataAttributesRequest instantiates a new LLMObsDatasetUpdateDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetUpdateDataAttributesRequest() *LLMObsDatasetUpdateDataAttributesRequest {
	this := LLMObsDatasetUpdateDataAttributesRequest{}
	return &this
}

// NewLLMObsDatasetUpdateDataAttributesRequestWithDefaults instantiates a new LLMObsDatasetUpdateDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetUpdateDataAttributesRequestWithDefaults() *LLMObsDatasetUpdateDataAttributesRequest {
	this := LLMObsDatasetUpdateDataAttributesRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsDatasetUpdateDataAttributesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetUpdateDataAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsDatasetUpdateDataAttributesRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsDatasetUpdateDataAttributesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LLMObsDatasetUpdateDataAttributesRequest) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetUpdateDataAttributesRequest) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LLMObsDatasetUpdateDataAttributesRequest) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LLMObsDatasetUpdateDataAttributesRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LLMObsDatasetUpdateDataAttributesRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetUpdateDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LLMObsDatasetUpdateDataAttributesRequest) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LLMObsDatasetUpdateDataAttributesRequest) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetUpdateDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetUpdateDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                `json:"description,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Name        *string                `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "metadata", "name"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Metadata = all.Metadata
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
