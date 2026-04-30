// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsExtendedGroupBy Group by configuration for Product Analytics Extended queries.
type ProductAnalyticsExtendedGroupBy struct {
	// Facet name to group by.
	Facet string `json:"facet"`
	// Maximum number of groups to return.
	Limit *int32 `json:"limit,omitempty"`
	// Whether to exclude events missing the group-by facet.
	ShouldExcludeMissing *bool `json:"should_exclude_missing,omitempty"`
	// Options for sorting group by results.
	Sort *FormulaAndFunctionEventQueryGroupBySort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsExtendedGroupBy instantiates a new ProductAnalyticsExtendedGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsExtendedGroupBy(facet string) *ProductAnalyticsExtendedGroupBy {
	this := ProductAnalyticsExtendedGroupBy{}
	this.Facet = facet
	return &this
}

// NewProductAnalyticsExtendedGroupByWithDefaults instantiates a new ProductAnalyticsExtendedGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsExtendedGroupByWithDefaults() *ProductAnalyticsExtendedGroupBy {
	this := ProductAnalyticsExtendedGroupBy{}
	return &this
}

// GetFacet returns the Facet field value.
func (o *ProductAnalyticsExtendedGroupBy) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsExtendedGroupBy) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *ProductAnalyticsExtendedGroupBy) SetFacet(v string) {
	o.Facet = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProductAnalyticsExtendedGroupBy) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsExtendedGroupBy) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProductAnalyticsExtendedGroupBy) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ProductAnalyticsExtendedGroupBy) SetLimit(v int32) {
	o.Limit = &v
}

// GetShouldExcludeMissing returns the ShouldExcludeMissing field value if set, zero value otherwise.
func (o *ProductAnalyticsExtendedGroupBy) GetShouldExcludeMissing() bool {
	if o == nil || o.ShouldExcludeMissing == nil {
		var ret bool
		return ret
	}
	return *o.ShouldExcludeMissing
}

// GetShouldExcludeMissingOk returns a tuple with the ShouldExcludeMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsExtendedGroupBy) GetShouldExcludeMissingOk() (*bool, bool) {
	if o == nil || o.ShouldExcludeMissing == nil {
		return nil, false
	}
	return o.ShouldExcludeMissing, true
}

// HasShouldExcludeMissing returns a boolean if a field has been set.
func (o *ProductAnalyticsExtendedGroupBy) HasShouldExcludeMissing() bool {
	return o != nil && o.ShouldExcludeMissing != nil
}

// SetShouldExcludeMissing gets a reference to the given bool and assigns it to the ShouldExcludeMissing field.
func (o *ProductAnalyticsExtendedGroupBy) SetShouldExcludeMissing(v bool) {
	o.ShouldExcludeMissing = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ProductAnalyticsExtendedGroupBy) GetSort() FormulaAndFunctionEventQueryGroupBySort {
	if o == nil || o.Sort == nil {
		var ret FormulaAndFunctionEventQueryGroupBySort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsExtendedGroupBy) GetSortOk() (*FormulaAndFunctionEventQueryGroupBySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ProductAnalyticsExtendedGroupBy) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given FormulaAndFunctionEventQueryGroupBySort and assigns it to the Sort field.
func (o *ProductAnalyticsExtendedGroupBy) SetSort(v FormulaAndFunctionEventQueryGroupBySort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsExtendedGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet"] = o.Facet
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.ShouldExcludeMissing != nil {
		toSerialize["should_exclude_missing"] = o.ShouldExcludeMissing
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
func (o *ProductAnalyticsExtendedGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet                *string                                  `json:"facet"`
		Limit                *int32                                   `json:"limit,omitempty"`
		ShouldExcludeMissing *bool                                    `json:"should_exclude_missing,omitempty"`
		Sort                 *FormulaAndFunctionEventQueryGroupBySort `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet", "limit", "should_exclude_missing", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Facet = *all.Facet
	o.Limit = all.Limit
	o.ShouldExcludeMissing = all.ShouldExcludeMissing
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
