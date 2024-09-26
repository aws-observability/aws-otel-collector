// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBrowserTestResultFullCheck Object describing the browser test configuration.
type SyntheticsBrowserTestResultFullCheck struct {
	// Configuration object for a Synthetic test.
	Config SyntheticsTestConfig `json:"config"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsBrowserTestResultFullCheck instantiates a new SyntheticsBrowserTestResultFullCheck object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsBrowserTestResultFullCheck(config SyntheticsTestConfig) *SyntheticsBrowserTestResultFullCheck {
	this := SyntheticsBrowserTestResultFullCheck{}
	this.Config = config
	return &this
}

// NewSyntheticsBrowserTestResultFullCheckWithDefaults instantiates a new SyntheticsBrowserTestResultFullCheck object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsBrowserTestResultFullCheckWithDefaults() *SyntheticsBrowserTestResultFullCheck {
	this := SyntheticsBrowserTestResultFullCheck{}
	return &this
}

// GetConfig returns the Config field value.
func (o *SyntheticsBrowserTestResultFullCheck) GetConfig() SyntheticsTestConfig {
	if o == nil {
		var ret SyntheticsTestConfig
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultFullCheck) GetConfigOk() (*SyntheticsTestConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value.
func (o *SyntheticsBrowserTestResultFullCheck) SetConfig(v SyntheticsTestConfig) {
	o.Config = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsBrowserTestResultFullCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["config"] = o.Config

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsBrowserTestResultFullCheck) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config *SyntheticsTestConfig `json:"config"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Config == nil {
		return fmt.Errorf("required field config missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Config.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Config = *all.Config

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
