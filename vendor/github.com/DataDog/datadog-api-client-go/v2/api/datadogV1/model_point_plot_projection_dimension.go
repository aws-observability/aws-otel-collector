// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotProjectionDimension Dimension mapping for the point plot projection.
type PointPlotProjectionDimension struct {
	// Alias for the column.
	Alias *string `json:"alias,omitempty"`
	// Source column name from the dataset.
	Column string `json:"column"`
	// Dimension of the point plot.
	Dimension PointPlotDimension `json:"dimension"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPointPlotProjectionDimension instantiates a new PointPlotProjectionDimension object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPointPlotProjectionDimension(column string, dimension PointPlotDimension) *PointPlotProjectionDimension {
	this := PointPlotProjectionDimension{}
	this.Column = column
	this.Dimension = dimension
	return &this
}

// NewPointPlotProjectionDimensionWithDefaults instantiates a new PointPlotProjectionDimension object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPointPlotProjectionDimensionWithDefaults() *PointPlotProjectionDimension {
	this := PointPlotProjectionDimension{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *PointPlotProjectionDimension) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotProjectionDimension) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *PointPlotProjectionDimension) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *PointPlotProjectionDimension) SetAlias(v string) {
	o.Alias = &v
}

// GetColumn returns the Column field value.
func (o *PointPlotProjectionDimension) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *PointPlotProjectionDimension) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *PointPlotProjectionDimension) SetColumn(v string) {
	o.Column = v
}

// GetDimension returns the Dimension field value.
func (o *PointPlotProjectionDimension) GetDimension() PointPlotDimension {
	if o == nil {
		var ret PointPlotDimension
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *PointPlotProjectionDimension) GetDimensionOk() (*PointPlotDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value.
func (o *PointPlotProjectionDimension) SetDimension(v PointPlotDimension) {
	o.Dimension = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PointPlotProjectionDimension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["column"] = o.Column
	toSerialize["dimension"] = o.Dimension

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PointPlotProjectionDimension) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias     *string             `json:"alias,omitempty"`
		Column    *string             `json:"column"`
		Dimension *PointPlotDimension `json:"dimension"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Dimension == nil {
		return fmt.Errorf("required field dimension missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias", "column", "dimension"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Alias = all.Alias
	o.Column = *all.Column
	if !all.Dimension.IsValid() {
		hasInvalidField = true
	} else {
		o.Dimension = *all.Dimension
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
