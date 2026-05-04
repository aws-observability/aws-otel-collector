// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueUpdateDataAttributesRequest Attributes for updating an LLM Observability annotation queue. All fields are optional.
type LLMObsAnnotationQueueUpdateDataAttributesRequest struct {
	// Updated description of the annotation queue.
	Description *string `json:"description,omitempty"`
	// Updated name of the annotation queue.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueUpdateDataAttributesRequest instantiates a new LLMObsAnnotationQueueUpdateDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueUpdateDataAttributesRequest() *LLMObsAnnotationQueueUpdateDataAttributesRequest {
	this := LLMObsAnnotationQueueUpdateDataAttributesRequest{}
	return &this
}

// NewLLMObsAnnotationQueueUpdateDataAttributesRequestWithDefaults instantiates a new LLMObsAnnotationQueueUpdateDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueUpdateDataAttributesRequestWithDefaults() *LLMObsAnnotationQueueUpdateDataAttributesRequest {
	this := LLMObsAnnotationQueueUpdateDataAttributesRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueUpdateDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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
func (o *LLMObsAnnotationQueueUpdateDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
