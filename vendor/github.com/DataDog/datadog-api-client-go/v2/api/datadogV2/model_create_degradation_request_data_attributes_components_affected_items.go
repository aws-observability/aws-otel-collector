// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateDegradationRequestDataAttributesComponentsAffectedItems A component affected by a degradation.
type CreateDegradationRequestDataAttributesComponentsAffectedItems struct {
	// The ID of the component. Must be a component of type `component`.
	Id uuid.UUID `json:"id"`
	// The name of the component.
	Name *string `json:"name,omitempty"`
	// The status of the component.
	Status StatusPagesComponentDataAttributesStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateDegradationRequestDataAttributesComponentsAffectedItems instantiates a new CreateDegradationRequestDataAttributesComponentsAffectedItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateDegradationRequestDataAttributesComponentsAffectedItems(id uuid.UUID, status StatusPagesComponentDataAttributesStatus) *CreateDegradationRequestDataAttributesComponentsAffectedItems {
	this := CreateDegradationRequestDataAttributesComponentsAffectedItems{}
	this.Id = id
	this.Status = status
	return &this
}

// NewCreateDegradationRequestDataAttributesComponentsAffectedItemsWithDefaults instantiates a new CreateDegradationRequestDataAttributesComponentsAffectedItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateDegradationRequestDataAttributesComponentsAffectedItemsWithDefaults() *CreateDegradationRequestDataAttributesComponentsAffectedItems {
	this := CreateDegradationRequestDataAttributesComponentsAffectedItems{}
	return &this
}

// GetId returns the Id field value.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) SetId(v uuid.UUID) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) GetStatus() StatusPagesComponentDataAttributesStatus {
	if o == nil {
		var ret StatusPagesComponentDataAttributesStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) GetStatusOk() (*StatusPagesComponentDataAttributesStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) SetStatus(v StatusPagesComponentDataAttributesStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateDegradationRequestDataAttributesComponentsAffectedItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateDegradationRequestDataAttributesComponentsAffectedItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *uuid.UUID                                `json:"id"`
		Name   *string                                   `json:"name,omitempty"`
		Status *StatusPagesComponentDataAttributesStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Name = all.Name
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
