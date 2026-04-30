// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestBranchCreateAttributes Attributes for creating a change request branch.
type ChangeRequestBranchCreateAttributes struct {
	// The name of the branch to create.
	BranchName string `json:"branch_name"`
	// The repository identifier in the format owner/repository.
	RepoId string `json:"repo_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestBranchCreateAttributes instantiates a new ChangeRequestBranchCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestBranchCreateAttributes(branchName string, repoId string) *ChangeRequestBranchCreateAttributes {
	this := ChangeRequestBranchCreateAttributes{}
	this.BranchName = branchName
	this.RepoId = repoId
	return &this
}

// NewChangeRequestBranchCreateAttributesWithDefaults instantiates a new ChangeRequestBranchCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestBranchCreateAttributesWithDefaults() *ChangeRequestBranchCreateAttributes {
	this := ChangeRequestBranchCreateAttributes{}
	return &this
}

// GetBranchName returns the BranchName field value.
func (o *ChangeRequestBranchCreateAttributes) GetBranchName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestBranchCreateAttributes) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchName, true
}

// SetBranchName sets field value.
func (o *ChangeRequestBranchCreateAttributes) SetBranchName(v string) {
	o.BranchName = v
}

// GetRepoId returns the RepoId field value.
func (o *ChangeRequestBranchCreateAttributes) GetRepoId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestBranchCreateAttributes) GetRepoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepoId, true
}

// SetRepoId sets field value.
func (o *ChangeRequestBranchCreateAttributes) SetRepoId(v string) {
	o.RepoId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestBranchCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["branch_name"] = o.BranchName
	toSerialize["repo_id"] = o.RepoId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestBranchCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BranchName *string `json:"branch_name"`
		RepoId     *string `json:"repo_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BranchName == nil {
		return fmt.Errorf("required field branch_name missing")
	}
	if all.RepoId == nil {
		return fmt.Errorf("required field repo_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branch_name", "repo_id"})
	} else {
		return err
	}
	o.BranchName = *all.BranchName
	o.RepoId = *all.RepoId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
