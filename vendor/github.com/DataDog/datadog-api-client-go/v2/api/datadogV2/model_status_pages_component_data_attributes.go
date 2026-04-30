// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesComponentDataAttributes The attributes of a component.
type StatusPagesComponentDataAttributes struct {
	// If the component is of type `group`, the components within the group.
	Components []StatusPagesComponentDataAttributesComponentsItems `json:"components,omitempty"`
	// Timestamp of when the component was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp of when the component was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The name of the component.
	Name *string `json:"name,omitempty"`
	// The zero-indexed position of the component.
	Position *int64 `json:"position,omitempty"`
	// The status of the component.
	Status *StatusPagesComponentDataAttributesStatus `json:"status,omitempty"`
	// The type of the component.
	Type CreateComponentRequestDataAttributesType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPagesComponentDataAttributes instantiates a new StatusPagesComponentDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPagesComponentDataAttributes(typeVar CreateComponentRequestDataAttributesType) *StatusPagesComponentDataAttributes {
	this := StatusPagesComponentDataAttributes{}
	this.Type = typeVar
	return &this
}

// NewStatusPagesComponentDataAttributesWithDefaults instantiates a new StatusPagesComponentDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPagesComponentDataAttributesWithDefaults() *StatusPagesComponentDataAttributes {
	this := StatusPagesComponentDataAttributes{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *StatusPagesComponentDataAttributes) GetComponents() []StatusPagesComponentDataAttributesComponentsItems {
	if o == nil || o.Components == nil {
		var ret []StatusPagesComponentDataAttributesComponentsItems
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataAttributes) GetComponentsOk() (*[]StatusPagesComponentDataAttributesComponentsItems, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *StatusPagesComponentDataAttributes) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []StatusPagesComponentDataAttributesComponentsItems and assigns it to the Components field.
func (o *StatusPagesComponentDataAttributes) SetComponents(v []StatusPagesComponentDataAttributesComponentsItems) {
	o.Components = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StatusPagesComponentDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StatusPagesComponentDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *StatusPagesComponentDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *StatusPagesComponentDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *StatusPagesComponentDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *StatusPagesComponentDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StatusPagesComponentDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StatusPagesComponentDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StatusPagesComponentDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *StatusPagesComponentDataAttributes) GetPosition() int64 {
	if o == nil || o.Position == nil {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataAttributes) GetPositionOk() (*int64, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *StatusPagesComponentDataAttributes) HasPosition() bool {
	return o != nil && o.Position != nil
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *StatusPagesComponentDataAttributes) SetPosition(v int64) {
	o.Position = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StatusPagesComponentDataAttributes) GetStatus() StatusPagesComponentDataAttributesStatus {
	if o == nil || o.Status == nil {
		var ret StatusPagesComponentDataAttributesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataAttributes) GetStatusOk() (*StatusPagesComponentDataAttributesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StatusPagesComponentDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given StatusPagesComponentDataAttributesStatus and assigns it to the Status field.
func (o *StatusPagesComponentDataAttributes) SetStatus(v StatusPagesComponentDataAttributesStatus) {
	o.Status = &v
}

// GetType returns the Type field value.
func (o *StatusPagesComponentDataAttributes) GetType() CreateComponentRequestDataAttributesType {
	if o == nil {
		var ret CreateComponentRequestDataAttributesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataAttributes) GetTypeOk() (*CreateComponentRequestDataAttributesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StatusPagesComponentDataAttributes) SetType(v CreateComponentRequestDataAttributesType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPagesComponentDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatusPagesComponentDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components []StatusPagesComponentDataAttributesComponentsItems `json:"components,omitempty"`
		CreatedAt  *time.Time                                          `json:"created_at,omitempty"`
		ModifiedAt *time.Time                                          `json:"modified_at,omitempty"`
		Name       *string                                             `json:"name,omitempty"`
		Position   *int64                                              `json:"position,omitempty"`
		Status     *StatusPagesComponentDataAttributesStatus           `json:"status,omitempty"`
		Type       *CreateComponentRequestDataAttributesType           `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components", "created_at", "modified_at", "name", "position", "status", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Components = all.Components
	o.CreatedAt = all.CreatedAt
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.Position = all.Position
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
