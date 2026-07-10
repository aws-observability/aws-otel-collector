// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceRepositoryInfoResponseAttributes Attributes of the service repository information.
type ServiceRepositoryInfoResponseAttributes struct {
	// The SHA of the commit associated with the service version.
	CommitSha *string `json:"commit_sha,omitempty"`
	// The URL of the source code repository.
	RepositoryUrl *string `json:"repository_url,omitempty"`
	// The status of the service repository info lookup.
	Status ServiceRepositoryInfoStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceRepositoryInfoResponseAttributes instantiates a new ServiceRepositoryInfoResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceRepositoryInfoResponseAttributes(status ServiceRepositoryInfoStatus) *ServiceRepositoryInfoResponseAttributes {
	this := ServiceRepositoryInfoResponseAttributes{}
	this.Status = status
	return &this
}

// NewServiceRepositoryInfoResponseAttributesWithDefaults instantiates a new ServiceRepositoryInfoResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceRepositoryInfoResponseAttributesWithDefaults() *ServiceRepositoryInfoResponseAttributes {
	this := ServiceRepositoryInfoResponseAttributes{}
	return &this
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise.
func (o *ServiceRepositoryInfoResponseAttributes) GetCommitSha() string {
	if o == nil || o.CommitSha == nil {
		var ret string
		return ret
	}
	return *o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRepositoryInfoResponseAttributes) GetCommitShaOk() (*string, bool) {
	if o == nil || o.CommitSha == nil {
		return nil, false
	}
	return o.CommitSha, true
}

// HasCommitSha returns a boolean if a field has been set.
func (o *ServiceRepositoryInfoResponseAttributes) HasCommitSha() bool {
	return o != nil && o.CommitSha != nil
}

// SetCommitSha gets a reference to the given string and assigns it to the CommitSha field.
func (o *ServiceRepositoryInfoResponseAttributes) SetCommitSha(v string) {
	o.CommitSha = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *ServiceRepositoryInfoResponseAttributes) GetRepositoryUrl() string {
	if o == nil || o.RepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRepositoryInfoResponseAttributes) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || o.RepositoryUrl == nil {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *ServiceRepositoryInfoResponseAttributes) HasRepositoryUrl() bool {
	return o != nil && o.RepositoryUrl != nil
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *ServiceRepositoryInfoResponseAttributes) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// GetStatus returns the Status field value.
func (o *ServiceRepositoryInfoResponseAttributes) GetStatus() ServiceRepositoryInfoStatus {
	if o == nil {
		var ret ServiceRepositoryInfoStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ServiceRepositoryInfoResponseAttributes) GetStatusOk() (*ServiceRepositoryInfoStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ServiceRepositoryInfoResponseAttributes) SetStatus(v ServiceRepositoryInfoStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceRepositoryInfoResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CommitSha != nil {
		toSerialize["commit_sha"] = o.CommitSha
	}
	if o.RepositoryUrl != nil {
		toSerialize["repository_url"] = o.RepositoryUrl
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceRepositoryInfoResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommitSha     *string                      `json:"commit_sha,omitempty"`
		RepositoryUrl *string                      `json:"repository_url,omitempty"`
		Status        *ServiceRepositoryInfoStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commit_sha", "repository_url", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CommitSha = all.CommitSha
	o.RepositoryUrl = all.RepositoryUrl
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
