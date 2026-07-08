// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CycloneDXMetadata Metadata about the BOM, including the scanned asset and the scanner tool.
type CycloneDXMetadata struct {
	// The asset that was scanned (for example, a host or container image).
	Component CycloneDXMetadataComponent `json:"component"`
	// Information about the scanner tool that produced this BOM.
	Tools CycloneDXMetadataTools `json:"tools"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCycloneDXMetadata instantiates a new CycloneDXMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCycloneDXMetadata(component CycloneDXMetadataComponent, tools CycloneDXMetadataTools) *CycloneDXMetadata {
	this := CycloneDXMetadata{}
	this.Component = component
	this.Tools = tools
	return &this
}

// NewCycloneDXMetadataWithDefaults instantiates a new CycloneDXMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCycloneDXMetadataWithDefaults() *CycloneDXMetadata {
	this := CycloneDXMetadata{}
	return &this
}

// GetComponent returns the Component field value.
func (o *CycloneDXMetadata) GetComponent() CycloneDXMetadataComponent {
	if o == nil {
		var ret CycloneDXMetadataComponent
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *CycloneDXMetadata) GetComponentOk() (*CycloneDXMetadataComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *CycloneDXMetadata) SetComponent(v CycloneDXMetadataComponent) {
	o.Component = v
}

// GetTools returns the Tools field value.
func (o *CycloneDXMetadata) GetTools() CycloneDXMetadataTools {
	if o == nil {
		var ret CycloneDXMetadataTools
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value
// and a boolean to check if the value has been set.
func (o *CycloneDXMetadata) GetToolsOk() (*CycloneDXMetadataTools, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tools, true
}

// SetTools sets field value.
func (o *CycloneDXMetadata) SetTools(v CycloneDXMetadataTools) {
	o.Tools = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CycloneDXMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["tools"] = o.Tools

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CycloneDXMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *CycloneDXMetadataComponent `json:"component"`
		Tools     *CycloneDXMetadataTools     `json:"tools"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Tools == nil {
		return fmt.Errorf("required field tools missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"component", "tools"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Component.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Component = *all.Component
	if all.Tools.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tools = *all.Tools

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
