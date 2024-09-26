// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowInstanceCreateMeta Additional information for creating a workflow instance.
type WorkflowInstanceCreateMeta struct {
	// The input parameters to the workflow.
	Payload map[string]interface{} `json:"payload,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowInstanceCreateMeta instantiates a new WorkflowInstanceCreateMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowInstanceCreateMeta() *WorkflowInstanceCreateMeta {
	this := WorkflowInstanceCreateMeta{}
	return &this
}

// NewWorkflowInstanceCreateMetaWithDefaults instantiates a new WorkflowInstanceCreateMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowInstanceCreateMetaWithDefaults() *WorkflowInstanceCreateMeta {
	this := WorkflowInstanceCreateMeta{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *WorkflowInstanceCreateMeta) GetPayload() map[string]interface{} {
	if o == nil || o.Payload == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceCreateMeta) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *WorkflowInstanceCreateMeta) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *WorkflowInstanceCreateMeta) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowInstanceCreateMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowInstanceCreateMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Payload map[string]interface{} `json:"payload,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"payload"})
	} else {
		return err
	}
	o.Payload = all.Payload

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
