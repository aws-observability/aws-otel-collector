// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestBindingItems Object describing the binding used for a mobile test.
type SyntheticsMobileTestBindingItems struct {
	// List of principals for a mobile test binding.
	Principals []string `json:"principals,omitempty"`
	// The definition of `SyntheticsMobileTestBindingItemsRole` object.
	Role *SyntheticsMobileTestBindingItemsRole `json:"role,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileTestBindingItems instantiates a new SyntheticsMobileTestBindingItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileTestBindingItems() *SyntheticsMobileTestBindingItems {
	this := SyntheticsMobileTestBindingItems{}
	return &this
}

// NewSyntheticsMobileTestBindingItemsWithDefaults instantiates a new SyntheticsMobileTestBindingItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileTestBindingItemsWithDefaults() *SyntheticsMobileTestBindingItems {
	this := SyntheticsMobileTestBindingItems{}
	return &this
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *SyntheticsMobileTestBindingItems) GetPrincipals() []string {
	if o == nil || o.Principals == nil {
		var ret []string
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestBindingItems) GetPrincipalsOk() (*[]string, bool) {
	if o == nil || o.Principals == nil {
		return nil, false
	}
	return &o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *SyntheticsMobileTestBindingItems) HasPrincipals() bool {
	return o != nil && o.Principals != nil
}

// SetPrincipals gets a reference to the given []string and assigns it to the Principals field.
func (o *SyntheticsMobileTestBindingItems) SetPrincipals(v []string) {
	o.Principals = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SyntheticsMobileTestBindingItems) GetRole() SyntheticsMobileTestBindingItemsRole {
	if o == nil || o.Role == nil {
		var ret SyntheticsMobileTestBindingItemsRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestBindingItems) GetRoleOk() (*SyntheticsMobileTestBindingItemsRole, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SyntheticsMobileTestBindingItems) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given SyntheticsMobileTestBindingItemsRole and assigns it to the Role field.
func (o *SyntheticsMobileTestBindingItems) SetRole(v SyntheticsMobileTestBindingItemsRole) {
	o.Role = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileTestBindingItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Principals != nil {
		toSerialize["principals"] = o.Principals
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileTestBindingItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Principals []string                              `json:"principals,omitempty"`
		Role       *SyntheticsMobileTestBindingItemsRole `json:"role,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"principals", "role"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Principals = all.Principals
	if all.Role != nil && !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = all.Role
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
