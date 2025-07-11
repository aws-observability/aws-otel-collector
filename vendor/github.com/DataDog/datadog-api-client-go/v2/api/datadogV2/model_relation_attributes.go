// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationAttributes Relation attributes.
type RelationAttributes struct {
	// Relation entity reference.
	From *RelationEntity `json:"from,omitempty"`
	// Relation entity reference.
	To *RelationEntity `json:"to,omitempty"`
	// Supported relation types.
	Type *RelationType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRelationAttributes instantiates a new RelationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelationAttributes() *RelationAttributes {
	this := RelationAttributes{}
	return &this
}

// NewRelationAttributesWithDefaults instantiates a new RelationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelationAttributesWithDefaults() *RelationAttributes {
	this := RelationAttributes{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *RelationAttributes) GetFrom() RelationEntity {
	if o == nil || o.From == nil {
		var ret RelationEntity
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationAttributes) GetFromOk() (*RelationEntity, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *RelationAttributes) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given RelationEntity and assigns it to the From field.
func (o *RelationAttributes) SetFrom(v RelationEntity) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *RelationAttributes) GetTo() RelationEntity {
	if o == nil || o.To == nil {
		var ret RelationEntity
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationAttributes) GetToOk() (*RelationEntity, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RelationAttributes) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given RelationEntity and assigns it to the To field.
func (o *RelationAttributes) SetTo(v RelationEntity) {
	o.To = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationAttributes) GetType() RelationType {
	if o == nil || o.Type == nil {
		var ret RelationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationAttributes) GetTypeOk() (*RelationType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given RelationType and assigns it to the Type field.
func (o *RelationAttributes) SetType(v RelationType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RelationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
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
func (o *RelationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From *RelationEntity `json:"from,omitempty"`
		To   *RelationEntity `json:"to,omitempty"`
		Type *RelationType   `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "to", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.From != nil && all.From.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.From = all.From
	if all.To != nil && all.To.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.To = all.To
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
