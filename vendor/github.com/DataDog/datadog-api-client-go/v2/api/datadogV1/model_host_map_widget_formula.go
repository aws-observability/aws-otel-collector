// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetFormula Formula for the infrastructure host map widget that specifies both the expression
// and the visual dimension it populates.
type HostMapWidgetFormula struct {
	// Expression alias.
	Alias *string `json:"alias,omitempty"`
	// Visual dimension driven by a formula in the infrastructure host map widget.
	Dimension HostMapWidgetDimension `json:"dimension"`
	// String expression built from queries, formulas, and functions.
	Formula string `json:"formula"`
	// Number format options for the widget.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHostMapWidgetFormula instantiates a new HostMapWidgetFormula object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostMapWidgetFormula(dimension HostMapWidgetDimension, formula string) *HostMapWidgetFormula {
	this := HostMapWidgetFormula{}
	this.Dimension = dimension
	this.Formula = formula
	return &this
}

// NewHostMapWidgetFormulaWithDefaults instantiates a new HostMapWidgetFormula object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostMapWidgetFormulaWithDefaults() *HostMapWidgetFormula {
	this := HostMapWidgetFormula{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *HostMapWidgetFormula) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetFormula) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *HostMapWidgetFormula) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *HostMapWidgetFormula) SetAlias(v string) {
	o.Alias = &v
}

// GetDimension returns the Dimension field value.
func (o *HostMapWidgetFormula) GetDimension() HostMapWidgetDimension {
	if o == nil {
		var ret HostMapWidgetDimension
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetFormula) GetDimensionOk() (*HostMapWidgetDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value.
func (o *HostMapWidgetFormula) SetDimension(v HostMapWidgetDimension) {
	o.Dimension = v
}

// GetFormula returns the Formula field value.
func (o *HostMapWidgetFormula) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetFormula) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *HostMapWidgetFormula) SetFormula(v string) {
	o.Formula = v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *HostMapWidgetFormula) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetFormula) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *HostMapWidgetFormula) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *HostMapWidgetFormula) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostMapWidgetFormula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["dimension"] = o.Dimension
	toSerialize["formula"] = o.Formula
	if o.NumberFormat != nil {
		toSerialize["number_format"] = o.NumberFormat
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostMapWidgetFormula) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias        *string                 `json:"alias,omitempty"`
		Dimension    *HostMapWidgetDimension `json:"dimension"`
		Formula      *string                 `json:"formula"`
		NumberFormat *WidgetNumberFormat     `json:"number_format,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Dimension == nil {
		return fmt.Errorf("required field dimension missing")
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias", "dimension", "formula", "number_format"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if !all.Dimension.IsValid() {
		hasInvalidField = true
	} else {
		o.Dimension = *all.Dimension
	}
	o.Formula = *all.Formula
	if all.NumberFormat != nil && all.NumberFormat.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumberFormat = all.NumberFormat

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
