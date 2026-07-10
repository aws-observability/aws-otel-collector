// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsTimeseriesMetric A timeseries metric containing timestamps, series values, and optional unit metadata.
type CommitmentsTimeseriesMetric struct {
	// Timeseries data as a map of series names to their corresponding value arrays.
	Series map[string][]float64 `json:"series"`
	// Unix timestamps in seconds for the timeseries data points.
	Times []int64 `json:"times"`
	// Unit metadata for a numeric metric.
	Unit *CommitmentsUnit `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsTimeseriesMetric instantiates a new CommitmentsTimeseriesMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsTimeseriesMetric(series map[string][]float64, times []int64) *CommitmentsTimeseriesMetric {
	this := CommitmentsTimeseriesMetric{}
	this.Series = series
	this.Times = times
	return &this
}

// NewCommitmentsTimeseriesMetricWithDefaults instantiates a new CommitmentsTimeseriesMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsTimeseriesMetricWithDefaults() *CommitmentsTimeseriesMetric {
	this := CommitmentsTimeseriesMetric{}
	return &this
}

// GetSeries returns the Series field value.
func (o *CommitmentsTimeseriesMetric) GetSeries() map[string][]float64 {
	if o == nil {
		var ret map[string][]float64
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *CommitmentsTimeseriesMetric) GetSeriesOk() (*map[string][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Series, true
}

// SetSeries sets field value.
func (o *CommitmentsTimeseriesMetric) SetSeries(v map[string][]float64) {
	o.Series = v
}

// GetTimes returns the Times field value.
func (o *CommitmentsTimeseriesMetric) GetTimes() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value
// and a boolean to check if the value has been set.
func (o *CommitmentsTimeseriesMetric) GetTimesOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Times, true
}

// SetTimes sets field value.
func (o *CommitmentsTimeseriesMetric) SetTimes(v []int64) {
	o.Times = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CommitmentsTimeseriesMetric) GetUnit() CommitmentsUnit {
	if o == nil || o.Unit == nil {
		var ret CommitmentsUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsTimeseriesMetric) GetUnitOk() (*CommitmentsUnit, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CommitmentsTimeseriesMetric) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given CommitmentsUnit and assigns it to the Unit field.
func (o *CommitmentsTimeseriesMetric) SetUnit(v CommitmentsUnit) {
	o.Unit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsTimeseriesMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["series"] = o.Series
	toSerialize["times"] = o.Times
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsTimeseriesMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Series *map[string][]float64 `json:"series"`
		Times  *[]int64              `json:"times"`
		Unit   *CommitmentsUnit      `json:"unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Series == nil {
		return fmt.Errorf("required field series missing")
	}
	if all.Times == nil {
		return fmt.Errorf("required field times missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"series", "times", "unit"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Series = *all.Series
	o.Times = *all.Times
	if all.Unit != nil && all.Unit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
