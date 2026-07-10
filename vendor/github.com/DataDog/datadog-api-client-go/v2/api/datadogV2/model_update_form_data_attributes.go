// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFormDataAttributes The attributes for updating a form.
type UpdateFormDataAttributes struct {
	// The fields to update on a form. At least one field must be provided.
	FormUpdate FormUpdateAttributes `json:"form_update"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateFormDataAttributes instantiates a new UpdateFormDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateFormDataAttributes(formUpdate FormUpdateAttributes) *UpdateFormDataAttributes {
	this := UpdateFormDataAttributes{}
	this.FormUpdate = formUpdate
	return &this
}

// NewUpdateFormDataAttributesWithDefaults instantiates a new UpdateFormDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateFormDataAttributesWithDefaults() *UpdateFormDataAttributes {
	this := UpdateFormDataAttributes{}
	return &this
}

// GetFormUpdate returns the FormUpdate field value.
func (o *UpdateFormDataAttributes) GetFormUpdate() FormUpdateAttributes {
	if o == nil {
		var ret FormUpdateAttributes
		return ret
	}
	return o.FormUpdate
}

// GetFormUpdateOk returns a tuple with the FormUpdate field value
// and a boolean to check if the value has been set.
func (o *UpdateFormDataAttributes) GetFormUpdateOk() (*FormUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormUpdate, true
}

// SetFormUpdate sets field value.
func (o *UpdateFormDataAttributes) SetFormUpdate(v FormUpdateAttributes) {
	o.FormUpdate = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateFormDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["form_update"] = o.FormUpdate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateFormDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FormUpdate *FormUpdateAttributes `json:"form_update"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FormUpdate == nil {
		return fmt.Errorf("required field form_update missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"form_update"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.FormUpdate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FormUpdate = *all.FormUpdate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
