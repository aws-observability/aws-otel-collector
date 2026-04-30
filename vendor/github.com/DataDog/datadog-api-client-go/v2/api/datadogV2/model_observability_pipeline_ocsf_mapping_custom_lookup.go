// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOcsfMappingCustomLookup Lookup table configuration for mapping source values to destination values.
type ObservabilityPipelineOcsfMappingCustomLookup struct {
	// The default value to use if no lookup match is found.
	Default interface{} `json:"default,omitempty"`
	// A list of lookup table entries for value transformation.
	Table []ObservabilityPipelineOcsfMappingCustomLookupTableEntry `json:"table,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineOcsfMappingCustomLookup instantiates a new ObservabilityPipelineOcsfMappingCustomLookup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineOcsfMappingCustomLookup() *ObservabilityPipelineOcsfMappingCustomLookup {
	this := ObservabilityPipelineOcsfMappingCustomLookup{}
	return &this
}

// NewObservabilityPipelineOcsfMappingCustomLookupWithDefaults instantiates a new ObservabilityPipelineOcsfMappingCustomLookup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineOcsfMappingCustomLookupWithDefaults() *ObservabilityPipelineOcsfMappingCustomLookup {
	this := ObservabilityPipelineOcsfMappingCustomLookup{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomLookup) GetDefault() interface{} {
	if o == nil || o.Default == nil {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookup) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookup) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *ObservabilityPipelineOcsfMappingCustomLookup) SetDefault(v interface{}) {
	o.Default = v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomLookup) GetTable() []ObservabilityPipelineOcsfMappingCustomLookupTableEntry {
	if o == nil || o.Table == nil {
		var ret []ObservabilityPipelineOcsfMappingCustomLookupTableEntry
		return ret
	}
	return o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookup) GetTableOk() (*[]ObservabilityPipelineOcsfMappingCustomLookupTableEntry, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return &o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookup) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given []ObservabilityPipelineOcsfMappingCustomLookupTableEntry and assigns it to the Table field.
func (o *ObservabilityPipelineOcsfMappingCustomLookup) SetTable(v []ObservabilityPipelineOcsfMappingCustomLookupTableEntry) {
	o.Table = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineOcsfMappingCustomLookup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineOcsfMappingCustomLookup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default interface{}                                              `json:"default,omitempty"`
		Table   []ObservabilityPipelineOcsfMappingCustomLookupTableEntry `json:"table,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default", "table"})
	} else {
		return err
	}
	o.Default = all.Default
	o.Table = all.Table

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
