// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionAggregateQueryFilter Filter definition for aggregate filtered queries.
type MonitorFormulaAndFunctionAggregateQueryFilter struct {
	// Attribute from the base query to filter on.
	BaseAttribute string `json:"base_attribute"`
	// Whether to exclude matching records instead of including them.
	Exclude *bool `json:"exclude,omitempty"`
	// Attribute from the filter query to match against.
	FilterAttribute string `json:"filter_attribute"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionAggregateQueryFilter instantiates a new MonitorFormulaAndFunctionAggregateQueryFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionAggregateQueryFilter(baseAttribute string, filterAttribute string) *MonitorFormulaAndFunctionAggregateQueryFilter {
	this := MonitorFormulaAndFunctionAggregateQueryFilter{}
	this.BaseAttribute = baseAttribute
	var exclude bool = false
	this.Exclude = &exclude
	this.FilterAttribute = filterAttribute
	return &this
}

// NewMonitorFormulaAndFunctionAggregateQueryFilterWithDefaults instantiates a new MonitorFormulaAndFunctionAggregateQueryFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionAggregateQueryFilterWithDefaults() *MonitorFormulaAndFunctionAggregateQueryFilter {
	this := MonitorFormulaAndFunctionAggregateQueryFilter{}
	var exclude bool = false
	this.Exclude = &exclude
	return &this
}

// GetBaseAttribute returns the BaseAttribute field value.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) GetBaseAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BaseAttribute
}

// GetBaseAttributeOk returns a tuple with the BaseAttribute field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) GetBaseAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseAttribute, true
}

// SetBaseAttribute sets field value.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) SetBaseAttribute(v string) {
	o.BaseAttribute = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) GetExclude() bool {
	if o == nil || o.Exclude == nil {
		var ret bool
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) GetExcludeOk() (*bool, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) HasExclude() bool {
	return o != nil && o.Exclude != nil
}

// SetExclude gets a reference to the given bool and assigns it to the Exclude field.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) SetExclude(v bool) {
	o.Exclude = &v
}

// GetFilterAttribute returns the FilterAttribute field value.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) GetFilterAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FilterAttribute
}

// GetFilterAttributeOk returns a tuple with the FilterAttribute field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) GetFilterAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterAttribute, true
}

// SetFilterAttribute sets field value.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) SetFilterAttribute(v string) {
	o.FilterAttribute = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionAggregateQueryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_attribute"] = o.BaseAttribute
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	toSerialize["filter_attribute"] = o.FilterAttribute
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionAggregateQueryFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseAttribute   *string `json:"base_attribute"`
		Exclude         *bool   `json:"exclude,omitempty"`
		FilterAttribute *string `json:"filter_attribute"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseAttribute == nil {
		return fmt.Errorf("required field base_attribute missing")
	}
	if all.FilterAttribute == nil {
		return fmt.Errorf("required field filter_attribute missing")
	}
	o.BaseAttribute = *all.BaseAttribute
	o.Exclude = all.Exclude
	o.FilterAttribute = *all.FilterAttribute

	return nil
}
