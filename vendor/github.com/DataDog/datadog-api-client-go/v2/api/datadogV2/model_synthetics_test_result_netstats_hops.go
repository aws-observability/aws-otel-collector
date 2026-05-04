// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultNetstatsHops Statistics about the number of hops for a network test.
type SyntheticsTestResultNetstatsHops struct {
	// Average number of hops.
	Avg *float64 `json:"avg,omitempty"`
	// Maximum number of hops.
	Max *int64 `json:"max,omitempty"`
	// Minimum number of hops.
	Min *int64 `json:"min,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultNetstatsHops instantiates a new SyntheticsTestResultNetstatsHops object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultNetstatsHops() *SyntheticsTestResultNetstatsHops {
	this := SyntheticsTestResultNetstatsHops{}
	return &this
}

// NewSyntheticsTestResultNetstatsHopsWithDefaults instantiates a new SyntheticsTestResultNetstatsHops object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultNetstatsHopsWithDefaults() *SyntheticsTestResultNetstatsHops {
	this := SyntheticsTestResultNetstatsHops{}
	return &this
}

// GetAvg returns the Avg field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetstatsHops) GetAvg() float64 {
	if o == nil || o.Avg == nil {
		var ret float64
		return ret
	}
	return *o.Avg
}

// GetAvgOk returns a tuple with the Avg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetstatsHops) GetAvgOk() (*float64, bool) {
	if o == nil || o.Avg == nil {
		return nil, false
	}
	return o.Avg, true
}

// HasAvg returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetstatsHops) HasAvg() bool {
	return o != nil && o.Avg != nil
}

// SetAvg gets a reference to the given float64 and assigns it to the Avg field.
func (o *SyntheticsTestResultNetstatsHops) SetAvg(v float64) {
	o.Avg = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetstatsHops) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetstatsHops) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetstatsHops) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *SyntheticsTestResultNetstatsHops) SetMax(v int64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetstatsHops) GetMin() int64 {
	if o == nil || o.Min == nil {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetstatsHops) GetMinOk() (*int64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetstatsHops) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *SyntheticsTestResultNetstatsHops) SetMin(v int64) {
	o.Min = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultNetstatsHops) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Avg != nil {
		toSerialize["avg"] = o.Avg
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultNetstatsHops) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Avg *float64 `json:"avg,omitempty"`
		Max *int64   `json:"max,omitempty"`
		Min *int64   `json:"min,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avg", "max", "min"})
	} else {
		return err
	}
	o.Avg = all.Avg
	o.Max = all.Max
	o.Min = all.Min

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
