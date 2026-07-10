// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAzureOpenAIMetadata Azure OpenAI-specific metadata for an integration account or inference request.
type LLMObsAzureOpenAIMetadata struct {
	// The Azure OpenAI deployment ID.
	DeploymentId *string `json:"deployment_id,omitempty"`
	// The model version deployed in Azure.
	ModelVersion *string `json:"model_version,omitempty"`
	// The Azure OpenAI resource name.
	ResourceName *string `json:"resource_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAzureOpenAIMetadata instantiates a new LLMObsAzureOpenAIMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAzureOpenAIMetadata() *LLMObsAzureOpenAIMetadata {
	this := LLMObsAzureOpenAIMetadata{}
	return &this
}

// NewLLMObsAzureOpenAIMetadataWithDefaults instantiates a new LLMObsAzureOpenAIMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAzureOpenAIMetadataWithDefaults() *LLMObsAzureOpenAIMetadata {
	this := LLMObsAzureOpenAIMetadata{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *LLMObsAzureOpenAIMetadata) GetDeploymentId() string {
	if o == nil || o.DeploymentId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAzureOpenAIMetadata) GetDeploymentIdOk() (*string, bool) {
	if o == nil || o.DeploymentId == nil {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *LLMObsAzureOpenAIMetadata) HasDeploymentId() bool {
	return o != nil && o.DeploymentId != nil
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *LLMObsAzureOpenAIMetadata) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetModelVersion returns the ModelVersion field value if set, zero value otherwise.
func (o *LLMObsAzureOpenAIMetadata) GetModelVersion() string {
	if o == nil || o.ModelVersion == nil {
		var ret string
		return ret
	}
	return *o.ModelVersion
}

// GetModelVersionOk returns a tuple with the ModelVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAzureOpenAIMetadata) GetModelVersionOk() (*string, bool) {
	if o == nil || o.ModelVersion == nil {
		return nil, false
	}
	return o.ModelVersion, true
}

// HasModelVersion returns a boolean if a field has been set.
func (o *LLMObsAzureOpenAIMetadata) HasModelVersion() bool {
	return o != nil && o.ModelVersion != nil
}

// SetModelVersion gets a reference to the given string and assigns it to the ModelVersion field.
func (o *LLMObsAzureOpenAIMetadata) SetModelVersion(v string) {
	o.ModelVersion = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *LLMObsAzureOpenAIMetadata) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAzureOpenAIMetadata) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *LLMObsAzureOpenAIMetadata) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *LLMObsAzureOpenAIMetadata) SetResourceName(v string) {
	o.ResourceName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAzureOpenAIMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeploymentId != nil {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	if o.ModelVersion != nil {
		toSerialize["model_version"] = o.ModelVersion
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAzureOpenAIMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeploymentId *string `json:"deployment_id,omitempty"`
		ModelVersion *string `json:"model_version,omitempty"`
		ResourceName *string `json:"resource_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deployment_id", "model_version", "resource_name"})
	} else {
		return err
	}
	o.DeploymentId = all.DeploymentId
	o.ModelVersion = all.ModelVersion
	o.ResourceName = all.ResourceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
