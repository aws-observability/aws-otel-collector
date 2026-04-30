// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeatUserDataAttributes Attributes of a user assigned to a seat, including their email, name, and assignment timestamp.
type SeatUserDataAttributes struct {
	// The date and time the seat was assigned.
	AssignedAt datadog.NullableTime `json:"assigned_at,omitempty"`
	// The email of the user.
	Email datadog.NullableString `json:"email,omitempty"`
	// The name of the user.
	Name datadog.NullableString `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeatUserDataAttributes instantiates a new SeatUserDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeatUserDataAttributes() *SeatUserDataAttributes {
	this := SeatUserDataAttributes{}
	return &this
}

// NewSeatUserDataAttributesWithDefaults instantiates a new SeatUserDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeatUserDataAttributesWithDefaults() *SeatUserDataAttributes {
	this := SeatUserDataAttributes{}
	return &this
}

// GetAssignedAt returns the AssignedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeatUserDataAttributes) GetAssignedAt() time.Time {
	if o == nil || o.AssignedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AssignedAt.Get()
}

// GetAssignedAtOk returns a tuple with the AssignedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SeatUserDataAttributes) GetAssignedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedAt.Get(), o.AssignedAt.IsSet()
}

// HasAssignedAt returns a boolean if a field has been set.
func (o *SeatUserDataAttributes) HasAssignedAt() bool {
	return o != nil && o.AssignedAt.IsSet()
}

// SetAssignedAt gets a reference to the given datadog.NullableTime and assigns it to the AssignedAt field.
func (o *SeatUserDataAttributes) SetAssignedAt(v time.Time) {
	o.AssignedAt.Set(&v)
}

// SetAssignedAtNil sets the value for AssignedAt to be an explicit nil.
func (o *SeatUserDataAttributes) SetAssignedAtNil() {
	o.AssignedAt.Set(nil)
}

// UnsetAssignedAt ensures that no value is present for AssignedAt, not even an explicit nil.
func (o *SeatUserDataAttributes) UnsetAssignedAt() {
	o.AssignedAt.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeatUserDataAttributes) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SeatUserDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *SeatUserDataAttributes) HasEmail() bool {
	return o != nil && o.Email.IsSet()
}

// SetEmail gets a reference to the given datadog.NullableString and assigns it to the Email field.
func (o *SeatUserDataAttributes) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil.
func (o *SeatUserDataAttributes) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil.
func (o *SeatUserDataAttributes) UnsetEmail() {
	o.Email.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeatUserDataAttributes) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SeatUserDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SeatUserDataAttributes) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *SeatUserDataAttributes) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *SeatUserDataAttributes) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *SeatUserDataAttributes) UnsetName() {
	o.Name.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SeatUserDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignedAt.IsSet() {
		toSerialize["assigned_at"] = o.AssignedAt.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeatUserDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignedAt datadog.NullableTime   `json:"assigned_at,omitempty"`
		Email      datadog.NullableString `json:"email,omitempty"`
		Name       datadog.NullableString `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assigned_at", "email", "name"})
	} else {
		return err
	}
	o.AssignedAt = all.AssignedAt
	o.Email = all.Email
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
