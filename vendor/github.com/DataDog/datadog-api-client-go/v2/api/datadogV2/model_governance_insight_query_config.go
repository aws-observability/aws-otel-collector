// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightQueryConfig Query execution context that allows the frontend to execute insight queries directly.
type GovernanceInsightQueryConfig struct {
	// The chart type the frontend should use to render the insight.
	ChartType *string `json:"chart_type,omitempty"`
	// The window used for the previous value comparison, for example `week` or `month`.
	ComparisonShift string `json:"comparison_shift"`
	// The default value to display when no data is available.
	DefaultValue *int64 `json:"default_value,omitempty"`
	// Whether an increase in the value is good, bad, or neutral. One of `neutral`,
	// `increase_better`, or `decrease_better`.
	Directionality *string `json:"directionality,omitempty"`
	// The number of days the insight value is computed over.
	EffectiveTimeWindowDays int64 `json:"effective_time_window_days"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceInsightQueryConfig instantiates a new GovernanceInsightQueryConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceInsightQueryConfig(comparisonShift string, effectiveTimeWindowDays int64) *GovernanceInsightQueryConfig {
	this := GovernanceInsightQueryConfig{}
	this.ComparisonShift = comparisonShift
	this.EffectiveTimeWindowDays = effectiveTimeWindowDays
	return &this
}

// NewGovernanceInsightQueryConfigWithDefaults instantiates a new GovernanceInsightQueryConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceInsightQueryConfigWithDefaults() *GovernanceInsightQueryConfig {
	this := GovernanceInsightQueryConfig{}
	return &this
}

// GetChartType returns the ChartType field value if set, zero value otherwise.
func (o *GovernanceInsightQueryConfig) GetChartType() string {
	if o == nil || o.ChartType == nil {
		var ret string
		return ret
	}
	return *o.ChartType
}

// GetChartTypeOk returns a tuple with the ChartType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightQueryConfig) GetChartTypeOk() (*string, bool) {
	if o == nil || o.ChartType == nil {
		return nil, false
	}
	return o.ChartType, true
}

// HasChartType returns a boolean if a field has been set.
func (o *GovernanceInsightQueryConfig) HasChartType() bool {
	return o != nil && o.ChartType != nil
}

// SetChartType gets a reference to the given string and assigns it to the ChartType field.
func (o *GovernanceInsightQueryConfig) SetChartType(v string) {
	o.ChartType = &v
}

// GetComparisonShift returns the ComparisonShift field value.
func (o *GovernanceInsightQueryConfig) GetComparisonShift() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ComparisonShift
}

// GetComparisonShiftOk returns a tuple with the ComparisonShift field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightQueryConfig) GetComparisonShiftOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparisonShift, true
}

// SetComparisonShift sets field value.
func (o *GovernanceInsightQueryConfig) SetComparisonShift(v string) {
	o.ComparisonShift = v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *GovernanceInsightQueryConfig) GetDefaultValue() int64 {
	if o == nil || o.DefaultValue == nil {
		var ret int64
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightQueryConfig) GetDefaultValueOk() (*int64, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *GovernanceInsightQueryConfig) HasDefaultValue() bool {
	return o != nil && o.DefaultValue != nil
}

// SetDefaultValue gets a reference to the given int64 and assigns it to the DefaultValue field.
func (o *GovernanceInsightQueryConfig) SetDefaultValue(v int64) {
	o.DefaultValue = &v
}

// GetDirectionality returns the Directionality field value if set, zero value otherwise.
func (o *GovernanceInsightQueryConfig) GetDirectionality() string {
	if o == nil || o.Directionality == nil {
		var ret string
		return ret
	}
	return *o.Directionality
}

// GetDirectionalityOk returns a tuple with the Directionality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightQueryConfig) GetDirectionalityOk() (*string, bool) {
	if o == nil || o.Directionality == nil {
		return nil, false
	}
	return o.Directionality, true
}

// HasDirectionality returns a boolean if a field has been set.
func (o *GovernanceInsightQueryConfig) HasDirectionality() bool {
	return o != nil && o.Directionality != nil
}

// SetDirectionality gets a reference to the given string and assigns it to the Directionality field.
func (o *GovernanceInsightQueryConfig) SetDirectionality(v string) {
	o.Directionality = &v
}

// GetEffectiveTimeWindowDays returns the EffectiveTimeWindowDays field value.
func (o *GovernanceInsightQueryConfig) GetEffectiveTimeWindowDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EffectiveTimeWindowDays
}

// GetEffectiveTimeWindowDaysOk returns a tuple with the EffectiveTimeWindowDays field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightQueryConfig) GetEffectiveTimeWindowDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveTimeWindowDays, true
}

// SetEffectiveTimeWindowDays sets field value.
func (o *GovernanceInsightQueryConfig) SetEffectiveTimeWindowDays(v int64) {
	o.EffectiveTimeWindowDays = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceInsightQueryConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChartType != nil {
		toSerialize["chart_type"] = o.ChartType
	}
	toSerialize["comparison_shift"] = o.ComparisonShift
	if o.DefaultValue != nil {
		toSerialize["default_value"] = o.DefaultValue
	}
	if o.Directionality != nil {
		toSerialize["directionality"] = o.Directionality
	}
	toSerialize["effective_time_window_days"] = o.EffectiveTimeWindowDays

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceInsightQueryConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChartType               *string `json:"chart_type,omitempty"`
		ComparisonShift         *string `json:"comparison_shift"`
		DefaultValue            *int64  `json:"default_value,omitempty"`
		Directionality          *string `json:"directionality,omitempty"`
		EffectiveTimeWindowDays *int64  `json:"effective_time_window_days"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ComparisonShift == nil {
		return fmt.Errorf("required field comparison_shift missing")
	}
	if all.EffectiveTimeWindowDays == nil {
		return fmt.Errorf("required field effective_time_window_days missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"chart_type", "comparison_shift", "default_value", "directionality", "effective_time_window_days"})
	} else {
		return err
	}
	o.ChartType = all.ChartType
	o.ComparisonShift = *all.ComparisonShift
	o.DefaultValue = all.DefaultValue
	o.Directionality = all.Directionality
	o.EffectiveTimeWindowDays = *all.EffectiveTimeWindowDays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
