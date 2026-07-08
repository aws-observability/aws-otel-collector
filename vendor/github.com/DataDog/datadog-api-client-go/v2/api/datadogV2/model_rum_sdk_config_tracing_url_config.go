// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSdkConfigTracingUrlConfig Configuration for a URL that should have distributed tracing enabled.
type RumSdkConfigTracingUrlConfig struct {
	// A match option used for URL or origin pattern matching.
	Match RumSdkConfigMatchOption `json:"match"`
	// The list of trace propagator types to use for this URL.
	PropagatorTypes []RumSdkConfigTracingUrlPropagatorType `json:"propagator_types"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSdkConfigTracingUrlConfig instantiates a new RumSdkConfigTracingUrlConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSdkConfigTracingUrlConfig(match RumSdkConfigMatchOption, propagatorTypes []RumSdkConfigTracingUrlPropagatorType) *RumSdkConfigTracingUrlConfig {
	this := RumSdkConfigTracingUrlConfig{}
	this.Match = match
	this.PropagatorTypes = propagatorTypes
	return &this
}

// NewRumSdkConfigTracingUrlConfigWithDefaults instantiates a new RumSdkConfigTracingUrlConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSdkConfigTracingUrlConfigWithDefaults() *RumSdkConfigTracingUrlConfig {
	this := RumSdkConfigTracingUrlConfig{}
	return &this
}

// GetMatch returns the Match field value.
func (o *RumSdkConfigTracingUrlConfig) GetMatch() RumSdkConfigMatchOption {
	if o == nil {
		var ret RumSdkConfigMatchOption
		return ret
	}
	return o.Match
}

// GetMatchOk returns a tuple with the Match field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigTracingUrlConfig) GetMatchOk() (*RumSdkConfigMatchOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Match, true
}

// SetMatch sets field value.
func (o *RumSdkConfigTracingUrlConfig) SetMatch(v RumSdkConfigMatchOption) {
	o.Match = v
}

// GetPropagatorTypes returns the PropagatorTypes field value.
func (o *RumSdkConfigTracingUrlConfig) GetPropagatorTypes() []RumSdkConfigTracingUrlPropagatorType {
	if o == nil {
		var ret []RumSdkConfigTracingUrlPropagatorType
		return ret
	}
	return o.PropagatorTypes
}

// GetPropagatorTypesOk returns a tuple with the PropagatorTypes field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigTracingUrlConfig) GetPropagatorTypesOk() (*[]RumSdkConfigTracingUrlPropagatorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropagatorTypes, true
}

// SetPropagatorTypes sets field value.
func (o *RumSdkConfigTracingUrlConfig) SetPropagatorTypes(v []RumSdkConfigTracingUrlPropagatorType) {
	o.PropagatorTypes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSdkConfigTracingUrlConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["match"] = o.Match
	toSerialize["propagator_types"] = o.PropagatorTypes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSdkConfigTracingUrlConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Match           *RumSdkConfigMatchOption                `json:"match"`
		PropagatorTypes *[]RumSdkConfigTracingUrlPropagatorType `json:"propagator_types"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Match == nil {
		return fmt.Errorf("required field match missing")
	}
	if all.PropagatorTypes == nil {
		return fmt.Errorf("required field propagator_types missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"match", "propagator_types"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Match.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Match = *all.Match
	o.PropagatorTypes = *all.PropagatorTypes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
