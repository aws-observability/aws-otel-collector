// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesFDDRule A faulty deployment detection rule to evaluate as part of a deployment gate evaluation.
type DeploymentGatesFDDRule struct {
	// Rule-level dry run. When enabled, the rule is evaluated normally but it always returns `pass`. The real result is visible in the Datadog UI.
	DryRun *bool `json:"dry_run,omitempty"`
	// Human-readable name for this rule.
	Name string `json:"name"`
	// Options for a `faulty_deployment_detection` rule.
	Options *DeploymentGatesFDDRuleOptions `json:"options,omitempty"`
	// The type identifier for a faulty deployment detection rule.
	Type DeploymentGatesFDDRuleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGatesFDDRule instantiates a new DeploymentGatesFDDRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGatesFDDRule(name string, typeVar DeploymentGatesFDDRuleType) *DeploymentGatesFDDRule {
	this := DeploymentGatesFDDRule{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewDeploymentGatesFDDRuleWithDefaults instantiates a new DeploymentGatesFDDRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGatesFDDRuleWithDefaults() *DeploymentGatesFDDRule {
	this := DeploymentGatesFDDRule{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeploymentGatesFDDRule) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesFDDRule) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeploymentGatesFDDRule) HasDryRun() bool {
	return o != nil && o.DryRun != nil
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeploymentGatesFDDRule) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetName returns the Name field value.
func (o *DeploymentGatesFDDRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesFDDRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DeploymentGatesFDDRule) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DeploymentGatesFDDRule) GetOptions() DeploymentGatesFDDRuleOptions {
	if o == nil || o.Options == nil {
		var ret DeploymentGatesFDDRuleOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesFDDRule) GetOptionsOk() (*DeploymentGatesFDDRuleOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DeploymentGatesFDDRule) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given DeploymentGatesFDDRuleOptions and assigns it to the Options field.
func (o *DeploymentGatesFDDRule) SetOptions(v DeploymentGatesFDDRuleOptions) {
	o.Options = &v
}

// GetType returns the Type field value.
func (o *DeploymentGatesFDDRule) GetType() DeploymentGatesFDDRuleType {
	if o == nil {
		var ret DeploymentGatesFDDRuleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesFDDRule) GetTypeOk() (*DeploymentGatesFDDRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DeploymentGatesFDDRule) SetType(v DeploymentGatesFDDRuleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGatesFDDRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	toSerialize["name"] = o.Name
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGatesFDDRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DryRun  *bool                          `json:"dry_run,omitempty"`
		Name    *string                        `json:"name"`
		Options *DeploymentGatesFDDRuleOptions `json:"options,omitempty"`
		Type    *DeploymentGatesFDDRuleType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dry_run", "name", "options", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DryRun = all.DryRun
	o.Name = *all.Name
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
