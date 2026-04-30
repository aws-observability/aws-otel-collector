// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionReturnCriteria Return criteria for retention queries.
type RetentionReturnCriteria struct {
	// Product Analytics event query.
	BaseQuery ProductAnalyticsBaseQuery `json:"base_query"`
	// Time interval for return criteria.
	TimeInterval *RetentionReturnCriteriaTimeInterval `json:"time_interval,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionReturnCriteria instantiates a new RetentionReturnCriteria object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionReturnCriteria(baseQuery ProductAnalyticsBaseQuery) *RetentionReturnCriteria {
	this := RetentionReturnCriteria{}
	this.BaseQuery = baseQuery
	return &this
}

// NewRetentionReturnCriteriaWithDefaults instantiates a new RetentionReturnCriteria object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionReturnCriteriaWithDefaults() *RetentionReturnCriteria {
	this := RetentionReturnCriteria{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *RetentionReturnCriteria) GetBaseQuery() ProductAnalyticsBaseQuery {
	if o == nil {
		var ret ProductAnalyticsBaseQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *RetentionReturnCriteria) GetBaseQueryOk() (*ProductAnalyticsBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *RetentionReturnCriteria) SetBaseQuery(v ProductAnalyticsBaseQuery) {
	o.BaseQuery = v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *RetentionReturnCriteria) GetTimeInterval() RetentionReturnCriteriaTimeInterval {
	if o == nil || o.TimeInterval == nil {
		var ret RetentionReturnCriteriaTimeInterval
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionReturnCriteria) GetTimeIntervalOk() (*RetentionReturnCriteriaTimeInterval, bool) {
	if o == nil || o.TimeInterval == nil {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *RetentionReturnCriteria) HasTimeInterval() bool {
	return o != nil && o.TimeInterval != nil
}

// SetTimeInterval gets a reference to the given RetentionReturnCriteriaTimeInterval and assigns it to the TimeInterval field.
func (o *RetentionReturnCriteria) SetTimeInterval(v RetentionReturnCriteriaTimeInterval) {
	o.TimeInterval = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionReturnCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	if o.TimeInterval != nil {
		toSerialize["time_interval"] = o.TimeInterval
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionReturnCriteria) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery    *ProductAnalyticsBaseQuery           `json:"base_query"`
		TimeInterval *RetentionReturnCriteriaTimeInterval `json:"time_interval,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
	}

	hasInvalidField := false
	if all.BaseQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BaseQuery = *all.BaseQuery
	if all.TimeInterval != nil && all.TimeInterval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeInterval = all.TimeInterval

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
