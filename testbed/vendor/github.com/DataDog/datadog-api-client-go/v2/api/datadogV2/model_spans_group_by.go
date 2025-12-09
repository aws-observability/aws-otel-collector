// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansGroupBy A group by rule.
type SpansGroupBy struct {
	// The name of the facet to use (required).
	Facet string `json:"facet"`
	// Used to perform a histogram computation (only for measure facets).
	// Note: At most 100 buckets are allowed, the number of buckets is (max - min)/interval.
	Histogram *SpansGroupByHistogram `json:"histogram,omitempty"`
	// The maximum buckets to return for this group by.
	Limit *int64 `json:"limit,omitempty"`
	// The value to use for spans that don't have the facet used to group by.
	Missing *SpansGroupByMissing `json:"missing,omitempty"`
	// A sort rule.
	Sort *SpansAggregateSort `json:"sort,omitempty"`
	// A resulting object to put the given computes in over all the matching records.
	Total *SpansGroupByTotal `json:"total,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpansGroupBy instantiates a new SpansGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansGroupBy(facet string) *SpansGroupBy {
	this := SpansGroupBy{}
	this.Facet = facet
	var limit int64 = 10
	this.Limit = &limit
	return &this
}

// NewSpansGroupByWithDefaults instantiates a new SpansGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansGroupByWithDefaults() *SpansGroupBy {
	this := SpansGroupBy{}
	var limit int64 = 10
	this.Limit = &limit
	return &this
}

// GetFacet returns the Facet field value.
func (o *SpansGroupBy) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *SpansGroupBy) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *SpansGroupBy) SetFacet(v string) {
	o.Facet = v
}

// GetHistogram returns the Histogram field value if set, zero value otherwise.
func (o *SpansGroupBy) GetHistogram() SpansGroupByHistogram {
	if o == nil || o.Histogram == nil {
		var ret SpansGroupByHistogram
		return ret
	}
	return *o.Histogram
}

// GetHistogramOk returns a tuple with the Histogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansGroupBy) GetHistogramOk() (*SpansGroupByHistogram, bool) {
	if o == nil || o.Histogram == nil {
		return nil, false
	}
	return o.Histogram, true
}

// HasHistogram returns a boolean if a field has been set.
func (o *SpansGroupBy) HasHistogram() bool {
	return o != nil && o.Histogram != nil
}

// SetHistogram gets a reference to the given SpansGroupByHistogram and assigns it to the Histogram field.
func (o *SpansGroupBy) SetHistogram(v SpansGroupByHistogram) {
	o.Histogram = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SpansGroupBy) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansGroupBy) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SpansGroupBy) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *SpansGroupBy) SetLimit(v int64) {
	o.Limit = &v
}

// GetMissing returns the Missing field value if set, zero value otherwise.
func (o *SpansGroupBy) GetMissing() SpansGroupByMissing {
	if o == nil || o.Missing == nil {
		var ret SpansGroupByMissing
		return ret
	}
	return *o.Missing
}

// GetMissingOk returns a tuple with the Missing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansGroupBy) GetMissingOk() (*SpansGroupByMissing, bool) {
	if o == nil || o.Missing == nil {
		return nil, false
	}
	return o.Missing, true
}

// HasMissing returns a boolean if a field has been set.
func (o *SpansGroupBy) HasMissing() bool {
	return o != nil && o.Missing != nil
}

// SetMissing gets a reference to the given SpansGroupByMissing and assigns it to the Missing field.
func (o *SpansGroupBy) SetMissing(v SpansGroupByMissing) {
	o.Missing = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *SpansGroupBy) GetSort() SpansAggregateSort {
	if o == nil || o.Sort == nil {
		var ret SpansAggregateSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansGroupBy) GetSortOk() (*SpansAggregateSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *SpansGroupBy) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given SpansAggregateSort and assigns it to the Sort field.
func (o *SpansGroupBy) SetSort(v SpansAggregateSort) {
	o.Sort = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SpansGroupBy) GetTotal() SpansGroupByTotal {
	if o == nil || o.Total == nil {
		var ret SpansGroupByTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansGroupBy) GetTotalOk() (*SpansGroupByTotal, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SpansGroupBy) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given SpansGroupByTotal and assigns it to the Total field.
func (o *SpansGroupBy) SetTotal(v SpansGroupByTotal) {
	o.Total = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet"] = o.Facet
	if o.Histogram != nil {
		toSerialize["histogram"] = o.Histogram
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Missing != nil {
		toSerialize["missing"] = o.Missing
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet     *string                `json:"facet"`
		Histogram *SpansGroupByHistogram `json:"histogram,omitempty"`
		Limit     *int64                 `json:"limit,omitempty"`
		Missing   *SpansGroupByMissing   `json:"missing,omitempty"`
		Sort      *SpansAggregateSort    `json:"sort,omitempty"`
		Total     *SpansGroupByTotal     `json:"total,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet", "histogram", "limit", "missing", "sort", "total"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Facet = *all.Facet
	if all.Histogram != nil && all.Histogram.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Histogram = all.Histogram
	o.Limit = all.Limit
	o.Missing = all.Missing
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	o.Total = all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
