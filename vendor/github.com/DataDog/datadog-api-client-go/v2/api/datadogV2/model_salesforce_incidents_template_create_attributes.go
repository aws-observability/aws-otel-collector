// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SalesforceIncidentsTemplateCreateAttributes Salesforce incident template attributes for a create request.
type SalesforceIncidentsTemplateCreateAttributes struct {
	// Long-form description body for Salesforce incidents created from this template.
	Description string `json:"description"`
	// Human-readable name for this incident template. Must be unique within your organization.
	Name string `json:"name"`
	// The Salesforce user ID that owns incidents created from this template.
	OwnerId string `json:"owner_id"`
	// Priority of the Salesforce incident created from this template.
	Priority SalesforceIncidentsTemplatePriority `json:"priority"`
	// The Datadog-assigned ID of the Salesforce organization this template belongs to.
	SalesforceOrgId uuid.UUID `json:"salesforce_org_id"`
	// Subject line for Salesforce incidents created from this template.
	Subject string `json:"subject"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSalesforceIncidentsTemplateCreateAttributes instantiates a new SalesforceIncidentsTemplateCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSalesforceIncidentsTemplateCreateAttributes(description string, name string, ownerId string, priority SalesforceIncidentsTemplatePriority, salesforceOrgId uuid.UUID, subject string) *SalesforceIncidentsTemplateCreateAttributes {
	this := SalesforceIncidentsTemplateCreateAttributes{}
	this.Description = description
	this.Name = name
	this.OwnerId = ownerId
	this.Priority = priority
	this.SalesforceOrgId = salesforceOrgId
	this.Subject = subject
	return &this
}

// NewSalesforceIncidentsTemplateCreateAttributesWithDefaults instantiates a new SalesforceIncidentsTemplateCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSalesforceIncidentsTemplateCreateAttributesWithDefaults() *SalesforceIncidentsTemplateCreateAttributes {
	this := SalesforceIncidentsTemplateCreateAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetPriority returns the Priority field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetPriority() SalesforceIncidentsTemplatePriority {
	if o == nil {
		var ret SalesforceIncidentsTemplatePriority
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetPriorityOk() (*SalesforceIncidentsTemplatePriority, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) SetPriority(v SalesforceIncidentsTemplatePriority) {
	o.Priority = v
}

// GetSalesforceOrgId returns the SalesforceOrgId field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetSalesforceOrgId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.SalesforceOrgId
}

// GetSalesforceOrgIdOk returns a tuple with the SalesforceOrgId field value
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetSalesforceOrgIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesforceOrgId, true
}

// SetSalesforceOrgId sets field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) SetSalesforceOrgId(v uuid.UUID) {
	o.SalesforceOrgId = v
}

// GetSubject returns the Subject field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsTemplateCreateAttributes) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value.
func (o *SalesforceIncidentsTemplateCreateAttributes) SetSubject(v string) {
	o.Subject = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SalesforceIncidentsTemplateCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["owner_id"] = o.OwnerId
	toSerialize["priority"] = o.Priority
	toSerialize["salesforce_org_id"] = o.SalesforceOrgId
	toSerialize["subject"] = o.Subject

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SalesforceIncidentsTemplateCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description     *string                              `json:"description"`
		Name            *string                              `json:"name"`
		OwnerId         *string                              `json:"owner_id"`
		Priority        *SalesforceIncidentsTemplatePriority `json:"priority"`
		SalesforceOrgId *uuid.UUID                           `json:"salesforce_org_id"`
		Subject         *string                              `json:"subject"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OwnerId == nil {
		return fmt.Errorf("required field owner_id missing")
	}
	if all.Priority == nil {
		return fmt.Errorf("required field priority missing")
	}
	if all.SalesforceOrgId == nil {
		return fmt.Errorf("required field salesforce_org_id missing")
	}
	if all.Subject == nil {
		return fmt.Errorf("required field subject missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "owner_id", "priority", "salesforce_org_id", "subject"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	o.Name = *all.Name
	o.OwnerId = *all.OwnerId
	if !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = *all.Priority
	}
	o.SalesforceOrgId = *all.SalesforceOrgId
	o.Subject = *all.Subject

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
