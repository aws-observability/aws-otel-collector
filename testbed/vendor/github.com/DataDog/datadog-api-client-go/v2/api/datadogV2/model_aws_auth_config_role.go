// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAuthConfigRole AWS Authentication config to integrate your account using an IAM role.
type AWSAuthConfigRole struct {
	// AWS IAM External ID for associated role.
	ExternalId *string `json:"external_id,omitempty"`
	// AWS IAM Role name.
	RoleName string `json:"role_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSAuthConfigRole instantiates a new AWSAuthConfigRole object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAuthConfigRole(roleName string) *AWSAuthConfigRole {
	this := AWSAuthConfigRole{}
	this.RoleName = roleName
	return &this
}

// NewAWSAuthConfigRoleWithDefaults instantiates a new AWSAuthConfigRole object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAuthConfigRoleWithDefaults() *AWSAuthConfigRole {
	this := AWSAuthConfigRole{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AWSAuthConfigRole) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAuthConfigRole) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AWSAuthConfigRole) HasExternalId() bool {
	return o != nil && o.ExternalId != nil
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AWSAuthConfigRole) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetRoleName returns the RoleName field value.
func (o *AWSAuthConfigRole) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AWSAuthConfigRole) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value.
func (o *AWSAuthConfigRole) SetRoleName(v string) {
	o.RoleName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAuthConfigRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	toSerialize["role_name"] = o.RoleName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSAuthConfigRole) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExternalId *string `json:"external_id,omitempty"`
		RoleName   *string `json:"role_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RoleName == nil {
		return fmt.Errorf("required field role_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"external_id", "role_name"})
	} else {
		return err
	}
	o.ExternalId = all.ExternalId
	o.RoleName = *all.RoleName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
