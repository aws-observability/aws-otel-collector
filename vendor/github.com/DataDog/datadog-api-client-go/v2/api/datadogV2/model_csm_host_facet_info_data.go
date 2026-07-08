// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmHostFacetInfoData The data wrapper for a facet info response.
type CsmHostFacetInfoData struct {
	// Attributes of a facet info response, containing the value distribution for the requested facet.
	Attributes CsmHostFacetInfoAttributes `json:"attributes"`
	// The identifier of the facet.
	Id string `json:"id"`
	// Metadata for the facet info response.
	Meta CsmHostFacetInfoMeta `json:"meta"`
	// The JSON:API type for facet info resources. The value should always be `facet_info`.
	Type CsmFacetInfoType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmHostFacetInfoData instantiates a new CsmHostFacetInfoData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmHostFacetInfoData(attributes CsmHostFacetInfoAttributes, id string, meta CsmHostFacetInfoMeta, typeVar CsmFacetInfoType) *CsmHostFacetInfoData {
	this := CsmHostFacetInfoData{}
	this.Attributes = attributes
	this.Id = id
	this.Meta = meta
	this.Type = typeVar
	return &this
}

// NewCsmHostFacetInfoDataWithDefaults instantiates a new CsmHostFacetInfoData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmHostFacetInfoDataWithDefaults() *CsmHostFacetInfoData {
	this := CsmHostFacetInfoData{}
	var typeVar CsmFacetInfoType = CSMFACETINFOTYPE_FACET_INFO
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CsmHostFacetInfoData) GetAttributes() CsmHostFacetInfoAttributes {
	if o == nil {
		var ret CsmHostFacetInfoAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CsmHostFacetInfoData) GetAttributesOk() (*CsmHostFacetInfoAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CsmHostFacetInfoData) SetAttributes(v CsmHostFacetInfoAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *CsmHostFacetInfoData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CsmHostFacetInfoData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CsmHostFacetInfoData) SetId(v string) {
	o.Id = v
}

// GetMeta returns the Meta field value.
func (o *CsmHostFacetInfoData) GetMeta() CsmHostFacetInfoMeta {
	if o == nil {
		var ret CsmHostFacetInfoMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *CsmHostFacetInfoData) GetMetaOk() (*CsmHostFacetInfoMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *CsmHostFacetInfoData) SetMeta(v CsmHostFacetInfoMeta) {
	o.Meta = v
}

// GetType returns the Type field value.
func (o *CsmHostFacetInfoData) GetType() CsmFacetInfoType {
	if o == nil {
		var ret CsmFacetInfoType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CsmHostFacetInfoData) GetTypeOk() (*CsmFacetInfoType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CsmHostFacetInfoData) SetType(v CsmFacetInfoType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmHostFacetInfoData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["meta"] = o.Meta
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmHostFacetInfoData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CsmHostFacetInfoAttributes `json:"attributes"`
		Id         *string                     `json:"id"`
		Meta       *CsmHostFacetInfoMeta       `json:"meta"`
		Type       *CsmFacetInfoType           `json:"type"`
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
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "meta", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta
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
