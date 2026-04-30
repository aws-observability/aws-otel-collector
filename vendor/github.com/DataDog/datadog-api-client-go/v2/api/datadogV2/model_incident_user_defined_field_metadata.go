// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldMetadata Metadata for autocomplete-type user-defined fields, describing how to populate autocomplete options.
type IncidentUserDefinedFieldMetadata struct {
	// The category of the autocomplete source.
	Category string `json:"category"`
	// The query parameter used to limit the number of autocomplete results.
	SearchLimitParam string `json:"search_limit_param"`
	// Additional query parameters to include in the search URL.
	SearchParams map[string]interface{} `json:"search_params"`
	// The query parameter used to pass typed input to the search URL.
	SearchQueryParam string `json:"search_query_param"`
	// The JSON path to the results in the response body.
	SearchResultPath string `json:"search_result_path"`
	// The URL used to populate autocomplete options.
	SearchUrl string `json:"search_url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedFieldMetadata instantiates a new IncidentUserDefinedFieldMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedFieldMetadata(category string, searchLimitParam string, searchParams map[string]interface{}, searchQueryParam string, searchResultPath string, searchUrl string) *IncidentUserDefinedFieldMetadata {
	this := IncidentUserDefinedFieldMetadata{}
	this.Category = category
	this.SearchLimitParam = searchLimitParam
	this.SearchParams = searchParams
	this.SearchQueryParam = searchQueryParam
	this.SearchResultPath = searchResultPath
	this.SearchUrl = searchUrl
	return &this
}

// NewIncidentUserDefinedFieldMetadataWithDefaults instantiates a new IncidentUserDefinedFieldMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedFieldMetadataWithDefaults() *IncidentUserDefinedFieldMetadata {
	this := IncidentUserDefinedFieldMetadata{}
	return &this
}

// GetCategory returns the Category field value.
func (o *IncidentUserDefinedFieldMetadata) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldMetadata) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *IncidentUserDefinedFieldMetadata) SetCategory(v string) {
	o.Category = v
}

// GetSearchLimitParam returns the SearchLimitParam field value.
func (o *IncidentUserDefinedFieldMetadata) GetSearchLimitParam() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SearchLimitParam
}

// GetSearchLimitParamOk returns a tuple with the SearchLimitParam field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldMetadata) GetSearchLimitParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchLimitParam, true
}

// SetSearchLimitParam sets field value.
func (o *IncidentUserDefinedFieldMetadata) SetSearchLimitParam(v string) {
	o.SearchLimitParam = v
}

// GetSearchParams returns the SearchParams field value.
func (o *IncidentUserDefinedFieldMetadata) GetSearchParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SearchParams
}

// GetSearchParamsOk returns a tuple with the SearchParams field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldMetadata) GetSearchParamsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchParams, true
}

// SetSearchParams sets field value.
func (o *IncidentUserDefinedFieldMetadata) SetSearchParams(v map[string]interface{}) {
	o.SearchParams = v
}

// GetSearchQueryParam returns the SearchQueryParam field value.
func (o *IncidentUserDefinedFieldMetadata) GetSearchQueryParam() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SearchQueryParam
}

// GetSearchQueryParamOk returns a tuple with the SearchQueryParam field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldMetadata) GetSearchQueryParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchQueryParam, true
}

// SetSearchQueryParam sets field value.
func (o *IncidentUserDefinedFieldMetadata) SetSearchQueryParam(v string) {
	o.SearchQueryParam = v
}

// GetSearchResultPath returns the SearchResultPath field value.
func (o *IncidentUserDefinedFieldMetadata) GetSearchResultPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SearchResultPath
}

// GetSearchResultPathOk returns a tuple with the SearchResultPath field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldMetadata) GetSearchResultPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchResultPath, true
}

// SetSearchResultPath sets field value.
func (o *IncidentUserDefinedFieldMetadata) SetSearchResultPath(v string) {
	o.SearchResultPath = v
}

// GetSearchUrl returns the SearchUrl field value.
func (o *IncidentUserDefinedFieldMetadata) GetSearchUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SearchUrl
}

// GetSearchUrlOk returns a tuple with the SearchUrl field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldMetadata) GetSearchUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchUrl, true
}

// SetSearchUrl sets field value.
func (o *IncidentUserDefinedFieldMetadata) SetSearchUrl(v string) {
	o.SearchUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedFieldMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["search_limit_param"] = o.SearchLimitParam
	toSerialize["search_params"] = o.SearchParams
	toSerialize["search_query_param"] = o.SearchQueryParam
	toSerialize["search_result_path"] = o.SearchResultPath
	toSerialize["search_url"] = o.SearchUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedFieldMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category         *string                 `json:"category"`
		SearchLimitParam *string                 `json:"search_limit_param"`
		SearchParams     *map[string]interface{} `json:"search_params"`
		SearchQueryParam *string                 `json:"search_query_param"`
		SearchResultPath *string                 `json:"search_result_path"`
		SearchUrl        *string                 `json:"search_url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.SearchLimitParam == nil {
		return fmt.Errorf("required field search_limit_param missing")
	}
	if all.SearchParams == nil {
		return fmt.Errorf("required field search_params missing")
	}
	if all.SearchQueryParam == nil {
		return fmt.Errorf("required field search_query_param missing")
	}
	if all.SearchResultPath == nil {
		return fmt.Errorf("required field search_result_path missing")
	}
	if all.SearchUrl == nil {
		return fmt.Errorf("required field search_url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "search_limit_param", "search_params", "search_query_param", "search_result_path", "search_url"})
	} else {
		return err
	}
	o.Category = *all.Category
	o.SearchLimitParam = *all.SearchLimitParam
	o.SearchParams = *all.SearchParams
	o.SearchQueryParam = *all.SearchQueryParam
	o.SearchResultPath = *all.SearchResultPath
	o.SearchUrl = *all.SearchUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableIncidentUserDefinedFieldMetadata handles when a null is used for IncidentUserDefinedFieldMetadata.
type NullableIncidentUserDefinedFieldMetadata struct {
	value *IncidentUserDefinedFieldMetadata
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentUserDefinedFieldMetadata) Get() *IncidentUserDefinedFieldMetadata {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentUserDefinedFieldMetadata) Set(val *IncidentUserDefinedFieldMetadata) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentUserDefinedFieldMetadata) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableIncidentUserDefinedFieldMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentUserDefinedFieldMetadata initializes the struct as if Set has been called.
func NewNullableIncidentUserDefinedFieldMetadata(val *IncidentUserDefinedFieldMetadata) *NullableIncidentUserDefinedFieldMetadata {
	return &NullableIncidentUserDefinedFieldMetadata{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentUserDefinedFieldMetadata) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentUserDefinedFieldMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
