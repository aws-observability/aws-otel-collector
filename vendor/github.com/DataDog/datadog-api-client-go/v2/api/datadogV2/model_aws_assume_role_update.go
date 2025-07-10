// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAssumeRoleUpdate The definition of `AWSAssumeRoleUpdate` object.
type AWSAssumeRoleUpdate struct {
	// AWS account the connection is created for
	AccountId *string `json:"account_id,omitempty"`
	// The `AWSAssumeRoleUpdate` `generate_new_external_id`.
	GenerateNewExternalId *bool `json:"generate_new_external_id,omitempty"`
	// Role to assume
	Role *string `json:"role,omitempty"`
	// The definition of `AWSAssumeRoleType` object.
	Type AWSAssumeRoleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSAssumeRoleUpdate instantiates a new AWSAssumeRoleUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAssumeRoleUpdate(typeVar AWSAssumeRoleType) *AWSAssumeRoleUpdate {
	this := AWSAssumeRoleUpdate{}
	this.Type = typeVar
	return &this
}

// NewAWSAssumeRoleUpdateWithDefaults instantiates a new AWSAssumeRoleUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAssumeRoleUpdateWithDefaults() *AWSAssumeRoleUpdate {
	this := AWSAssumeRoleUpdate{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AWSAssumeRoleUpdate) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAssumeRoleUpdate) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AWSAssumeRoleUpdate) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AWSAssumeRoleUpdate) SetAccountId(v string) {
	o.AccountId = &v
}

// GetGenerateNewExternalId returns the GenerateNewExternalId field value if set, zero value otherwise.
func (o *AWSAssumeRoleUpdate) GetGenerateNewExternalId() bool {
	if o == nil || o.GenerateNewExternalId == nil {
		var ret bool
		return ret
	}
	return *o.GenerateNewExternalId
}

// GetGenerateNewExternalIdOk returns a tuple with the GenerateNewExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAssumeRoleUpdate) GetGenerateNewExternalIdOk() (*bool, bool) {
	if o == nil || o.GenerateNewExternalId == nil {
		return nil, false
	}
	return o.GenerateNewExternalId, true
}

// HasGenerateNewExternalId returns a boolean if a field has been set.
func (o *AWSAssumeRoleUpdate) HasGenerateNewExternalId() bool {
	return o != nil && o.GenerateNewExternalId != nil
}

// SetGenerateNewExternalId gets a reference to the given bool and assigns it to the GenerateNewExternalId field.
func (o *AWSAssumeRoleUpdate) SetGenerateNewExternalId(v bool) {
	o.GenerateNewExternalId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AWSAssumeRoleUpdate) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAssumeRoleUpdate) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AWSAssumeRoleUpdate) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AWSAssumeRoleUpdate) SetRole(v string) {
	o.Role = &v
}

// GetType returns the Type field value.
func (o *AWSAssumeRoleUpdate) GetType() AWSAssumeRoleType {
	if o == nil {
		var ret AWSAssumeRoleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSAssumeRoleUpdate) GetTypeOk() (*AWSAssumeRoleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AWSAssumeRoleUpdate) SetType(v AWSAssumeRoleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAssumeRoleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.GenerateNewExternalId != nil {
		toSerialize["generate_new_external_id"] = o.GenerateNewExternalId
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSAssumeRoleUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId             *string            `json:"account_id,omitempty"`
		GenerateNewExternalId *bool              `json:"generate_new_external_id,omitempty"`
		Role                  *string            `json:"role,omitempty"`
		Type                  *AWSAssumeRoleType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "generate_new_external_id", "role", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = all.AccountId
	o.GenerateNewExternalId = all.GenerateNewExternalId
	o.Role = all.Role
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
