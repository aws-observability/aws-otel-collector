// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineQuotaProcessor The Quota Processor measures logging traffic for logs that match a specified filter. When the configured daily quota is met, the processor can drop or alert.
type ObservabilityPipelineQuotaProcessor struct {
	// If set to `true`, logs that matched the quota filter and sent after the quota has been met are dropped; only logs that did not match the filter query continue through the pipeline.
	DropEvents bool `json:"drop_events"`
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// If `true`, the processor skips quota checks when partition fields are missing from the logs.
	IgnoreWhenMissingPartitions *bool `json:"ignore_when_missing_partitions,omitempty"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The maximum amount of data or number of events allowed before the quota is enforced. Can be specified in bytes or events.
	Limit ObservabilityPipelineQuotaProcessorLimit `json:"limit"`
	// Name of the quota.
	Name string `json:"name"`
	// The action to take when the quota is exceeded. Options:
	// - `drop`: Drop the event.
	// - `no_action`: Let the event pass through.
	// - `overflow_routing`: Route to an overflow destination.
	OverflowAction *ObservabilityPipelineQuotaProcessorOverflowAction `json:"overflow_action,omitempty"`
	// A list of alternate quota rules that apply to specific sets of events, identified by matching field values. Each override can define a custom limit.
	Overrides []ObservabilityPipelineQuotaProcessorOverride `json:"overrides,omitempty"`
	// A list of fields used to segment log traffic for quota enforcement. Quotas are tracked independently by unique combinations of these field values.
	PartitionFields []string `json:"partition_fields,omitempty"`
	// The processor type. The value should always be `quota`.
	Type ObservabilityPipelineQuotaProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineQuotaProcessor instantiates a new ObservabilityPipelineQuotaProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineQuotaProcessor(dropEvents bool, id string, include string, inputs []string, limit ObservabilityPipelineQuotaProcessorLimit, name string, typeVar ObservabilityPipelineQuotaProcessorType) *ObservabilityPipelineQuotaProcessor {
	this := ObservabilityPipelineQuotaProcessor{}
	this.DropEvents = dropEvents
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Limit = limit
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineQuotaProcessorWithDefaults instantiates a new ObservabilityPipelineQuotaProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineQuotaProcessorWithDefaults() *ObservabilityPipelineQuotaProcessor {
	this := ObservabilityPipelineQuotaProcessor{}
	var typeVar ObservabilityPipelineQuotaProcessorType = OBSERVABILITYPIPELINEQUOTAPROCESSORTYPE_QUOTA
	this.Type = typeVar
	return &this
}

// GetDropEvents returns the DropEvents field value.
func (o *ObservabilityPipelineQuotaProcessor) GetDropEvents() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DropEvents
}

// GetDropEventsOk returns a tuple with the DropEvents field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetDropEventsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DropEvents, true
}

// SetDropEvents sets field value.
func (o *ObservabilityPipelineQuotaProcessor) SetDropEvents(v bool) {
	o.DropEvents = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineQuotaProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineQuotaProcessor) SetId(v string) {
	o.Id = v
}

// GetIgnoreWhenMissingPartitions returns the IgnoreWhenMissingPartitions field value if set, zero value otherwise.
func (o *ObservabilityPipelineQuotaProcessor) GetIgnoreWhenMissingPartitions() bool {
	if o == nil || o.IgnoreWhenMissingPartitions == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreWhenMissingPartitions
}

// GetIgnoreWhenMissingPartitionsOk returns a tuple with the IgnoreWhenMissingPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetIgnoreWhenMissingPartitionsOk() (*bool, bool) {
	if o == nil || o.IgnoreWhenMissingPartitions == nil {
		return nil, false
	}
	return o.IgnoreWhenMissingPartitions, true
}

// HasIgnoreWhenMissingPartitions returns a boolean if a field has been set.
func (o *ObservabilityPipelineQuotaProcessor) HasIgnoreWhenMissingPartitions() bool {
	return o != nil && o.IgnoreWhenMissingPartitions != nil
}

// SetIgnoreWhenMissingPartitions gets a reference to the given bool and assigns it to the IgnoreWhenMissingPartitions field.
func (o *ObservabilityPipelineQuotaProcessor) SetIgnoreWhenMissingPartitions(v bool) {
	o.IgnoreWhenMissingPartitions = &v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineQuotaProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineQuotaProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineQuotaProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineQuotaProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetLimit returns the Limit field value.
func (o *ObservabilityPipelineQuotaProcessor) GetLimit() ObservabilityPipelineQuotaProcessorLimit {
	if o == nil {
		var ret ObservabilityPipelineQuotaProcessorLimit
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetLimitOk() (*ObservabilityPipelineQuotaProcessorLimit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *ObservabilityPipelineQuotaProcessor) SetLimit(v ObservabilityPipelineQuotaProcessorLimit) {
	o.Limit = v
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineQuotaProcessor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineQuotaProcessor) SetName(v string) {
	o.Name = v
}

// GetOverflowAction returns the OverflowAction field value if set, zero value otherwise.
func (o *ObservabilityPipelineQuotaProcessor) GetOverflowAction() ObservabilityPipelineQuotaProcessorOverflowAction {
	if o == nil || o.OverflowAction == nil {
		var ret ObservabilityPipelineQuotaProcessorOverflowAction
		return ret
	}
	return *o.OverflowAction
}

// GetOverflowActionOk returns a tuple with the OverflowAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetOverflowActionOk() (*ObservabilityPipelineQuotaProcessorOverflowAction, bool) {
	if o == nil || o.OverflowAction == nil {
		return nil, false
	}
	return o.OverflowAction, true
}

// HasOverflowAction returns a boolean if a field has been set.
func (o *ObservabilityPipelineQuotaProcessor) HasOverflowAction() bool {
	return o != nil && o.OverflowAction != nil
}

// SetOverflowAction gets a reference to the given ObservabilityPipelineQuotaProcessorOverflowAction and assigns it to the OverflowAction field.
func (o *ObservabilityPipelineQuotaProcessor) SetOverflowAction(v ObservabilityPipelineQuotaProcessorOverflowAction) {
	o.OverflowAction = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *ObservabilityPipelineQuotaProcessor) GetOverrides() []ObservabilityPipelineQuotaProcessorOverride {
	if o == nil || o.Overrides == nil {
		var ret []ObservabilityPipelineQuotaProcessorOverride
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetOverridesOk() (*[]ObservabilityPipelineQuotaProcessorOverride, bool) {
	if o == nil || o.Overrides == nil {
		return nil, false
	}
	return &o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *ObservabilityPipelineQuotaProcessor) HasOverrides() bool {
	return o != nil && o.Overrides != nil
}

// SetOverrides gets a reference to the given []ObservabilityPipelineQuotaProcessorOverride and assigns it to the Overrides field.
func (o *ObservabilityPipelineQuotaProcessor) SetOverrides(v []ObservabilityPipelineQuotaProcessorOverride) {
	o.Overrides = v
}

// GetPartitionFields returns the PartitionFields field value if set, zero value otherwise.
func (o *ObservabilityPipelineQuotaProcessor) GetPartitionFields() []string {
	if o == nil || o.PartitionFields == nil {
		var ret []string
		return ret
	}
	return o.PartitionFields
}

// GetPartitionFieldsOk returns a tuple with the PartitionFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetPartitionFieldsOk() (*[]string, bool) {
	if o == nil || o.PartitionFields == nil {
		return nil, false
	}
	return &o.PartitionFields, true
}

// HasPartitionFields returns a boolean if a field has been set.
func (o *ObservabilityPipelineQuotaProcessor) HasPartitionFields() bool {
	return o != nil && o.PartitionFields != nil
}

// SetPartitionFields gets a reference to the given []string and assigns it to the PartitionFields field.
func (o *ObservabilityPipelineQuotaProcessor) SetPartitionFields(v []string) {
	o.PartitionFields = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineQuotaProcessor) GetType() ObservabilityPipelineQuotaProcessorType {
	if o == nil {
		var ret ObservabilityPipelineQuotaProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetTypeOk() (*ObservabilityPipelineQuotaProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineQuotaProcessor) SetType(v ObservabilityPipelineQuotaProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineQuotaProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["drop_events"] = o.DropEvents
	toSerialize["id"] = o.Id
	if o.IgnoreWhenMissingPartitions != nil {
		toSerialize["ignore_when_missing_partitions"] = o.IgnoreWhenMissingPartitions
	}
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["limit"] = o.Limit
	toSerialize["name"] = o.Name
	if o.OverflowAction != nil {
		toSerialize["overflow_action"] = o.OverflowAction
	}
	if o.Overrides != nil {
		toSerialize["overrides"] = o.Overrides
	}
	if o.PartitionFields != nil {
		toSerialize["partition_fields"] = o.PartitionFields
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineQuotaProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DropEvents                  *bool                                              `json:"drop_events"`
		Id                          *string                                            `json:"id"`
		IgnoreWhenMissingPartitions *bool                                              `json:"ignore_when_missing_partitions,omitempty"`
		Include                     *string                                            `json:"include"`
		Inputs                      *[]string                                          `json:"inputs"`
		Limit                       *ObservabilityPipelineQuotaProcessorLimit          `json:"limit"`
		Name                        *string                                            `json:"name"`
		OverflowAction              *ObservabilityPipelineQuotaProcessorOverflowAction `json:"overflow_action,omitempty"`
		Overrides                   []ObservabilityPipelineQuotaProcessorOverride      `json:"overrides,omitempty"`
		PartitionFields             []string                                           `json:"partition_fields,omitempty"`
		Type                        *ObservabilityPipelineQuotaProcessorType           `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DropEvents == nil {
		return fmt.Errorf("required field drop_events missing")
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
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"drop_events", "id", "ignore_when_missing_partitions", "include", "inputs", "limit", "name", "overflow_action", "overrides", "partition_fields", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DropEvents = *all.DropEvents
	o.Id = *all.Id
	o.IgnoreWhenMissingPartitions = all.IgnoreWhenMissingPartitions
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	if all.Limit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Limit = *all.Limit
	o.Name = *all.Name
	if all.OverflowAction != nil && !all.OverflowAction.IsValid() {
		hasInvalidField = true
	} else {
		o.OverflowAction = all.OverflowAction
	}
	o.Overrides = all.Overrides
	o.PartitionFields = all.PartitionFields
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
