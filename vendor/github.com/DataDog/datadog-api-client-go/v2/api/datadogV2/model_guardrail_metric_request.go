// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuardrailMetricRequest Guardrail metric request payload.
type GuardrailMetricRequest struct {
	// The metric ID to monitor.
	MetricId string `json:"metric_id"`
	// Action to perform when a guardrail threshold is triggered.
	TriggerAction GuardrailTriggerAction `json:"trigger_action"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuardrailMetricRequest instantiates a new GuardrailMetricRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuardrailMetricRequest(metricId string, triggerAction GuardrailTriggerAction) *GuardrailMetricRequest {
	this := GuardrailMetricRequest{}
	this.MetricId = metricId
	this.TriggerAction = triggerAction
	return &this
}

// NewGuardrailMetricRequestWithDefaults instantiates a new GuardrailMetricRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuardrailMetricRequestWithDefaults() *GuardrailMetricRequest {
	this := GuardrailMetricRequest{}
	return &this
}

// GetMetricId returns the MetricId field value.
func (o *GuardrailMetricRequest) GetMetricId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *GuardrailMetricRequest) GetMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value.
func (o *GuardrailMetricRequest) SetMetricId(v string) {
	o.MetricId = v
}

// GetTriggerAction returns the TriggerAction field value.
func (o *GuardrailMetricRequest) GetTriggerAction() GuardrailTriggerAction {
	if o == nil {
		var ret GuardrailTriggerAction
		return ret
	}
	return o.TriggerAction
}

// GetTriggerActionOk returns a tuple with the TriggerAction field value
// and a boolean to check if the value has been set.
func (o *GuardrailMetricRequest) GetTriggerActionOk() (*GuardrailTriggerAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerAction, true
}

// SetTriggerAction sets field value.
func (o *GuardrailMetricRequest) SetTriggerAction(v GuardrailTriggerAction) {
	o.TriggerAction = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuardrailMetricRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["metric_id"] = o.MetricId
	toSerialize["trigger_action"] = o.TriggerAction

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuardrailMetricRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetricId      *string                 `json:"metric_id"`
		TriggerAction *GuardrailTriggerAction `json:"trigger_action"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MetricId == nil {
		return fmt.Errorf("required field metric_id missing")
	}
	if all.TriggerAction == nil {
		return fmt.Errorf("required field trigger_action missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metric_id", "trigger_action"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MetricId = *all.MetricId
	if !all.TriggerAction.IsValid() {
		hasInvalidField = true
	} else {
		o.TriggerAction = *all.TriggerAction
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
