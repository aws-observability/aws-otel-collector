// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultGit Git information associated with the test result.
type SyntheticsTestResultGit struct {
	// Git branch name.
	Branch *string `json:"branch,omitempty"`
	// Details of the Git commit associated with the test result.
	Commit *SyntheticsTestResultGitCommit `json:"commit,omitempty"`
	// Git repository URL.
	RepositoryUrl *string `json:"repository_url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultGit instantiates a new SyntheticsTestResultGit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultGit() *SyntheticsTestResultGit {
	this := SyntheticsTestResultGit{}
	return &this
}

// NewSyntheticsTestResultGitWithDefaults instantiates a new SyntheticsTestResultGit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultGitWithDefaults() *SyntheticsTestResultGit {
	this := SyntheticsTestResultGit{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *SyntheticsTestResultGit) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGit) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *SyntheticsTestResultGit) HasBranch() bool {
	return o != nil && o.Branch != nil
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *SyntheticsTestResultGit) SetBranch(v string) {
	o.Branch = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *SyntheticsTestResultGit) GetCommit() SyntheticsTestResultGitCommit {
	if o == nil || o.Commit == nil {
		var ret SyntheticsTestResultGitCommit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGit) GetCommitOk() (*SyntheticsTestResultGitCommit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *SyntheticsTestResultGit) HasCommit() bool {
	return o != nil && o.Commit != nil
}

// SetCommit gets a reference to the given SyntheticsTestResultGitCommit and assigns it to the Commit field.
func (o *SyntheticsTestResultGit) SetCommit(v SyntheticsTestResultGitCommit) {
	o.Commit = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *SyntheticsTestResultGit) GetRepositoryUrl() string {
	if o == nil || o.RepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGit) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || o.RepositoryUrl == nil {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *SyntheticsTestResultGit) HasRepositoryUrl() bool {
	return o != nil && o.RepositoryUrl != nil
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *SyntheticsTestResultGit) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultGit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.RepositoryUrl != nil {
		toSerialize["repository_url"] = o.RepositoryUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultGit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Branch        *string                        `json:"branch,omitempty"`
		Commit        *SyntheticsTestResultGitCommit `json:"commit,omitempty"`
		RepositoryUrl *string                        `json:"repository_url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branch", "commit", "repository_url"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Branch = all.Branch
	if all.Commit != nil && all.Commit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Commit = all.Commit
	o.RepositoryUrl = all.RepositoryUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
