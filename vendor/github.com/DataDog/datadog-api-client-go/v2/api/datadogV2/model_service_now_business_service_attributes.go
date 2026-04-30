// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowBusinessServiceAttributes Attributes of a ServiceNow business service
type ServiceNowBusinessServiceAttributes struct {
	// The ID of the ServiceNow instance
	InstanceId uuid.UUID `json:"instance_id"`
	// The name of the business service
	ServiceName string `json:"service_name"`
	// The system ID of the business service in ServiceNow
	ServiceSysId string `json:"service_sys_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceNowBusinessServiceAttributes instantiates a new ServiceNowBusinessServiceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceNowBusinessServiceAttributes(instanceId uuid.UUID, serviceName string, serviceSysId string) *ServiceNowBusinessServiceAttributes {
	this := ServiceNowBusinessServiceAttributes{}
	this.InstanceId = instanceId
	this.ServiceName = serviceName
	this.ServiceSysId = serviceSysId
	return &this
}

// NewServiceNowBusinessServiceAttributesWithDefaults instantiates a new ServiceNowBusinessServiceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceNowBusinessServiceAttributesWithDefaults() *ServiceNowBusinessServiceAttributes {
	this := ServiceNowBusinessServiceAttributes{}
	return &this
}

// GetInstanceId returns the InstanceId field value.
func (o *ServiceNowBusinessServiceAttributes) GetInstanceId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *ServiceNowBusinessServiceAttributes) GetInstanceIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value.
func (o *ServiceNowBusinessServiceAttributes) SetInstanceId(v uuid.UUID) {
	o.InstanceId = v
}

// GetServiceName returns the ServiceName field value.
func (o *ServiceNowBusinessServiceAttributes) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *ServiceNowBusinessServiceAttributes) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value.
func (o *ServiceNowBusinessServiceAttributes) SetServiceName(v string) {
	o.ServiceName = v
}

// GetServiceSysId returns the ServiceSysId field value.
func (o *ServiceNowBusinessServiceAttributes) GetServiceSysId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceSysId
}

// GetServiceSysIdOk returns a tuple with the ServiceSysId field value
// and a boolean to check if the value has been set.
func (o *ServiceNowBusinessServiceAttributes) GetServiceSysIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceSysId, true
}

// SetServiceSysId sets field value.
func (o *ServiceNowBusinessServiceAttributes) SetServiceSysId(v string) {
	o.ServiceSysId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceNowBusinessServiceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["instance_id"] = o.InstanceId
	toSerialize["service_name"] = o.ServiceName
	toSerialize["service_sys_id"] = o.ServiceSysId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceNowBusinessServiceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceId   *uuid.UUID `json:"instance_id"`
		ServiceName  *string    `json:"service_name"`
		ServiceSysId *string    `json:"service_sys_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InstanceId == nil {
		return fmt.Errorf("required field instance_id missing")
	}
	if all.ServiceName == nil {
		return fmt.Errorf("required field service_name missing")
	}
	if all.ServiceSysId == nil {
		return fmt.Errorf("required field service_sys_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"instance_id", "service_name", "service_sys_id"})
	} else {
		return err
	}
	o.InstanceId = *all.InstanceId
	o.ServiceName = *all.ServiceName
	o.ServiceSysId = *all.ServiceSysId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
