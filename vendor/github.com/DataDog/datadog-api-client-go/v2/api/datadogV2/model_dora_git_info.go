// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORAGitInfo Git info for DORA Metrics events.
type DORAGitInfo struct {
	// Git Commit SHA.
	CommitSha string `json:"commit_sha"`
	// Git Repository URL
	RepositoryUrl string `json:"repository_url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORAGitInfo instantiates a new DORAGitInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORAGitInfo(commitSha string, repositoryUrl string) *DORAGitInfo {
	this := DORAGitInfo{}
	this.CommitSha = commitSha
	this.RepositoryUrl = repositoryUrl
	return &this
}

// NewDORAGitInfoWithDefaults instantiates a new DORAGitInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORAGitInfoWithDefaults() *DORAGitInfo {
	this := DORAGitInfo{}
	return &this
}

// GetCommitSha returns the CommitSha field value.
func (o *DORAGitInfo) GetCommitSha() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value
// and a boolean to check if the value has been set.
func (o *DORAGitInfo) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitSha, true
}

// SetCommitSha sets field value.
func (o *DORAGitInfo) SetCommitSha(v string) {
	o.CommitSha = v
}

// GetRepositoryUrl returns the RepositoryUrl field value.
func (o *DORAGitInfo) GetRepositoryUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value
// and a boolean to check if the value has been set.
func (o *DORAGitInfo) GetRepositoryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryUrl, true
}

// SetRepositoryUrl sets field value.
func (o *DORAGitInfo) SetRepositoryUrl(v string) {
	o.RepositoryUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORAGitInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["commit_sha"] = o.CommitSha
	toSerialize["repository_url"] = o.RepositoryUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORAGitInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommitSha     *string `json:"commit_sha"`
		RepositoryUrl *string `json:"repository_url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CommitSha == nil {
		return fmt.Errorf("required field commit_sha missing")
	}
	if all.RepositoryUrl == nil {
		return fmt.Errorf("required field repository_url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commit_sha", "repository_url"})
	} else {
		return err
	}
	o.CommitSha = *all.CommitSha
	o.RepositoryUrl = *all.RepositoryUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
