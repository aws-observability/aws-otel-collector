// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatMapWidgetXAxis X Axis controls for the heat map widget.
type HeatMapWidgetXAxis struct {
	// Number of time buckets to target, also known as the resolution
	// of the time bins. This is only applicable for distribution of
	// points (group distributions use the roll-up modifier).
	NumBuckets *int64 `json:"num_buckets,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHeatMapWidgetXAxis instantiates a new HeatMapWidgetXAxis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatMapWidgetXAxis() *HeatMapWidgetXAxis {
	this := HeatMapWidgetXAxis{}
	return &this
}

// NewHeatMapWidgetXAxisWithDefaults instantiates a new HeatMapWidgetXAxis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatMapWidgetXAxisWithDefaults() *HeatMapWidgetXAxis {
	this := HeatMapWidgetXAxis{}
	return &this
}

// GetNumBuckets returns the NumBuckets field value if set, zero value otherwise.
func (o *HeatMapWidgetXAxis) GetNumBuckets() int64 {
	if o == nil || o.NumBuckets == nil {
		var ret int64
		return ret
	}
	return *o.NumBuckets
}

// GetNumBucketsOk returns a tuple with the NumBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatMapWidgetXAxis) GetNumBucketsOk() (*int64, bool) {
	if o == nil || o.NumBuckets == nil {
		return nil, false
	}
	return o.NumBuckets, true
}

// HasNumBuckets returns a boolean if a field has been set.
func (o *HeatMapWidgetXAxis) HasNumBuckets() bool {
	return o != nil && o.NumBuckets != nil
}

// SetNumBuckets gets a reference to the given int64 and assigns it to the NumBuckets field.
func (o *HeatMapWidgetXAxis) SetNumBuckets(v int64) {
	o.NumBuckets = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatMapWidgetXAxis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NumBuckets != nil {
		toSerialize["num_buckets"] = o.NumBuckets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatMapWidgetXAxis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NumBuckets *int64 `json:"num_buckets,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"num_buckets"})
	} else {
		return err
	}
	o.NumBuckets = all.NumBuckets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
