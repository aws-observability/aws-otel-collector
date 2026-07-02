// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleBasedViewComplianceFramework Compliance framework mapping for a rule.
type RuleBasedViewComplianceFramework struct {
	// Identifier of the control inside the requirement.
	Control *string `json:"control,omitempty"`
	// Handle of the compliance framework.
	Framework *string `json:"framework,omitempty"`
	// Whether the framework is a Datadog default framework. `true` indicates a Datadog framework and `false` indicates a custom framework.
	IsDefault *bool `json:"is_default,omitempty"`
	// Optional message describing the framework mapping for the rule.
	Message *string `json:"message,omitempty"`
	// Name of the requirement that contains the control.
	Requirement *string `json:"requirement,omitempty"`
	// Version of the compliance framework.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRuleBasedViewComplianceFramework instantiates a new RuleBasedViewComplianceFramework object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRuleBasedViewComplianceFramework() *RuleBasedViewComplianceFramework {
	this := RuleBasedViewComplianceFramework{}
	return &this
}

// NewRuleBasedViewComplianceFrameworkWithDefaults instantiates a new RuleBasedViewComplianceFramework object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRuleBasedViewComplianceFrameworkWithDefaults() *RuleBasedViewComplianceFramework {
	this := RuleBasedViewComplianceFramework{}
	return &this
}

// GetControl returns the Control field value if set, zero value otherwise.
func (o *RuleBasedViewComplianceFramework) GetControl() string {
	if o == nil || o.Control == nil {
		var ret string
		return ret
	}
	return *o.Control
}

// GetControlOk returns a tuple with the Control field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleBasedViewComplianceFramework) GetControlOk() (*string, bool) {
	if o == nil || o.Control == nil {
		return nil, false
	}
	return o.Control, true
}

// HasControl returns a boolean if a field has been set.
func (o *RuleBasedViewComplianceFramework) HasControl() bool {
	return o != nil && o.Control != nil
}

// SetControl gets a reference to the given string and assigns it to the Control field.
func (o *RuleBasedViewComplianceFramework) SetControl(v string) {
	o.Control = &v
}

// GetFramework returns the Framework field value if set, zero value otherwise.
func (o *RuleBasedViewComplianceFramework) GetFramework() string {
	if o == nil || o.Framework == nil {
		var ret string
		return ret
	}
	return *o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleBasedViewComplianceFramework) GetFrameworkOk() (*string, bool) {
	if o == nil || o.Framework == nil {
		return nil, false
	}
	return o.Framework, true
}

// HasFramework returns a boolean if a field has been set.
func (o *RuleBasedViewComplianceFramework) HasFramework() bool {
	return o != nil && o.Framework != nil
}

// SetFramework gets a reference to the given string and assigns it to the Framework field.
func (o *RuleBasedViewComplianceFramework) SetFramework(v string) {
	o.Framework = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *RuleBasedViewComplianceFramework) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleBasedViewComplianceFramework) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *RuleBasedViewComplianceFramework) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *RuleBasedViewComplianceFramework) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RuleBasedViewComplianceFramework) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleBasedViewComplianceFramework) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RuleBasedViewComplianceFramework) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RuleBasedViewComplianceFramework) SetMessage(v string) {
	o.Message = &v
}

// GetRequirement returns the Requirement field value if set, zero value otherwise.
func (o *RuleBasedViewComplianceFramework) GetRequirement() string {
	if o == nil || o.Requirement == nil {
		var ret string
		return ret
	}
	return *o.Requirement
}

// GetRequirementOk returns a tuple with the Requirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleBasedViewComplianceFramework) GetRequirementOk() (*string, bool) {
	if o == nil || o.Requirement == nil {
		return nil, false
	}
	return o.Requirement, true
}

// HasRequirement returns a boolean if a field has been set.
func (o *RuleBasedViewComplianceFramework) HasRequirement() bool {
	return o != nil && o.Requirement != nil
}

// SetRequirement gets a reference to the given string and assigns it to the Requirement field.
func (o *RuleBasedViewComplianceFramework) SetRequirement(v string) {
	o.Requirement = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RuleBasedViewComplianceFramework) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleBasedViewComplianceFramework) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RuleBasedViewComplianceFramework) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RuleBasedViewComplianceFramework) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RuleBasedViewComplianceFramework) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Control != nil {
		toSerialize["control"] = o.Control
	}
	if o.Framework != nil {
		toSerialize["framework"] = o.Framework
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Requirement != nil {
		toSerialize["requirement"] = o.Requirement
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RuleBasedViewComplianceFramework) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Control     *string `json:"control,omitempty"`
		Framework   *string `json:"framework,omitempty"`
		IsDefault   *bool   `json:"is_default,omitempty"`
		Message     *string `json:"message,omitempty"`
		Requirement *string `json:"requirement,omitempty"`
		Version     *string `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"control", "framework", "is_default", "message", "requirement", "version"})
	} else {
		return err
	}
	o.Control = all.Control
	o.Framework = all.Framework
	o.IsDefault = all.IsDefault
	o.Message = all.Message
	o.Requirement = all.Requirement
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
