// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionFilterAllAttributes The attributes of the retention filter.
type RetentionFilterAllAttributes struct {
	// The creation timestamp of the retention filter.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The creator of the retention filter.
	CreatedBy *string `json:"created_by,omitempty"`
	// Shows whether the filter can be edited.
	Editable *bool `json:"editable,omitempty"`
	// The status of the retention filter (Enabled/Disabled).
	Enabled *bool `json:"enabled,omitempty"`
	// The execution order of the retention filter.
	ExecutionOrder *int64 `json:"execution_order,omitempty"`
	// The spans filter used to index spans.
	Filter *SpansFilter `json:"filter,omitempty"`
	// The type of retention filter.
	FilterType *RetentionFilterAllType `json:"filter_type,omitempty"`
	// The modification timestamp of the retention filter.
	ModifiedAt *int64 `json:"modified_at,omitempty"`
	// The modifier of the retention filter.
	ModifiedBy *string `json:"modified_by,omitempty"`
	// The name of the retention filter.
	Name *string `json:"name,omitempty"`
	// Sample rate to apply to spans going through this retention filter.
	// A value of 1.0 keeps all spans matching the query.
	Rate *float64 `json:"rate,omitempty"`
	// Sample rate to apply to traces containing spans going through this retention filter.
	// A value of 1.0 keeps all traces with spans matching the query.
	TraceRate *float64 `json:"trace_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRetentionFilterAllAttributes instantiates a new RetentionFilterAllAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionFilterAllAttributes() *RetentionFilterAllAttributes {
	this := RetentionFilterAllAttributes{}
	var filterType RetentionFilterAllType = RETENTIONFILTERALLTYPE_SPANS_SAMPLING_PROCESSOR
	this.FilterType = &filterType
	return &this
}

// NewRetentionFilterAllAttributesWithDefaults instantiates a new RetentionFilterAllAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionFilterAllAttributesWithDefaults() *RetentionFilterAllAttributes {
	this := RetentionFilterAllAttributes{}
	var filterType RetentionFilterAllType = RETENTIONFILTERALLTYPE_SPANS_SAMPLING_PROCESSOR
	this.FilterType = &filterType
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *RetentionFilterAllAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *RetentionFilterAllAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasEditable() bool {
	return o != nil && o.Editable != nil
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *RetentionFilterAllAttributes) SetEditable(v bool) {
	o.Editable = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RetentionFilterAllAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExecutionOrder returns the ExecutionOrder field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetExecutionOrder() int64 {
	if o == nil || o.ExecutionOrder == nil {
		var ret int64
		return ret
	}
	return *o.ExecutionOrder
}

// GetExecutionOrderOk returns a tuple with the ExecutionOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetExecutionOrderOk() (*int64, bool) {
	if o == nil || o.ExecutionOrder == nil {
		return nil, false
	}
	return o.ExecutionOrder, true
}

// HasExecutionOrder returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasExecutionOrder() bool {
	return o != nil && o.ExecutionOrder != nil
}

// SetExecutionOrder gets a reference to the given int64 and assigns it to the ExecutionOrder field.
func (o *RetentionFilterAllAttributes) SetExecutionOrder(v int64) {
	o.ExecutionOrder = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetFilter() SpansFilter {
	if o == nil || o.Filter == nil {
		var ret SpansFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetFilterOk() (*SpansFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given SpansFilter and assigns it to the Filter field.
func (o *RetentionFilterAllAttributes) SetFilter(v SpansFilter) {
	o.Filter = &v
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetFilterType() RetentionFilterAllType {
	if o == nil || o.FilterType == nil {
		var ret RetentionFilterAllType
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetFilterTypeOk() (*RetentionFilterAllType, bool) {
	if o == nil || o.FilterType == nil {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasFilterType() bool {
	return o != nil && o.FilterType != nil
}

// SetFilterType gets a reference to the given RetentionFilterAllType and assigns it to the FilterType field.
func (o *RetentionFilterAllAttributes) SetFilterType(v RetentionFilterAllType) {
	o.FilterType = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetModifiedAt() int64 {
	if o == nil || o.ModifiedAt == nil {
		var ret int64
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetModifiedAtOk() (*int64, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given int64 and assigns it to the ModifiedAt field.
func (o *RetentionFilterAllAttributes) SetModifiedAt(v int64) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy != nil
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *RetentionFilterAllAttributes) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RetentionFilterAllAttributes) SetName(v string) {
	o.Name = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetRate() float64 {
	if o == nil || o.Rate == nil {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetRateOk() (*float64, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasRate() bool {
	return o != nil && o.Rate != nil
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *RetentionFilterAllAttributes) SetRate(v float64) {
	o.Rate = &v
}

// GetTraceRate returns the TraceRate field value if set, zero value otherwise.
func (o *RetentionFilterAllAttributes) GetTraceRate() float64 {
	if o == nil || o.TraceRate == nil {
		var ret float64
		return ret
	}
	return *o.TraceRate
}

// GetTraceRateOk returns a tuple with the TraceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterAllAttributes) GetTraceRateOk() (*float64, bool) {
	if o == nil || o.TraceRate == nil {
		return nil, false
	}
	return o.TraceRate, true
}

// HasTraceRate returns a boolean if a field has been set.
func (o *RetentionFilterAllAttributes) HasTraceRate() bool {
	return o != nil && o.TraceRate != nil
}

// SetTraceRate gets a reference to the given float64 and assigns it to the TraceRate field.
func (o *RetentionFilterAllAttributes) SetTraceRate(v float64) {
	o.TraceRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionFilterAllAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Editable != nil {
		toSerialize["editable"] = o.Editable
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExecutionOrder != nil {
		toSerialize["execution_order"] = o.ExecutionOrder
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.FilterType != nil {
		toSerialize["filter_type"] = o.FilterType
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.ModifiedBy != nil {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.TraceRate != nil {
		toSerialize["trace_rate"] = o.TraceRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionFilterAllAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *int64                  `json:"created_at,omitempty"`
		CreatedBy      *string                 `json:"created_by,omitempty"`
		Editable       *bool                   `json:"editable,omitempty"`
		Enabled        *bool                   `json:"enabled,omitempty"`
		ExecutionOrder *int64                  `json:"execution_order,omitempty"`
		Filter         *SpansFilter            `json:"filter,omitempty"`
		FilterType     *RetentionFilterAllType `json:"filter_type,omitempty"`
		ModifiedAt     *int64                  `json:"modified_at,omitempty"`
		ModifiedBy     *string                 `json:"modified_by,omitempty"`
		Name           *string                 `json:"name,omitempty"`
		Rate           *float64                `json:"rate,omitempty"`
		TraceRate      *float64                `json:"trace_rate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "editable", "enabled", "execution_order", "filter", "filter_type", "modified_at", "modified_by", "name", "rate", "trace_rate"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Editable = all.Editable
	o.Enabled = all.Enabled
	o.ExecutionOrder = all.ExecutionOrder
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	if all.FilterType != nil && !all.FilterType.IsValid() {
		hasInvalidField = true
	} else {
		o.FilterType = all.FilterType
	}
	o.ModifiedAt = all.ModifiedAt
	o.ModifiedBy = all.ModifiedBy
	o.Name = all.Name
	o.Rate = all.Rate
	o.TraceRate = all.TraceRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
