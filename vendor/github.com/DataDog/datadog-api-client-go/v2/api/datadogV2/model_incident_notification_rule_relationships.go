// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRuleRelationships The notification rule's resource relationships.
type IncidentNotificationRuleRelationships struct {
	// Relationship to user.
	CreatedByUser *RelationshipToUser `json:"created_by_user,omitempty"`
	// Relationship to an incident type.
	IncidentType *RelationshipToIncidentType `json:"incident_type,omitempty"`
	// Relationship to user.
	LastModifiedByUser *RelationshipToUser `json:"last_modified_by_user,omitempty"`
	// A relationship reference to a notification template.
	NotificationTemplate *RelationshipToIncidentNotificationTemplate `json:"notification_template,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationRuleRelationships instantiates a new IncidentNotificationRuleRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationRuleRelationships() *IncidentNotificationRuleRelationships {
	this := IncidentNotificationRuleRelationships{}
	return &this
}

// NewIncidentNotificationRuleRelationshipsWithDefaults instantiates a new IncidentNotificationRuleRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationRuleRelationshipsWithDefaults() *IncidentNotificationRuleRelationships {
	this := IncidentNotificationRuleRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *IncidentNotificationRuleRelationships) GetCreatedByUser() RelationshipToUser {
	if o == nil || o.CreatedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleRelationships) GetCreatedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *IncidentNotificationRuleRelationships) HasCreatedByUser() bool {
	return o != nil && o.CreatedByUser != nil
}

// SetCreatedByUser gets a reference to the given RelationshipToUser and assigns it to the CreatedByUser field.
func (o *IncidentNotificationRuleRelationships) SetCreatedByUser(v RelationshipToUser) {
	o.CreatedByUser = &v
}

// GetIncidentType returns the IncidentType field value if set, zero value otherwise.
func (o *IncidentNotificationRuleRelationships) GetIncidentType() RelationshipToIncidentType {
	if o == nil || o.IncidentType == nil {
		var ret RelationshipToIncidentType
		return ret
	}
	return *o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleRelationships) GetIncidentTypeOk() (*RelationshipToIncidentType, bool) {
	if o == nil || o.IncidentType == nil {
		return nil, false
	}
	return o.IncidentType, true
}

// HasIncidentType returns a boolean if a field has been set.
func (o *IncidentNotificationRuleRelationships) HasIncidentType() bool {
	return o != nil && o.IncidentType != nil
}

// SetIncidentType gets a reference to the given RelationshipToIncidentType and assigns it to the IncidentType field.
func (o *IncidentNotificationRuleRelationships) SetIncidentType(v RelationshipToIncidentType) {
	o.IncidentType = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *IncidentNotificationRuleRelationships) GetLastModifiedByUser() RelationshipToUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleRelationships) GetLastModifiedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *IncidentNotificationRuleRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given RelationshipToUser and assigns it to the LastModifiedByUser field.
func (o *IncidentNotificationRuleRelationships) SetLastModifiedByUser(v RelationshipToUser) {
	o.LastModifiedByUser = &v
}

// GetNotificationTemplate returns the NotificationTemplate field value if set, zero value otherwise.
func (o *IncidentNotificationRuleRelationships) GetNotificationTemplate() RelationshipToIncidentNotificationTemplate {
	if o == nil || o.NotificationTemplate == nil {
		var ret RelationshipToIncidentNotificationTemplate
		return ret
	}
	return *o.NotificationTemplate
}

// GetNotificationTemplateOk returns a tuple with the NotificationTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleRelationships) GetNotificationTemplateOk() (*RelationshipToIncidentNotificationTemplate, bool) {
	if o == nil || o.NotificationTemplate == nil {
		return nil, false
	}
	return o.NotificationTemplate, true
}

// HasNotificationTemplate returns a boolean if a field has been set.
func (o *IncidentNotificationRuleRelationships) HasNotificationTemplate() bool {
	return o != nil && o.NotificationTemplate != nil
}

// SetNotificationTemplate gets a reference to the given RelationshipToIncidentNotificationTemplate and assigns it to the NotificationTemplate field.
func (o *IncidentNotificationRuleRelationships) SetNotificationTemplate(v RelationshipToIncidentNotificationTemplate) {
	o.NotificationTemplate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationRuleRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if o.IncidentType != nil {
		toSerialize["incident_type"] = o.IncidentType
	}
	if o.LastModifiedByUser != nil {
		toSerialize["last_modified_by_user"] = o.LastModifiedByUser
	}
	if o.NotificationTemplate != nil {
		toSerialize["notification_template"] = o.NotificationTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNotificationRuleRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByUser        *RelationshipToUser                         `json:"created_by_user,omitempty"`
		IncidentType         *RelationshipToIncidentType                 `json:"incident_type,omitempty"`
		LastModifiedByUser   *RelationshipToUser                         `json:"last_modified_by_user,omitempty"`
		NotificationTemplate *RelationshipToIncidentNotificationTemplate `json:"notification_template,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by_user", "incident_type", "last_modified_by_user", "notification_template"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedByUser != nil && all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedByUser = all.CreatedByUser
	if all.IncidentType != nil && all.IncidentType.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentType = all.IncidentType
	if all.LastModifiedByUser != nil && all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedByUser = all.LastModifiedByUser
	if all.NotificationTemplate != nil && all.NotificationTemplate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NotificationTemplate = all.NotificationTemplate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
