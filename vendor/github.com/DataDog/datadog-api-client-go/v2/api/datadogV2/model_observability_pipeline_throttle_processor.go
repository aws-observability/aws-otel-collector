// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineThrottleProcessor The `throttle` processor limits the number of events that pass through over a given time window.
type ObservabilityPipelineThrottleProcessor struct {
	// Optional list of fields used to group events before the threshold has been reached.
	GroupBy []string `json:"group_by,omitempty"`
	// The unique identifier for this processor.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the input for this processor.
	Inputs []string `json:"inputs"`
	// the number of events allowed in a given time window. Events sent after the threshold has been reached, are dropped.
	Threshold int64 `json:"threshold"`
	// The processor type. The value should always be `throttle`.
	Type ObservabilityPipelineThrottleProcessorType `json:"type"`
	// The time window in seconds over which the threshold applies.
	Window float64 `json:"window"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineThrottleProcessor instantiates a new ObservabilityPipelineThrottleProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineThrottleProcessor(id string, include string, inputs []string, threshold int64, typeVar ObservabilityPipelineThrottleProcessorType, window float64) *ObservabilityPipelineThrottleProcessor {
	this := ObservabilityPipelineThrottleProcessor{}
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Threshold = threshold
	this.Type = typeVar
	this.Window = window
	return &this
}

// NewObservabilityPipelineThrottleProcessorWithDefaults instantiates a new ObservabilityPipelineThrottleProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineThrottleProcessorWithDefaults() *ObservabilityPipelineThrottleProcessor {
	this := ObservabilityPipelineThrottleProcessor{}
	var typeVar ObservabilityPipelineThrottleProcessorType = OBSERVABILITYPIPELINETHROTTLEPROCESSORTYPE_THROTTLE
	this.Type = typeVar
	return &this
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ObservabilityPipelineThrottleProcessor) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineThrottleProcessor) GetGroupByOk() (*[]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ObservabilityPipelineThrottleProcessor) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *ObservabilityPipelineThrottleProcessor) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineThrottleProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineThrottleProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineThrottleProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineThrottleProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineThrottleProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineThrottleProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineThrottleProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineThrottleProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineThrottleProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetThreshold returns the Threshold field value.
func (o *ObservabilityPipelineThrottleProcessor) GetThreshold() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineThrottleProcessor) GetThresholdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value.
func (o *ObservabilityPipelineThrottleProcessor) SetThreshold(v int64) {
	o.Threshold = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineThrottleProcessor) GetType() ObservabilityPipelineThrottleProcessorType {
	if o == nil {
		var ret ObservabilityPipelineThrottleProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineThrottleProcessor) GetTypeOk() (*ObservabilityPipelineThrottleProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineThrottleProcessor) SetType(v ObservabilityPipelineThrottleProcessorType) {
	o.Type = v
}

// GetWindow returns the Window field value.
func (o *ObservabilityPipelineThrottleProcessor) GetWindow() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineThrottleProcessor) GetWindowOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value.
func (o *ObservabilityPipelineThrottleProcessor) SetWindow(v float64) {
	o.Window = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineThrottleProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["threshold"] = o.Threshold
	toSerialize["type"] = o.Type
	toSerialize["window"] = o.Window

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineThrottleProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupBy   []string                                    `json:"group_by,omitempty"`
		Id        *string                                     `json:"id"`
		Include   *string                                     `json:"include"`
		Inputs    *[]string                                   `json:"inputs"`
		Threshold *int64                                      `json:"threshold"`
		Type      *ObservabilityPipelineThrottleProcessorType `json:"type"`
		Window    *float64                                    `json:"window"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Threshold == nil {
		return fmt.Errorf("required field threshold missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Window == nil {
		return fmt.Errorf("required field window missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group_by", "id", "include", "inputs", "threshold", "type", "window"})
	} else {
		return err
	}

	hasInvalidField := false
	o.GroupBy = all.GroupBy
	o.Id = *all.Id
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	o.Threshold = *all.Threshold
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Window = *all.Window

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
