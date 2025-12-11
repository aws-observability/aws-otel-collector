// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestsSearchRequestAttributes Attributes for the flaky tests search request.
type FlakyTestsSearchRequestAttributes struct {
	// Search filter settings.
	Filter *FlakyTestsSearchFilter `json:"filter,omitempty"`
	// Pagination attributes for listing flaky tests.
	Page *FlakyTestsSearchPageOptions `json:"page,omitempty"`
	// Parameter for sorting flaky test results. The default sort is by ascending Fully Qualified Name (FQN). The FQN is the concatenation of the test module, suite, and name.
	Sort *FlakyTestsSearchSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlakyTestsSearchRequestAttributes instantiates a new FlakyTestsSearchRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlakyTestsSearchRequestAttributes() *FlakyTestsSearchRequestAttributes {
	this := FlakyTestsSearchRequestAttributes{}
	return &this
}

// NewFlakyTestsSearchRequestAttributesWithDefaults instantiates a new FlakyTestsSearchRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlakyTestsSearchRequestAttributesWithDefaults() *FlakyTestsSearchRequestAttributes {
	this := FlakyTestsSearchRequestAttributes{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *FlakyTestsSearchRequestAttributes) GetFilter() FlakyTestsSearchFilter {
	if o == nil || o.Filter == nil {
		var ret FlakyTestsSearchFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestsSearchRequestAttributes) GetFilterOk() (*FlakyTestsSearchFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *FlakyTestsSearchRequestAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given FlakyTestsSearchFilter and assigns it to the Filter field.
func (o *FlakyTestsSearchRequestAttributes) SetFilter(v FlakyTestsSearchFilter) {
	o.Filter = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *FlakyTestsSearchRequestAttributes) GetPage() FlakyTestsSearchPageOptions {
	if o == nil || o.Page == nil {
		var ret FlakyTestsSearchPageOptions
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestsSearchRequestAttributes) GetPageOk() (*FlakyTestsSearchPageOptions, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *FlakyTestsSearchRequestAttributes) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given FlakyTestsSearchPageOptions and assigns it to the Page field.
func (o *FlakyTestsSearchRequestAttributes) SetPage(v FlakyTestsSearchPageOptions) {
	o.Page = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *FlakyTestsSearchRequestAttributes) GetSort() FlakyTestsSearchSort {
	if o == nil || o.Sort == nil {
		var ret FlakyTestsSearchSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlakyTestsSearchRequestAttributes) GetSortOk() (*FlakyTestsSearchSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *FlakyTestsSearchRequestAttributes) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given FlakyTestsSearchSort and assigns it to the Sort field.
func (o *FlakyTestsSearchRequestAttributes) SetSort(v FlakyTestsSearchSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FlakyTestsSearchRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
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
func (o *FlakyTestsSearchRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter *FlakyTestsSearchFilter      `json:"filter,omitempty"`
		Page   *FlakyTestsSearchPageOptions `json:"page,omitempty"`
		Sort   *FlakyTestsSearchSort        `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "page", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page
	if all.Sort != nil && !all.Sort.IsValid() {
		hasInvalidField = true
	} else {
		o.Sort = all.Sort
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
