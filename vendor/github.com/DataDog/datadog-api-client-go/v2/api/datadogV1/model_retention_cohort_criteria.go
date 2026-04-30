// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionCohortCriteria Cohort criteria for retention queries.
type RetentionCohortCriteria struct {
	// Product Analytics event query.
	BaseQuery ProductAnalyticsBaseQuery `json:"base_query"`
	// Time interval for cohort criteria.
	TimeInterval RetentionCohortCriteriaTimeInterval `json:"time_interval"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionCohortCriteria instantiates a new RetentionCohortCriteria object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionCohortCriteria(baseQuery ProductAnalyticsBaseQuery, timeInterval RetentionCohortCriteriaTimeInterval) *RetentionCohortCriteria {
	this := RetentionCohortCriteria{}
	this.BaseQuery = baseQuery
	this.TimeInterval = timeInterval
	return &this
}

// NewRetentionCohortCriteriaWithDefaults instantiates a new RetentionCohortCriteria object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionCohortCriteriaWithDefaults() *RetentionCohortCriteria {
	this := RetentionCohortCriteria{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *RetentionCohortCriteria) GetBaseQuery() ProductAnalyticsBaseQuery {
	if o == nil {
		var ret ProductAnalyticsBaseQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *RetentionCohortCriteria) GetBaseQueryOk() (*ProductAnalyticsBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *RetentionCohortCriteria) SetBaseQuery(v ProductAnalyticsBaseQuery) {
	o.BaseQuery = v
}

// GetTimeInterval returns the TimeInterval field value.
func (o *RetentionCohortCriteria) GetTimeInterval() RetentionCohortCriteriaTimeInterval {
	if o == nil {
		var ret RetentionCohortCriteriaTimeInterval
		return ret
	}
	return o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value
// and a boolean to check if the value has been set.
func (o *RetentionCohortCriteria) GetTimeIntervalOk() (*RetentionCohortCriteriaTimeInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeInterval, true
}

// SetTimeInterval sets field value.
func (o *RetentionCohortCriteria) SetTimeInterval(v RetentionCohortCriteriaTimeInterval) {
	o.TimeInterval = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionCohortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	toSerialize["time_interval"] = o.TimeInterval
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionCohortCriteria) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery    *ProductAnalyticsBaseQuery           `json:"base_query"`
		TimeInterval *RetentionCohortCriteriaTimeInterval `json:"time_interval"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
	}
	if all.TimeInterval == nil {
		return fmt.Errorf("required field time_interval missing")
	}

	hasInvalidField := false
	if all.BaseQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BaseQuery = *all.BaseQuery
	if all.TimeInterval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeInterval = *all.TimeInterval

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
