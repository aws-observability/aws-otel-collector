// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Step A Step is a sub-component of a workflow. Each Step performs an action.
type Step struct {
	// The unique identifier of an action.
	ActionId string `json:"actionId"`
	// Used to create conditions before running subsequent actions.
	CompletionGate *CompletionGate `json:"completionGate,omitempty"`
	// The unique identifier of a connection defined in the spec.
	ConnectionLabel *string `json:"connectionLabel,omitempty"`
	// The definition of `StepDisplay` object.
	Display *StepDisplay `json:"display,omitempty"`
	// The `Step` `errorHandlers`.
	ErrorHandlers []ErrorHandler `json:"errorHandlers,omitempty"`
	// Name of the step.
	Name string `json:"name"`
	// A list of subsequent actions to run.
	OutboundEdges []OutboundEdge `json:"outboundEdges,omitempty"`
	// A list of inputs for an action.
	Parameters []Parameter `json:"parameters,omitempty"`
	// Used to merge multiple branches into a single branch.
	ReadinessGate *ReadinessGate `json:"readinessGate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStep instantiates a new Step object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStep(actionId string, name string) *Step {
	this := Step{}
	this.ActionId = actionId
	this.Name = name
	return &this
}

// NewStepWithDefaults instantiates a new Step object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStepWithDefaults() *Step {
	this := Step{}
	return &this
}

// GetActionId returns the ActionId field value.
func (o *Step) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *Step) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value.
func (o *Step) SetActionId(v string) {
	o.ActionId = v
}

// GetCompletionGate returns the CompletionGate field value if set, zero value otherwise.
func (o *Step) GetCompletionGate() CompletionGate {
	if o == nil || o.CompletionGate == nil {
		var ret CompletionGate
		return ret
	}
	return *o.CompletionGate
}

// GetCompletionGateOk returns a tuple with the CompletionGate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Step) GetCompletionGateOk() (*CompletionGate, bool) {
	if o == nil || o.CompletionGate == nil {
		return nil, false
	}
	return o.CompletionGate, true
}

// HasCompletionGate returns a boolean if a field has been set.
func (o *Step) HasCompletionGate() bool {
	return o != nil && o.CompletionGate != nil
}

// SetCompletionGate gets a reference to the given CompletionGate and assigns it to the CompletionGate field.
func (o *Step) SetCompletionGate(v CompletionGate) {
	o.CompletionGate = &v
}

// GetConnectionLabel returns the ConnectionLabel field value if set, zero value otherwise.
func (o *Step) GetConnectionLabel() string {
	if o == nil || o.ConnectionLabel == nil {
		var ret string
		return ret
	}
	return *o.ConnectionLabel
}

// GetConnectionLabelOk returns a tuple with the ConnectionLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Step) GetConnectionLabelOk() (*string, bool) {
	if o == nil || o.ConnectionLabel == nil {
		return nil, false
	}
	return o.ConnectionLabel, true
}

// HasConnectionLabel returns a boolean if a field has been set.
func (o *Step) HasConnectionLabel() bool {
	return o != nil && o.ConnectionLabel != nil
}

// SetConnectionLabel gets a reference to the given string and assigns it to the ConnectionLabel field.
func (o *Step) SetConnectionLabel(v string) {
	o.ConnectionLabel = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *Step) GetDisplay() StepDisplay {
	if o == nil || o.Display == nil {
		var ret StepDisplay
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Step) GetDisplayOk() (*StepDisplay, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *Step) HasDisplay() bool {
	return o != nil && o.Display != nil
}

// SetDisplay gets a reference to the given StepDisplay and assigns it to the Display field.
func (o *Step) SetDisplay(v StepDisplay) {
	o.Display = &v
}

// GetErrorHandlers returns the ErrorHandlers field value if set, zero value otherwise.
func (o *Step) GetErrorHandlers() []ErrorHandler {
	if o == nil || o.ErrorHandlers == nil {
		var ret []ErrorHandler
		return ret
	}
	return o.ErrorHandlers
}

// GetErrorHandlersOk returns a tuple with the ErrorHandlers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Step) GetErrorHandlersOk() (*[]ErrorHandler, bool) {
	if o == nil || o.ErrorHandlers == nil {
		return nil, false
	}
	return &o.ErrorHandlers, true
}

// HasErrorHandlers returns a boolean if a field has been set.
func (o *Step) HasErrorHandlers() bool {
	return o != nil && o.ErrorHandlers != nil
}

// SetErrorHandlers gets a reference to the given []ErrorHandler and assigns it to the ErrorHandlers field.
func (o *Step) SetErrorHandlers(v []ErrorHandler) {
	o.ErrorHandlers = v
}

// GetName returns the Name field value.
func (o *Step) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Step) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Step) SetName(v string) {
	o.Name = v
}

// GetOutboundEdges returns the OutboundEdges field value if set, zero value otherwise.
func (o *Step) GetOutboundEdges() []OutboundEdge {
	if o == nil || o.OutboundEdges == nil {
		var ret []OutboundEdge
		return ret
	}
	return o.OutboundEdges
}

// GetOutboundEdgesOk returns a tuple with the OutboundEdges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Step) GetOutboundEdgesOk() (*[]OutboundEdge, bool) {
	if o == nil || o.OutboundEdges == nil {
		return nil, false
	}
	return &o.OutboundEdges, true
}

// HasOutboundEdges returns a boolean if a field has been set.
func (o *Step) HasOutboundEdges() bool {
	return o != nil && o.OutboundEdges != nil
}

// SetOutboundEdges gets a reference to the given []OutboundEdge and assigns it to the OutboundEdges field.
func (o *Step) SetOutboundEdges(v []OutboundEdge) {
	o.OutboundEdges = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Step) GetParameters() []Parameter {
	if o == nil || o.Parameters == nil {
		var ret []Parameter
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Step) GetParametersOk() (*[]Parameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Step) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []Parameter and assigns it to the Parameters field.
func (o *Step) SetParameters(v []Parameter) {
	o.Parameters = v
}

// GetReadinessGate returns the ReadinessGate field value if set, zero value otherwise.
func (o *Step) GetReadinessGate() ReadinessGate {
	if o == nil || o.ReadinessGate == nil {
		var ret ReadinessGate
		return ret
	}
	return *o.ReadinessGate
}

// GetReadinessGateOk returns a tuple with the ReadinessGate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Step) GetReadinessGateOk() (*ReadinessGate, bool) {
	if o == nil || o.ReadinessGate == nil {
		return nil, false
	}
	return o.ReadinessGate, true
}

// HasReadinessGate returns a boolean if a field has been set.
func (o *Step) HasReadinessGate() bool {
	return o != nil && o.ReadinessGate != nil
}

// SetReadinessGate gets a reference to the given ReadinessGate and assigns it to the ReadinessGate field.
func (o *Step) SetReadinessGate(v ReadinessGate) {
	o.ReadinessGate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Step) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["actionId"] = o.ActionId
	if o.CompletionGate != nil {
		toSerialize["completionGate"] = o.CompletionGate
	}
	if o.ConnectionLabel != nil {
		toSerialize["connectionLabel"] = o.ConnectionLabel
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.ErrorHandlers != nil {
		toSerialize["errorHandlers"] = o.ErrorHandlers
	}
	toSerialize["name"] = o.Name
	if o.OutboundEdges != nil {
		toSerialize["outboundEdges"] = o.OutboundEdges
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.ReadinessGate != nil {
		toSerialize["readinessGate"] = o.ReadinessGate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Step) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionId        *string         `json:"actionId"`
		CompletionGate  *CompletionGate `json:"completionGate,omitempty"`
		ConnectionLabel *string         `json:"connectionLabel,omitempty"`
		Display         *StepDisplay    `json:"display,omitempty"`
		ErrorHandlers   []ErrorHandler  `json:"errorHandlers,omitempty"`
		Name            *string         `json:"name"`
		OutboundEdges   []OutboundEdge  `json:"outboundEdges,omitempty"`
		Parameters      []Parameter     `json:"parameters,omitempty"`
		ReadinessGate   *ReadinessGate  `json:"readinessGate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActionId == nil {
		return fmt.Errorf("required field actionId missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actionId", "completionGate", "connectionLabel", "display", "errorHandlers", "name", "outboundEdges", "parameters", "readinessGate"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ActionId = *all.ActionId
	if all.CompletionGate != nil && all.CompletionGate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CompletionGate = all.CompletionGate
	o.ConnectionLabel = all.ConnectionLabel
	if all.Display != nil && all.Display.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Display = all.Display
	o.ErrorHandlers = all.ErrorHandlers
	o.Name = *all.Name
	o.OutboundEdges = all.OutboundEdges
	o.Parameters = all.Parameters
	if all.ReadinessGate != nil && all.ReadinessGate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReadinessGate = all.ReadinessGate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
