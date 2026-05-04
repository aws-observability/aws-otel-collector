// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineQuotaProcessor The `quota` processor measures logging traffic for logs that match a specified filter. When the configured daily quota is met, the processor can drop or alert.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineQuotaProcessor struct {
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// If set to `true`, logs that match the quota filter and are sent after the quota is exceeded are dropped. Logs that do not match the filter continue through the pipeline. **Note**: You can set either `drop_events` or `overflow_action`, but not both.
	DropEvents *bool `json:"drop_events,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// If `true`, the processor skips quota checks when partition fields are missing from the logs.
	IgnoreWhenMissingPartitions *bool `json:"ignore_when_missing_partitions,omitempty"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// The maximum amount of data or number of events allowed before the quota is enforced. Can be specified in bytes or events.
	Limit ObservabilityPipelineQuotaProcessorLimit `json:"limit"`
	// Name of the quota.
	Name string `json:"name"`
	// The action to take when the quota or bucket limit is exceeded. Options:
	// - `drop`: Drop the event.
	// - `no_action`: Let the event pass through.
	// - `overflow_routing`: Route to an overflow destination.
	OverflowAction *ObservabilityPipelineQuotaProcessorOverflowAction `json:"overflow_action,omitempty"`
	// A list of alternate quota rules that apply to specific sets of events, identified by matching field values. Each override can define a custom limit.
	Overrides []ObservabilityPipelineQuotaProcessorOverride `json:"overrides,omitempty"`
	// A list of fields used to segment log traffic for quota enforcement. Quotas are tracked independently by unique combinations of these field values.
	PartitionFields []string `json:"partition_fields,omitempty"`
	// The action to take when the quota or bucket limit is exceeded. Options:
	// - `drop`: Drop the event.
	// - `no_action`: Let the event pass through.
	// - `overflow_routing`: Route to an overflow destination.
	TooManyBucketsAction *ObservabilityPipelineQuotaProcessorOverflowAction `json:"too_many_buckets_action,omitempty"`
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
func NewObservabilityPipelineQuotaProcessor(enabled bool, id string, include string, limit ObservabilityPipelineQuotaProcessorLimit, name string, typeVar ObservabilityPipelineQuotaProcessorType) *ObservabilityPipelineQuotaProcessor {
	this := ObservabilityPipelineQuotaProcessor{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
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

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineQuotaProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineQuotaProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineQuotaProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDropEvents returns the DropEvents field value if set, zero value otherwise.
func (o *ObservabilityPipelineQuotaProcessor) GetDropEvents() bool {
	if o == nil || o.DropEvents == nil {
		var ret bool
		return ret
	}
	return *o.DropEvents
}

// GetDropEventsOk returns a tuple with the DropEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetDropEventsOk() (*bool, bool) {
	if o == nil || o.DropEvents == nil {
		return nil, false
	}
	return o.DropEvents, true
}

// HasDropEvents returns a boolean if a field has been set.
func (o *ObservabilityPipelineQuotaProcessor) HasDropEvents() bool {
	return o != nil && o.DropEvents != nil
}

// SetDropEvents gets a reference to the given bool and assigns it to the DropEvents field.
func (o *ObservabilityPipelineQuotaProcessor) SetDropEvents(v bool) {
	o.DropEvents = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineQuotaProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineQuotaProcessor) SetEnabled(v bool) {
	o.Enabled = v
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

// GetTooManyBucketsAction returns the TooManyBucketsAction field value if set, zero value otherwise.
func (o *ObservabilityPipelineQuotaProcessor) GetTooManyBucketsAction() ObservabilityPipelineQuotaProcessorOverflowAction {
	if o == nil || o.TooManyBucketsAction == nil {
		var ret ObservabilityPipelineQuotaProcessorOverflowAction
		return ret
	}
	return *o.TooManyBucketsAction
}

// GetTooManyBucketsActionOk returns a tuple with the TooManyBucketsAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessor) GetTooManyBucketsActionOk() (*ObservabilityPipelineQuotaProcessorOverflowAction, bool) {
	if o == nil || o.TooManyBucketsAction == nil {
		return nil, false
	}
	return o.TooManyBucketsAction, true
}

// HasTooManyBucketsAction returns a boolean if a field has been set.
func (o *ObservabilityPipelineQuotaProcessor) HasTooManyBucketsAction() bool {
	return o != nil && o.TooManyBucketsAction != nil
}

// SetTooManyBucketsAction gets a reference to the given ObservabilityPipelineQuotaProcessorOverflowAction and assigns it to the TooManyBucketsAction field.
func (o *ObservabilityPipelineQuotaProcessor) SetTooManyBucketsAction(v ObservabilityPipelineQuotaProcessorOverflowAction) {
	o.TooManyBucketsAction = &v
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
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.DropEvents != nil {
		toSerialize["drop_events"] = o.DropEvents
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	if o.IgnoreWhenMissingPartitions != nil {
		toSerialize["ignore_when_missing_partitions"] = o.IgnoreWhenMissingPartitions
	}
	toSerialize["include"] = o.Include
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
	if o.TooManyBucketsAction != nil {
		toSerialize["too_many_buckets_action"] = o.TooManyBucketsAction
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
		DisplayName                 *string                                            `json:"display_name,omitempty"`
		DropEvents                  *bool                                              `json:"drop_events,omitempty"`
		Enabled                     *bool                                              `json:"enabled"`
		Id                          *string                                            `json:"id"`
		IgnoreWhenMissingPartitions *bool                                              `json:"ignore_when_missing_partitions,omitempty"`
		Include                     *string                                            `json:"include"`
		Limit                       *ObservabilityPipelineQuotaProcessorLimit          `json:"limit"`
		Name                        *string                                            `json:"name"`
		OverflowAction              *ObservabilityPipelineQuotaProcessorOverflowAction `json:"overflow_action,omitempty"`
		Overrides                   []ObservabilityPipelineQuotaProcessorOverride      `json:"overrides,omitempty"`
		PartitionFields             []string                                           `json:"partition_fields,omitempty"`
		TooManyBucketsAction        *ObservabilityPipelineQuotaProcessorOverflowAction `json:"too_many_buckets_action,omitempty"`
		Type                        *ObservabilityPipelineQuotaProcessorType           `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "drop_events", "enabled", "id", "ignore_when_missing_partitions", "include", "limit", "name", "overflow_action", "overrides", "partition_fields", "too_many_buckets_action", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	o.DropEvents = all.DropEvents
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.IgnoreWhenMissingPartitions = all.IgnoreWhenMissingPartitions
	o.Include = *all.Include
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
	if all.TooManyBucketsAction != nil && !all.TooManyBucketsAction.IsValid() {
		hasInvalidField = true
	} else {
		o.TooManyBucketsAction = all.TooManyBucketsAction
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
