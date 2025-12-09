// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRuleCreateData Notification rule data for a create request.
type IncidentNotificationRuleCreateData struct {
	// The attributes for creating a notification rule.
	Attributes IncidentNotificationRuleCreateAttributes `json:"attributes"`
	// The definition of `NotificationRuleCreateDataRelationships` object.
	Relationships *IncidentNotificationRuleCreateDataRelationships `json:"relationships,omitempty"`
	// Notification rules resource type.
	Type IncidentNotificationRuleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationRuleCreateData instantiates a new IncidentNotificationRuleCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationRuleCreateData(attributes IncidentNotificationRuleCreateAttributes, typeVar IncidentNotificationRuleType) *IncidentNotificationRuleCreateData {
	this := IncidentNotificationRuleCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewIncidentNotificationRuleCreateDataWithDefaults instantiates a new IncidentNotificationRuleCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationRuleCreateDataWithDefaults() *IncidentNotificationRuleCreateData {
	this := IncidentNotificationRuleCreateData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentNotificationRuleCreateData) GetAttributes() IncidentNotificationRuleCreateAttributes {
	if o == nil {
		var ret IncidentNotificationRuleCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleCreateData) GetAttributesOk() (*IncidentNotificationRuleCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentNotificationRuleCreateData) SetAttributes(v IncidentNotificationRuleCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentNotificationRuleCreateData) GetRelationships() IncidentNotificationRuleCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentNotificationRuleCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleCreateData) GetRelationshipsOk() (*IncidentNotificationRuleCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentNotificationRuleCreateData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given IncidentNotificationRuleCreateDataRelationships and assigns it to the Relationships field.
func (o *IncidentNotificationRuleCreateData) SetRelationships(v IncidentNotificationRuleCreateDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *IncidentNotificationRuleCreateData) GetType() IncidentNotificationRuleType {
	if o == nil {
		var ret IncidentNotificationRuleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleCreateData) GetTypeOk() (*IncidentNotificationRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentNotificationRuleCreateData) SetType(v IncidentNotificationRuleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationRuleCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
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
func (o *IncidentNotificationRuleCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentNotificationRuleCreateAttributes        `json:"attributes"`
		Relationships *IncidentNotificationRuleCreateDataRelationships `json:"relationships,omitempty"`
		Type          *IncidentNotificationRuleType                    `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
