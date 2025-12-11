// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Estimation Recommended resource values for a Spark driver or executor, derived from recent real usage metrics. Used by SPA to propose more efficient pod sizing.
type Estimation struct {
	// CPU usage statistics derived from historical Spark job metrics. Provides multiple estimates so users can choose between conservative and cost-saving risk profiles.
	Cpu *Cpu `json:"cpu,omitempty"`
	// Recommended ephemeral storage allocation (in MiB). Derived from job temporary storage patterns.
	EphemeralStorage *int64 `json:"ephemeral_storage,omitempty"`
	// Recommended JVM heap size (in MiB).
	Heap *int64 `json:"heap,omitempty"`
	// Recommended total memory allocation (in MiB). Includes both heap and overhead.
	Memory *int64 `json:"memory,omitempty"`
	// Recommended JVM overhead (in MiB). Computed as total memory - heap.
	Overhead *int64 `json:"overhead,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEstimation instantiates a new Estimation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEstimation() *Estimation {
	this := Estimation{}
	return &this
}

// NewEstimationWithDefaults instantiates a new Estimation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEstimationWithDefaults() *Estimation {
	this := Estimation{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Estimation) GetCpu() Cpu {
	if o == nil || o.Cpu == nil {
		var ret Cpu
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimation) GetCpuOk() (*Cpu, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Estimation) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given Cpu and assigns it to the Cpu field.
func (o *Estimation) SetCpu(v Cpu) {
	o.Cpu = &v
}

// GetEphemeralStorage returns the EphemeralStorage field value if set, zero value otherwise.
func (o *Estimation) GetEphemeralStorage() int64 {
	if o == nil || o.EphemeralStorage == nil {
		var ret int64
		return ret
	}
	return *o.EphemeralStorage
}

// GetEphemeralStorageOk returns a tuple with the EphemeralStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimation) GetEphemeralStorageOk() (*int64, bool) {
	if o == nil || o.EphemeralStorage == nil {
		return nil, false
	}
	return o.EphemeralStorage, true
}

// HasEphemeralStorage returns a boolean if a field has been set.
func (o *Estimation) HasEphemeralStorage() bool {
	return o != nil && o.EphemeralStorage != nil
}

// SetEphemeralStorage gets a reference to the given int64 and assigns it to the EphemeralStorage field.
func (o *Estimation) SetEphemeralStorage(v int64) {
	o.EphemeralStorage = &v
}

// GetHeap returns the Heap field value if set, zero value otherwise.
func (o *Estimation) GetHeap() int64 {
	if o == nil || o.Heap == nil {
		var ret int64
		return ret
	}
	return *o.Heap
}

// GetHeapOk returns a tuple with the Heap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimation) GetHeapOk() (*int64, bool) {
	if o == nil || o.Heap == nil {
		return nil, false
	}
	return o.Heap, true
}

// HasHeap returns a boolean if a field has been set.
func (o *Estimation) HasHeap() bool {
	return o != nil && o.Heap != nil
}

// SetHeap gets a reference to the given int64 and assigns it to the Heap field.
func (o *Estimation) SetHeap(v int64) {
	o.Heap = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Estimation) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimation) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Estimation) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *Estimation) SetMemory(v int64) {
	o.Memory = &v
}

// GetOverhead returns the Overhead field value if set, zero value otherwise.
func (o *Estimation) GetOverhead() int64 {
	if o == nil || o.Overhead == nil {
		var ret int64
		return ret
	}
	return *o.Overhead
}

// GetOverheadOk returns a tuple with the Overhead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Estimation) GetOverheadOk() (*int64, bool) {
	if o == nil || o.Overhead == nil {
		return nil, false
	}
	return o.Overhead, true
}

// HasOverhead returns a boolean if a field has been set.
func (o *Estimation) HasOverhead() bool {
	return o != nil && o.Overhead != nil
}

// SetOverhead gets a reference to the given int64 and assigns it to the Overhead field.
func (o *Estimation) SetOverhead(v int64) {
	o.Overhead = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Estimation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.EphemeralStorage != nil {
		toSerialize["ephemeral_storage"] = o.EphemeralStorage
	}
	if o.Heap != nil {
		toSerialize["heap"] = o.Heap
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Overhead != nil {
		toSerialize["overhead"] = o.Overhead
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Estimation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cpu              *Cpu   `json:"cpu,omitempty"`
		EphemeralStorage *int64 `json:"ephemeral_storage,omitempty"`
		Heap             *int64 `json:"heap,omitempty"`
		Memory           *int64 `json:"memory,omitempty"`
		Overhead         *int64 `json:"overhead,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cpu", "ephemeral_storage", "heap", "memory", "overhead"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Cpu != nil && all.Cpu.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cpu = all.Cpu
	o.EphemeralStorage = all.EphemeralStorage
	o.Heap = all.Heap
	o.Memory = all.Memory
	o.Overhead = all.Overhead

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
