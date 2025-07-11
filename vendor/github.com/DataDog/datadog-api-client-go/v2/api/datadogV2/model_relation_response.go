// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationResponse Relation response data.
type RelationResponse struct {
	// Relation attributes.
	Attributes *RelationAttributes `json:"attributes,omitempty"`
	// Relation ID.
	Id *string `json:"id,omitempty"`
	// Relation metadata.
	Meta *RelationMeta `json:"meta,omitempty"`
	// Relation relationships.
	Relationships *RelationRelationships `json:"relationships,omitempty"`
	// Relation subtype.
	Subtype *string `json:"subtype,omitempty"`
	// Relation type.
	Type *RelationResponseType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRelationResponse instantiates a new RelationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelationResponse() *RelationResponse {
	this := RelationResponse{}
	return &this
}

// NewRelationResponseWithDefaults instantiates a new RelationResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelationResponseWithDefaults() *RelationResponse {
	this := RelationResponse{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RelationResponse) GetAttributes() RelationAttributes {
	if o == nil || o.Attributes == nil {
		var ret RelationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationResponse) GetAttributesOk() (*RelationAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RelationResponse) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given RelationAttributes and assigns it to the Attributes field.
func (o *RelationResponse) SetAttributes(v RelationAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelationResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelationResponse) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelationResponse) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RelationResponse) GetMeta() RelationMeta {
	if o == nil || o.Meta == nil {
		var ret RelationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationResponse) GetMetaOk() (*RelationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RelationResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given RelationMeta and assigns it to the Meta field.
func (o *RelationResponse) SetMeta(v RelationMeta) {
	o.Meta = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *RelationResponse) GetRelationships() RelationRelationships {
	if o == nil || o.Relationships == nil {
		var ret RelationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationResponse) GetRelationshipsOk() (*RelationRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *RelationResponse) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given RelationRelationships and assigns it to the Relationships field.
func (o *RelationResponse) SetRelationships(v RelationRelationships) {
	o.Relationships = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *RelationResponse) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationResponse) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *RelationResponse) HasSubtype() bool {
	return o != nil && o.Subtype != nil
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *RelationResponse) SetSubtype(v string) {
	o.Subtype = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationResponse) GetType() RelationResponseType {
	if o == nil || o.Type == nil {
		var ret RelationResponseType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationResponse) GetTypeOk() (*RelationResponseType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationResponse) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given RelationResponseType and assigns it to the Type field.
func (o *RelationResponse) SetType(v RelationResponseType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RelationResponse) MarshalJSON() ([]byte, error) {
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
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
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
func (o *RelationResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *RelationAttributes    `json:"attributes,omitempty"`
		Id            *string                `json:"id,omitempty"`
		Meta          *RelationMeta          `json:"meta,omitempty"`
		Relationships *RelationRelationships `json:"relationships,omitempty"`
		Subtype       *string                `json:"subtype,omitempty"`
		Type          *RelationResponseType  `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "meta", "relationships", "subtype", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
	o.Subtype = all.Subtype
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
