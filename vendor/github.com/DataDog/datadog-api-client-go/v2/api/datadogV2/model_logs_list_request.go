// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsListRequest The request for a logs list.
type LogsListRequest struct {
	// The search and filter query settings
	Filter *LogsQueryFilter `json:"filter,omitempty"`
	// Global query options that are used during the query.
	// Note: These fields are currently deprecated and do not affect the query results.
	// Deprecated
	Options *LogsQueryOptions `json:"options,omitempty"`
	// Paging attributes for listing logs.
	Page *LogsListRequestPage `json:"page,omitempty"`
	// Sort parameters when querying logs.
	Sort *LogsSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsListRequest instantiates a new LogsListRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsListRequest() *LogsListRequest {
	this := LogsListRequest{}
	return &this
}

// NewLogsListRequestWithDefaults instantiates a new LogsListRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsListRequestWithDefaults() *LogsListRequest {
	this := LogsListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *LogsListRequest) GetFilter() LogsQueryFilter {
	if o == nil || o.Filter == nil {
		var ret LogsQueryFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsListRequest) GetFilterOk() (*LogsQueryFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *LogsListRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given LogsQueryFilter and assigns it to the Filter field.
func (o *LogsListRequest) SetFilter(v LogsQueryFilter) {
	o.Filter = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
// Deprecated
func (o *LogsListRequest) GetOptions() LogsQueryOptions {
	if o == nil || o.Options == nil {
		var ret LogsQueryOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LogsListRequest) GetOptionsOk() (*LogsQueryOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *LogsListRequest) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given LogsQueryOptions and assigns it to the Options field.
// Deprecated
func (o *LogsListRequest) SetOptions(v LogsQueryOptions) {
	o.Options = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *LogsListRequest) GetPage() LogsListRequestPage {
	if o == nil || o.Page == nil {
		var ret LogsListRequestPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsListRequest) GetPageOk() (*LogsListRequestPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *LogsListRequest) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given LogsListRequestPage and assigns it to the Page field.
func (o *LogsListRequest) SetPage(v LogsListRequestPage) {
	o.Page = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *LogsListRequest) GetSort() LogsSort {
	if o == nil || o.Sort == nil {
		var ret LogsSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsListRequest) GetSortOk() (*LogsSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *LogsListRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given LogsSort and assigns it to the Sort field.
func (o *LogsListRequest) SetSort(v LogsSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsListRequest) MarshalJSON() ([]byte, error) {
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
func (o *LogsListRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter  *LogsQueryFilter     `json:"filter,omitempty"`
		Options *LogsQueryOptions    `json:"options,omitempty"`
		Page    *LogsListRequestPage `json:"page,omitempty"`
		Sort    *LogsSort            `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
