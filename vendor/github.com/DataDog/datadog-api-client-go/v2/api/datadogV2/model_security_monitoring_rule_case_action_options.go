// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleCaseActionOptions Options for the rule action
type SecurityMonitoringRuleCaseActionOptions struct {
	// Duration of the action in seconds. 0 indicates no expiration.
	Duration *int64 `json:"duration,omitempty"`
	// Used with the case action of type 'flag_ip'. The value specified in this field is applied as a flag to the IP addresses.
	FlaggedIpType *SecurityMonitoringRuleCaseActionOptionsFlaggedIPType `json:"flaggedIPType,omitempty"`
	// Used with the case action of type 'user_behavior'. The value specified in this field is applied as a risk tag to all users affected by the rule.
	UserBehaviorName *string `json:"userBehaviorName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleCaseActionOptions instantiates a new SecurityMonitoringRuleCaseActionOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleCaseActionOptions() *SecurityMonitoringRuleCaseActionOptions {
	this := SecurityMonitoringRuleCaseActionOptions{}
	return &this
}

// NewSecurityMonitoringRuleCaseActionOptionsWithDefaults instantiates a new SecurityMonitoringRuleCaseActionOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleCaseActionOptionsWithDefaults() *SecurityMonitoringRuleCaseActionOptions {
	this := SecurityMonitoringRuleCaseActionOptions{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCaseActionOptions) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCaseActionOptions) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCaseActionOptions) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *SecurityMonitoringRuleCaseActionOptions) SetDuration(v int64) {
	o.Duration = &v
}

// GetFlaggedIpType returns the FlaggedIpType field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCaseActionOptions) GetFlaggedIpType() SecurityMonitoringRuleCaseActionOptionsFlaggedIPType {
	if o == nil || o.FlaggedIpType == nil {
		var ret SecurityMonitoringRuleCaseActionOptionsFlaggedIPType
		return ret
	}
	return *o.FlaggedIpType
}

// GetFlaggedIpTypeOk returns a tuple with the FlaggedIpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCaseActionOptions) GetFlaggedIpTypeOk() (*SecurityMonitoringRuleCaseActionOptionsFlaggedIPType, bool) {
	if o == nil || o.FlaggedIpType == nil {
		return nil, false
	}
	return o.FlaggedIpType, true
}

// HasFlaggedIpType returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCaseActionOptions) HasFlaggedIpType() bool {
	return o != nil && o.FlaggedIpType != nil
}

// SetFlaggedIpType gets a reference to the given SecurityMonitoringRuleCaseActionOptionsFlaggedIPType and assigns it to the FlaggedIpType field.
func (o *SecurityMonitoringRuleCaseActionOptions) SetFlaggedIpType(v SecurityMonitoringRuleCaseActionOptionsFlaggedIPType) {
	o.FlaggedIpType = &v
}

// GetUserBehaviorName returns the UserBehaviorName field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCaseActionOptions) GetUserBehaviorName() string {
	if o == nil || o.UserBehaviorName == nil {
		var ret string
		return ret
	}
	return *o.UserBehaviorName
}

// GetUserBehaviorNameOk returns a tuple with the UserBehaviorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCaseActionOptions) GetUserBehaviorNameOk() (*string, bool) {
	if o == nil || o.UserBehaviorName == nil {
		return nil, false
	}
	return o.UserBehaviorName, true
}

// HasUserBehaviorName returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCaseActionOptions) HasUserBehaviorName() bool {
	return o != nil && o.UserBehaviorName != nil
}

// SetUserBehaviorName gets a reference to the given string and assigns it to the UserBehaviorName field.
func (o *SecurityMonitoringRuleCaseActionOptions) SetUserBehaviorName(v string) {
	o.UserBehaviorName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleCaseActionOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.FlaggedIpType != nil {
		toSerialize["flaggedIPType"] = o.FlaggedIpType
	}
	if o.UserBehaviorName != nil {
		toSerialize["userBehaviorName"] = o.UserBehaviorName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleCaseActionOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration         *int64                                                `json:"duration,omitempty"`
		FlaggedIpType    *SecurityMonitoringRuleCaseActionOptionsFlaggedIPType `json:"flaggedIPType,omitempty"`
		UserBehaviorName *string                                               `json:"userBehaviorName,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration", "flaggedIPType", "userBehaviorName"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Duration = all.Duration
	if all.FlaggedIpType != nil && !all.FlaggedIpType.IsValid() {
		hasInvalidField = true
	} else {
		o.FlaggedIpType = all.FlaggedIpType
	}
	o.UserBehaviorName = all.UserBehaviorName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
