// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppCreatePipelineEventRequestAttributes Attributes of the pipeline event to create.
type CIAppCreatePipelineEventRequestAttributes struct {
	// The Datadog environment.
	Env *string `json:"env,omitempty"`
	// The name of the CI provider. By default, this is "custom".
	ProviderName *string `json:"provider_name,omitempty"`
	// Details of the CI pipeline event.
	Resource CIAppCreatePipelineEventRequestAttributesResource `json:"resource"`
	// If the CI provider is SaaS, use this to differentiate between instances.
	Service *string `json:"service,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppCreatePipelineEventRequestAttributes instantiates a new CIAppCreatePipelineEventRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppCreatePipelineEventRequestAttributes(resource CIAppCreatePipelineEventRequestAttributesResource) *CIAppCreatePipelineEventRequestAttributes {
	this := CIAppCreatePipelineEventRequestAttributes{}
	this.Resource = resource
	return &this
}

// NewCIAppCreatePipelineEventRequestAttributesWithDefaults instantiates a new CIAppCreatePipelineEventRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppCreatePipelineEventRequestAttributesWithDefaults() *CIAppCreatePipelineEventRequestAttributes {
	this := CIAppCreatePipelineEventRequestAttributes{}
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *CIAppCreatePipelineEventRequestAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppCreatePipelineEventRequestAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *CIAppCreatePipelineEventRequestAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *CIAppCreatePipelineEventRequestAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *CIAppCreatePipelineEventRequestAttributes) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppCreatePipelineEventRequestAttributes) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *CIAppCreatePipelineEventRequestAttributes) HasProviderName() bool {
	return o != nil && o.ProviderName != nil
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *CIAppCreatePipelineEventRequestAttributes) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetResource returns the Resource field value.
func (o *CIAppCreatePipelineEventRequestAttributes) GetResource() CIAppCreatePipelineEventRequestAttributesResource {
	if o == nil {
		var ret CIAppCreatePipelineEventRequestAttributesResource
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *CIAppCreatePipelineEventRequestAttributes) GetResourceOk() (*CIAppCreatePipelineEventRequestAttributesResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value.
func (o *CIAppCreatePipelineEventRequestAttributes) SetResource(v CIAppCreatePipelineEventRequestAttributesResource) {
	o.Resource = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *CIAppCreatePipelineEventRequestAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppCreatePipelineEventRequestAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CIAppCreatePipelineEventRequestAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *CIAppCreatePipelineEventRequestAttributes) SetService(v string) {
	o.Service = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppCreatePipelineEventRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.ProviderName != nil {
		toSerialize["provider_name"] = o.ProviderName
	}
	toSerialize["resource"] = o.Resource
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppCreatePipelineEventRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Env          *string                                            `json:"env,omitempty"`
		ProviderName *string                                            `json:"provider_name,omitempty"`
		Resource     *CIAppCreatePipelineEventRequestAttributesResource `json:"resource"`
		Service      *string                                            `json:"service,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Resource == nil {
		return fmt.Errorf("required field resource missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"env", "provider_name", "resource", "service"})
	} else {
		return err
	}
	o.Env = all.Env
	o.ProviderName = all.ProviderName
	o.Resource = *all.Resource
	o.Service = all.Service

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
