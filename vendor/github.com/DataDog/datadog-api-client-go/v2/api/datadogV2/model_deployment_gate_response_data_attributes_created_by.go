// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateResponseDataAttributesCreatedBy Information about the user who created the deployment gate.
type DeploymentGateResponseDataAttributesCreatedBy struct {
	// The handle of the user who created the deployment rule.
	Handle *string `json:"handle,omitempty"`
	// The ID of the user who created the deployment rule.
	Id string `json:"id"`
	// The name of the user who created the deployment rule.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGateResponseDataAttributesCreatedBy instantiates a new DeploymentGateResponseDataAttributesCreatedBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGateResponseDataAttributesCreatedBy(id string) *DeploymentGateResponseDataAttributesCreatedBy {
	this := DeploymentGateResponseDataAttributesCreatedBy{}
	this.Id = id
	return &this
}

// NewDeploymentGateResponseDataAttributesCreatedByWithDefaults instantiates a new DeploymentGateResponseDataAttributesCreatedBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGateResponseDataAttributesCreatedByWithDefaults() *DeploymentGateResponseDataAttributesCreatedBy {
	this := DeploymentGateResponseDataAttributesCreatedBy{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *DeploymentGateResponseDataAttributesCreatedBy) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributesCreatedBy) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *DeploymentGateResponseDataAttributesCreatedBy) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *DeploymentGateResponseDataAttributesCreatedBy) SetHandle(v string) {
	o.Handle = &v
}

// GetId returns the Id field value.
func (o *DeploymentGateResponseDataAttributesCreatedBy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributesCreatedBy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DeploymentGateResponseDataAttributesCreatedBy) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentGateResponseDataAttributesCreatedBy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributesCreatedBy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentGateResponseDataAttributesCreatedBy) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentGateResponseDataAttributesCreatedBy) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGateResponseDataAttributesCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	toSerialize["id"] = o.Id
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGateResponseDataAttributesCreatedBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle *string `json:"handle,omitempty"`
		Id     *string `json:"id"`
		Name   *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "id", "name"})
	} else {
		return err
	}
	o.Handle = all.Handle
	o.Id = *all.Id
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
