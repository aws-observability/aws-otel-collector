// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionEventQueryGroupByFields Flat group by configuration using multiple event facet fields.
type FormulaAndFunctionEventQueryGroupByFields struct {
	// List of event facets to group by.
	Fields []string `json:"fields"`
	// Number of groups to return.
	Limit *int64 `json:"limit,omitempty"`
	// Options for sorting group by results.
	Sort *FormulaAndFunctionEventQueryGroupBySort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormulaAndFunctionEventQueryGroupByFields instantiates a new FormulaAndFunctionEventQueryGroupByFields object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaAndFunctionEventQueryGroupByFields(fields []string) *FormulaAndFunctionEventQueryGroupByFields {
	this := FormulaAndFunctionEventQueryGroupByFields{}
	this.Fields = fields
	return &this
}

// NewFormulaAndFunctionEventQueryGroupByFieldsWithDefaults instantiates a new FormulaAndFunctionEventQueryGroupByFields object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaAndFunctionEventQueryGroupByFieldsWithDefaults() *FormulaAndFunctionEventQueryGroupByFields {
	this := FormulaAndFunctionEventQueryGroupByFields{}
	return &this
}

// GetFields returns the Fields field value.
func (o *FormulaAndFunctionEventQueryGroupByFields) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryGroupByFields) GetFieldsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value.
func (o *FormulaAndFunctionEventQueryGroupByFields) SetFields(v []string) {
	o.Fields = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryGroupByFields) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryGroupByFields) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryGroupByFields) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *FormulaAndFunctionEventQueryGroupByFields) SetLimit(v int64) {
	o.Limit = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryGroupByFields) GetSort() FormulaAndFunctionEventQueryGroupBySort {
	if o == nil || o.Sort == nil {
		var ret FormulaAndFunctionEventQueryGroupBySort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryGroupByFields) GetSortOk() (*FormulaAndFunctionEventQueryGroupBySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryGroupByFields) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given FormulaAndFunctionEventQueryGroupBySort and assigns it to the Sort field.
func (o *FormulaAndFunctionEventQueryGroupByFields) SetSort(v FormulaAndFunctionEventQueryGroupBySort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaAndFunctionEventQueryGroupByFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fields"] = o.Fields
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaAndFunctionEventQueryGroupByFields) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields *[]string                                `json:"fields"`
		Limit  *int64                                   `json:"limit,omitempty"`
		Sort   *FormulaAndFunctionEventQueryGroupBySort `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fields == nil {
		return fmt.Errorf("required field fields missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "limit", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Fields = *all.Fields
	o.Limit = all.Limit
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
