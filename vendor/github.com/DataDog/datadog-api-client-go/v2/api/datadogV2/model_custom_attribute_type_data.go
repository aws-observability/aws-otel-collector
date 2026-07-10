// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomAttributeTypeData Type-specific configuration for the custom attribute. For SELECT-type attributes, this contains the list of allowed options.
type CustomAttributeTypeData struct {
	// Options for SELECT type custom attributes.
	Options []CustomAttributeSelectOption `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomAttributeTypeData instantiates a new CustomAttributeTypeData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomAttributeTypeData() *CustomAttributeTypeData {
	this := CustomAttributeTypeData{}
	return &this
}

// NewCustomAttributeTypeDataWithDefaults instantiates a new CustomAttributeTypeData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomAttributeTypeDataWithDefaults() *CustomAttributeTypeData {
	this := CustomAttributeTypeData{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CustomAttributeTypeData) GetOptions() []CustomAttributeSelectOption {
	if o == nil || o.Options == nil {
		var ret []CustomAttributeSelectOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeTypeData) GetOptionsOk() (*[]CustomAttributeSelectOption, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CustomAttributeTypeData) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given []CustomAttributeSelectOption and assigns it to the Options field.
func (o *CustomAttributeTypeData) SetOptions(v []CustomAttributeSelectOption) {
	o.Options = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomAttributeTypeData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomAttributeTypeData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Options []CustomAttributeSelectOption `json:"options,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"options"})
	} else {
		return err
	}
	o.Options = all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
