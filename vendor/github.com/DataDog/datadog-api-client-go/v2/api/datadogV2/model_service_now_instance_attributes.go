// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowInstanceAttributes Attributes of a ServiceNow instance
type ServiceNowInstanceAttributes struct {
	// The name of the ServiceNow instance
	InstanceName string `json:"instance_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceNowInstanceAttributes instantiates a new ServiceNowInstanceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceNowInstanceAttributes(instanceName string) *ServiceNowInstanceAttributes {
	this := ServiceNowInstanceAttributes{}
	this.InstanceName = instanceName
	return &this
}

// NewServiceNowInstanceAttributesWithDefaults instantiates a new ServiceNowInstanceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceNowInstanceAttributesWithDefaults() *ServiceNowInstanceAttributes {
	this := ServiceNowInstanceAttributes{}
	return &this
}

// GetInstanceName returns the InstanceName field value.
func (o *ServiceNowInstanceAttributes) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *ServiceNowInstanceAttributes) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value.
func (o *ServiceNowInstanceAttributes) SetInstanceName(v string) {
	o.InstanceName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceNowInstanceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["instance_name"] = o.InstanceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceNowInstanceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceName *string `json:"instance_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InstanceName == nil {
		return fmt.Errorf("required field instance_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"instance_name"})
	} else {
		return err
	}
	o.InstanceName = *all.InstanceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
