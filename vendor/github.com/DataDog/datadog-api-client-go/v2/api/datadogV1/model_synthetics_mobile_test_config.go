// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestConfig Configuration object for a Synthetic mobile test.
type SyntheticsMobileTestConfig struct {
	// Initial application arguments for a mobile test.
	InitialApplicationArguments map[string]string `json:"initialApplicationArguments,omitempty"`
	// Array of variables used for the test steps.
	Variables []SyntheticsConfigVariable `json:"variables,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileTestConfig instantiates a new SyntheticsMobileTestConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileTestConfig() *SyntheticsMobileTestConfig {
	this := SyntheticsMobileTestConfig{}
	return &this
}

// NewSyntheticsMobileTestConfigWithDefaults instantiates a new SyntheticsMobileTestConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileTestConfigWithDefaults() *SyntheticsMobileTestConfig {
	this := SyntheticsMobileTestConfig{}
	return &this
}

// GetInitialApplicationArguments returns the InitialApplicationArguments field value if set, zero value otherwise.
func (o *SyntheticsMobileTestConfig) GetInitialApplicationArguments() map[string]string {
	if o == nil || o.InitialApplicationArguments == nil {
		var ret map[string]string
		return ret
	}
	return o.InitialApplicationArguments
}

// GetInitialApplicationArgumentsOk returns a tuple with the InitialApplicationArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestConfig) GetInitialApplicationArgumentsOk() (*map[string]string, bool) {
	if o == nil || o.InitialApplicationArguments == nil {
		return nil, false
	}
	return &o.InitialApplicationArguments, true
}

// HasInitialApplicationArguments returns a boolean if a field has been set.
func (o *SyntheticsMobileTestConfig) HasInitialApplicationArguments() bool {
	return o != nil && o.InitialApplicationArguments != nil
}

// SetInitialApplicationArguments gets a reference to the given map[string]string and assigns it to the InitialApplicationArguments field.
func (o *SyntheticsMobileTestConfig) SetInitialApplicationArguments(v map[string]string) {
	o.InitialApplicationArguments = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *SyntheticsMobileTestConfig) GetVariables() []SyntheticsConfigVariable {
	if o == nil || o.Variables == nil {
		var ret []SyntheticsConfigVariable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestConfig) GetVariablesOk() (*[]SyntheticsConfigVariable, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *SyntheticsMobileTestConfig) HasVariables() bool {
	return o != nil && o.Variables != nil
}

// SetVariables gets a reference to the given []SyntheticsConfigVariable and assigns it to the Variables field.
func (o *SyntheticsMobileTestConfig) SetVariables(v []SyntheticsConfigVariable) {
	o.Variables = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileTestConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.InitialApplicationArguments != nil {
		toSerialize["initialApplicationArguments"] = o.InitialApplicationArguments
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileTestConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InitialApplicationArguments map[string]string          `json:"initialApplicationArguments,omitempty"`
		Variables                   []SyntheticsConfigVariable `json:"variables,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"initialApplicationArguments", "variables"})
	} else {
		return err
	}
	o.InitialApplicationArguments = all.InitialApplicationArguments
	o.Variables = all.Variables

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
