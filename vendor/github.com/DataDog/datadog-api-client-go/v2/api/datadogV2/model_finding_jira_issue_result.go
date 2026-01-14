// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingJiraIssueResult Result of the Jira issue creation.
type FindingJiraIssueResult struct {
	// Account ID of the Jira issue.
	AccountId *string `json:"account_id,omitempty"`
	// Unique identifier of the Jira issue.
	IssueId *string `json:"issue_id,omitempty"`
	// Key of the Jira issue.
	IssueKey *string `json:"issue_key,omitempty"`
	// URL of the Jira issue.
	IssueUrl *string `json:"issue_url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFindingJiraIssueResult instantiates a new FindingJiraIssueResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFindingJiraIssueResult() *FindingJiraIssueResult {
	this := FindingJiraIssueResult{}
	return &this
}

// NewFindingJiraIssueResultWithDefaults instantiates a new FindingJiraIssueResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFindingJiraIssueResultWithDefaults() *FindingJiraIssueResult {
	this := FindingJiraIssueResult{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *FindingJiraIssueResult) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingJiraIssueResult) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *FindingJiraIssueResult) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *FindingJiraIssueResult) SetAccountId(v string) {
	o.AccountId = &v
}

// GetIssueId returns the IssueId field value if set, zero value otherwise.
func (o *FindingJiraIssueResult) GetIssueId() string {
	if o == nil || o.IssueId == nil {
		var ret string
		return ret
	}
	return *o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingJiraIssueResult) GetIssueIdOk() (*string, bool) {
	if o == nil || o.IssueId == nil {
		return nil, false
	}
	return o.IssueId, true
}

// HasIssueId returns a boolean if a field has been set.
func (o *FindingJiraIssueResult) HasIssueId() bool {
	return o != nil && o.IssueId != nil
}

// SetIssueId gets a reference to the given string and assigns it to the IssueId field.
func (o *FindingJiraIssueResult) SetIssueId(v string) {
	o.IssueId = &v
}

// GetIssueKey returns the IssueKey field value if set, zero value otherwise.
func (o *FindingJiraIssueResult) GetIssueKey() string {
	if o == nil || o.IssueKey == nil {
		var ret string
		return ret
	}
	return *o.IssueKey
}

// GetIssueKeyOk returns a tuple with the IssueKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingJiraIssueResult) GetIssueKeyOk() (*string, bool) {
	if o == nil || o.IssueKey == nil {
		return nil, false
	}
	return o.IssueKey, true
}

// HasIssueKey returns a boolean if a field has been set.
func (o *FindingJiraIssueResult) HasIssueKey() bool {
	return o != nil && o.IssueKey != nil
}

// SetIssueKey gets a reference to the given string and assigns it to the IssueKey field.
func (o *FindingJiraIssueResult) SetIssueKey(v string) {
	o.IssueKey = &v
}

// GetIssueUrl returns the IssueUrl field value if set, zero value otherwise.
func (o *FindingJiraIssueResult) GetIssueUrl() string {
	if o == nil || o.IssueUrl == nil {
		var ret string
		return ret
	}
	return *o.IssueUrl
}

// GetIssueUrlOk returns a tuple with the IssueUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingJiraIssueResult) GetIssueUrlOk() (*string, bool) {
	if o == nil || o.IssueUrl == nil {
		return nil, false
	}
	return o.IssueUrl, true
}

// HasIssueUrl returns a boolean if a field has been set.
func (o *FindingJiraIssueResult) HasIssueUrl() bool {
	return o != nil && o.IssueUrl != nil
}

// SetIssueUrl gets a reference to the given string and assigns it to the IssueUrl field.
func (o *FindingJiraIssueResult) SetIssueUrl(v string) {
	o.IssueUrl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FindingJiraIssueResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FindingJiraIssueResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId *string `json:"account_id,omitempty"`
		IssueId   *string `json:"issue_id,omitempty"`
		IssueKey  *string `json:"issue_key,omitempty"`
		IssueUrl  *string `json:"issue_url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "issue_id", "issue_key", "issue_url"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.IssueId = all.IssueId
	o.IssueKey = all.IssueKey
	o.IssueUrl = all.IssueUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
