// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsTimeseriesResponseAttributes Attributes of a timeseries analytics response, containing series data, timestamps, and interval definitions.
type ProductAnalyticsTimeseriesResponseAttributes struct {
	// Interval definitions describing the time buckets used in the response.
	Intervals []ProductAnalyticsInterval `json:"intervals,omitempty"`
	// The list of series, each corresponding to a query or group-by combination.
	Series []ProductAnalyticsSerie `json:"series,omitempty"`
	// Timestamps for each data point (epoch milliseconds).
	Times []int64 `json:"times,omitempty"`
	// Values for each series at each time point.
	Values [][]*float64 `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsTimeseriesResponseAttributes instantiates a new ProductAnalyticsTimeseriesResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsTimeseriesResponseAttributes() *ProductAnalyticsTimeseriesResponseAttributes {
	this := ProductAnalyticsTimeseriesResponseAttributes{}
	return &this
}

// NewProductAnalyticsTimeseriesResponseAttributesWithDefaults instantiates a new ProductAnalyticsTimeseriesResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsTimeseriesResponseAttributesWithDefaults() *ProductAnalyticsTimeseriesResponseAttributes {
	this := ProductAnalyticsTimeseriesResponseAttributes{}
	return &this
}

// GetIntervals returns the Intervals field value if set, zero value otherwise.
func (o *ProductAnalyticsTimeseriesResponseAttributes) GetIntervals() []ProductAnalyticsInterval {
	if o == nil || o.Intervals == nil {
		var ret []ProductAnalyticsInterval
		return ret
	}
	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsTimeseriesResponseAttributes) GetIntervalsOk() (*[]ProductAnalyticsInterval, bool) {
	if o == nil || o.Intervals == nil {
		return nil, false
	}
	return &o.Intervals, true
}

// HasIntervals returns a boolean if a field has been set.
func (o *ProductAnalyticsTimeseriesResponseAttributes) HasIntervals() bool {
	return o != nil && o.Intervals != nil
}

// SetIntervals gets a reference to the given []ProductAnalyticsInterval and assigns it to the Intervals field.
func (o *ProductAnalyticsTimeseriesResponseAttributes) SetIntervals(v []ProductAnalyticsInterval) {
	o.Intervals = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *ProductAnalyticsTimeseriesResponseAttributes) GetSeries() []ProductAnalyticsSerie {
	if o == nil || o.Series == nil {
		var ret []ProductAnalyticsSerie
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsTimeseriesResponseAttributes) GetSeriesOk() (*[]ProductAnalyticsSerie, bool) {
	if o == nil || o.Series == nil {
		return nil, false
	}
	return &o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *ProductAnalyticsTimeseriesResponseAttributes) HasSeries() bool {
	return o != nil && o.Series != nil
}

// SetSeries gets a reference to the given []ProductAnalyticsSerie and assigns it to the Series field.
func (o *ProductAnalyticsTimeseriesResponseAttributes) SetSeries(v []ProductAnalyticsSerie) {
	o.Series = v
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *ProductAnalyticsTimeseriesResponseAttributes) GetTimes() []int64 {
	if o == nil || o.Times == nil {
		var ret []int64
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsTimeseriesResponseAttributes) GetTimesOk() (*[]int64, bool) {
	if o == nil || o.Times == nil {
		return nil, false
	}
	return &o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *ProductAnalyticsTimeseriesResponseAttributes) HasTimes() bool {
	return o != nil && o.Times != nil
}

// SetTimes gets a reference to the given []int64 and assigns it to the Times field.
func (o *ProductAnalyticsTimeseriesResponseAttributes) SetTimes(v []int64) {
	o.Times = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ProductAnalyticsTimeseriesResponseAttributes) GetValues() [][]*float64 {
	if o == nil || o.Values == nil {
		var ret [][]*float64
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsTimeseriesResponseAttributes) GetValuesOk() (*[][]*float64, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ProductAnalyticsTimeseriesResponseAttributes) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given [][]*float64 and assigns it to the Values field.
func (o *ProductAnalyticsTimeseriesResponseAttributes) SetValues(v [][]*float64) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsTimeseriesResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Intervals != nil {
		toSerialize["intervals"] = o.Intervals
	}
	if o.Series != nil {
		toSerialize["series"] = o.Series
	}
	if o.Times != nil {
		toSerialize["times"] = o.Times
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsTimeseriesResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Intervals []ProductAnalyticsInterval `json:"intervals,omitempty"`
		Series    []ProductAnalyticsSerie    `json:"series,omitempty"`
		Times     []int64                    `json:"times,omitempty"`
		Values    [][]*float64               `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"intervals", "series", "times", "values"})
	} else {
		return err
	}
	o.Intervals = all.Intervals
	o.Series = all.Series
	o.Times = all.Times
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
