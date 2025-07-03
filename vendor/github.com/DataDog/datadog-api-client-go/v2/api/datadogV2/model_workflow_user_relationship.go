// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowUserRelationship The definition of `WorkflowUserRelationship` object.
type WorkflowUserRelationship struct {
	// The definition of `WorkflowUserRelationshipData` object.
	Data *WorkflowUserRelationshipData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowUserRelationship instantiates a new WorkflowUserRelationship object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowUserRelationship() *WorkflowUserRelationship {
	this := WorkflowUserRelationship{}
	return &this
}

// NewWorkflowUserRelationshipWithDefaults instantiates a new WorkflowUserRelationship object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowUserRelationshipWithDefaults() *WorkflowUserRelationship {
	this := WorkflowUserRelationship{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WorkflowUserRelationship) GetData() WorkflowUserRelationshipData {
	if o == nil || o.Data == nil {
		var ret WorkflowUserRelationshipData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowUserRelationship) GetDataOk() (*WorkflowUserRelationshipData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WorkflowUserRelationship) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given WorkflowUserRelationshipData and assigns it to the Data field.
func (o *WorkflowUserRelationship) SetData(v WorkflowUserRelationshipData) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowUserRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowUserRelationship) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *WorkflowUserRelationshipData `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
