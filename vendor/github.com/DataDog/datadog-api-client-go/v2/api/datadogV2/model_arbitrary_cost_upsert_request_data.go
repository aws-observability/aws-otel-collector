// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryCostUpsertRequestData The definition of `ArbitraryCostUpsertRequestData` object.
type ArbitraryCostUpsertRequestData struct {
	// The definition of `ArbitraryCostUpsertRequestDataAttributes` object.
	Attributes *ArbitraryCostUpsertRequestDataAttributes `json:"attributes,omitempty"`
	// The `ArbitraryCostUpsertRequestData` `id`.
	Id *string `json:"id,omitempty"`
	// Upsert arbitrary rule resource type.
	Type ArbitraryCostUpsertRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArbitraryCostUpsertRequestData instantiates a new ArbitraryCostUpsertRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArbitraryCostUpsertRequestData(typeVar ArbitraryCostUpsertRequestDataType) *ArbitraryCostUpsertRequestData {
	this := ArbitraryCostUpsertRequestData{}
	this.Type = typeVar
	return &this
}

// NewArbitraryCostUpsertRequestDataWithDefaults instantiates a new ArbitraryCostUpsertRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArbitraryCostUpsertRequestDataWithDefaults() *ArbitraryCostUpsertRequestData {
	this := ArbitraryCostUpsertRequestData{}
	var typeVar ArbitraryCostUpsertRequestDataType = ARBITRARYCOSTUPSERTREQUESTDATATYPE_UPSERT_ARBITRARY_RULE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ArbitraryCostUpsertRequestData) GetAttributes() ArbitraryCostUpsertRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ArbitraryCostUpsertRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestData) GetAttributesOk() (*ArbitraryCostUpsertRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ArbitraryCostUpsertRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ArbitraryCostUpsertRequestDataAttributes and assigns it to the Attributes field.
func (o *ArbitraryCostUpsertRequestData) SetAttributes(v ArbitraryCostUpsertRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArbitraryCostUpsertRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArbitraryCostUpsertRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ArbitraryCostUpsertRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *ArbitraryCostUpsertRequestData) GetType() ArbitraryCostUpsertRequestDataType {
	if o == nil {
		var ret ArbitraryCostUpsertRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestData) GetTypeOk() (*ArbitraryCostUpsertRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ArbitraryCostUpsertRequestData) SetType(v ArbitraryCostUpsertRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArbitraryCostUpsertRequestData) MarshalJSON() ([]byte, error) {
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
func (o *ArbitraryCostUpsertRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ArbitraryCostUpsertRequestDataAttributes `json:"attributes,omitempty"`
		Id         *string                                   `json:"id,omitempty"`
		Type       *ArbitraryCostUpsertRequestDataType       `json:"type"`
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
