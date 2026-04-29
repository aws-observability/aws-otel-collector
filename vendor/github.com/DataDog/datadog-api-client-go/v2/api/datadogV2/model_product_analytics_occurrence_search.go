// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsOccurrenceSearch Search parameters for an occurrence query.
type ProductAnalyticsOccurrenceSearch struct {
	// Filter for occurrence-based queries.
	Occurrences *ProductAnalyticsOccurrenceFilter `json:"occurrences,omitempty"`
	// The search query using Datadog search syntax.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsOccurrenceSearch instantiates a new ProductAnalyticsOccurrenceSearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsOccurrenceSearch() *ProductAnalyticsOccurrenceSearch {
	this := ProductAnalyticsOccurrenceSearch{}
	return &this
}

// NewProductAnalyticsOccurrenceSearchWithDefaults instantiates a new ProductAnalyticsOccurrenceSearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsOccurrenceSearchWithDefaults() *ProductAnalyticsOccurrenceSearch {
	this := ProductAnalyticsOccurrenceSearch{}
	return &this
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *ProductAnalyticsOccurrenceSearch) GetOccurrences() ProductAnalyticsOccurrenceFilter {
	if o == nil || o.Occurrences == nil {
		var ret ProductAnalyticsOccurrenceFilter
		return ret
	}
	return *o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsOccurrenceSearch) GetOccurrencesOk() (*ProductAnalyticsOccurrenceFilter, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *ProductAnalyticsOccurrenceSearch) HasOccurrences() bool {
	return o != nil && o.Occurrences != nil
}

// SetOccurrences gets a reference to the given ProductAnalyticsOccurrenceFilter and assigns it to the Occurrences field.
func (o *ProductAnalyticsOccurrenceSearch) SetOccurrences(v ProductAnalyticsOccurrenceFilter) {
	o.Occurrences = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ProductAnalyticsOccurrenceSearch) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsOccurrenceSearch) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ProductAnalyticsOccurrenceSearch) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ProductAnalyticsOccurrenceSearch) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsOccurrenceSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsOccurrenceSearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Occurrences *ProductAnalyticsOccurrenceFilter `json:"occurrences,omitempty"`
		Query       *string                           `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"occurrences", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Occurrences != nil && all.Occurrences.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Occurrences = all.Occurrences
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
