// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpanAttributes Attributes of an LLM Observability span.
type LLMObsSpanAttributes struct {
	// Duration of the span in nanoseconds.
	Duration float64 `json:"duration"`
	// Evaluation metrics keyed by evaluator name.
	Evaluation map[string]LLMObsSpanEvaluationMetric `json:"evaluation,omitempty"`
	// Input or output content of an LLM Observability span.
	Input *LLMObsSpanIO `json:"input,omitempty"`
	// Detected intent of the span.
	Intent *string `json:"intent,omitempty"`
	// Arbitrary metadata associated with the span.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Numeric metrics associated with the span (e.g., token counts).
	Metrics map[string]float64 `json:"metrics,omitempty"`
	// Name of the ML application this span belongs to.
	MlApp string `json:"ml_app"`
	// Name of the model used in this span.
	ModelName *string `json:"model_name,omitempty"`
	// Provider of the model used in this span.
	ModelProvider *string `json:"model_provider,omitempty"`
	// Name of the span.
	Name string `json:"name"`
	// Input or output content of an LLM Observability span.
	Output *LLMObsSpanIO `json:"output,omitempty"`
	// Identifier of the parent span, if any.
	ParentId *string `json:"parent_id,omitempty"`
	// Unique identifier of the span.
	SpanId string `json:"span_id"`
	// Kind of span (e.g., llm, agent, tool, task, workflow).
	SpanKind string `json:"span_kind"`
	// Start time of the span in nanoseconds since Unix epoch.
	StartNs int64 `json:"start_ns"`
	// Status of the span (e.g., ok, error).
	Status string `json:"status"`
	// Tags associated with the span.
	Tags []string `json:"tags,omitempty"`
	// Tool definitions available to the span.
	ToolDefinitions []LLMObsSpanToolDefinition `json:"tool_definitions,omitempty"`
	// Trace identifier this span belongs to.
	TraceId string `json:"trace_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSpanAttributes instantiates a new LLMObsSpanAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSpanAttributes(duration float64, mlApp string, name string, spanId string, spanKind string, startNs int64, status string, traceId string) *LLMObsSpanAttributes {
	this := LLMObsSpanAttributes{}
	this.Duration = duration
	this.MlApp = mlApp
	this.Name = name
	this.SpanId = spanId
	this.SpanKind = spanKind
	this.StartNs = startNs
	this.Status = status
	this.TraceId = traceId
	return &this
}

// NewLLMObsSpanAttributesWithDefaults instantiates a new LLMObsSpanAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSpanAttributesWithDefaults() *LLMObsSpanAttributes {
	this := LLMObsSpanAttributes{}
	return &this
}

// GetDuration returns the Duration field value.
func (o *LLMObsSpanAttributes) GetDuration() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetDurationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value.
func (o *LLMObsSpanAttributes) SetDuration(v float64) {
	o.Duration = v
}

// GetEvaluation returns the Evaluation field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetEvaluation() map[string]LLMObsSpanEvaluationMetric {
	if o == nil || o.Evaluation == nil {
		var ret map[string]LLMObsSpanEvaluationMetric
		return ret
	}
	return o.Evaluation
}

// GetEvaluationOk returns a tuple with the Evaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetEvaluationOk() (*map[string]LLMObsSpanEvaluationMetric, bool) {
	if o == nil || o.Evaluation == nil {
		return nil, false
	}
	return &o.Evaluation, true
}

// HasEvaluation returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasEvaluation() bool {
	return o != nil && o.Evaluation != nil
}

// SetEvaluation gets a reference to the given map[string]LLMObsSpanEvaluationMetric and assigns it to the Evaluation field.
func (o *LLMObsSpanAttributes) SetEvaluation(v map[string]LLMObsSpanEvaluationMetric) {
	o.Evaluation = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetInput() LLMObsSpanIO {
	if o == nil || o.Input == nil {
		var ret LLMObsSpanIO
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetInputOk() (*LLMObsSpanIO, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasInput() bool {
	return o != nil && o.Input != nil
}

// SetInput gets a reference to the given LLMObsSpanIO and assigns it to the Input field.
func (o *LLMObsSpanAttributes) SetInput(v LLMObsSpanIO) {
	o.Input = &v
}

// GetIntent returns the Intent field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetIntent() string {
	if o == nil || o.Intent == nil {
		var ret string
		return ret
	}
	return *o.Intent
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetIntentOk() (*string, bool) {
	if o == nil || o.Intent == nil {
		return nil, false
	}
	return o.Intent, true
}

// HasIntent returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasIntent() bool {
	return o != nil && o.Intent != nil
}

// SetIntent gets a reference to the given string and assigns it to the Intent field.
func (o *LLMObsSpanAttributes) SetIntent(v string) {
	o.Intent = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LLMObsSpanAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetMetrics() map[string]float64 {
	if o == nil || o.Metrics == nil {
		var ret map[string]float64
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetMetricsOk() (*map[string]float64, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given map[string]float64 and assigns it to the Metrics field.
func (o *LLMObsSpanAttributes) SetMetrics(v map[string]float64) {
	o.Metrics = v
}

// GetMlApp returns the MlApp field value.
func (o *LLMObsSpanAttributes) GetMlApp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MlApp
}

// GetMlAppOk returns a tuple with the MlApp field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetMlAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MlApp, true
}

// SetMlApp sets field value.
func (o *LLMObsSpanAttributes) SetMlApp(v string) {
	o.MlApp = v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetModelName() string {
	if o == nil || o.ModelName == nil {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetModelNameOk() (*string, bool) {
	if o == nil || o.ModelName == nil {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasModelName() bool {
	return o != nil && o.ModelName != nil
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *LLMObsSpanAttributes) SetModelName(v string) {
	o.ModelName = &v
}

// GetModelProvider returns the ModelProvider field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetModelProvider() string {
	if o == nil || o.ModelProvider == nil {
		var ret string
		return ret
	}
	return *o.ModelProvider
}

// GetModelProviderOk returns a tuple with the ModelProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetModelProviderOk() (*string, bool) {
	if o == nil || o.ModelProvider == nil {
		return nil, false
	}
	return o.ModelProvider, true
}

// HasModelProvider returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasModelProvider() bool {
	return o != nil && o.ModelProvider != nil
}

// SetModelProvider gets a reference to the given string and assigns it to the ModelProvider field.
func (o *LLMObsSpanAttributes) SetModelProvider(v string) {
	o.ModelProvider = &v
}

// GetName returns the Name field value.
func (o *LLMObsSpanAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsSpanAttributes) SetName(v string) {
	o.Name = v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetOutput() LLMObsSpanIO {
	if o == nil || o.Output == nil {
		var ret LLMObsSpanIO
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetOutputOk() (*LLMObsSpanIO, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasOutput() bool {
	return o != nil && o.Output != nil
}

// SetOutput gets a reference to the given LLMObsSpanIO and assigns it to the Output field.
func (o *LLMObsSpanAttributes) SetOutput(v LLMObsSpanIO) {
	o.Output = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasParentId() bool {
	return o != nil && o.ParentId != nil
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *LLMObsSpanAttributes) SetParentId(v string) {
	o.ParentId = &v
}

// GetSpanId returns the SpanId field value.
func (o *LLMObsSpanAttributes) GetSpanId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetSpanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanId, true
}

// SetSpanId sets field value.
func (o *LLMObsSpanAttributes) SetSpanId(v string) {
	o.SpanId = v
}

// GetSpanKind returns the SpanKind field value.
func (o *LLMObsSpanAttributes) GetSpanKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpanKind
}

// GetSpanKindOk returns a tuple with the SpanKind field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetSpanKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanKind, true
}

// SetSpanKind sets field value.
func (o *LLMObsSpanAttributes) SetSpanKind(v string) {
	o.SpanKind = v
}

// GetStartNs returns the StartNs field value.
func (o *LLMObsSpanAttributes) GetStartNs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartNs
}

// GetStartNsOk returns a tuple with the StartNs field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetStartNsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartNs, true
}

// SetStartNs sets field value.
func (o *LLMObsSpanAttributes) SetStartNs(v int64) {
	o.StartNs = v
}

// GetStatus returns the Status field value.
func (o *LLMObsSpanAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *LLMObsSpanAttributes) SetStatus(v string) {
	o.Status = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsSpanAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetToolDefinitions returns the ToolDefinitions field value if set, zero value otherwise.
func (o *LLMObsSpanAttributes) GetToolDefinitions() []LLMObsSpanToolDefinition {
	if o == nil || o.ToolDefinitions == nil {
		var ret []LLMObsSpanToolDefinition
		return ret
	}
	return o.ToolDefinitions
}

// GetToolDefinitionsOk returns a tuple with the ToolDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetToolDefinitionsOk() (*[]LLMObsSpanToolDefinition, bool) {
	if o == nil || o.ToolDefinitions == nil {
		return nil, false
	}
	return &o.ToolDefinitions, true
}

// HasToolDefinitions returns a boolean if a field has been set.
func (o *LLMObsSpanAttributes) HasToolDefinitions() bool {
	return o != nil && o.ToolDefinitions != nil
}

// SetToolDefinitions gets a reference to the given []LLMObsSpanToolDefinition and assigns it to the ToolDefinitions field.
func (o *LLMObsSpanAttributes) SetToolDefinitions(v []LLMObsSpanToolDefinition) {
	o.ToolDefinitions = v
}

// GetTraceId returns the TraceId field value.
func (o *LLMObsSpanAttributes) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpanAttributes) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value.
func (o *LLMObsSpanAttributes) SetTraceId(v string) {
	o.TraceId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSpanAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["duration"] = o.Duration
	if o.Evaluation != nil {
		toSerialize["evaluation"] = o.Evaluation
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.Intent != nil {
		toSerialize["intent"] = o.Intent
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	toSerialize["ml_app"] = o.MlApp
	if o.ModelName != nil {
		toSerialize["model_name"] = o.ModelName
	}
	if o.ModelProvider != nil {
		toSerialize["model_provider"] = o.ModelProvider
	}
	toSerialize["name"] = o.Name
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	if o.ParentId != nil {
		toSerialize["parent_id"] = o.ParentId
	}
	toSerialize["span_id"] = o.SpanId
	toSerialize["span_kind"] = o.SpanKind
	toSerialize["start_ns"] = o.StartNs
	toSerialize["status"] = o.Status
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ToolDefinitions != nil {
		toSerialize["tool_definitions"] = o.ToolDefinitions
	}
	toSerialize["trace_id"] = o.TraceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsSpanAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration        *float64                              `json:"duration"`
		Evaluation      map[string]LLMObsSpanEvaluationMetric `json:"evaluation,omitempty"`
		Input           *LLMObsSpanIO                         `json:"input,omitempty"`
		Intent          *string                               `json:"intent,omitempty"`
		Metadata        map[string]interface{}                `json:"metadata,omitempty"`
		Metrics         map[string]float64                    `json:"metrics,omitempty"`
		MlApp           *string                               `json:"ml_app"`
		ModelName       *string                               `json:"model_name,omitempty"`
		ModelProvider   *string                               `json:"model_provider,omitempty"`
		Name            *string                               `json:"name"`
		Output          *LLMObsSpanIO                         `json:"output,omitempty"`
		ParentId        *string                               `json:"parent_id,omitempty"`
		SpanId          *string                               `json:"span_id"`
		SpanKind        *string                               `json:"span_kind"`
		StartNs         *int64                                `json:"start_ns"`
		Status          *string                               `json:"status"`
		Tags            []string                              `json:"tags,omitempty"`
		ToolDefinitions []LLMObsSpanToolDefinition            `json:"tool_definitions,omitempty"`
		TraceId         *string                               `json:"trace_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Duration == nil {
		return fmt.Errorf("required field duration missing")
	}
	if all.MlApp == nil {
		return fmt.Errorf("required field ml_app missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SpanId == nil {
		return fmt.Errorf("required field span_id missing")
	}
	if all.SpanKind == nil {
		return fmt.Errorf("required field span_kind missing")
	}
	if all.StartNs == nil {
		return fmt.Errorf("required field start_ns missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.TraceId == nil {
		return fmt.Errorf("required field trace_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration", "evaluation", "input", "intent", "metadata", "metrics", "ml_app", "model_name", "model_provider", "name", "output", "parent_id", "span_id", "span_kind", "start_ns", "status", "tags", "tool_definitions", "trace_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Duration = *all.Duration
	o.Evaluation = all.Evaluation
	if all.Input != nil && all.Input.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Input = all.Input
	o.Intent = all.Intent
	o.Metadata = all.Metadata
	o.Metrics = all.Metrics
	o.MlApp = *all.MlApp
	o.ModelName = all.ModelName
	o.ModelProvider = all.ModelProvider
	o.Name = *all.Name
	if all.Output != nil && all.Output.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Output = all.Output
	o.ParentId = all.ParentId
	o.SpanId = *all.SpanId
	o.SpanKind = *all.SpanKind
	o.StartNs = *all.StartNs
	o.Status = *all.Status
	o.Tags = all.Tags
	o.ToolDefinitions = all.ToolDefinitions
	o.TraceId = *all.TraceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
