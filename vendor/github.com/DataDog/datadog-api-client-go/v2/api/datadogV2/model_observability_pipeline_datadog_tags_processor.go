// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatadogTagsProcessor The `datadog_tags` processor includes or excludes specific Datadog tags in your logs.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineDatadogTagsProcessor struct {
	// The action to take on tags with matching keys.
	Action ObservabilityPipelineDatadogTagsProcessorAction `json:"action"`
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of tag keys.
	Keys []string `json:"keys"`
	// The processing mode.
	Mode ObservabilityPipelineDatadogTagsProcessorMode `json:"mode"`
	// The processor type. The value should always be `datadog_tags`.
	Type ObservabilityPipelineDatadogTagsProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineDatadogTagsProcessor instantiates a new ObservabilityPipelineDatadogTagsProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineDatadogTagsProcessor(action ObservabilityPipelineDatadogTagsProcessorAction, enabled bool, id string, include string, keys []string, mode ObservabilityPipelineDatadogTagsProcessorMode, typeVar ObservabilityPipelineDatadogTagsProcessorType) *ObservabilityPipelineDatadogTagsProcessor {
	this := ObservabilityPipelineDatadogTagsProcessor{}
	this.Action = action
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Keys = keys
	this.Mode = mode
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineDatadogTagsProcessorWithDefaults instantiates a new ObservabilityPipelineDatadogTagsProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineDatadogTagsProcessorWithDefaults() *ObservabilityPipelineDatadogTagsProcessor {
	this := ObservabilityPipelineDatadogTagsProcessor{}
	var typeVar ObservabilityPipelineDatadogTagsProcessorType = OBSERVABILITYPIPELINEDATADOGTAGSPROCESSORTYPE_DATADOG_TAGS
	this.Type = typeVar
	return &this
}

// GetAction returns the Action field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetAction() ObservabilityPipelineDatadogTagsProcessorAction {
	if o == nil {
		var ret ObservabilityPipelineDatadogTagsProcessorAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetActionOk() (*ObservabilityPipelineDatadogTagsProcessorAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) SetAction(v ObservabilityPipelineDatadogTagsProcessorAction) {
	o.Action = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatadogTagsProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineDatadogTagsProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) SetInclude(v string) {
	o.Include = v
}

// GetKeys returns the Keys field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keys, true
}

// SetKeys sets field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) SetKeys(v []string) {
	o.Keys = v
}

// GetMode returns the Mode field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetMode() ObservabilityPipelineDatadogTagsProcessorMode {
	if o == nil {
		var ret ObservabilityPipelineDatadogTagsProcessorMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetModeOk() (*ObservabilityPipelineDatadogTagsProcessorMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) SetMode(v ObservabilityPipelineDatadogTagsProcessorMode) {
	o.Mode = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetType() ObservabilityPipelineDatadogTagsProcessorType {
	if o == nil {
		var ret ObservabilityPipelineDatadogTagsProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogTagsProcessor) GetTypeOk() (*ObservabilityPipelineDatadogTagsProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineDatadogTagsProcessor) SetType(v ObservabilityPipelineDatadogTagsProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineDatadogTagsProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["keys"] = o.Keys
	toSerialize["mode"] = o.Mode
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineDatadogTagsProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action      *ObservabilityPipelineDatadogTagsProcessorAction `json:"action"`
		DisplayName *string                                          `json:"display_name,omitempty"`
		Enabled     *bool                                            `json:"enabled"`
		Id          *string                                          `json:"id"`
		Include     *string                                          `json:"include"`
		Keys        *[]string                                        `json:"keys"`
		Mode        *ObservabilityPipelineDatadogTagsProcessorMode   `json:"mode"`
		Type        *ObservabilityPipelineDatadogTagsProcessorType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "display_name", "enabled", "id", "include", "keys", "mode", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = *all.Include
	o.Keys = *all.Keys
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
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
