// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesRuleResponse The result of a single rule evaluation.
type DeploymentGatesRuleResponse struct {
	// Whether this rule was evaluated in dry-run mode.
	DryRun *bool `json:"dry_run,omitempty"`
	// The name of the rule.
	Name *string `json:"name,omitempty"`
	// The reason for the rule result, if applicable.
	Reason *string `json:"reason,omitempty"`
	// The overall status of the gate evaluation.
	// - `in_progress`: The evaluation is still running.
	// - `pass`: All rules passed successfully and the deployment is allowed to proceed.
	// - `fail`: One or more rules did not pass; the deployment should not proceed.
	Status *DeploymentGatesEvaluationResultResponseAttributesGateStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGatesRuleResponse instantiates a new DeploymentGatesRuleResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGatesRuleResponse() *DeploymentGatesRuleResponse {
	this := DeploymentGatesRuleResponse{}
	return &this
}

// NewDeploymentGatesRuleResponseWithDefaults instantiates a new DeploymentGatesRuleResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGatesRuleResponseWithDefaults() *DeploymentGatesRuleResponse {
	this := DeploymentGatesRuleResponse{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeploymentGatesRuleResponse) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesRuleResponse) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeploymentGatesRuleResponse) HasDryRun() bool {
	return o != nil && o.DryRun != nil
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeploymentGatesRuleResponse) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentGatesRuleResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesRuleResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentGatesRuleResponse) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentGatesRuleResponse) SetName(v string) {
	o.Name = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DeploymentGatesRuleResponse) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesRuleResponse) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DeploymentGatesRuleResponse) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DeploymentGatesRuleResponse) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentGatesRuleResponse) GetStatus() DeploymentGatesEvaluationResultResponseAttributesGateStatus {
	if o == nil || o.Status == nil {
		var ret DeploymentGatesEvaluationResultResponseAttributesGateStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesRuleResponse) GetStatusOk() (*DeploymentGatesEvaluationResultResponseAttributesGateStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentGatesRuleResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given DeploymentGatesEvaluationResultResponseAttributesGateStatus and assigns it to the Status field.
func (o *DeploymentGatesRuleResponse) SetStatus(v DeploymentGatesEvaluationResultResponseAttributesGateStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGatesRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGatesRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DryRun *bool                                                        `json:"dry_run,omitempty"`
		Name   *string                                                      `json:"name,omitempty"`
		Reason *string                                                      `json:"reason,omitempty"`
		Status *DeploymentGatesEvaluationResultResponseAttributesGateStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dry_run", "name", "reason", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DryRun = all.DryRun
	o.Name = all.Name
	o.Reason = all.Reason
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
