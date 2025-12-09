// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCURConfigPostData AWS CUR config Post data.
type AwsCURConfigPostData struct {
	// Attributes for AWS CUR config Post Request.
	Attributes *AwsCURConfigPostRequestAttributes `json:"attributes,omitempty"`
	// Type of AWS CUR config Post Request.
	Type AwsCURConfigPostRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsCURConfigPostData instantiates a new AwsCURConfigPostData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsCURConfigPostData(typeVar AwsCURConfigPostRequestType) *AwsCURConfigPostData {
	this := AwsCURConfigPostData{}
	this.Type = typeVar
	return &this
}

// NewAwsCURConfigPostDataWithDefaults instantiates a new AwsCURConfigPostData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsCURConfigPostDataWithDefaults() *AwsCURConfigPostData {
	this := AwsCURConfigPostData{}
	var typeVar AwsCURConfigPostRequestType = AWSCURCONFIGPOSTREQUESTTYPE_AWS_CUR_CONFIG_POST_REQUEST
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AwsCURConfigPostData) GetAttributes() AwsCURConfigPostRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret AwsCURConfigPostRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPostData) GetAttributesOk() (*AwsCURConfigPostRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AwsCURConfigPostData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given AwsCURConfigPostRequestAttributes and assigns it to the Attributes field.
func (o *AwsCURConfigPostData) SetAttributes(v AwsCURConfigPostRequestAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *AwsCURConfigPostData) GetType() AwsCURConfigPostRequestType {
	if o == nil {
		var ret AwsCURConfigPostRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPostData) GetTypeOk() (*AwsCURConfigPostRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AwsCURConfigPostData) SetType(v AwsCURConfigPostRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsCURConfigPostData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsCURConfigPostData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AwsCURConfigPostRequestAttributes `json:"attributes,omitempty"`
		Type       *AwsCURConfigPostRequestType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
