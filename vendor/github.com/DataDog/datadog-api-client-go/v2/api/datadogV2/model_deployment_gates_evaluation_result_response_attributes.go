// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesEvaluationResultResponseAttributes Attributes for a deployment gate evaluation result response.
type DeploymentGatesEvaluationResultResponseAttributes struct {
	// Whether the gate was evaluated in dry-run mode.
	DryRun bool `json:"dry_run"`
	// The unique identifier of the gate evaluation.
	EvaluationId string `json:"evaluation_id"`
	// A URL to view the evaluation details in the Datadog UI.
	EvaluationUrl string `json:"evaluation_url"`
	// The unique identifier of the deployment gate.
	GateId uuid.UUID `json:"gate_id"`
	// The overall status of the gate evaluation.
	// - `in_progress`: The evaluation is still running.
	// - `pass`: All rules passed successfully and the deployment is allowed to proceed.
	// - `fail`: One or more rules did not pass; the deployment should not proceed.
	GateStatus DeploymentGatesEvaluationResultResponseAttributesGateStatus `json:"gate_status"`
	// The results of individual rule evaluations.
	Rules []DeploymentGatesRuleResponse `json:"rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGatesEvaluationResultResponseAttributes instantiates a new DeploymentGatesEvaluationResultResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGatesEvaluationResultResponseAttributes(dryRun bool, evaluationId string, evaluationUrl string, gateId uuid.UUID, gateStatus DeploymentGatesEvaluationResultResponseAttributesGateStatus, rules []DeploymentGatesRuleResponse) *DeploymentGatesEvaluationResultResponseAttributes {
	this := DeploymentGatesEvaluationResultResponseAttributes{}
	this.DryRun = dryRun
	this.EvaluationId = evaluationId
	this.EvaluationUrl = evaluationUrl
	this.GateId = gateId
	this.GateStatus = gateStatus
	this.Rules = rules
	return &this
}

// NewDeploymentGatesEvaluationResultResponseAttributesWithDefaults instantiates a new DeploymentGatesEvaluationResultResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGatesEvaluationResultResponseAttributesWithDefaults() *DeploymentGatesEvaluationResultResponseAttributes {
	this := DeploymentGatesEvaluationResultResponseAttributes{}
	return &this
}

// GetDryRun returns the DryRun field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DryRun, true
}

// SetDryRun sets field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) SetDryRun(v bool) {
	o.DryRun = v
}

// GetEvaluationId returns the EvaluationId field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetEvaluationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvaluationId
}

// GetEvaluationIdOk returns a tuple with the EvaluationId field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetEvaluationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationId, true
}

// SetEvaluationId sets field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) SetEvaluationId(v string) {
	o.EvaluationId = v
}

// GetEvaluationUrl returns the EvaluationUrl field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetEvaluationUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvaluationUrl
}

// GetEvaluationUrlOk returns a tuple with the EvaluationUrl field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetEvaluationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationUrl, true
}

// SetEvaluationUrl sets field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) SetEvaluationUrl(v string) {
	o.EvaluationUrl = v
}

// GetGateId returns the GateId field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetGateId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.GateId
}

// GetGateIdOk returns a tuple with the GateId field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetGateIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GateId, true
}

// SetGateId sets field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) SetGateId(v uuid.UUID) {
	o.GateId = v
}

// GetGateStatus returns the GateStatus field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetGateStatus() DeploymentGatesEvaluationResultResponseAttributesGateStatus {
	if o == nil {
		var ret DeploymentGatesEvaluationResultResponseAttributesGateStatus
		return ret
	}
	return o.GateStatus
}

// GetGateStatusOk returns a tuple with the GateStatus field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetGateStatusOk() (*DeploymentGatesEvaluationResultResponseAttributesGateStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GateStatus, true
}

// SetGateStatus sets field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) SetGateStatus(v DeploymentGatesEvaluationResultResponseAttributesGateStatus) {
	o.GateStatus = v
}

// GetRules returns the Rules field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetRules() []DeploymentGatesRuleResponse {
	if o == nil {
		var ret []DeploymentGatesRuleResponse
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationResultResponseAttributes) GetRulesOk() (*[]DeploymentGatesRuleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *DeploymentGatesEvaluationResultResponseAttributes) SetRules(v []DeploymentGatesRuleResponse) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGatesEvaluationResultResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dry_run"] = o.DryRun
	toSerialize["evaluation_id"] = o.EvaluationId
	toSerialize["evaluation_url"] = o.EvaluationUrl
	toSerialize["gate_id"] = o.GateId
	toSerialize["gate_status"] = o.GateStatus
	toSerialize["rules"] = o.Rules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGatesEvaluationResultResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DryRun        *bool                                                        `json:"dry_run"`
		EvaluationId  *string                                                      `json:"evaluation_id"`
		EvaluationUrl *string                                                      `json:"evaluation_url"`
		GateId        *uuid.UUID                                                   `json:"gate_id"`
		GateStatus    *DeploymentGatesEvaluationResultResponseAttributesGateStatus `json:"gate_status"`
		Rules         *[]DeploymentGatesRuleResponse                               `json:"rules"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DryRun == nil {
		return fmt.Errorf("required field dry_run missing")
	}
	if all.EvaluationId == nil {
		return fmt.Errorf("required field evaluation_id missing")
	}
	if all.EvaluationUrl == nil {
		return fmt.Errorf("required field evaluation_url missing")
	}
	if all.GateId == nil {
		return fmt.Errorf("required field gate_id missing")
	}
	if all.GateStatus == nil {
		return fmt.Errorf("required field gate_status missing")
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dry_run", "evaluation_id", "evaluation_url", "gate_id", "gate_status", "rules"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DryRun = *all.DryRun
	o.EvaluationId = *all.EvaluationId
	o.EvaluationUrl = *all.EvaluationUrl
	o.GateId = *all.GateId
	if !all.GateStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.GateStatus = *all.GateStatus
	}
	o.Rules = *all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
