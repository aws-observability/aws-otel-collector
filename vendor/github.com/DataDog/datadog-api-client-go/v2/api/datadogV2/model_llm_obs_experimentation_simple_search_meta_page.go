// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationSimpleSearchMetaPage Page metadata.
type LLMObsExperimentationSimpleSearchMetaPage struct {
	// Current page number.
	Current *int32 `json:"current,omitempty"`
	// Page size used for this response.
	Limit *int32 `json:"limit,omitempty"`
	// Total number of matching results (capped at the maximum search limit).
	TotalCount *int32 `json:"total_count,omitempty"`
	// Total number of pages available.
	TotalPages *int32 `json:"total_pages,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationSimpleSearchMetaPage instantiates a new LLMObsExperimentationSimpleSearchMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationSimpleSearchMetaPage() *LLMObsExperimentationSimpleSearchMetaPage {
	this := LLMObsExperimentationSimpleSearchMetaPage{}
	return &this
}

// NewLLMObsExperimentationSimpleSearchMetaPageWithDefaults instantiates a new LLMObsExperimentationSimpleSearchMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationSimpleSearchMetaPageWithDefaults() *LLMObsExperimentationSimpleSearchMetaPage {
	this := LLMObsExperimentationSimpleSearchMetaPage{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *LLMObsExperimentationSimpleSearchMetaPage) GetCurrent() int32 {
	if o == nil || o.Current == nil {
		var ret int32
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSimpleSearchMetaPage) GetCurrentOk() (*int32, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *LLMObsExperimentationSimpleSearchMetaPage) HasCurrent() bool {
	return o != nil && o.Current != nil
}

// SetCurrent gets a reference to the given int32 and assigns it to the Current field.
func (o *LLMObsExperimentationSimpleSearchMetaPage) SetCurrent(v int32) {
	o.Current = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *LLMObsExperimentationSimpleSearchMetaPage) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSimpleSearchMetaPage) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *LLMObsExperimentationSimpleSearchMetaPage) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *LLMObsExperimentationSimpleSearchMetaPage) SetLimit(v int32) {
	o.Limit = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *LLMObsExperimentationSimpleSearchMetaPage) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSimpleSearchMetaPage) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *LLMObsExperimentationSimpleSearchMetaPage) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *LLMObsExperimentationSimpleSearchMetaPage) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *LLMObsExperimentationSimpleSearchMetaPage) GetTotalPages() int32 {
	if o == nil || o.TotalPages == nil {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSimpleSearchMetaPage) GetTotalPagesOk() (*int32, bool) {
	if o == nil || o.TotalPages == nil {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *LLMObsExperimentationSimpleSearchMetaPage) HasTotalPages() bool {
	return o != nil && o.TotalPages != nil
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *LLMObsExperimentationSimpleSearchMetaPage) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationSimpleSearchMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	if o.TotalPages != nil {
		toSerialize["total_pages"] = o.TotalPages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationSimpleSearchMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Current    *int32 `json:"current,omitempty"`
		Limit      *int32 `json:"limit,omitempty"`
		TotalCount *int32 `json:"total_count,omitempty"`
		TotalPages *int32 `json:"total_pages,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"current", "limit", "total_count", "total_pages"})
	} else {
		return err
	}
	o.Current = all.Current
	o.Limit = all.Limit
	o.TotalCount = all.TotalCount
	o.TotalPages = all.TotalPages

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
