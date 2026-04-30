// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueTemplateCreateRequestAttributes Attributes for creating a Jira issue template
type JiraIssueTemplateCreateRequestAttributes struct {
	// Custom fields for the Jira issue template
	Fields map[string]interface{} `json:"fields,omitempty"`
	// The ID of the Jira issue type
	IssueTypeId *string `json:"issue_type_id,omitempty"`
	// Reference to the Jira account
	JiraAccount *JiraIssueTemplateCreateRequestAttributesJiraAccount `json:"jira-account,omitempty"`
	// The name of the issue template
	Name *string `json:"name,omitempty"`
	// The ID of the Jira project
	ProjectId *string `json:"project_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssueTemplateCreateRequestAttributes instantiates a new JiraIssueTemplateCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssueTemplateCreateRequestAttributes() *JiraIssueTemplateCreateRequestAttributes {
	this := JiraIssueTemplateCreateRequestAttributes{}
	return &this
}

// NewJiraIssueTemplateCreateRequestAttributesWithDefaults instantiates a new JiraIssueTemplateCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueTemplateCreateRequestAttributesWithDefaults() *JiraIssueTemplateCreateRequestAttributes {
	this := JiraIssueTemplateCreateRequestAttributes{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *JiraIssueTemplateCreateRequestAttributes) GetFields() map[string]interface{} {
	if o == nil || o.Fields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateCreateRequestAttributes) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *JiraIssueTemplateCreateRequestAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *JiraIssueTemplateCreateRequestAttributes) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetIssueTypeId returns the IssueTypeId field value if set, zero value otherwise.
func (o *JiraIssueTemplateCreateRequestAttributes) GetIssueTypeId() string {
	if o == nil || o.IssueTypeId == nil {
		var ret string
		return ret
	}
	return *o.IssueTypeId
}

// GetIssueTypeIdOk returns a tuple with the IssueTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateCreateRequestAttributes) GetIssueTypeIdOk() (*string, bool) {
	if o == nil || o.IssueTypeId == nil {
		return nil, false
	}
	return o.IssueTypeId, true
}

// HasIssueTypeId returns a boolean if a field has been set.
func (o *JiraIssueTemplateCreateRequestAttributes) HasIssueTypeId() bool {
	return o != nil && o.IssueTypeId != nil
}

// SetIssueTypeId gets a reference to the given string and assigns it to the IssueTypeId field.
func (o *JiraIssueTemplateCreateRequestAttributes) SetIssueTypeId(v string) {
	o.IssueTypeId = &v
}

// GetJiraAccount returns the JiraAccount field value if set, zero value otherwise.
func (o *JiraIssueTemplateCreateRequestAttributes) GetJiraAccount() JiraIssueTemplateCreateRequestAttributesJiraAccount {
	if o == nil || o.JiraAccount == nil {
		var ret JiraIssueTemplateCreateRequestAttributesJiraAccount
		return ret
	}
	return *o.JiraAccount
}

// GetJiraAccountOk returns a tuple with the JiraAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateCreateRequestAttributes) GetJiraAccountOk() (*JiraIssueTemplateCreateRequestAttributesJiraAccount, bool) {
	if o == nil || o.JiraAccount == nil {
		return nil, false
	}
	return o.JiraAccount, true
}

// HasJiraAccount returns a boolean if a field has been set.
func (o *JiraIssueTemplateCreateRequestAttributes) HasJiraAccount() bool {
	return o != nil && o.JiraAccount != nil
}

// SetJiraAccount gets a reference to the given JiraIssueTemplateCreateRequestAttributesJiraAccount and assigns it to the JiraAccount field.
func (o *JiraIssueTemplateCreateRequestAttributes) SetJiraAccount(v JiraIssueTemplateCreateRequestAttributesJiraAccount) {
	o.JiraAccount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JiraIssueTemplateCreateRequestAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateCreateRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JiraIssueTemplateCreateRequestAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JiraIssueTemplateCreateRequestAttributes) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *JiraIssueTemplateCreateRequestAttributes) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateCreateRequestAttributes) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *JiraIssueTemplateCreateRequestAttributes) HasProjectId() bool {
	return o != nil && o.ProjectId != nil
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *JiraIssueTemplateCreateRequestAttributes) SetProjectId(v string) {
	o.ProjectId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssueTemplateCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.IssueTypeId != nil {
		toSerialize["issue_type_id"] = o.IssueTypeId
	}
	if o.JiraAccount != nil {
		toSerialize["jira-account"] = o.JiraAccount
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssueTemplateCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields      map[string]interface{}                               `json:"fields,omitempty"`
		IssueTypeId *string                                              `json:"issue_type_id,omitempty"`
		JiraAccount *JiraIssueTemplateCreateRequestAttributesJiraAccount `json:"jira-account,omitempty"`
		Name        *string                                              `json:"name,omitempty"`
		ProjectId   *string                                              `json:"project_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "issue_type_id", "jira-account", "name", "project_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Fields = all.Fields
	o.IssueTypeId = all.IssueTypeId
	if all.JiraAccount != nil && all.JiraAccount.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JiraAccount = all.JiraAccount
	o.Name = all.Name
	o.ProjectId = all.ProjectId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
