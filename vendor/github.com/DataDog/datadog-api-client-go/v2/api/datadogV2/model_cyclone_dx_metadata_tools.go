// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CycloneDXMetadataTools Information about the scanner tool that produced this BOM.
type CycloneDXMetadataTools struct {
	// The scanner tool components. Must contain exactly one element.
	Components []CycloneDXToolComponent `json:"components"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCycloneDXMetadataTools instantiates a new CycloneDXMetadataTools object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCycloneDXMetadataTools(components []CycloneDXToolComponent) *CycloneDXMetadataTools {
	this := CycloneDXMetadataTools{}
	this.Components = components
	return &this
}

// NewCycloneDXMetadataToolsWithDefaults instantiates a new CycloneDXMetadataTools object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCycloneDXMetadataToolsWithDefaults() *CycloneDXMetadataTools {
	this := CycloneDXMetadataTools{}
	return &this
}

// GetComponents returns the Components field value.
func (o *CycloneDXMetadataTools) GetComponents() []CycloneDXToolComponent {
	if o == nil {
		var ret []CycloneDXToolComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *CycloneDXMetadataTools) GetComponentsOk() (*[]CycloneDXToolComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value.
func (o *CycloneDXMetadataTools) SetComponents(v []CycloneDXToolComponent) {
	o.Components = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CycloneDXMetadataTools) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["components"] = o.Components

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CycloneDXMetadataTools) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components *[]CycloneDXToolComponent `json:"components"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Components == nil {
		return fmt.Errorf("required field components missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components"})
	} else {
		return err
	}
	o.Components = *all.Components

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
