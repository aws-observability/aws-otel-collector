// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationsFilterRequest Request body for filtering cost recommendations.
type RecommendationsFilterRequest struct {
	// Filter expression applied to the recommendations.
	Filter *string `json:"filter,omitempty"`
	// Ordered list of sort clauses applied to the result set.
	Sort []RecommendationsFilterRequestSortItems `json:"sort,omitempty"`
	// Active view name (for example, `active`, `dismissed`, `open`, `in-progress`, or `completed`).
	View *string `json:"view,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRecommendationsFilterRequest instantiates a new RecommendationsFilterRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendationsFilterRequest() *RecommendationsFilterRequest {
	this := RecommendationsFilterRequest{}
	return &this
}

// NewRecommendationsFilterRequestWithDefaults instantiates a new RecommendationsFilterRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecommendationsFilterRequestWithDefaults() *RecommendationsFilterRequest {
	this := RecommendationsFilterRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RecommendationsFilterRequest) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsFilterRequest) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RecommendationsFilterRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *RecommendationsFilterRequest) SetFilter(v string) {
	o.Filter = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *RecommendationsFilterRequest) GetSort() []RecommendationsFilterRequestSortItems {
	if o == nil || o.Sort == nil {
		var ret []RecommendationsFilterRequestSortItems
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsFilterRequest) GetSortOk() (*[]RecommendationsFilterRequestSortItems, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return &o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *RecommendationsFilterRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given []RecommendationsFilterRequestSortItems and assigns it to the Sort field.
func (o *RecommendationsFilterRequest) SetSort(v []RecommendationsFilterRequestSortItems) {
	o.Sort = v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *RecommendationsFilterRequest) GetView() string {
	if o == nil || o.View == nil {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsFilterRequest) GetViewOk() (*string, bool) {
	if o == nil || o.View == nil {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *RecommendationsFilterRequest) HasView() bool {
	return o != nil && o.View != nil
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *RecommendationsFilterRequest) SetView(v string) {
	o.View = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecommendationsFilterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.View != nil {
		toSerialize["view"] = o.View
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RecommendationsFilterRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter *string                                 `json:"filter,omitempty"`
		Sort   []RecommendationsFilterRequestSortItems `json:"sort,omitempty"`
		View   *string                                 `json:"view,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "sort", "view"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.Sort = all.Sort
	o.View = all.View

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
