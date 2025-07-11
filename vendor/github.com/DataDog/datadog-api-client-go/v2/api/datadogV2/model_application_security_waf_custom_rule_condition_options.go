// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleConditionOptions Options for the operator of this condition.
type ApplicationSecurityWafCustomRuleConditionOptions struct {
	// Evaluate the value as case sensitive.
	CaseSensitive *bool `json:"case_sensitive,omitempty"`
	// Only evaluate this condition if the value has a minimum amount of characters.
	MinLength *int64 `json:"min_length,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafCustomRuleConditionOptions instantiates a new ApplicationSecurityWafCustomRuleConditionOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafCustomRuleConditionOptions() *ApplicationSecurityWafCustomRuleConditionOptions {
	this := ApplicationSecurityWafCustomRuleConditionOptions{}
	var caseSensitive bool = false
	this.CaseSensitive = &caseSensitive
	var minLength int64 = 0
	this.MinLength = &minLength
	return &this
}

// NewApplicationSecurityWafCustomRuleConditionOptionsWithDefaults instantiates a new ApplicationSecurityWafCustomRuleConditionOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafCustomRuleConditionOptionsWithDefaults() *ApplicationSecurityWafCustomRuleConditionOptions {
	this := ApplicationSecurityWafCustomRuleConditionOptions{}
	var caseSensitive bool = false
	this.CaseSensitive = &caseSensitive
	var minLength int64 = 0
	this.MinLength = &minLength
	return &this
}

// GetCaseSensitive returns the CaseSensitive field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleConditionOptions) GetCaseSensitive() bool {
	if o == nil || o.CaseSensitive == nil {
		var ret bool
		return ret
	}
	return *o.CaseSensitive
}

// GetCaseSensitiveOk returns a tuple with the CaseSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleConditionOptions) GetCaseSensitiveOk() (*bool, bool) {
	if o == nil || o.CaseSensitive == nil {
		return nil, false
	}
	return o.CaseSensitive, true
}

// HasCaseSensitive returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleConditionOptions) HasCaseSensitive() bool {
	return o != nil && o.CaseSensitive != nil
}

// SetCaseSensitive gets a reference to the given bool and assigns it to the CaseSensitive field.
func (o *ApplicationSecurityWafCustomRuleConditionOptions) SetCaseSensitive(v bool) {
	o.CaseSensitive = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleConditionOptions) GetMinLength() int64 {
	if o == nil || o.MinLength == nil {
		var ret int64
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleConditionOptions) GetMinLengthOk() (*int64, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleConditionOptions) HasMinLength() bool {
	return o != nil && o.MinLength != nil
}

// SetMinLength gets a reference to the given int64 and assigns it to the MinLength field.
func (o *ApplicationSecurityWafCustomRuleConditionOptions) SetMinLength(v int64) {
	o.MinLength = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafCustomRuleConditionOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CaseSensitive != nil {
		toSerialize["case_sensitive"] = o.CaseSensitive
	}
	if o.MinLength != nil {
		toSerialize["min_length"] = o.MinLength
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafCustomRuleConditionOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaseSensitive *bool  `json:"case_sensitive,omitempty"`
		MinLength     *int64 `json:"min_length,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"case_sensitive", "min_length"})
	} else {
		return err
	}
	o.CaseSensitive = all.CaseSensitive
	o.MinLength = all.MinLength

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
