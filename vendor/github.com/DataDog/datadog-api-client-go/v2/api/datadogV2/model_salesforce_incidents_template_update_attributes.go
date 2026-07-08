// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SalesforceIncidentsTemplateUpdateAttributes Salesforce incident template attributes for an update request.
type SalesforceIncidentsTemplateUpdateAttributes struct {
	// Long-form description body for Salesforce incidents created from this template.
	Description *string `json:"description,omitempty"`
	// Human-readable name for this incident template.
	Name *string `json:"name,omitempty"`
	// The Salesforce user ID that owns incidents created from this template.
	OwnerId *string `json:"owner_id,omitempty"`
	// Priority of the Salesforce incident created from this template.
	Priority *SalesforceIncidentsTemplatePriority `json:"priority,omitempty"`
	// The Datadog-assigned ID of the Salesforce organization this template belongs to.
	SalesforceOrgId *uuid.UUID `json:"salesforce_org_id,omitempty"`
	// Subject line for Salesforce incidents created from this template.
	Subject *string `json:"subject,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSalesforceIncidentsTemplateUpdateAttributes instantiates a new SalesforceIncidentsTemplateUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSalesforceIncidentsTemplateUpdateAttributes() *SalesforceIncidentsTemplateUpdateAttributes {
	this := SalesforceIncidentsTemplateUpdateAttributes{}
	return &this
}

// NewSalesforceIncidentsTemplateUpdateAttributesWithDefaults instantiates a new SalesforceIncidentsTemplateUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSalesforceIncidentsTemplateUpdateAttributesWithDefaults() *SalesforceIncidentsTemplateUpdateAttributes {
	this := SalesforceIncidentsTemplateUpdateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SalesforceIncidentsTemplateUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SalesforceIncidentsTemplateUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) HasOwnerId() bool {
	return o != nil && o.OwnerId != nil
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *SalesforceIncidentsTemplateUpdateAttributes) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetPriority() SalesforceIncidentsTemplatePriority {
	if o == nil || o.Priority == nil {
		var ret SalesforceIncidentsTemplatePriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetPriorityOk() (*SalesforceIncidentsTemplatePriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given SalesforceIncidentsTemplatePriority and assigns it to the Priority field.
func (o *SalesforceIncidentsTemplateUpdateAttributes) SetPriority(v SalesforceIncidentsTemplatePriority) {
	o.Priority = &v
}

// GetSalesforceOrgId returns the SalesforceOrgId field value if set, zero value otherwise.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetSalesforceOrgId() uuid.UUID {
	if o == nil || o.SalesforceOrgId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.SalesforceOrgId
}

// GetSalesforceOrgIdOk returns a tuple with the SalesforceOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetSalesforceOrgIdOk() (*uuid.UUID, bool) {
	if o == nil || o.SalesforceOrgId == nil {
		return nil, false
	}
	return o.SalesforceOrgId, true
}

// HasSalesforceOrgId returns a boolean if a field has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) HasSalesforceOrgId() bool {
	return o != nil && o.SalesforceOrgId != nil
}

// SetSalesforceOrgId gets a reference to the given uuid.UUID and assigns it to the SalesforceOrgId field.
func (o *SalesforceIncidentsTemplateUpdateAttributes) SetSalesforceOrgId(v uuid.UUID) {
	o.SalesforceOrgId = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SalesforceIncidentsTemplateUpdateAttributes) HasSubject() bool {
	return o != nil && o.Subject != nil
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SalesforceIncidentsTemplateUpdateAttributes) SetSubject(v string) {
	o.Subject = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SalesforceIncidentsTemplateUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OwnerId != nil {
		toSerialize["owner_id"] = o.OwnerId
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.SalesforceOrgId != nil {
		toSerialize["salesforce_org_id"] = o.SalesforceOrgId
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SalesforceIncidentsTemplateUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description     *string                              `json:"description,omitempty"`
		Name            *string                              `json:"name,omitempty"`
		OwnerId         *string                              `json:"owner_id,omitempty"`
		Priority        *SalesforceIncidentsTemplatePriority `json:"priority,omitempty"`
		SalesforceOrgId *uuid.UUID                           `json:"salesforce_org_id,omitempty"`
		Subject         *string                              `json:"subject,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "owner_id", "priority", "salesforce_org_id", "subject"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Name = all.Name
	o.OwnerId = all.OwnerId
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
	o.SalesforceOrgId = all.SalesforceOrgId
	o.Subject = all.Subject

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
