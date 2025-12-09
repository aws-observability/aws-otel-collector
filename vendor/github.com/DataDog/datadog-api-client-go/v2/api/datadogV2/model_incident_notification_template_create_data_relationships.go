// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationTemplateCreateDataRelationships The definition of `NotificationTemplateCreateDataRelationships` object.
type IncidentNotificationTemplateCreateDataRelationships struct {
	// Relationship to an incident type.
	IncidentType *RelationshipToIncidentType `json:"incident_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationTemplateCreateDataRelationships instantiates a new IncidentNotificationTemplateCreateDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationTemplateCreateDataRelationships() *IncidentNotificationTemplateCreateDataRelationships {
	this := IncidentNotificationTemplateCreateDataRelationships{}
	return &this
}

// NewIncidentNotificationTemplateCreateDataRelationshipsWithDefaults instantiates a new IncidentNotificationTemplateCreateDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationTemplateCreateDataRelationshipsWithDefaults() *IncidentNotificationTemplateCreateDataRelationships {
	this := IncidentNotificationTemplateCreateDataRelationships{}
	return &this
}

// GetIncidentType returns the IncidentType field value if set, zero value otherwise.
func (o *IncidentNotificationTemplateCreateDataRelationships) GetIncidentType() RelationshipToIncidentType {
	if o == nil || o.IncidentType == nil {
		var ret RelationshipToIncidentType
		return ret
	}
	return *o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateCreateDataRelationships) GetIncidentTypeOk() (*RelationshipToIncidentType, bool) {
	if o == nil || o.IncidentType == nil {
		return nil, false
	}
	return o.IncidentType, true
}

// HasIncidentType returns a boolean if a field has been set.
func (o *IncidentNotificationTemplateCreateDataRelationships) HasIncidentType() bool {
	return o != nil && o.IncidentType != nil
}

// SetIncidentType gets a reference to the given RelationshipToIncidentType and assigns it to the IncidentType field.
func (o *IncidentNotificationTemplateCreateDataRelationships) SetIncidentType(v RelationshipToIncidentType) {
	o.IncidentType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationTemplateCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncidentType != nil {
		toSerialize["incident_type"] = o.IncidentType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNotificationTemplateCreateDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentType *RelationshipToIncidentType `json:"incident_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incident_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IncidentType != nil && all.IncidentType.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncidentType = all.IncidentType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
