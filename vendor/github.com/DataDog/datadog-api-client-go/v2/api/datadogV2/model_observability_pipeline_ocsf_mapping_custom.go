// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOcsfMappingCustom Custom OCSF mapping configuration for transforming logs.
type ObservabilityPipelineOcsfMappingCustom struct {
	// A list of field mapping rules for transforming log fields to OCSF schema fields.
	Mapping []ObservabilityPipelineOcsfMappingCustomFieldMapping `json:"mapping"`
	// Metadata for the custom OCSF mapping.
	Metadata ObservabilityPipelineOcsfMappingCustomMetadata `json:"metadata"`
	// The version of the custom mapping configuration.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineOcsfMappingCustom instantiates a new ObservabilityPipelineOcsfMappingCustom object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineOcsfMappingCustom(mapping []ObservabilityPipelineOcsfMappingCustomFieldMapping, metadata ObservabilityPipelineOcsfMappingCustomMetadata, version int64) *ObservabilityPipelineOcsfMappingCustom {
	this := ObservabilityPipelineOcsfMappingCustom{}
	this.Mapping = mapping
	this.Metadata = metadata
	this.Version = version
	return &this
}

// NewObservabilityPipelineOcsfMappingCustomWithDefaults instantiates a new ObservabilityPipelineOcsfMappingCustom object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineOcsfMappingCustomWithDefaults() *ObservabilityPipelineOcsfMappingCustom {
	this := ObservabilityPipelineOcsfMappingCustom{}
	return &this
}

// GetMapping returns the Mapping field value.
func (o *ObservabilityPipelineOcsfMappingCustom) GetMapping() []ObservabilityPipelineOcsfMappingCustomFieldMapping {
	if o == nil {
		var ret []ObservabilityPipelineOcsfMappingCustomFieldMapping
		return ret
	}
	return o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustom) GetMappingOk() (*[]ObservabilityPipelineOcsfMappingCustomFieldMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapping, true
}

// SetMapping sets field value.
func (o *ObservabilityPipelineOcsfMappingCustom) SetMapping(v []ObservabilityPipelineOcsfMappingCustomFieldMapping) {
	o.Mapping = v
}

// GetMetadata returns the Metadata field value.
func (o *ObservabilityPipelineOcsfMappingCustom) GetMetadata() ObservabilityPipelineOcsfMappingCustomMetadata {
	if o == nil {
		var ret ObservabilityPipelineOcsfMappingCustomMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustom) GetMetadataOk() (*ObservabilityPipelineOcsfMappingCustomMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value.
func (o *ObservabilityPipelineOcsfMappingCustom) SetMetadata(v ObservabilityPipelineOcsfMappingCustomMetadata) {
	o.Metadata = v
}

// GetVersion returns the Version field value.
func (o *ObservabilityPipelineOcsfMappingCustom) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOcsfMappingCustom) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *ObservabilityPipelineOcsfMappingCustom) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineOcsfMappingCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["mapping"] = o.Mapping
	toSerialize["metadata"] = o.Metadata
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineOcsfMappingCustom) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mapping  *[]ObservabilityPipelineOcsfMappingCustomFieldMapping `json:"mapping"`
		Metadata *ObservabilityPipelineOcsfMappingCustomMetadata       `json:"metadata"`
		Version  *int64                                                `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mapping == nil {
		return fmt.Errorf("required field mapping missing")
	}
	if all.Metadata == nil {
		return fmt.Errorf("required field metadata missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"mapping", "metadata", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Mapping = *all.Mapping
	if all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = *all.Metadata
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
