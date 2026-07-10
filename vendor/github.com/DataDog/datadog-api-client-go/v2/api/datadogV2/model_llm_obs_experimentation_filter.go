// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationFilter Filter criteria for an experimentation search request.
type LLMObsExperimentationFilter struct {
	// When `true`, include soft-deleted entities alongside active ones.
	IncludeDeleted *bool `json:"include_deleted,omitempty"`
	// When `true`, return only soft-deleted entities.
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Free-text search query.
	Query *string `json:"query,omitempty"`
	// Entity types to search. Valid values are `projects`, `datasets`, `dataset_records`, `experiments`, and `experiment_runs`.
	Scope []string `json:"scope"`
	// Filter dataset records by a specific dataset version.
	Version datadog.NullableInt64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationFilter instantiates a new LLMObsExperimentationFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationFilter(scope []string) *LLMObsExperimentationFilter {
	this := LLMObsExperimentationFilter{}
	var includeDeleted bool = false
	this.IncludeDeleted = &includeDeleted
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	this.Scope = scope
	return &this
}

// NewLLMObsExperimentationFilterWithDefaults instantiates a new LLMObsExperimentationFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationFilterWithDefaults() *LLMObsExperimentationFilter {
	this := LLMObsExperimentationFilter{}
	var includeDeleted bool = false
	this.IncludeDeleted = &includeDeleted
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// GetIncludeDeleted returns the IncludeDeleted field value if set, zero value otherwise.
func (o *LLMObsExperimentationFilter) GetIncludeDeleted() bool {
	if o == nil || o.IncludeDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IncludeDeleted
}

// GetIncludeDeletedOk returns a tuple with the IncludeDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationFilter) GetIncludeDeletedOk() (*bool, bool) {
	if o == nil || o.IncludeDeleted == nil {
		return nil, false
	}
	return o.IncludeDeleted, true
}

// HasIncludeDeleted returns a boolean if a field has been set.
func (o *LLMObsExperimentationFilter) HasIncludeDeleted() bool {
	return o != nil && o.IncludeDeleted != nil
}

// SetIncludeDeleted gets a reference to the given bool and assigns it to the IncludeDeleted field.
func (o *LLMObsExperimentationFilter) SetIncludeDeleted(v bool) {
	o.IncludeDeleted = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *LLMObsExperimentationFilter) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationFilter) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *LLMObsExperimentationFilter) HasIsDeleted() bool {
	return o != nil && o.IsDeleted != nil
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *LLMObsExperimentationFilter) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *LLMObsExperimentationFilter) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationFilter) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *LLMObsExperimentationFilter) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *LLMObsExperimentationFilter) SetQuery(v string) {
	o.Query = &v
}

// GetScope returns the Scope field value.
func (o *LLMObsExperimentationFilter) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationFilter) GetScopeOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *LLMObsExperimentationFilter) SetScope(v []string) {
	o.Scope = v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentationFilter) GetVersion() int64 {
	if o == nil || o.Version.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentationFilter) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *LLMObsExperimentationFilter) HasVersion() bool {
	return o != nil && o.Version.IsSet()
}

// SetVersion gets a reference to the given datadog.NullableInt64 and assigns it to the Version field.
func (o *LLMObsExperimentationFilter) SetVersion(v int64) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil.
func (o *LLMObsExperimentationFilter) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil.
func (o *LLMObsExperimentationFilter) UnsetVersion() {
	o.Version.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludeDeleted != nil {
		toSerialize["include_deleted"] = o.IncludeDeleted
	}
	if o.IsDeleted != nil {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	toSerialize["scope"] = o.Scope
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeDeleted *bool                 `json:"include_deleted,omitempty"`
		IsDeleted      *bool                 `json:"is_deleted,omitempty"`
		Query          *string               `json:"query,omitempty"`
		Scope          *[]string             `json:"scope"`
		Version        datadog.NullableInt64 `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_deleted", "is_deleted", "query", "scope", "version"})
	} else {
		return err
	}
	o.IncludeDeleted = all.IncludeDeleted
	o.IsDeleted = all.IsDeleted
	o.Query = all.Query
	o.Scope = *all.Scope
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
