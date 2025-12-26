// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateDeploymentRuleParamsDataAttributes Parameters for updating a deployment rule.
type UpdateDeploymentRuleParamsDataAttributes struct {
	// Whether to run this rule in dry-run mode.
	DryRun bool `json:"dry_run"`
	// The name of the deployment rule.
	Name string `json:"name"`
	// Options for deployment rule response representing either faulty deployment detection or monitor options.
	Options DeploymentRulesOptions `json:"options"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateDeploymentRuleParamsDataAttributes instantiates a new UpdateDeploymentRuleParamsDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateDeploymentRuleParamsDataAttributes(dryRun bool, name string, options DeploymentRulesOptions) *UpdateDeploymentRuleParamsDataAttributes {
	this := UpdateDeploymentRuleParamsDataAttributes{}
	this.DryRun = dryRun
	this.Name = name
	this.Options = options
	return &this
}

// NewUpdateDeploymentRuleParamsDataAttributesWithDefaults instantiates a new UpdateDeploymentRuleParamsDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateDeploymentRuleParamsDataAttributesWithDefaults() *UpdateDeploymentRuleParamsDataAttributes {
	this := UpdateDeploymentRuleParamsDataAttributes{}
	return &this
}

// GetDryRun returns the DryRun field value.
func (o *UpdateDeploymentRuleParamsDataAttributes) GetDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value
// and a boolean to check if the value has been set.
func (o *UpdateDeploymentRuleParamsDataAttributes) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DryRun, true
}

// SetDryRun sets field value.
func (o *UpdateDeploymentRuleParamsDataAttributes) SetDryRun(v bool) {
	o.DryRun = v
}

// GetName returns the Name field value.
func (o *UpdateDeploymentRuleParamsDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateDeploymentRuleParamsDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *UpdateDeploymentRuleParamsDataAttributes) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value.
func (o *UpdateDeploymentRuleParamsDataAttributes) GetOptions() DeploymentRulesOptions {
	if o == nil {
		var ret DeploymentRulesOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UpdateDeploymentRuleParamsDataAttributes) GetOptionsOk() (*DeploymentRulesOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *UpdateDeploymentRuleParamsDataAttributes) SetOptions(v DeploymentRulesOptions) {
	o.Options = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateDeploymentRuleParamsDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dry_run"] = o.DryRun
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateDeploymentRuleParamsDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DryRun  *bool                   `json:"dry_run"`
		Name    *string                 `json:"name"`
		Options *DeploymentRulesOptions `json:"options"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DryRun == nil {
		return fmt.Errorf("required field dry_run missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dry_run", "name", "options"})
	} else {
		return err
	}
	o.DryRun = *all.DryRun
	o.Name = *all.Name
	o.Options = *all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
