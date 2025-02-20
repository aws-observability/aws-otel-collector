// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOHistorySLIData An object that holds an SLI value and its associated data. It can represent an SLO's overall SLI value.
// This can also represent the SLI value for a specific monitor in multi-monitor SLOs, or a group in grouped SLOs.
type SLOHistorySLIData struct {
	// A mapping of threshold `timeframe` to the remaining error budget.
	ErrorBudgetRemaining map[string]float64 `json:"error_budget_remaining,omitempty"`
	// An array of error objects returned while querying the history data for the service level objective.
	Errors []SLOHistoryResponseErrorWithType `json:"errors,omitempty"`
	// For groups in a grouped SLO, this is the group name.
	Group *string `json:"group,omitempty"`
	// The state transition history for `monitor` or `time-slice` SLOs. It is represented as
	// an array of pairs. Each pair is an array containing the timestamp of the transition
	// as an integer in Unix epoch format in the first element, and the state as an integer in the
	// second element. An integer value of `0` for state means uptime, `1` means downtime, and `2` means no data.
	// Periods of no data count as uptime in time-slice SLOs, while for monitor SLOs, no data is counted
	// either as uptime or downtime depending on monitor settings. See
	// [SLO documentation](https://docs.datadoghq.com/service_management/service_level_objectives/monitor/#missing-data)
	// for detailed information.
	History [][]float64 `json:"history,omitempty"`
	// For `monitor` based SLOs, this is the last modified timestamp in epoch seconds of the monitor.
	MonitorModified *int64 `json:"monitor_modified,omitempty"`
	// For `monitor` based SLOs, this describes the type of monitor.
	MonitorType *string `json:"monitor_type,omitempty"`
	// For groups in a grouped SLO, this is the group name. For monitors in a multi-monitor SLO, this is the monitor name.
	Name *string `json:"name,omitempty"`
	// A mapping of threshold `timeframe` to number of accurate decimals, regardless of the from && to timestamp.
	Precision map[string]float64 `json:"precision,omitempty"`
	// For `monitor` based SLOs, when `true` this indicates that a replay is in progress to give an accurate uptime
	// calculation.
	Preview *bool `json:"preview,omitempty"`
	// The current SLI value of the SLO over the history window.
	SliValue datadog.NullableFloat64 `json:"sli_value,omitempty"`
	// The amount of decimal places the SLI value is accurate to for the given from `&&` to timestamp.
	SpanPrecision *float64 `json:"span_precision,omitempty"`
	// Use `sli_value` instead.
	// Deprecated
	Uptime datadog.NullableFloat64 `json:"uptime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSLOHistorySLIData instantiates a new SLOHistorySLIData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOHistorySLIData() *SLOHistorySLIData {
	this := SLOHistorySLIData{}
	return &this
}

// NewSLOHistorySLIDataWithDefaults instantiates a new SLOHistorySLIData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOHistorySLIDataWithDefaults() *SLOHistorySLIData {
	this := SLOHistorySLIData{}
	return &this
}

// GetErrorBudgetRemaining returns the ErrorBudgetRemaining field value if set, zero value otherwise.
func (o *SLOHistorySLIData) GetErrorBudgetRemaining() map[string]float64 {
	if o == nil || o.ErrorBudgetRemaining == nil {
		var ret map[string]float64
		return ret
	}
	return o.ErrorBudgetRemaining
}

// GetErrorBudgetRemainingOk returns a tuple with the ErrorBudgetRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistorySLIData) GetErrorBudgetRemainingOk() (*map[string]float64, bool) {
	if o == nil || o.ErrorBudgetRemaining == nil {
		return nil, false
	}
	return &o.ErrorBudgetRemaining, true
}

// HasErrorBudgetRemaining returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasErrorBudgetRemaining() bool {
	return o != nil && o.ErrorBudgetRemaining != nil
}

// SetErrorBudgetRemaining gets a reference to the given map[string]float64 and assigns it to the ErrorBudgetRemaining field.
func (o *SLOHistorySLIData) SetErrorBudgetRemaining(v map[string]float64) {
	o.ErrorBudgetRemaining = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SLOHistorySLIData) GetErrors() []SLOHistoryResponseErrorWithType {
	if o == nil || o.Errors == nil {
		var ret []SLOHistoryResponseErrorWithType
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistorySLIData) GetErrorsOk() (*[]SLOHistoryResponseErrorWithType, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given []SLOHistoryResponseErrorWithType and assigns it to the Errors field.
func (o *SLOHistorySLIData) SetErrors(v []SLOHistoryResponseErrorWithType) {
	o.Errors = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *SLOHistorySLIData) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistorySLIData) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *SLOHistorySLIData) SetGroup(v string) {
	o.Group = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *SLOHistorySLIData) GetHistory() [][]float64 {
	if o == nil || o.History == nil {
		var ret [][]float64
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistorySLIData) GetHistoryOk() (*[][]float64, bool) {
	if o == nil || o.History == nil {
		return nil, false
	}
	return &o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasHistory() bool {
	return o != nil && o.History != nil
}

// SetHistory gets a reference to the given [][]float64 and assigns it to the History field.
func (o *SLOHistorySLIData) SetHistory(v [][]float64) {
	o.History = v
}

// GetMonitorModified returns the MonitorModified field value if set, zero value otherwise.
func (o *SLOHistorySLIData) GetMonitorModified() int64 {
	if o == nil || o.MonitorModified == nil {
		var ret int64
		return ret
	}
	return *o.MonitorModified
}

// GetMonitorModifiedOk returns a tuple with the MonitorModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistorySLIData) GetMonitorModifiedOk() (*int64, bool) {
	if o == nil || o.MonitorModified == nil {
		return nil, false
	}
	return o.MonitorModified, true
}

// HasMonitorModified returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasMonitorModified() bool {
	return o != nil && o.MonitorModified != nil
}

// SetMonitorModified gets a reference to the given int64 and assigns it to the MonitorModified field.
func (o *SLOHistorySLIData) SetMonitorModified(v int64) {
	o.MonitorModified = &v
}

// GetMonitorType returns the MonitorType field value if set, zero value otherwise.
func (o *SLOHistorySLIData) GetMonitorType() string {
	if o == nil || o.MonitorType == nil {
		var ret string
		return ret
	}
	return *o.MonitorType
}

// GetMonitorTypeOk returns a tuple with the MonitorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistorySLIData) GetMonitorTypeOk() (*string, bool) {
	if o == nil || o.MonitorType == nil {
		return nil, false
	}
	return o.MonitorType, true
}

// HasMonitorType returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasMonitorType() bool {
	return o != nil && o.MonitorType != nil
}

// SetMonitorType gets a reference to the given string and assigns it to the MonitorType field.
func (o *SLOHistorySLIData) SetMonitorType(v string) {
	o.MonitorType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SLOHistorySLIData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistorySLIData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SLOHistorySLIData) SetName(v string) {
	o.Name = &v
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *SLOHistorySLIData) GetPrecision() map[string]float64 {
	if o == nil || o.Precision == nil {
		var ret map[string]float64
		return ret
	}
	return o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistorySLIData) GetPrecisionOk() (*map[string]float64, bool) {
	if o == nil || o.Precision == nil {
		return nil, false
	}
	return &o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasPrecision() bool {
	return o != nil && o.Precision != nil
}

// SetPrecision gets a reference to the given map[string]float64 and assigns it to the Precision field.
func (o *SLOHistorySLIData) SetPrecision(v map[string]float64) {
	o.Precision = v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *SLOHistorySLIData) GetPreview() bool {
	if o == nil || o.Preview == nil {
		var ret bool
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistorySLIData) GetPreviewOk() (*bool, bool) {
	if o == nil || o.Preview == nil {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasPreview() bool {
	return o != nil && o.Preview != nil
}

// SetPreview gets a reference to the given bool and assigns it to the Preview field.
func (o *SLOHistorySLIData) SetPreview(v bool) {
	o.Preview = &v
}

// GetSliValue returns the SliValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SLOHistorySLIData) GetSliValue() float64 {
	if o == nil || o.SliValue.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SliValue.Get()
}

// GetSliValueOk returns a tuple with the SliValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SLOHistorySLIData) GetSliValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SliValue.Get(), o.SliValue.IsSet()
}

// HasSliValue returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasSliValue() bool {
	return o != nil && o.SliValue.IsSet()
}

// SetSliValue gets a reference to the given datadog.NullableFloat64 and assigns it to the SliValue field.
func (o *SLOHistorySLIData) SetSliValue(v float64) {
	o.SliValue.Set(&v)
}

// SetSliValueNil sets the value for SliValue to be an explicit nil.
func (o *SLOHistorySLIData) SetSliValueNil() {
	o.SliValue.Set(nil)
}

// UnsetSliValue ensures that no value is present for SliValue, not even an explicit nil.
func (o *SLOHistorySLIData) UnsetSliValue() {
	o.SliValue.Unset()
}

// GetSpanPrecision returns the SpanPrecision field value if set, zero value otherwise.
func (o *SLOHistorySLIData) GetSpanPrecision() float64 {
	if o == nil || o.SpanPrecision == nil {
		var ret float64
		return ret
	}
	return *o.SpanPrecision
}

// GetSpanPrecisionOk returns a tuple with the SpanPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistorySLIData) GetSpanPrecisionOk() (*float64, bool) {
	if o == nil || o.SpanPrecision == nil {
		return nil, false
	}
	return o.SpanPrecision, true
}

// HasSpanPrecision returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasSpanPrecision() bool {
	return o != nil && o.SpanPrecision != nil
}

// SetSpanPrecision gets a reference to the given float64 and assigns it to the SpanPrecision field.
func (o *SLOHistorySLIData) SetSpanPrecision(v float64) {
	o.SpanPrecision = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *SLOHistorySLIData) GetUptime() float64 {
	if o == nil || o.Uptime.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Uptime.Get()
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
// Deprecated
func (o *SLOHistorySLIData) GetUptimeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uptime.Get(), o.Uptime.IsSet()
}

// HasUptime returns a boolean if a field has been set.
func (o *SLOHistorySLIData) HasUptime() bool {
	return o != nil && o.Uptime.IsSet()
}

// SetUptime gets a reference to the given datadog.NullableFloat64 and assigns it to the Uptime field.
// Deprecated
func (o *SLOHistorySLIData) SetUptime(v float64) {
	o.Uptime.Set(&v)
}

// SetUptimeNil sets the value for Uptime to be an explicit nil.
func (o *SLOHistorySLIData) SetUptimeNil() {
	o.Uptime.Set(nil)
}

// UnsetUptime ensures that no value is present for Uptime, not even an explicit nil.
func (o *SLOHistorySLIData) UnsetUptime() {
	o.Uptime.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOHistorySLIData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ErrorBudgetRemaining != nil {
		toSerialize["error_budget_remaining"] = o.ErrorBudgetRemaining
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.History != nil {
		toSerialize["history"] = o.History
	}
	if o.MonitorModified != nil {
		toSerialize["monitor_modified"] = o.MonitorModified
	}
	if o.MonitorType != nil {
		toSerialize["monitor_type"] = o.MonitorType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Precision != nil {
		toSerialize["precision"] = o.Precision
	}
	if o.Preview != nil {
		toSerialize["preview"] = o.Preview
	}
	if o.SliValue.IsSet() {
		toSerialize["sli_value"] = o.SliValue.Get()
	}
	if o.SpanPrecision != nil {
		toSerialize["span_precision"] = o.SpanPrecision
	}
	if o.Uptime.IsSet() {
		toSerialize["uptime"] = o.Uptime.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOHistorySLIData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorBudgetRemaining map[string]float64                `json:"error_budget_remaining,omitempty"`
		Errors               []SLOHistoryResponseErrorWithType `json:"errors,omitempty"`
		Group                *string                           `json:"group,omitempty"`
		History              [][]float64                       `json:"history,omitempty"`
		MonitorModified      *int64                            `json:"monitor_modified,omitempty"`
		MonitorType          *string                           `json:"monitor_type,omitempty"`
		Name                 *string                           `json:"name,omitempty"`
		Precision            map[string]float64                `json:"precision,omitempty"`
		Preview              *bool                             `json:"preview,omitempty"`
		SliValue             datadog.NullableFloat64           `json:"sli_value,omitempty"`
		SpanPrecision        *float64                          `json:"span_precision,omitempty"`
		Uptime               datadog.NullableFloat64           `json:"uptime,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error_budget_remaining", "errors", "group", "history", "monitor_modified", "monitor_type", "name", "precision", "preview", "sli_value", "span_precision", "uptime"})
	} else {
		return err
	}
	o.ErrorBudgetRemaining = all.ErrorBudgetRemaining
	o.Errors = all.Errors
	o.Group = all.Group
	o.History = all.History
	o.MonitorModified = all.MonitorModified
	o.MonitorType = all.MonitorType
	o.Name = all.Name
	o.Precision = all.Precision
	o.Preview = all.Preview
	o.SliValue = all.SliValue
	o.SpanPrecision = all.SpanPrecision
	o.Uptime = all.Uptime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
