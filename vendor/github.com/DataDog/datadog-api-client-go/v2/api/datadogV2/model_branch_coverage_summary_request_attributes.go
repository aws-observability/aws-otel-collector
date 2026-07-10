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
	// Deprecated: use `repository_url` instead. The repository URL.
	// Deprecated
	RepositoryId *string `json:"repository_id,omitempty"`
	// The repository URL. Accepts a full URL with or without a scheme (for example, `https://github.com/org/repo` or `github.com/org/repo`).
	RepositoryUrl *string `json:"repository_url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBranchCoverageSummaryRequestAttributes instantiates a new BranchCoverageSummaryRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBranchCoverageSummaryRequestAttributes(branch string) *BranchCoverageSummaryRequestAttributes {
	this := BranchCoverageSummaryRequestAttributes{}
	this.Branch = branch
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

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
// Deprecated
func (o *BranchCoverageSummaryRequestAttributes) GetRepositoryId() string {
	if o == nil || o.RepositoryId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BranchCoverageSummaryRequestAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *BranchCoverageSummaryRequestAttributes) HasRepositoryId() bool {
	return o != nil && o.RepositoryId != nil
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
// Deprecated
func (o *BranchCoverageSummaryRequestAttributes) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *BranchCoverageSummaryRequestAttributes) GetRepositoryUrl() string {
	if o == nil || o.RepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchCoverageSummaryRequestAttributes) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || o.RepositoryUrl == nil {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *BranchCoverageSummaryRequestAttributes) HasRepositoryUrl() bool {
	return o != nil && o.RepositoryUrl != nil
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *BranchCoverageSummaryRequestAttributes) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BranchCoverageSummaryRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["branch"] = o.Branch
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
func (o *BranchCoverageSummaryRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Branch        *string `json:"branch"`
		RepositoryId  *string `json:"repository_id,omitempty"`
		RepositoryUrl *string `json:"repository_url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Branch == nil {
		return fmt.Errorf("required field branch missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branch", "repository_id", "repository_url"})
	} else {
		return err
	}
	o.Branch = *all.Branch
	o.RepositoryId = all.RepositoryId
	o.RepositoryUrl = all.RepositoryUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
