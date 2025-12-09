// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateBucketAttributes A bucket values.
type SpansAggregateBucketAttributes struct {
	// The key, value pairs for each group by.
	By map[string]interface{} `json:"by,omitempty"`
	// The compute data.
	Compute interface{} `json:"compute,omitempty"`
	// A map of the metric name -> value for regular compute or list of values for a timeseries.
	Computes map[string]SpansAggregateBucketValue `json:"computes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpansAggregateBucketAttributes instantiates a new SpansAggregateBucketAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansAggregateBucketAttributes() *SpansAggregateBucketAttributes {
	this := SpansAggregateBucketAttributes{}
	return &this
}

// NewSpansAggregateBucketAttributesWithDefaults instantiates a new SpansAggregateBucketAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansAggregateBucketAttributesWithDefaults() *SpansAggregateBucketAttributes {
	this := SpansAggregateBucketAttributes{}
	return &this
}

// GetBy returns the By field value if set, zero value otherwise.
func (o *SpansAggregateBucketAttributes) GetBy() map[string]interface{} {
	if o == nil || o.By == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.By
}

// GetByOk returns a tuple with the By field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateBucketAttributes) GetByOk() (*map[string]interface{}, bool) {
	if o == nil || o.By == nil {
		return nil, false
	}
	return &o.By, true
}

// HasBy returns a boolean if a field has been set.
func (o *SpansAggregateBucketAttributes) HasBy() bool {
	return o != nil && o.By != nil
}

// SetBy gets a reference to the given map[string]interface{} and assigns it to the By field.
func (o *SpansAggregateBucketAttributes) SetBy(v map[string]interface{}) {
	o.By = v
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *SpansAggregateBucketAttributes) GetCompute() interface{} {
	if o == nil || o.Compute == nil {
		var ret interface{}
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateBucketAttributes) GetComputeOk() (*interface{}, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return &o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *SpansAggregateBucketAttributes) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given interface{} and assigns it to the Compute field.
func (o *SpansAggregateBucketAttributes) SetCompute(v interface{}) {
	o.Compute = v
}

// GetComputes returns the Computes field value if set, zero value otherwise.
func (o *SpansAggregateBucketAttributes) GetComputes() map[string]SpansAggregateBucketValue {
	if o == nil || o.Computes == nil {
		var ret map[string]SpansAggregateBucketValue
		return ret
	}
	return o.Computes
}

// GetComputesOk returns a tuple with the Computes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateBucketAttributes) GetComputesOk() (*map[string]SpansAggregateBucketValue, bool) {
	if o == nil || o.Computes == nil {
		return nil, false
	}
	return &o.Computes, true
}

// HasComputes returns a boolean if a field has been set.
func (o *SpansAggregateBucketAttributes) HasComputes() bool {
	return o != nil && o.Computes != nil
}

// SetComputes gets a reference to the given map[string]SpansAggregateBucketValue and assigns it to the Computes field.
func (o *SpansAggregateBucketAttributes) SetComputes(v map[string]SpansAggregateBucketValue) {
	o.Computes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansAggregateBucketAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.By != nil {
		toSerialize["by"] = o.By
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	if o.Computes != nil {
		toSerialize["computes"] = o.Computes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansAggregateBucketAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		By       map[string]interface{}               `json:"by,omitempty"`
		Compute  interface{}                          `json:"compute,omitempty"`
		Computes map[string]SpansAggregateBucketValue `json:"computes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"by", "compute", "computes"})
	} else {
		return err
	}
	o.By = all.By
	o.Compute = all.Compute
	o.Computes = all.Computes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
