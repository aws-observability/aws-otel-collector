// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentResponseRelationships The incident's relationships from a response.
type IncidentResponseRelationships struct {
	// A relationship reference for attachments.
	Attachments *RelationshipToIncidentAttachment `json:"attachments,omitempty"`
	// Relationship to user.
	CommanderUser NullableNullableRelationshipToUser `json:"commander_user,omitempty"`
	// Relationship to user.
	CreatedByUser *RelationshipToUser `json:"created_by_user,omitempty"`
	// Relationship to user.
	DeclaredByUser *RelationshipToUser `json:"declared_by_user,omitempty"`
	// Relationship to impacts.
	Impacts *RelationshipToIncidentImpacts `json:"impacts,omitempty"`
	// A relationship reference for multiple integration metadata objects.
	Integrations *RelationshipToIncidentIntegrationMetadatas `json:"integrations,omitempty"`
	// Relationship to user.
	LastModifiedByUser *RelationshipToUser `json:"last_modified_by_user,omitempty"`
	// Relationship to incident responders.
	Responders *RelationshipToIncidentResponders `json:"responders,omitempty"`
	// Relationship to incident user defined fields.
	UserDefinedFields *RelationshipToIncidentUserDefinedFields `json:"user_defined_fields,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentResponseRelationships instantiates a new IncidentResponseRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentResponseRelationships() *IncidentResponseRelationships {
	this := IncidentResponseRelationships{}
	return &this
}

// NewIncidentResponseRelationshipsWithDefaults instantiates a new IncidentResponseRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentResponseRelationshipsWithDefaults() *IncidentResponseRelationships {
	this := IncidentResponseRelationships{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetAttachments() RelationshipToIncidentAttachment {
	if o == nil || o.Attachments == nil {
		var ret RelationshipToIncidentAttachment
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetAttachmentsOk() (*RelationshipToIncidentAttachment, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasAttachments() bool {
	return o != nil && o.Attachments != nil
}

// SetAttachments gets a reference to the given RelationshipToIncidentAttachment and assigns it to the Attachments field.
func (o *IncidentResponseRelationships) SetAttachments(v RelationshipToIncidentAttachment) {
	o.Attachments = &v
}

// GetCommanderUser returns the CommanderUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseRelationships) GetCommanderUser() NullableRelationshipToUser {
	if o == nil || o.CommanderUser.Get() == nil {
		var ret NullableRelationshipToUser
		return ret
	}
	return *o.CommanderUser.Get()
}

// GetCommanderUserOk returns a tuple with the CommanderUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseRelationships) GetCommanderUserOk() (*NullableRelationshipToUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommanderUser.Get(), o.CommanderUser.IsSet()
}

// HasCommanderUser returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasCommanderUser() bool {
	return o != nil && o.CommanderUser.IsSet()
}

// SetCommanderUser gets a reference to the given NullableNullableRelationshipToUser and assigns it to the CommanderUser field.
func (o *IncidentResponseRelationships) SetCommanderUser(v NullableRelationshipToUser) {
	o.CommanderUser.Set(&v)
}

// SetCommanderUserNil sets the value for CommanderUser to be an explicit nil.
func (o *IncidentResponseRelationships) SetCommanderUserNil() {
	o.CommanderUser.Set(nil)
}

// UnsetCommanderUser ensures that no value is present for CommanderUser, not even an explicit nil.
func (o *IncidentResponseRelationships) UnsetCommanderUser() {
	o.CommanderUser.Unset()
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetCreatedByUser() RelationshipToUser {
	if o == nil || o.CreatedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetCreatedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasCreatedByUser() bool {
	return o != nil && o.CreatedByUser != nil
}

// SetCreatedByUser gets a reference to the given RelationshipToUser and assigns it to the CreatedByUser field.
func (o *IncidentResponseRelationships) SetCreatedByUser(v RelationshipToUser) {
	o.CreatedByUser = &v
}

// GetDeclaredByUser returns the DeclaredByUser field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetDeclaredByUser() RelationshipToUser {
	if o == nil || o.DeclaredByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.DeclaredByUser
}

// GetDeclaredByUserOk returns a tuple with the DeclaredByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetDeclaredByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.DeclaredByUser == nil {
		return nil, false
	}
	return o.DeclaredByUser, true
}

// HasDeclaredByUser returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasDeclaredByUser() bool {
	return o != nil && o.DeclaredByUser != nil
}

// SetDeclaredByUser gets a reference to the given RelationshipToUser and assigns it to the DeclaredByUser field.
func (o *IncidentResponseRelationships) SetDeclaredByUser(v RelationshipToUser) {
	o.DeclaredByUser = &v
}

// GetImpacts returns the Impacts field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetImpacts() RelationshipToIncidentImpacts {
	if o == nil || o.Impacts == nil {
		var ret RelationshipToIncidentImpacts
		return ret
	}
	return *o.Impacts
}

// GetImpactsOk returns a tuple with the Impacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetImpactsOk() (*RelationshipToIncidentImpacts, bool) {
	if o == nil || o.Impacts == nil {
		return nil, false
	}
	return o.Impacts, true
}

// HasImpacts returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasImpacts() bool {
	return o != nil && o.Impacts != nil
}

// SetImpacts gets a reference to the given RelationshipToIncidentImpacts and assigns it to the Impacts field.
func (o *IncidentResponseRelationships) SetImpacts(v RelationshipToIncidentImpacts) {
	o.Impacts = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetIntegrations() RelationshipToIncidentIntegrationMetadatas {
	if o == nil || o.Integrations == nil {
		var ret RelationshipToIncidentIntegrationMetadatas
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetIntegrationsOk() (*RelationshipToIncidentIntegrationMetadatas, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasIntegrations() bool {
	return o != nil && o.Integrations != nil
}

// SetIntegrations gets a reference to the given RelationshipToIncidentIntegrationMetadatas and assigns it to the Integrations field.
func (o *IncidentResponseRelationships) SetIntegrations(v RelationshipToIncidentIntegrationMetadatas) {
	o.Integrations = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetLastModifiedByUser() RelationshipToUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetLastModifiedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given RelationshipToUser and assigns it to the LastModifiedByUser field.
func (o *IncidentResponseRelationships) SetLastModifiedByUser(v RelationshipToUser) {
	o.LastModifiedByUser = &v
}

// GetResponders returns the Responders field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetResponders() RelationshipToIncidentResponders {
	if o == nil || o.Responders == nil {
		var ret RelationshipToIncidentResponders
		return ret
	}
	return *o.Responders
}

// GetRespondersOk returns a tuple with the Responders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetRespondersOk() (*RelationshipToIncidentResponders, bool) {
	if o == nil || o.Responders == nil {
		return nil, false
	}
	return o.Responders, true
}

// HasResponders returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasResponders() bool {
	return o != nil && o.Responders != nil
}

// SetResponders gets a reference to the given RelationshipToIncidentResponders and assigns it to the Responders field.
func (o *IncidentResponseRelationships) SetResponders(v RelationshipToIncidentResponders) {
	o.Responders = &v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *IncidentResponseRelationships) GetUserDefinedFields() RelationshipToIncidentUserDefinedFields {
	if o == nil || o.UserDefinedFields == nil {
		var ret RelationshipToIncidentUserDefinedFields
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseRelationships) GetUserDefinedFieldsOk() (*RelationshipToIncidentUserDefinedFields, bool) {
	if o == nil || o.UserDefinedFields == nil {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *IncidentResponseRelationships) HasUserDefinedFields() bool {
	return o != nil && o.UserDefinedFields != nil
}

// SetUserDefinedFields gets a reference to the given RelationshipToIncidentUserDefinedFields and assigns it to the UserDefinedFields field.
func (o *IncidentResponseRelationships) SetUserDefinedFields(v RelationshipToIncidentUserDefinedFields) {
	o.UserDefinedFields = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentResponseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.CommanderUser.IsSet() {
		toSerialize["commander_user"] = o.CommanderUser.Get()
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if o.DeclaredByUser != nil {
		toSerialize["declared_by_user"] = o.DeclaredByUser
	}
	if o.Impacts != nil {
		toSerialize["impacts"] = o.Impacts
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}
	if o.LastModifiedByUser != nil {
		toSerialize["last_modified_by_user"] = o.LastModifiedByUser
	}
	if o.Responders != nil {
		toSerialize["responders"] = o.Responders
	}
	if o.UserDefinedFields != nil {
		toSerialize["user_defined_fields"] = o.UserDefinedFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentResponseRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attachments        *RelationshipToIncidentAttachment           `json:"attachments,omitempty"`
		CommanderUser      NullableNullableRelationshipToUser          `json:"commander_user,omitempty"`
		CreatedByUser      *RelationshipToUser                         `json:"created_by_user,omitempty"`
		DeclaredByUser     *RelationshipToUser                         `json:"declared_by_user,omitempty"`
		Impacts            *RelationshipToIncidentImpacts              `json:"impacts,omitempty"`
		Integrations       *RelationshipToIncidentIntegrationMetadatas `json:"integrations,omitempty"`
		LastModifiedByUser *RelationshipToUser                         `json:"last_modified_by_user,omitempty"`
		Responders         *RelationshipToIncidentResponders           `json:"responders,omitempty"`
		UserDefinedFields  *RelationshipToIncidentUserDefinedFields    `json:"user_defined_fields,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attachments", "commander_user", "created_by_user", "declared_by_user", "impacts", "integrations", "last_modified_by_user", "responders", "user_defined_fields"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attachments != nil && all.Attachments.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attachments = all.Attachments
	o.CommanderUser = all.CommanderUser
	if all.CreatedByUser != nil && all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedByUser = all.CreatedByUser
	if all.DeclaredByUser != nil && all.DeclaredByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DeclaredByUser = all.DeclaredByUser
	if all.Impacts != nil && all.Impacts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Impacts = all.Impacts
	if all.Integrations != nil && all.Integrations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Integrations = all.Integrations
	if all.LastModifiedByUser != nil && all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedByUser = all.LastModifiedByUser
	if all.Responders != nil && all.Responders.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Responders = all.Responders
	if all.UserDefinedFields != nil && all.UserDefinedFields.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UserDefinedFields = all.UserDefinedFields

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
