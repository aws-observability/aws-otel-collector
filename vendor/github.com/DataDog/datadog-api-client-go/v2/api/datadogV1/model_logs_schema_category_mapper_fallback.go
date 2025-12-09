// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSchemaCategoryMapperFallback Used to override hardcoded category values with a value pulled from a source attribute on the log.
type LogsSchemaCategoryMapperFallback struct {
	// Fallback sources used to populate value of field.
	Sources map[string][]string `json:"sources,omitempty"`
	// Values that define when the fallback is used.
	Values map[string]string `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsSchemaCategoryMapperFallback instantiates a new LogsSchemaCategoryMapperFallback object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsSchemaCategoryMapperFallback() *LogsSchemaCategoryMapperFallback {
	this := LogsSchemaCategoryMapperFallback{}
	return &this
}

// NewLogsSchemaCategoryMapperFallbackWithDefaults instantiates a new LogsSchemaCategoryMapperFallback object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsSchemaCategoryMapperFallbackWithDefaults() *LogsSchemaCategoryMapperFallback {
	this := LogsSchemaCategoryMapperFallback{}
	return &this
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *LogsSchemaCategoryMapperFallback) GetSources() map[string][]string {
	if o == nil || o.Sources == nil {
		var ret map[string][]string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsSchemaCategoryMapperFallback) GetSourcesOk() (*map[string][]string, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *LogsSchemaCategoryMapperFallback) HasSources() bool {
	return o != nil && o.Sources != nil
}

// SetSources gets a reference to the given map[string][]string and assigns it to the Sources field.
func (o *LogsSchemaCategoryMapperFallback) SetSources(v map[string][]string) {
	o.Sources = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *LogsSchemaCategoryMapperFallback) GetValues() map[string]string {
	if o == nil || o.Values == nil {
		var ret map[string]string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsSchemaCategoryMapperFallback) GetValuesOk() (*map[string]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *LogsSchemaCategoryMapperFallback) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given map[string]string and assigns it to the Values field.
func (o *LogsSchemaCategoryMapperFallback) SetValues(v map[string]string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsSchemaCategoryMapperFallback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsSchemaCategoryMapperFallback) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Sources map[string][]string `json:"sources,omitempty"`
		Values  map[string]string   `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"sources", "values"})
	} else {
		return err
	}
	o.Sources = all.Sources
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
