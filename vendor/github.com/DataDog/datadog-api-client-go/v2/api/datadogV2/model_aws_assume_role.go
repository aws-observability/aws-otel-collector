// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAssumeRole The definition of `AWSAssumeRole` object.
type AWSAssumeRole struct {
	// AWS account the connection is created for
	AccountId string `json:"account_id"`
	// External ID used to scope which connection can be used to assume the role
	ExternalId *string `json:"external_id,omitempty"`
	// AWS account that will assume the role
	PrincipalId *string `json:"principal_id,omitempty"`
	// Role to assume
	Role string `json:"role"`
	// The definition of `AWSAssumeRoleType` object.
	Type AWSAssumeRoleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSAssumeRole instantiates a new AWSAssumeRole object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAssumeRole(accountId string, role string, typeVar AWSAssumeRoleType) *AWSAssumeRole {
	this := AWSAssumeRole{}
	this.AccountId = accountId
	this.Role = role
	this.Type = typeVar
	return &this
}

// NewAWSAssumeRoleWithDefaults instantiates a new AWSAssumeRole object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAssumeRoleWithDefaults() *AWSAssumeRole {
	this := AWSAssumeRole{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *AWSAssumeRole) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AWSAssumeRole) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *AWSAssumeRole) SetAccountId(v string) {
	o.AccountId = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AWSAssumeRole) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAssumeRole) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AWSAssumeRole) HasExternalId() bool {
	return o != nil && o.ExternalId != nil
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AWSAssumeRole) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetPrincipalId returns the PrincipalId field value if set, zero value otherwise.
func (o *AWSAssumeRole) GetPrincipalId() string {
	if o == nil || o.PrincipalId == nil {
		var ret string
		return ret
	}
	return *o.PrincipalId
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAssumeRole) GetPrincipalIdOk() (*string, bool) {
	if o == nil || o.PrincipalId == nil {
		return nil, false
	}
	return o.PrincipalId, true
}

// HasPrincipalId returns a boolean if a field has been set.
func (o *AWSAssumeRole) HasPrincipalId() bool {
	return o != nil && o.PrincipalId != nil
}

// SetPrincipalId gets a reference to the given string and assigns it to the PrincipalId field.
func (o *AWSAssumeRole) SetPrincipalId(v string) {
	o.PrincipalId = &v
}

// GetRole returns the Role field value.
func (o *AWSAssumeRole) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AWSAssumeRole) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *AWSAssumeRole) SetRole(v string) {
	o.Role = v
}

// GetType returns the Type field value.
func (o *AWSAssumeRole) GetType() AWSAssumeRoleType {
	if o == nil {
		var ret AWSAssumeRoleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSAssumeRole) GetTypeOk() (*AWSAssumeRoleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AWSAssumeRole) SetType(v AWSAssumeRoleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAssumeRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.PrincipalId != nil {
		toSerialize["principal_id"] = o.PrincipalId
	}
	toSerialize["role"] = o.Role
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSAssumeRole) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId   *string            `json:"account_id"`
		ExternalId  *string            `json:"external_id,omitempty"`
		PrincipalId *string            `json:"principal_id,omitempty"`
		Role        *string            `json:"role"`
		Type        *AWSAssumeRoleType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "external_id", "principal_id", "role", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = *all.AccountId
	o.ExternalId = all.ExternalId
	o.PrincipalId = all.PrincipalId
	o.Role = *all.Role
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
