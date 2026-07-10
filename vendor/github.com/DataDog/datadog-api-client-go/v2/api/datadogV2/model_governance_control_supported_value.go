// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlSupportedValue A supported value for an enumerated parameter.
type GovernanceControlSupportedValue struct {
	// The human-readable label for the value.
	Label string `json:"label"`
	// The machine-readable value.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceControlSupportedValue instantiates a new GovernanceControlSupportedValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceControlSupportedValue(label string, value string) *GovernanceControlSupportedValue {
	this := GovernanceControlSupportedValue{}
	this.Label = label
	this.Value = value
	return &this
}

// NewGovernanceControlSupportedValueWithDefaults instantiates a new GovernanceControlSupportedValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceControlSupportedValueWithDefaults() *GovernanceControlSupportedValue {
	this := GovernanceControlSupportedValue{}
	return &this
}

// GetLabel returns the Label field value.
func (o *GovernanceControlSupportedValue) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlSupportedValue) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *GovernanceControlSupportedValue) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value.
func (o *GovernanceControlSupportedValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlSupportedValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *GovernanceControlSupportedValue) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceControlSupportedValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["label"] = o.Label
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceControlSupportedValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label *string `json:"label"`
		Value *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"label", "value"})
	} else {
		return err
	}
	o.Label = *all.Label
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
