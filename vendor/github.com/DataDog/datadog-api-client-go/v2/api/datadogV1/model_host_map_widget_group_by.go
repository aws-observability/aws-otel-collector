// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetGroupBy Defines a grouping dimension for the infrastructure host map.
type HostMapWidgetGroupBy struct {
	// Column name from the entity table (for example, `cloud_provider`, `tags`, `labels`).
	Column string `json:"column"`
	// Key within the column for nested attribute types (for example, `service` within `tags`).
	Key *string `json:"key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHostMapWidgetGroupBy instantiates a new HostMapWidgetGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostMapWidgetGroupBy(column string) *HostMapWidgetGroupBy {
	this := HostMapWidgetGroupBy{}
	this.Column = column
	return &this
}

// NewHostMapWidgetGroupByWithDefaults instantiates a new HostMapWidgetGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostMapWidgetGroupByWithDefaults() *HostMapWidgetGroupBy {
	this := HostMapWidgetGroupBy{}
	return &this
}

// GetColumn returns the Column field value.
func (o *HostMapWidgetGroupBy) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetGroupBy) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *HostMapWidgetGroupBy) SetColumn(v string) {
	o.Column = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *HostMapWidgetGroupBy) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetGroupBy) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *HostMapWidgetGroupBy) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *HostMapWidgetGroupBy) SetKey(v string) {
	o.Key = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostMapWidgetGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostMapWidgetGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string `json:"column"`
		Key    *string `json:"key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column", "key"})
	} else {
		return err
	}
	o.Column = *all.Column
	o.Key = all.Key

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
