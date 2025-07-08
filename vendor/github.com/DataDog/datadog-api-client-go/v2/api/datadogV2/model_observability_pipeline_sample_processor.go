// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSampleProcessor The `sample` processor allows probabilistic sampling of logs at a fixed rate.
type ObservabilityPipelineSampleProcessor struct {
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The percentage of logs to sample.
	Percentage *float64 `json:"percentage,omitempty"`
	// Number of events to sample (1 in N).
	Rate *int64 `json:"rate,omitempty"`
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
func NewObservabilityPipelineSampleProcessor(id string, include string, inputs []string, typeVar ObservabilityPipelineSampleProcessorType) *ObservabilityPipelineSampleProcessor {
	this := ObservabilityPipelineSampleProcessor{}
	this.Id = id
	this.Include = include
	this.Inputs = inputs
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

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineSampleProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSampleProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineSampleProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *ObservabilityPipelineSampleProcessor) GetPercentage() float64 {
	if o == nil || o.Percentage == nil {
		var ret float64
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSampleProcessor) GetPercentageOk() (*float64, bool) {
	if o == nil || o.Percentage == nil {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *ObservabilityPipelineSampleProcessor) HasPercentage() bool {
	return o != nil && o.Percentage != nil
}

// SetPercentage gets a reference to the given float64 and assigns it to the Percentage field.
func (o *ObservabilityPipelineSampleProcessor) SetPercentage(v float64) {
	o.Percentage = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ObservabilityPipelineSampleProcessor) GetRate() int64 {
	if o == nil || o.Rate == nil {
		var ret int64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSampleProcessor) GetRateOk() (*int64, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ObservabilityPipelineSampleProcessor) HasRate() bool {
	return o != nil && o.Rate != nil
}

// SetRate gets a reference to the given int64 and assigns it to the Rate field.
func (o *ObservabilityPipelineSampleProcessor) SetRate(v int64) {
	o.Rate = &v
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
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	if o.Percentage != nil {
		toSerialize["percentage"] = o.Percentage
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSampleProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id         *string                                   `json:"id"`
		Include    *string                                   `json:"include"`
		Inputs     *[]string                                 `json:"inputs"`
		Percentage *float64                                  `json:"percentage,omitempty"`
		Rate       *int64                                    `json:"rate,omitempty"`
		Type       *ObservabilityPipelineSampleProcessorType `json:"type"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "include", "inputs", "percentage", "rate", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	o.Percentage = all.Percentage
	o.Rate = all.Rate
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
