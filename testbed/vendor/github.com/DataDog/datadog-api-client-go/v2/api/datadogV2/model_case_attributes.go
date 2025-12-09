// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseAttributes Case resource attributes
type CaseAttributes struct {
	// Timestamp of when the case was archived
	ArchivedAt datadog.NullableTime `json:"archived_at,omitempty"`
	// The definition of `CaseObjectAttributes` object.
	Attributes map[string][]string `json:"attributes,omitempty"`
	// Timestamp of when the case was closed
	ClosedAt datadog.NullableTime `json:"closed_at,omitempty"`
	// Timestamp of when the case was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Case custom attributes
	CustomAttributes map[string]CustomAttributeValue `json:"custom_attributes,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// Jira issue attached to case
	JiraIssue NullableJiraIssue `json:"jira_issue,omitempty"`
	// Key
	Key *string `json:"key,omitempty"`
	// Timestamp of when the case was last modified
	ModifiedAt datadog.NullableTime `json:"modified_at,omitempty"`
	// Case priority
	Priority *CasePriority `json:"priority,omitempty"`
	// ServiceNow ticket attached to case
	ServiceNowTicket NullableServiceNowTicket `json:"service_now_ticket,omitempty"`
	// Case status
	Status *CaseStatus `json:"status,omitempty"`
	// Title
	Title *string `json:"title,omitempty"`
	// Case type
	// Deprecated
	Type *CaseType `json:"type,omitempty"`
	// Case type UUID
	TypeId *string `json:"type_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseAttributes instantiates a new CaseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseAttributes() *CaseAttributes {
	this := CaseAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// NewCaseAttributesWithDefaults instantiates a new CaseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseAttributesWithDefaults() *CaseAttributes {
	this := CaseAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseAttributes) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt.Get()
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseAttributes) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchivedAt.Get(), o.ArchivedAt.IsSet()
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *CaseAttributes) HasArchivedAt() bool {
	return o != nil && o.ArchivedAt.IsSet()
}

// SetArchivedAt gets a reference to the given datadog.NullableTime and assigns it to the ArchivedAt field.
func (o *CaseAttributes) SetArchivedAt(v time.Time) {
	o.ArchivedAt.Set(&v)
}

// SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil.
func (o *CaseAttributes) SetArchivedAtNil() {
	o.ArchivedAt.Set(nil)
}

// UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil.
func (o *CaseAttributes) UnsetArchivedAt() {
	o.ArchivedAt.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CaseAttributes) GetAttributes() map[string][]string {
	if o == nil || o.Attributes == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAttributes) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CaseAttributes) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *CaseAttributes) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseAttributes) GetClosedAt() time.Time {
	if o == nil || o.ClosedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt.Get()
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseAttributes) GetClosedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClosedAt.Get(), o.ClosedAt.IsSet()
}

// HasClosedAt returns a boolean if a field has been set.
func (o *CaseAttributes) HasClosedAt() bool {
	return o != nil && o.ClosedAt.IsSet()
}

// SetClosedAt gets a reference to the given datadog.NullableTime and assigns it to the ClosedAt field.
func (o *CaseAttributes) SetClosedAt(v time.Time) {
	o.ClosedAt.Set(&v)
}

// SetClosedAtNil sets the value for ClosedAt to be an explicit nil.
func (o *CaseAttributes) SetClosedAtNil() {
	o.ClosedAt.Set(nil)
}

// UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil.
func (o *CaseAttributes) UnsetClosedAt() {
	o.ClosedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CaseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CaseAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CaseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise.
func (o *CaseAttributes) GetCustomAttributes() map[string]CustomAttributeValue {
	if o == nil || o.CustomAttributes == nil {
		var ret map[string]CustomAttributeValue
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAttributes) GetCustomAttributesOk() (*map[string]CustomAttributeValue, bool) {
	if o == nil || o.CustomAttributes == nil {
		return nil, false
	}
	return &o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *CaseAttributes) HasCustomAttributes() bool {
	return o != nil && o.CustomAttributes != nil
}

// SetCustomAttributes gets a reference to the given map[string]CustomAttributeValue and assigns it to the CustomAttributes field.
func (o *CaseAttributes) SetCustomAttributes(v map[string]CustomAttributeValue) {
	o.CustomAttributes = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CaseAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CaseAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CaseAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetJiraIssue returns the JiraIssue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseAttributes) GetJiraIssue() JiraIssue {
	if o == nil || o.JiraIssue.Get() == nil {
		var ret JiraIssue
		return ret
	}
	return *o.JiraIssue.Get()
}

// GetJiraIssueOk returns a tuple with the JiraIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseAttributes) GetJiraIssueOk() (*JiraIssue, bool) {
	if o == nil {
		return nil, false
	}
	return o.JiraIssue.Get(), o.JiraIssue.IsSet()
}

// HasJiraIssue returns a boolean if a field has been set.
func (o *CaseAttributes) HasJiraIssue() bool {
	return o != nil && o.JiraIssue.IsSet()
}

// SetJiraIssue gets a reference to the given NullableJiraIssue and assigns it to the JiraIssue field.
func (o *CaseAttributes) SetJiraIssue(v JiraIssue) {
	o.JiraIssue.Set(&v)
}

// SetJiraIssueNil sets the value for JiraIssue to be an explicit nil.
func (o *CaseAttributes) SetJiraIssueNil() {
	o.JiraIssue.Set(nil)
}

// UnsetJiraIssue ensures that no value is present for JiraIssue, not even an explicit nil.
func (o *CaseAttributes) UnsetJiraIssue() {
	o.JiraIssue.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CaseAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CaseAttributes) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CaseAttributes) SetKey(v string) {
	o.Key = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *CaseAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt.IsSet()
}

