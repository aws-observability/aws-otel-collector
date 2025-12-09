// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTypeRelationships The incident type's resource relationships.
type IncidentTypeRelationships struct {
	// Relationship to user.
	CreatedByUser *RelationshipToUser `json:"created_by_user,omitempty"`
	// A reference to a Google Meet Configuration resource.
	GoogleMeetConfiguration NullableGoogleMeetConfigurationReference `json:"google_meet_configuration,omitempty"`
	// Relationship to user.
	LastModifiedByUser *RelationshipToUser `json:"last_modified_by_user,omitempty"`
	// A reference to a Microsoft Teams Configuration resource.
	MicrosoftTeamsConfiguration NullableMicrosoftTeamsConfigurationReference `json:"microsoft_teams_configuration,omitempty"`
	// A reference to a Zoom configuration resource.
	ZoomConfiguration NullableZoomConfigurationReference `json:"zoom_configuration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTypeRelationships instantiates a new IncidentTypeRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTypeRelationships() *IncidentTypeRelationships {
	this := IncidentTypeRelationships{}
	return &this
}

// NewIncidentTypeRelationshipsWithDefaults instantiates a new IncidentTypeRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTypeRelationshipsWithDefaults() *IncidentTypeRelationships {
	this := IncidentTypeRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *IncidentTypeRelationships) GetCreatedByUser() RelationshipToUser {
	if o == nil || o.CreatedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeRelationships) GetCreatedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *IncidentTypeRelationships) HasCreatedByUser() bool {
	return o != nil && o.CreatedByUser != nil
}

// SetCreatedByUser gets a reference to the given RelationshipToUser and assigns it to the CreatedByUser field.
func (o *IncidentTypeRelationships) SetCreatedByUser(v RelationshipToUser) {
	o.CreatedByUser = &v
}

// GetGoogleMeetConfiguration returns the GoogleMeetConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentTypeRelationships) GetGoogleMeetConfiguration() GoogleMeetConfigurationReference {
	if o == nil || o.GoogleMeetConfiguration.Get() == nil {
		var ret GoogleMeetConfigurationReference
		return ret
	}
	return *o.GoogleMeetConfiguration.Get()
}

// GetGoogleMeetConfigurationOk returns a tuple with the GoogleMeetConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentTypeRelationships) GetGoogleMeetConfigurationOk() (*GoogleMeetConfigurationReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.GoogleMeetConfiguration.Get(), o.GoogleMeetConfiguration.IsSet()
}

// HasGoogleMeetConfiguration returns a boolean if a field has been set.
func (o *IncidentTypeRelationships) HasGoogleMeetConfiguration() bool {
	return o != nil && o.GoogleMeetConfiguration.IsSet()
}

// SetGoogleMeetConfiguration gets a reference to the given NullableGoogleMeetConfigurationReference and assigns it to the GoogleMeetConfiguration field.
func (o *IncidentTypeRelationships) SetGoogleMeetConfiguration(v GoogleMeetConfigurationReference) {
	o.GoogleMeetConfiguration.Set(&v)
}

// SetGoogleMeetConfigurationNil sets the value for GoogleMeetConfiguration to be an explicit nil.
func (o *IncidentTypeRelationships) SetGoogleMeetConfigurationNil() {
	o.GoogleMeetConfiguration.Set(nil)
}

// UnsetGoogleMeetConfiguration ensures that no value is present for GoogleMeetConfiguration, not even an explicit nil.
func (o *IncidentTypeRelationships) UnsetGoogleMeetConfiguration() {
	o.GoogleMeetConfiguration.Unset()
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *IncidentTypeRelationships) GetLastModifiedByUser() RelationshipToUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeRelationships) GetLastModifiedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *IncidentTypeRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given RelationshipToUser and assigns it to the LastModifiedByUser field.
func (o *IncidentTypeRelationships) SetLastModifiedByUser(v RelationshipToUser) {
	o.LastModifiedByUser = &v
}

// GetMicrosoftTeamsConfiguration returns the MicrosoftTeamsConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentTypeRelationships) GetMicrosoftTeamsConfiguration() MicrosoftTeamsConfigurationReference {
	if o == nil || o.MicrosoftTeamsConfiguration.Get() == nil {
		var ret MicrosoftTeamsConfigurationReference
		return ret
	}
	return *o.MicrosoftTeamsConfiguration.Get()
}

// GetMicrosoftTeamsConfigurationOk returns a tuple with the MicrosoftTeamsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentTypeRelationships) GetMicrosoftTeamsConfigurationOk() (*MicrosoftTeamsConfigurationReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.MicrosoftTeamsConfiguration.Get(), o.MicrosoftTeamsConfiguration.IsSet()
}

// HasMicrosoftTeamsConfiguration returns a boolean if a field has been set.
func (o *IncidentTypeRelationships) HasMicrosoftTeamsConfiguration() bool {
	return o != nil && o.MicrosoftTeamsConfiguration.IsSet()
}

// SetMicrosoftTeamsConfiguration gets a reference to the given NullableMicrosoftTeamsConfigurationReference and assigns it to the MicrosoftTeamsConfiguration field.
func (o *IncidentTypeRelationships) SetMicrosoftTeamsConfiguration(v MicrosoftTeamsConfigurationReference) {
	o.MicrosoftTeamsConfiguration.Set(&v)
}

// SetMicrosoftTeamsConfigurationNil sets the value for MicrosoftTeamsConfiguration to be an explicit nil.
func (o *IncidentTypeRelationships) SetMicrosoftTeamsConfigurationNil() {
	o.MicrosoftTeamsConfiguration.Set(nil)
}

// UnsetMicrosoftTeamsConfiguration ensures that no value is present for MicrosoftTeamsConfiguration, not even an explicit nil.
func (o *IncidentTypeRelationships) UnsetMicrosoftTeamsConfiguration() {
	o.MicrosoftTeamsConfiguration.Unset()
}

// GetZoomConfiguration returns the ZoomConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentTypeRelationships) GetZoomConfiguration() ZoomConfigurationReference {
	if o == nil || o.ZoomConfiguration.Get() == nil {
		var ret ZoomConfigurationReference
		return ret
	}
	return *o.ZoomConfiguration.Get()
}

// GetZoomConfigurationOk returns a tuple with the ZoomConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentTypeRelationships) GetZoomConfigurationOk() (*ZoomConfigurationReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZoomConfiguration.Get(), o.ZoomConfiguration.IsSet()
}

// HasZoomConfiguration returns a boolean if a field has been set.
func (o *IncidentTypeRelationships) HasZoomConfiguration() bool {
	return o != nil && o.ZoomConfiguration.IsSet()
}

// SetZoomConfiguration gets a reference to the given NullableZoomConfigurationReference and assigns it to the ZoomConfiguration field.
func (o *IncidentTypeRelationships) SetZoomConfiguration(v ZoomConfigurationReference) {
	o.ZoomConfiguration.Set(&v)
}

// SetZoomConfigurationNil sets the value for ZoomConfiguration to be an explicit nil.
func (o *IncidentTypeRelationships) SetZoomConfigurationNil() {
	o.ZoomConfiguration.Set(nil)
}

// UnsetZoomConfiguration ensures that no value is present for ZoomConfiguration, not even an explicit nil.
func (o *IncidentTypeRelationships) UnsetZoomConfiguration() {
	o.ZoomConfiguration.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTypeRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if o.GoogleMeetConfiguration.IsSet() {
		toSerialize["google_meet_configuration"] = o.GoogleMeetConfiguration.Get()
	}
	if o.LastModifiedByUser != nil {
		toSerialize["last_modified_by_user"] = o.LastModifiedByUser
	}
	if o.MicrosoftTeamsConfiguration.IsSet() {
		toSerialize["microsoft_teams_configuration"] = o.MicrosoftTeamsConfiguration.Get()
	}
	if o.ZoomConfiguration.IsSet() {
		toSerialize["zoom_configuration"] = o.ZoomConfiguration.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTypeRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByUser               *RelationshipToUser                          `json:"created_by_user,omitempty"`
		GoogleMeetConfiguration     NullableGoogleMeetConfigurationReference     `json:"google_meet_configuration,omitempty"`
		LastModifiedByUser          *RelationshipToUser                          `json:"last_modified_by_user,omitempty"`
		MicrosoftTeamsConfiguration NullableMicrosoftTeamsConfigurationReference `json:"microsoft_teams_configuration,omitempty"`
		ZoomConfiguration           NullableZoomConfigurationReference           `json:"zoom_configuration,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by_user", "google_meet_configuration", "last_modified_by_user", "microsoft_teams_configuration", "zoom_configuration"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedByUser != nil && all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedByUser = all.CreatedByUser
	o.GoogleMeetConfiguration = all.GoogleMeetConfiguration
	if all.LastModifiedByUser != nil && all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedByUser = all.LastModifiedByUser
	o.MicrosoftTeamsConfiguration = all.MicrosoftTeamsConfiguration
	o.ZoomConfiguration = all.ZoomConfiguration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
