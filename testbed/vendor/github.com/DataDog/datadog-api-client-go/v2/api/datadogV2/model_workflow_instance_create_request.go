// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowInstanceCreateRequest Request used to create a workflow instance.
type WorkflowInstanceCreateRequest struct {
	// Additional information for creating a workflow instance.
	Meta *WorkflowInstanceCreateMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowInstanceCreateRequest instantiates a new WorkflowInstanceCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowInstanceCreateRequest() *WorkflowInstanceCreateRequest {
	this := WorkflowInstanceCreateRequest{}
	return &this
}

// NewWorkflowInstanceCreateRequestWithDefaults instantiates a new WorkflowInstanceCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowInstanceCreateRequestWithDefaults() *WorkflowInstanceCreateRequest {
	this := WorkflowInstanceCreateRequest{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *WorkflowInstanceCreateRequest) GetMeta() WorkflowInstanceCreateMeta {
	if o == nil || o.Meta == nil {
		var ret WorkflowInstanceCreateMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceCreateRequest) GetMetaOk() (*WorkflowInstanceCreateMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *WorkflowInstanceCreateRequest) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given WorkflowInstanceCreateMeta and assigns it to the Meta field.
func (o *WorkflowInstanceCreateRequest) SetMeta(v WorkflowInstanceCreateMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowInstanceCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowInstanceCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Meta *WorkflowInstanceCreateMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"meta"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
