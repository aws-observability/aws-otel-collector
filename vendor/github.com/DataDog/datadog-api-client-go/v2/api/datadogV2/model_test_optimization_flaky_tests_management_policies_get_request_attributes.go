// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes Attributes for requesting Flaky Tests Management policies.
type TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes struct {
	// The repository identifier.
	RepositoryId string `json:"repository_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes instantiates a new TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes(repositoryId string) *TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes {
	this := TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes{}
	this.RepositoryId = repositoryId
	return &this
}

// NewTestOptimizationFlakyTestsManagementPoliciesGetRequestAttributesWithDefaults instantiates a new TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationFlakyTestsManagementPoliciesGetRequestAttributesWithDefaults() *TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes {
	this := TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes{}
	return &this
}

// GetRepositoryId returns the RepositoryId field value.
func (o *TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes) GetRepositoryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryId, true
}

// SetRepositoryId sets field value.
func (o *TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes) SetRepositoryId(v string) {
	o.RepositoryId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["repository_id"] = o.RepositoryId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RepositoryId *string `json:"repository_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RepositoryId == nil {
		return fmt.Errorf("required field repository_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"repository_id"})
	} else {
		return err
	}
	o.RepositoryId = *all.RepositoryId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
