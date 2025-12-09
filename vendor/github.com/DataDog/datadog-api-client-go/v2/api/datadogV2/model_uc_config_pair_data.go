// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UCConfigPairData The definition of `UCConfigPairData` object.
type UCConfigPairData struct {
	// The definition of `UCConfigPairDataAttributes` object.
	Attributes *UCConfigPairDataAttributes `json:"attributes,omitempty"`
	// The `UCConfigPairData` `id`.
	Id *string `json:"id,omitempty"`
	// Azure UC configs resource type.
	Type UCConfigPairDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUCConfigPairData instantiates a new UCConfigPairData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUCConfigPairData(typeVar UCConfigPairDataType) *UCConfigPairData {
	this := UCConfigPairData{}
	this.Type = typeVar
	return &this
}

// NewUCConfigPairDataWithDefaults instantiates a new UCConfigPairData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUCConfigPairDataWithDefaults() *UCConfigPairData {
	this := UCConfigPairData{}
	var typeVar UCConfigPairDataType = UCCONFIGPAIRDATATYPE_AZURE_UC_CONFIGS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UCConfigPairData) GetAttributes() UCConfigPairDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret UCConfigPairDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairData) GetAttributesOk() (*UCConfigPairDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UCConfigPairData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UCConfigPairDataAttributes and assigns it to the Attributes field.
func (o *UCConfigPairData) SetAttributes(v UCConfigPairDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UCConfigPairData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UCConfigPairData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UCConfigPairData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UCConfigPairData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *UCConfigPairData) GetType() UCConfigPairDataType {
	if o == nil {
		var ret UCConfigPairDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UCConfigPairData) GetTypeOk() (*UCConfigPairDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UCConfigPairData) SetType(v UCConfigPairDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UCConfigPairData) MarshalJSON() ([]byte, error) {
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
func (o *UCConfigPairData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UCConfigPairDataAttributes `json:"attributes,omitempty"`
		Id         *string                     `json:"id,omitempty"`
		Type       *UCConfigPairDataType       `json:"type"`
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
