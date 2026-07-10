// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssigneeResponseDataAttributes Attributes of the assignee response.
type AssigneeResponseDataAttributes struct {
	// Unique identifier of the Datadog user assigned to the security findings. Omitted when the findings were unassigned.
	AssigneeId *string `json:"assignee_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssigneeResponseDataAttributes instantiates a new AssigneeResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssigneeResponseDataAttributes() *AssigneeResponseDataAttributes {
	this := AssigneeResponseDataAttributes{}
	return &this
}

// NewAssigneeResponseDataAttributesWithDefaults instantiates a new AssigneeResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssigneeResponseDataAttributesWithDefaults() *AssigneeResponseDataAttributes {
	this := AssigneeResponseDataAttributes{}
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *AssigneeResponseDataAttributes) GetAssigneeId() string {
	if o == nil || o.AssigneeId == nil {
		var ret string
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeResponseDataAttributes) GetAssigneeIdOk() (*string, bool) {
	if o == nil || o.AssigneeId == nil {
		return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *AssigneeResponseDataAttributes) HasAssigneeId() bool {
	return o != nil && o.AssigneeId != nil
}

// SetAssigneeId gets a reference to the given string and assigns it to the AssigneeId field.
func (o *AssigneeResponseDataAttributes) SetAssigneeId(v string) {
	o.AssigneeId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssigneeResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssigneeId != nil {
		toSerialize["assignee_id"] = o.AssigneeId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AssigneeResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssigneeId *string `json:"assignee_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee_id"})
	} else {
		return err
	}
	o.AssigneeId = all.AssigneeId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
