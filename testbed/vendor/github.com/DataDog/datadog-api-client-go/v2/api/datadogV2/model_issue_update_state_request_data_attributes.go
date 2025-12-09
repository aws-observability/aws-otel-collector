// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueUpdateStateRequestDataAttributes Object describing an issue state update request.
type IssueUpdateStateRequestDataAttributes struct {
	// State of the issue
	State IssueState `json:"state"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueUpdateStateRequestDataAttributes instantiates a new IssueUpdateStateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueUpdateStateRequestDataAttributes(state IssueState) *IssueUpdateStateRequestDataAttributes {
	this := IssueUpdateStateRequestDataAttributes{}
	this.State = state
	return &this
}

// NewIssueUpdateStateRequestDataAttributesWithDefaults instantiates a new IssueUpdateStateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueUpdateStateRequestDataAttributesWithDefaults() *IssueUpdateStateRequestDataAttributes {
	this := IssueUpdateStateRequestDataAttributes{}
	return &this
}

// GetState returns the State field value.
func (o *IssueUpdateStateRequestDataAttributes) GetState() IssueState {
	if o == nil {
		var ret IssueState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *IssueUpdateStateRequestDataAttributes) GetStateOk() (*IssueState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *IssueUpdateStateRequestDataAttributes) SetState(v IssueState) {
	o.State = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueUpdateStateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueUpdateStateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		State *IssueState `json:"state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"state"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
