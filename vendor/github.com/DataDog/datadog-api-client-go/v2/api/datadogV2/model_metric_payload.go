// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricPayload The metrics' payload.
type MetricPayload struct {
	// A list of timeseries to submit to Datadog.
	Series []MetricSeries `json:"series"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricPayload instantiates a new MetricPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricPayload(series []MetricSeries) *MetricPayload {
	this := MetricPayload{}
	this.Series = series
	return &this
}

// NewMetricPayloadWithDefaults instantiates a new MetricPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricPayloadWithDefaults() *MetricPayload {
	this := MetricPayload{}
	return &this
}

// GetSeries returns the Series field value.
func (o *MetricPayload) GetSeries() []MetricSeries {
	if o == nil {
		var ret []MetricSeries
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *MetricPayload) GetSeriesOk() (*[]MetricSeries, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Series, true
}

// SetSeries sets field value.
func (o *MetricPayload) SetSeries(v []MetricSeries) {
	o.Series = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["series"] = o.Series

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Series *[]MetricSeries `json:"series"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Series == nil {
		return fmt.Errorf("required field series missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"series"})
	} else {
		return err
	}
	o.Series = *all.Series

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
