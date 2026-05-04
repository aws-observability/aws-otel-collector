// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSampleProcessor The `sample` processor allows probabilistic sampling of logs at a fixed rate.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineSampleProcessor struct {
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// Optional list of fields to group events by. Each group is sampled independently.
	GroupBy []string `json:"group_by,omitempty"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// The percentage of logs to sample.
	Percentage float64 `json:"percentage"`
	// The processor type. The value should always be `sample`.
	Type ObservabilityPipelineSampleProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSampleProcessor instantiates a new ObservabilityPipelineSampleProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSampleProcessor(enabled bool, id string, include string, percentage float64, typeVar ObservabilityPipelineSampleProcessorType) *ObservabilityPipelineSampleProcessor {
	this := ObservabilityPipelineSampleProcessor{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Percentage = percentage
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSampleProcessorWithDefaults instantiates a new ObservabilityPipelineSampleProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSampleProcessorWithDefaults() *ObservabilityPipelineSampleProcessor {
	this := ObservabilityPipelineSampleProcessor{}
	var typeVar ObservabilityPipelineSampleProcessorType = OBSERVABILITYPIPELINESAMPLEPROCESSORTYPE_SAMPLE
	this.Type = typeVar
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineSampleProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSampleProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineSampleProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineSampleProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineSampleProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSampleProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineSampleProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ObservabilityPipelineSampleProcessor) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSampleProcessor) GetGroupByOk() (*[]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ObservabilityPipelineSampleProcessor) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *ObservabilityPipelineSampleProcessor) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSampleProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSampleProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSampleProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineSampleProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSampleProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineSampleProcessor) SetInclude(v string) {
	o.Include = v
}

// GetPercentage returns the Percentage field value.
func (o *ObservabilityPipelineSampleProcessor) GetPercentage() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSampleProcessor) GetPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value.
func (o *ObservabilityPipelineSampleProcessor) SetPercentage(v float64) {
	o.Percentage = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSampleProcessor) GetType() ObservabilityPipelineSampleProcessorType {
	if o == nil {
		var ret ObservabilityPipelineSampleProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSampleProcessor) GetTypeOk() (*ObservabilityPipelineSampleProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSampleProcessor) SetType(v ObservabilityPipelineSampleProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSampleProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["enabled"] = o.Enabled
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["percentage"] = o.Percentage
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSampleProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string                                   `json:"display_name,omitempty"`
		Enabled     *bool                                     `json:"enabled"`
		GroupBy     []string                                  `json:"group_by,omitempty"`
		Id          *string                                   `json:"id"`
		Include     *string                                   `json:"include"`
		Percentage  *float64                                  `json:"percentage"`
		Type        *ObservabilityPipelineSampleProcessorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Percentage == nil {
		return fmt.Errorf("required field percentage missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "enabled", "group_by", "id", "include", "percentage", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.GroupBy = all.GroupBy
	o.Id = *all.Id
	o.Include = *all.Include
	o.Percentage = *all.Percentage
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
