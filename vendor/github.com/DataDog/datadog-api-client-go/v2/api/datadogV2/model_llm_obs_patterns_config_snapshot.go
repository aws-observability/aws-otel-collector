// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsConfigSnapshot Snapshot of the configuration used for a patterns run.
type LLMObsPatternsConfigSnapshot struct {
	// Integration account ID used for a bring-your-own-model run.
	AccountId *string `json:"account_id,omitempty"`
	// Query that selected the spans for the run.
	EvpQuery *string `json:"evp_query,omitempty"`
	// Depth of the topic hierarchy generated.
	HierarchyDepth *int32 `json:"hierarchy_depth,omitempty"`
	// Integration provider used for a bring-your-own-model run.
	IntegrationProvider *string `json:"integration_provider,omitempty"`
	// Model name used for a bring-your-own-model run.
	ModelName *string `json:"model_name,omitempty"`
	// Maximum number of records processed for the run.
	NumRecords *int32 `json:"num_records,omitempty"`
	// Fraction of matching spans sampled for the run.
	SamplingRatio *float64 `json:"sampling_ratio,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsConfigSnapshot instantiates a new LLMObsPatternsConfigSnapshot object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsConfigSnapshot() *LLMObsPatternsConfigSnapshot {
	this := LLMObsPatternsConfigSnapshot{}
	return &this
}

// NewLLMObsPatternsConfigSnapshotWithDefaults instantiates a new LLMObsPatternsConfigSnapshot object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsConfigSnapshotWithDefaults() *LLMObsPatternsConfigSnapshot {
	this := LLMObsPatternsConfigSnapshot{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigSnapshot) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigSnapshot) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigSnapshot) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *LLMObsPatternsConfigSnapshot) SetAccountId(v string) {
	o.AccountId = &v
}

// GetEvpQuery returns the EvpQuery field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigSnapshot) GetEvpQuery() string {
	if o == nil || o.EvpQuery == nil {
		var ret string
		return ret
	}
	return *o.EvpQuery
}

// GetEvpQueryOk returns a tuple with the EvpQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigSnapshot) GetEvpQueryOk() (*string, bool) {
	if o == nil || o.EvpQuery == nil {
		return nil, false
	}
	return o.EvpQuery, true
}

// HasEvpQuery returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigSnapshot) HasEvpQuery() bool {
	return o != nil && o.EvpQuery != nil
}

// SetEvpQuery gets a reference to the given string and assigns it to the EvpQuery field.
func (o *LLMObsPatternsConfigSnapshot) SetEvpQuery(v string) {
	o.EvpQuery = &v
}

// GetHierarchyDepth returns the HierarchyDepth field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigSnapshot) GetHierarchyDepth() int32 {
	if o == nil || o.HierarchyDepth == nil {
		var ret int32
		return ret
	}
	return *o.HierarchyDepth
}

// GetHierarchyDepthOk returns a tuple with the HierarchyDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigSnapshot) GetHierarchyDepthOk() (*int32, bool) {
	if o == nil || o.HierarchyDepth == nil {
		return nil, false
	}
	return o.HierarchyDepth, true
}

// HasHierarchyDepth returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigSnapshot) HasHierarchyDepth() bool {
	return o != nil && o.HierarchyDepth != nil
}

// SetHierarchyDepth gets a reference to the given int32 and assigns it to the HierarchyDepth field.
func (o *LLMObsPatternsConfigSnapshot) SetHierarchyDepth(v int32) {
	o.HierarchyDepth = &v
}

// GetIntegrationProvider returns the IntegrationProvider field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigSnapshot) GetIntegrationProvider() string {
	if o == nil || o.IntegrationProvider == nil {
		var ret string
		return ret
	}
	return *o.IntegrationProvider
}

// GetIntegrationProviderOk returns a tuple with the IntegrationProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigSnapshot) GetIntegrationProviderOk() (*string, bool) {
	if o == nil || o.IntegrationProvider == nil {
		return nil, false
	}
	return o.IntegrationProvider, true
}

// HasIntegrationProvider returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigSnapshot) HasIntegrationProvider() bool {
	return o != nil && o.IntegrationProvider != nil
}

// SetIntegrationProvider gets a reference to the given string and assigns it to the IntegrationProvider field.
func (o *LLMObsPatternsConfigSnapshot) SetIntegrationProvider(v string) {
	o.IntegrationProvider = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigSnapshot) GetModelName() string {
	if o == nil || o.ModelName == nil {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigSnapshot) GetModelNameOk() (*string, bool) {
	if o == nil || o.ModelName == nil {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigSnapshot) HasModelName() bool {
	return o != nil && o.ModelName != nil
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *LLMObsPatternsConfigSnapshot) SetModelName(v string) {
	o.ModelName = &v
}

// GetNumRecords returns the NumRecords field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigSnapshot) GetNumRecords() int32 {
	if o == nil || o.NumRecords == nil {
		var ret int32
		return ret
	}
	return *o.NumRecords
}

// GetNumRecordsOk returns a tuple with the NumRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigSnapshot) GetNumRecordsOk() (*int32, bool) {
	if o == nil || o.NumRecords == nil {
		return nil, false
	}
	return o.NumRecords, true
}

// HasNumRecords returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigSnapshot) HasNumRecords() bool {
	return o != nil && o.NumRecords != nil
}

// SetNumRecords gets a reference to the given int32 and assigns it to the NumRecords field.
func (o *LLMObsPatternsConfigSnapshot) SetNumRecords(v int32) {
	o.NumRecords = &v
}

// GetSamplingRatio returns the SamplingRatio field value if set, zero value otherwise.
func (o *LLMObsPatternsConfigSnapshot) GetSamplingRatio() float64 {
	if o == nil || o.SamplingRatio == nil {
		var ret float64
		return ret
	}
	return *o.SamplingRatio
}

// GetSamplingRatioOk returns a tuple with the SamplingRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsConfigSnapshot) GetSamplingRatioOk() (*float64, bool) {
	if o == nil || o.SamplingRatio == nil {
		return nil, false
	}
	return o.SamplingRatio, true
}

// HasSamplingRatio returns a boolean if a field has been set.
func (o *LLMObsPatternsConfigSnapshot) HasSamplingRatio() bool {
	return o != nil && o.SamplingRatio != nil
}

// SetSamplingRatio gets a reference to the given float64 and assigns it to the SamplingRatio field.
func (o *LLMObsPatternsConfigSnapshot) SetSamplingRatio(v float64) {
	o.SamplingRatio = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsConfigSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.EvpQuery != nil {
		toSerialize["evp_query"] = o.EvpQuery
	}
	if o.HierarchyDepth != nil {
		toSerialize["hierarchy_depth"] = o.HierarchyDepth
	}
	if o.IntegrationProvider != nil {
		toSerialize["integration_provider"] = o.IntegrationProvider
	}
	if o.ModelName != nil {
		toSerialize["model_name"] = o.ModelName
	}
	if o.NumRecords != nil {
		toSerialize["num_records"] = o.NumRecords
	}
	if o.SamplingRatio != nil {
		toSerialize["sampling_ratio"] = o.SamplingRatio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsConfigSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId           *string  `json:"account_id,omitempty"`
		EvpQuery            *string  `json:"evp_query,omitempty"`
		HierarchyDepth      *int32   `json:"hierarchy_depth,omitempty"`
		IntegrationProvider *string  `json:"integration_provider,omitempty"`
		ModelName           *string  `json:"model_name,omitempty"`
		NumRecords          *int32   `json:"num_records,omitempty"`
		SamplingRatio       *float64 `json:"sampling_ratio,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "evp_query", "hierarchy_depth", "integration_provider", "model_name", "num_records", "sampling_ratio"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.EvpQuery = all.EvpQuery
	o.HierarchyDepth = all.HierarchyDepth
	o.IntegrationProvider = all.IntegrationProvider
	o.ModelName = all.ModelName
	o.NumRecords = all.NumRecords
	o.SamplingRatio = all.SamplingRatio

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
