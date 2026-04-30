// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectColumnsConfig Project columns configuration.
type ProjectColumnsConfig struct {
	// List of column configurations for the project board view.
	Columns []ProjectColumnsConfigColumnsItems `json:"columns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectColumnsConfig instantiates a new ProjectColumnsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectColumnsConfig() *ProjectColumnsConfig {
	this := ProjectColumnsConfig{}
	return &this
}

// NewProjectColumnsConfigWithDefaults instantiates a new ProjectColumnsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectColumnsConfigWithDefaults() *ProjectColumnsConfig {
	this := ProjectColumnsConfig{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *ProjectColumnsConfig) GetColumns() []ProjectColumnsConfigColumnsItems {
	if o == nil || o.Columns == nil {
		var ret []ProjectColumnsConfigColumnsItems
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectColumnsConfig) GetColumnsOk() (*[]ProjectColumnsConfigColumnsItems, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *ProjectColumnsConfig) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []ProjectColumnsConfigColumnsItems and assigns it to the Columns field.
func (o *ProjectColumnsConfig) SetColumns(v []ProjectColumnsConfigColumnsItems) {
	o.Columns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectColumnsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectColumnsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns []ProjectColumnsConfigColumnsItems `json:"columns,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns"})
	} else {
		return err
	}
	o.Columns = all.Columns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
