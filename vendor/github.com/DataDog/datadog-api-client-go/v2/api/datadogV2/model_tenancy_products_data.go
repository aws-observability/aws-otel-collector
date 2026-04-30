// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TenancyProductsData A single OCI tenancy product resource object containing the tenancy ID, type, and product attributes.
type TenancyProductsData struct {
	// Attributes of an OCI tenancy product resource, containing the list of available products and their enablement status.
	Attributes *TenancyProductsDataAttributes `json:"attributes,omitempty"`
	// The OCID of the OCI tenancy.
	Id *string `json:"id,omitempty"`
	// OCI tenancy product resource type.
	Type TenancyProductsDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTenancyProductsData instantiates a new TenancyProductsData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTenancyProductsData(typeVar TenancyProductsDataType) *TenancyProductsData {
	this := TenancyProductsData{}
	this.Type = typeVar
	return &this
}

// NewTenancyProductsDataWithDefaults instantiates a new TenancyProductsData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTenancyProductsDataWithDefaults() *TenancyProductsData {
	this := TenancyProductsData{}
	var typeVar TenancyProductsDataType = TENANCYPRODUCTSDATATYPE_OCI_TENANCY_PRODUCT
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TenancyProductsData) GetAttributes() TenancyProductsDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TenancyProductsDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyProductsData) GetAttributesOk() (*TenancyProductsDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TenancyProductsData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given TenancyProductsDataAttributes and assigns it to the Attributes field.
func (o *TenancyProductsData) SetAttributes(v TenancyProductsDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TenancyProductsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyProductsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TenancyProductsData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TenancyProductsData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *TenancyProductsData) GetType() TenancyProductsDataType {
	if o == nil {
		var ret TenancyProductsDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TenancyProductsData) GetTypeOk() (*TenancyProductsDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TenancyProductsData) SetType(v TenancyProductsDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TenancyProductsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TenancyProductsData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *TenancyProductsDataAttributes `json:"attributes,omitempty"`
		Id         *string                        `json:"id,omitempty"`
		Type       *TenancyProductsDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
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
