// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowTicketCreateAttributes ServiceNow ticket creation attributes
type ServiceNowTicketCreateAttributes struct {
	// ServiceNow assignment group
	AssignmentGroup *string `json:"assignment_group,omitempty"`
	// ServiceNow instance name
	InstanceName string `json:"instance_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceNowTicketCreateAttributes instantiates a new ServiceNowTicketCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceNowTicketCreateAttributes(instanceName string) *ServiceNowTicketCreateAttributes {
	this := ServiceNowTicketCreateAttributes{}
	this.InstanceName = instanceName
	return &this
}

// NewServiceNowTicketCreateAttributesWithDefaults instantiates a new ServiceNowTicketCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceNowTicketCreateAttributesWithDefaults() *ServiceNowTicketCreateAttributes {
	this := ServiceNowTicketCreateAttributes{}
	return &this
}

// GetAssignmentGroup returns the AssignmentGroup field value if set, zero value otherwise.
func (o *ServiceNowTicketCreateAttributes) GetAssignmentGroup() string {
	if o == nil || o.AssignmentGroup == nil {
		var ret string
		return ret
	}
	return *o.AssignmentGroup
}

// GetAssignmentGroupOk returns a tuple with the AssignmentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowTicketCreateAttributes) GetAssignmentGroupOk() (*string, bool) {
	if o == nil || o.AssignmentGroup == nil {
		return nil, false
	}
	return o.AssignmentGroup, true
}

// HasAssignmentGroup returns a boolean if a field has been set.
func (o *ServiceNowTicketCreateAttributes) HasAssignmentGroup() bool {
	return o != nil && o.AssignmentGroup != nil
}

// SetAssignmentGroup gets a reference to the given string and assigns it to the AssignmentGroup field.
func (o *ServiceNowTicketCreateAttributes) SetAssignmentGroup(v string) {
	o.AssignmentGroup = &v
}

// GetInstanceName returns the InstanceName field value.
func (o *ServiceNowTicketCreateAttributes) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *ServiceNowTicketCreateAttributes) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value.
func (o *ServiceNowTicketCreateAttributes) SetInstanceName(v string) {
	o.InstanceName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceNowTicketCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignmentGroup != nil {
		toSerialize["assignment_group"] = o.AssignmentGroup
	}
	toSerialize["instance_name"] = o.InstanceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceNowTicketCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentGroup *string `json:"assignment_group,omitempty"`
		InstanceName    *string `json:"instance_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InstanceName == nil {
		return fmt.Errorf("required field instance_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment_group", "instance_name"})
	} else {
		return err
	}
	o.AssignmentGroup = all.AssignmentGroup
	o.InstanceName = *all.InstanceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
