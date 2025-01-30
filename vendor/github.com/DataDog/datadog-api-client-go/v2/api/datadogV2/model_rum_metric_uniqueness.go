// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumMetricUniqueness The rule to count updatable events. Is only set if `event_type` is `sessions` or `views`.
type RumMetricUniqueness struct {
	// When to count updatable events. `match` when the event is first seen, or `end` when the event is complete.
	When RumMetricUniquenessWhen `json:"when"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumMetricUniqueness instantiates a new RumMetricUniqueness object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumMetricUniqueness(when RumMetricUniquenessWhen) *RumMetricUniqueness {
	this := RumMetricUniqueness{}
	this.When = when
	return &this
}

// NewRumMetricUniquenessWithDefaults instantiates a new RumMetricUniqueness object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumMetricUniquenessWithDefaults() *RumMetricUniqueness {
	this := RumMetricUniqueness{}
	return &this
}

// GetWhen returns the When field value.
func (o *RumMetricUniqueness) GetWhen() RumMetricUniquenessWhen {
	if o == nil {
		var ret RumMetricUniquenessWhen
		return ret
	}
	return o.When
}

// GetWhenOk returns a tuple with the When field value
// and a boolean to check if the value has been set.
func (o *RumMetricUniqueness) GetWhenOk() (*RumMetricUniquenessWhen, bool) {
	if o == nil {
		return nil, false
	}
	return &o.When, true
}

// SetWhen sets field value.
func (o *RumMetricUniqueness) SetWhen(v RumMetricUniquenessWhen) {
	o.When = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumMetricUniqueness) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["when"] = o.When

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumMetricUniqueness) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		When *RumMetricUniquenessWhen `json:"when"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.When == nil {
		return fmt.Errorf("required field when missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"when"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.When.IsValid() {
		hasInvalidField = true
	} else {
		o.When = *all.When
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
