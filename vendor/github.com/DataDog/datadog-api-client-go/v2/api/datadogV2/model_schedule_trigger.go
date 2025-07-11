// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleTrigger Trigger a workflow from a Schedule. The workflow must be published.
type ScheduleTrigger struct {
	// Recurrence rule expression for scheduling.
	RruleExpression string `json:"rruleExpression"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleTrigger instantiates a new ScheduleTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleTrigger(rruleExpression string) *ScheduleTrigger {
	this := ScheduleTrigger{}
	this.RruleExpression = rruleExpression
	return &this
}

// NewScheduleTriggerWithDefaults instantiates a new ScheduleTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleTriggerWithDefaults() *ScheduleTrigger {
	this := ScheduleTrigger{}
	return &this
}

// GetRruleExpression returns the RruleExpression field value.
func (o *ScheduleTrigger) GetRruleExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RruleExpression
}

// GetRruleExpressionOk returns a tuple with the RruleExpression field value
// and a boolean to check if the value has been set.
func (o *ScheduleTrigger) GetRruleExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RruleExpression, true
}

// SetRruleExpression sets field value.
func (o *ScheduleTrigger) SetRruleExpression(v string) {
	o.RruleExpression = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rruleExpression"] = o.RruleExpression

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RruleExpression *string `json:"rruleExpression"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RruleExpression == nil {
		return fmt.Errorf("required field rruleExpression missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rruleExpression"})
	} else {
		return err
	}
	o.RruleExpression = *all.RruleExpression

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
