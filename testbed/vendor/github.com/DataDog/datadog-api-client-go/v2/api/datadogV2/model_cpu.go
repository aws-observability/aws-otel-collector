// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Cpu CPU usage statistics derived from historical Spark job metrics. Provides multiple estimates so users can choose between conservative and cost-saving risk profiles.
type Cpu struct {
	// Maximum CPU usage observed for the job, expressed in millicores. This represents the upper bound of usage.
	Max *int64 `json:"max,omitempty"`
	// 75th percentile of CPU usage (millicores). Represents a cost-saving configuration while covering most workloads.
	P75 *int64 `json:"p75,omitempty"`
	// 95th percentile of CPU usage (millicores). Balances performance and cost, providing a safer margin than p75.
	P95 *int64 `json:"p95,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCpu instantiates a new Cpu object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCpu() *Cpu {
	this := Cpu{}
	return &this
}

// NewCpuWithDefaults instantiates a new Cpu object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCpuWithDefaults() *Cpu {
	this := Cpu{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Cpu) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Cpu) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *Cpu) SetMax(v int64) {
	o.Max = &v
}

// GetP75 returns the P75 field value if set, zero value otherwise.
func (o *Cpu) GetP75() int64 {
	if o == nil || o.P75 == nil {
		var ret int64
		return ret
	}
	return *o.P75
}

// GetP75Ok returns a tuple with the P75 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu) GetP75Ok() (*int64, bool) {
	if o == nil || o.P75 == nil {
		return nil, false
	}
	return o.P75, true
}

// HasP75 returns a boolean if a field has been set.
func (o *Cpu) HasP75() bool {
	return o != nil && o.P75 != nil
}

// SetP75 gets a reference to the given int64 and assigns it to the P75 field.
func (o *Cpu) SetP75(v int64) {
	o.P75 = &v
}

// GetP95 returns the P95 field value if set, zero value otherwise.
func (o *Cpu) GetP95() int64 {
	if o == nil || o.P95 == nil {
		var ret int64
		return ret
	}
	return *o.P95
}

// GetP95Ok returns a tuple with the P95 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cpu) GetP95Ok() (*int64, bool) {
	if o == nil || o.P95 == nil {
		return nil, false
	}
	return o.P95, true
}

// HasP95 returns a boolean if a field has been set.
func (o *Cpu) HasP95() bool {
	return o != nil && o.P95 != nil
}

// SetP95 gets a reference to the given int64 and assigns it to the P95 field.
func (o *Cpu) SetP95(v int64) {
	o.P95 = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Cpu) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.P75 != nil {
		toSerialize["p75"] = o.P75
	}
	if o.P95 != nil {
		toSerialize["p95"] = o.P95
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Cpu) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Max *int64 `json:"max,omitempty"`
		P75 *int64 `json:"p75,omitempty"`
		P95 *int64 `json:"p95,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max", "p75", "p95"})
	} else {
		return err
	}
	o.Max = all.Max
	o.P75 = all.P75
	o.P95 = all.P95

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
