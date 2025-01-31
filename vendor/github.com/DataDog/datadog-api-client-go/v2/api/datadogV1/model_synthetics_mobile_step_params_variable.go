// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParamsVariable Variable object for `extractVariable` step type.
type SyntheticsMobileStepParamsVariable struct {
	// An example for the variable.
	Example string `json:"example"`
	// The variable name.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileStepParamsVariable instantiates a new SyntheticsMobileStepParamsVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileStepParamsVariable(example string, name string) *SyntheticsMobileStepParamsVariable {
	this := SyntheticsMobileStepParamsVariable{}
	this.Example = example
	this.Name = name
	return &this
}

// NewSyntheticsMobileStepParamsVariableWithDefaults instantiates a new SyntheticsMobileStepParamsVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileStepParamsVariableWithDefaults() *SyntheticsMobileStepParamsVariable {
	this := SyntheticsMobileStepParamsVariable{}
	return &this
}

// GetExample returns the Example field value.
func (o *SyntheticsMobileStepParamsVariable) GetExample() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Example
}

// GetExampleOk returns a tuple with the Example field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsVariable) GetExampleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Example, true
}

// SetExample sets field value.
func (o *SyntheticsMobileStepParamsVariable) SetExample(v string) {
	o.Example = v
}

// GetName returns the Name field value.
func (o *SyntheticsMobileStepParamsVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsMobileStepParamsVariable) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileStepParamsVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["example"] = o.Example
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileStepParamsVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Example *string `json:"example"`
		Name    *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Example == nil {
		return fmt.Errorf("required field example missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"example", "name"})
	} else {
		return err
	}
	o.Example = *all.Example
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
