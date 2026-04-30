// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowAssignmentGroupAttributes Attributes of a ServiceNow assignment group
type ServiceNowAssignmentGroupAttributes struct {
	// The name of the assignment group
	AssignmentGroupName string `json:"assignment_group_name"`
	// The system ID of the assignment group in ServiceNow
	AssignmentGroupSysId string `json:"assignment_group_sys_id"`
	// The ID of the ServiceNow instance
	InstanceId uuid.UUID `json:"instance_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceNowAssignmentGroupAttributes instantiates a new ServiceNowAssignmentGroupAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceNowAssignmentGroupAttributes(assignmentGroupName string, assignmentGroupSysId string, instanceId uuid.UUID) *ServiceNowAssignmentGroupAttributes {
	this := ServiceNowAssignmentGroupAttributes{}
	this.AssignmentGroupName = assignmentGroupName
	this.AssignmentGroupSysId = assignmentGroupSysId
	this.InstanceId = instanceId
	return &this
}

// NewServiceNowAssignmentGroupAttributesWithDefaults instantiates a new ServiceNowAssignmentGroupAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceNowAssignmentGroupAttributesWithDefaults() *ServiceNowAssignmentGroupAttributes {
	this := ServiceNowAssignmentGroupAttributes{}
	return &this
}

// GetAssignmentGroupName returns the AssignmentGroupName field value.
func (o *ServiceNowAssignmentGroupAttributes) GetAssignmentGroupName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AssignmentGroupName
}

// GetAssignmentGroupNameOk returns a tuple with the AssignmentGroupName field value
// and a boolean to check if the value has been set.
func (o *ServiceNowAssignmentGroupAttributes) GetAssignmentGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentGroupName, true
}

// SetAssignmentGroupName sets field value.
func (o *ServiceNowAssignmentGroupAttributes) SetAssignmentGroupName(v string) {
	o.AssignmentGroupName = v
}

// GetAssignmentGroupSysId returns the AssignmentGroupSysId field value.
func (o *ServiceNowAssignmentGroupAttributes) GetAssignmentGroupSysId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AssignmentGroupSysId
}

// GetAssignmentGroupSysIdOk returns a tuple with the AssignmentGroupSysId field value
// and a boolean to check if the value has been set.
func (o *ServiceNowAssignmentGroupAttributes) GetAssignmentGroupSysIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentGroupSysId, true
}

// SetAssignmentGroupSysId sets field value.
func (o *ServiceNowAssignmentGroupAttributes) SetAssignmentGroupSysId(v string) {
	o.AssignmentGroupSysId = v
}

// GetInstanceId returns the InstanceId field value.
func (o *ServiceNowAssignmentGroupAttributes) GetInstanceId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *ServiceNowAssignmentGroupAttributes) GetInstanceIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value.
func (o *ServiceNowAssignmentGroupAttributes) SetInstanceId(v uuid.UUID) {
	o.InstanceId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceNowAssignmentGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["assignment_group_name"] = o.AssignmentGroupName
	toSerialize["assignment_group_sys_id"] = o.AssignmentGroupSysId
	toSerialize["instance_id"] = o.InstanceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceNowAssignmentGroupAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentGroupName  *string    `json:"assignment_group_name"`
		AssignmentGroupSysId *string    `json:"assignment_group_sys_id"`
		InstanceId           *uuid.UUID `json:"instance_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AssignmentGroupName == nil {
		return fmt.Errorf("required field assignment_group_name missing")
	}
	if all.AssignmentGroupSysId == nil {
		return fmt.Errorf("required field assignment_group_sys_id missing")
	}
	if all.InstanceId == nil {
		return fmt.Errorf("required field instance_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment_group_name", "assignment_group_sys_id", "instance_id"})
	} else {
		return err
	}
	o.AssignmentGroupName = *all.AssignmentGroupName
	o.AssignmentGroupSysId = *all.AssignmentGroupSysId
	o.InstanceId = *all.InstanceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
