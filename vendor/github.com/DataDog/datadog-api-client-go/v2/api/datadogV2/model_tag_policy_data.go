// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyData A tag policy resource.
type TagPolicyData struct {
	// The attributes of a tag policy resource.
	Attributes TagPolicyAttributes `json:"attributes"`
	// The unique identifier of the tag policy.
	Id string `json:"id"`
	// Related resources for a tag policy. Only present when the corresponding `include` query parameter is supplied.
	Relationships *TagPolicyRelationships `json:"relationships,omitempty"`
	// JSON:API resource type for a tag policy.
	Type TagPolicyResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPolicyData instantiates a new TagPolicyData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPolicyData(attributes TagPolicyAttributes, id string, typeVar TagPolicyResourceType) *TagPolicyData {
	this := TagPolicyData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewTagPolicyDataWithDefaults instantiates a new TagPolicyData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPolicyDataWithDefaults() *TagPolicyData {
	this := TagPolicyData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *TagPolicyData) GetAttributes() TagPolicyAttributes {
	if o == nil {
		var ret TagPolicyAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TagPolicyData) GetAttributesOk() (*TagPolicyAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *TagPolicyData) SetAttributes(v TagPolicyAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *TagPolicyData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagPolicyData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TagPolicyData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TagPolicyData) GetRelationships() TagPolicyRelationships {
	if o == nil || o.Relationships == nil {
		var ret TagPolicyRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyData) GetRelationshipsOk() (*TagPolicyRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TagPolicyData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given TagPolicyRelationships and assigns it to the Relationships field.
func (o *TagPolicyData) SetRelationships(v TagPolicyRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *TagPolicyData) GetType() TagPolicyResourceType {
	if o == nil {
		var ret TagPolicyResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TagPolicyData) GetTypeOk() (*TagPolicyResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TagPolicyData) SetType(v TagPolicyResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPolicyData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagPolicyData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *TagPolicyAttributes    `json:"attributes"`
		Id            *string                 `json:"id"`
		Relationships *TagPolicyRelationships `json:"relationships,omitempty"`
		Type          *TagPolicyResourceType  `json:"type"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
