// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueTemplateDataAttributes Attributes of a Jira issue template
type JiraIssueTemplateDataAttributes struct {
	// Custom fields for the Jira issue template
	Fields map[string]interface{} `json:"fields"`
	// The ID of the Jira issue type
	IssueTypeId string `json:"issue_type_id"`
	// The name of the issue template
	Name string `json:"name"`
	// The ID of the Jira project
	ProjectId string `json:"project_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssueTemplateDataAttributes instantiates a new JiraIssueTemplateDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssueTemplateDataAttributes(fields map[string]interface{}, issueTypeId string, name string, projectId string) *JiraIssueTemplateDataAttributes {
	this := JiraIssueTemplateDataAttributes{}
	this.Fields = fields
	this.IssueTypeId = issueTypeId
	this.Name = name
	this.ProjectId = projectId
	return &this
}

// NewJiraIssueTemplateDataAttributesWithDefaults instantiates a new JiraIssueTemplateDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueTemplateDataAttributesWithDefaults() *JiraIssueTemplateDataAttributes {
	this := JiraIssueTemplateDataAttributes{}
	return &this
}

// GetFields returns the Fields field value.
func (o *JiraIssueTemplateDataAttributes) GetFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateDataAttributes) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value.
func (o *JiraIssueTemplateDataAttributes) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetIssueTypeId returns the IssueTypeId field value.
func (o *JiraIssueTemplateDataAttributes) GetIssueTypeId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueTypeId
}

// GetIssueTypeIdOk returns a tuple with the IssueTypeId field value
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateDataAttributes) GetIssueTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueTypeId, true
}

// SetIssueTypeId sets field value.
func (o *JiraIssueTemplateDataAttributes) SetIssueTypeId(v string) {
	o.IssueTypeId = v
}

// GetName returns the Name field value.
func (o *JiraIssueTemplateDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *JiraIssueTemplateDataAttributes) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value.
func (o *JiraIssueTemplateDataAttributes) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateDataAttributes) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *JiraIssueTemplateDataAttributes) SetProjectId(v string) {
	o.ProjectId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssueTemplateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fields"] = o.Fields
	toSerialize["issue_type_id"] = o.IssueTypeId
	toSerialize["name"] = o.Name
	toSerialize["project_id"] = o.ProjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssueTemplateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields      *map[string]interface{} `json:"fields"`
		IssueTypeId *string                 `json:"issue_type_id"`
		Name        *string                 `json:"name"`
		ProjectId   *string                 `json:"project_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fields == nil {
		return fmt.Errorf("required field fields missing")
	}
	if all.IssueTypeId == nil {
		return fmt.Errorf("required field issue_type_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "issue_type_id", "name", "project_id"})
	} else {
		return err
	}
	o.Fields = *all.Fields
	o.IssueTypeId = *all.IssueTypeId
	o.Name = *all.Name
	o.ProjectId = *all.ProjectId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
