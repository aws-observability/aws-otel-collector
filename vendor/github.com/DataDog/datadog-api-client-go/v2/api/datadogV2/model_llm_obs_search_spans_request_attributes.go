// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSearchSpansRequestAttributes Attributes of an LLM Observability spans search request.
type LLMObsSearchSpansRequestAttributes struct {
	// Filter criteria for an LLM Observability span search.
	Filter *LLMObsSpanFilter `json:"filter,omitempty"`
	// Additional options for a span search request.
	Options *LLMObsSpanSearchOptions `json:"options,omitempty"`
	// Pagination settings for a span search request.
	Page *LLMObsSpanPageQuery `json:"page,omitempty"`
	// Sort order for the results. Use `-` prefix for descending order.
	Sort *string `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSearchSpansRequestAttributes instantiates a new LLMObsSearchSpansRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSearchSpansRequestAttributes() *LLMObsSearchSpansRequestAttributes {
	this := LLMObsSearchSpansRequestAttributes{}
	return &this
}

// NewLLMObsSearchSpansRequestAttributesWithDefaults instantiates a new LLMObsSearchSpansRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSearchSpansRequestAttributesWithDefaults() *LLMObsSearchSpansRequestAttributes {
	this := LLMObsSearchSpansRequestAttributes{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *LLMObsSearchSpansRequestAttributes) GetFilter() LLMObsSpanFilter {
	if o == nil || o.Filter == nil {
		var ret LLMObsSpanFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSearchSpansRequestAttributes) GetFilterOk() (*LLMObsSpanFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *LLMObsSearchSpansRequestAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given LLMObsSpanFilter and assigns it to the Filter field.
func (o *LLMObsSearchSpansRequestAttributes) SetFilter(v LLMObsSpanFilter) {
	o.Filter = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *LLMObsSearchSpansRequestAttributes) GetOptions() LLMObsSpanSearchOptions {
	if o == nil || o.Options == nil {
		var ret LLMObsSpanSearchOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSearchSpansRequestAttributes) GetOptionsOk() (*LLMObsSpanSearchOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *LLMObsSearchSpansRequestAttributes) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given LLMObsSpanSearchOptions and assigns it to the Options field.
func (o *LLMObsSearchSpansRequestAttributes) SetOptions(v LLMObsSpanSearchOptions) {
	o.Options = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *LLMObsSearchSpansRequestAttributes) GetPage() LLMObsSpanPageQuery {
	if o == nil || o.Page == nil {
		var ret LLMObsSpanPageQuery
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSearchSpansRequestAttributes) GetPageOk() (*LLMObsSpanPageQuery, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *LLMObsSearchSpansRequestAttributes) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given LLMObsSpanPageQuery and assigns it to the Page field.
func (o *LLMObsSearchSpansRequestAttributes) SetPage(v LLMObsSpanPageQuery) {
	o.Page = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *LLMObsSearchSpansRequestAttributes) GetSort() string {
	if o == nil || o.Sort == nil {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSearchSpansRequestAttributes) GetSortOk() (*string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *LLMObsSearchSpansRequestAttributes) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *LLMObsSearchSpansRequestAttributes) SetSort(v string) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSearchSpansRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsSearchSpansRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter  *LLMObsSpanFilter        `json:"filter,omitempty"`
		Options *LLMObsSpanSearchOptions `json:"options,omitempty"`
		Page    *LLMObsSpanPageQuery     `json:"page,omitempty"`
		Sort    *string                  `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "options", "page", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page
	o.Sort = all.Sort

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
