// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlMitigationDefinition The definition of a mitigation available for a control.
type GovernanceControlMitigationDefinition struct {
	// The verb describing the mitigation action, such as `revoke` or `delete`.
	ActionVerb string `json:"action_verb"`
	// A human-readable description of the mitigation.
	Description string `json:"description"`
	// The execution modes the mitigation supports, such as `manual` or `automatic`.
	ExecutionModes []string `json:"execution_modes,omitempty"`
	// The feature flags that gate the mitigation.
	FeatureFlags []string `json:"feature_flags"`
	// The unique identifier of the mitigation.
	Id string `json:"id"`
	// A warning shown to the user before applying the mitigation manually.
	ManualMitigationWarning string `json:"manual_mitigation_warning"`
	// The permissions required to apply the mitigation.
	Permissions []string `json:"permissions"`
	// Whether the mitigation requires AI to be enabled.
	RequiresAi bool `json:"requires_ai"`
	// An array of parameter definitions.
	SupportedParameters []GovernanceControlParameterDefinition `json:"supported_parameters"`
	// A short, human-readable name for the mitigation.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceControlMitigationDefinition instantiates a new GovernanceControlMitigationDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceControlMitigationDefinition(actionVerb string, description string, featureFlags []string, id string, manualMitigationWarning string, permissions []string, requiresAi bool, supportedParameters []GovernanceControlParameterDefinition, title string) *GovernanceControlMitigationDefinition {
	this := GovernanceControlMitigationDefinition{}
	this.ActionVerb = actionVerb
	this.Description = description
	this.FeatureFlags = featureFlags
	this.Id = id
	this.ManualMitigationWarning = manualMitigationWarning
	this.Permissions = permissions
	this.RequiresAi = requiresAi
	this.SupportedParameters = supportedParameters
	this.Title = title
	return &this
}

// NewGovernanceControlMitigationDefinitionWithDefaults instantiates a new GovernanceControlMitigationDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceControlMitigationDefinitionWithDefaults() *GovernanceControlMitigationDefinition {
	this := GovernanceControlMitigationDefinition{}
	return &this
}

// GetActionVerb returns the ActionVerb field value.
func (o *GovernanceControlMitigationDefinition) GetActionVerb() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionVerb
}

// GetActionVerbOk returns a tuple with the ActionVerb field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetActionVerbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionVerb, true
}

// SetActionVerb sets field value.
func (o *GovernanceControlMitigationDefinition) SetActionVerb(v string) {
	o.ActionVerb = v
}

// GetDescription returns the Description field value.
func (o *GovernanceControlMitigationDefinition) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GovernanceControlMitigationDefinition) SetDescription(v string) {
	o.Description = v
}

// GetExecutionModes returns the ExecutionModes field value if set, zero value otherwise.
func (o *GovernanceControlMitigationDefinition) GetExecutionModes() []string {
	if o == nil || o.ExecutionModes == nil {
		var ret []string
		return ret
	}
	return o.ExecutionModes
}

// GetExecutionModesOk returns a tuple with the ExecutionModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetExecutionModesOk() (*[]string, bool) {
	if o == nil || o.ExecutionModes == nil {
		return nil, false
	}
	return &o.ExecutionModes, true
}

// HasExecutionModes returns a boolean if a field has been set.
func (o *GovernanceControlMitigationDefinition) HasExecutionModes() bool {
	return o != nil && o.ExecutionModes != nil
}

// SetExecutionModes gets a reference to the given []string and assigns it to the ExecutionModes field.
func (o *GovernanceControlMitigationDefinition) SetExecutionModes(v []string) {
	o.ExecutionModes = v
}

// GetFeatureFlags returns the FeatureFlags field value.
func (o *GovernanceControlMitigationDefinition) GetFeatureFlags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetFeatureFlagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlags, true
}

// SetFeatureFlags sets field value.
func (o *GovernanceControlMitigationDefinition) SetFeatureFlags(v []string) {
	o.FeatureFlags = v
}

// GetId returns the Id field value.
func (o *GovernanceControlMitigationDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *GovernanceControlMitigationDefinition) SetId(v string) {
	o.Id = v
}

// GetManualMitigationWarning returns the ManualMitigationWarning field value.
func (o *GovernanceControlMitigationDefinition) GetManualMitigationWarning() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ManualMitigationWarning
}

