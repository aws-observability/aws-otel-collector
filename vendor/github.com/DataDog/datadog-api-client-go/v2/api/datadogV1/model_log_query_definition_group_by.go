// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogQueryDefinitionGroupBy Defined items in the group.
type LogQueryDefinitionGroupBy struct {
	// Facet name.
	Facet string `json:"facet"`
	// Maximum number of items in the group.
	Limit *int64 `json:"limit,omitempty"`
	// Define a sorting method.
	Sort *LogQueryDefinitionGroupBySort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogQueryDefinitionGroupBy instantiates a new LogQueryDefinitionGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogQueryDefinitionGroupBy(facet string) *LogQueryDefinitionGroupBy {
	this := LogQueryDefinitionGroupBy{}
	this.Facet = facet
	return &this
}

// NewLogQueryDefinitionGroupByWithDefaults instantiates a new LogQueryDefinitionGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogQueryDefinitionGroupByWithDefaults() *LogQueryDefinitionGroupBy {
	this := LogQueryDefinitionGroupBy{}
	return &this
}

// GetFacet returns the Facet field value.
func (o *LogQueryDefinitionGroupBy) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *LogQueryDefinitionGroupBy) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *LogQueryDefinitionGroupBy) SetFacet(v string) {
	o.Facet = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *LogQueryDefinitionGroupBy) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogQueryDefinitionGroupBy) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *LogQueryDefinitionGroupBy) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *LogQueryDefinitionGroupBy) SetLimit(v int64) {
	o.Limit = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *LogQueryDefinitionGroupBy) GetSort() LogQueryDefinitionGroupBySort {
	if o == nil || o.Sort == nil {
		var ret LogQueryDefinitionGroupBySort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogQueryDefinitionGroupBy) GetSortOk() (*LogQueryDefinitionGroupBySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *LogQueryDefinitionGroupBy) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given LogQueryDefinitionGroupBySort and assigns it to the Sort field.
func (o *LogQueryDefinitionGroupBy) SetSort(v LogQueryDefinitionGroupBySort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogQueryDefinitionGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet"] = o.Facet
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
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
func (o *LogQueryDefinitionGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet *string                        `json:"facet"`
		Limit *int64                         `json:"limit,omitempty"`
		Sort  *LogQueryDefinitionGroupBySort `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet", "limit", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Facet = *all.Facet
	o.Limit = all.Limit
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
