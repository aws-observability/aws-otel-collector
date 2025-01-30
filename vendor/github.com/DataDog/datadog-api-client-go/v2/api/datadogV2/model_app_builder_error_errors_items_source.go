// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppBuilderErrorErrorsItemsSource The definition of `AppBuilderErrorErrorsItemsSource` object.
type AppBuilderErrorErrorsItemsSource struct {
	// The `source` `parameter`.
	Parameter *string `json:"parameter,omitempty"`
	// The `source` `pointer`.
	Pointer *string `json:"pointer,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAppBuilderErrorErrorsItemsSource instantiates a new AppBuilderErrorErrorsItemsSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAppBuilderErrorErrorsItemsSource() *AppBuilderErrorErrorsItemsSource {
	this := AppBuilderErrorErrorsItemsSource{}
	return &this
}

// NewAppBuilderErrorErrorsItemsSourceWithDefaults instantiates a new AppBuilderErrorErrorsItemsSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAppBuilderErrorErrorsItemsSourceWithDefaults() *AppBuilderErrorErrorsItemsSource {
	this := AppBuilderErrorErrorsItemsSource{}
	return &this
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *AppBuilderErrorErrorsItemsSource) GetParameter() string {
	if o == nil || o.Parameter == nil {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBuilderErrorErrorsItemsSource) GetParameterOk() (*string, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *AppBuilderErrorErrorsItemsSource) HasParameter() bool {
	return o != nil && o.Parameter != nil
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *AppBuilderErrorErrorsItemsSource) SetParameter(v string) {
	o.Parameter = &v
}

// GetPointer returns the Pointer field value if set, zero value otherwise.
func (o *AppBuilderErrorErrorsItemsSource) GetPointer() string {
	if o == nil || o.Pointer == nil {
		var ret string
		return ret
	}
	return *o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBuilderErrorErrorsItemsSource) GetPointerOk() (*string, bool) {
	if o == nil || o.Pointer == nil {
		return nil, false
	}
	return o.Pointer, true
}

// HasPointer returns a boolean if a field has been set.
func (o *AppBuilderErrorErrorsItemsSource) HasPointer() bool {
	return o != nil && o.Pointer != nil
}

// SetPointer gets a reference to the given string and assigns it to the Pointer field.
func (o *AppBuilderErrorErrorsItemsSource) SetPointer(v string) {
	o.Pointer = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AppBuilderErrorErrorsItemsSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	if o.Pointer != nil {
		toSerialize["pointer"] = o.Pointer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AppBuilderErrorErrorsItemsSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Parameter *string `json:"parameter,omitempty"`
		Pointer   *string `json:"pointer,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"parameter", "pointer"})
	} else {
		return err
	}
	o.Parameter = all.Parameter
	o.Pointer = all.Pointer

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
