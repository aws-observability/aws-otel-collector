// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BranchCoverageSummaryRequestAttributes Attributes for requesting code coverage summary for a branch.
type BranchCoverageSummaryRequestAttributes struct {
	// The branch name.
	Branch string `json:"branch"`
	// The repository identifier.
	RepositoryId string `json:"repository_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBranchCoverageSummaryRequestAttributes instantiates a new BranchCoverageSummaryRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBranchCoverageSummaryRequestAttributes(branch string, repositoryId string) *BranchCoverageSummaryRequestAttributes {
	this := BranchCoverageSummaryRequestAttributes{}
	this.Branch = branch
	this.RepositoryId = repositoryId
	return &this
}

// NewBranchCoverageSummaryRequestAttributesWithDefaults instantiates a new BranchCoverageSummaryRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBranchCoverageSummaryRequestAttributesWithDefaults() *BranchCoverageSummaryRequestAttributes {
	this := BranchCoverageSummaryRequestAttributes{}
	return &this
}

// GetBranch returns the Branch field value.
func (o *BranchCoverageSummaryRequestAttributes) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *BranchCoverageSummaryRequestAttributes) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value.
func (o *BranchCoverageSummaryRequestAttributes) SetBranch(v string) {
	o.Branch = v
}

// GetRepositoryId returns the RepositoryId field value.
func (o *BranchCoverageSummaryRequestAttributes) GetRepositoryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
func (o *BranchCoverageSummaryRequestAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryId, true
}

// SetRepositoryId sets field value.
func (o *BranchCoverageSummaryRequestAttributes) SetRepositoryId(v string) {
	o.RepositoryId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BranchCoverageSummaryRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["branch"] = o.Branch
	toSerialize["repository_id"] = o.RepositoryId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BranchCoverageSummaryRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Branch       *string `json:"branch"`
		RepositoryId *string `json:"repository_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Branch == nil {
		return fmt.Errorf("required field branch missing")
	}
	if all.RepositoryId == nil {
		return fmt.Errorf("required field repository_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branch", "repository_id"})
	} else {
		return err
	}
	o.Branch = *all.Branch
	o.RepositoryId = *all.RepositoryId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
