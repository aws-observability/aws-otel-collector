// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowTemplateCreateRequestAttributes Attributes for creating a ServiceNow template
type ServiceNowTemplateCreateRequestAttributes struct {
	// The ID of the assignment group
	AssignmentGroupId *uuid.UUID `json:"assignment_group_id,omitempty"`
	// The ID of the business service
	BusinessServiceId *uuid.UUID `json:"business_service_id,omitempty"`
	// Custom field mappings for the template
	FieldsMapping map[string]string `json:"fields_mapping,omitempty"`
	// The handle name of the template
	HandleName string `json:"handle_name"`
	// The ID of the ServiceNow instance
	InstanceId uuid.UUID `json:"instance_id"`
	// The name of the destination ServiceNow table
	ServicenowTablename string `json:"servicenow_tablename"`
	// The ID of the user
	UserId *uuid.UUID `json:"user_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceNowTemplateCreateRequestAttributes instantiates a new ServiceNowTemplateCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceNowTemplateCreateRequestAttributes(handleName string, instanceId uuid.UUID, servicenowTablename string) *ServiceNowTemplateCreateRequestAttributes {
	this := ServiceNowTemplateCreateRequestAttributes{}
	this.HandleName = handleName
	this.InstanceId = instanceId
	this.ServicenowTablename = servicenowTablename
	return &this
}

// NewServiceNowTemplateCreateRequestAttributesWithDefaults instantiates a new ServiceNowTemplateCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceNowTemplateCreateRequestAttributesWithDefaults() *ServiceNowTemplateCreateRequestAttributes {
	this := ServiceNowTemplateCreateRequestAttributes{}
	return &this
}

// GetAssignmentGroupId returns the AssignmentGroupId field value if set, zero value otherwise.
func (o *ServiceNowTemplateCreateRequestAttributes) GetAssignmentGroupId() uuid.UUID {
	if o == nil || o.AssignmentGroupId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.AssignmentGroupId
}

// GetAssignmentGroupIdOk returns a tuple with the AssignmentGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) GetAssignmentGroupIdOk() (*uuid.UUID, bool) {
	if o == nil || o.AssignmentGroupId == nil {
		return nil, false
	}
	return o.AssignmentGroupId, true
}

// HasAssignmentGroupId returns a boolean if a field has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) HasAssignmentGroupId() bool {
	return o != nil && o.AssignmentGroupId != nil
}

// SetAssignmentGroupId gets a reference to the given uuid.UUID and assigns it to the AssignmentGroupId field.
func (o *ServiceNowTemplateCreateRequestAttributes) SetAssignmentGroupId(v uuid.UUID) {
	o.AssignmentGroupId = &v
}

// GetBusinessServiceId returns the BusinessServiceId field value if set, zero value otherwise.
func (o *ServiceNowTemplateCreateRequestAttributes) GetBusinessServiceId() uuid.UUID {
	if o == nil || o.BusinessServiceId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.BusinessServiceId
}

// GetBusinessServiceIdOk returns a tuple with the BusinessServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) GetBusinessServiceIdOk() (*uuid.UUID, bool) {
	if o == nil || o.BusinessServiceId == nil {
		return nil, false
	}
	return o.BusinessServiceId, true
}

// HasBusinessServiceId returns a boolean if a field has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) HasBusinessServiceId() bool {
	return o != nil && o.BusinessServiceId != nil
}

// SetBusinessServiceId gets a reference to the given uuid.UUID and assigns it to the BusinessServiceId field.
func (o *ServiceNowTemplateCreateRequestAttributes) SetBusinessServiceId(v uuid.UUID) {
	o.BusinessServiceId = &v
}

// GetFieldsMapping returns the FieldsMapping field value if set, zero value otherwise.
func (o *ServiceNowTemplateCreateRequestAttributes) GetFieldsMapping() map[string]string {
	if o == nil || o.FieldsMapping == nil {
		var ret map[string]string
		return ret
	}
	return o.FieldsMapping
}

// GetFieldsMappingOk returns a tuple with the FieldsMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) GetFieldsMappingOk() (*map[string]string, bool) {
	if o == nil || o.FieldsMapping == nil {
		return nil, false
	}
	return &o.FieldsMapping, true
}

// HasFieldsMapping returns a boolean if a field has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) HasFieldsMapping() bool {
	return o != nil && o.FieldsMapping != nil
}

// SetFieldsMapping gets a reference to the given map[string]string and assigns it to the FieldsMapping field.
func (o *ServiceNowTemplateCreateRequestAttributes) SetFieldsMapping(v map[string]string) {
	o.FieldsMapping = v
}

// GetHandleName returns the HandleName field value.
func (o *ServiceNowTemplateCreateRequestAttributes) GetHandleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.HandleName
}

// GetHandleNameOk returns a tuple with the HandleName field value
// and a boolean to check if the value has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) GetHandleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandleName, true
}

// SetHandleName sets field value.
func (o *ServiceNowTemplateCreateRequestAttributes) SetHandleName(v string) {
	o.HandleName = v
}

// GetInstanceId returns the InstanceId field value.
func (o *ServiceNowTemplateCreateRequestAttributes) GetInstanceId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) GetInstanceIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value.
func (o *ServiceNowTemplateCreateRequestAttributes) SetInstanceId(v uuid.UUID) {
	o.InstanceId = v
}

// GetServicenowTablename returns the ServicenowTablename field value.
func (o *ServiceNowTemplateCreateRequestAttributes) GetServicenowTablename() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServicenowTablename
}

// GetServicenowTablenameOk returns a tuple with the ServicenowTablename field value
// and a boolean to check if the value has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) GetServicenowTablenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServicenowTablename, true
}

// SetServicenowTablename sets field value.
func (o *ServiceNowTemplateCreateRequestAttributes) SetServicenowTablename(v string) {
	o.ServicenowTablename = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ServiceNowTemplateCreateRequestAttributes) GetUserId() uuid.UUID {
	if o == nil || o.UserId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) GetUserIdOk() (*uuid.UUID, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ServiceNowTemplateCreateRequestAttributes) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given uuid.UUID and assigns it to the UserId field.
func (o *ServiceNowTemplateCreateRequestAttributes) SetUserId(v uuid.UUID) {
	o.UserId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceNowTemplateCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignmentGroupId != nil {
		toSerialize["assignment_group_id"] = o.AssignmentGroupId
	}
	if o.BusinessServiceId != nil {
		toSerialize["business_service_id"] = o.BusinessServiceId
	}
	if o.FieldsMapping != nil {
		toSerialize["fields_mapping"] = o.FieldsMapping
	}
	toSerialize["handle_name"] = o.HandleName
	toSerialize["instance_id"] = o.InstanceId
	toSerialize["servicenow_tablename"] = o.ServicenowTablename
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceNowTemplateCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentGroupId   *uuid.UUID        `json:"assignment_group_id,omitempty"`
		BusinessServiceId   *uuid.UUID        `json:"business_service_id,omitempty"`
		FieldsMapping       map[string]string `json:"fields_mapping,omitempty"`
		HandleName          *string           `json:"handle_name"`
		InstanceId          *uuid.UUID        `json:"instance_id"`
		ServicenowTablename *string           `json:"servicenow_tablename"`
		UserId              *uuid.UUID        `json:"user_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HandleName == nil {
		return fmt.Errorf("required field handle_name missing")
	}
	if all.InstanceId == nil {
		return fmt.Errorf("required field instance_id missing")
	}
	if all.ServicenowTablename == nil {
		return fmt.Errorf("required field servicenow_tablename missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment_group_id", "business_service_id", "fields_mapping", "handle_name", "instance_id", "servicenow_tablename", "user_id"})
	} else {
		return err
	}
	o.AssignmentGroupId = all.AssignmentGroupId
	o.BusinessServiceId = all.BusinessServiceId
	o.FieldsMapping = all.FieldsMapping
	o.HandleName = *all.HandleName
	o.InstanceId = *all.InstanceId
	o.ServicenowTablename = *all.ServicenowTablename
	o.UserId = all.UserId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
