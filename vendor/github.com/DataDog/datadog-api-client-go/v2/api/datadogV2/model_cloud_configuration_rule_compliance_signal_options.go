// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudConfigurationRuleComplianceSignalOptions How to generate compliance signals. Useful for cloud_configuration rules only.
type CloudConfigurationRuleComplianceSignalOptions struct {
	// Whether signals will be sent.
	UserActivationStatus datadog.NullableBool `json:"userActivationStatus,omitempty"`
	// Fields to use to group findings by when sending signals.
	UserGroupByFields datadog.NullableList[string] `json:"userGroupByFields,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCloudConfigurationRuleComplianceSignalOptions instantiates a new CloudConfigurationRuleComplianceSignalOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudConfigurationRuleComplianceSignalOptions() *CloudConfigurationRuleComplianceSignalOptions {
	this := CloudConfigurationRuleComplianceSignalOptions{}
	return &this
}

// NewCloudConfigurationRuleComplianceSignalOptionsWithDefaults instantiates a new CloudConfigurationRuleComplianceSignalOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudConfigurationRuleComplianceSignalOptionsWithDefaults() *CloudConfigurationRuleComplianceSignalOptions {
	this := CloudConfigurationRuleComplianceSignalOptions{}
	return &this
}

// GetUserActivationStatus returns the UserActivationStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudConfigurationRuleComplianceSignalOptions) GetUserActivationStatus() bool {
	if o == nil || o.UserActivationStatus.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UserActivationStatus.Get()
}

// GetUserActivationStatusOk returns a tuple with the UserActivationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CloudConfigurationRuleComplianceSignalOptions) GetUserActivationStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserActivationStatus.Get(), o.UserActivationStatus.IsSet()
}

// HasUserActivationStatus returns a boolean if a field has been set.
func (o *CloudConfigurationRuleComplianceSignalOptions) HasUserActivationStatus() bool {
	return o != nil && o.UserActivationStatus.IsSet()
}

// SetUserActivationStatus gets a reference to the given datadog.NullableBool and assigns it to the UserActivationStatus field.
func (o *CloudConfigurationRuleComplianceSignalOptions) SetUserActivationStatus(v bool) {
	o.UserActivationStatus.Set(&v)
}

// SetUserActivationStatusNil sets the value for UserActivationStatus to be an explicit nil.
func (o *CloudConfigurationRuleComplianceSignalOptions) SetUserActivationStatusNil() {
	o.UserActivationStatus.Set(nil)
}

// UnsetUserActivationStatus ensures that no value is present for UserActivationStatus, not even an explicit nil.
func (o *CloudConfigurationRuleComplianceSignalOptions) UnsetUserActivationStatus() {
	o.UserActivationStatus.Unset()
}

// GetUserGroupByFields returns the UserGroupByFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudConfigurationRuleComplianceSignalOptions) GetUserGroupByFields() []string {
	if o == nil || o.UserGroupByFields.Get() == nil {
		var ret []string
		return ret
	}
	return *o.UserGroupByFields.Get()
}

// GetUserGroupByFieldsOk returns a tuple with the UserGroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CloudConfigurationRuleComplianceSignalOptions) GetUserGroupByFieldsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserGroupByFields.Get(), o.UserGroupByFields.IsSet()
}

// HasUserGroupByFields returns a boolean if a field has been set.
func (o *CloudConfigurationRuleComplianceSignalOptions) HasUserGroupByFields() bool {
	return o != nil && o.UserGroupByFields.IsSet()
}

// SetUserGroupByFields gets a reference to the given datadog.NullableList[string] and assigns it to the UserGroupByFields field.
func (o *CloudConfigurationRuleComplianceSignalOptions) SetUserGroupByFields(v []string) {
	o.UserGroupByFields.Set(&v)
}

// SetUserGroupByFieldsNil sets the value for UserGroupByFields to be an explicit nil.
func (o *CloudConfigurationRuleComplianceSignalOptions) SetUserGroupByFieldsNil() {
	o.UserGroupByFields.Set(nil)
}

// UnsetUserGroupByFields ensures that no value is present for UserGroupByFields, not even an explicit nil.
func (o *CloudConfigurationRuleComplianceSignalOptions) UnsetUserGroupByFields() {
	o.UserGroupByFields.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudConfigurationRuleComplianceSignalOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.UserActivationStatus.IsSet() {
		toSerialize["userActivationStatus"] = o.UserActivationStatus.Get()
	}
	if o.UserGroupByFields.IsSet() {
		toSerialize["userGroupByFields"] = o.UserGroupByFields.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudConfigurationRuleComplianceSignalOptions) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		UserActivationStatus datadog.NullableBool         `json:"userActivationStatus,omitempty"`
		UserGroupByFields    datadog.NullableList[string] `json:"userGroupByFields,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"userActivationStatus", "userGroupByFields"})
	} else {
		return err
	}
	o.UserActivationStatus = all.UserActivationStatus
	o.UserGroupByFields = all.UserGroupByFields
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
