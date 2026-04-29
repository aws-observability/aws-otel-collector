// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentHandleDataResponse Data object representing an incident handle in a response.
type IncidentHandleDataResponse struct {
	// Incident handle attributes for responses
	Attributes IncidentHandleAttributesResponse `json:"attributes"`
	// The ID of the incident handle
	Id string `json:"id"`
	// Relationships associated with an incident handle response, including linked users and incident type.
	Relationships NullableIncidentHandleRelationships `json:"relationships,omitempty"`
	// Incident handle resource type
	Type IncidentHandleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentHandleDataResponse instantiates a new IncidentHandleDataResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentHandleDataResponse(attributes IncidentHandleAttributesResponse, id string, typeVar IncidentHandleType) *IncidentHandleDataResponse {
	this := IncidentHandleDataResponse{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentHandleDataResponseWithDefaults instantiates a new IncidentHandleDataResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentHandleDataResponseWithDefaults() *IncidentHandleDataResponse {
	this := IncidentHandleDataResponse{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentHandleDataResponse) GetAttributes() IncidentHandleAttributesResponse {
	if o == nil {
		var ret IncidentHandleAttributesResponse
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleDataResponse) GetAttributesOk() (*IncidentHandleAttributesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentHandleDataResponse) SetAttributes(v IncidentHandleAttributesResponse) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *IncidentHandleDataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleDataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentHandleDataResponse) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentHandleDataResponse) GetRelationships() IncidentHandleRelationships {
	if o == nil || o.Relationships.Get() == nil {
		var ret IncidentHandleRelationships
		return ret
	}
	return *o.Relationships.Get()
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentHandleDataResponse) GetRelationshipsOk() (*IncidentHandleRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relationships.Get(), o.Relationships.IsSet()
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentHandleDataResponse) HasRelationships() bool {
	return o != nil && o.Relationships.IsSet()
}

// SetRelationships gets a reference to the given NullableIncidentHandleRelationships and assigns it to the Relationships field.
func (o *IncidentHandleDataResponse) SetRelationships(v IncidentHandleRelationships) {
	o.Relationships.Set(&v)
}

// SetRelationshipsNil sets the value for Relationships to be an explicit nil.
func (o *IncidentHandleDataResponse) SetRelationshipsNil() {
	o.Relationships.Set(nil)
}

// UnsetRelationships ensures that no value is present for Relationships, not even an explicit nil.
func (o *IncidentHandleDataResponse) UnsetRelationships() {
	o.Relationships.Unset()
}

// GetType returns the Type field value.
func (o *IncidentHandleDataResponse) GetType() IncidentHandleType {
	if o == nil {
		var ret IncidentHandleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleDataResponse) GetTypeOk() (*IncidentHandleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentHandleDataResponse) SetType(v IncidentHandleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentHandleDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
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
func (o *IncidentHandleDataResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentHandleAttributesResponse   `json:"attributes"`
		Id            *string                             `json:"id"`
		Relationships NullableIncidentHandleRelationships `json:"relationships,omitempty"`
		Type          *IncidentHandleType                 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
	o.Id = *all.Id
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