// GetManualMitigationWarningOk returns a tuple with the ManualMitigationWarning field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetManualMitigationWarningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManualMitigationWarning, true
}

// SetManualMitigationWarning sets field value.
func (o *GovernanceControlMitigationDefinition) SetManualMitigationWarning(v string) {
	o.ManualMitigationWarning = v
}

// GetPermissions returns the Permissions field value.
func (o *GovernanceControlMitigationDefinition) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetPermissionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value.
func (o *GovernanceControlMitigationDefinition) SetPermissions(v []string) {
	o.Permissions = v
}

// GetRequiresAi returns the RequiresAi field value.
func (o *GovernanceControlMitigationDefinition) GetRequiresAi() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.RequiresAi
}

// GetRequiresAiOk returns a tuple with the RequiresAi field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetRequiresAiOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiresAi, true
}

// SetRequiresAi sets field value.
func (o *GovernanceControlMitigationDefinition) SetRequiresAi(v bool) {
	o.RequiresAi = v
}

// GetSupportedParameters returns the SupportedParameters field value.
func (o *GovernanceControlMitigationDefinition) GetSupportedParameters() []GovernanceControlParameterDefinition {
	if o == nil {
		var ret []GovernanceControlParameterDefinition
		return ret
	}
	return o.SupportedParameters
}

// GetSupportedParametersOk returns a tuple with the SupportedParameters field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetSupportedParametersOk() (*[]GovernanceControlParameterDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedParameters, true
}

// SetSupportedParameters sets field value.
func (o *GovernanceControlMitigationDefinition) SetSupportedParameters(v []GovernanceControlParameterDefinition) {
	o.SupportedParameters = v
}

// GetTitle returns the Title field value.
func (o *GovernanceControlMitigationDefinition) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlMitigationDefinition) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *GovernanceControlMitigationDefinition) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceControlMitigationDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action_verb"] = o.ActionVerb
	toSerialize["description"] = o.Description
	if o.ExecutionModes != nil {
		toSerialize["execution_modes"] = o.ExecutionModes
	}
	toSerialize["feature_flags"] = o.FeatureFlags
	toSerialize["id"] = o.Id
	toSerialize["manual_mitigation_warning"] = o.ManualMitigationWarning
	toSerialize["permissions"] = o.Permissions
	toSerialize["requires_ai"] = o.RequiresAi
	toSerialize["supported_parameters"] = o.SupportedParameters
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceControlMitigationDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionVerb              *string                                 `json:"action_verb"`
		Description             *string                                 `json:"description"`
		ExecutionModes          []string                                `json:"execution_modes,omitempty"`
		FeatureFlags            *[]string                               `json:"feature_flags"`
		Id                      *string                                 `json:"id"`
		ManualMitigationWarning *string                                 `json:"manual_mitigation_warning"`
		Permissions             *[]string                               `json:"permissions"`
		RequiresAi              *bool                                   `json:"requires_ai"`
		SupportedParameters     *[]GovernanceControlParameterDefinition `json:"supported_parameters"`
		Title                   *string                                 `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActionVerb == nil {
		return fmt.Errorf("required field action_verb missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.FeatureFlags == nil {
		return fmt.Errorf("required field feature_flags missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.ManualMitigationWarning == nil {
		return fmt.Errorf("required field manual_mitigation_warning missing")
	}
	if all.Permissions == nil {
		return fmt.Errorf("required field permissions missing")
	}
	if all.RequiresAi == nil {
		return fmt.Errorf("required field requires_ai missing")
	}
	if all.SupportedParameters == nil {
		return fmt.Errorf("required field supported_parameters missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action_verb", "description", "execution_modes", "feature_flags", "id", "manual_mitigation_warning", "permissions", "requires_ai", "supported_parameters", "title"})
	} else {
		return err
	}
	o.ActionVerb = *all.ActionVerb
	o.Description = *all.Description
	o.ExecutionModes = all.ExecutionModes
	o.FeatureFlags = *all.FeatureFlags
	o.Id = *all.Id
	o.ManualMitigationWarning = *all.ManualMitigationWarning
	o.Permissions = *all.Permissions
	o.RequiresAi = *all.RequiresAi
	o.SupportedParameters = *all.SupportedParameters
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
