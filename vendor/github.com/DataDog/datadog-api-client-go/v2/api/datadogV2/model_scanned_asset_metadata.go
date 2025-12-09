// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScannedAssetMetadata The metadata of a scanned asset.
type ScannedAssetMetadata struct {
	// The attributes of a scanned asset metadata.
	Attributes ScannedAssetMetadataAttributes `json:"attributes"`
	// The ID of the scanned asset metadata.
	Id string `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScannedAssetMetadata instantiates a new ScannedAssetMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScannedAssetMetadata(attributes ScannedAssetMetadataAttributes, id string) *ScannedAssetMetadata {
	this := ScannedAssetMetadata{}
	this.Attributes = attributes
	this.Id = id
	return &this
}

// NewScannedAssetMetadataWithDefaults instantiates a new ScannedAssetMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScannedAssetMetadataWithDefaults() *ScannedAssetMetadata {
	this := ScannedAssetMetadata{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ScannedAssetMetadata) GetAttributes() ScannedAssetMetadataAttributes {
	if o == nil {
		var ret ScannedAssetMetadataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ScannedAssetMetadata) GetAttributesOk() (*ScannedAssetMetadataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ScannedAssetMetadata) SetAttributes(v ScannedAssetMetadataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ScannedAssetMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScannedAssetMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ScannedAssetMetadata) SetId(v string) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScannedAssetMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScannedAssetMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ScannedAssetMetadataAttributes `json:"attributes"`
		Id         *string                         `json:"id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
