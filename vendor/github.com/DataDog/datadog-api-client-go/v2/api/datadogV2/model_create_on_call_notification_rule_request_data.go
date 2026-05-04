// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateOnCallNotificationRuleRequestData Data for creating an on-call notification rule
type CreateOnCallNotificationRuleRequestData struct {
	// Attributes for creating or modifying an on-call notification rule.
	Attributes *OnCallNotificationRuleRequestAttributes `json:"attributes,omitempty"`
	// Relationship object for creating a notification rule
	Relationships *OnCallNotificationRuleRelationships `json:"relationships,omitempty"`
	// Indicates that the resource is of type 'notification_rules'.
	Type OnCallNotificationRuleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateOnCallNotificationRuleRequestData instantiates a new CreateOnCallNotificationRuleRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateOnCallNotificationRuleRequestData(typeVar OnCallNotificationRuleType) *CreateOnCallNotificationRuleRequestData {
	this := CreateOnCallNotificationRuleRequestData{}
	this.Type = typeVar
	return &this
}

// NewCreateOnCallNotificationRuleRequestDataWithDefaults instantiates a new CreateOnCallNotificationRuleRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateOnCallNotificationRuleRequestDataWithDefaults() *CreateOnCallNotificationRuleRequestData {
	this := CreateOnCallNotificationRuleRequestData{}
	var typeVar OnCallNotificationRuleType = ONCALLNOTIFICATIONRULETYPE_NOTIFICATION_RULES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateOnCallNotificationRuleRequestData) GetAttributes() OnCallNotificationRuleRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret OnCallNotificationRuleRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOnCallNotificationRuleRequestData) GetAttributesOk() (*OnCallNotificationRuleRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateOnCallNotificationRuleRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given OnCallNotificationRuleRequestAttributes and assigns it to the Attributes field.
func (o *CreateOnCallNotificationRuleRequestData) SetAttributes(v OnCallNotificationRuleRequestAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CreateOnCallNotificationRuleRequestData) GetRelationships() OnCallNotificationRuleRelationships {
	if o == nil || o.Relationships == nil {
		var ret OnCallNotificationRuleRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOnCallNotificationRuleRequestData) GetRelationshipsOk() (*OnCallNotificationRuleRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CreateOnCallNotificationRuleRequestData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given OnCallNotificationRuleRelationships and assigns it to the Relationships field.
func (o *CreateOnCallNotificationRuleRequestData) SetRelationships(v OnCallNotificationRuleRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *CreateOnCallNotificationRuleRequestData) GetType() OnCallNotificationRuleType {
	if o == nil {
		var ret OnCallNotificationRuleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateOnCallNotificationRuleRequestData) GetTypeOk() (*OnCallNotificationRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateOnCallNotificationRuleRequestData) SetType(v OnCallNotificationRuleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateOnCallNotificationRuleRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
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
func (o *CreateOnCallNotificationRuleRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *OnCallNotificationRuleRequestAttributes `json:"attributes,omitempty"`
		Relationships *OnCallNotificationRuleRelationships     `json:"relationships,omitempty"`
		Type          *OnCallNotificationRuleType              `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
