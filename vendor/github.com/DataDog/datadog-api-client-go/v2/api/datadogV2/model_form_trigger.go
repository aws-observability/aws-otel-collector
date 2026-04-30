// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormTrigger Trigger a workflow from a Form.
type FormTrigger struct {
	// The form UUID.
	FormId *string `json:"formId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormTrigger instantiates a new FormTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormTrigger() *FormTrigger {
	this := FormTrigger{}
	return &this
}

// NewFormTriggerWithDefaults instantiates a new FormTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormTriggerWithDefaults() *FormTrigger {
	this := FormTrigger{}
	return &this
}

// GetFormId returns the FormId field value if set, zero value otherwise.
func (o *FormTrigger) GetFormId() string {
	if o == nil || o.FormId == nil {
		var ret string
		return ret
	}
	return *o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormTrigger) GetFormIdOk() (*string, bool) {
	if o == nil || o.FormId == nil {
		return nil, false
	}
	return o.FormId, true
}

// HasFormId returns a boolean if a field has been set.
func (o *FormTrigger) HasFormId() bool {
	return o != nil && o.FormId != nil
}

// SetFormId gets a reference to the given string and assigns it to the FormId field.
func (o *FormTrigger) SetFormId(v string) {
	o.FormId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FormId != nil {
		toSerialize["formId"] = o.FormId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FormId *string `json:"formId,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formId"})
	} else {
		return err
	}
	o.FormId = all.FormId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
