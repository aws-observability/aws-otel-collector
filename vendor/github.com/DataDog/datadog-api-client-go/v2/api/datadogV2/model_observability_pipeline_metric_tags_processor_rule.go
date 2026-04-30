// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricTagsProcessorRule Defines a rule for filtering metric tags based on key patterns.
type ObservabilityPipelineMetricTagsProcessorRule struct {
	// The action to take on tags with matching keys.
	Action ObservabilityPipelineMetricTagsProcessorRuleAction `json:"action"`
	// A Datadog search query used to determine which metrics this rule targets.
	Include string `json:"include"`
	// A list of tag keys to include or exclude.
	Keys []string `json:"keys"`
	// The processing mode for tag filtering.
	Mode ObservabilityPipelineMetricTagsProcessorRuleMode `json:"mode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineMetricTagsProcessorRule instantiates a new ObservabilityPipelineMetricTagsProcessorRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineMetricTagsProcessorRule(action ObservabilityPipelineMetricTagsProcessorRuleAction, include string, keys []string, mode ObservabilityPipelineMetricTagsProcessorRuleMode) *ObservabilityPipelineMetricTagsProcessorRule {
	this := ObservabilityPipelineMetricTagsProcessorRule{}
	this.Action = action
	this.Include = include
	this.Keys = keys
	this.Mode = mode
	return &this
}

// NewObservabilityPipelineMetricTagsProcessorRuleWithDefaults instantiates a new ObservabilityPipelineMetricTagsProcessorRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineMetricTagsProcessorRuleWithDefaults() *ObservabilityPipelineMetricTagsProcessorRule {
	this := ObservabilityPipelineMetricTagsProcessorRule{}
	return &this
}

// GetAction returns the Action field value.
func (o *ObservabilityPipelineMetricTagsProcessorRule) GetAction() ObservabilityPipelineMetricTagsProcessorRuleAction {
	if o == nil {
		var ret ObservabilityPipelineMetricTagsProcessorRuleAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricTagsProcessorRule) GetActionOk() (*ObservabilityPipelineMetricTagsProcessorRuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *ObservabilityPipelineMetricTagsProcessorRule) SetAction(v ObservabilityPipelineMetricTagsProcessorRuleAction) {
	o.Action = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineMetricTagsProcessorRule) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricTagsProcessorRule) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineMetricTagsProcessorRule) SetInclude(v string) {
	o.Include = v
}

// GetKeys returns the Keys field value.
func (o *ObservabilityPipelineMetricTagsProcessorRule) GetKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricTagsProcessorRule) GetKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keys, true
}

// SetKeys sets field value.
func (o *ObservabilityPipelineMetricTagsProcessorRule) SetKeys(v []string) {
	o.Keys = v
}

// GetMode returns the Mode field value.
func (o *ObservabilityPipelineMetricTagsProcessorRule) GetMode() ObservabilityPipelineMetricTagsProcessorRuleMode {
	if o == nil {
		var ret ObservabilityPipelineMetricTagsProcessorRuleMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMetricTagsProcessorRule) GetModeOk() (*ObservabilityPipelineMetricTagsProcessorRuleMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ObservabilityPipelineMetricTagsProcessorRule) SetMode(v ObservabilityPipelineMetricTagsProcessorRuleMode) {
	o.Mode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineMetricTagsProcessorRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["include"] = o.Include
	toSerialize["keys"] = o.Keys
	toSerialize["mode"] = o.Mode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineMetricTagsProcessorRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action  *ObservabilityPipelineMetricTagsProcessorRuleAction `json:"action"`
		Include *string                                             `json:"include"`
		Keys    *[]string                                           `json:"keys"`
		Mode    *ObservabilityPipelineMetricTagsProcessorRuleMode   `json:"mode"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Keys == nil {
		return fmt.Errorf("required field keys missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "include", "keys", "mode"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.Include = *all.Include
	o.Keys = *all.Keys
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
