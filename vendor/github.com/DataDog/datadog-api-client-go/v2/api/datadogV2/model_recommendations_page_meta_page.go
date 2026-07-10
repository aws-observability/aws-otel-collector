// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationsPageMetaPage Pagination metadata for a page of cost recommendations.
type RecommendationsPageMetaPage struct {
	// The filter expression that was applied to produce this page.
	Filter *string `json:"filter,omitempty"`
	// Opaque token used to fetch the next page; absent on the last page.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Number of items returned in this page (1–10000).
	PageSize *int32 `json:"page_size,omitempty"`
	// Pagination token echoed back from the request.
	PageToken *string `json:"page_token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRecommendationsPageMetaPage instantiates a new RecommendationsPageMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendationsPageMetaPage() *RecommendationsPageMetaPage {
	this := RecommendationsPageMetaPage{}
	return &this
}

// NewRecommendationsPageMetaPageWithDefaults instantiates a new RecommendationsPageMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecommendationsPageMetaPageWithDefaults() *RecommendationsPageMetaPage {
	this := RecommendationsPageMetaPage{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RecommendationsPageMetaPage) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsPageMetaPage) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RecommendationsPageMetaPage) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *RecommendationsPageMetaPage) SetFilter(v string) {
	o.Filter = &v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *RecommendationsPageMetaPage) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsPageMetaPage) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *RecommendationsPageMetaPage) HasNextPageToken() bool {
	return o != nil && o.NextPageToken != nil
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *RecommendationsPageMetaPage) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *RecommendationsPageMetaPage) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsPageMetaPage) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *RecommendationsPageMetaPage) HasPageSize() bool {
	return o != nil && o.PageSize != nil
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *RecommendationsPageMetaPage) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *RecommendationsPageMetaPage) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsPageMetaPage) GetPageTokenOk() (*string, bool) {
	if o == nil || o.PageToken == nil {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *RecommendationsPageMetaPage) HasPageToken() bool {
	return o != nil && o.PageToken != nil
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *RecommendationsPageMetaPage) SetPageToken(v string) {
	o.PageToken = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecommendationsPageMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RecommendationsPageMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter        *string `json:"filter,omitempty"`
		NextPageToken *string `json:"next_page_token,omitempty"`
		PageSize      *int32  `json:"page_size,omitempty"`
		PageToken     *string `json:"page_token,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "next_page_token", "page_size", "page_token"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.NextPageToken = all.NextPageToken
	o.PageSize = all.PageSize
	o.PageToken = all.PageToken

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
