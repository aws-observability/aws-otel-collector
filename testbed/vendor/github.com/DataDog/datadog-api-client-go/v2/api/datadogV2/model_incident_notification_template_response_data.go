// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationTemplateResponseData Notification template data from a response.
type IncidentNotificationTemplateResponseData struct {
	// The notification template's attributes.
	Attributes *IncidentNotificationTemplateAttributes `json:"attributes,omitempty"`
	// The unique identifier of the notification template.
	Id uuid.UUID `json:"id"`
	// The notification template's resource relationships.
	Relationships *IncidentNotificationTemplateRelationships `json:"relationships,omitempty"`
	// Notification templates resource type.
	Type IncidentNotificationTemplateType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationTemplateResponseData instantiates a new IncidentNotificationTemplateResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationTemplateResponseData(id uuid.UUID, typeVar IncidentNotificationTemplateType) *IncidentNotificationTemplateResponseData {
	this := IncidentNotificationTemplateResponseData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentNotificationTemplateResponseDataWithDefaults instantiates a new IncidentNotificationTemplateResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationTemplateResponseDataWithDefaults() *IncidentNotificationTemplateResponseData {
	this := IncidentNotificationTemplateResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentNotificationTemplateResponseData) GetAttributes() IncidentNotificationTemplateAttributes {
	if o == nil || o.Attributes == nil {
		var ret IncidentNotificationTemplateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateResponseData) GetAttributesOk() (*IncidentNotificationTemplateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentNotificationTemplateResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given IncidentNotificationTemplateAttributes and assigns it to the Attributes field.
func (o *IncidentNotificationTemplateResponseData) SetAttributes(v IncidentNotificationTemplateAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *IncidentNotificationTemplateResponseData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateResponseData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentNotificationTemplateResponseData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentNotificationTemplateResponseData) GetRelationships() IncidentNotificationTemplateRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentNotificationTemplateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateResponseData) GetRelationshipsOk() (*IncidentNotificationTemplateRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentNotificationTemplateResponseData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given IncidentNotificationTemplateRelationships and assigns it to the Relationships field.
func (o *IncidentNotificationTemplateResponseData) SetRelationships(v IncidentNotificationTemplateRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *IncidentNotificationTemplateResponseData) GetType() IncidentNotificationTemplateType {
	if o == nil {
		var ret IncidentNotificationTemplateType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateResponseData) GetTypeOk() (*IncidentNotificationTemplateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentNotificationTemplateResponseData) SetType(v IncidentNotificationTemplateType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationTemplateResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
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
func (o *IncidentNotificationTemplateResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentNotificationTemplateAttributes    `json:"attributes,omitempty"`
		Id            *uuid.UUID                                 `json:"id"`
		Relationships *IncidentNotificationTemplateRelationships `json:"relationships,omitempty"`
		Type          *IncidentNotificationTemplateType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
