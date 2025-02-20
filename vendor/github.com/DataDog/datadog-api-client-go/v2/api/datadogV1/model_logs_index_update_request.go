// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsIndexUpdateRequest Object for updating a Datadog Log index.
type LogsIndexUpdateRequest struct {
	// The number of log events you can send in this index per day before you are rate-limited.
	DailyLimit *int64 `json:"daily_limit,omitempty"`
	// Object containing options to override the default daily limit reset time.
	DailyLimitReset *LogsDailyLimitReset `json:"daily_limit_reset,omitempty"`
	// A percentage threshold of the daily quota at which a Datadog warning event is generated.
	DailyLimitWarningThresholdPercentage *float64 `json:"daily_limit_warning_threshold_percentage,omitempty"`
	// If true, sets the `daily_limit` value to null and the index is not limited on a daily basis (any
	// specified `daily_limit` value in the request is ignored). If false or omitted, the index's current
	// `daily_limit` is maintained.
	DisableDailyLimit *bool `json:"disable_daily_limit,omitempty"`
	// An array of exclusion objects. The logs are tested against the query of each filter,
	// following the order of the array. Only the first matching active exclusion matters,
	// others (if any) are ignored.
	ExclusionFilters []LogsExclusion `json:"exclusion_filters,omitempty"`
	// Filter for logs.
	Filter LogsFilter `json:"filter"`
	// The total number of days logs are stored in Standard and Flex Tier before being deleted from the index.
	// If Standard Tier is enabled on this index, logs are first retained in Standard Tier for the number of days specified through `num_retention_days`,
	// and then stored in Flex Tier until the number of days specified in `num_flex_logs_retention_days` is reached.
	// The available values depend on retention plans specified in your organization's contract/subscriptions.
	//
	// **Note**: Changing this value affects all logs already in this index. It may also affect billing.
	NumFlexLogsRetentionDays *int64 `json:"num_flex_logs_retention_days,omitempty"`
	// The number of days logs are stored in Standard Tier before aging into the Flex Tier or being deleted from the index.
	// The available values depend on retention plans specified in your organization's contract/subscriptions.
	//
	// **Note**: Changing this value affects all logs already in this index. It may also affect billing.
	NumRetentionDays *int64 `json:"num_retention_days,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsIndexUpdateRequest instantiates a new LogsIndexUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsIndexUpdateRequest(filter LogsFilter) *LogsIndexUpdateRequest {
	this := LogsIndexUpdateRequest{}
	this.Filter = filter
	return &this
}

// NewLogsIndexUpdateRequestWithDefaults instantiates a new LogsIndexUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsIndexUpdateRequestWithDefaults() *LogsIndexUpdateRequest {
	this := LogsIndexUpdateRequest{}
	return &this
}

// GetDailyLimit returns the DailyLimit field value if set, zero value otherwise.
func (o *LogsIndexUpdateRequest) GetDailyLimit() int64 {
	if o == nil || o.DailyLimit == nil {
		var ret int64
		return ret
	}
	return *o.DailyLimit
}

// GetDailyLimitOk returns a tuple with the DailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndexUpdateRequest) GetDailyLimitOk() (*int64, bool) {
	if o == nil || o.DailyLimit == nil {
		return nil, false
	}
	return o.DailyLimit, true
}

// HasDailyLimit returns a boolean if a field has been set.
func (o *LogsIndexUpdateRequest) HasDailyLimit() bool {
	return o != nil && o.DailyLimit != nil
}

// SetDailyLimit gets a reference to the given int64 and assigns it to the DailyLimit field.
func (o *LogsIndexUpdateRequest) SetDailyLimit(v int64) {
	o.DailyLimit = &v
}

// GetDailyLimitReset returns the DailyLimitReset field value if set, zero value otherwise.
func (o *LogsIndexUpdateRequest) GetDailyLimitReset() LogsDailyLimitReset {
	if o == nil || o.DailyLimitReset == nil {
		var ret LogsDailyLimitReset
		return ret
	}
	return *o.DailyLimitReset
}

// GetDailyLimitResetOk returns a tuple with the DailyLimitReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndexUpdateRequest) GetDailyLimitResetOk() (*LogsDailyLimitReset, bool) {
	if o == nil || o.DailyLimitReset == nil {
		return nil, false
	}
	return o.DailyLimitReset, true
}

// HasDailyLimitReset returns a boolean if a field has been set.
func (o *LogsIndexUpdateRequest) HasDailyLimitReset() bool {
	return o != nil && o.DailyLimitReset != nil
}

// SetDailyLimitReset gets a reference to the given LogsDailyLimitReset and assigns it to the DailyLimitReset field.
func (o *LogsIndexUpdateRequest) SetDailyLimitReset(v LogsDailyLimitReset) {
	o.DailyLimitReset = &v
}

// GetDailyLimitWarningThresholdPercentage returns the DailyLimitWarningThresholdPercentage field value if set, zero value otherwise.
func (o *LogsIndexUpdateRequest) GetDailyLimitWarningThresholdPercentage() float64 {
	if o == nil || o.DailyLimitWarningThresholdPercentage == nil {
		var ret float64
		return ret
	}
	return *o.DailyLimitWarningThresholdPercentage
}

// GetDailyLimitWarningThresholdPercentageOk returns a tuple with the DailyLimitWarningThresholdPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndexUpdateRequest) GetDailyLimitWarningThresholdPercentageOk() (*float64, bool) {
	if o == nil || o.DailyLimitWarningThresholdPercentage == nil {
		return nil, false
	}
	return o.DailyLimitWarningThresholdPercentage, true
}

// HasDailyLimitWarningThresholdPercentage returns a boolean if a field has been set.
func (o *LogsIndexUpdateRequest) HasDailyLimitWarningThresholdPercentage() bool {
	return o != nil && o.DailyLimitWarningThresholdPercentage != nil
}

// SetDailyLimitWarningThresholdPercentage gets a reference to the given float64 and assigns it to the DailyLimitWarningThresholdPercentage field.
func (o *LogsIndexUpdateRequest) SetDailyLimitWarningThresholdPercentage(v float64) {
	o.DailyLimitWarningThresholdPercentage = &v
}

// GetDisableDailyLimit returns the DisableDailyLimit field value if set, zero value otherwise.
func (o *LogsIndexUpdateRequest) GetDisableDailyLimit() bool {
	if o == nil || o.DisableDailyLimit == nil {
		var ret bool
		return ret
	}
	return *o.DisableDailyLimit
}

// GetDisableDailyLimitOk returns a tuple with the DisableDailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndexUpdateRequest) GetDisableDailyLimitOk() (*bool, bool) {
	if o == nil || o.DisableDailyLimit == nil {
		return nil, false
	}
	return o.DisableDailyLimit, true
}

// HasDisableDailyLimit returns a boolean if a field has been set.
func (o *LogsIndexUpdateRequest) HasDisableDailyLimit() bool {
	return o != nil && o.DisableDailyLimit != nil
}

// SetDisableDailyLimit gets a reference to the given bool and assigns it to the DisableDailyLimit field.
func (o *LogsIndexUpdateRequest) SetDisableDailyLimit(v bool) {
	o.DisableDailyLimit = &v
}

// GetExclusionFilters returns the ExclusionFilters field value if set, zero value otherwise.
func (o *LogsIndexUpdateRequest) GetExclusionFilters() []LogsExclusion {
	if o == nil || o.ExclusionFilters == nil {
		var ret []LogsExclusion
		return ret
	}
	return o.ExclusionFilters
}

// GetExclusionFiltersOk returns a tuple with the ExclusionFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndexUpdateRequest) GetExclusionFiltersOk() (*[]LogsExclusion, bool) {
	if o == nil || o.ExclusionFilters == nil {
		return nil, false
	}
	return &o.ExclusionFilters, true
}

// HasExclusionFilters returns a boolean if a field has been set.
func (o *LogsIndexUpdateRequest) HasExclusionFilters() bool {
	return o != nil && o.ExclusionFilters != nil
}

// SetExclusionFilters gets a reference to the given []LogsExclusion and assigns it to the ExclusionFilters field.
func (o *LogsIndexUpdateRequest) SetExclusionFilters(v []LogsExclusion) {
	o.ExclusionFilters = v
}

// GetFilter returns the Filter field value.
func (o *LogsIndexUpdateRequest) GetFilter() LogsFilter {
	if o == nil {
		var ret LogsFilter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *LogsIndexUpdateRequest) GetFilterOk() (*LogsFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *LogsIndexUpdateRequest) SetFilter(v LogsFilter) {
	o.Filter = v
}

// GetNumFlexLogsRetentionDays returns the NumFlexLogsRetentionDays field value if set, zero value otherwise.
func (o *LogsIndexUpdateRequest) GetNumFlexLogsRetentionDays() int64 {
	if o == nil || o.NumFlexLogsRetentionDays == nil {
		var ret int64
		return ret
	}
	return *o.NumFlexLogsRetentionDays
}

// GetNumFlexLogsRetentionDaysOk returns a tuple with the NumFlexLogsRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndexUpdateRequest) GetNumFlexLogsRetentionDaysOk() (*int64, bool) {
	if o == nil || o.NumFlexLogsRetentionDays == nil {
		return nil, false
	}
	return o.NumFlexLogsRetentionDays, true
}

// HasNumFlexLogsRetentionDays returns a boolean if a field has been set.
func (o *LogsIndexUpdateRequest) HasNumFlexLogsRetentionDays() bool {
	return o != nil && o.NumFlexLogsRetentionDays != nil
}

// SetNumFlexLogsRetentionDays gets a reference to the given int64 and assigns it to the NumFlexLogsRetentionDays field.
func (o *LogsIndexUpdateRequest) SetNumFlexLogsRetentionDays(v int64) {
	o.NumFlexLogsRetentionDays = &v
}

// GetNumRetentionDays returns the NumRetentionDays field value if set, zero value otherwise.
func (o *LogsIndexUpdateRequest) GetNumRetentionDays() int64 {
	if o == nil || o.NumRetentionDays == nil {
		var ret int64
		return ret
	}
	return *o.NumRetentionDays
}

// GetNumRetentionDaysOk returns a tuple with the NumRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndexUpdateRequest) GetNumRetentionDaysOk() (*int64, bool) {
	if o == nil || o.NumRetentionDays == nil {
		return nil, false
	}
	return o.NumRetentionDays, true
}

// HasNumRetentionDays returns a boolean if a field has been set.
func (o *LogsIndexUpdateRequest) HasNumRetentionDays() bool {
	return o != nil && o.NumRetentionDays != nil
}

// SetNumRetentionDays gets a reference to the given int64 and assigns it to the NumRetentionDays field.
func (o *LogsIndexUpdateRequest) SetNumRetentionDays(v int64) {
	o.NumRetentionDays = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsIndexUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DailyLimit != nil {
		toSerialize["daily_limit"] = o.DailyLimit
	}
	if o.DailyLimitReset != nil {
		toSerialize["daily_limit_reset"] = o.DailyLimitReset
	}
	if o.DailyLimitWarningThresholdPercentage != nil {
		toSerialize["daily_limit_warning_threshold_percentage"] = o.DailyLimitWarningThresholdPercentage
	}
	if o.DisableDailyLimit != nil {
		toSerialize["disable_daily_limit"] = o.DisableDailyLimit
	}
	if o.ExclusionFilters != nil {
		toSerialize["exclusion_filters"] = o.ExclusionFilters
	}
	toSerialize["filter"] = o.Filter
	if o.NumFlexLogsRetentionDays != nil {
		toSerialize["num_flex_logs_retention_days"] = o.NumFlexLogsRetentionDays
	}
	if o.NumRetentionDays != nil {
		toSerialize["num_retention_days"] = o.NumRetentionDays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsIndexUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DailyLimit                           *int64               `json:"daily_limit,omitempty"`
		DailyLimitReset                      *LogsDailyLimitReset `json:"daily_limit_reset,omitempty"`
		DailyLimitWarningThresholdPercentage *float64             `json:"daily_limit_warning_threshold_percentage,omitempty"`
		DisableDailyLimit                    *bool                `json:"disable_daily_limit,omitempty"`
		ExclusionFilters                     []LogsExclusion      `json:"exclusion_filters,omitempty"`
		Filter                               *LogsFilter          `json:"filter"`
		NumFlexLogsRetentionDays             *int64               `json:"num_flex_logs_retention_days,omitempty"`
		NumRetentionDays                     *int64               `json:"num_retention_days,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"daily_limit", "daily_limit_reset", "daily_limit_warning_threshold_percentage", "disable_daily_limit", "exclusion_filters", "filter", "num_flex_logs_retention_days", "num_retention_days"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DailyLimit = all.DailyLimit
	if all.DailyLimitReset != nil && all.DailyLimitReset.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DailyLimitReset = all.DailyLimitReset
	o.DailyLimitWarningThresholdPercentage = all.DailyLimitWarningThresholdPercentage
	o.DisableDailyLimit = all.DisableDailyLimit
	o.ExclusionFilters = all.ExclusionFilters
	if all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = *all.Filter
	o.NumFlexLogsRetentionDays = all.NumFlexLogsRetentionDays
	o.NumRetentionDays = all.NumRetentionDays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
