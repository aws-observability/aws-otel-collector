// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssignSeatsUserResponseData The response data object containing attributes of the seat assignment result.
type AssignSeatsUserResponseData struct {
	// Attributes of the assign seats response, including the list of users assigned and the product code.
	Attributes *AssignSeatsUserResponseDataAttributes `json:"attributes,omitempty"`
	// The ID of the assign seats user response.
	Id *string `json:"id,omitempty"`
	// Seat assignments resource type.
	Type *SeatAssignmentsDataType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssignSeatsUserResponseData instantiates a new AssignSeatsUserResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssignSeatsUserResponseData() *AssignSeatsUserResponseData {
	this := AssignSeatsUserResponseData{}
	var typeVar SeatAssignmentsDataType = SEATASSIGNMENTSDATATYPE_SEAT_ASSIGNMENTS
	this.Type = &typeVar
	return &this
}

// NewAssignSeatsUserResponseDataWithDefaults instantiates a new AssignSeatsUserResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssignSeatsUserResponseDataWithDefaults() *AssignSeatsUserResponseData {
	this := AssignSeatsUserResponseData{}
	var typeVar SeatAssignmentsDataType = SEATASSIGNMENTSDATATYPE_SEAT_ASSIGNMENTS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AssignSeatsUserResponseData) GetAttributes() AssignSeatsUserResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret AssignSeatsUserResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignSeatsUserResponseData) GetAttributesOk() (*AssignSeatsUserResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AssignSeatsUserResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given AssignSeatsUserResponseDataAttributes and assigns it to the Attributes field.
func (o *AssignSeatsUserResponseData) SetAttributes(v AssignSeatsUserResponseDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssignSeatsUserResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignSeatsUserResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssignSeatsUserResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssignSeatsUserResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssignSeatsUserResponseData) GetType() SeatAssignmentsDataType {
	if o == nil || o.Type == nil {
		var ret SeatAssignmentsDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignSeatsUserResponseData) GetTypeOk() (*SeatAssignmentsDataType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssignSeatsUserResponseData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SeatAssignmentsDataType and assigns it to the Type field.
func (o *AssignSeatsUserResponseData) SetType(v SeatAssignmentsDataType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssignSeatsUserResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *AssignSeatsUserResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AssignSeatsUserResponseDataAttributes `json:"attributes,omitempty"`
		Id         *string                                `json:"id,omitempty"`
		Type       *SeatAssignmentsDataType               `json:"type,omitempty"`
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
