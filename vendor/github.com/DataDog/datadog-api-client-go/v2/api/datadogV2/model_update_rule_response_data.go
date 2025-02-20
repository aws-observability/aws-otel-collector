// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateRuleResponseData The data for a rule update response.
type UpdateRuleResponseData struct {
	// Details of a rule.
	Attributes *RuleAttributes `json:"attributes,omitempty"`
	// The unique ID for a scorecard rule.
	Id *string `json:"id,omitempty"`
	// Scorecard create rule response relationship.
	Relationships *RelationshipToRule `json:"relationships,omitempty"`
	// The JSON:API type for scorecard rules.
	Type *RuleType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateRuleResponseData instantiates a new UpdateRuleResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateRuleResponseData() *UpdateRuleResponseData {
	this := UpdateRuleResponseData{}
	var typeVar RuleType = RULETYPE_RULE
	this.Type = &typeVar
	return &this
}

// NewUpdateRuleResponseDataWithDefaults instantiates a new UpdateRuleResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateRuleResponseDataWithDefaults() *UpdateRuleResponseData {
	this := UpdateRuleResponseData{}
	var typeVar RuleType = RULETYPE_RULE
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateRuleResponseData) GetAttributes() RuleAttributes {
	if o == nil || o.Attributes == nil {
		var ret RuleAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRuleResponseData) GetAttributesOk() (*RuleAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateRuleResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given RuleAttributes and assigns it to the Attributes field.
func (o *UpdateRuleResponseData) SetAttributes(v RuleAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateRuleResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRuleResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateRuleResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateRuleResponseData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *UpdateRuleResponseData) GetRelationships() RelationshipToRule {
	if o == nil || o.Relationships == nil {
		var ret RelationshipToRule
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRuleResponseData) GetRelationshipsOk() (*RelationshipToRule, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *UpdateRuleResponseData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given RelationshipToRule and assigns it to the Relationships field.
func (o *UpdateRuleResponseData) SetRelationships(v RelationshipToRule) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateRuleResponseData) GetType() RuleType {
	if o == nil || o.Type == nil {
		var ret RuleType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRuleResponseData) GetTypeOk() (*RuleType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateRuleResponseData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given RuleType and assigns it to the Type field.
func (o *UpdateRuleResponseData) SetType(v RuleType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateRuleResponseData) MarshalJSON() ([]byte, error) {
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
func (o *UpdateRuleResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *RuleAttributes     `json:"attributes,omitempty"`
		Id            *string             `json:"id,omitempty"`
		Relationships *RelationshipToRule `json:"relationships,omitempty"`
		Type          *RuleType           `json:"type,omitempty"`
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
