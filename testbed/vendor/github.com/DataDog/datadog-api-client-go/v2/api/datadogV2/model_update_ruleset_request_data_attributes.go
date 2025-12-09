// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateRulesetRequestDataAttributes The definition of `UpdateRulesetRequestDataAttributes` object.
type UpdateRulesetRequestDataAttributes struct {
	// The `attributes` `enabled`.
	Enabled bool `json:"enabled"`
	// The `attributes` `last_version`.
	LastVersion *int64 `json:"last_version,omitempty"`
	// The `attributes` `rules`.
	Rules []UpdateRulesetRequestDataAttributesRulesItems `json:"rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateRulesetRequestDataAttributes instantiates a new UpdateRulesetRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateRulesetRequestDataAttributes(enabled bool, rules []UpdateRulesetRequestDataAttributesRulesItems) *UpdateRulesetRequestDataAttributes {
	this := UpdateRulesetRequestDataAttributes{}
	this.Enabled = enabled
	this.Rules = rules
	return &this
}

// NewUpdateRulesetRequestDataAttributesWithDefaults instantiates a new UpdateRulesetRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateRulesetRequestDataAttributesWithDefaults() *UpdateRulesetRequestDataAttributes {
	this := UpdateRulesetRequestDataAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *UpdateRulesetRequestDataAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *UpdateRulesetRequestDataAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLastVersion returns the LastVersion field value if set, zero value otherwise.
func (o *UpdateRulesetRequestDataAttributes) GetLastVersion() int64 {
	if o == nil || o.LastVersion == nil {
		var ret int64
		return ret
	}
	return *o.LastVersion
}

// GetLastVersionOk returns a tuple with the LastVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributes) GetLastVersionOk() (*int64, bool) {
	if o == nil || o.LastVersion == nil {
		return nil, false
	}
	return o.LastVersion, true
}

// HasLastVersion returns a boolean if a field has been set.
func (o *UpdateRulesetRequestDataAttributes) HasLastVersion() bool {
	return o != nil && o.LastVersion != nil
}

// SetLastVersion gets a reference to the given int64 and assigns it to the LastVersion field.
func (o *UpdateRulesetRequestDataAttributes) SetLastVersion(v int64) {
	o.LastVersion = &v
}

// GetRules returns the Rules field value.
func (o *UpdateRulesetRequestDataAttributes) GetRules() []UpdateRulesetRequestDataAttributesRulesItems {
	if o == nil {
		var ret []UpdateRulesetRequestDataAttributesRulesItems
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributes) GetRulesOk() (*[]UpdateRulesetRequestDataAttributesRulesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *UpdateRulesetRequestDataAttributes) SetRules(v []UpdateRulesetRequestDataAttributesRulesItems) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateRulesetRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.LastVersion != nil {
		toSerialize["last_version"] = o.LastVersion
	}
	toSerialize["rules"] = o.Rules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateRulesetRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled     *bool                                           `json:"enabled"`
		LastVersion *int64                                          `json:"last_version,omitempty"`
		Rules       *[]UpdateRulesetRequestDataAttributesRulesItems `json:"rules"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "last_version", "rules"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.LastVersion = all.LastVersion
	o.Rules = *all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
