// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesComponentGroupAttributesComponentsItems A component within a component group.
type StatusPagesComponentGroupAttributesComponentsItems struct {
	// The ID of the grouped component.
	Id *uuid.UUID `json:"id,omitempty"`
	// The name of the grouped component.
	Name *string `json:"name,omitempty"`
	// The zero-indexed position of the grouped component. Relative to the other components in the group.
	Position *int64 `json:"position,omitempty"`
	// The status of the component.
	Status *StatusPagesComponentGroupAttributesComponentsItemsStatus `json:"status,omitempty"`
	// The type of the component.
	Type *StatusPagesComponentGroupAttributesComponentsItemsType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPagesComponentGroupAttributesComponentsItems instantiates a new StatusPagesComponentGroupAttributesComponentsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPagesComponentGroupAttributesComponentsItems() *StatusPagesComponentGroupAttributesComponentsItems {
	this := StatusPagesComponentGroupAttributesComponentsItems{}
	return &this
}

// NewStatusPagesComponentGroupAttributesComponentsItemsWithDefaults instantiates a new StatusPagesComponentGroupAttributesComponentsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPagesComponentGroupAttributesComponentsItemsWithDefaults() *StatusPagesComponentGroupAttributesComponentsItems {
	this := StatusPagesComponentGroupAttributesComponentsItems{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StatusPagesComponentGroupAttributesComponentsItems) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupAttributesComponentsItems) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StatusPagesComponentGroupAttributesComponentsItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *StatusPagesComponentGroupAttributesComponentsItems) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StatusPagesComponentGroupAttributesComponentsItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupAttributesComponentsItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StatusPagesComponentGroupAttributesComponentsItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StatusPagesComponentGroupAttributesComponentsItems) SetName(v string) {
	o.Name = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *StatusPagesComponentGroupAttributesComponentsItems) GetPosition() int64 {
	if o == nil || o.Position == nil {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupAttributesComponentsItems) GetPositionOk() (*int64, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *StatusPagesComponentGroupAttributesComponentsItems) HasPosition() bool {
	return o != nil && o.Position != nil
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *StatusPagesComponentGroupAttributesComponentsItems) SetPosition(v int64) {
	o.Position = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StatusPagesComponentGroupAttributesComponentsItems) GetStatus() StatusPagesComponentGroupAttributesComponentsItemsStatus {
	if o == nil || o.Status == nil {
		var ret StatusPagesComponentGroupAttributesComponentsItemsStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupAttributesComponentsItems) GetStatusOk() (*StatusPagesComponentGroupAttributesComponentsItemsStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StatusPagesComponentGroupAttributesComponentsItems) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given StatusPagesComponentGroupAttributesComponentsItemsStatus and assigns it to the Status field.
func (o *StatusPagesComponentGroupAttributesComponentsItems) SetStatus(v StatusPagesComponentGroupAttributesComponentsItemsStatus) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StatusPagesComponentGroupAttributesComponentsItems) GetType() StatusPagesComponentGroupAttributesComponentsItemsType {
	if o == nil || o.Type == nil {
		var ret StatusPagesComponentGroupAttributesComponentsItemsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupAttributesComponentsItems) GetTypeOk() (*StatusPagesComponentGroupAttributesComponentsItemsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StatusPagesComponentGroupAttributesComponentsItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given StatusPagesComponentGroupAttributesComponentsItemsType and assigns it to the Type field.
func (o *StatusPagesComponentGroupAttributesComponentsItems) SetType(v StatusPagesComponentGroupAttributesComponentsItemsType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPagesComponentGroupAttributesComponentsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatusPagesComponentGroupAttributesComponentsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id       *uuid.UUID                                                `json:"id,omitempty"`
		Name     *string                                                   `json:"name,omitempty"`
		Position *int64                                                    `json:"position,omitempty"`
		Status   *StatusPagesComponentGroupAttributesComponentsItemsStatus `json:"status,omitempty"`
		Type     *StatusPagesComponentGroupAttributesComponentsItemsType   `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "position", "status", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = all.Name
	o.Position = all.Position
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
