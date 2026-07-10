// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsConfigAttributes Attributes of an LLM Observability patterns configuration.
type LLMObsPatternsConfigAttributes struct {
	// Integration account ID for a bring-your-own-model configuration.
	AccountId datadog.NullableString `json:"account_id,omitempty"`
	// Timestamp when the configuration was created.
	CreatedAt time.Time `json:"created_at"`
	// Query that selects the spans the patterns run analyzes.
	EvpQuery string `json:"evp_query"`
	// Depth of the topic hierarchy to generate.
	HierarchyDepth int32 `json:"hierarchy_depth"`
	// Integration provider for a bring-your-own-model configuration.
	IntegrationProvider datadog.NullableString `json:"integration_provider,omitempty"`
	// Model name for a bring-your-own-model configuration.
	ModelName datadog.NullableString `json:"model_name,omitempty"`
	// Name of the configuration.
	Name string `json:"name"`
	// Maximum number of records to process for the run.
	NumRecords int32 `json:"num_records"`
	// Fraction of matching spans to sample for the run.
	SamplingRatio float64 `json:"sampling_ratio"`
	// Scope of the configuration.
	Scope string `json:"scope"`
	// Template used to guide topic generation.
	Template datadog.NullableString `json:"template,omitempty"`
	// Timestamp when the configuration was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsConfigAttributes instantiates a new LLMObsPatternsConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsConfigAttributes(createdAt time.Time, evpQuery string, hierarchyDepth int32, name string, numRecords int32, samplingRatio float64, scope string, updatedAt time.Time) *LLMObsPatternsConfigAttributes {
	this := LLMObsPatternsConfigAttributes{}
	this.CreatedAt = createdAt
	this.EvpQuery = evpQuery
	this.HierarchyDepth = hierarchyDepth
	this.Name = name
	this.NumRecords = numRecords
	this.SamplingRatio = samplingRatio
	this.Scope = scope
	this.UpdatedAt = updatedAt
	return &this
}

// NewLLMObsPatternsConfigAttributesWithDefaults instantiates a new LLMObsPatternsConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsConfigAttributesWithDefaults() *LLMObsPatternsConfigAttributes {
	this := LLMObsPatternsConfigAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsPatternsConfigAttributes) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsPatternsConfigAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigAttributes) HasAccountId() bool {
	return o != nil && o.AccountId.IsSet()
}

// SetAccountId gets a reference to the given datadog.NullableString and assigns it to the AccountId field.
func (o *LLMObsPatternsConfigAttributes) SetAccountId(v string) {
	o.AccountId.Set(&v)
}

// SetAccountIdNil sets the value for AccountId to be an explicit nil.
func (o *LLMObsPatternsConfigAttributes) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil.
func (o *LLMObsPatternsConfigAttributes) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsPatternsConfigAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsPatternsConfigAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEvpQuery returns the EvpQuery field value.
func (o *LLMObsPatternsConfigAttributes) GetEvpQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvpQuery
}

// GetEvpQueryOk returns a tuple with the EvpQuery field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigAttributes) GetEvpQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvpQuery, true
}

// SetEvpQuery sets field value.
func (o *LLMObsPatternsConfigAttributes) SetEvpQuery(v string) {
	o.EvpQuery = v
}

// GetHierarchyDepth returns the HierarchyDepth field value.
func (o *LLMObsPatternsConfigAttributes) GetHierarchyDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.HierarchyDepth
}

// GetHierarchyDepthOk returns a tuple with the HierarchyDepth field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigAttributes) GetHierarchyDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HierarchyDepth, true
}

// SetHierarchyDepth sets field value.
func (o *LLMObsPatternsConfigAttributes) SetHierarchyDepth(v int32) {
	o.HierarchyDepth = v
}

// GetIntegrationProvider returns the IntegrationProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsPatternsConfigAttributes) GetIntegrationProvider() string {
	if o == nil || o.IntegrationProvider.Get() == nil {
		var ret string
		return ret
	}
	return *o.IntegrationProvider.Get()
}

// GetIntegrationProviderOk returns a tuple with the IntegrationProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsPatternsConfigAttributes) GetIntegrationProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntegrationProvider.Get(), o.IntegrationProvider.IsSet()
}

// HasIntegrationProvider returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigAttributes) HasIntegrationProvider() bool {
	return o != nil && o.IntegrationProvider.IsSet()
}

// SetIntegrationProvider gets a reference to the given datadog.NullableString and assigns it to the IntegrationProvider field.
func (o *LLMObsPatternsConfigAttributes) SetIntegrationProvider(v string) {
	o.IntegrationProvider.Set(&v)
}

// SetIntegrationProviderNil sets the value for IntegrationProvider to be an explicit nil.
func (o *LLMObsPatternsConfigAttributes) SetIntegrationProviderNil() {
	o.IntegrationProvider.Set(nil)
}

// UnsetIntegrationProvider ensures that no value is present for IntegrationProvider, not even an explicit nil.
func (o *LLMObsPatternsConfigAttributes) UnsetIntegrationProvider() {
	o.IntegrationProvider.Unset()
}

