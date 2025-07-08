// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableFile Defines a static enrichment table loaded from a CSV file.
type ObservabilityPipelineEnrichmentTableFile struct {
	// File encoding format.
	Encoding ObservabilityPipelineEnrichmentTableFileEncoding `json:"encoding"`
	// Key fields used to look up enrichment values.
	Key []ObservabilityPipelineEnrichmentTableFileKeyItems `json:"key"`
	// Path to the CSV file.
	Path string `json:"path"`
	// Schema defining column names and their types.
	Schema []ObservabilityPipelineEnrichmentTableFileSchemaItems `json:"schema"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineEnrichmentTableFile instantiates a new ObservabilityPipelineEnrichmentTableFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineEnrichmentTableFile(encoding ObservabilityPipelineEnrichmentTableFileEncoding, key []ObservabilityPipelineEnrichmentTableFileKeyItems, path string, schema []ObservabilityPipelineEnrichmentTableFileSchemaItems) *ObservabilityPipelineEnrichmentTableFile {
	this := ObservabilityPipelineEnrichmentTableFile{}
	this.Encoding = encoding
	this.Key = key
	this.Path = path
	this.Schema = schema
	return &this
}

// NewObservabilityPipelineEnrichmentTableFileWithDefaults instantiates a new ObservabilityPipelineEnrichmentTableFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineEnrichmentTableFileWithDefaults() *ObservabilityPipelineEnrichmentTableFile {
	this := ObservabilityPipelineEnrichmentTableFile{}
	return &this
}

// GetEncoding returns the Encoding field value.
func (o *ObservabilityPipelineEnrichmentTableFile) GetEncoding() ObservabilityPipelineEnrichmentTableFileEncoding {
	if o == nil {
		var ret ObservabilityPipelineEnrichmentTableFileEncoding
		return ret
	}
	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFile) GetEncodingOk() (*ObservabilityPipelineEnrichmentTableFileEncoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value.
func (o *ObservabilityPipelineEnrichmentTableFile) SetEncoding(v ObservabilityPipelineEnrichmentTableFileEncoding) {
	o.Encoding = v
}

// GetKey returns the Key field value.
func (o *ObservabilityPipelineEnrichmentTableFile) GetKey() []ObservabilityPipelineEnrichmentTableFileKeyItems {
	if o == nil {
		var ret []ObservabilityPipelineEnrichmentTableFileKeyItems
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFile) GetKeyOk() (*[]ObservabilityPipelineEnrichmentTableFileKeyItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *ObservabilityPipelineEnrichmentTableFile) SetKey(v []ObservabilityPipelineEnrichmentTableFileKeyItems) {
	o.Key = v
}

// GetPath returns the Path field value.
func (o *ObservabilityPipelineEnrichmentTableFile) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFile) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *ObservabilityPipelineEnrichmentTableFile) SetPath(v string) {
	o.Path = v
}

// GetSchema returns the Schema field value.
func (o *ObservabilityPipelineEnrichmentTableFile) GetSchema() []ObservabilityPipelineEnrichmentTableFileSchemaItems {
	if o == nil {
		var ret []ObservabilityPipelineEnrichmentTableFileSchemaItems
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFile) GetSchemaOk() (*[]ObservabilityPipelineEnrichmentTableFileSchemaItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value.
func (o *ObservabilityPipelineEnrichmentTableFile) SetSchema(v []ObservabilityPipelineEnrichmentTableFileSchemaItems) {
	o.Schema = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineEnrichmentTableFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["encoding"] = o.Encoding
	toSerialize["key"] = o.Key
	toSerialize["path"] = o.Path
	toSerialize["schema"] = o.Schema

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineEnrichmentTableFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Encoding *ObservabilityPipelineEnrichmentTableFileEncoding      `json:"encoding"`
		Key      *[]ObservabilityPipelineEnrichmentTableFileKeyItems    `json:"key"`
		Path     *string                                                `json:"path"`
		Schema   *[]ObservabilityPipelineEnrichmentTableFileSchemaItems `json:"schema"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Encoding == nil {
		return fmt.Errorf("required field encoding missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	if all.Schema == nil {
		return fmt.Errorf("required field schema missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"encoding", "key", "path", "schema"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Encoding.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Encoding = *all.Encoding
	o.Key = *all.Key
	o.Path = *all.Path
	o.Schema = *all.Schema

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
