// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSNamespacesResponseAttributes AWS Namespaces response attributes.
type AWSNamespacesResponseAttributes struct {
	// AWS CloudWatch namespace.
	Namespaces []string `json:"namespaces"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSNamespacesResponseAttributes instantiates a new AWSNamespacesResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSNamespacesResponseAttributes(namespaces []string) *AWSNamespacesResponseAttributes {
	this := AWSNamespacesResponseAttributes{}
	this.Namespaces = namespaces
	return &this
}

// NewAWSNamespacesResponseAttributesWithDefaults instantiates a new AWSNamespacesResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSNamespacesResponseAttributesWithDefaults() *AWSNamespacesResponseAttributes {
	this := AWSNamespacesResponseAttributes{}
	return &this
}

// GetNamespaces returns the Namespaces field value.
func (o *AWSNamespacesResponseAttributes) GetNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value
// and a boolean to check if the value has been set.
func (o *AWSNamespacesResponseAttributes) GetNamespacesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespaces, true
}

// SetNamespaces sets field value.
func (o *AWSNamespacesResponseAttributes) SetNamespaces(v []string) {
	o.Namespaces = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSNamespacesResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["namespaces"] = o.Namespaces

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSNamespacesResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Namespaces *[]string `json:"namespaces"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Namespaces == nil {
		return fmt.Errorf("required field namespaces missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"namespaces"})
	} else {
		return err
	}
	o.Namespaces = *all.Namespaces

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
