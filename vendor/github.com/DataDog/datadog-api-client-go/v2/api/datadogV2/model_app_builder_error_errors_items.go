// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppBuilderErrorErrorsItems The definition of `AppBuilderErrorErrorsItems` object.
type AppBuilderErrorErrorsItems struct {
	// The `items` `detail`.
	Detail *string `json:"detail,omitempty"`
	// The definition of `AppBuilderErrorErrorsItemsSource` object.
	Source *AppBuilderErrorErrorsItemsSource `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAppBuilderErrorErrorsItems instantiates a new AppBuilderErrorErrorsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAppBuilderErrorErrorsItems() *AppBuilderErrorErrorsItems {
	this := AppBuilderErrorErrorsItems{}
	return &this
}

// NewAppBuilderErrorErrorsItemsWithDefaults instantiates a new AppBuilderErrorErrorsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAppBuilderErrorErrorsItemsWithDefaults() *AppBuilderErrorErrorsItems {
	this := AppBuilderErrorErrorsItems{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *AppBuilderErrorErrorsItems) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBuilderErrorErrorsItems) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *AppBuilderErrorErrorsItems) HasDetail() bool {
	return o != nil && o.Detail != nil
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *AppBuilderErrorErrorsItems) SetDetail(v string) {
	o.Detail = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AppBuilderErrorErrorsItems) GetSource() AppBuilderErrorErrorsItemsSource {
	if o == nil || o.Source == nil {
		var ret AppBuilderErrorErrorsItemsSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBuilderErrorErrorsItems) GetSourceOk() (*AppBuilderErrorErrorsItemsSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AppBuilderErrorErrorsItems) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given AppBuilderErrorErrorsItemsSource and assigns it to the Source field.
func (o *AppBuilderErrorErrorsItems) SetSource(v AppBuilderErrorErrorsItemsSource) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AppBuilderErrorErrorsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AppBuilderErrorErrorsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Detail *string                           `json:"detail,omitempty"`
		Source *AppBuilderErrorErrorsItemsSource `json:"source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"detail", "source"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Detail = all.Detail
	if all.Source != nil && all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
