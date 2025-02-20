// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsGlobalVariableOptions Options for the Global Variable for MFA.
type SyntheticsGlobalVariableOptions struct {
	// Parameters for the TOTP/MFA variable
	TotpParameters *SyntheticsGlobalVariableTOTPParameters `json:"totp_parameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsGlobalVariableOptions instantiates a new SyntheticsGlobalVariableOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsGlobalVariableOptions() *SyntheticsGlobalVariableOptions {
	this := SyntheticsGlobalVariableOptions{}
	return &this
}

// NewSyntheticsGlobalVariableOptionsWithDefaults instantiates a new SyntheticsGlobalVariableOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsGlobalVariableOptionsWithDefaults() *SyntheticsGlobalVariableOptions {
	this := SyntheticsGlobalVariableOptions{}
	return &this
}

// GetTotpParameters returns the TotpParameters field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariableOptions) GetTotpParameters() SyntheticsGlobalVariableTOTPParameters {
	if o == nil || o.TotpParameters == nil {
		var ret SyntheticsGlobalVariableTOTPParameters
		return ret
	}
	return *o.TotpParameters
}

// GetTotpParametersOk returns a tuple with the TotpParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableOptions) GetTotpParametersOk() (*SyntheticsGlobalVariableTOTPParameters, bool) {
	if o == nil || o.TotpParameters == nil {
		return nil, false
	}
	return o.TotpParameters, true
}

// HasTotpParameters returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariableOptions) HasTotpParameters() bool {
	return o != nil && o.TotpParameters != nil
}

// SetTotpParameters gets a reference to the given SyntheticsGlobalVariableTOTPParameters and assigns it to the TotpParameters field.
func (o *SyntheticsGlobalVariableOptions) SetTotpParameters(v SyntheticsGlobalVariableTOTPParameters) {
	o.TotpParameters = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsGlobalVariableOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TotpParameters != nil {
		toSerialize["totp_parameters"] = o.TotpParameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsGlobalVariableOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotpParameters *SyntheticsGlobalVariableTOTPParameters `json:"totp_parameters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"totp_parameters"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.TotpParameters != nil && all.TotpParameters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TotpParameters = all.TotpParameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
