// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueCreateAttributes Jira issue creation attributes
type JiraIssueCreateAttributes struct {
	// Additional Jira fields
	Fields map[string]interface{} `json:"fields,omitempty"`
	// Jira issue type ID
	IssueTypeId string `json:"issue_type_id"`
	// Jira account ID
	JiraAccountId string `json:"jira_account_id"`
	// Jira project ID
	ProjectId string `json:"project_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssueCreateAttributes instantiates a new JiraIssueCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssueCreateAttributes(issueTypeId string, jiraAccountId string, projectId string) *JiraIssueCreateAttributes {
	this := JiraIssueCreateAttributes{}
	this.IssueTypeId = issueTypeId
	this.JiraAccountId = jiraAccountId
	this.ProjectId = projectId
	return &this
}

// NewJiraIssueCreateAttributesWithDefaults instantiates a new JiraIssueCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueCreateAttributesWithDefaults() *JiraIssueCreateAttributes {
	this := JiraIssueCreateAttributes{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *JiraIssueCreateAttributes) GetFields() map[string]interface{} {
	if o == nil || o.Fields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueCreateAttributes) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *JiraIssueCreateAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *JiraIssueCreateAttributes) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetIssueTypeId returns the IssueTypeId field value.
func (o *JiraIssueCreateAttributes) GetIssueTypeId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueTypeId
}

// GetIssueTypeIdOk returns a tuple with the IssueTypeId field value
// and a boolean to check if the value has been set.
func (o *JiraIssueCreateAttributes) GetIssueTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueTypeId, true
}

// SetIssueTypeId sets field value.
func (o *JiraIssueCreateAttributes) SetIssueTypeId(v string) {
	o.IssueTypeId = v
}

// GetJiraAccountId returns the JiraAccountId field value.
func (o *JiraIssueCreateAttributes) GetJiraAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JiraAccountId
}

// GetJiraAccountIdOk returns a tuple with the JiraAccountId field value
// and a boolean to check if the value has been set.
func (o *JiraIssueCreateAttributes) GetJiraAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JiraAccountId, true
}

// SetJiraAccountId sets field value.
func (o *JiraIssueCreateAttributes) SetJiraAccountId(v string) {
	o.JiraAccountId = v
}

// GetProjectId returns the ProjectId field value.
func (o *JiraIssueCreateAttributes) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *JiraIssueCreateAttributes) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *JiraIssueCreateAttributes) SetProjectId(v string) {
	o.ProjectId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssueCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	toSerialize["issue_type_id"] = o.IssueTypeId
	toSerialize["jira_account_id"] = o.JiraAccountId
	toSerialize["project_id"] = o.ProjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssueCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields        map[string]interface{} `json:"fields,omitempty"`
		IssueTypeId   *string                `json:"issue_type_id"`
		JiraAccountId *string                `json:"jira_account_id"`
		ProjectId     *string                `json:"project_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IssueTypeId == nil {
		return fmt.Errorf("required field issue_type_id missing")
	}
	if all.JiraAccountId == nil {
		return fmt.Errorf("required field jira_account_id missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "issue_type_id", "jira_account_id", "project_id"})
	} else {
		return err
	}
	o.Fields = all.Fields
	o.IssueTypeId = *all.IssueTypeId
	o.JiraAccountId = *all.JiraAccountId
	o.ProjectId = *all.ProjectId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
