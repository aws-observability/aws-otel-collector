// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigVertexAIOptions Google Vertex AI-specific options for LLM provider configuration.
type LLMObsCustomEvalConfigVertexAIOptions struct {
	// Google Cloud region.
	Location *string `json:"location,omitempty"`
	// Google Cloud project ID.
	Project *string `json:"project,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCustomEvalConfigVertexAIOptions instantiates a new LLMObsCustomEvalConfigVertexAIOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCustomEvalConfigVertexAIOptions() *LLMObsCustomEvalConfigVertexAIOptions {
	this := LLMObsCustomEvalConfigVertexAIOptions{}
	return &this
}

// NewLLMObsCustomEvalConfigVertexAIOptionsWithDefaults instantiates a new LLMObsCustomEvalConfigVertexAIOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCustomEvalConfigVertexAIOptionsWithDefaults() *LLMObsCustomEvalConfigVertexAIOptions {
	this := LLMObsCustomEvalConfigVertexAIOptions{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigVertexAIOptions) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigVertexAIOptions) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigVertexAIOptions) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *LLMObsCustomEvalConfigVertexAIOptions) SetLocation(v string) {
	o.Location = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *LLMObsCustomEvalConfigVertexAIOptions) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCustomEvalConfigVertexAIOptions) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *LLMObsCustomEvalConfigVertexAIOptions) HasProject() bool {
	return o != nil && o.Project != nil
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *LLMObsCustomEvalConfigVertexAIOptions) SetProject(v string) {
	o.Project = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCustomEvalConfigVertexAIOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCustomEvalConfigVertexAIOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Location *string `json:"location,omitempty"`
		Project  *string `json:"project,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"location", "project"})
	} else {
		return err
	}
	o.Location = all.Location
	o.Project = all.Project

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
