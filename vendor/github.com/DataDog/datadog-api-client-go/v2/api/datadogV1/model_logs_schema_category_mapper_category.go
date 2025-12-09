// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSchemaCategoryMapperCategory Object describing the logs filter with corresponding category ID and name assignment.
type LogsSchemaCategoryMapperCategory struct {
	// Filter for logs.
	Filter LogsFilter `json:"filter"`
	// ID to inject into the category.
	Id int64 `json:"id"`
	// Value to assign to target schema field.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsSchemaCategoryMapperCategory instantiates a new LogsSchemaCategoryMapperCategory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsSchemaCategoryMapperCategory(filter LogsFilter, id int64, name string) *LogsSchemaCategoryMapperCategory {
	this := LogsSchemaCategoryMapperCategory{}
	this.Filter = filter
	this.Id = id
	this.Name = name
	return &this
}

// NewLogsSchemaCategoryMapperCategoryWithDefaults instantiates a new LogsSchemaCategoryMapperCategory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsSchemaCategoryMapperCategoryWithDefaults() *LogsSchemaCategoryMapperCategory {
	this := LogsSchemaCategoryMapperCategory{}
	return &this
}

// GetFilter returns the Filter field value.
func (o *LogsSchemaCategoryMapperCategory) GetFilter() LogsFilter {
	if o == nil {
		var ret LogsFilter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaCategoryMapperCategory) GetFilterOk() (*LogsFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *LogsSchemaCategoryMapperCategory) SetFilter(v LogsFilter) {
	o.Filter = v
}

// GetId returns the Id field value.
func (o *LogsSchemaCategoryMapperCategory) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaCategoryMapperCategory) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LogsSchemaCategoryMapperCategory) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *LogsSchemaCategoryMapperCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaCategoryMapperCategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LogsSchemaCategoryMapperCategory) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsSchemaCategoryMapperCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["filter"] = o.Filter
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsSchemaCategoryMapperCategory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter *LogsFilter `json:"filter"`
		Id     *int64      `json:"id"`
		Name   *string     `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "id", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = *all.Filter
	o.Id = *all.Id
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
