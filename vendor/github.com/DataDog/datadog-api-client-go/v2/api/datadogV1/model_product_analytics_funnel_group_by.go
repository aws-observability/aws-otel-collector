// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelGroupBy Group by configuration for user journey funnel.
type ProductAnalyticsFunnelGroupBy struct {
	// Facet to group by.
	Facet string `json:"facet"`
	// Maximum number of groups.
	Limit *int64 `json:"limit,omitempty"`
	// Whether to exclude missing values.
	ShouldExcludeMissing *bool `json:"should_exclude_missing,omitempty"`
	// Sort configuration for user journey funnel group by.
	Sort *ProductAnalyticsFunnelGroupBySort `json:"sort,omitempty"`
	// Target for user journey search.
	Target *UserJourneySearchTarget `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelGroupBy instantiates a new ProductAnalyticsFunnelGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelGroupBy(facet string) *ProductAnalyticsFunnelGroupBy {
	this := ProductAnalyticsFunnelGroupBy{}
	this.Facet = facet
	return &this
}

// NewProductAnalyticsFunnelGroupByWithDefaults instantiates a new ProductAnalyticsFunnelGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelGroupByWithDefaults() *ProductAnalyticsFunnelGroupBy {
	this := ProductAnalyticsFunnelGroupBy{}
	return &this
}

// GetFacet returns the Facet field value.
func (o *ProductAnalyticsFunnelGroupBy) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelGroupBy) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *ProductAnalyticsFunnelGroupBy) SetFacet(v string) {
	o.Facet = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelGroupBy) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelGroupBy) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelGroupBy) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ProductAnalyticsFunnelGroupBy) SetLimit(v int64) {
	o.Limit = &v
}

// GetShouldExcludeMissing returns the ShouldExcludeMissing field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelGroupBy) GetShouldExcludeMissing() bool {
	if o == nil || o.ShouldExcludeMissing == nil {
		var ret bool
		return ret
	}
	return *o.ShouldExcludeMissing
}

// GetShouldExcludeMissingOk returns a tuple with the ShouldExcludeMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelGroupBy) GetShouldExcludeMissingOk() (*bool, bool) {
	if o == nil || o.ShouldExcludeMissing == nil {
		return nil, false
	}
	return o.ShouldExcludeMissing, true
}

// HasShouldExcludeMissing returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelGroupBy) HasShouldExcludeMissing() bool {
	return o != nil && o.ShouldExcludeMissing != nil
}

// SetShouldExcludeMissing gets a reference to the given bool and assigns it to the ShouldExcludeMissing field.
func (o *ProductAnalyticsFunnelGroupBy) SetShouldExcludeMissing(v bool) {
	o.ShouldExcludeMissing = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelGroupBy) GetSort() ProductAnalyticsFunnelGroupBySort {
	if o == nil || o.Sort == nil {
		var ret ProductAnalyticsFunnelGroupBySort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelGroupBy) GetSortOk() (*ProductAnalyticsFunnelGroupBySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelGroupBy) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ProductAnalyticsFunnelGroupBySort and assigns it to the Sort field.
func (o *ProductAnalyticsFunnelGroupBy) SetSort(v ProductAnalyticsFunnelGroupBySort) {
	o.Sort = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelGroupBy) GetTarget() UserJourneySearchTarget {
	if o == nil || o.Target == nil {
		var ret UserJourneySearchTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelGroupBy) GetTargetOk() (*UserJourneySearchTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelGroupBy) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given UserJourneySearchTarget and assigns it to the Target field.
func (o *ProductAnalyticsFunnelGroupBy) SetTarget(v UserJourneySearchTarget) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelGroupBy) MarshalJSON() ([]byte, error) {
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
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet                *string                            `json:"facet"`
		Limit                *int64                             `json:"limit,omitempty"`
		ShouldExcludeMissing *bool                              `json:"should_exclude_missing,omitempty"`
		Sort                 *ProductAnalyticsFunnelGroupBySort `json:"sort,omitempty"`
		Target               *UserJourneySearchTarget           `json:"target,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet", "limit", "should_exclude_missing", "sort", "target"})
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
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
