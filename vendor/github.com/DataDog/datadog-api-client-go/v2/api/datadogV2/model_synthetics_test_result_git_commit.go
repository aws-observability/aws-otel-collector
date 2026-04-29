// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultGitCommit Details of the Git commit associated with the test result.
type SyntheticsTestResultGitCommit struct {
	// A Git user (author or committer).
	Author *SyntheticsTestResultGitUser `json:"author,omitempty"`
	// A Git user (author or committer).
	Committer *SyntheticsTestResultGitUser `json:"committer,omitempty"`
	// Commit message.
	Message *string `json:"message,omitempty"`
	// Commit SHA.
	Sha *string `json:"sha,omitempty"`
	// URL of the commit.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultGitCommit instantiates a new SyntheticsTestResultGitCommit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultGitCommit() *SyntheticsTestResultGitCommit {
	this := SyntheticsTestResultGitCommit{}
	return &this
}

// NewSyntheticsTestResultGitCommitWithDefaults instantiates a new SyntheticsTestResultGitCommit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultGitCommitWithDefaults() *SyntheticsTestResultGitCommit {
	this := SyntheticsTestResultGitCommit{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *SyntheticsTestResultGitCommit) GetAuthor() SyntheticsTestResultGitUser {
	if o == nil || o.Author == nil {
		var ret SyntheticsTestResultGitUser
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGitCommit) GetAuthorOk() (*SyntheticsTestResultGitUser, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *SyntheticsTestResultGitCommit) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given SyntheticsTestResultGitUser and assigns it to the Author field.
func (o *SyntheticsTestResultGitCommit) SetAuthor(v SyntheticsTestResultGitUser) {
	o.Author = &v
}

// GetCommitter returns the Committer field value if set, zero value otherwise.
func (o *SyntheticsTestResultGitCommit) GetCommitter() SyntheticsTestResultGitUser {
	if o == nil || o.Committer == nil {
		var ret SyntheticsTestResultGitUser
		return ret
	}
	return *o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGitCommit) GetCommitterOk() (*SyntheticsTestResultGitUser, bool) {
	if o == nil || o.Committer == nil {
		return nil, false
	}
	return o.Committer, true
}

// HasCommitter returns a boolean if a field has been set.
func (o *SyntheticsTestResultGitCommit) HasCommitter() bool {
	return o != nil && o.Committer != nil
}

// SetCommitter gets a reference to the given SyntheticsTestResultGitUser and assigns it to the Committer field.
func (o *SyntheticsTestResultGitCommit) SetCommitter(v SyntheticsTestResultGitUser) {
	o.Committer = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsTestResultGitCommit) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGitCommit) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultGitCommit) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsTestResultGitCommit) SetMessage(v string) {
	o.Message = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *SyntheticsTestResultGitCommit) GetSha() string {
	if o == nil || o.Sha == nil {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGitCommit) GetShaOk() (*string, bool) {
	if o == nil || o.Sha == nil {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *SyntheticsTestResultGitCommit) HasSha() bool {
	return o != nil && o.Sha != nil
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *SyntheticsTestResultGitCommit) SetSha(v string) {
	o.Sha = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticsTestResultGitCommit) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGitCommit) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticsTestResultGitCommit) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SyntheticsTestResultGitCommit) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultGitCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Committer != nil {
		toSerialize["committer"] = o.Committer
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Sha != nil {
		toSerialize["sha"] = o.Sha
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultGitCommit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author    *SyntheticsTestResultGitUser `json:"author,omitempty"`
		Committer *SyntheticsTestResultGitUser `json:"committer,omitempty"`
		Message   *string                      `json:"message,omitempty"`
		Sha       *string                      `json:"sha,omitempty"`
		Url       *string                      `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "committer", "message", "sha", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Author != nil && all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = all.Author
	if all.Committer != nil && all.Committer.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Committer = all.Committer
	o.Message = all.Message
	o.Sha = all.Sha
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
