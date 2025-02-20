// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueResult Jira issue information
type JiraIssueResult struct {
	// Jira issue ID
	IssueId *string `json:"issue_id,omitempty"`
	// Jira issue key
	IssueKey *string `json:"issue_key,omitempty"`
	// Jira issue URL
	IssueUrl *string `json:"issue_url,omitempty"`
	// Jira project key
	ProjectKey *string `json:"project_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssueResult instantiates a new JiraIssueResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssueResult() *JiraIssueResult {
	this := JiraIssueResult{}
	return &this
}

// NewJiraIssueResultWithDefaults instantiates a new JiraIssueResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueResultWithDefaults() *JiraIssueResult {
	this := JiraIssueResult{}
	return &this
}

// GetIssueId returns the IssueId field value if set, zero value otherwise.
func (o *JiraIssueResult) GetIssueId() string {
	if o == nil || o.IssueId == nil {
		var ret string
		return ret
	}
	return *o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueResult) GetIssueIdOk() (*string, bool) {
	if o == nil || o.IssueId == nil {
		return nil, false
	}
	return o.IssueId, true
}

// HasIssueId returns a boolean if a field has been set.
func (o *JiraIssueResult) HasIssueId() bool {
	return o != nil && o.IssueId != nil
}

// SetIssueId gets a reference to the given string and assigns it to the IssueId field.
func (o *JiraIssueResult) SetIssueId(v string) {
	o.IssueId = &v
}

// GetIssueKey returns the IssueKey field value if set, zero value otherwise.
func (o *JiraIssueResult) GetIssueKey() string {
	if o == nil || o.IssueKey == nil {
		var ret string
		return ret
	}
	return *o.IssueKey
}

// GetIssueKeyOk returns a tuple with the IssueKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueResult) GetIssueKeyOk() (*string, bool) {
	if o == nil || o.IssueKey == nil {
		return nil, false
	}
	return o.IssueKey, true
}

// HasIssueKey returns a boolean if a field has been set.
func (o *JiraIssueResult) HasIssueKey() bool {
	return o != nil && o.IssueKey != nil
}

// SetIssueKey gets a reference to the given string and assigns it to the IssueKey field.
func (o *JiraIssueResult) SetIssueKey(v string) {
	o.IssueKey = &v
}

// GetIssueUrl returns the IssueUrl field value if set, zero value otherwise.
func (o *JiraIssueResult) GetIssueUrl() string {
	if o == nil || o.IssueUrl == nil {
		var ret string
		return ret
	}
	return *o.IssueUrl
}

// GetIssueUrlOk returns a tuple with the IssueUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueResult) GetIssueUrlOk() (*string, bool) {
	if o == nil || o.IssueUrl == nil {
		return nil, false
	}
	return o.IssueUrl, true
}

// HasIssueUrl returns a boolean if a field has been set.
func (o *JiraIssueResult) HasIssueUrl() bool {
	return o != nil && o.IssueUrl != nil
}

// SetIssueUrl gets a reference to the given string and assigns it to the IssueUrl field.
func (o *JiraIssueResult) SetIssueUrl(v string) {
	o.IssueUrl = &v
}

// GetProjectKey returns the ProjectKey field value if set, zero value otherwise.
func (o *JiraIssueResult) GetProjectKey() string {
	if o == nil || o.ProjectKey == nil {
		var ret string
		return ret
	}
	return *o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueResult) GetProjectKeyOk() (*string, bool) {
	if o == nil || o.ProjectKey == nil {
		return nil, false
	}
	return o.ProjectKey, true
}

// HasProjectKey returns a boolean if a field has been set.
func (o *JiraIssueResult) HasProjectKey() bool {
	return o != nil && o.ProjectKey != nil
}

// SetProjectKey gets a reference to the given string and assigns it to the ProjectKey field.
func (o *JiraIssueResult) SetProjectKey(v string) {
	o.ProjectKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssueResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IssueId != nil {
		toSerialize["issue_id"] = o.IssueId
	}
	if o.IssueKey != nil {
		toSerialize["issue_key"] = o.IssueKey
	}
	if o.IssueUrl != nil {
		toSerialize["issue_url"] = o.IssueUrl
	}
	if o.ProjectKey != nil {
		toSerialize["project_key"] = o.ProjectKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssueResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IssueId    *string `json:"issue_id,omitempty"`
		IssueKey   *string `json:"issue_key,omitempty"`
		IssueUrl   *string `json:"issue_url,omitempty"`
		ProjectKey *string `json:"project_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"issue_id", "issue_key", "issue_url", "project_key"})
	} else {
		return err
	}
	o.IssueId = all.IssueId
	o.IssueKey = all.IssueKey
	o.IssueUrl = all.IssueUrl
	o.ProjectKey = all.ProjectKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
