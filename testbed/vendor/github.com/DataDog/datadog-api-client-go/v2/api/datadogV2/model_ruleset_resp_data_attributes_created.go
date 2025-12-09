// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesetRespDataAttributesCreated The definition of `RulesetRespDataAttributesCreated` object.
type RulesetRespDataAttributesCreated struct {
	// The `created` `nanos`.
	Nanos *int32 `json:"nanos,omitempty"`
	// The `created` `seconds`.
	Seconds *int64 `json:"seconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRulesetRespDataAttributesCreated instantiates a new RulesetRespDataAttributesCreated object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRulesetRespDataAttributesCreated() *RulesetRespDataAttributesCreated {
	this := RulesetRespDataAttributesCreated{}
	return &this
}

// NewRulesetRespDataAttributesCreatedWithDefaults instantiates a new RulesetRespDataAttributesCreated object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRulesetRespDataAttributesCreatedWithDefaults() *RulesetRespDataAttributesCreated {
	this := RulesetRespDataAttributesCreated{}
	return &this
}

// GetNanos returns the Nanos field value if set, zero value otherwise.
func (o *RulesetRespDataAttributesCreated) GetNanos() int32 {
	if o == nil || o.Nanos == nil {
		var ret int32
		return ret
	}
	return *o.Nanos
}

// GetNanosOk returns a tuple with the Nanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributesCreated) GetNanosOk() (*int32, bool) {
	if o == nil || o.Nanos == nil {
		return nil, false
	}
	return o.Nanos, true
}

// HasNanos returns a boolean if a field has been set.
func (o *RulesetRespDataAttributesCreated) HasNanos() bool {
	return o != nil && o.Nanos != nil
}

// SetNanos gets a reference to the given int32 and assigns it to the Nanos field.
func (o *RulesetRespDataAttributesCreated) SetNanos(v int32) {
	o.Nanos = &v
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *RulesetRespDataAttributesCreated) GetSeconds() int64 {
	if o == nil || o.Seconds == nil {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributesCreated) GetSecondsOk() (*int64, bool) {
	if o == nil || o.Seconds == nil {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *RulesetRespDataAttributesCreated) HasSeconds() bool {
	return o != nil && o.Seconds != nil
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *RulesetRespDataAttributesCreated) SetSeconds(v int64) {
	o.Seconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RulesetRespDataAttributesCreated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Nanos != nil {
		toSerialize["nanos"] = o.Nanos
	}
	if o.Seconds != nil {
		toSerialize["seconds"] = o.Seconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RulesetRespDataAttributesCreated) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Nanos   *int32 `json:"nanos,omitempty"`
		Seconds *int64 `json:"seconds,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"nanos", "seconds"})
	} else {
		return err
	}
	o.Nanos = all.Nanos
	o.Seconds = all.Seconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
