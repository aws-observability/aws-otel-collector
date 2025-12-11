// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAPIWaitStep The Wait step used in a Synthetic multi-step API test.
type SyntheticsAPIWaitStep struct {
	// ID of the step.
	Id *string `json:"id,omitempty"`
	// The name of the step.
	Name string `json:"name"`
	// The subtype of the Synthetic multi-step API wait step.
	Subtype SyntheticsAPIWaitStepSubtype `json:"subtype"`
	// The time to wait in seconds. Minimum value: 0. Maximum value: 180.
	Value int32 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsAPIWaitStep instantiates a new SyntheticsAPIWaitStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAPIWaitStep(name string, subtype SyntheticsAPIWaitStepSubtype, value int32) *SyntheticsAPIWaitStep {
	this := SyntheticsAPIWaitStep{}
	this.Name = name
	this.Subtype = subtype
	this.Value = value
	return &this
}

// NewSyntheticsAPIWaitStepWithDefaults instantiates a new SyntheticsAPIWaitStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAPIWaitStepWithDefaults() *SyntheticsAPIWaitStep {
	this := SyntheticsAPIWaitStep{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsAPIWaitStep) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPIWaitStep) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsAPIWaitStep) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsAPIWaitStep) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value.
func (o *SyntheticsAPIWaitStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPIWaitStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsAPIWaitStep) SetName(v string) {
	o.Name = v
}

// GetSubtype returns the Subtype field value.
func (o *SyntheticsAPIWaitStep) GetSubtype() SyntheticsAPIWaitStepSubtype {
	if o == nil {
		var ret SyntheticsAPIWaitStepSubtype
		return ret
	}
	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPIWaitStep) GetSubtypeOk() (*SyntheticsAPIWaitStepSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value.
func (o *SyntheticsAPIWaitStep) SetSubtype(v SyntheticsAPIWaitStepSubtype) {
	o.Subtype = v
}

// GetValue returns the Value field value.
func (o *SyntheticsAPIWaitStep) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPIWaitStep) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *SyntheticsAPIWaitStep) SetValue(v int32) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAPIWaitStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["subtype"] = o.Subtype
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAPIWaitStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id      *string                       `json:"id,omitempty"`
		Name    *string                       `json:"name"`
		Subtype *SyntheticsAPIWaitStepSubtype `json:"subtype"`
		Value   *int32                        `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Subtype == nil {
		return fmt.Errorf("required field subtype missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "subtype", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = *all.Name
	if !all.Subtype.IsValid() {
		hasInvalidField = true
	} else {
		o.Subtype = *all.Subtype
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
