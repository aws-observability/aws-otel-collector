// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricCustomAggregation A time and space aggregation combination for use in query.
type MetricCustomAggregation struct {
	// A space aggregation for use in query.
	Space MetricCustomSpaceAggregation `json:"space"`
	// A time aggregation for use in query.
	Time MetricCustomTimeAggregation `json:"time"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricCustomAggregation instantiates a new MetricCustomAggregation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricCustomAggregation(space MetricCustomSpaceAggregation, time MetricCustomTimeAggregation) *MetricCustomAggregation {
	this := MetricCustomAggregation{}
	this.Space = space
	this.Time = time
	return &this
}

// NewMetricCustomAggregationWithDefaults instantiates a new MetricCustomAggregation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricCustomAggregationWithDefaults() *MetricCustomAggregation {
	this := MetricCustomAggregation{}
	return &this
}

// GetSpace returns the Space field value.
func (o *MetricCustomAggregation) GetSpace() MetricCustomSpaceAggregation {
	if o == nil {
		var ret MetricCustomSpaceAggregation
		return ret
	}
	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *MetricCustomAggregation) GetSpaceOk() (*MetricCustomSpaceAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value.
func (o *MetricCustomAggregation) SetSpace(v MetricCustomSpaceAggregation) {
	o.Space = v
}

// GetTime returns the Time field value.
func (o *MetricCustomAggregation) GetTime() MetricCustomTimeAggregation {
	if o == nil {
		var ret MetricCustomTimeAggregation
		return ret
	}
	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *MetricCustomAggregation) GetTimeOk() (*MetricCustomTimeAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value.
func (o *MetricCustomAggregation) SetTime(v MetricCustomTimeAggregation) {
	o.Time = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricCustomAggregation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["space"] = o.Space
	toSerialize["time"] = o.Time

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricCustomAggregation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Space *MetricCustomSpaceAggregation `json:"space"`
		Time  *MetricCustomTimeAggregation  `json:"time"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Space == nil {
		return fmt.Errorf("required field space missing")
	}
	if all.Time == nil {
		return fmt.Errorf("required field time missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"space", "time"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Space.IsValid() {
		hasInvalidField = true
	} else {
		o.Space = *all.Space
	}
	if !all.Time.IsValid() {
		hasInvalidField = true
	} else {
		o.Time = *all.Time
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
