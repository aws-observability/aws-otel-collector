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
	// The repository identifier.
	RepositoryId string `json:"repository_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitCoverageSummaryRequestAttributes instantiates a new CommitCoverageSummaryRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitCoverageSummaryRequestAttributes(commitSha string, repositoryId string) *CommitCoverageSummaryRequestAttributes {
	this := CommitCoverageSummaryRequestAttributes{}
	this.CommitSha = commitSha
	this.RepositoryId = repositoryId
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

// GetRepositoryId returns the RepositoryId field value.
func (o *CommitCoverageSummaryRequestAttributes) GetRepositoryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
func (o *CommitCoverageSummaryRequestAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryId, true
}

// SetRepositoryId sets field value.
func (o *CommitCoverageSummaryRequestAttributes) SetRepositoryId(v string) {
	o.RepositoryId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitCoverageSummaryRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["commit_sha"] = o.CommitSha
	toSerialize["repository_id"] = o.RepositoryId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitCoverageSummaryRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommitSha    *string `json:"commit_sha"`
		RepositoryId *string `json:"repository_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CommitSha == nil {
		return fmt.Errorf("required field commit_sha missing")
	}
	if all.RepositoryId == nil {
		return fmt.Errorf("required field repository_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commit_sha", "repository_id"})
	} else {
		return err
	}
	o.CommitSha = *all.CommitSha
	o.RepositoryId = *all.RepositoryId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
