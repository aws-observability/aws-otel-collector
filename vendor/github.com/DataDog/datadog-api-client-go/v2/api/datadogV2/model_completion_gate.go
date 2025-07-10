// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompletionGate Used to create conditions before running subsequent actions.
type CompletionGate struct {
	// The definition of `CompletionCondition` object.
	CompletionCondition CompletionCondition `json:"completionCondition"`
	// The definition of `RetryStrategy` object.
	RetryStrategy RetryStrategy `json:"retryStrategy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCompletionGate instantiates a new CompletionGate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCompletionGate(completionCondition CompletionCondition, retryStrategy RetryStrategy) *CompletionGate {
	this := CompletionGate{}
	this.CompletionCondition = completionCondition
	this.RetryStrategy = retryStrategy
	return &this
}

// NewCompletionGateWithDefaults instantiates a new CompletionGate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCompletionGateWithDefaults() *CompletionGate {
	this := CompletionGate{}
	return &this
}

// GetCompletionCondition returns the CompletionCondition field value.
func (o *CompletionGate) GetCompletionCondition() CompletionCondition {
	if o == nil {
		var ret CompletionCondition
		return ret
	}
	return o.CompletionCondition
}

// GetCompletionConditionOk returns a tuple with the CompletionCondition field value
// and a boolean to check if the value has been set.
func (o *CompletionGate) GetCompletionConditionOk() (*CompletionCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletionCondition, true
}

// SetCompletionCondition sets field value.
func (o *CompletionGate) SetCompletionCondition(v CompletionCondition) {
	o.CompletionCondition = v
}

// GetRetryStrategy returns the RetryStrategy field value.
func (o *CompletionGate) GetRetryStrategy() RetryStrategy {
	if o == nil {
		var ret RetryStrategy
		return ret
	}
	return o.RetryStrategy
}

// GetRetryStrategyOk returns a tuple with the RetryStrategy field value
// and a boolean to check if the value has been set.
func (o *CompletionGate) GetRetryStrategyOk() (*RetryStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetryStrategy, true
}

// SetRetryStrategy sets field value.
func (o *CompletionGate) SetRetryStrategy(v RetryStrategy) {
	o.RetryStrategy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CompletionGate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["completionCondition"] = o.CompletionCondition
	toSerialize["retryStrategy"] = o.RetryStrategy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CompletionGate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompletionCondition *CompletionCondition `json:"completionCondition"`
		RetryStrategy       *RetryStrategy       `json:"retryStrategy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CompletionCondition == nil {
		return fmt.Errorf("required field completionCondition missing")
	}
	if all.RetryStrategy == nil {
		return fmt.Errorf("required field retryStrategy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completionCondition", "retryStrategy"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CompletionCondition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CompletionCondition = *all.CompletionCondition
	if all.RetryStrategy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RetryStrategy = *all.RetryStrategy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
