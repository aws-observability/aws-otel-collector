// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionQueryWithRelationships Restriction query object returned by the API.
type RestrictionQueryWithRelationships struct {
	// Attributes of the restriction query.
	Attributes *RestrictionQueryAttributes `json:"attributes,omitempty"`
	// ID of the restriction query.
	Id *string `json:"id,omitempty"`
	// Relationships of the user object.
	Relationships *UserRelationships `json:"relationships,omitempty"`
	// Restriction query resource type.
	Type *LogsRestrictionQueriesType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestrictionQueryWithRelationships instantiates a new RestrictionQueryWithRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestrictionQueryWithRelationships() *RestrictionQueryWithRelationships {
	this := RestrictionQueryWithRelationships{}
	var typeVar LogsRestrictionQueriesType = LOGSRESTRICTIONQUERIESTYPE_LOGS_RESTRICTION_QUERIES
	this.Type = &typeVar
	return &this
}

// NewRestrictionQueryWithRelationshipsWithDefaults instantiates a new RestrictionQueryWithRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestrictionQueryWithRelationshipsWithDefaults() *RestrictionQueryWithRelationships {
	this := RestrictionQueryWithRelationships{}
	var typeVar LogsRestrictionQueriesType = LOGSRESTRICTIONQUERIESTYPE_LOGS_RESTRICTION_QUERIES
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RestrictionQueryWithRelationships) GetAttributes() RestrictionQueryAttributes {
	if o == nil || o.Attributes == nil {
		var ret RestrictionQueryAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryWithRelationships) GetAttributesOk() (*RestrictionQueryAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RestrictionQueryWithRelationships) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given RestrictionQueryAttributes and assigns it to the Attributes field.
func (o *RestrictionQueryWithRelationships) SetAttributes(v RestrictionQueryAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RestrictionQueryWithRelationships) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryWithRelationships) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RestrictionQueryWithRelationships) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RestrictionQueryWithRelationships) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *RestrictionQueryWithRelationships) GetRelationships() UserRelationships {
	if o == nil || o.Relationships == nil {
		var ret UserRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryWithRelationships) GetRelationshipsOk() (*UserRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *RestrictionQueryWithRelationships) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given UserRelationships and assigns it to the Relationships field.
func (o *RestrictionQueryWithRelationships) SetRelationships(v UserRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RestrictionQueryWithRelationships) GetType() LogsRestrictionQueriesType {
	if o == nil || o.Type == nil {
		var ret LogsRestrictionQueriesType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryWithRelationships) GetTypeOk() (*LogsRestrictionQueriesType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RestrictionQueryWithRelationships) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given LogsRestrictionQueriesType and assigns it to the Type field.
func (o *RestrictionQueryWithRelationships) SetType(v LogsRestrictionQueriesType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestrictionQueryWithRelationships) MarshalJSON() ([]byte, error) {
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
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
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
func (o *RestrictionQueryWithRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *RestrictionQueryAttributes `json:"attributes,omitempty"`
		Id            *string                     `json:"id,omitempty"`
		Relationships *UserRelationships          `json:"relationships,omitempty"`
		Type          *LogsRestrictionQueriesType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
