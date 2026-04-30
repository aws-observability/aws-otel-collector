// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionReferenceTableColumn A column definition for reference table queries.
type MonitorFormulaAndFunctionReferenceTableColumn struct {
	// Optional alias for the column.
	Alias *string `json:"alias,omitempty"`
	// Name of the column.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionReferenceTableColumn instantiates a new MonitorFormulaAndFunctionReferenceTableColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionReferenceTableColumn(name string) *MonitorFormulaAndFunctionReferenceTableColumn {
	this := MonitorFormulaAndFunctionReferenceTableColumn{}
	this.Name = name
	return &this
}

// NewMonitorFormulaAndFunctionReferenceTableColumnWithDefaults instantiates a new MonitorFormulaAndFunctionReferenceTableColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionReferenceTableColumnWithDefaults() *MonitorFormulaAndFunctionReferenceTableColumn {
	this := MonitorFormulaAndFunctionReferenceTableColumn{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionReferenceTableColumn) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionReferenceTableColumn) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionReferenceTableColumn) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *MonitorFormulaAndFunctionReferenceTableColumn) SetAlias(v string) {
	o.Alias = &v
}

// GetName returns the Name field value.
func (o *MonitorFormulaAndFunctionReferenceTableColumn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionReferenceTableColumn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MonitorFormulaAndFunctionReferenceTableColumn) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionReferenceTableColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionReferenceTableColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias *string `json:"alias,omitempty"`
		Name  *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	o.Alias = all.Alias
	o.Name = *all.Name

	return nil
}
