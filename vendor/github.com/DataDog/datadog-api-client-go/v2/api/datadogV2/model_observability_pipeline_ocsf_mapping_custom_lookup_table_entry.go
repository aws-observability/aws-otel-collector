// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOcsfMappingCustomLookupTableEntry A single entry in a lookup table for value transformation.
type ObservabilityPipelineOcsfMappingCustomLookupTableEntry struct {
	// The substring to match in the source value.
	Contains *string `json:"contains,omitempty"`
	// The exact value to match in the source.
	Equals interface{} `json:"equals,omitempty"`
	// The source field to match against.
	EqualsSource *string `json:"equals_source,omitempty"`
	// A regex pattern to match in the source value.
	Matches *string `json:"matches,omitempty"`
	// A regex pattern that must not match the source value.
	NotMatches *string `json:"not_matches,omitempty"`
	// The value to use when a match is found.
	Value interface{} `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineOcsfMappingCustomLookupTableEntry instantiates a new ObservabilityPipelineOcsfMappingCustomLookupTableEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineOcsfMappingCustomLookupTableEntry() *ObservabilityPipelineOcsfMappingCustomLookupTableEntry {
	this := ObservabilityPipelineOcsfMappingCustomLookupTableEntry{}
	return &this
}

// NewObservabilityPipelineOcsfMappingCustomLookupTableEntryWithDefaults instantiates a new ObservabilityPipelineOcsfMappingCustomLookupTableEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineOcsfMappingCustomLookupTableEntryWithDefaults() *ObservabilityPipelineOcsfMappingCustomLookupTableEntry {
	this := ObservabilityPipelineOcsfMappingCustomLookupTableEntry{}
	return &this
}

// GetContains returns the Contains field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetContains() string {
	if o == nil || o.Contains == nil {
		var ret string
		return ret
	}
	return *o.Contains
}

// GetContainsOk returns a tuple with the Contains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetContainsOk() (*string, bool) {
	if o == nil || o.Contains == nil {
		return nil, false
	}
	return o.Contains, true
}

// HasContains returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) HasContains() bool {
	return o != nil && o.Contains != nil
}

// SetContains gets a reference to the given string and assigns it to the Contains field.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) SetContains(v string) {
	o.Contains = &v
}

// GetEquals returns the Equals field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetEquals() interface{} {
	if o == nil || o.Equals == nil {
		var ret interface{}
		return ret
	}
	return o.Equals
}

// GetEqualsOk returns a tuple with the Equals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetEqualsOk() (*interface{}, bool) {
	if o == nil || o.Equals == nil {
		return nil, false
	}
	return &o.Equals, true
}

// HasEquals returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) HasEquals() bool {
	return o != nil && o.Equals != nil
}

// SetEquals gets a reference to the given interface{} and assigns it to the Equals field.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) SetEquals(v interface{}) {
	o.Equals = v
}

// GetEqualsSource returns the EqualsSource field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetEqualsSource() string {
	if o == nil || o.EqualsSource == nil {
		var ret string
		return ret
	}
	return *o.EqualsSource
}

// GetEqualsSourceOk returns a tuple with the EqualsSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetEqualsSourceOk() (*string, bool) {
	if o == nil || o.EqualsSource == nil {
		return nil, false
	}
	return o.EqualsSource, true
}

// HasEqualsSource returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) HasEqualsSource() bool {
	return o != nil && o.EqualsSource != nil
}

// SetEqualsSource gets a reference to the given string and assigns it to the EqualsSource field.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) SetEqualsSource(v string) {
	o.EqualsSource = &v
}

// GetMatches returns the Matches field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetMatches() string {
	if o == nil || o.Matches == nil {
		var ret string
		return ret
	}
	return *o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetMatchesOk() (*string, bool) {
	if o == nil || o.Matches == nil {
		return nil, false
	}
	return o.Matches, true
}

// HasMatches returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) HasMatches() bool {
	return o != nil && o.Matches != nil
}

// SetMatches gets a reference to the given string and assigns it to the Matches field.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) SetMatches(v string) {
	o.Matches = &v
}

// GetNotMatches returns the NotMatches field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetNotMatches() string {
	if o == nil || o.NotMatches == nil {
		var ret string
		return ret
	}
	return *o.NotMatches
}

// GetNotMatchesOk returns a tuple with the NotMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetNotMatchesOk() (*string, bool) {
	if o == nil || o.NotMatches == nil {
		return nil, false
	}
	return o.NotMatches, true
}

// HasNotMatches returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) HasNotMatches() bool {
	return o != nil && o.NotMatches != nil
}

// SetNotMatches gets a reference to the given string and assigns it to the NotMatches field.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) SetNotMatches(v string) {
	o.NotMatches = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) SetValue(v interface{}) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineOcsfMappingCustomLookupTableEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Contains != nil {
		toSerialize["contains"] = o.Contains
	}
	if o.Equals != nil {
		toSerialize["equals"] = o.Equals
	}
	if o.EqualsSource != nil {
		toSerialize["equals_source"] = o.EqualsSource
	}
	if o.Matches != nil {
		toSerialize["matches"] = o.Matches
	}
	if o.NotMatches != nil {
		toSerialize["not_matches"] = o.NotMatches
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
func (o *ObservabilityPipelineOcsfMappingCustomLookupTableEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Contains     *string     `json:"contains,omitempty"`
		Equals       interface{} `json:"equals,omitempty"`
		EqualsSource *string     `json:"equals_source,omitempty"`
		Matches      *string     `json:"matches,omitempty"`
		NotMatches   *string     `json:"not_matches,omitempty"`
		Value        interface{} `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"contains", "equals", "equals_source", "matches", "not_matches", "value"})
	} else {
		return err
	}
	o.Contains = all.Contains
	o.Equals = all.Equals
	o.EqualsSource = all.EqualsSource
	o.Matches = all.Matches
	o.NotMatches = all.NotMatches
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
