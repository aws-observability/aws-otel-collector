// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationFlakyTestsManagementPoliciesAttributes Attributes of the Flaky Tests Management policies for a repository.
type TestOptimizationFlakyTestsManagementPoliciesAttributes struct {
	// Configuration for the attempt-to-fix Flaky Tests Management policy.
	AttemptToFix *TestOptimizationFlakyTestsManagementPoliciesAttemptToFix `json:"attempt_to_fix,omitempty"`
	// Configuration for the disabled Flaky Tests Management policy.
	Disabled *TestOptimizationFlakyTestsManagementPoliciesDisabled `json:"disabled,omitempty"`
	// Configuration for the quarantined Flaky Tests Management policy.
	Quarantined *TestOptimizationFlakyTestsManagementPoliciesQuarantined `json:"quarantined,omitempty"`
	// The repository identifier.
	RepositoryId *string `json:"repository_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationFlakyTestsManagementPoliciesAttributes instantiates a new TestOptimizationFlakyTestsManagementPoliciesAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationFlakyTestsManagementPoliciesAttributes() *TestOptimizationFlakyTestsManagementPoliciesAttributes {
	this := TestOptimizationFlakyTestsManagementPoliciesAttributes{}
	return &this
}

// NewTestOptimizationFlakyTestsManagementPoliciesAttributesWithDefaults instantiates a new TestOptimizationFlakyTestsManagementPoliciesAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationFlakyTestsManagementPoliciesAttributesWithDefaults() *TestOptimizationFlakyTestsManagementPoliciesAttributes {
	this := TestOptimizationFlakyTestsManagementPoliciesAttributes{}
	return &this
}

// GetAttemptToFix returns the AttemptToFix field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) GetAttemptToFix() TestOptimizationFlakyTestsManagementPoliciesAttemptToFix {
	if o == nil || o.AttemptToFix == nil {
		var ret TestOptimizationFlakyTestsManagementPoliciesAttemptToFix
		return ret
	}
	return *o.AttemptToFix
}

// GetAttemptToFixOk returns a tuple with the AttemptToFix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) GetAttemptToFixOk() (*TestOptimizationFlakyTestsManagementPoliciesAttemptToFix, bool) {
	if o == nil || o.AttemptToFix == nil {
		return nil, false
	}
	return o.AttemptToFix, true
}

// HasAttemptToFix returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) HasAttemptToFix() bool {
	return o != nil && o.AttemptToFix != nil
}

// SetAttemptToFix gets a reference to the given TestOptimizationFlakyTestsManagementPoliciesAttemptToFix and assigns it to the AttemptToFix field.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) SetAttemptToFix(v TestOptimizationFlakyTestsManagementPoliciesAttemptToFix) {
	o.AttemptToFix = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) GetDisabled() TestOptimizationFlakyTestsManagementPoliciesDisabled {
	if o == nil || o.Disabled == nil {
		var ret TestOptimizationFlakyTestsManagementPoliciesDisabled
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) GetDisabledOk() (*TestOptimizationFlakyTestsManagementPoliciesDisabled, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given TestOptimizationFlakyTestsManagementPoliciesDisabled and assigns it to the Disabled field.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) SetDisabled(v TestOptimizationFlakyTestsManagementPoliciesDisabled) {
	o.Disabled = &v
}

// GetQuarantined returns the Quarantined field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) GetQuarantined() TestOptimizationFlakyTestsManagementPoliciesQuarantined {
	if o == nil || o.Quarantined == nil {
		var ret TestOptimizationFlakyTestsManagementPoliciesQuarantined
		return ret
	}
	return *o.Quarantined
}

// GetQuarantinedOk returns a tuple with the Quarantined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) GetQuarantinedOk() (*TestOptimizationFlakyTestsManagementPoliciesQuarantined, bool) {
	if o == nil || o.Quarantined == nil {
		return nil, false
	}
	return o.Quarantined, true
}

// HasQuarantined returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) HasQuarantined() bool {
	return o != nil && o.Quarantined != nil
}

// SetQuarantined gets a reference to the given TestOptimizationFlakyTestsManagementPoliciesQuarantined and assigns it to the Quarantined field.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) SetQuarantined(v TestOptimizationFlakyTestsManagementPoliciesQuarantined) {
	o.Quarantined = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) GetRepositoryId() string {
	if o == nil || o.RepositoryId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) HasRepositoryId() bool {
	return o != nil && o.RepositoryId != nil
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationFlakyTestsManagementPoliciesAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AttemptToFix != nil {
		toSerialize["attempt_to_fix"] = o.AttemptToFix
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Quarantined != nil {
		toSerialize["quarantined"] = o.Quarantined
	}
	if o.RepositoryId != nil {
		toSerialize["repository_id"] = o.RepositoryId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationFlakyTestsManagementPoliciesAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AttemptToFix *TestOptimizationFlakyTestsManagementPoliciesAttemptToFix `json:"attempt_to_fix,omitempty"`
		Disabled     *TestOptimizationFlakyTestsManagementPoliciesDisabled     `json:"disabled,omitempty"`
		Quarantined  *TestOptimizationFlakyTestsManagementPoliciesQuarantined  `json:"quarantined,omitempty"`
		RepositoryId *string                                                   `json:"repository_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attempt_to_fix", "disabled", "quarantined", "repository_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AttemptToFix != nil && all.AttemptToFix.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AttemptToFix = all.AttemptToFix
	if all.Disabled != nil && all.Disabled.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Disabled = all.Disabled
	if all.Quarantined != nil && all.Quarantined.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Quarantined = all.Quarantined
	o.RepositoryId = all.RepositoryId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
