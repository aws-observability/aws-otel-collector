// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseAssignAttributes Case assign attributes
type CaseAssignAttributes struct {
	// Assignee's UUID
	AssigneeId string `json:"assignee_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseAssignAttributes instantiates a new CaseAssignAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseAssignAttributes(assigneeId string) *CaseAssignAttributes {
	this := CaseAssignAttributes{}
	this.AssigneeId = assigneeId
	return &this
}

// NewCaseAssignAttributesWithDefaults instantiates a new CaseAssignAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseAssignAttributesWithDefaults() *CaseAssignAttributes {
	this := CaseAssignAttributes{}
	return &this
}

// GetAssigneeId returns the AssigneeId field value.
func (o *CaseAssignAttributes) GetAssigneeId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value
// and a boolean to check if the value has been set.
func (o *CaseAssignAttributes) GetAssigneeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssigneeId, true
}

// SetAssigneeId sets field value.
func (o *CaseAssignAttributes) SetAssigneeId(v string) {
	o.AssigneeId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseAssignAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["assignee_id"] = o.AssigneeId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseAssignAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssigneeId *string `json:"assignee_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AssigneeId == nil {
		return fmt.Errorf("required field assignee_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee_id"})
	} else {
		return err
	}
	o.AssigneeId = *all.AssigneeId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
