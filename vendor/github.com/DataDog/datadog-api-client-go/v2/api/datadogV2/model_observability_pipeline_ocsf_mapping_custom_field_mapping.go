// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOcsfMappingCustomFieldMapping Defines a single field mapping rule for transforming a source field to an OCSF destination field.
type ObservabilityPipelineOcsfMappingCustomFieldMapping struct {
	// The default value to use if the source field is missing or empty.
	Default interface{} `json:"default,omitempty"`
	// The destination OCSF field path.
	Dest string `json:"dest"`
	// Lookup table configuration for mapping source values to destination values.
	Lookup *ObservabilityPipelineOcsfMappingCustomLookup `json:"lookup,omitempty"`
	// The source field path from the log event.
	Source interface{} `json:"source,omitempty"`
	// Multiple source field paths for combined mapping.
	Sources interface{} `json:"sources,omitempty"`
	// A static value to use for the destination field.
	Value interface{} `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineOcsfMappingCustomFieldMapping instantiates a new ObservabilityPipelineOcsfMappingCustomFieldMapping object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineOcsfMappingCustomFieldMapping(dest string) *ObservabilityPipelineOcsfMappingCustomFieldMapping {
	this := ObservabilityPipelineOcsfMappingCustomFieldMapping{}
	this.Dest = dest
	return &this
}

// NewObservabilityPipelineOcsfMappingCustomFieldMappingWithDefaults instantiates a new ObservabilityPipelineOcsfMappingCustomFieldMapping object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineOcsfMappingCustomFieldMappingWithDefaults() *ObservabilityPipelineOcsfMappingCustomFieldMapping {
	this := ObservabilityPipelineOcsfMappingCustomFieldMapping{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetDefault() interface{} {
	if o == nil || o.Default == nil {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) SetDefault(v interface{}) {
	o.Default = v
}

// GetDest returns the Dest field value.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetDest() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Dest
}

// GetDestOk returns a tuple with the Dest field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetDestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dest, true
}

// SetDest sets field value.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) SetDest(v string) {
	o.Dest = v
}

// GetLookup returns the Lookup field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetLookup() ObservabilityPipelineOcsfMappingCustomLookup {
	if o == nil || o.Lookup == nil {
		var ret ObservabilityPipelineOcsfMappingCustomLookup
		return ret
	}
	return *o.Lookup
}

// GetLookupOk returns a tuple with the Lookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetLookupOk() (*ObservabilityPipelineOcsfMappingCustomLookup, bool) {
	if o == nil || o.Lookup == nil {
		return nil, false
	}
	return o.Lookup, true
}

// HasLookup returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) HasLookup() bool {
	return o != nil && o.Lookup != nil
}

// SetLookup gets a reference to the given ObservabilityPipelineOcsfMappingCustomLookup and assigns it to the Lookup field.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) SetLookup(v ObservabilityPipelineOcsfMappingCustomLookup) {
	o.Lookup = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetSource() interface{} {
	if o == nil || o.Source == nil {
		var ret interface{}
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetSourceOk() (*interface{}, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return &o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given interface{} and assigns it to the Source field.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) SetSource(v interface{}) {
	o.Source = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetSources() interface{} {
	if o == nil || o.Sources == nil {
		var ret interface{}
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetSourcesOk() (*interface{}, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) HasSources() bool {
	return o != nil && o.Sources != nil
}

// SetSources gets a reference to the given interface{} and assigns it to the Sources field.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) SetSources(v interface{}) {
	o.Sources = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) SetValue(v interface{}) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineOcsfMappingCustomFieldMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	toSerialize["dest"] = o.Dest
	if o.Lookup != nil {
		toSerialize["lookup"] = o.Lookup
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineOcsfMappingCustomFieldMapping) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default interface{}                                   `json:"default,omitempty"`
		Dest    *string                                       `json:"dest"`
		Lookup  *ObservabilityPipelineOcsfMappingCustomLookup `json:"lookup,omitempty"`
		Source  interface{}                                   `json:"source,omitempty"`
		Sources interface{}                                   `json:"sources,omitempty"`
		Value   interface{}                                   `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Dest == nil {
		return fmt.Errorf("required field dest missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default", "dest", "lookup", "source", "sources", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Default = all.Default
	o.Dest = *all.Dest
	if all.Lookup != nil && all.Lookup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Lookup = all.Lookup
	o.Source = all.Source
	o.Sources = all.Sources
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
