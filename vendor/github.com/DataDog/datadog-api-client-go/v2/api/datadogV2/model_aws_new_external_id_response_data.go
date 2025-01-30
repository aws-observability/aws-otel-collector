// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSNewExternalIDResponseData AWS External ID response body.
type AWSNewExternalIDResponseData struct {
	// AWS External ID response body.
	Attributes *AWSNewExternalIDResponseAttributes `json:"attributes,omitempty"`
	// The `AWSNewExternalIDResponseData` `id`.
	Id string `json:"id"`
	// The `AWSNewExternalIDResponseData` `type`.
	Type AWSNewExternalIDResponseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSNewExternalIDResponseData instantiates a new AWSNewExternalIDResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSNewExternalIDResponseData(id string, typeVar AWSNewExternalIDResponseDataType) *AWSNewExternalIDResponseData {
	this := AWSNewExternalIDResponseData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewAWSNewExternalIDResponseDataWithDefaults instantiates a new AWSNewExternalIDResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSNewExternalIDResponseDataWithDefaults() *AWSNewExternalIDResponseData {
	this := AWSNewExternalIDResponseData{}
	var id string = "external_id"
	this.Id = id
	var typeVar AWSNewExternalIDResponseDataType = AWSNEWEXTERNALIDRESPONSEDATATYPE_EXTERNAL_ID
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AWSNewExternalIDResponseData) GetAttributes() AWSNewExternalIDResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret AWSNewExternalIDResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSNewExternalIDResponseData) GetAttributesOk() (*AWSNewExternalIDResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AWSNewExternalIDResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given AWSNewExternalIDResponseAttributes and assigns it to the Attributes field.
func (o *AWSNewExternalIDResponseData) SetAttributes(v AWSNewExternalIDResponseAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *AWSNewExternalIDResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AWSNewExternalIDResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AWSNewExternalIDResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *AWSNewExternalIDResponseData) GetType() AWSNewExternalIDResponseDataType {
	if o == nil {
		var ret AWSNewExternalIDResponseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSNewExternalIDResponseData) GetTypeOk() (*AWSNewExternalIDResponseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AWSNewExternalIDResponseData) SetType(v AWSNewExternalIDResponseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSNewExternalIDResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSNewExternalIDResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AWSNewExternalIDResponseAttributes `json:"attributes,omitempty"`
		Id         *string                             `json:"id"`
		Type       *AWSNewExternalIDResponseDataType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
