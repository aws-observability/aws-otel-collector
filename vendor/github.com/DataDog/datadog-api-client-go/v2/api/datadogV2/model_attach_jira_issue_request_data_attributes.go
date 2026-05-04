// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AttachJiraIssueRequestDataAttributes Attributes of the Jira issue to attach security findings to.
type AttachJiraIssueRequestDataAttributes struct {
	// URL of the Jira issue to attach security findings to.
	JiraIssueUrl string `json:"jira_issue_url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAttachJiraIssueRequestDataAttributes instantiates a new AttachJiraIssueRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAttachJiraIssueRequestDataAttributes(jiraIssueUrl string) *AttachJiraIssueRequestDataAttributes {
	this := AttachJiraIssueRequestDataAttributes{}
	this.JiraIssueUrl = jiraIssueUrl
	return &this
}

// NewAttachJiraIssueRequestDataAttributesWithDefaults instantiates a new AttachJiraIssueRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAttachJiraIssueRequestDataAttributesWithDefaults() *AttachJiraIssueRequestDataAttributes {
	this := AttachJiraIssueRequestDataAttributes{}
	return &this
}

// GetJiraIssueUrl returns the JiraIssueUrl field value.
func (o *AttachJiraIssueRequestDataAttributes) GetJiraIssueUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JiraIssueUrl
}

// GetJiraIssueUrlOk returns a tuple with the JiraIssueUrl field value
// and a boolean to check if the value has been set.
func (o *AttachJiraIssueRequestDataAttributes) GetJiraIssueUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JiraIssueUrl, true
}

// SetJiraIssueUrl sets field value.
func (o *AttachJiraIssueRequestDataAttributes) SetJiraIssueUrl(v string) {
	o.JiraIssueUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AttachJiraIssueRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["jira_issue_url"] = o.JiraIssueUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AttachJiraIssueRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JiraIssueUrl *string `json:"jira_issue_url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JiraIssueUrl == nil {
		return fmt.Errorf("required field jira_issue_url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"jira_issue_url"})
	} else {
		return err
	}
	o.JiraIssueUrl = *all.JiraIssueUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
