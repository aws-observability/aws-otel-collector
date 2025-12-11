// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueCaseAttributes Object containing the information of a case.
type IssueCaseAttributes struct {
	// Timestamp of when the case was archived.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
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
	Insights []IssueCaseInsight `json:"insights,omitempty"`
	// Jira issue of the case.
	JiraIssue *IssueCaseJiraIssue `json:"jira_issue,omitempty"`
	// Key of the case.
	Key *string `json:"key,omitempty"`
	// Timestamp of when the case was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Case priority
	Priority *CasePriority `json:"priority,omitempty"`
	// Case status
	Status *CaseStatus `json:"status,omitempty"`
	// Title of the case.
	Title *string `json:"title,omitempty"`
	// Type of the case.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueCaseAttributes instantiates a new IssueCaseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueCaseAttributes() *IssueCaseAttributes {
	this := IssueCaseAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// NewIssueCaseAttributesWithDefaults instantiates a new IssueCaseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueCaseAttributesWithDefaults() *IssueCaseAttributes {
	this := IssueCaseAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasArchivedAt() bool {
	return o != nil && o.ArchivedAt != nil
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *IssueCaseAttributes) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetClosedAt() time.Time {
	if o == nil || o.ClosedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetClosedAtOk() (*time.Time, bool) {
	if o == nil || o.ClosedAt == nil {
		return nil, false
	}
	return o.ClosedAt, true
}

// HasClosedAt returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasClosedAt() bool {
	return o != nil && o.ClosedAt != nil
}

// SetClosedAt gets a reference to the given time.Time and assigns it to the ClosedAt field.
func (o *IssueCaseAttributes) SetClosedAt(v time.Time) {
	o.ClosedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IssueCaseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreationSource returns the CreationSource field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetCreationSource() string {
	if o == nil || o.CreationSource == nil {
		var ret string
		return ret
	}
	return *o.CreationSource
}

// GetCreationSourceOk returns a tuple with the CreationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetCreationSourceOk() (*string, bool) {
	if o == nil || o.CreationSource == nil {
		return nil, false
	}
	return o.CreationSource, true
}

// HasCreationSource returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasCreationSource() bool {
	return o != nil && o.CreationSource != nil
}

// SetCreationSource gets a reference to the given string and assigns it to the CreationSource field.
func (o *IssueCaseAttributes) SetCreationSource(v string) {
	o.CreationSource = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IssueCaseAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetDueDate() string {
	if o == nil || o.DueDate == nil {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetDueDateOk() (*string, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasDueDate() bool {
	return o != nil && o.DueDate != nil
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *IssueCaseAttributes) SetDueDate(v string) {
	o.DueDate = &v
}

// GetInsights returns the Insights field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetInsights() []IssueCaseInsight {
	if o == nil || o.Insights == nil {
		var ret []IssueCaseInsight
		return ret
	}
	return o.Insights
}

// GetInsightsOk returns a tuple with the Insights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetInsightsOk() (*[]IssueCaseInsight, bool) {
	if o == nil || o.Insights == nil {
		return nil, false
	}
	return &o.Insights, true
}

// HasInsights returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasInsights() bool {
	return o != nil && o.Insights != nil
}

// SetInsights gets a reference to the given []IssueCaseInsight and assigns it to the Insights field.
func (o *IssueCaseAttributes) SetInsights(v []IssueCaseInsight) {
	o.Insights = v
}

// GetJiraIssue returns the JiraIssue field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetJiraIssue() IssueCaseJiraIssue {
	if o == nil || o.JiraIssue == nil {
		var ret IssueCaseJiraIssue
		return ret
	}
	return *o.JiraIssue
}

// GetJiraIssueOk returns a tuple with the JiraIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetJiraIssueOk() (*IssueCaseJiraIssue, bool) {
	if o == nil || o.JiraIssue == nil {
		return nil, false
	}
	return o.JiraIssue, true
}

// HasJiraIssue returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasJiraIssue() bool {
	return o != nil && o.JiraIssue != nil
}

// SetJiraIssue gets a reference to the given IssueCaseJiraIssue and assigns it to the JiraIssue field.
func (o *IssueCaseAttributes) SetJiraIssue(v IssueCaseJiraIssue) {
	o.JiraIssue = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *IssueCaseAttributes) SetKey(v string) {
	o.Key = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *IssueCaseAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetPriority() CasePriority {
	if o == nil || o.Priority == nil {
		var ret CasePriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetPriorityOk() (*CasePriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given CasePriority and assigns it to the Priority field.
func (o *IssueCaseAttributes) SetPriority(v CasePriority) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetStatus() CaseStatus {
	if o == nil || o.Status == nil {
		var ret CaseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetStatusOk() (*CaseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given CaseStatus and assigns it to the Status field.
func (o *IssueCaseAttributes) SetStatus(v CaseStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *IssueCaseAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IssueCaseAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IssueCaseAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IssueCaseAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueCaseAttributes) MarshalJSON() ([]byte, error) {
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
func (o *IssueCaseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchivedAt     *time.Time          `json:"archived_at,omitempty"`
		ClosedAt       *time.Time          `json:"closed_at,omitempty"`
		CreatedAt      *time.Time          `json:"created_at,omitempty"`
		CreationSource *string             `json:"creation_source,omitempty"`
		Description    *string             `json:"description,omitempty"`
		DueDate        *string             `json:"due_date,omitempty"`
		Insights       []IssueCaseInsight  `json:"insights,omitempty"`
		JiraIssue      *IssueCaseJiraIssue `json:"jira_issue,omitempty"`
		Key            *string             `json:"key,omitempty"`
		ModifiedAt     *time.Time          `json:"modified_at,omitempty"`
		Priority       *CasePriority       `json:"priority,omitempty"`
		Status         *CaseStatus         `json:"status,omitempty"`
		Title          *string             `json:"title,omitempty"`
		Type           *string             `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archived_at", "closed_at", "created_at", "creation_source", "description", "due_date", "insights", "jira_issue", "key", "modified_at", "priority", "status", "title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ArchivedAt = all.ArchivedAt
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
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
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
