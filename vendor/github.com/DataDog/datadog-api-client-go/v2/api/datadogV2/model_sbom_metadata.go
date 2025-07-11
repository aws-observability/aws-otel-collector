// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SBOMMetadata Provides additional information about a BOM.
type SBOMMetadata struct {
	// The component that the BOM describes.
	Component *SBOMMetadataComponent `json:"component,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSBOMMetadata instantiates a new SBOMMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSBOMMetadata() *SBOMMetadata {
	this := SBOMMetadata{}
	return &this
}

// NewSBOMMetadataWithDefaults instantiates a new SBOMMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSBOMMetadataWithDefaults() *SBOMMetadata {
	this := SBOMMetadata{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *SBOMMetadata) GetComponent() SBOMMetadataComponent {
	if o == nil || o.Component == nil {
		var ret SBOMMetadataComponent
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBOMMetadata) GetComponentOk() (*SBOMMetadataComponent, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *SBOMMetadata) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given SBOMMetadataComponent and assigns it to the Component field.
func (o *SBOMMetadata) SetComponent(v SBOMMetadataComponent) {
	o.Component = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SBOMMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SBOMMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *SBOMMetadataComponent `json:"component,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"component"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Component != nil && all.Component.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Component = all.Component

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
