// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingCaseResponseDataAttributes Attributes of the case.
type FindingCaseResponseDataAttributes struct {
	// Timestamp of when the case was archived.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// Relationship to user.
	AssignedTo *RelationshipToUser `json:"assigned_to,omitempty"`
	//
	Attributes map[string][]string `json:"attributes,omitempty"`
	// Timestamp of when the case was closed.
	ClosedAt *time.Time `json:"closed_at,omitempty"`
	// Timestamp of when the case was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Source of the case creation.
	CreationSource *string `json:"creation_source,omitempty"`
	// Description of the case.
	Description *string `json:"description,omitempty"`
	// Due date of the case.
	DueDate *string `json:"due_date,omitempty"`
	// Insights of the case.
	Insights []CaseInsightsItems `json:"insights,omitempty"`
	// Jira issue associated with the case.
	JiraIssue *FindingJiraIssue `json:"jira_issue,omitempty"`
	// Key of the case.
	Key *string `json:"key,omitempty"`
	// Timestamp of when the case was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Priority of the case.
	Priority *string `json:"priority,omitempty"`
	// Status of the case.
	Status *string `json:"status,omitempty"`
	// Status group of the case.
	StatusGroup *string `json:"status_group,omitempty"`
	// Status name of the case.
	StatusName *string `json:"status_name,omitempty"`
	// Title of the case.
	Title *string `json:"title,omitempty"`
	// Type of the case. For security cases, this is always "SECURITY".
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFindingCaseResponseDataAttributes instantiates a new FindingCaseResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFindingCaseResponseDataAttributes() *FindingCaseResponseDataAttributes {
	this := FindingCaseResponseDataAttributes{}
	return &this
}

// NewFindingCaseResponseDataAttributesWithDefaults instantiates a new FindingCaseResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFindingCaseResponseDataAttributesWithDefaults() *FindingCaseResponseDataAttributes {
	this := FindingCaseResponseDataAttributes{}
	return &this
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasArchivedAt() bool {
	return o != nil && o.ArchivedAt != nil
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *FindingCaseResponseDataAttributes) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetAssignedTo returns the AssignedTo field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetAssignedTo() RelationshipToUser {
	if o == nil || o.AssignedTo == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.AssignedTo
}

// GetAssignedToOk returns a tuple with the AssignedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetAssignedToOk() (*RelationshipToUser, bool) {
	if o == nil || o.AssignedTo == nil {
		return nil, false
	}
	return o.AssignedTo, true
}

// HasAssignedTo returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasAssignedTo() bool {
	return o != nil && o.AssignedTo != nil
}

// SetAssignedTo gets a reference to the given RelationshipToUser and assigns it to the AssignedTo field.
func (o *FindingCaseResponseDataAttributes) SetAssignedTo(v RelationshipToUser) {
	o.AssignedTo = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetAttributes() map[string][]string {
	if o == nil || o.Attributes == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *FindingCaseResponseDataAttributes) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetClosedAt() time.Time {
	if o == nil || o.ClosedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetClosedAtOk() (*time.Time, bool) {
	if o == nil || o.ClosedAt == nil {
		return nil, false
	}
	return o.ClosedAt, true
}

// HasClosedAt returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasClosedAt() bool {
	return o != nil && o.ClosedAt != nil
}

// SetClosedAt gets a reference to the given time.Time and assigns it to the ClosedAt field.
func (o *FindingCaseResponseDataAttributes) SetClosedAt(v time.Time) {
	o.ClosedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FindingCaseResponseDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreationSource returns the CreationSource field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetCreationSource() string {
	if o == nil || o.CreationSource == nil {
		var ret string
		return ret
	}
	return *o.CreationSource
}

// GetCreationSourceOk returns a tuple with the CreationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetCreationSourceOk() (*string, bool) {
	if o == nil || o.CreationSource == nil {
		return nil, false
	}
	return o.CreationSource, true
}

// HasCreationSource returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasCreationSource() bool {
	return o != nil && o.CreationSource != nil
}

// SetCreationSource gets a reference to the given string and assigns it to the CreationSource field.
func (o *FindingCaseResponseDataAttributes) SetCreationSource(v string) {
	o.CreationSource = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FindingCaseResponseDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetDueDate() string {
	if o == nil || o.DueDate == nil {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetDueDateOk() (*string, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasDueDate() bool {
	return o != nil && o.DueDate != nil
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *FindingCaseResponseDataAttributes) SetDueDate(v string) {
	o.DueDate = &v
}

// GetInsights returns the Insights field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetInsights() []CaseInsightsItems {
	if o == nil || o.Insights == nil {
		var ret []CaseInsightsItems
		return ret
	}
	return o.Insights
}

// GetInsightsOk returns a tuple with the Insights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetInsightsOk() (*[]CaseInsightsItems, bool) {
	if o == nil || o.Insights == nil {
		return nil, false
	}
	return &o.Insights, true
}

// HasInsights returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasInsights() bool {
	return o != nil && o.Insights != nil
}

// SetInsights gets a reference to the given []CaseInsightsItems and assigns it to the Insights field.
func (o *FindingCaseResponseDataAttributes) SetInsights(v []CaseInsightsItems) {
	o.Insights = v
}

// GetJiraIssue returns the JiraIssue field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetJiraIssue() FindingJiraIssue {
	if o == nil || o.JiraIssue == nil {
		var ret FindingJiraIssue
		return ret
	}
	return *o.JiraIssue
}

