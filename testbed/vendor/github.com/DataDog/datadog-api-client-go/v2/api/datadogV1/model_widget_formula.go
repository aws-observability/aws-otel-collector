// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetFormula Formula to be used in a widget query.
type WidgetFormula struct {
	// Expression alias.
	Alias *string `json:"alias,omitempty"`
	// Define a display mode for the table cell.
	CellDisplayMode *TableWidgetCellDisplayMode `json:"cell_display_mode,omitempty"`
	// Cell display mode options for the widget formula. (only if `cell_display_mode` is set to `trend`).
	CellDisplayModeOptions *WidgetFormulaCellDisplayModeOptions `json:"cell_display_mode_options,omitempty"`
	// List of conditional formats.
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty"`
	// String expression built from queries, formulas, and functions.
	Formula string `json:"formula"`
	// Options for limiting results returned.
	Limit *WidgetFormulaLimit `json:"limit,omitempty"`
	// Number format options for the widget.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// Styling options for widget formulas.
	Style *WidgetFormulaStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetFormula instantiates a new WidgetFormula object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetFormula(formula string) *WidgetFormula {
	this := WidgetFormula{}
	this.Formula = formula
	return &this
}

// NewWidgetFormulaWithDefaults instantiates a new WidgetFormula object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetFormulaWithDefaults() *WidgetFormula {
	this := WidgetFormula{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *WidgetFormula) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *WidgetFormula) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *WidgetFormula) SetAlias(v string) {
	o.Alias = &v
}

// GetCellDisplayMode returns the CellDisplayMode field value if set, zero value otherwise.
func (o *WidgetFormula) GetCellDisplayMode() TableWidgetCellDisplayMode {
	if o == nil || o.CellDisplayMode == nil {
		var ret TableWidgetCellDisplayMode
		return ret
	}
	return *o.CellDisplayMode
}

// GetCellDisplayModeOk returns a tuple with the CellDisplayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetCellDisplayModeOk() (*TableWidgetCellDisplayMode, bool) {
	if o == nil || o.CellDisplayMode == nil {
		return nil, false
	}
	return o.CellDisplayMode, true
}

// HasCellDisplayMode returns a boolean if a field has been set.
func (o *WidgetFormula) HasCellDisplayMode() bool {
	return o != nil && o.CellDisplayMode != nil
}

// SetCellDisplayMode gets a reference to the given TableWidgetCellDisplayMode and assigns it to the CellDisplayMode field.
func (o *WidgetFormula) SetCellDisplayMode(v TableWidgetCellDisplayMode) {
	o.CellDisplayMode = &v
}

// GetCellDisplayModeOptions returns the CellDisplayModeOptions field value if set, zero value otherwise.
func (o *WidgetFormula) GetCellDisplayModeOptions() WidgetFormulaCellDisplayModeOptions {
	if o == nil || o.CellDisplayModeOptions == nil {
		var ret WidgetFormulaCellDisplayModeOptions
		return ret
	}
	return *o.CellDisplayModeOptions
}

// GetCellDisplayModeOptionsOk returns a tuple with the CellDisplayModeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetCellDisplayModeOptionsOk() (*WidgetFormulaCellDisplayModeOptions, bool) {
	if o == nil || o.CellDisplayModeOptions == nil {
		return nil, false
	}
	return o.CellDisplayModeOptions, true
}

// HasCellDisplayModeOptions returns a boolean if a field has been set.
func (o *WidgetFormula) HasCellDisplayModeOptions() bool {
	return o != nil && o.CellDisplayModeOptions != nil
}

// SetCellDisplayModeOptions gets a reference to the given WidgetFormulaCellDisplayModeOptions and assigns it to the CellDisplayModeOptions field.
func (o *WidgetFormula) SetCellDisplayModeOptions(v WidgetFormulaCellDisplayModeOptions) {
	o.CellDisplayModeOptions = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *WidgetFormula) GetConditionalFormats() []WidgetConditionalFormat {
	if o == nil || o.ConditionalFormats == nil {
		var ret []WidgetConditionalFormat
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetConditionalFormatsOk() (*[]WidgetConditionalFormat, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *WidgetFormula) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []WidgetConditionalFormat and assigns it to the ConditionalFormats field.
func (o *WidgetFormula) SetConditionalFormats(v []WidgetConditionalFormat) {
	o.ConditionalFormats = v
}

// GetFormula returns the Formula field value.
func (o *WidgetFormula) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *WidgetFormula) SetFormula(v string) {
	o.Formula = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *WidgetFormula) GetLimit() WidgetFormulaLimit {
	if o == nil || o.Limit == nil {
		var ret WidgetFormulaLimit
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetLimitOk() (*WidgetFormulaLimit, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *WidgetFormula) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given WidgetFormulaLimit and assigns it to the Limit field.
func (o *WidgetFormula) SetLimit(v WidgetFormulaLimit) {
	o.Limit = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *WidgetFormula) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *WidgetFormula) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *WidgetFormula) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *WidgetFormula) GetStyle() WidgetFormulaStyle {
	if o == nil || o.Style == nil {
		var ret WidgetFormulaStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetStyleOk() (*WidgetFormulaStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *WidgetFormula) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given WidgetFormulaStyle and assigns it to the Style field.
func (o *WidgetFormula) SetStyle(v WidgetFormulaStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetFormula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.CellDisplayMode != nil {
		toSerialize["cell_display_mode"] = o.CellDisplayMode
	}
	if o.CellDisplayModeOptions != nil {
		toSerialize["cell_display_mode_options"] = o.CellDisplayModeOptions
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["formula"] = o.Formula
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NumberFormat != nil {
		toSerialize["number_format"] = o.NumberFormat
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetFormula) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias                  *string                              `json:"alias,omitempty"`
		CellDisplayMode        *TableWidgetCellDisplayMode          `json:"cell_display_mode,omitempty"`
		CellDisplayModeOptions *WidgetFormulaCellDisplayModeOptions `json:"cell_display_mode_options,omitempty"`
		ConditionalFormats     []WidgetConditionalFormat            `json:"conditional_formats,omitempty"`
		Formula                *string                              `json:"formula"`
		Limit                  *WidgetFormulaLimit                  `json:"limit,omitempty"`
		NumberFormat           *WidgetNumberFormat                  `json:"number_format,omitempty"`
		Style                  *WidgetFormulaStyle                  `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias", "cell_display_mode", "cell_display_mode_options", "conditional_formats", "formula", "limit", "number_format", "style"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if all.CellDisplayMode != nil && !all.CellDisplayMode.IsValid() {
		hasInvalidField = true
	} else {
		o.CellDisplayMode = all.CellDisplayMode
	}
	if all.CellDisplayModeOptions != nil && all.CellDisplayModeOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CellDisplayModeOptions = all.CellDisplayModeOptions
	o.ConditionalFormats = all.ConditionalFormats
	o.Formula = *all.Formula
	if all.Limit != nil && all.Limit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Limit = all.Limit
	if all.NumberFormat != nil && all.NumberFormat.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumberFormat = all.NumberFormat
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
