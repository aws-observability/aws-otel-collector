// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOTimeSliceCondition The time-slice condition, composed of 3 parts: 1. the metric timeseries query, 2. the comparator,
// and 3. the threshold. Optionally, a fourth part, the query interval, can be provided.
type SLOTimeSliceCondition struct {
	// The comparator used to compare the SLI value to the threshold.
	Comparator SLOTimeSliceComparator `json:"comparator"`
	// The queries and formula used to calculate the SLI value.
	Query SLOTimeSliceQuery `json:"query"`
	// The interval used when querying data, which defines the size of a time slice.
	// Two values are allowed: 60 (1 minute) and 300 (5 minutes).
	// If not provided, the value defaults to 300 (5 minutes).
	QueryIntervalSeconds *SLOTimeSliceInterval `json:"query_interval_seconds,omitempty"`
	// The threshold value to which each SLI value will be compared.
	Threshold float64 `json:"threshold"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSLOTimeSliceCondition instantiates a new SLOTimeSliceCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOTimeSliceCondition(comparator SLOTimeSliceComparator, query SLOTimeSliceQuery, threshold float64) *SLOTimeSliceCondition {
	this := SLOTimeSliceCondition{}
	this.Comparator = comparator
	this.Query = query
	this.Threshold = threshold
	return &this
}

// NewSLOTimeSliceConditionWithDefaults instantiates a new SLOTimeSliceCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOTimeSliceConditionWithDefaults() *SLOTimeSliceCondition {
	this := SLOTimeSliceCondition{}
	return &this
}

// GetComparator returns the Comparator field value.
func (o *SLOTimeSliceCondition) GetComparator() SLOTimeSliceComparator {
	if o == nil {
		var ret SLOTimeSliceComparator
		return ret
	}
	return o.Comparator
}

// GetComparatorOk returns a tuple with the Comparator field value
// and a boolean to check if the value has been set.
func (o *SLOTimeSliceCondition) GetComparatorOk() (*SLOTimeSliceComparator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comparator, true
}

// SetComparator sets field value.
func (o *SLOTimeSliceCondition) SetComparator(v SLOTimeSliceComparator) {
	o.Comparator = v
}

// GetQuery returns the Query field value.
func (o *SLOTimeSliceCondition) GetQuery() SLOTimeSliceQuery {
	if o == nil {
		var ret SLOTimeSliceQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SLOTimeSliceCondition) GetQueryOk() (*SLOTimeSliceQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *SLOTimeSliceCondition) SetQuery(v SLOTimeSliceQuery) {
	o.Query = v
}

// GetQueryIntervalSeconds returns the QueryIntervalSeconds field value if set, zero value otherwise.
func (o *SLOTimeSliceCondition) GetQueryIntervalSeconds() SLOTimeSliceInterval {
	if o == nil || o.QueryIntervalSeconds == nil {
		var ret SLOTimeSliceInterval
		return ret
	}
	return *o.QueryIntervalSeconds
}

// GetQueryIntervalSecondsOk returns a tuple with the QueryIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOTimeSliceCondition) GetQueryIntervalSecondsOk() (*SLOTimeSliceInterval, bool) {
	if o == nil || o.QueryIntervalSeconds == nil {
		return nil, false
	}
	return o.QueryIntervalSeconds, true
}

// HasQueryIntervalSeconds returns a boolean if a field has been set.
func (o *SLOTimeSliceCondition) HasQueryIntervalSeconds() bool {
	return o != nil && o.QueryIntervalSeconds != nil
}

// SetQueryIntervalSeconds gets a reference to the given SLOTimeSliceInterval and assigns it to the QueryIntervalSeconds field.
func (o *SLOTimeSliceCondition) SetQueryIntervalSeconds(v SLOTimeSliceInterval) {
	o.QueryIntervalSeconds = &v
}

// GetThreshold returns the Threshold field value.
func (o *SLOTimeSliceCondition) GetThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *SLOTimeSliceCondition) GetThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value.
func (o *SLOTimeSliceCondition) SetThreshold(v float64) {
	o.Threshold = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOTimeSliceCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["comparator"] = o.Comparator
	toSerialize["query"] = o.Query
	if o.QueryIntervalSeconds != nil {
		toSerialize["query_interval_seconds"] = o.QueryIntervalSeconds
	}
	toSerialize["threshold"] = o.Threshold

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOTimeSliceCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Comparator           *SLOTimeSliceComparator `json:"comparator"`
		Query                *SLOTimeSliceQuery      `json:"query"`
		QueryIntervalSeconds *SLOTimeSliceInterval   `json:"query_interval_seconds,omitempty"`
		Threshold            *float64                `json:"threshold"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Comparator == nil {
		return fmt.Errorf("required field comparator missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Threshold == nil {
		return fmt.Errorf("required field threshold missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"comparator", "query", "query_interval_seconds", "threshold"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Comparator.IsValid() {
		hasInvalidField = true
	} else {
		o.Comparator = *all.Comparator
	}
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	if all.QueryIntervalSeconds != nil && !all.QueryIntervalSeconds.IsValid() {
		hasInvalidField = true
	} else {
		o.QueryIntervalSeconds = all.QueryIntervalSeconds
	}
	o.Threshold = *all.Threshold

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
