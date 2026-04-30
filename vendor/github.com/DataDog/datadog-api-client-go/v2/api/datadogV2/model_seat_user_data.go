// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeatUserData A seat user resource object containing its ID, type, and associated attributes.
type SeatUserData struct {
	// Attributes of a user assigned to a seat, including their email, name, and assignment timestamp.
	Attributes *SeatUserDataAttributes `json:"attributes,omitempty"`
	// The ID of the seat user.
	Id datadog.NullableString `json:"id,omitempty"`
	// Seat users resource type.
	Type *SeatUserDataType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeatUserData instantiates a new SeatUserData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeatUserData() *SeatUserData {
	this := SeatUserData{}
	var typeVar SeatUserDataType = SEATUSERDATATYPE_SEAT_USERS
	this.Type = &typeVar
	return &this
}

// NewSeatUserDataWithDefaults instantiates a new SeatUserData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeatUserDataWithDefaults() *SeatUserData {
	this := SeatUserData{}
	var typeVar SeatUserDataType = SEATUSERDATATYPE_SEAT_USERS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SeatUserData) GetAttributes() SeatUserDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret SeatUserDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeatUserData) GetAttributesOk() (*SeatUserDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SeatUserData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given SeatUserDataAttributes and assigns it to the Attributes field.
func (o *SeatUserData) SetAttributes(v SeatUserDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeatUserData) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SeatUserData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SeatUserData) HasId() bool {
	return o != nil && o.Id.IsSet()
}

// SetId gets a reference to the given datadog.NullableString and assigns it to the Id field.
func (o *SeatUserData) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil.
func (o *SeatUserData) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil.
func (o *SeatUserData) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SeatUserData) GetType() SeatUserDataType {
	if o == nil || o.Type == nil {
		var ret SeatUserDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeatUserData) GetTypeOk() (*SeatUserDataType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SeatUserData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SeatUserDataType and assigns it to the Type field.
func (o *SeatUserData) SetType(v SeatUserDataType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeatUserData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
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
func (o *SeatUserData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SeatUserDataAttributes `json:"attributes,omitempty"`
		Id         datadog.NullableString  `json:"id,omitempty"`
		Type       *SeatUserDataType       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
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
