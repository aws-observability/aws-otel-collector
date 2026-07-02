// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityPolicyRulesetOverride Deprecated: Override WAF ruleset parameters. Use `protectionPresets` instead.
//
// Deprecated: This model is deprecated.
type ApplicationSecurityPolicyRulesetOverride struct {
	// When blocking is enabled, the ruleset will block the traffic it matches.
	Blocking bool `json:"blocking"`
	// When false, this ruleset will not match any traffic.
	Enabled bool `json:"enabled"`
	// The identifier of the ruleset to override.
	Id string `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityPolicyRulesetOverride instantiates a new ApplicationSecurityPolicyRulesetOverride object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityPolicyRulesetOverride(blocking bool, enabled bool, id string) *ApplicationSecurityPolicyRulesetOverride {
	this := ApplicationSecurityPolicyRulesetOverride{}
	this.Blocking = blocking
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewApplicationSecurityPolicyRulesetOverrideWithDefaults instantiates a new ApplicationSecurityPolicyRulesetOverride object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityPolicyRulesetOverrideWithDefaults() *ApplicationSecurityPolicyRulesetOverride {
	this := ApplicationSecurityPolicyRulesetOverride{}
	return &this
}

// GetBlocking returns the Blocking field value.
func (o *ApplicationSecurityPolicyRulesetOverride) GetBlocking() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Blocking
}

// GetBlockingOk returns a tuple with the Blocking field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyRulesetOverride) GetBlockingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blocking, true
}

// SetBlocking sets field value.
func (o *ApplicationSecurityPolicyRulesetOverride) SetBlocking(v bool) {
	o.Blocking = v
}

// GetEnabled returns the Enabled field value.
func (o *ApplicationSecurityPolicyRulesetOverride) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyRulesetOverride) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ApplicationSecurityPolicyRulesetOverride) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ApplicationSecurityPolicyRulesetOverride) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyRulesetOverride) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ApplicationSecurityPolicyRulesetOverride) SetId(v string) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityPolicyRulesetOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["blocking"] = o.Blocking
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityPolicyRulesetOverride) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Blocking *bool   `json:"blocking"`
		Enabled  *bool   `json:"enabled"`
		Id       *string `json:"id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Blocking == nil {
		return fmt.Errorf("required field blocking missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"blocking", "enabled", "id"})
	} else {
		return err
	}
	o.Blocking = *all.Blocking
	o.Enabled = *all.Enabled
	o.Id = *all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
