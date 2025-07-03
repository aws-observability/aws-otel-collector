// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPHeaderUpdate The definition of `HTTPHeaderUpdate` object.
type HTTPHeaderUpdate struct {
	// Should the header be deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// The `HTTPHeaderUpdate` `name`.
	Name string `json:"name"`
	// The `HTTPHeaderUpdate` `value`.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHTTPHeaderUpdate instantiates a new HTTPHeaderUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHTTPHeaderUpdate(name string) *HTTPHeaderUpdate {
	this := HTTPHeaderUpdate{}
	this.Name = name
	return &this
}

// NewHTTPHeaderUpdateWithDefaults instantiates a new HTTPHeaderUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHTTPHeaderUpdateWithDefaults() *HTTPHeaderUpdate {
	this := HTTPHeaderUpdate{}
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *HTTPHeaderUpdate) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPHeaderUpdate) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *HTTPHeaderUpdate) HasDeleted() bool {
	return o != nil && o.Deleted != nil
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *HTTPHeaderUpdate) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetName returns the Name field value.
func (o *HTTPHeaderUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HTTPHeaderUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *HTTPHeaderUpdate) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HTTPHeaderUpdate) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPHeaderUpdate) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HTTPHeaderUpdate) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HTTPHeaderUpdate) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HTTPHeaderUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	toSerialize["name"] = o.Name
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HTTPHeaderUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Deleted *bool   `json:"deleted,omitempty"`
		Name    *string `json:"name"`
		Value   *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deleted", "name", "value"})
	} else {
		return err
	}
	o.Deleted = all.Deleted
	o.Name = *all.Name
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
