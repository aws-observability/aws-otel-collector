// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabMetricStatRange The range of values for a specific metric statistic.
type ModelLabMetricStatRange struct {
	// The maximum value of the statistic.
	Max float64 `json:"max"`
	// The minimum value of the statistic.
	Min float64 `json:"min"`
	// The metric statistic name.
	Stat string `json:"stat"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabMetricStatRange instantiates a new ModelLabMetricStatRange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabMetricStatRange(max float64, min float64, stat string) *ModelLabMetricStatRange {
	this := ModelLabMetricStatRange{}
	this.Max = max
	this.Min = min
	this.Stat = stat
	return &this
}

// NewModelLabMetricStatRangeWithDefaults instantiates a new ModelLabMetricStatRange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabMetricStatRangeWithDefaults() *ModelLabMetricStatRange {
	this := ModelLabMetricStatRange{}
	return &this
}

// GetMax returns the Max field value.
func (o *ModelLabMetricStatRange) GetMax() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *ModelLabMetricStatRange) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value.
func (o *ModelLabMetricStatRange) SetMax(v float64) {
	o.Max = v
}

// GetMin returns the Min field value.
func (o *ModelLabMetricStatRange) GetMin() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *ModelLabMetricStatRange) GetMinOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value.
func (o *ModelLabMetricStatRange) SetMin(v float64) {
	o.Min = v
}

// GetStat returns the Stat field value.
func (o *ModelLabMetricStatRange) GetStat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Stat
}

// GetStatOk returns a tuple with the Stat field value
// and a boolean to check if the value has been set.
func (o *ModelLabMetricStatRange) GetStatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stat, true
}

// SetStat sets field value.
func (o *ModelLabMetricStatRange) SetStat(v string) {
	o.Stat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabMetricStatRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["max"] = o.Max
	toSerialize["min"] = o.Min
	toSerialize["stat"] = o.Stat

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabMetricStatRange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Max  *float64 `json:"max"`
		Min  *float64 `json:"min"`
		Stat *string  `json:"stat"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Max == nil {
		return fmt.Errorf("required field max missing")
	}
	if all.Min == nil {
		return fmt.Errorf("required field min missing")
	}
	if all.Stat == nil {
		return fmt.Errorf("required field stat missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max", "min", "stat"})
	} else {
		return err
	}
	o.Max = *all.Max
	o.Min = *all.Min
	o.Stat = *all.Stat

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
