// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCURConfigPatchRequestAttributes Attributes for AWS CUR config Patch Request.
type AwsCURConfigPatchRequestAttributes struct {
	// The account filtering configuration.
	AccountFilters *AccountFilteringConfig `json:"account_filters,omitempty"`
	// Whether or not the Cloud Cost Management account is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsCURConfigPatchRequestAttributes instantiates a new AwsCURConfigPatchRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsCURConfigPatchRequestAttributes() *AwsCURConfigPatchRequestAttributes {
	this := AwsCURConfigPatchRequestAttributes{}
	return &this
}

// NewAwsCURConfigPatchRequestAttributesWithDefaults instantiates a new AwsCURConfigPatchRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsCURConfigPatchRequestAttributesWithDefaults() *AwsCURConfigPatchRequestAttributes {
	this := AwsCURConfigPatchRequestAttributes{}
	return &this
}

// GetAccountFilters returns the AccountFilters field value if set, zero value otherwise.
func (o *AwsCURConfigPatchRequestAttributes) GetAccountFilters() AccountFilteringConfig {
	if o == nil || o.AccountFilters == nil {
		var ret AccountFilteringConfig
		return ret
	}
	return *o.AccountFilters
}

// GetAccountFiltersOk returns a tuple with the AccountFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPatchRequestAttributes) GetAccountFiltersOk() (*AccountFilteringConfig, bool) {
	if o == nil || o.AccountFilters == nil {
		return nil, false
	}
	return o.AccountFilters, true
}

// HasAccountFilters returns a boolean if a field has been set.
func (o *AwsCURConfigPatchRequestAttributes) HasAccountFilters() bool {
	return o != nil && o.AccountFilters != nil
}

// SetAccountFilters gets a reference to the given AccountFilteringConfig and assigns it to the AccountFilters field.
func (o *AwsCURConfigPatchRequestAttributes) SetAccountFilters(v AccountFilteringConfig) {
	o.AccountFilters = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *AwsCURConfigPatchRequestAttributes) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCURConfigPatchRequestAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *AwsCURConfigPatchRequestAttributes) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *AwsCURConfigPatchRequestAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsCURConfigPatchRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountFilters != nil {
		toSerialize["account_filters"] = o.AccountFilters
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsCURConfigPatchRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountFilters *AccountFilteringConfig `json:"account_filters,omitempty"`
		IsEnabled      *bool                   `json:"is_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_filters", "is_enabled"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AccountFilters != nil && all.AccountFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AccountFilters = all.AccountFilters
	o.IsEnabled = all.IsEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
