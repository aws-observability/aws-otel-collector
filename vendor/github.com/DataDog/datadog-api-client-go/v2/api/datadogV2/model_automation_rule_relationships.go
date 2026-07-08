// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleRelationships Related resources for the automation rule, including the users who created and last modified it.
type AutomationRuleRelationships struct {
	// Relationship to user.
	CreatedBy NullableNullableUserRelationship `json:"created_by,omitempty"`
	// Relationship to user.
	ModifiedBy NullableNullableUserRelationship `json:"modified_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutomationRuleRelationships instantiates a new AutomationRuleRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomationRuleRelationships() *AutomationRuleRelationships {
	this := AutomationRuleRelationships{}
	return &this
}

// NewAutomationRuleRelationshipsWithDefaults instantiates a new AutomationRuleRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutomationRuleRelationshipsWithDefaults() *AutomationRuleRelationships {
	this := AutomationRuleRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationRuleRelationships) GetCreatedBy() NullableUserRelationship {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret NullableUserRelationship
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AutomationRuleRelationships) GetCreatedByOk() (*NullableUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AutomationRuleRelationships) HasCreatedBy() bool {
	return o != nil && o.CreatedBy.IsSet()
}

// SetCreatedBy gets a reference to the given NullableNullableUserRelationship and assigns it to the CreatedBy field.
func (o *AutomationRuleRelationships) SetCreatedBy(v NullableUserRelationship) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil.
func (o *AutomationRuleRelationships) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil.
func (o *AutomationRuleRelationships) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationRuleRelationships) GetModifiedBy() NullableUserRelationship {
	if o == nil || o.ModifiedBy.Get() == nil {
		var ret NullableUserRelationship
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AutomationRuleRelationships) GetModifiedByOk() (*NullableUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *AutomationRuleRelationships) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy.IsSet()
}

// SetModifiedBy gets a reference to the given NullableNullableUserRelationship and assigns it to the ModifiedBy field.
func (o *AutomationRuleRelationships) SetModifiedBy(v NullableUserRelationship) {
	o.ModifiedBy.Set(&v)
}

// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil.
func (o *AutomationRuleRelationships) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil.
func (o *AutomationRuleRelationships) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o AutomationRuleRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedBy.IsSet() {
		toSerialize["created_by"] = o.CreatedBy.Get()
	}
	if o.ModifiedBy.IsSet() {
		toSerialize["modified_by"] = o.ModifiedBy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutomationRuleRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy  NullableNullableUserRelationship `json:"created_by,omitempty"`
		ModifiedBy NullableNullableUserRelationship `json:"modified_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "modified_by"})
	} else {
		return err
	}
	o.CreatedBy = all.CreatedBy
	o.ModifiedBy = all.ModifiedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
