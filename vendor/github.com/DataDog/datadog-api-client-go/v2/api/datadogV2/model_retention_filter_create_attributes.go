// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionFilterCreateAttributes The object describing the configuration of the retention filter to create/update.
type RetentionFilterCreateAttributes struct {
	// Enable/Disable the retention filter.
	Enabled bool `json:"enabled"`
	// The spans filter. Spans matching this filter will be indexed and stored.
	Filter SpansFilterCreate `json:"filter"`
	// The type of retention filter. The value should always be spans-sampling-processor.
	FilterType RetentionFilterType `json:"filter_type"`
	// The name of the retention filter.
	Name string `json:"name"`
	// Sample rate to apply to spans going through this retention filter.
	// A value of 1.0 keeps all spans matching the query.
	Rate float64 `json:"rate"`
	// Sample rate to apply to traces containing spans going through this retention filter.
	// A value of 1.0 keeps all traces with spans matching the query.
	TraceRate *float64 `json:"trace_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRetentionFilterCreateAttributes instantiates a new RetentionFilterCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionFilterCreateAttributes(enabled bool, filter SpansFilterCreate, filterType RetentionFilterType, name string, rate float64) *RetentionFilterCreateAttributes {
	this := RetentionFilterCreateAttributes{}
	this.Enabled = enabled
	this.Filter = filter
	this.FilterType = filterType
	this.Name = name
	this.Rate = rate
	return &this
}

// NewRetentionFilterCreateAttributesWithDefaults instantiates a new RetentionFilterCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionFilterCreateAttributesWithDefaults() *RetentionFilterCreateAttributes {
	this := RetentionFilterCreateAttributes{}
	var filterType RetentionFilterType = RETENTIONFILTERTYPE_SPANS_SAMPLING_PROCESSOR
	this.FilterType = filterType
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *RetentionFilterCreateAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RetentionFilterCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *RetentionFilterCreateAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFilter returns the Filter field value.
func (o *RetentionFilterCreateAttributes) GetFilter() SpansFilterCreate {
	if o == nil {
		var ret SpansFilterCreate
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *RetentionFilterCreateAttributes) GetFilterOk() (*SpansFilterCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *RetentionFilterCreateAttributes) SetFilter(v SpansFilterCreate) {
	o.Filter = v
}

// GetFilterType returns the FilterType field value.
func (o *RetentionFilterCreateAttributes) GetFilterType() RetentionFilterType {
	if o == nil {
		var ret RetentionFilterType
		return ret
	}
	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *RetentionFilterCreateAttributes) GetFilterTypeOk() (*RetentionFilterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value.
func (o *RetentionFilterCreateAttributes) SetFilterType(v RetentionFilterType) {
	o.FilterType = v
}

// GetName returns the Name field value.
func (o *RetentionFilterCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RetentionFilterCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RetentionFilterCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetRate returns the Rate field value.
func (o *RetentionFilterCreateAttributes) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *RetentionFilterCreateAttributes) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value.
func (o *RetentionFilterCreateAttributes) SetRate(v float64) {
	o.Rate = v
}

// GetTraceRate returns the TraceRate field value if set, zero value otherwise.
func (o *RetentionFilterCreateAttributes) GetTraceRate() float64 {
	if o == nil || o.TraceRate == nil {
		var ret float64
		return ret
	}
	return *o.TraceRate
}

// GetTraceRateOk returns a tuple with the TraceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionFilterCreateAttributes) GetTraceRateOk() (*float64, bool) {
	if o == nil || o.TraceRate == nil {
		return nil, false
	}
	return o.TraceRate, true
}

// HasTraceRate returns a boolean if a field has been set.
func (o *RetentionFilterCreateAttributes) HasTraceRate() bool {
	return o != nil && o.TraceRate != nil
}

// SetTraceRate gets a reference to the given float64 and assigns it to the TraceRate field.
func (o *RetentionFilterCreateAttributes) SetTraceRate(v float64) {
	o.TraceRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionFilterCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["filter"] = o.Filter
	toSerialize["filter_type"] = o.FilterType
	toSerialize["name"] = o.Name
	toSerialize["rate"] = o.Rate
	if o.TraceRate != nil {
		toSerialize["trace_rate"] = o.TraceRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionFilterCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled    *bool                `json:"enabled"`
		Filter     *SpansFilterCreate   `json:"filter"`
		FilterType *RetentionFilterType `json:"filter_type"`
		Name       *string              `json:"name"`
		Rate       *float64             `json:"rate"`
		TraceRate  *float64             `json:"trace_rate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
	}
	if all.FilterType == nil {
		return fmt.Errorf("required field filter_type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Rate == nil {
		return fmt.Errorf("required field rate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "filter", "filter_type", "name", "rate", "trace_rate"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = *all.Enabled
	if all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = *all.Filter
	if !all.FilterType.IsValid() {
		hasInvalidField = true
	} else {
		o.FilterType = *all.FilterType
	}
	o.Name = *all.Name
	o.Rate = *all.Rate
	o.TraceRate = all.TraceRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
