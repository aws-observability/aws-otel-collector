// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCurConfigResponseDataAttributesAccountFilters The definition of `AwsCurConfigResponseDataAttributesAccountFilters` object.
type AwsCurConfigResponseDataAttributesAccountFilters struct {
	// The `account_filters` `excluded_accounts`.
	ExcludedAccounts []string `json:"excluded_accounts,omitempty"`
	// The `account_filters` `include_new_accounts`.
	IncludeNewAccounts datadog.NullableBool `json:"include_new_accounts,omitempty"`
	// The `account_filters` `included_accounts`.
	IncludedAccounts []string `json:"included_accounts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsCurConfigResponseDataAttributesAccountFilters instantiates a new AwsCurConfigResponseDataAttributesAccountFilters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsCurConfigResponseDataAttributesAccountFilters() *AwsCurConfigResponseDataAttributesAccountFilters {
	this := AwsCurConfigResponseDataAttributesAccountFilters{}
	return &this
}

// NewAwsCurConfigResponseDataAttributesAccountFiltersWithDefaults instantiates a new AwsCurConfigResponseDataAttributesAccountFilters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsCurConfigResponseDataAttributesAccountFiltersWithDefaults() *AwsCurConfigResponseDataAttributesAccountFilters {
	this := AwsCurConfigResponseDataAttributesAccountFilters{}
	return &this
}

// GetExcludedAccounts returns the ExcludedAccounts field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) GetExcludedAccounts() []string {
	if o == nil || o.ExcludedAccounts == nil {
		var ret []string
		return ret
	}
	return o.ExcludedAccounts
}

// GetExcludedAccountsOk returns a tuple with the ExcludedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) GetExcludedAccountsOk() (*[]string, bool) {
	if o == nil || o.ExcludedAccounts == nil {
		return nil, false
	}
	return &o.ExcludedAccounts, true
}

// HasExcludedAccounts returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) HasExcludedAccounts() bool {
	return o != nil && o.ExcludedAccounts != nil
}

// SetExcludedAccounts gets a reference to the given []string and assigns it to the ExcludedAccounts field.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) SetExcludedAccounts(v []string) {
	o.ExcludedAccounts = v
}

// GetIncludeNewAccounts returns the IncludeNewAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsCurConfigResponseDataAttributesAccountFilters) GetIncludeNewAccounts() bool {
	if o == nil || o.IncludeNewAccounts.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludeNewAccounts.Get()
}

// GetIncludeNewAccountsOk returns a tuple with the IncludeNewAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) GetIncludeNewAccountsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeNewAccounts.Get(), o.IncludeNewAccounts.IsSet()
}

// HasIncludeNewAccounts returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) HasIncludeNewAccounts() bool {
	return o != nil && o.IncludeNewAccounts.IsSet()
}

// SetIncludeNewAccounts gets a reference to the given datadog.NullableBool and assigns it to the IncludeNewAccounts field.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) SetIncludeNewAccounts(v bool) {
	o.IncludeNewAccounts.Set(&v)
}

// SetIncludeNewAccountsNil sets the value for IncludeNewAccounts to be an explicit nil.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) SetIncludeNewAccountsNil() {
	o.IncludeNewAccounts.Set(nil)
}

// UnsetIncludeNewAccounts ensures that no value is present for IncludeNewAccounts, not even an explicit nil.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) UnsetIncludeNewAccounts() {
	o.IncludeNewAccounts.Unset()
}

// GetIncludedAccounts returns the IncludedAccounts field value if set, zero value otherwise.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) GetIncludedAccounts() []string {
	if o == nil || o.IncludedAccounts == nil {
		var ret []string
		return ret
	}
	return o.IncludedAccounts
}

// GetIncludedAccountsOk returns a tuple with the IncludedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) GetIncludedAccountsOk() (*[]string, bool) {
	if o == nil || o.IncludedAccounts == nil {
		return nil, false
	}
	return &o.IncludedAccounts, true
}

// HasIncludedAccounts returns a boolean if a field has been set.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) HasIncludedAccounts() bool {
	return o != nil && o.IncludedAccounts != nil
}

// SetIncludedAccounts gets a reference to the given []string and assigns it to the IncludedAccounts field.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) SetIncludedAccounts(v []string) {
	o.IncludedAccounts = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsCurConfigResponseDataAttributesAccountFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExcludedAccounts != nil {
		toSerialize["excluded_accounts"] = o.ExcludedAccounts
	}
	if o.IncludeNewAccounts.IsSet() {
		toSerialize["include_new_accounts"] = o.IncludeNewAccounts.Get()
	}
	if o.IncludedAccounts != nil {
		toSerialize["included_accounts"] = o.IncludedAccounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsCurConfigResponseDataAttributesAccountFilters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExcludedAccounts   []string             `json:"excluded_accounts,omitempty"`
		IncludeNewAccounts datadog.NullableBool `json:"include_new_accounts,omitempty"`
		IncludedAccounts   []string             `json:"included_accounts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"excluded_accounts", "include_new_accounts", "included_accounts"})
	} else {
		return err
	}
	o.ExcludedAccounts = all.ExcludedAccounts
	o.IncludeNewAccounts = all.IncludeNewAccounts
	o.IncludedAccounts = all.IncludedAccounts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
