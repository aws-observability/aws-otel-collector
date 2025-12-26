// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateDeploymentRuleParamsDataAttributes Parameters for creating a deployment rule.
type CreateDeploymentRuleParamsDataAttributes struct {
	// Whether this rule is run in dry-run mode.
	DryRun *bool `json:"dry_run,omitempty"`
	// The name of the deployment rule.
	Name string `json:"name"`
	// Options for deployment rule response representing either faulty deployment detection or monitor options.
	Options DeploymentRulesOptions `json:"options"`
	// The type of the deployment rule (faulty_deployment_detection or monitor).
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateDeploymentRuleParamsDataAttributes instantiates a new CreateDeploymentRuleParamsDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateDeploymentRuleParamsDataAttributes(name string, options DeploymentRulesOptions, typeVar string) *CreateDeploymentRuleParamsDataAttributes {
	this := CreateDeploymentRuleParamsDataAttributes{}
	var dryRun bool = false
	this.DryRun = &dryRun
	this.Name = name
	this.Options = options
	this.Type = typeVar
	return &this
}

// NewCreateDeploymentRuleParamsDataAttributesWithDefaults instantiates a new CreateDeploymentRuleParamsDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateDeploymentRuleParamsDataAttributesWithDefaults() *CreateDeploymentRuleParamsDataAttributes {
	this := CreateDeploymentRuleParamsDataAttributes{}
	var dryRun bool = false
	this.DryRun = &dryRun
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateDeploymentRuleParamsDataAttributes) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeploymentRuleParamsDataAttributes) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateDeploymentRuleParamsDataAttributes) HasDryRun() bool {
	return o != nil && o.DryRun != nil
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateDeploymentRuleParamsDataAttributes) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetName returns the Name field value.
func (o *CreateDeploymentRuleParamsDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDeploymentRuleParamsDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateDeploymentRuleParamsDataAttributes) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value.
func (o *CreateDeploymentRuleParamsDataAttributes) GetOptions() DeploymentRulesOptions {
	if o == nil {
		var ret DeploymentRulesOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *CreateDeploymentRuleParamsDataAttributes) GetOptionsOk() (*DeploymentRulesOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *CreateDeploymentRuleParamsDataAttributes) SetOptions(v DeploymentRulesOptions) {
	o.Options = v
}

// GetType returns the Type field value.
func (o *CreateDeploymentRuleParamsDataAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateDeploymentRuleParamsDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateDeploymentRuleParamsDataAttributes) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateDeploymentRuleParamsDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateDeploymentRuleParamsDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DryRun  *bool                   `json:"dry_run,omitempty"`
		Name    *string                 `json:"name"`
		Options *DeploymentRulesOptions `json:"options"`
		Type    *string                 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dry_run", "name", "options", "type"})
	} else {
		return err
	}
	o.DryRun = all.DryRun
	o.Name = *all.Name
	o.Options = *all.Options
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
