// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleQueryPayload Payload to test a rule query with the expected result.
type SecurityMonitoringRuleQueryPayload struct {
	// Expected result of the test.
	ExpectedResult *bool `json:"expectedResult,omitempty"`
	// Index of the query under test.
	Index *int64 `json:"index,omitempty"`
	// Payload used to test the rule query.
	Payload *SecurityMonitoringRuleQueryPayloadData `json:"payload,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleQueryPayload instantiates a new SecurityMonitoringRuleQueryPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleQueryPayload() *SecurityMonitoringRuleQueryPayload {
	this := SecurityMonitoringRuleQueryPayload{}
	return &this
}

// NewSecurityMonitoringRuleQueryPayloadWithDefaults instantiates a new SecurityMonitoringRuleQueryPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleQueryPayloadWithDefaults() *SecurityMonitoringRuleQueryPayload {
	this := SecurityMonitoringRuleQueryPayload{}
	return &this
}

// GetExpectedResult returns the ExpectedResult field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQueryPayload) GetExpectedResult() bool {
	if o == nil || o.ExpectedResult == nil {
		var ret bool
		return ret
	}
	return *o.ExpectedResult
}

// GetExpectedResultOk returns a tuple with the ExpectedResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQueryPayload) GetExpectedResultOk() (*bool, bool) {
	if o == nil || o.ExpectedResult == nil {
		return nil, false
	}
	return o.ExpectedResult, true
}

// HasExpectedResult returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQueryPayload) HasExpectedResult() bool {
	return o != nil && o.ExpectedResult != nil
}

// SetExpectedResult gets a reference to the given bool and assigns it to the ExpectedResult field.
func (o *SecurityMonitoringRuleQueryPayload) SetExpectedResult(v bool) {
	o.ExpectedResult = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQueryPayload) GetIndex() int64 {
	if o == nil || o.Index == nil {
		var ret int64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQueryPayload) GetIndexOk() (*int64, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQueryPayload) HasIndex() bool {
	return o != nil && o.Index != nil
}

// SetIndex gets a reference to the given int64 and assigns it to the Index field.
func (o *SecurityMonitoringRuleQueryPayload) SetIndex(v int64) {
	o.Index = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQueryPayload) GetPayload() SecurityMonitoringRuleQueryPayloadData {
	if o == nil || o.Payload == nil {
		var ret SecurityMonitoringRuleQueryPayloadData
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQueryPayload) GetPayloadOk() (*SecurityMonitoringRuleQueryPayloadData, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQueryPayload) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given SecurityMonitoringRuleQueryPayloadData and assigns it to the Payload field.
func (o *SecurityMonitoringRuleQueryPayload) SetPayload(v SecurityMonitoringRuleQueryPayloadData) {
	o.Payload = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleQueryPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExpectedResult != nil {
		toSerialize["expectedResult"] = o.ExpectedResult
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleQueryPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExpectedResult *bool                                   `json:"expectedResult,omitempty"`
		Index          *int64                                  `json:"index,omitempty"`
		Payload        *SecurityMonitoringRuleQueryPayloadData `json:"payload,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expectedResult", "index", "payload"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExpectedResult = all.ExpectedResult
	o.Index = all.Index
	if all.Payload != nil && all.Payload.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Payload = all.Payload

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
