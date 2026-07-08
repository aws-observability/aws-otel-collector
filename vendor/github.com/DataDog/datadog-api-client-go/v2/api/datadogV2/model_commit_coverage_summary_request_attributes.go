// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitCoverageSummaryRequestAttributes Attributes for requesting code coverage summary for a commit.
type CommitCoverageSummaryRequestAttributes struct {
	// The commit SHA (40-character hexadecimal string).
	CommitSha string `json:"commit_sha"`
	// Deprecated: use `repository_url` instead. The repository URL.
	// Deprecated
	RepositoryId *string `json:"repository_id,omitempty"`
	// The repository URL. Accepts a full URL with or without a scheme (for example, `https://github.com/org/repo` or `github.com/org/repo`).
	RepositoryUrl *string `json:"repository_url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitCoverageSummaryRequestAttributes instantiates a new CommitCoverageSummaryRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitCoverageSummaryRequestAttributes(commitSha string) *CommitCoverageSummaryRequestAttributes {
	this := CommitCoverageSummaryRequestAttributes{}
	this.CommitSha = commitSha
	return &this
}

// NewCommitCoverageSummaryRequestAttributesWithDefaults instantiates a new CommitCoverageSummaryRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitCoverageSummaryRequestAttributesWithDefaults() *CommitCoverageSummaryRequestAttributes {
	this := CommitCoverageSummaryRequestAttributes{}
	return &this
}

// GetCommitSha returns the CommitSha field value.
func (o *CommitCoverageSummaryRequestAttributes) GetCommitSha() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value
// and a boolean to check if the value has been set.
func (o *CommitCoverageSummaryRequestAttributes) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitSha, true
}

// SetCommitSha sets field value.
func (o *CommitCoverageSummaryRequestAttributes) SetCommitSha(v string) {
	o.CommitSha = v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
// Deprecated
func (o *CommitCoverageSummaryRequestAttributes) GetRepositoryId() string {
	if o == nil || o.RepositoryId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CommitCoverageSummaryRequestAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *CommitCoverageSummaryRequestAttributes) HasRepositoryId() bool {
	return o != nil && o.RepositoryId != nil
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
// Deprecated
func (o *CommitCoverageSummaryRequestAttributes) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *CommitCoverageSummaryRequestAttributes) GetRepositoryUrl() string {
	if o == nil || o.RepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitCoverageSummaryRequestAttributes) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || o.RepositoryUrl == nil {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *CommitCoverageSummaryRequestAttributes) HasRepositoryUrl() bool {
	return o != nil && o.RepositoryUrl != nil
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *CommitCoverageSummaryRequestAttributes) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitCoverageSummaryRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["commit_sha"] = o.CommitSha
	if o.RepositoryId != nil {
		toSerialize["repository_id"] = o.RepositoryId
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
func (o *CommitCoverageSummaryRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommitSha     *string `json:"commit_sha"`
		RepositoryId  *string `json:"repository_id,omitempty"`
		RepositoryUrl *string `json:"repository_url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CommitSha == nil {
		return fmt.Errorf("required field commit_sha missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commit_sha", "repository_id", "repository_url"})
	} else {
		return err
	}
	o.CommitSha = *all.CommitSha
	o.RepositoryId = all.RepositoryId
	o.RepositoryUrl = all.RepositoryUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
