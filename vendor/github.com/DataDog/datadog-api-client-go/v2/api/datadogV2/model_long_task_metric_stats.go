// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LongTaskMetricStats Statistical distribution (average, min, max) of a long task metric across sampled views.
type LongTaskMetricStats struct {
	// Average value across sampled views.
	Average float64 `json:"average"`
	// Maximum value across sampled views.
	Max float64 `json:"max"`
	// Minimum value across sampled views.
	Min float64 `json:"min"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLongTaskMetricStats instantiates a new LongTaskMetricStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLongTaskMetricStats(average float64, max float64, min float64) *LongTaskMetricStats {
	this := LongTaskMetricStats{}
	this.Average = average
	this.Max = max
	this.Min = min
	return &this
}

// NewLongTaskMetricStatsWithDefaults instantiates a new LongTaskMetricStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLongTaskMetricStatsWithDefaults() *LongTaskMetricStats {
	this := LongTaskMetricStats{}
	return &this
}

// GetAverage returns the Average field value.
func (o *LongTaskMetricStats) GetAverage() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Average
}

// GetAverageOk returns a tuple with the Average field value
// and a boolean to check if the value has been set.
func (o *LongTaskMetricStats) GetAverageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Average, true
}

// SetAverage sets field value.
func (o *LongTaskMetricStats) SetAverage(v float64) {
	o.Average = v
}

// GetMax returns the Max field value.
func (o *LongTaskMetricStats) GetMax() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *LongTaskMetricStats) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value.
func (o *LongTaskMetricStats) SetMax(v float64) {
	o.Max = v
}

// GetMin returns the Min field value.
func (o *LongTaskMetricStats) GetMin() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *LongTaskMetricStats) GetMinOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value.
func (o *LongTaskMetricStats) SetMin(v float64) {
	o.Min = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LongTaskMetricStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["average"] = o.Average
	toSerialize["max"] = o.Max
	toSerialize["min"] = o.Min

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LongTaskMetricStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Average *float64 `json:"average"`
		Max     *float64 `json:"max"`
		Min     *float64 `json:"min"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Average == nil {
		return fmt.Errorf("required field average missing")
	}
	if all.Max == nil {
		return fmt.Errorf("required field max missing")
	}
	if all.Min == nil {
		return fmt.Errorf("required field min missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"average", "max", "min"})
	} else {
		return err
	}
	o.Average = *all.Average
	o.Max = *all.Max
	o.Min = *all.Min

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
