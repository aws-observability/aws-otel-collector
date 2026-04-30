// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesEvaluationResponseAttributes Attributes for a deployment gate evaluation response.
type DeploymentGatesEvaluationResponseAttributes struct {
	// The unique identifier of the gate evaluation.
	EvaluationId string `json:"evaluation_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGatesEvaluationResponseAttributes instantiates a new DeploymentGatesEvaluationResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGatesEvaluationResponseAttributes(evaluationId string) *DeploymentGatesEvaluationResponseAttributes {
	this := DeploymentGatesEvaluationResponseAttributes{}
	this.EvaluationId = evaluationId
	return &this
}

// NewDeploymentGatesEvaluationResponseAttributesWithDefaults instantiates a new DeploymentGatesEvaluationResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGatesEvaluationResponseAttributesWithDefaults() *DeploymentGatesEvaluationResponseAttributes {
	this := DeploymentGatesEvaluationResponseAttributes{}
	return &this
}

// GetEvaluationId returns the EvaluationId field value.
func (o *DeploymentGatesEvaluationResponseAttributes) GetEvaluationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvaluationId
}

// GetEvaluationIdOk returns a tuple with the EvaluationId field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationResponseAttributes) GetEvaluationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationId, true
}

// SetEvaluationId sets field value.
func (o *DeploymentGatesEvaluationResponseAttributes) SetEvaluationId(v string) {
	o.EvaluationId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGatesEvaluationResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["evaluation_id"] = o.EvaluationId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGatesEvaluationResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EvaluationId *string `json:"evaluation_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EvaluationId == nil {
		return fmt.Errorf("required field evaluation_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"evaluation_id"})
	} else {
		return err
	}
	o.EvaluationId = *all.EvaluationId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
