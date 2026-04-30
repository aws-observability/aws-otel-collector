// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentPatchRequestAttributes Attributes for patching a DORA deployment event.
type DORADeploymentPatchRequestAttributes struct {
	// Indicates whether the deployment resulted in a change failure.
	ChangeFailure *bool `json:"change_failure,omitempty"`
	// Remediation details for the deployment. Optional, but required to calculate failed deployment recovery time.
	Remediation *DORADeploymentPatchRemediation `json:"remediation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORADeploymentPatchRequestAttributes instantiates a new DORADeploymentPatchRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORADeploymentPatchRequestAttributes() *DORADeploymentPatchRequestAttributes {
	this := DORADeploymentPatchRequestAttributes{}
	return &this
}

// NewDORADeploymentPatchRequestAttributesWithDefaults instantiates a new DORADeploymentPatchRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORADeploymentPatchRequestAttributesWithDefaults() *DORADeploymentPatchRequestAttributes {
	this := DORADeploymentPatchRequestAttributes{}
	return &this
}

// GetChangeFailure returns the ChangeFailure field value if set, zero value otherwise.
func (o *DORADeploymentPatchRequestAttributes) GetChangeFailure() bool {
	if o == nil || o.ChangeFailure == nil {
		var ret bool
		return ret
	}
	return *o.ChangeFailure
}

// GetChangeFailureOk returns a tuple with the ChangeFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchRequestAttributes) GetChangeFailureOk() (*bool, bool) {
	if o == nil || o.ChangeFailure == nil {
		return nil, false
	}
	return o.ChangeFailure, true
}

// HasChangeFailure returns a boolean if a field has been set.
func (o *DORADeploymentPatchRequestAttributes) HasChangeFailure() bool {
	return o != nil && o.ChangeFailure != nil
}

// SetChangeFailure gets a reference to the given bool and assigns it to the ChangeFailure field.
func (o *DORADeploymentPatchRequestAttributes) SetChangeFailure(v bool) {
	o.ChangeFailure = &v
}

// GetRemediation returns the Remediation field value if set, zero value otherwise.
func (o *DORADeploymentPatchRequestAttributes) GetRemediation() DORADeploymentPatchRemediation {
	if o == nil || o.Remediation == nil {
		var ret DORADeploymentPatchRemediation
		return ret
	}
	return *o.Remediation
}

// GetRemediationOk returns a tuple with the Remediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchRequestAttributes) GetRemediationOk() (*DORADeploymentPatchRemediation, bool) {
	if o == nil || o.Remediation == nil {
		return nil, false
	}
	return o.Remediation, true
}

// HasRemediation returns a boolean if a field has been set.
func (o *DORADeploymentPatchRequestAttributes) HasRemediation() bool {
	return o != nil && o.Remediation != nil
}

// SetRemediation gets a reference to the given DORADeploymentPatchRemediation and assigns it to the Remediation field.
func (o *DORADeploymentPatchRequestAttributes) SetRemediation(v DORADeploymentPatchRemediation) {
	o.Remediation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORADeploymentPatchRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChangeFailure != nil {
		toSerialize["change_failure"] = o.ChangeFailure
	}
	if o.Remediation != nil {
		toSerialize["remediation"] = o.Remediation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORADeploymentPatchRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeFailure *bool                           `json:"change_failure,omitempty"`
		Remediation   *DORADeploymentPatchRemediation `json:"remediation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_failure", "remediation"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChangeFailure = all.ChangeFailure
	if all.Remediation != nil && all.Remediation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Remediation = all.Remediation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
