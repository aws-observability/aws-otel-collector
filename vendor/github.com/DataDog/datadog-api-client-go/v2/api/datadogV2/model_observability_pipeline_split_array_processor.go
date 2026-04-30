// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplitArrayProcessor The `split_array` processor splits array fields into separate events based on configured rules.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineSplitArrayProcessor struct {
	// A list of array split configurations.
	Arrays []ObservabilityPipelineSplitArrayProcessorArrayConfig `json:"arrays"`
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets. For split_array, this should typically be `*`.
	Include string `json:"include"`
	// The processor type. The value should always be `split_array`.
	Type ObservabilityPipelineSplitArrayProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSplitArrayProcessor instantiates a new ObservabilityPipelineSplitArrayProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSplitArrayProcessor(arrays []ObservabilityPipelineSplitArrayProcessorArrayConfig, enabled bool, id string, include string, typeVar ObservabilityPipelineSplitArrayProcessorType) *ObservabilityPipelineSplitArrayProcessor {
	this := ObservabilityPipelineSplitArrayProcessor{}
	this.Arrays = arrays
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSplitArrayProcessorWithDefaults instantiates a new ObservabilityPipelineSplitArrayProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSplitArrayProcessorWithDefaults() *ObservabilityPipelineSplitArrayProcessor {
	this := ObservabilityPipelineSplitArrayProcessor{}
	var typeVar ObservabilityPipelineSplitArrayProcessorType = OBSERVABILITYPIPELINESPLITARRAYPROCESSORTYPE_SPLIT_ARRAY
	this.Type = typeVar
	return &this
}

// GetArrays returns the Arrays field value.
func (o *ObservabilityPipelineSplitArrayProcessor) GetArrays() []ObservabilityPipelineSplitArrayProcessorArrayConfig {
	if o == nil {
		var ret []ObservabilityPipelineSplitArrayProcessorArrayConfig
		return ret
	}
	return o.Arrays
}

// GetArraysOk returns a tuple with the Arrays field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplitArrayProcessor) GetArraysOk() (*[]ObservabilityPipelineSplitArrayProcessorArrayConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arrays, true
}

// SetArrays sets field value.
func (o *ObservabilityPipelineSplitArrayProcessor) SetArrays(v []ObservabilityPipelineSplitArrayProcessorArrayConfig) {
	o.Arrays = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplitArrayProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplitArrayProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplitArrayProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineSplitArrayProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineSplitArrayProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplitArrayProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineSplitArrayProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSplitArrayProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplitArrayProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSplitArrayProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineSplitArrayProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplitArrayProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineSplitArrayProcessor) SetInclude(v string) {
	o.Include = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSplitArrayProcessor) GetType() ObservabilityPipelineSplitArrayProcessorType {
	if o == nil {
		var ret ObservabilityPipelineSplitArrayProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplitArrayProcessor) GetTypeOk() (*ObservabilityPipelineSplitArrayProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSplitArrayProcessor) SetType(v ObservabilityPipelineSplitArrayProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSplitArrayProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["arrays"] = o.Arrays
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSplitArrayProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arrays      *[]ObservabilityPipelineSplitArrayProcessorArrayConfig `json:"arrays"`
		DisplayName *string                                                `json:"display_name,omitempty"`
		Enabled     *bool                                                  `json:"enabled"`
		Id          *string                                                `json:"id"`
		Include     *string                                                `json:"include"`
		Type        *ObservabilityPipelineSplitArrayProcessorType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Arrays == nil {
		return fmt.Errorf("required field arrays missing")
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arrays", "display_name", "enabled", "id", "include", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Arrays = *all.Arrays
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = *all.Include
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
