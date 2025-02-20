// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAccountUpdateRequestData AWS Account Update Request data.
type AWSAccountUpdateRequestData struct {
	// The AWS Account Integration Config to be updated.
	Attributes AWSAccountUpdateRequestAttributes `json:"attributes"`
	// Unique Datadog ID of the AWS Account Integration Config.
	// To get the config ID for an account, use the [List all AWS integrations](https://docs.datadoghq.com/api/latest/aws-integration/#list-all-aws-integrations)
	// endpoint and query by AWS Account ID.
	Id *string `json:"id,omitempty"`
	// AWS Account resource type.
	Type AWSAccountType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSAccountUpdateRequestData instantiates a new AWSAccountUpdateRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAccountUpdateRequestData(attributes AWSAccountUpdateRequestAttributes, typeVar AWSAccountType) *AWSAccountUpdateRequestData {
	this := AWSAccountUpdateRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewAWSAccountUpdateRequestDataWithDefaults instantiates a new AWSAccountUpdateRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAccountUpdateRequestDataWithDefaults() *AWSAccountUpdateRequestData {
	this := AWSAccountUpdateRequestData{}
	var typeVar AWSAccountType = AWSACCOUNTTYPE_ACCOUNT
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AWSAccountUpdateRequestData) GetAttributes() AWSAccountUpdateRequestAttributes {
	if o == nil {
		var ret AWSAccountUpdateRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestData) GetAttributesOk() (*AWSAccountUpdateRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AWSAccountUpdateRequestData) SetAttributes(v AWSAccountUpdateRequestAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AWSAccountUpdateRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AWSAccountUpdateRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AWSAccountUpdateRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *AWSAccountUpdateRequestData) GetType() AWSAccountType {
	if o == nil {
		var ret AWSAccountType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestData) GetTypeOk() (*AWSAccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AWSAccountUpdateRequestData) SetType(v AWSAccountType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAccountUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
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
func (o *AWSAccountUpdateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AWSAccountUpdateRequestAttributes `json:"attributes"`
		Id         *string                            `json:"id,omitempty"`
		Type       *AWSAccountType                    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
