// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ErrorHandler Used to handle errors in an action.
type ErrorHandler struct {
	// The `ErrorHandler` `fallbackStepName`.
	FallbackStepName string `json:"fallbackStepName"`
	// The definition of `RetryStrategy` object.
	RetryStrategy RetryStrategy `json:"retryStrategy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewErrorHandler instantiates a new ErrorHandler object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewErrorHandler(fallbackStepName string, retryStrategy RetryStrategy) *ErrorHandler {
	this := ErrorHandler{}
	this.FallbackStepName = fallbackStepName
	this.RetryStrategy = retryStrategy
	return &this
}

// NewErrorHandlerWithDefaults instantiates a new ErrorHandler object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewErrorHandlerWithDefaults() *ErrorHandler {
	this := ErrorHandler{}
	return &this
}

// GetFallbackStepName returns the FallbackStepName field value.
func (o *ErrorHandler) GetFallbackStepName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FallbackStepName
}

// GetFallbackStepNameOk returns a tuple with the FallbackStepName field value
// and a boolean to check if the value has been set.
func (o *ErrorHandler) GetFallbackStepNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FallbackStepName, true
}

// SetFallbackStepName sets field value.
func (o *ErrorHandler) SetFallbackStepName(v string) {
	o.FallbackStepName = v
}

// GetRetryStrategy returns the RetryStrategy field value.
func (o *ErrorHandler) GetRetryStrategy() RetryStrategy {
	if o == nil {
		var ret RetryStrategy
		return ret
	}
	return o.RetryStrategy
}

// GetRetryStrategyOk returns a tuple with the RetryStrategy field value
// and a boolean to check if the value has been set.
func (o *ErrorHandler) GetRetryStrategyOk() (*RetryStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetryStrategy, true
}

// SetRetryStrategy sets field value.
func (o *ErrorHandler) SetRetryStrategy(v RetryStrategy) {
	o.RetryStrategy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ErrorHandler) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fallbackStepName"] = o.FallbackStepName
	toSerialize["retryStrategy"] = o.RetryStrategy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ErrorHandler) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FallbackStepName *string        `json:"fallbackStepName"`
		RetryStrategy    *RetryStrategy `json:"retryStrategy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FallbackStepName == nil {
		return fmt.Errorf("required field fallbackStepName missing")
	}
	if all.RetryStrategy == nil {
		return fmt.Errorf("required field retryStrategy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fallbackStepName", "retryStrategy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FallbackStepName = *all.FallbackStepName
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
