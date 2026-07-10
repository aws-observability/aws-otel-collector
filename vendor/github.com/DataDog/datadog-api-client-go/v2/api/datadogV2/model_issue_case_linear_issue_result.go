// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueCaseLinearIssueResult Contains the identifiers and URL for a successfully created Linear issue.
type IssueCaseLinearIssueResult struct {
	// Linear account identifier.
	AccountId *string `json:"account_id,omitempty"`
	// Linear issue identifier.
	IssueId *string `json:"issue_id,omitempty"`
	// Linear issue key.
	IssueKey *string `json:"issue_key,omitempty"`
	// Linear issue URL.
	IssueUrl *string `json:"issue_url,omitempty"`
	// Linear team identifier.
	TeamId *string `json:"team_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueCaseLinearIssueResult instantiates a new IssueCaseLinearIssueResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueCaseLinearIssueResult() *IssueCaseLinearIssueResult {
	this := IssueCaseLinearIssueResult{}
	return &this
}

// NewIssueCaseLinearIssueResultWithDefaults instantiates a new IssueCaseLinearIssueResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueCaseLinearIssueResultWithDefaults() *IssueCaseLinearIssueResult {
	this := IssueCaseLinearIssueResult{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *IssueCaseLinearIssueResult) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseLinearIssueResult) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *IssueCaseLinearIssueResult) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *IssueCaseLinearIssueResult) SetAccountId(v string) {
	o.AccountId = &v
}

// GetIssueId returns the IssueId field value if set, zero value otherwise.
func (o *IssueCaseLinearIssueResult) GetIssueId() string {
	if o == nil || o.IssueId == nil {
		var ret string
		return ret
	}
	return *o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseLinearIssueResult) GetIssueIdOk() (*string, bool) {
	if o == nil || o.IssueId == nil {
		return nil, false
	}
	return o.IssueId, true
}

// HasIssueId returns a boolean if a field has been set.
func (o *IssueCaseLinearIssueResult) HasIssueId() bool {
	return o != nil && o.IssueId != nil
}

// SetIssueId gets a reference to the given string and assigns it to the IssueId field.
func (o *IssueCaseLinearIssueResult) SetIssueId(v string) {
	o.IssueId = &v
}

// GetIssueKey returns the IssueKey field value if set, zero value otherwise.
func (o *IssueCaseLinearIssueResult) GetIssueKey() string {
	if o == nil || o.IssueKey == nil {
		var ret string
		return ret
	}
	return *o.IssueKey
}

// GetIssueKeyOk returns a tuple with the IssueKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseLinearIssueResult) GetIssueKeyOk() (*string, bool) {
	if o == nil || o.IssueKey == nil {
		return nil, false
	}
	return o.IssueKey, true
}

// HasIssueKey returns a boolean if a field has been set.
func (o *IssueCaseLinearIssueResult) HasIssueKey() bool {
	return o != nil && o.IssueKey != nil
}

// SetIssueKey gets a reference to the given string and assigns it to the IssueKey field.
func (o *IssueCaseLinearIssueResult) SetIssueKey(v string) {
	o.IssueKey = &v
}

// GetIssueUrl returns the IssueUrl field value if set, zero value otherwise.
func (o *IssueCaseLinearIssueResult) GetIssueUrl() string {
	if o == nil || o.IssueUrl == nil {
		var ret string
		return ret
	}
	return *o.IssueUrl
}

// GetIssueUrlOk returns a tuple with the IssueUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseLinearIssueResult) GetIssueUrlOk() (*string, bool) {
	if o == nil || o.IssueUrl == nil {
		return nil, false
	}
	return o.IssueUrl, true
}

// HasIssueUrl returns a boolean if a field has been set.
func (o *IssueCaseLinearIssueResult) HasIssueUrl() bool {
	return o != nil && o.IssueUrl != nil
}

// SetIssueUrl gets a reference to the given string and assigns it to the IssueUrl field.
func (o *IssueCaseLinearIssueResult) SetIssueUrl(v string) {
	o.IssueUrl = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *IssueCaseLinearIssueResult) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseLinearIssueResult) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *IssueCaseLinearIssueResult) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *IssueCaseLinearIssueResult) SetTeamId(v string) {
	o.TeamId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueCaseLinearIssueResult) MarshalJSON() ([]byte, error) {
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
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueCaseLinearIssueResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId *string `json:"account_id,omitempty"`
		IssueId   *string `json:"issue_id,omitempty"`
		IssueKey  *string `json:"issue_key,omitempty"`
		IssueUrl  *string `json:"issue_url,omitempty"`
		TeamId    *string `json:"team_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "issue_id", "issue_key", "issue_url", "team_id"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.IssueId = all.IssueId
	o.IssueKey = all.IssueKey
	o.IssueUrl = all.IssueUrl
	o.TeamId = all.TeamId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
