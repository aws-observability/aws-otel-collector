// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSIntegrationIamPermissionsResponseAttributes AWS Integration IAM Permissions response attributes.
type AWSIntegrationIamPermissionsResponseAttributes struct {
	// List of AWS IAM permissions required for the integration.
	Permissions []string `json:"permissions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSIntegrationIamPermissionsResponseAttributes instantiates a new AWSIntegrationIamPermissionsResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSIntegrationIamPermissionsResponseAttributes(permissions []string) *AWSIntegrationIamPermissionsResponseAttributes {
	this := AWSIntegrationIamPermissionsResponseAttributes{}
	this.Permissions = permissions
	return &this
}

// NewAWSIntegrationIamPermissionsResponseAttributesWithDefaults instantiates a new AWSIntegrationIamPermissionsResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSIntegrationIamPermissionsResponseAttributesWithDefaults() *AWSIntegrationIamPermissionsResponseAttributes {
	this := AWSIntegrationIamPermissionsResponseAttributes{}
	return &this
}

// GetPermissions returns the Permissions field value.
func (o *AWSIntegrationIamPermissionsResponseAttributes) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *AWSIntegrationIamPermissionsResponseAttributes) GetPermissionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value.
func (o *AWSIntegrationIamPermissionsResponseAttributes) SetPermissions(v []string) {
	o.Permissions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSIntegrationIamPermissionsResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["permissions"] = o.Permissions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSIntegrationIamPermissionsResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Permissions *[]string `json:"permissions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Permissions == nil {
		return fmt.Errorf("required field permissions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"permissions"})
	} else {
		return err
	}
	o.Permissions = *all.Permissions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
