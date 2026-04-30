// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatCreateOrganizationHandleRequestData Organization handle data for a create request.
type GoogleChatCreateOrganizationHandleRequestData struct {
	// Organization handle attributes for a create request.
	Attributes GoogleChatCreateOrganizationHandleRequestAttributes `json:"attributes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatCreateOrganizationHandleRequestData instantiates a new GoogleChatCreateOrganizationHandleRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatCreateOrganizationHandleRequestData(attributes GoogleChatCreateOrganizationHandleRequestAttributes) *GoogleChatCreateOrganizationHandleRequestData {
	this := GoogleChatCreateOrganizationHandleRequestData{}
	this.Attributes = attributes
	return &this
}

// NewGoogleChatCreateOrganizationHandleRequestDataWithDefaults instantiates a new GoogleChatCreateOrganizationHandleRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatCreateOrganizationHandleRequestDataWithDefaults() *GoogleChatCreateOrganizationHandleRequestData {
	this := GoogleChatCreateOrganizationHandleRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *GoogleChatCreateOrganizationHandleRequestData) GetAttributes() GoogleChatCreateOrganizationHandleRequestAttributes {
	if o == nil {
		var ret GoogleChatCreateOrganizationHandleRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GoogleChatCreateOrganizationHandleRequestData) GetAttributesOk() (*GoogleChatCreateOrganizationHandleRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *GoogleChatCreateOrganizationHandleRequestData) SetAttributes(v GoogleChatCreateOrganizationHandleRequestAttributes) {
	o.Attributes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatCreateOrganizationHandleRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatCreateOrganizationHandleRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *GoogleChatCreateOrganizationHandleRequestAttributes `json:"attributes"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
