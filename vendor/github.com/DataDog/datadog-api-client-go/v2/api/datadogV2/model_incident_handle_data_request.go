// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentHandleDataRequest Data object representing an incident handle in a create or update request.
type IncidentHandleDataRequest struct {
	// Incident handle attributes for requests
	Attributes IncidentHandleAttributesRequest `json:"attributes"`
	// The ID of the incident handle (required for PUT requests)
	Id *string `json:"id,omitempty"`
	// Relationships to associate with an incident handle in a create or update request.
	Relationships NullableIncidentHandleRelationshipsRequest `json:"relationships,omitempty"`
	// Incident handle resource type
	Type IncidentHandleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentHandleDataRequest instantiates a new IncidentHandleDataRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentHandleDataRequest(attributes IncidentHandleAttributesRequest, typeVar IncidentHandleType) *IncidentHandleDataRequest {
	this := IncidentHandleDataRequest{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewIncidentHandleDataRequestWithDefaults instantiates a new IncidentHandleDataRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentHandleDataRequestWithDefaults() *IncidentHandleDataRequest {
	this := IncidentHandleDataRequest{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentHandleDataRequest) GetAttributes() IncidentHandleAttributesRequest {
	if o == nil {
		var ret IncidentHandleAttributesRequest
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleDataRequest) GetAttributesOk() (*IncidentHandleAttributesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentHandleDataRequest) SetAttributes(v IncidentHandleAttributesRequest) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IncidentHandleDataRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentHandleDataRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IncidentHandleDataRequest) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IncidentHandleDataRequest) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentHandleDataRequest) GetRelationships() IncidentHandleRelationshipsRequest {
	if o == nil || o.Relationships.Get() == nil {
		var ret IncidentHandleRelationshipsRequest
		return ret
	}
	return *o.Relationships.Get()
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentHandleDataRequest) GetRelationshipsOk() (*IncidentHandleRelationshipsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relationships.Get(), o.Relationships.IsSet()
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentHandleDataRequest) HasRelationships() bool {
	return o != nil && o.Relationships.IsSet()
}

// SetRelationships gets a reference to the given NullableIncidentHandleRelationshipsRequest and assigns it to the Relationships field.
func (o *IncidentHandleDataRequest) SetRelationships(v IncidentHandleRelationshipsRequest) {
	o.Relationships.Set(&v)
}

// SetRelationshipsNil sets the value for Relationships to be an explicit nil.
func (o *IncidentHandleDataRequest) SetRelationshipsNil() {
	o.Relationships.Set(nil)
}

// UnsetRelationships ensures that no value is present for Relationships, not even an explicit nil.
func (o *IncidentHandleDataRequest) UnsetRelationships() {
	o.Relationships.Unset()
}

// GetType returns the Type field value.
func (o *IncidentHandleDataRequest) GetType() IncidentHandleType {
	if o == nil {
		var ret IncidentHandleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleDataRequest) GetTypeOk() (*IncidentHandleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentHandleDataRequest) SetType(v IncidentHandleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentHandleDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Relationships.IsSet() {
		toSerialize["relationships"] = o.Relationships.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentHandleDataRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentHandleAttributesRequest           `json:"attributes"`
		Id            *string                                    `json:"id,omitempty"`
		Relationships NullableIncidentHandleRelationshipsRequest `json:"relationships,omitempty"`
		Type          *IncidentHandleType                        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = all.Id
	o.Relationships = all.Relationships
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