// GetModelName returns the ModelName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsPatternsConfigAttributes) GetModelName() string {
	if o == nil || o.ModelName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModelName.Get()
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsPatternsConfigAttributes) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModelName.Get(), o.ModelName.IsSet()
}

// HasModelName returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigAttributes) HasModelName() bool {
	return o != nil && o.ModelName.IsSet()
}

// SetModelName gets a reference to the given datadog.NullableString and assigns it to the ModelName field.
func (o *LLMObsPatternsConfigAttributes) SetModelName(v string) {
	o.ModelName.Set(&v)
}

// SetModelNameNil sets the value for ModelName to be an explicit nil.
func (o *LLMObsPatternsConfigAttributes) SetModelNameNil() {
	o.ModelName.Set(nil)
}

// UnsetModelName ensures that no value is present for ModelName, not even an explicit nil.
func (o *LLMObsPatternsConfigAttributes) UnsetModelName() {
	o.ModelName.Unset()
}

// GetName returns the Name field value.
func (o *LLMObsPatternsConfigAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsPatternsConfigAttributes) SetName(v string) {
	o.Name = v
}

// GetNumRecords returns the NumRecords field value.
func (o *LLMObsPatternsConfigAttributes) GetNumRecords() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.NumRecords
}

// GetNumRecordsOk returns a tuple with the NumRecords field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigAttributes) GetNumRecordsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumRecords, true
}

// SetNumRecords sets field value.
func (o *LLMObsPatternsConfigAttributes) SetNumRecords(v int32) {
	o.NumRecords = v
}

// GetSamplingRatio returns the SamplingRatio field value.
func (o *LLMObsPatternsConfigAttributes) GetSamplingRatio() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.SamplingRatio
}

// GetSamplingRatioOk returns a tuple with the SamplingRatio field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigAttributes) GetSamplingRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SamplingRatio, true
}

// SetSamplingRatio sets field value.
func (o *LLMObsPatternsConfigAttributes) SetSamplingRatio(v float64) {
	o.SamplingRatio = v
}

// GetScope returns the Scope field value.
func (o *LLMObsPatternsConfigAttributes) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigAttributes) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *LLMObsPatternsConfigAttributes) SetScope(v string) {
	o.Scope = v
}

// GetTemplate returns the Template field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsPatternsConfigAttributes) GetTemplate() string {
	if o == nil || o.Template.Get() == nil {
		var ret string
		return ret
	}
	return *o.Template.Get()
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsPatternsConfigAttributes) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Template.Get(), o.Template.IsSet()
}

// HasTemplate returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigAttributes) HasTemplate() bool {
	return o != nil && o.Template.IsSet()
}

// SetTemplate gets a reference to the given datadog.NullableString and assigns it to the Template field.
func (o *LLMObsPatternsConfigAttributes) SetTemplate(v string) {
	o.Template.Set(&v)
}

// SetTemplateNil sets the value for Template to be an explicit nil.
func (o *LLMObsPatternsConfigAttributes) SetTemplateNil() {
	o.Template.Set(nil)
}

// UnsetTemplate ensures that no value is present for Template, not even an explicit nil.
func (o *LLMObsPatternsConfigAttributes) UnsetTemplate() {
	o.Template.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *LLMObsPatternsConfigAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *LLMObsPatternsConfigAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["evp_query"] = o.EvpQuery
	toSerialize["hierarchy_depth"] = o.HierarchyDepth
	if o.IntegrationProvider.IsSet() {
		toSerialize["integration_provider"] = o.IntegrationProvider.Get()
	}
	if o.ModelName.IsSet() {
		toSerialize["model_name"] = o.ModelName.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["num_records"] = o.NumRecords
	toSerialize["sampling_ratio"] = o.SamplingRatio
	toSerialize["scope"] = o.Scope
	if o.Template.IsSet() {
		toSerialize["template"] = o.Template.Get()
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId           datadog.NullableString `json:"account_id,omitempty"`
		CreatedAt           *time.Time             `json:"created_at"`
		EvpQuery            *string                `json:"evp_query"`
		HierarchyDepth      *int32                 `json:"hierarchy_depth"`
		IntegrationProvider datadog.NullableString `json:"integration_provider,omitempty"`
		ModelName           datadog.NullableString `json:"model_name,omitempty"`
		Name                *string                `json:"name"`
		NumRecords          *int32                 `json:"num_records"`
		SamplingRatio       *float64               `json:"sampling_ratio"`
		Scope               *string                `json:"scope"`
		Template            datadog.NullableString `json:"template,omitempty"`
		UpdatedAt           *time.Time             `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
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
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "created_at", "evp_query", "hierarchy_depth", "integration_provider", "model_name", "name", "num_records", "sampling_ratio", "scope", "template", "updated_at"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.CreatedAt = *all.CreatedAt
	o.EvpQuery = *all.EvpQuery
	o.HierarchyDepth = *all.HierarchyDepth
	o.IntegrationProvider = all.IntegrationProvider
	o.ModelName = all.ModelName
	o.Name = *all.Name
	o.NumRecords = *all.NumRecords
	o.SamplingRatio = *all.SamplingRatio
	o.Scope = *all.Scope
	o.Template = all.Template
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
