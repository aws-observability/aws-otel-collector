// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsConfigUpsertRequestAttributes Attributes for creating or updating an LLM Observability patterns configuration.
type LLMObsPatternsConfigUpsertRequestAttributes struct {
	// Integration account ID for a bring-your-own-model configuration.
	AccountId *string `json:"account_id,omitempty"`
	// The ID of an existing configuration to update. If omitted, a new configuration is created.
	ConfigId *string `json:"config_id,omitempty"`
	// Query that selects the spans the patterns run analyzes.
	EvpQuery string `json:"evp_query"`
	// Depth of the topic hierarchy to generate.
	HierarchyDepth int32 `json:"hierarchy_depth"`
	// Integration provider for a bring-your-own-model configuration.
	IntegrationProvider *string `json:"integration_provider,omitempty"`
	// Model name for a bring-your-own-model configuration.
	ModelName *string `json:"model_name,omitempty"`
	// Name of the configuration.
	Name string `json:"name"`
	// Maximum number of records to process for the run.
	NumRecords int32 `json:"num_records"`
	// Fraction of matching spans to sample for the run.
	SamplingRatio float64 `json:"sampling_ratio"`
	// Scope of the configuration.
	Scope *string `json:"scope,omitempty"`
	// Template used to guide topic generation.
	Template *string `json:"template,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsConfigUpsertRequestAttributes instantiates a new LLMObsPatternsConfigUpsertRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsConfigUpsertRequestAttributes(evpQuery string, hierarchyDepth int32, name string, numRecords int32, samplingRatio float64) *LLMObsPatternsConfigUpsertRequestAttributes {
	this := LLMObsPatternsConfigUpsertRequestAttributes{}
	this.EvpQuery = evpQuery
	this.HierarchyDepth = hierarchyDepth
	this.Name = name
	this.NumRecords = numRecords
	this.SamplingRatio = samplingRatio
	return &this
}

// NewLLMObsPatternsConfigUpsertRequestAttributesWithDefaults instantiates a new LLMObsPatternsConfigUpsertRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsConfigUpsertRequestAttributesWithDefaults() *LLMObsPatternsConfigUpsertRequestAttributes {
	this := LLMObsPatternsConfigUpsertRequestAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetAccountId(v string) {
	o.AccountId = &v
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetConfigId() string {
	if o == nil || o.ConfigId == nil {
		var ret string
		return ret
	}
	return *o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetConfigIdOk() (*string, bool) {
	if o == nil || o.ConfigId == nil {
		return nil, false
	}
	return o.ConfigId, true
}

// HasConfigId returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) HasConfigId() bool {
	return o != nil && o.ConfigId != nil
}

// SetConfigId gets a reference to the given string and assigns it to the ConfigId field.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetConfigId(v string) {
	o.ConfigId = &v
}

// GetEvpQuery returns the EvpQuery field value.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetEvpQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvpQuery
}

// GetEvpQueryOk returns a tuple with the EvpQuery field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetEvpQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvpQuery, true
}

// SetEvpQuery sets field value.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetEvpQuery(v string) {
	o.EvpQuery = v
}

// GetHierarchyDepth returns the HierarchyDepth field value.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetHierarchyDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.HierarchyDepth
}

// GetHierarchyDepthOk returns a tuple with the HierarchyDepth field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetHierarchyDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HierarchyDepth, true
}

// SetHierarchyDepth sets field value.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetHierarchyDepth(v int32) {
	o.HierarchyDepth = v
}

// GetIntegrationProvider returns the IntegrationProvider field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetIntegrationProvider() string {
	if o == nil || o.IntegrationProvider == nil {
		var ret string
		return ret
	}
	return *o.IntegrationProvider
}

// GetIntegrationProviderOk returns a tuple with the IntegrationProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetIntegrationProviderOk() (*string, bool) {
	if o == nil || o.IntegrationProvider == nil {
		return nil, false
	}
	return o.IntegrationProvider, true
}

// HasIntegrationProvider returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) HasIntegrationProvider() bool {
	return o != nil && o.IntegrationProvider != nil
}

// SetIntegrationProvider gets a reference to the given string and assigns it to the IntegrationProvider field.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetIntegrationProvider(v string) {
	o.IntegrationProvider = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetModelName() string {
	if o == nil || o.ModelName == nil {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetModelNameOk() (*string, bool) {
	if o == nil || o.ModelName == nil {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) HasModelName() bool {
	return o != nil && o.ModelName != nil
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetModelName(v string) {
	o.ModelName = &v
}

// GetName returns the Name field value.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetNumRecords returns the NumRecords field value.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetNumRecords() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.NumRecords
}

// GetNumRecordsOk returns a tuple with the NumRecords field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetNumRecordsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumRecords, true
}

// SetNumRecords sets field value.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetNumRecords(v int32) {
	o.NumRecords = v
}

// GetSamplingRatio returns the SamplingRatio field value.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetSamplingRatio() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.SamplingRatio
}

// GetSamplingRatioOk returns a tuple with the SamplingRatio field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetSamplingRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SamplingRatio, true
}

// SetSamplingRatio sets field value.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetSamplingRatio(v float64) {
	o.SamplingRatio = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetScope(v string) {
	o.Scope = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) HasTemplate() bool {
	return o != nil && o.Template != nil
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) SetTemplate(v string) {
	o.Template = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsConfigUpsertRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.ConfigId != nil {
		toSerialize["config_id"] = o.ConfigId
	}
	toSerialize["evp_query"] = o.EvpQuery
	toSerialize["hierarchy_depth"] = o.HierarchyDepth
	if o.IntegrationProvider != nil {
		toSerialize["integration_provider"] = o.IntegrationProvider
	}
	if o.ModelName != nil {
		toSerialize["model_name"] = o.ModelName
	}
	toSerialize["name"] = o.Name
	toSerialize["num_records"] = o.NumRecords
	toSerialize["sampling_ratio"] = o.SamplingRatio
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsConfigUpsertRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId           *string  `json:"account_id,omitempty"`
		ConfigId            *string  `json:"config_id,omitempty"`
		EvpQuery            *string  `json:"evp_query"`
		HierarchyDepth      *int32   `json:"hierarchy_depth"`
		IntegrationProvider *string  `json:"integration_provider,omitempty"`
		ModelName           *string  `json:"model_name,omitempty"`
		Name                *string  `json:"name"`
		NumRecords          *int32   `json:"num_records"`
		SamplingRatio       *float64 `json:"sampling_ratio"`
		Scope               *string  `json:"scope,omitempty"`
		Template            *string  `json:"template,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EvpQuery == nil {
		return fmt.Errorf("required field evp_query missing")
	}
	if all.HierarchyDepth == nil {
		return fmt.Errorf("required field hierarchy_depth missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.NumRecords == nil {
		return fmt.Errorf("required field num_records missing")
	}
	if all.SamplingRatio == nil {
		return fmt.Errorf("required field sampling_ratio missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "config_id", "evp_query", "hierarchy_depth", "integration_provider", "model_name", "name", "num_records", "sampling_ratio", "scope", "template"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.ConfigId = all.ConfigId
	o.EvpQuery = *all.EvpQuery
	o.HierarchyDepth = *all.HierarchyDepth
	o.IntegrationProvider = all.IntegrationProvider
	o.ModelName = all.ModelName
	o.Name = *all.Name
	o.NumRecords = *all.NumRecords
	o.SamplingRatio = *all.SamplingRatio
	o.Scope = all.Scope
	o.Template = all.Template

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
