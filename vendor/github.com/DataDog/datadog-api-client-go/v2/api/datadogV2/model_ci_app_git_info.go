// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppGitInfo If pipelines are triggered due to actions to a Git repository, then all payloads must contain this.
// Note that either `tag` or `branch` has to be provided, but not both.
type CIAppGitInfo struct {
	// The commit author email.
	AuthorEmail string `json:"author_email"`
	// The commit author name.
	AuthorName datadog.NullableString `json:"author_name,omitempty"`
	// The commit author timestamp in RFC3339 format.
	AuthorTime datadog.NullableString `json:"author_time,omitempty"`
	// The branch name (if a tag use the tag parameter).
	Branch datadog.NullableString `json:"branch,omitempty"`
	// The commit timestamp in RFC3339 format.
	CommitTime datadog.NullableString `json:"commit_time,omitempty"`
	// The committer email.
	CommitterEmail datadog.NullableString `json:"committer_email,omitempty"`
	// The committer name.
	CommitterName datadog.NullableString `json:"committer_name,omitempty"`
	// The Git repository's default branch.
	DefaultBranch datadog.NullableString `json:"default_branch,omitempty"`
	// The commit message.
	Message datadog.NullableString `json:"message,omitempty"`
	// The URL of the repository.
	RepositoryUrl string `json:"repository_url"`
	// The git commit SHA.
	Sha string `json:"sha"`
	// The tag name (if a branch use the branch parameter).
	Tag datadog.NullableString `json:"tag,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppGitInfo instantiates a new CIAppGitInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppGitInfo(authorEmail string, repositoryUrl string, sha string) *CIAppGitInfo {
	this := CIAppGitInfo{}
	this.AuthorEmail = authorEmail
	this.RepositoryUrl = repositoryUrl
	this.Sha = sha
	return &this
}

// NewCIAppGitInfoWithDefaults instantiates a new CIAppGitInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppGitInfoWithDefaults() *CIAppGitInfo {
	this := CIAppGitInfo{}
	return &this
}

// GetAuthorEmail returns the AuthorEmail field value.
func (o *CIAppGitInfo) GetAuthorEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AuthorEmail
}

// GetAuthorEmailOk returns a tuple with the AuthorEmail field value
// and a boolean to check if the value has been set.
func (o *CIAppGitInfo) GetAuthorEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorEmail, true
}

// SetAuthorEmail sets field value.
func (o *CIAppGitInfo) SetAuthorEmail(v string) {
	o.AuthorEmail = v
}

// GetAuthorName returns the AuthorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppGitInfo) GetAuthorName() string {
	if o == nil || o.AuthorName.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorName.Get()
}

// GetAuthorNameOk returns a tuple with the AuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppGitInfo) GetAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorName.Get(), o.AuthorName.IsSet()
}

// HasAuthorName returns a boolean if a field has been set.
func (o *CIAppGitInfo) HasAuthorName() bool {
	return o != nil && o.AuthorName.IsSet()
}

// SetAuthorName gets a reference to the given datadog.NullableString and assigns it to the AuthorName field.
func (o *CIAppGitInfo) SetAuthorName(v string) {
	o.AuthorName.Set(&v)
}

// SetAuthorNameNil sets the value for AuthorName to be an explicit nil.
func (o *CIAppGitInfo) SetAuthorNameNil() {
	o.AuthorName.Set(nil)
}

// UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil.
func (o *CIAppGitInfo) UnsetAuthorName() {
	o.AuthorName.Unset()
}

// GetAuthorTime returns the AuthorTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppGitInfo) GetAuthorTime() string {
	if o == nil || o.AuthorTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorTime.Get()
}

// GetAuthorTimeOk returns a tuple with the AuthorTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppGitInfo) GetAuthorTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorTime.Get(), o.AuthorTime.IsSet()
}

// HasAuthorTime returns a boolean if a field has been set.
func (o *CIAppGitInfo) HasAuthorTime() bool {
	return o != nil && o.AuthorTime.IsSet()
}

// SetAuthorTime gets a reference to the given datadog.NullableString and assigns it to the AuthorTime field.
func (o *CIAppGitInfo) SetAuthorTime(v string) {
	o.AuthorTime.Set(&v)
}

// SetAuthorTimeNil sets the value for AuthorTime to be an explicit nil.
func (o *CIAppGitInfo) SetAuthorTimeNil() {
	o.AuthorTime.Set(nil)
}

// UnsetAuthorTime ensures that no value is present for AuthorTime, not even an explicit nil.
func (o *CIAppGitInfo) UnsetAuthorTime() {
	o.AuthorTime.Unset()
}

// GetBranch returns the Branch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppGitInfo) GetBranch() string {
	if o == nil || o.Branch.Get() == nil {
		var ret string
		return ret
	}
	return *o.Branch.Get()
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppGitInfo) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branch.Get(), o.Branch.IsSet()
}

// HasBranch returns a boolean if a field has been set.
func (o *CIAppGitInfo) HasBranch() bool {
	return o != nil && o.Branch.IsSet()
}

// SetBranch gets a reference to the given datadog.NullableString and assigns it to the Branch field.
func (o *CIAppGitInfo) SetBranch(v string) {
	o.Branch.Set(&v)
}

// SetBranchNil sets the value for Branch to be an explicit nil.
func (o *CIAppGitInfo) SetBranchNil() {
	o.Branch.Set(nil)
}

// UnsetBranch ensures that no value is present for Branch, not even an explicit nil.
func (o *CIAppGitInfo) UnsetBranch() {
	o.Branch.Unset()
}

// GetCommitTime returns the CommitTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppGitInfo) GetCommitTime() string {
	if o == nil || o.CommitTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.CommitTime.Get()
}

// GetCommitTimeOk returns a tuple with the CommitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppGitInfo) GetCommitTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitTime.Get(), o.CommitTime.IsSet()
}

// HasCommitTime returns a boolean if a field has been set.
func (o *CIAppGitInfo) HasCommitTime() bool {
	return o != nil && o.CommitTime.IsSet()
}

// SetCommitTime gets a reference to the given datadog.NullableString and assigns it to the CommitTime field.
func (o *CIAppGitInfo) SetCommitTime(v string) {
	o.CommitTime.Set(&v)
}

// SetCommitTimeNil sets the value for CommitTime to be an explicit nil.
func (o *CIAppGitInfo) SetCommitTimeNil() {
	o.CommitTime.Set(nil)
}

// UnsetCommitTime ensures that no value is present for CommitTime, not even an explicit nil.
func (o *CIAppGitInfo) UnsetCommitTime() {
	o.CommitTime.Unset()
}

// GetCommitterEmail returns the CommitterEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppGitInfo) GetCommitterEmail() string {
	if o == nil || o.CommitterEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.CommitterEmail.Get()
}

// GetCommitterEmailOk returns a tuple with the CommitterEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppGitInfo) GetCommitterEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitterEmail.Get(), o.CommitterEmail.IsSet()
}

// HasCommitterEmail returns a boolean if a field has been set.
func (o *CIAppGitInfo) HasCommitterEmail() bool {
	return o != nil && o.CommitterEmail.IsSet()
}

// SetCommitterEmail gets a reference to the given datadog.NullableString and assigns it to the CommitterEmail field.
func (o *CIAppGitInfo) SetCommitterEmail(v string) {
	o.CommitterEmail.Set(&v)
}

// SetCommitterEmailNil sets the value for CommitterEmail to be an explicit nil.
func (o *CIAppGitInfo) SetCommitterEmailNil() {
	o.CommitterEmail.Set(nil)
}

// UnsetCommitterEmail ensures that no value is present for CommitterEmail, not even an explicit nil.
func (o *CIAppGitInfo) UnsetCommitterEmail() {
	o.CommitterEmail.Unset()
}

// GetCommitterName returns the CommitterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppGitInfo) GetCommitterName() string {
	if o == nil || o.CommitterName.Get() == nil {
		var ret string
		return ret
	}
	return *o.CommitterName.Get()
}

// GetCommitterNameOk returns a tuple with the CommitterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppGitInfo) GetCommitterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitterName.Get(), o.CommitterName.IsSet()
}

// HasCommitterName returns a boolean if a field has been set.
func (o *CIAppGitInfo) HasCommitterName() bool {
	return o != nil && o.CommitterName.IsSet()
}

// SetCommitterName gets a reference to the given datadog.NullableString and assigns it to the CommitterName field.
func (o *CIAppGitInfo) SetCommitterName(v string) {
	o.CommitterName.Set(&v)
}

// SetCommitterNameNil sets the value for CommitterName to be an explicit nil.
func (o *CIAppGitInfo) SetCommitterNameNil() {
	o.CommitterName.Set(nil)
}

// UnsetCommitterName ensures that no value is present for CommitterName, not even an explicit nil.
func (o *CIAppGitInfo) UnsetCommitterName() {
	o.CommitterName.Unset()
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppGitInfo) GetDefaultBranch() string {
	if o == nil || o.DefaultBranch.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultBranch.Get()
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppGitInfo) GetDefaultBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultBranch.Get(), o.DefaultBranch.IsSet()
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *CIAppGitInfo) HasDefaultBranch() bool {
	return o != nil && o.DefaultBranch.IsSet()
}

// SetDefaultBranch gets a reference to the given datadog.NullableString and assigns it to the DefaultBranch field.
func (o *CIAppGitInfo) SetDefaultBranch(v string) {
	o.DefaultBranch.Set(&v)
}

// SetDefaultBranchNil sets the value for DefaultBranch to be an explicit nil.
func (o *CIAppGitInfo) SetDefaultBranchNil() {
	o.DefaultBranch.Set(nil)
}

// UnsetDefaultBranch ensures that no value is present for DefaultBranch, not even an explicit nil.
func (o *CIAppGitInfo) UnsetDefaultBranch() {
	o.DefaultBranch.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppGitInfo) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppGitInfo) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *CIAppGitInfo) HasMessage() bool {
	return o != nil && o.Message.IsSet()
}

// SetMessage gets a reference to the given datadog.NullableString and assigns it to the Message field.
func (o *CIAppGitInfo) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil.
func (o *CIAppGitInfo) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil.
func (o *CIAppGitInfo) UnsetMessage() {
	o.Message.Unset()
}

// GetRepositoryUrl returns the RepositoryUrl field value.
func (o *CIAppGitInfo) GetRepositoryUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value
// and a boolean to check if the value has been set.
func (o *CIAppGitInfo) GetRepositoryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryUrl, true
}

// SetRepositoryUrl sets field value.
func (o *CIAppGitInfo) SetRepositoryUrl(v string) {
	o.RepositoryUrl = v
}

// GetSha returns the Sha field value.
func (o *CIAppGitInfo) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *CIAppGitInfo) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value.
func (o *CIAppGitInfo) SetSha(v string) {
	o.Sha = v
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CIAppGitInfo) GetTag() string {
	if o == nil || o.Tag.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CIAppGitInfo) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *CIAppGitInfo) HasTag() bool {
	return o != nil && o.Tag.IsSet()
}

// SetTag gets a reference to the given datadog.NullableString and assigns it to the Tag field.
func (o *CIAppGitInfo) SetTag(v string) {
	o.Tag.Set(&v)
}

// SetTagNil sets the value for Tag to be an explicit nil.
func (o *CIAppGitInfo) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil.
func (o *CIAppGitInfo) UnsetTag() {
	o.Tag.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppGitInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["author_email"] = o.AuthorEmail
	if o.AuthorName.IsSet() {
		toSerialize["author_name"] = o.AuthorName.Get()
	}
	if o.AuthorTime.IsSet() {
		toSerialize["author_time"] = o.AuthorTime.Get()
	}
	if o.Branch.IsSet() {
		toSerialize["branch"] = o.Branch.Get()
	}
	if o.CommitTime.IsSet() {
		toSerialize["commit_time"] = o.CommitTime.Get()
	}
	if o.CommitterEmail.IsSet() {
		toSerialize["committer_email"] = o.CommitterEmail.Get()
	}
	if o.CommitterName.IsSet() {
		toSerialize["committer_name"] = o.CommitterName.Get()
	}
	if o.DefaultBranch.IsSet() {
		toSerialize["default_branch"] = o.DefaultBranch.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	toSerialize["repository_url"] = o.RepositoryUrl
	toSerialize["sha"] = o.Sha
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppGitInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthorEmail    *string                `json:"author_email"`
		AuthorName     datadog.NullableString `json:"author_name,omitempty"`
		AuthorTime     datadog.NullableString `json:"author_time,omitempty"`
		Branch         datadog.NullableString `json:"branch,omitempty"`
		CommitTime     datadog.NullableString `json:"commit_time,omitempty"`
		CommitterEmail datadog.NullableString `json:"committer_email,omitempty"`
		CommitterName  datadog.NullableString `json:"committer_name,omitempty"`
		DefaultBranch  datadog.NullableString `json:"default_branch,omitempty"`
		Message        datadog.NullableString `json:"message,omitempty"`
		RepositoryUrl  *string                `json:"repository_url"`
		Sha            *string                `json:"sha"`
		Tag            datadog.NullableString `json:"tag,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthorEmail == nil {
		return fmt.Errorf("required field author_email missing")
	}
	if all.RepositoryUrl == nil {
		return fmt.Errorf("required field repository_url missing")
	}
	if all.Sha == nil {
		return fmt.Errorf("required field sha missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author_email", "author_name", "author_time", "branch", "commit_time", "committer_email", "committer_name", "default_branch", "message", "repository_url", "sha", "tag"})
	} else {
		return err
	}
	o.AuthorEmail = *all.AuthorEmail
	o.AuthorName = all.AuthorName
	o.AuthorTime = all.AuthorTime
	o.Branch = all.Branch
	o.CommitTime = all.CommitTime
	o.CommitterEmail = all.CommitterEmail
	o.CommitterName = all.CommitterName
	o.DefaultBranch = all.DefaultBranch
	o.Message = all.Message
	o.RepositoryUrl = *all.RepositoryUrl
	o.Sha = *all.Sha
	o.Tag = all.Tag

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableCIAppGitInfo handles when a null is used for CIAppGitInfo.
type NullableCIAppGitInfo struct {
	value *CIAppGitInfo
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppGitInfo) Get() *CIAppGitInfo {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppGitInfo) Set(val *CIAppGitInfo) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppGitInfo) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableCIAppGitInfo) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppGitInfo initializes the struct as if Set has been called.
func NewNullableCIAppGitInfo(val *CIAppGitInfo) *NullableCIAppGitInfo {
	return &NullableCIAppGitInfo{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppGitInfo) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppGitInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
