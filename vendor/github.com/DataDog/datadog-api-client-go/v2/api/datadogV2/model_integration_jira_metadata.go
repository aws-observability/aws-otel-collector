// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationJiraMetadata Metadata for connecting a case management project to a Jira project.
type IntegrationJiraMetadata struct {
	// The Jira account identifier.
	AccountId *string `json:"account_id,omitempty"`
	// The Jira issue type identifier to use when creating issues.
	IssueTypeId *string `json:"issue_type_id,omitempty"`
	// The Jira project identifier to associate with this case project.
	ProjectId *string `json:"project_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationJiraMetadata instantiates a new IntegrationJiraMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationJiraMetadata() *IntegrationJiraMetadata {
	this := IntegrationJiraMetadata{}
	return &this
}

// NewIntegrationJiraMetadataWithDefaults instantiates a new IntegrationJiraMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationJiraMetadataWithDefaults() *IntegrationJiraMetadata {
	this := IntegrationJiraMetadata{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *IntegrationJiraMetadata) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraMetadata) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *IntegrationJiraMetadata) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *IntegrationJiraMetadata) SetAccountId(v string) {
	o.AccountId = &v
}

// GetIssueTypeId returns the IssueTypeId field value if set, zero value otherwise.
func (o *IntegrationJiraMetadata) GetIssueTypeId() string {
	if o == nil || o.IssueTypeId == nil {
		var ret string
		return ret
	}
	return *o.IssueTypeId
}

// GetIssueTypeIdOk returns a tuple with the IssueTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraMetadata) GetIssueTypeIdOk() (*string, bool) {
	if o == nil || o.IssueTypeId == nil {
		return nil, false
	}
	return o.IssueTypeId, true
}

// HasIssueTypeId returns a boolean if a field has been set.
func (o *IntegrationJiraMetadata) HasIssueTypeId() bool {
	return o != nil && o.IssueTypeId != nil
}

// SetIssueTypeId gets a reference to the given string and assigns it to the IssueTypeId field.
func (o *IntegrationJiraMetadata) SetIssueTypeId(v string) {
	o.IssueTypeId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *IntegrationJiraMetadata) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraMetadata) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *IntegrationJiraMetadata) HasProjectId() bool {
	return o != nil && o.ProjectId != nil
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *IntegrationJiraMetadata) SetProjectId(v string) {
	o.ProjectId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationJiraMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.IssueTypeId != nil {
		toSerialize["issue_type_id"] = o.IssueTypeId
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
func (o *IntegrationJiraMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId   *string `json:"account_id,omitempty"`
		IssueTypeId *string `json:"issue_type_id,omitempty"`
		ProjectId   *string `json:"project_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "issue_type_id", "project_id"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.IssueTypeId = all.IssueTypeId
	o.ProjectId = all.ProjectId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
