// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamPermissionSettingUpdateAttributes Team permission setting update attributes
type TeamPermissionSettingUpdateAttributes struct {
	// What type of user is allowed to perform the specified action
	Value *TeamPermissionSettingValue `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamPermissionSettingUpdateAttributes instantiates a new TeamPermissionSettingUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamPermissionSettingUpdateAttributes() *TeamPermissionSettingUpdateAttributes {
	this := TeamPermissionSettingUpdateAttributes{}
	return &this
}

// NewTeamPermissionSettingUpdateAttributesWithDefaults instantiates a new TeamPermissionSettingUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamPermissionSettingUpdateAttributesWithDefaults() *TeamPermissionSettingUpdateAttributes {
	this := TeamPermissionSettingUpdateAttributes{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TeamPermissionSettingUpdateAttributes) GetValue() TeamPermissionSettingValue {
	if o == nil || o.Value == nil {
		var ret TeamPermissionSettingValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPermissionSettingUpdateAttributes) GetValueOk() (*TeamPermissionSettingValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TeamPermissionSettingUpdateAttributes) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given TeamPermissionSettingValue and assigns it to the Value field.
func (o *TeamPermissionSettingUpdateAttributes) SetValue(v TeamPermissionSettingValue) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamPermissionSettingUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamPermissionSettingUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value *TeamPermissionSettingValue `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Value != nil && !all.Value.IsValid() {
		hasInvalidField = true
	} else {
		o.Value = all.Value
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
