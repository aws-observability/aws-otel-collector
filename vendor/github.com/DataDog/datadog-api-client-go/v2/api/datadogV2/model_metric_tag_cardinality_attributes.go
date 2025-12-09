// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricTagCardinalityAttributes An object containing properties related to the tag key
type MetricTagCardinalityAttributes struct {
	// This describes the recent change in the tag keys cardinality
	CardinalityDelta *int64 `json:"cardinality_delta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricTagCardinalityAttributes instantiates a new MetricTagCardinalityAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricTagCardinalityAttributes() *MetricTagCardinalityAttributes {
	this := MetricTagCardinalityAttributes{}
	return &this
}

// NewMetricTagCardinalityAttributesWithDefaults instantiates a new MetricTagCardinalityAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricTagCardinalityAttributesWithDefaults() *MetricTagCardinalityAttributes {
	this := MetricTagCardinalityAttributes{}
	return &this
}

// GetCardinalityDelta returns the CardinalityDelta field value if set, zero value otherwise.
func (o *MetricTagCardinalityAttributes) GetCardinalityDelta() int64 {
	if o == nil || o.CardinalityDelta == nil {
		var ret int64
		return ret
	}
	return *o.CardinalityDelta
}

// GetCardinalityDeltaOk returns a tuple with the CardinalityDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagCardinalityAttributes) GetCardinalityDeltaOk() (*int64, bool) {
	if o == nil || o.CardinalityDelta == nil {
		return nil, false
	}
	return o.CardinalityDelta, true
}

// HasCardinalityDelta returns a boolean if a field has been set.
func (o *MetricTagCardinalityAttributes) HasCardinalityDelta() bool {
	return o != nil && o.CardinalityDelta != nil
}

// SetCardinalityDelta gets a reference to the given int64 and assigns it to the CardinalityDelta field.
func (o *MetricTagCardinalityAttributes) SetCardinalityDelta(v int64) {
	o.CardinalityDelta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricTagCardinalityAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CardinalityDelta != nil {
		toSerialize["cardinality_delta"] = o.CardinalityDelta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricTagCardinalityAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CardinalityDelta *int64 `json:"cardinality_delta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cardinality_delta"})
	} else {
		return err
	}
	o.CardinalityDelta = all.CardinalityDelta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
