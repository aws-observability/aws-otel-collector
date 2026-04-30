// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultVariables Variables captured during a test step.
type SyntheticsTestResultVariables struct {
	// Variables defined in the test configuration.
	Config []SyntheticsTestResultVariable `json:"config,omitempty"`
	// Variables extracted during the test execution.
	Extracted []SyntheticsTestResultVariable `json:"extracted,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultVariables instantiates a new SyntheticsTestResultVariables object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultVariables() *SyntheticsTestResultVariables {
	this := SyntheticsTestResultVariables{}
	return &this
}

// NewSyntheticsTestResultVariablesWithDefaults instantiates a new SyntheticsTestResultVariables object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultVariablesWithDefaults() *SyntheticsTestResultVariables {
	this := SyntheticsTestResultVariables{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariables) GetConfig() []SyntheticsTestResultVariable {
	if o == nil || o.Config == nil {
		var ret []SyntheticsTestResultVariable
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariables) GetConfigOk() (*[]SyntheticsTestResultVariable, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariables) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given []SyntheticsTestResultVariable and assigns it to the Config field.
func (o *SyntheticsTestResultVariables) SetConfig(v []SyntheticsTestResultVariable) {
	o.Config = v
}

// GetExtracted returns the Extracted field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariables) GetExtracted() []SyntheticsTestResultVariable {
	if o == nil || o.Extracted == nil {
		var ret []SyntheticsTestResultVariable
		return ret
	}
	return o.Extracted
}

// GetExtractedOk returns a tuple with the Extracted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariables) GetExtractedOk() (*[]SyntheticsTestResultVariable, bool) {
	if o == nil || o.Extracted == nil {
		return nil, false
	}
	return &o.Extracted, true
}

// HasExtracted returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariables) HasExtracted() bool {
	return o != nil && o.Extracted != nil
}

// SetExtracted gets a reference to the given []SyntheticsTestResultVariable and assigns it to the Extracted field.
func (o *SyntheticsTestResultVariables) SetExtracted(v []SyntheticsTestResultVariable) {
	o.Extracted = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultVariables) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Extracted != nil {
		toSerialize["extracted"] = o.Extracted
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultVariables) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config    []SyntheticsTestResultVariable `json:"config,omitempty"`
		Extracted []SyntheticsTestResultVariable `json:"extracted,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "extracted"})
	} else {
		return err
	}
	o.Config = all.Config
	o.Extracted = all.Extracted

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
