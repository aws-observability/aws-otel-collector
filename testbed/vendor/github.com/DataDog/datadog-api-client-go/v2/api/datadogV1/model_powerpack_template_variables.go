// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackTemplateVariables Powerpack template variables.
type PowerpackTemplateVariables struct {
	// Template variables controlled at the powerpack level.
	ControlledByPowerpack []PowerpackTemplateVariableContents `json:"controlled_by_powerpack,omitempty"`
	// Template variables controlled by the external resource, such as the dashboard this powerpack is on.
	ControlledExternally []PowerpackTemplateVariableContents `json:"controlled_externally,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPowerpackTemplateVariables instantiates a new PowerpackTemplateVariables object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackTemplateVariables() *PowerpackTemplateVariables {
	this := PowerpackTemplateVariables{}
	return &this
}

// NewPowerpackTemplateVariablesWithDefaults instantiates a new PowerpackTemplateVariables object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackTemplateVariablesWithDefaults() *PowerpackTemplateVariables {
	this := PowerpackTemplateVariables{}
	return &this
}

// GetControlledByPowerpack returns the ControlledByPowerpack field value if set, zero value otherwise.
func (o *PowerpackTemplateVariables) GetControlledByPowerpack() []PowerpackTemplateVariableContents {
	if o == nil || o.ControlledByPowerpack == nil {
		var ret []PowerpackTemplateVariableContents
		return ret
	}
	return o.ControlledByPowerpack
}

// GetControlledByPowerpackOk returns a tuple with the ControlledByPowerpack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackTemplateVariables) GetControlledByPowerpackOk() (*[]PowerpackTemplateVariableContents, bool) {
	if o == nil || o.ControlledByPowerpack == nil {
		return nil, false
	}
	return &o.ControlledByPowerpack, true
}

// HasControlledByPowerpack returns a boolean if a field has been set.
func (o *PowerpackTemplateVariables) HasControlledByPowerpack() bool {
	return o != nil && o.ControlledByPowerpack != nil
}

// SetControlledByPowerpack gets a reference to the given []PowerpackTemplateVariableContents and assigns it to the ControlledByPowerpack field.
func (o *PowerpackTemplateVariables) SetControlledByPowerpack(v []PowerpackTemplateVariableContents) {
	o.ControlledByPowerpack = v
}

// GetControlledExternally returns the ControlledExternally field value if set, zero value otherwise.
func (o *PowerpackTemplateVariables) GetControlledExternally() []PowerpackTemplateVariableContents {
	if o == nil || o.ControlledExternally == nil {
		var ret []PowerpackTemplateVariableContents
		return ret
	}
	return o.ControlledExternally
}

// GetControlledExternallyOk returns a tuple with the ControlledExternally field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackTemplateVariables) GetControlledExternallyOk() (*[]PowerpackTemplateVariableContents, bool) {
	if o == nil || o.ControlledExternally == nil {
		return nil, false
	}
	return &o.ControlledExternally, true
}

// HasControlledExternally returns a boolean if a field has been set.
func (o *PowerpackTemplateVariables) HasControlledExternally() bool {
	return o != nil && o.ControlledExternally != nil
}

// SetControlledExternally gets a reference to the given []PowerpackTemplateVariableContents and assigns it to the ControlledExternally field.
func (o *PowerpackTemplateVariables) SetControlledExternally(v []PowerpackTemplateVariableContents) {
	o.ControlledExternally = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackTemplateVariables) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ControlledByPowerpack != nil {
		toSerialize["controlled_by_powerpack"] = o.ControlledByPowerpack
	}
	if o.ControlledExternally != nil {
		toSerialize["controlled_externally"] = o.ControlledExternally
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackTemplateVariables) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ControlledByPowerpack []PowerpackTemplateVariableContents `json:"controlled_by_powerpack,omitempty"`
		ControlledExternally  []PowerpackTemplateVariableContents `json:"controlled_externally,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"controlled_by_powerpack", "controlled_externally"})
	} else {
		return err
	}
	o.ControlledByPowerpack = all.ControlledByPowerpack
	o.ControlledExternally = all.ControlledExternally

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
