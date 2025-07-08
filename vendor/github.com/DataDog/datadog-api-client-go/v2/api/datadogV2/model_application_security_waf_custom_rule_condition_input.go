// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleConditionInput Input from the request on which the condition should apply.
type ApplicationSecurityWafCustomRuleConditionInput struct {
	// Input from the request on which the condition should apply.
	Address ApplicationSecurityWafCustomRuleConditionInputAddress `json:"address"`
	// Specific path for the input.
	KeyPath []string `json:"key_path,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafCustomRuleConditionInput instantiates a new ApplicationSecurityWafCustomRuleConditionInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafCustomRuleConditionInput(address ApplicationSecurityWafCustomRuleConditionInputAddress) *ApplicationSecurityWafCustomRuleConditionInput {
	this := ApplicationSecurityWafCustomRuleConditionInput{}
	this.Address = address
	return &this
}

// NewApplicationSecurityWafCustomRuleConditionInputWithDefaults instantiates a new ApplicationSecurityWafCustomRuleConditionInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafCustomRuleConditionInputWithDefaults() *ApplicationSecurityWafCustomRuleConditionInput {
	this := ApplicationSecurityWafCustomRuleConditionInput{}
	return &this
}

// GetAddress returns the Address field value.
func (o *ApplicationSecurityWafCustomRuleConditionInput) GetAddress() ApplicationSecurityWafCustomRuleConditionInputAddress {
	if o == nil {
		var ret ApplicationSecurityWafCustomRuleConditionInputAddress
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleConditionInput) GetAddressOk() (*ApplicationSecurityWafCustomRuleConditionInputAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value.
func (o *ApplicationSecurityWafCustomRuleConditionInput) SetAddress(v ApplicationSecurityWafCustomRuleConditionInputAddress) {
	o.Address = v
}

// GetKeyPath returns the KeyPath field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleConditionInput) GetKeyPath() []string {
	if o == nil || o.KeyPath == nil {
		var ret []string
		return ret
	}
	return o.KeyPath
}

// GetKeyPathOk returns a tuple with the KeyPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleConditionInput) GetKeyPathOk() (*[]string, bool) {
	if o == nil || o.KeyPath == nil {
		return nil, false
	}
	return &o.KeyPath, true
}

// HasKeyPath returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleConditionInput) HasKeyPath() bool {
	return o != nil && o.KeyPath != nil
}

// SetKeyPath gets a reference to the given []string and assigns it to the KeyPath field.
func (o *ApplicationSecurityWafCustomRuleConditionInput) SetKeyPath(v []string) {
	o.KeyPath = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafCustomRuleConditionInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["address"] = o.Address
	if o.KeyPath != nil {
		toSerialize["key_path"] = o.KeyPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafCustomRuleConditionInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Address *ApplicationSecurityWafCustomRuleConditionInputAddress `json:"address"`
		KeyPath []string                                               `json:"key_path,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Address == nil {
		return fmt.Errorf("required field address missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"address", "key_path"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Address.IsValid() {
		hasInvalidField = true
	} else {
		o.Address = *all.Address
	}
	o.KeyPath = all.KeyPath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
