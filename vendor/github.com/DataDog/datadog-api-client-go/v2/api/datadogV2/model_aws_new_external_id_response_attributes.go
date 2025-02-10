// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSNewExternalIDResponseAttributes AWS External ID response body.
type AWSNewExternalIDResponseAttributes struct {
	// AWS IAM External ID for associated role.
	ExternalId string `json:"external_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSNewExternalIDResponseAttributes instantiates a new AWSNewExternalIDResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSNewExternalIDResponseAttributes(externalId string) *AWSNewExternalIDResponseAttributes {
	this := AWSNewExternalIDResponseAttributes{}
	this.ExternalId = externalId
	return &this
}

// NewAWSNewExternalIDResponseAttributesWithDefaults instantiates a new AWSNewExternalIDResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSNewExternalIDResponseAttributesWithDefaults() *AWSNewExternalIDResponseAttributes {
	this := AWSNewExternalIDResponseAttributes{}
	return &this
}

// GetExternalId returns the ExternalId field value.
func (o *AWSNewExternalIDResponseAttributes) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *AWSNewExternalIDResponseAttributes) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value.
func (o *AWSNewExternalIDResponseAttributes) SetExternalId(v string) {
	o.ExternalId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSNewExternalIDResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["external_id"] = o.ExternalId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSNewExternalIDResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExternalId *string `json:"external_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExternalId == nil {
		return fmt.Errorf("required field external_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"external_id"})
	} else {
		return err
	}
	o.ExternalId = *all.ExternalId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
