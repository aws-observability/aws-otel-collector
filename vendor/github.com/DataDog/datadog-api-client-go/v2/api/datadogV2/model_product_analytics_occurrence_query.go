// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsOccurrenceQuery A Product Analytics occurrence-filtered query.
type ProductAnalyticsOccurrenceQuery struct {
	// The data source identifier for occurrence queries.
	DataSource ProductAnalyticsOccurrenceQueryDataSource `json:"data_source"`
	// Search parameters for an occurrence query.
	Search ProductAnalyticsOccurrenceSearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsOccurrenceQuery instantiates a new ProductAnalyticsOccurrenceQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsOccurrenceQuery(dataSource ProductAnalyticsOccurrenceQueryDataSource, search ProductAnalyticsOccurrenceSearch) *ProductAnalyticsOccurrenceQuery {
	this := ProductAnalyticsOccurrenceQuery{}
	this.DataSource = dataSource
	this.Search = search
	return &this
}

// NewProductAnalyticsOccurrenceQueryWithDefaults instantiates a new ProductAnalyticsOccurrenceQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsOccurrenceQueryWithDefaults() *ProductAnalyticsOccurrenceQuery {
	this := ProductAnalyticsOccurrenceQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *ProductAnalyticsOccurrenceQuery) GetDataSource() ProductAnalyticsOccurrenceQueryDataSource {
	if o == nil {
		var ret ProductAnalyticsOccurrenceQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsOccurrenceQuery) GetDataSourceOk() (*ProductAnalyticsOccurrenceQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ProductAnalyticsOccurrenceQuery) SetDataSource(v ProductAnalyticsOccurrenceQueryDataSource) {
	o.DataSource = v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsOccurrenceQuery) GetSearch() ProductAnalyticsOccurrenceSearch {
	if o == nil {
		var ret ProductAnalyticsOccurrenceSearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsOccurrenceQuery) GetSearchOk() (*ProductAnalyticsOccurrenceSearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsOccurrenceQuery) SetSearch(v ProductAnalyticsOccurrenceSearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsOccurrenceQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["search"] = o.Search

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsOccurrenceQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *ProductAnalyticsOccurrenceQueryDataSource `json:"data_source"`
		Search     *ProductAnalyticsOccurrenceSearch          `json:"search"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "search"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
