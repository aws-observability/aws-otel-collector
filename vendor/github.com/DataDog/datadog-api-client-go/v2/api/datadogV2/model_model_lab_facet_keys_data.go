// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabFacetKeysData A facet keys JSON:API resource object.
type ModelLabFacetKeysData struct {
	// Available facet key names for filtering resources.
	Attributes ModelLabFacetKeysAttributes `json:"attributes"`
	// The unique identifier of the facet keys resource.
	Id string `json:"id"`
	// The JSON:API type for a facet keys resource.
	Type ModelLabFacetKeysType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabFacetKeysData instantiates a new ModelLabFacetKeysData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabFacetKeysData(attributes ModelLabFacetKeysAttributes, id string, typeVar ModelLabFacetKeysType) *ModelLabFacetKeysData {
	this := ModelLabFacetKeysData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewModelLabFacetKeysDataWithDefaults instantiates a new ModelLabFacetKeysData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabFacetKeysDataWithDefaults() *ModelLabFacetKeysData {
	this := ModelLabFacetKeysData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ModelLabFacetKeysData) GetAttributes() ModelLabFacetKeysAttributes {
	if o == nil {
		var ret ModelLabFacetKeysAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ModelLabFacetKeysData) GetAttributesOk() (*ModelLabFacetKeysAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ModelLabFacetKeysData) SetAttributes(v ModelLabFacetKeysAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ModelLabFacetKeysData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelLabFacetKeysData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ModelLabFacetKeysData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ModelLabFacetKeysData) GetType() ModelLabFacetKeysType {
	if o == nil {
		var ret ModelLabFacetKeysType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ModelLabFacetKeysData) GetTypeOk() (*ModelLabFacetKeysType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ModelLabFacetKeysData) SetType(v ModelLabFacetKeysType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabFacetKeysData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabFacetKeysData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ModelLabFacetKeysAttributes `json:"attributes"`
		Id         *string                      `json:"id"`
		Type       *ModelLabFacetKeysType       `json:"type"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
