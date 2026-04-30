// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsOccurrenceFilter Filter for occurrence-based queries.
type ProductAnalyticsOccurrenceFilter struct {
	// Additional metadata.
	Meta map[string]string `json:"meta,omitempty"`
	// Comparison operator (=, >=, <=, >, <).
	Operator string `json:"operator"`
	// The occurrence count threshold as a string.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsOccurrenceFilter instantiates a new ProductAnalyticsOccurrenceFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsOccurrenceFilter(operator string, value string) *ProductAnalyticsOccurrenceFilter {
	this := ProductAnalyticsOccurrenceFilter{}
	this.Operator = operator
	this.Value = value
	return &this
}

// NewProductAnalyticsOccurrenceFilterWithDefaults instantiates a new ProductAnalyticsOccurrenceFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsOccurrenceFilterWithDefaults() *ProductAnalyticsOccurrenceFilter {
	this := ProductAnalyticsOccurrenceFilter{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ProductAnalyticsOccurrenceFilter) GetMeta() map[string]string {
	if o == nil || o.Meta == nil {
		var ret map[string]string
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsOccurrenceFilter) GetMetaOk() (*map[string]string, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return &o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ProductAnalyticsOccurrenceFilter) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given map[string]string and assigns it to the Meta field.
func (o *ProductAnalyticsOccurrenceFilter) SetMeta(v map[string]string) {
	o.Meta = v
}

// GetOperator returns the Operator field value.
func (o *ProductAnalyticsOccurrenceFilter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsOccurrenceFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *ProductAnalyticsOccurrenceFilter) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value.
func (o *ProductAnalyticsOccurrenceFilter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsOccurrenceFilter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ProductAnalyticsOccurrenceFilter) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsOccurrenceFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsOccurrenceFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Meta     map[string]string `json:"meta,omitempty"`
		Operator *string           `json:"operator"`
		Value    *string           `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"meta", "operator", "value"})
	} else {
		return err
	}
	o.Meta = all.Meta
	o.Operator = *all.Operator
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
