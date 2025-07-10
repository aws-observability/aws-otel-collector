// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetryStrategyLinear The definition of `RetryStrategyLinear` object.
type RetryStrategyLinear struct {
	// The `RetryStrategyLinear` `interval`. The expected format is the number of seconds ending with an s. For example, 1 day is 86400s
	Interval string `json:"interval"`
	// The `RetryStrategyLinear` `maxRetries`.
	MaxRetries float64 `json:"maxRetries"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRetryStrategyLinear instantiates a new RetryStrategyLinear object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetryStrategyLinear(interval string, maxRetries float64) *RetryStrategyLinear {
	this := RetryStrategyLinear{}
	this.Interval = interval
	this.MaxRetries = maxRetries
	return &this
}

// NewRetryStrategyLinearWithDefaults instantiates a new RetryStrategyLinear object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetryStrategyLinearWithDefaults() *RetryStrategyLinear {
	this := RetryStrategyLinear{}
	return &this
}

// GetInterval returns the Interval field value.
func (o *RetryStrategyLinear) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *RetryStrategyLinear) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value.
func (o *RetryStrategyLinear) SetInterval(v string) {
	o.Interval = v
}

// GetMaxRetries returns the MaxRetries field value.
func (o *RetryStrategyLinear) GetMaxRetries() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value
// and a boolean to check if the value has been set.
func (o *RetryStrategyLinear) GetMaxRetriesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRetries, true
}

// SetMaxRetries sets field value.
func (o *RetryStrategyLinear) SetMaxRetries(v float64) {
	o.MaxRetries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetryStrategyLinear) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["interval"] = o.Interval
	toSerialize["maxRetries"] = o.MaxRetries

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetryStrategyLinear) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interval   *string  `json:"interval"`
		MaxRetries *float64 `json:"maxRetries"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Interval == nil {
		return fmt.Errorf("required field interval missing")
	}
	if all.MaxRetries == nil {
		return fmt.Errorf("required field maxRetries missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"interval", "maxRetries"})
	} else {
		return err
	}
	o.Interval = *all.Interval
	o.MaxRetries = *all.MaxRetries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
