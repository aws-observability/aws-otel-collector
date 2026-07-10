// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsIntegrationModel A model available for a given LLM provider integration and account.
type LLMObsIntegrationModel struct {
	// Whether the account has access to this model.
	HasAccess bool `json:"has_access"`
	// Unique identifier for the model entry.
	Id string `json:"id"`
	// The name of the LLM provider integration.
	Integration string `json:"integration"`
	// Human-readable name of the LLM provider integration.
	IntegrationDisplayName string `json:"integration_display_name"`
	// Whether the model supports structured output via JSON schema.
	JsonSchema bool `json:"json_schema"`
	// Human-readable model name.
	ModelDisplayName string `json:"model_display_name"`
	// Provider-specific model identifier used in inference calls.
	ModelId string `json:"model_id"`
	// The underlying model provider.
	Provider string `json:"provider"`
	// Human-readable name of the underlying model provider.
	ProviderDisplayName string `json:"provider_display_name"`
	// Map of region-specific model ID prefix overrides.
	RegionPrefixOverrides map[string]string `json:"region_prefix_overrides,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsIntegrationModel instantiates a new LLMObsIntegrationModel object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsIntegrationModel(hasAccess bool, id string, integration string, integrationDisplayName string, jsonSchema bool, modelDisplayName string, modelId string, provider string, providerDisplayName string) *LLMObsIntegrationModel {
	this := LLMObsIntegrationModel{}
	this.HasAccess = hasAccess
	this.Id = id
	this.Integration = integration
	this.IntegrationDisplayName = integrationDisplayName
	this.JsonSchema = jsonSchema
	this.ModelDisplayName = modelDisplayName
	this.ModelId = modelId
	this.Provider = provider
	this.ProviderDisplayName = providerDisplayName
	return &this
}

// NewLLMObsIntegrationModelWithDefaults instantiates a new LLMObsIntegrationModel object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsIntegrationModelWithDefaults() *LLMObsIntegrationModel {
	this := LLMObsIntegrationModel{}
	return &this
}

// GetHasAccess returns the HasAccess field value.
func (o *LLMObsIntegrationModel) GetHasAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationModel) GetHasAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasAccess, true
}

// SetHasAccess sets field value.
func (o *LLMObsIntegrationModel) SetHasAccess(v bool) {
	o.HasAccess = v
}

// GetId returns the Id field value.
func (o *LLMObsIntegrationModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsIntegrationModel) SetId(v string) {
	o.Id = v
}

// GetIntegration returns the Integration field value.
func (o *LLMObsIntegrationModel) GetIntegration() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationModel) GetIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *LLMObsIntegrationModel) SetIntegration(v string) {
	o.Integration = v
}

// GetIntegrationDisplayName returns the IntegrationDisplayName field value.
func (o *LLMObsIntegrationModel) GetIntegrationDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IntegrationDisplayName
}

// GetIntegrationDisplayNameOk returns a tuple with the IntegrationDisplayName field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationModel) GetIntegrationDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationDisplayName, true
}

// SetIntegrationDisplayName sets field value.
func (o *LLMObsIntegrationModel) SetIntegrationDisplayName(v string) {
	o.IntegrationDisplayName = v
}

// GetJsonSchema returns the JsonSchema field value.
func (o *LLMObsIntegrationModel) GetJsonSchema() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.JsonSchema
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationModel) GetJsonSchemaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonSchema, true
}

// SetJsonSchema sets field value.
func (o *LLMObsIntegrationModel) SetJsonSchema(v bool) {
	o.JsonSchema = v
}

// GetModelDisplayName returns the ModelDisplayName field value.
func (o *LLMObsIntegrationModel) GetModelDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelDisplayName
}

// GetModelDisplayNameOk returns a tuple with the ModelDisplayName field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationModel) GetModelDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelDisplayName, true
}

// SetModelDisplayName sets field value.
func (o *LLMObsIntegrationModel) SetModelDisplayName(v string) {
	o.ModelDisplayName = v
}

// GetModelId returns the ModelId field value.
func (o *LLMObsIntegrationModel) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationModel) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value.
func (o *LLMObsIntegrationModel) SetModelId(v string) {
	o.ModelId = v
}

// GetProvider returns the Provider field value.
func (o *LLMObsIntegrationModel) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationModel) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *LLMObsIntegrationModel) SetProvider(v string) {
	o.Provider = v
}

// GetProviderDisplayName returns the ProviderDisplayName field value.
func (o *LLMObsIntegrationModel) GetProviderDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProviderDisplayName
}

// GetProviderDisplayNameOk returns a tuple with the ProviderDisplayName field value
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationModel) GetProviderDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderDisplayName, true
}

// SetProviderDisplayName sets field value.
func (o *LLMObsIntegrationModel) SetProviderDisplayName(v string) {
	o.ProviderDisplayName = v
}

// GetRegionPrefixOverrides returns the RegionPrefixOverrides field value if set, zero value otherwise.
func (o *LLMObsIntegrationModel) GetRegionPrefixOverrides() map[string]string {
	if o == nil || o.RegionPrefixOverrides == nil {
		var ret map[string]string
		return ret
	}
	return o.RegionPrefixOverrides
}

// GetRegionPrefixOverridesOk returns a tuple with the RegionPrefixOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsIntegrationModel) GetRegionPrefixOverridesOk() (*map[string]string, bool) {
	if o == nil || o.RegionPrefixOverrides == nil {
		return nil, false
	}
	return &o.RegionPrefixOverrides, true
}

// HasRegionPrefixOverrides returns a boolean if a field has been set.
func (o *LLMObsIntegrationModel) HasRegionPrefixOverrides() bool {
	return o != nil && o.RegionPrefixOverrides != nil
}

// SetRegionPrefixOverrides gets a reference to the given map[string]string and assigns it to the RegionPrefixOverrides field.
func (o *LLMObsIntegrationModel) SetRegionPrefixOverrides(v map[string]string) {
	o.RegionPrefixOverrides = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsIntegrationModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["has_access"] = o.HasAccess
	toSerialize["id"] = o.Id
	toSerialize["integration"] = o.Integration
	toSerialize["integration_display_name"] = o.IntegrationDisplayName
	toSerialize["json_schema"] = o.JsonSchema
	toSerialize["model_display_name"] = o.ModelDisplayName
	toSerialize["model_id"] = o.ModelId
	toSerialize["provider"] = o.Provider
	toSerialize["provider_display_name"] = o.ProviderDisplayName
	if o.RegionPrefixOverrides != nil {
		toSerialize["region_prefix_overrides"] = o.RegionPrefixOverrides
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsIntegrationModel) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HasAccess              *bool             `json:"has_access"`
		Id                     *string           `json:"id"`
		Integration            *string           `json:"integration"`
		IntegrationDisplayName *string           `json:"integration_display_name"`
		JsonSchema             *bool             `json:"json_schema"`
		ModelDisplayName       *string           `json:"model_display_name"`
		ModelId                *string           `json:"model_id"`
		Provider               *string           `json:"provider"`
		ProviderDisplayName    *string           `json:"provider_display_name"`
		RegionPrefixOverrides  map[string]string `json:"region_prefix_overrides,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HasAccess == nil {
		return fmt.Errorf("required field has_access missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Integration == nil {
		return fmt.Errorf("required field integration missing")
	}
	if all.IntegrationDisplayName == nil {
		return fmt.Errorf("required field integration_display_name missing")
	}
	if all.JsonSchema == nil {
		return fmt.Errorf("required field json_schema missing")
	}
	if all.ModelDisplayName == nil {
		return fmt.Errorf("required field model_display_name missing")
	}
	if all.ModelId == nil {
		return fmt.Errorf("required field model_id missing")
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.ProviderDisplayName == nil {
		return fmt.Errorf("required field provider_display_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"has_access", "id", "integration", "integration_display_name", "json_schema", "model_display_name", "model_id", "provider", "provider_display_name", "region_prefix_overrides"})
	} else {
		return err
	}
	o.HasAccess = *all.HasAccess
	o.Id = *all.Id
	o.Integration = *all.Integration
	o.IntegrationDisplayName = *all.IntegrationDisplayName
	o.JsonSchema = *all.JsonSchema
	o.ModelDisplayName = *all.ModelDisplayName
	o.ModelId = *all.ModelId
	o.Provider = *all.Provider
	o.ProviderDisplayName = *all.ProviderDisplayName
	o.RegionPrefixOverrides = all.RegionPrefixOverrides

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
