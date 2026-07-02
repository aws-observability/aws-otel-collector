// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipInferenceListAttributes The attributes of the ownership inferences collection response.
type OwnershipInferenceListAttributes struct {
	// The list of inferences for a resource, with one inference per owner type.
	Items []OwnershipInferenceItem `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipInferenceListAttributes instantiates a new OwnershipInferenceListAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipInferenceListAttributes(items []OwnershipInferenceItem) *OwnershipInferenceListAttributes {
	this := OwnershipInferenceListAttributes{}
	this.Items = items
	return &this
}

// NewOwnershipInferenceListAttributesWithDefaults instantiates a new OwnershipInferenceListAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipInferenceListAttributesWithDefaults() *OwnershipInferenceListAttributes {
	this := OwnershipInferenceListAttributes{}
	return &this
}

// GetItems returns the Items field value.
func (o *OwnershipInferenceListAttributes) GetItems() []OwnershipInferenceItem {
	if o == nil {
		var ret []OwnershipInferenceItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *OwnershipInferenceListAttributes) GetItemsOk() (*[]OwnershipInferenceItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *OwnershipInferenceListAttributes) SetItems(v []OwnershipInferenceItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipInferenceListAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipInferenceListAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items *[]OwnershipInferenceItem `json:"items"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"items"})
	} else {
		return err
	}
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
