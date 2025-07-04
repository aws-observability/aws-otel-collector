// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetryStrategy The definition of `RetryStrategy` object.
type RetryStrategy struct {
	// The definition of `RetryStrategyKind` object.
	Kind RetryStrategyKind `json:"kind"`
	// The definition of `RetryStrategyLinear` object.
	Linear *RetryStrategyLinear `json:"linear,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRetryStrategy instantiates a new RetryStrategy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetryStrategy(kind RetryStrategyKind) *RetryStrategy {
	this := RetryStrategy{}
	this.Kind = kind
	return &this
}

// NewRetryStrategyWithDefaults instantiates a new RetryStrategy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetryStrategyWithDefaults() *RetryStrategy {
	this := RetryStrategy{}
	return &this
}

// GetKind returns the Kind field value.
func (o *RetryStrategy) GetKind() RetryStrategyKind {
	if o == nil {
		var ret RetryStrategyKind
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *RetryStrategy) GetKindOk() (*RetryStrategyKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *RetryStrategy) SetKind(v RetryStrategyKind) {
	o.Kind = v
}

// GetLinear returns the Linear field value if set, zero value otherwise.
func (o *RetryStrategy) GetLinear() RetryStrategyLinear {
	if o == nil || o.Linear == nil {
		var ret RetryStrategyLinear
		return ret
	}
	return *o.Linear
}

// GetLinearOk returns a tuple with the Linear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryStrategy) GetLinearOk() (*RetryStrategyLinear, bool) {
	if o == nil || o.Linear == nil {
		return nil, false
	}
	return o.Linear, true
}

// HasLinear returns a boolean if a field has been set.
func (o *RetryStrategy) HasLinear() bool {
	return o != nil && o.Linear != nil
}

// SetLinear gets a reference to the given RetryStrategyLinear and assigns it to the Linear field.
func (o *RetryStrategy) SetLinear(v RetryStrategyLinear) {
	o.Linear = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetryStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["kind"] = o.Kind
	if o.Linear != nil {
		toSerialize["linear"] = o.Linear
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetryStrategy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kind   *RetryStrategyKind   `json:"kind"`
		Linear *RetryStrategyLinear `json:"linear,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"kind", "linear"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = *all.Kind
	}
	if all.Linear != nil && all.Linear.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Linear = all.Linear

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
