// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuardrailMetric Guardrail metric details.
type GuardrailMetric struct {
	// The metric ID to monitor.
	MetricId string `json:"metric_id"`
	// Action to perform when a guardrail threshold is triggered.
	TriggerAction GuardrailTriggerAction `json:"trigger_action"`
	// The signal or system that triggered the action.
	TriggeredBy datadog.NullableString `json:"triggered_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuardrailMetric instantiates a new GuardrailMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuardrailMetric(metricId string, triggerAction GuardrailTriggerAction) *GuardrailMetric {
	this := GuardrailMetric{}
	this.MetricId = metricId
	this.TriggerAction = triggerAction
	return &this
}

// NewGuardrailMetricWithDefaults instantiates a new GuardrailMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuardrailMetricWithDefaults() *GuardrailMetric {
	this := GuardrailMetric{}
	return &this
}

// GetMetricId returns the MetricId field value.
func (o *GuardrailMetric) GetMetricId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *GuardrailMetric) GetMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value.
func (o *GuardrailMetric) SetMetricId(v string) {
	o.MetricId = v
}

// GetTriggerAction returns the TriggerAction field value.
func (o *GuardrailMetric) GetTriggerAction() GuardrailTriggerAction {
	if o == nil {
		var ret GuardrailTriggerAction
		return ret
	}
	return o.TriggerAction
}

// GetTriggerActionOk returns a tuple with the TriggerAction field value
// and a boolean to check if the value has been set.
func (o *GuardrailMetric) GetTriggerActionOk() (*GuardrailTriggerAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerAction, true
}

// SetTriggerAction sets field value.
func (o *GuardrailMetric) SetTriggerAction(v GuardrailTriggerAction) {
	o.TriggerAction = v
}

// GetTriggeredBy returns the TriggeredBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuardrailMetric) GetTriggeredBy() string {
	if o == nil || o.TriggeredBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.TriggeredBy.Get()
}

// GetTriggeredByOk returns a tuple with the TriggeredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GuardrailMetric) GetTriggeredByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TriggeredBy.Get(), o.TriggeredBy.IsSet()
}

// HasTriggeredBy returns a boolean if a field has been set.
func (o *GuardrailMetric) HasTriggeredBy() bool {
	return o != nil && o.TriggeredBy.IsSet()
}

// SetTriggeredBy gets a reference to the given datadog.NullableString and assigns it to the TriggeredBy field.
func (o *GuardrailMetric) SetTriggeredBy(v string) {
	o.TriggeredBy.Set(&v)
}

// SetTriggeredByNil sets the value for TriggeredBy to be an explicit nil.
func (o *GuardrailMetric) SetTriggeredByNil() {
	o.TriggeredBy.Set(nil)
}

// UnsetTriggeredBy ensures that no value is present for TriggeredBy, not even an explicit nil.
func (o *GuardrailMetric) UnsetTriggeredBy() {
	o.TriggeredBy.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o GuardrailMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["metric_id"] = o.MetricId
	toSerialize["trigger_action"] = o.TriggerAction
	if o.TriggeredBy.IsSet() {
		toSerialize["triggered_by"] = o.TriggeredBy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuardrailMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetricId      *string                 `json:"metric_id"`
		TriggerAction *GuardrailTriggerAction `json:"trigger_action"`
		TriggeredBy   datadog.NullableString  `json:"triggered_by,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"metric_id", "trigger_action", "triggered_by"})
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
	o.TriggeredBy = all.TriggeredBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
