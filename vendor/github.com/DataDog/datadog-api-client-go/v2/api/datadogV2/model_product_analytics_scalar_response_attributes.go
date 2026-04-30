// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsScalarResponseAttributes Attributes of a scalar analytics response, containing the result columns.
type ProductAnalyticsScalarResponseAttributes struct {
	// The list of result columns, each containing values and metadata.
	Columns []ProductAnalyticsScalarColumn `json:"columns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsScalarResponseAttributes instantiates a new ProductAnalyticsScalarResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsScalarResponseAttributes() *ProductAnalyticsScalarResponseAttributes {
	this := ProductAnalyticsScalarResponseAttributes{}
	return &this
}

// NewProductAnalyticsScalarResponseAttributesWithDefaults instantiates a new ProductAnalyticsScalarResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsScalarResponseAttributesWithDefaults() *ProductAnalyticsScalarResponseAttributes {
	this := ProductAnalyticsScalarResponseAttributes{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *ProductAnalyticsScalarResponseAttributes) GetColumns() []ProductAnalyticsScalarColumn {
	if o == nil || o.Columns == nil {
		var ret []ProductAnalyticsScalarColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsScalarResponseAttributes) GetColumnsOk() (*[]ProductAnalyticsScalarColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *ProductAnalyticsScalarResponseAttributes) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []ProductAnalyticsScalarColumn and assigns it to the Columns field.
func (o *ProductAnalyticsScalarResponseAttributes) SetColumns(v []ProductAnalyticsScalarColumn) {
	o.Columns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsScalarResponseAttributes) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsScalarResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns []ProductAnalyticsScalarColumn `json:"columns,omitempty"`
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
