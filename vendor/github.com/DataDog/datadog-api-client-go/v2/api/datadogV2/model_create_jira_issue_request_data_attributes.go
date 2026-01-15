// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateJiraIssueRequestDataAttributes Attributes of the Jira issue to create.
type CreateJiraIssueRequestDataAttributes struct {
	// Unique identifier of the user assigned to the Jira issue.
	AssigneeId *string `json:"assignee_id,omitempty"`
	// Description of the Jira issue. If not provided, the description will be automatically generated.
	Description *string `json:"description,omitempty"`
	// Custom fields of the Jira issue to create. For the list of available fields, see [Jira documentation](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-createmeta-projectidorkey-issuetypes-issuetypeid-get).
	Fields map[string]interface{} `json:"fields,omitempty"`
	// Case priority
	Priority *CasePriority `json:"priority,omitempty"`
	// Title of the Jira issue. If not provided, the title will be automatically generated.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateJiraIssueRequestDataAttributes instantiates a new CreateJiraIssueRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateJiraIssueRequestDataAttributes() *CreateJiraIssueRequestDataAttributes {
	this := CreateJiraIssueRequestDataAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// NewCreateJiraIssueRequestDataAttributesWithDefaults instantiates a new CreateJiraIssueRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateJiraIssueRequestDataAttributesWithDefaults() *CreateJiraIssueRequestDataAttributes {
	this := CreateJiraIssueRequestDataAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *CreateJiraIssueRequestDataAttributes) GetAssigneeId() string {
	if o == nil || o.AssigneeId == nil {
		var ret string
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestDataAttributes) GetAssigneeIdOk() (*string, bool) {
	if o == nil || o.AssigneeId == nil {
		return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *CreateJiraIssueRequestDataAttributes) HasAssigneeId() bool {
	return o != nil && o.AssigneeId != nil
}

// SetAssigneeId gets a reference to the given string and assigns it to the AssigneeId field.
func (o *CreateJiraIssueRequestDataAttributes) SetAssigneeId(v string) {
	o.AssigneeId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateJiraIssueRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateJiraIssueRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateJiraIssueRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *CreateJiraIssueRequestDataAttributes) GetFields() map[string]interface{} {
	if o == nil || o.Fields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestDataAttributes) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CreateJiraIssueRequestDataAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *CreateJiraIssueRequestDataAttributes) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CreateJiraIssueRequestDataAttributes) GetPriority() CasePriority {
	if o == nil || o.Priority == nil {
		var ret CasePriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestDataAttributes) GetPriorityOk() (*CasePriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CreateJiraIssueRequestDataAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given CasePriority and assigns it to the Priority field.
func (o *CreateJiraIssueRequestDataAttributes) SetPriority(v CasePriority) {
	o.Priority = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateJiraIssueRequestDataAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateJiraIssueRequestDataAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateJiraIssueRequestDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateJiraIssueRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssigneeId != nil {
		toSerialize["assignee_id"] = o.AssigneeId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateJiraIssueRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssigneeId  *string                `json:"assignee_id,omitempty"`
		Description *string                `json:"description,omitempty"`
		Fields      map[string]interface{} `json:"fields,omitempty"`
		Priority    *CasePriority          `json:"priority,omitempty"`
		Title       *string                `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee_id", "description", "fields", "priority", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssigneeId = all.AssigneeId
	o.Description = all.Description
	o.Fields = all.Fields
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