// GetJiraIssueOk returns a tuple with the JiraIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetJiraIssueOk() (*FindingJiraIssue, bool) {
	if o == nil || o.JiraIssue == nil {
		return nil, false
	}
	return o.JiraIssue, true
}

// HasJiraIssue returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasJiraIssue() bool {
	return o != nil && o.JiraIssue != nil
}

// SetJiraIssue gets a reference to the given FindingJiraIssue and assigns it to the JiraIssue field.
func (o *FindingCaseResponseDataAttributes) SetJiraIssue(v FindingJiraIssue) {
	o.JiraIssue = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FindingCaseResponseDataAttributes) SetKey(v string) {
	o.Key = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *FindingCaseResponseDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *FindingCaseResponseDataAttributes) SetPriority(v string) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FindingCaseResponseDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetStatusGroup returns the StatusGroup field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetStatusGroup() string {
	if o == nil || o.StatusGroup == nil {
		var ret string
		return ret
	}
	return *o.StatusGroup
}

// GetStatusGroupOk returns a tuple with the StatusGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetStatusGroupOk() (*string, bool) {
	if o == nil || o.StatusGroup == nil {
		return nil, false
	}
	return o.StatusGroup, true
}

// HasStatusGroup returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasStatusGroup() bool {
	return o != nil && o.StatusGroup != nil
}

// SetStatusGroup gets a reference to the given string and assigns it to the StatusGroup field.
func (o *FindingCaseResponseDataAttributes) SetStatusGroup(v string) {
	o.StatusGroup = &v
}

// GetStatusName returns the StatusName field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetStatusName() string {
	if o == nil || o.StatusName == nil {
		var ret string
		return ret
	}
	return *o.StatusName
}

// GetStatusNameOk returns a tuple with the StatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetStatusNameOk() (*string, bool) {
	if o == nil || o.StatusName == nil {
		return nil, false
	}
	return o.StatusName, true
}

// HasStatusName returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasStatusName() bool {
	return o != nil && o.StatusName != nil
}

// SetStatusName gets a reference to the given string and assigns it to the StatusName field.
func (o *FindingCaseResponseDataAttributes) SetStatusName(v string) {
	o.StatusName = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *FindingCaseResponseDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FindingCaseResponseDataAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FindingCaseResponseDataAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FindingCaseResponseDataAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FindingCaseResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ArchivedAt != nil {
		if o.ArchivedAt.Nanosecond() == 0 {
			toSerialize["archived_at"] = o.ArchivedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["archived_at"] = o.ArchivedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.AssignedTo != nil {
		toSerialize["assigned_to"] = o.AssignedTo
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ClosedAt != nil {
		if o.ClosedAt.Nanosecond() == 0 {
			toSerialize["closed_at"] = o.ClosedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["closed_at"] = o.ClosedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreationSource != nil {
		toSerialize["creation_source"] = o.CreationSource
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DueDate != nil {
		toSerialize["due_date"] = o.DueDate
	}
	if o.Insights != nil {
		toSerialize["insights"] = o.Insights
	}
	if o.JiraIssue != nil {
		toSerialize["jira_issue"] = o.JiraIssue
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusGroup != nil {
		toSerialize["status_group"] = o.StatusGroup
	}
	if o.StatusName != nil {
		toSerialize["status_name"] = o.StatusName
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FindingCaseResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchivedAt     *time.Time          `json:"archived_at,omitempty"`
		AssignedTo     *RelationshipToUser `json:"assigned_to,omitempty"`
		Attributes     map[string][]string `json:"attributes,omitempty"`
		ClosedAt       *time.Time          `json:"closed_at,omitempty"`
		CreatedAt      *time.Time          `json:"created_at,omitempty"`
		CreationSource *string             `json:"creation_source,omitempty"`
		Description    *string             `json:"description,omitempty"`
		DueDate        *string             `json:"due_date,omitempty"`
		Insights       []CaseInsightsItems `json:"insights,omitempty"`
		JiraIssue      *FindingJiraIssue   `json:"jira_issue,omitempty"`
		Key            *string             `json:"key,omitempty"`
		ModifiedAt     *time.Time          `json:"modified_at,omitempty"`
		Priority       *string             `json:"priority,omitempty"`
		Status         *string             `json:"status,omitempty"`
		StatusGroup    *string             `json:"status_group,omitempty"`
		StatusName     *string             `json:"status_name,omitempty"`
		Title          *string             `json:"title,omitempty"`
		Type           *string             `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archived_at", "assigned_to", "attributes", "closed_at", "created_at", "creation_source", "description", "due_date", "insights", "jira_issue", "key", "modified_at", "priority", "status", "status_group", "status_name", "title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ArchivedAt = all.ArchivedAt
	if all.AssignedTo != nil && all.AssignedTo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AssignedTo = all.AssignedTo
	o.Attributes = all.Attributes
	o.ClosedAt = all.ClosedAt
	o.CreatedAt = all.CreatedAt
	o.CreationSource = all.CreationSource
	o.Description = all.Description
	o.DueDate = all.DueDate
	o.Insights = all.Insights
	if all.JiraIssue != nil && all.JiraIssue.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JiraIssue = all.JiraIssue
	o.Key = all.Key
	o.ModifiedAt = all.ModifiedAt
	o.Priority = all.Priority
	o.Status = all.Status
	o.StatusGroup = all.StatusGroup
	o.StatusName = all.StatusName
	o.Title = all.Title
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