// SetModifiedAt gets a reference to the given datadog.NullableTime and assigns it to the ModifiedAt field.
func (o *CaseAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil.
func (o *CaseAttributes) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil.
func (o *CaseAttributes) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CaseAttributes) GetPriority() CasePriority {
	if o == nil || o.Priority == nil {
		var ret CasePriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAttributes) GetPriorityOk() (*CasePriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CaseAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given CasePriority and assigns it to the Priority field.
func (o *CaseAttributes) SetPriority(v CasePriority) {
	o.Priority = &v
}

// GetServiceNowTicket returns the ServiceNowTicket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseAttributes) GetServiceNowTicket() ServiceNowTicket {
	if o == nil || o.ServiceNowTicket.Get() == nil {
		var ret ServiceNowTicket
		return ret
	}
	return *o.ServiceNowTicket.Get()
}

// GetServiceNowTicketOk returns a tuple with the ServiceNowTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseAttributes) GetServiceNowTicketOk() (*ServiceNowTicket, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceNowTicket.Get(), o.ServiceNowTicket.IsSet()
}

// HasServiceNowTicket returns a boolean if a field has been set.
func (o *CaseAttributes) HasServiceNowTicket() bool {
	return o != nil && o.ServiceNowTicket.IsSet()
}

// SetServiceNowTicket gets a reference to the given NullableServiceNowTicket and assigns it to the ServiceNowTicket field.
func (o *CaseAttributes) SetServiceNowTicket(v ServiceNowTicket) {
	o.ServiceNowTicket.Set(&v)
}

// SetServiceNowTicketNil sets the value for ServiceNowTicket to be an explicit nil.
func (o *CaseAttributes) SetServiceNowTicketNil() {
	o.ServiceNowTicket.Set(nil)
}

// UnsetServiceNowTicket ensures that no value is present for ServiceNowTicket, not even an explicit nil.
func (o *CaseAttributes) UnsetServiceNowTicket() {
	o.ServiceNowTicket.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CaseAttributes) GetStatus() CaseStatus {
	if o == nil || o.Status == nil {
		var ret CaseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAttributes) GetStatusOk() (*CaseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CaseAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given CaseStatus and assigns it to the Status field.
func (o *CaseAttributes) SetStatus(v CaseStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CaseAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CaseAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CaseAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
// Deprecated
func (o *CaseAttributes) GetType() CaseType {
	if o == nil || o.Type == nil {
		var ret CaseType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CaseAttributes) GetTypeOk() (*CaseType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CaseAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CaseType and assigns it to the Type field.
// Deprecated
func (o *CaseAttributes) SetType(v CaseType) {
	o.Type = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *CaseAttributes) GetTypeId() string {
	if o == nil || o.TypeId == nil {
		var ret string
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAttributes) GetTypeIdOk() (*string, bool) {
	if o == nil || o.TypeId == nil {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *CaseAttributes) HasTypeId() bool {
	return o != nil && o.TypeId != nil
}

// SetTypeId gets a reference to the given string and assigns it to the TypeId field.
func (o *CaseAttributes) SetTypeId(v string) {
	o.TypeId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ArchivedAt.IsSet() {
		toSerialize["archived_at"] = o.ArchivedAt.Get()
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ClosedAt.IsSet() {
		toSerialize["closed_at"] = o.ClosedAt.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CustomAttributes != nil {
		toSerialize["custom_attributes"] = o.CustomAttributes
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.JiraIssue.IsSet() {
		toSerialize["jira_issue"] = o.JiraIssue.Get()
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.ServiceNowTicket.IsSet() {
		toSerialize["service_now_ticket"] = o.ServiceNowTicket.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.TypeId != nil {
		toSerialize["type_id"] = o.TypeId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchivedAt       datadog.NullableTime            `json:"archived_at,omitempty"`
		Attributes       map[string][]string             `json:"attributes,omitempty"`
		ClosedAt         datadog.NullableTime            `json:"closed_at,omitempty"`
		CreatedAt        *time.Time                      `json:"created_at,omitempty"`
		CustomAttributes map[string]CustomAttributeValue `json:"custom_attributes,omitempty"`
		Description      *string                         `json:"description,omitempty"`
		JiraIssue        NullableJiraIssue               `json:"jira_issue,omitempty"`
		Key              *string                         `json:"key,omitempty"`
		ModifiedAt       datadog.NullableTime            `json:"modified_at,omitempty"`
		Priority         *CasePriority                   `json:"priority,omitempty"`
		ServiceNowTicket NullableServiceNowTicket        `json:"service_now_ticket,omitempty"`
		Status           *CaseStatus                     `json:"status,omitempty"`
		Title            *string                         `json:"title,omitempty"`
		Type             *CaseType                       `json:"type,omitempty"`
		TypeId           *string                         `json:"type_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archived_at", "attributes", "closed_at", "created_at", "custom_attributes", "description", "jira_issue", "key", "modified_at", "priority", "service_now_ticket", "status", "title", "type", "type_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ArchivedAt = all.ArchivedAt
	o.Attributes = all.Attributes
	o.ClosedAt = all.ClosedAt
	o.CreatedAt = all.CreatedAt
	o.CustomAttributes = all.CustomAttributes
	o.Description = all.Description
	o.JiraIssue = all.JiraIssue
	o.Key = all.Key
	o.ModifiedAt = all.ModifiedAt
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
	o.ServiceNowTicket = all.ServiceNowTicket
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Title = all.Title
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.TypeId = all.TypeId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
