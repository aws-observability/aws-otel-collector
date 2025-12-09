// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Spec The spec defines what the workflow does.
type Spec struct {
	// A list of annotations used in the workflow. These are like sticky notes for your workflow!
	Annotations []Annotation `json:"annotations,omitempty"`
	// A list of connections or connection groups used in the workflow.
	ConnectionEnvs []ConnectionEnv `json:"connectionEnvs,omitempty"`
	// Unique identifier used to trigger workflows automatically in Datadog.
	Handle *string `json:"handle,omitempty"`
	// A list of input parameters for the workflow. These can be used as dynamic runtime values in your workflow.
	InputSchema *InputSchema `json:"inputSchema,omitempty"`
	// A list of output parameters for the workflow.
	OutputSchema *OutputSchema `json:"outputSchema,omitempty"`
	// A `Step` is a sub-component of a workflow. Each `Step` performs an action.
	Steps []Step `json:"steps,omitempty"`
	// The list of triggers that activate this workflow. At least one trigger is required, and each trigger type may appear at most once.
	Triggers []Trigger `json:"triggers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpec instantiates a new Spec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpec() *Spec {
	this := Spec{}
	return &this
}

// NewSpecWithDefaults instantiates a new Spec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpecWithDefaults() *Spec {
	this := Spec{}
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *Spec) GetAnnotations() []Annotation {
	if o == nil || o.Annotations == nil {
		var ret []Annotation
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetAnnotationsOk() (*[]Annotation, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *Spec) HasAnnotations() bool {
	return o != nil && o.Annotations != nil
}

// SetAnnotations gets a reference to the given []Annotation and assigns it to the Annotations field.
func (o *Spec) SetAnnotations(v []Annotation) {
	o.Annotations = v
}

// GetConnectionEnvs returns the ConnectionEnvs field value if set, zero value otherwise.
func (o *Spec) GetConnectionEnvs() []ConnectionEnv {
	if o == nil || o.ConnectionEnvs == nil {
		var ret []ConnectionEnv
		return ret
	}
	return o.ConnectionEnvs
}

// GetConnectionEnvsOk returns a tuple with the ConnectionEnvs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetConnectionEnvsOk() (*[]ConnectionEnv, bool) {
	if o == nil || o.ConnectionEnvs == nil {
		return nil, false
	}
	return &o.ConnectionEnvs, true
}

// HasConnectionEnvs returns a boolean if a field has been set.
func (o *Spec) HasConnectionEnvs() bool {
	return o != nil && o.ConnectionEnvs != nil
}

// SetConnectionEnvs gets a reference to the given []ConnectionEnv and assigns it to the ConnectionEnvs field.
func (o *Spec) SetConnectionEnvs(v []ConnectionEnv) {
	o.ConnectionEnvs = v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *Spec) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *Spec) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *Spec) SetHandle(v string) {
	o.Handle = &v
}

// GetInputSchema returns the InputSchema field value if set, zero value otherwise.
func (o *Spec) GetInputSchema() InputSchema {
	if o == nil || o.InputSchema == nil {
		var ret InputSchema
		return ret
	}
	return *o.InputSchema
}

// GetInputSchemaOk returns a tuple with the InputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetInputSchemaOk() (*InputSchema, bool) {
	if o == nil || o.InputSchema == nil {
		return nil, false
	}
	return o.InputSchema, true
}

// HasInputSchema returns a boolean if a field has been set.
func (o *Spec) HasInputSchema() bool {
	return o != nil && o.InputSchema != nil
}

// SetInputSchema gets a reference to the given InputSchema and assigns it to the InputSchema field.
func (o *Spec) SetInputSchema(v InputSchema) {
	o.InputSchema = &v
}

// GetOutputSchema returns the OutputSchema field value if set, zero value otherwise.
func (o *Spec) GetOutputSchema() OutputSchema {
	if o == nil || o.OutputSchema == nil {
		var ret OutputSchema
		return ret
	}
	return *o.OutputSchema
}

// GetOutputSchemaOk returns a tuple with the OutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetOutputSchemaOk() (*OutputSchema, bool) {
	if o == nil || o.OutputSchema == nil {
		return nil, false
	}
	return o.OutputSchema, true
}

// HasOutputSchema returns a boolean if a field has been set.
func (o *Spec) HasOutputSchema() bool {
	return o != nil && o.OutputSchema != nil
}

// SetOutputSchema gets a reference to the given OutputSchema and assigns it to the OutputSchema field.
func (o *Spec) SetOutputSchema(v OutputSchema) {
	o.OutputSchema = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *Spec) GetSteps() []Step {
	if o == nil || o.Steps == nil {
		var ret []Step
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetStepsOk() (*[]Step, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *Spec) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []Step and assigns it to the Steps field.
func (o *Spec) SetSteps(v []Step) {
	o.Steps = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *Spec) GetTriggers() []Trigger {
	if o == nil || o.Triggers == nil {
		var ret []Trigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetTriggersOk() (*[]Trigger, bool) {
	if o == nil || o.Triggers == nil {
		return nil, false
	}
	return &o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *Spec) HasTriggers() bool {
	return o != nil && o.Triggers != nil
}

// SetTriggers gets a reference to the given []Trigger and assigns it to the Triggers field.
func (o *Spec) SetTriggers(v []Trigger) {
	o.Triggers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Spec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.ConnectionEnvs != nil {
		toSerialize["connectionEnvs"] = o.ConnectionEnvs
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.InputSchema != nil {
		toSerialize["inputSchema"] = o.InputSchema
	}
	if o.OutputSchema != nil {
		toSerialize["outputSchema"] = o.OutputSchema
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.Triggers != nil {
		toSerialize["triggers"] = o.Triggers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Spec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Annotations    []Annotation    `json:"annotations,omitempty"`
		ConnectionEnvs []ConnectionEnv `json:"connectionEnvs,omitempty"`
		Handle         *string         `json:"handle,omitempty"`
		InputSchema    *InputSchema    `json:"inputSchema,omitempty"`
		OutputSchema   *OutputSchema   `json:"outputSchema,omitempty"`
		Steps          []Step          `json:"steps,omitempty"`
		Triggers       []Trigger       `json:"triggers,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotations", "connectionEnvs", "handle", "inputSchema", "outputSchema", "steps", "triggers"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Annotations = all.Annotations
	o.ConnectionEnvs = all.ConnectionEnvs
	o.Handle = all.Handle
	if all.InputSchema != nil && all.InputSchema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.InputSchema = all.InputSchema
	if all.OutputSchema != nil && all.OutputSchema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OutputSchema = all.OutputSchema
	o.Steps = all.Steps
	o.Triggers = all.Triggers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
