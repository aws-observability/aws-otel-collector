// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesEvaluationConfiguration Inline rule definitions for a deployment gate evaluation. When provided, rules are evaluated
// directly from this configuration instead of using the preconfigured gate rules.
// At least one rule is required.
type DeploymentGatesEvaluationConfiguration struct {
	// Gate-level dry run. When enabled, the rules are evaluated normally but the gate always returns `pass`. The real result is visible in the Datadog UI.
	DryRun *bool `json:"dry_run,omitempty"`
	// The list of rules to evaluate. At least one rule is required.
	Rules []DeploymentGatesEvaluationRule `json:"rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGatesEvaluationConfiguration instantiates a new DeploymentGatesEvaluationConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGatesEvaluationConfiguration(rules []DeploymentGatesEvaluationRule) *DeploymentGatesEvaluationConfiguration {
	this := DeploymentGatesEvaluationConfiguration{}
	this.Rules = rules
	return &this
}

// NewDeploymentGatesEvaluationConfigurationWithDefaults instantiates a new DeploymentGatesEvaluationConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGatesEvaluationConfigurationWithDefaults() *DeploymentGatesEvaluationConfiguration {
	this := DeploymentGatesEvaluationConfiguration{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeploymentGatesEvaluationConfiguration) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationConfiguration) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeploymentGatesEvaluationConfiguration) HasDryRun() bool {
	return o != nil && o.DryRun != nil
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeploymentGatesEvaluationConfiguration) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetRules returns the Rules field value.
func (o *DeploymentGatesEvaluationConfiguration) GetRules() []DeploymentGatesEvaluationRule {
	if o == nil {
		var ret []DeploymentGatesEvaluationRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationConfiguration) GetRulesOk() (*[]DeploymentGatesEvaluationRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *DeploymentGatesEvaluationConfiguration) SetRules(v []DeploymentGatesEvaluationRule) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGatesEvaluationConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	toSerialize["rules"] = o.Rules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGatesEvaluationConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DryRun *bool                            `json:"dry_run,omitempty"`
		Rules  *[]DeploymentGatesEvaluationRule `json:"rules"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dry_run", "rules"})
	} else {
		return err
	}
	o.DryRun = all.DryRun
	o.Rules = *all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
