// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionFilters Filters for retention queries.
type RetentionFilters struct {
	// Product Analytics/RUM audience filters.
	AudienceFilters *ProductAnalyticsAudienceFilters `json:"audience_filters,omitempty"`
	// String filter.
	StringFilter *string `json:"string_filter,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionFilters instantiates a new RetentionFilters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionFilters() *RetentionFilters {
	this := RetentionFilters{}
	return &this
}

// NewRetentionFiltersWithDefaults instantiates a new RetentionFilters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionFiltersWithDefaults() *RetentionFilters {
	this := RetentionFilters{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *RetentionFilters) GetAudienceFilters() ProductAnalyticsAudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret ProductAnalyticsAudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilters) GetAudienceFiltersOk() (*ProductAnalyticsAudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *RetentionFilters) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given ProductAnalyticsAudienceFilters and assigns it to the AudienceFilters field.
func (o *RetentionFilters) SetAudienceFilters(v ProductAnalyticsAudienceFilters) {
	o.AudienceFilters = &v
}

// GetStringFilter returns the StringFilter field value if set, zero value otherwise.
func (o *RetentionFilters) GetStringFilter() string {
	if o == nil || o.StringFilter == nil {
		var ret string
		return ret
	}
	return *o.StringFilter
}

// GetStringFilterOk returns a tuple with the StringFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilters) GetStringFilterOk() (*string, bool) {
	if o == nil || o.StringFilter == nil {
		return nil, false
	}
	return o.StringFilter, true
}

// HasStringFilter returns a boolean if a field has been set.
func (o *RetentionFilters) HasStringFilter() bool {
	return o != nil && o.StringFilter != nil
}

// SetStringFilter gets a reference to the given string and assigns it to the StringFilter field.
func (o *RetentionFilters) SetStringFilter(v string) {
	o.StringFilter = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	if o.StringFilter != nil {
		toSerialize["string_filter"] = o.StringFilter
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionFilters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *ProductAnalyticsAudienceFilters `json:"audience_filters,omitempty"`
		StringFilter    *string                          `json:"string_filter,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	o.StringFilter = all.StringFilter

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
