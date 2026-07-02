// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationAnalyticsAggregate Analytics aggregation parameters.
type LLMObsExperimentationAnalyticsAggregate struct {
	// List of metric computations to perform.
	Compute []LLMObsExperimentationAnalyticsCompute `json:"compute"`
	// Filter to a specific dataset version.
	DatasetVersion datadog.NullableInt64 `json:"dataset_version,omitempty"`
	// Fields to group results by.
	GroupBy []LLMObsExperimentationAnalyticsGroupBy `json:"group_by,omitempty"`
	// Data indexes to query. At least one is required.
	Indexes []string `json:"indexes"`
	// Maximum number of results to return.
	Limit datadog.NullableInt32 `json:"limit,omitempty"`
	// Search query for filtering analytics data.
	Search LLMObsExperimentationAnalyticsSearch `json:"search"`
	// Unix-millisecond time range for filtering analytics data.
	Time *LLMObsExperimentationAnalyticsTimeRange `json:"time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationAnalyticsAggregate instantiates a new LLMObsExperimentationAnalyticsAggregate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationAnalyticsAggregate(compute []LLMObsExperimentationAnalyticsCompute, indexes []string, search LLMObsExperimentationAnalyticsSearch) *LLMObsExperimentationAnalyticsAggregate {
	this := LLMObsExperimentationAnalyticsAggregate{}
	this.Compute = compute
	this.Indexes = indexes
	this.Search = search
	return &this
}

// NewLLMObsExperimentationAnalyticsAggregateWithDefaults instantiates a new LLMObsExperimentationAnalyticsAggregate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationAnalyticsAggregateWithDefaults() *LLMObsExperimentationAnalyticsAggregate {
	this := LLMObsExperimentationAnalyticsAggregate{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *LLMObsExperimentationAnalyticsAggregate) GetCompute() []LLMObsExperimentationAnalyticsCompute {
	if o == nil {
		var ret []LLMObsExperimentationAnalyticsCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsAggregate) GetComputeOk() (*[]LLMObsExperimentationAnalyticsCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *LLMObsExperimentationAnalyticsAggregate) SetCompute(v []LLMObsExperimentationAnalyticsCompute) {
	o.Compute = v
}

// GetDatasetVersion returns the DatasetVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentationAnalyticsAggregate) GetDatasetVersion() int64 {
	if o == nil || o.DatasetVersion.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DatasetVersion.Get()
}

// GetDatasetVersionOk returns a tuple with the DatasetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentationAnalyticsAggregate) GetDatasetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatasetVersion.Get(), o.DatasetVersion.IsSet()
}

// HasDatasetVersion returns a boolean if a field has been set.
func (o *LLMObsExperimentationAnalyticsAggregate) HasDatasetVersion() bool {
	return o != nil && o.DatasetVersion.IsSet()
}

// SetDatasetVersion gets a reference to the given datadog.NullableInt64 and assigns it to the DatasetVersion field.
func (o *LLMObsExperimentationAnalyticsAggregate) SetDatasetVersion(v int64) {
	o.DatasetVersion.Set(&v)
}

// SetDatasetVersionNil sets the value for DatasetVersion to be an explicit nil.
func (o *LLMObsExperimentationAnalyticsAggregate) SetDatasetVersionNil() {
	o.DatasetVersion.Set(nil)
}

// UnsetDatasetVersion ensures that no value is present for DatasetVersion, not even an explicit nil.
func (o *LLMObsExperimentationAnalyticsAggregate) UnsetDatasetVersion() {
	o.DatasetVersion.Unset()
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *LLMObsExperimentationAnalyticsAggregate) GetGroupBy() []LLMObsExperimentationAnalyticsGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []LLMObsExperimentationAnalyticsGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsAggregate) GetGroupByOk() (*[]LLMObsExperimentationAnalyticsGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *LLMObsExperimentationAnalyticsAggregate) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []LLMObsExperimentationAnalyticsGroupBy and assigns it to the GroupBy field.
func (o *LLMObsExperimentationAnalyticsAggregate) SetGroupBy(v []LLMObsExperimentationAnalyticsGroupBy) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value.
func (o *LLMObsExperimentationAnalyticsAggregate) GetIndexes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsAggregate) GetIndexesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// SetIndexes sets field value.
func (o *LLMObsExperimentationAnalyticsAggregate) SetIndexes(v []string) {
	o.Indexes = v
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentationAnalyticsAggregate) GetLimit() int32 {
	if o == nil || o.Limit.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentationAnalyticsAggregate) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *LLMObsExperimentationAnalyticsAggregate) HasLimit() bool {
	return o != nil && o.Limit.IsSet()
}

// SetLimit gets a reference to the given datadog.NullableInt32 and assigns it to the Limit field.
func (o *LLMObsExperimentationAnalyticsAggregate) SetLimit(v int32) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil.
func (o *LLMObsExperimentationAnalyticsAggregate) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil.
func (o *LLMObsExperimentationAnalyticsAggregate) UnsetLimit() {
	o.Limit.Unset()
}

// GetSearch returns the Search field value.
func (o *LLMObsExperimentationAnalyticsAggregate) GetSearch() LLMObsExperimentationAnalyticsSearch {
	if o == nil {
		var ret LLMObsExperimentationAnalyticsSearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsAggregate) GetSearchOk() (*LLMObsExperimentationAnalyticsSearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *LLMObsExperimentationAnalyticsAggregate) SetSearch(v LLMObsExperimentationAnalyticsSearch) {
	o.Search = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *LLMObsExperimentationAnalyticsAggregate) GetTime() LLMObsExperimentationAnalyticsTimeRange {
	if o == nil || o.Time == nil {
		var ret LLMObsExperimentationAnalyticsTimeRange
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsAggregate) GetTimeOk() (*LLMObsExperimentationAnalyticsTimeRange, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *LLMObsExperimentationAnalyticsAggregate) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given LLMObsExperimentationAnalyticsTimeRange and assigns it to the Time field.
func (o *LLMObsExperimentationAnalyticsAggregate) SetTime(v LLMObsExperimentationAnalyticsTimeRange) {
	o.Time = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationAnalyticsAggregate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	if o.DatasetVersion.IsSet() {
		toSerialize["dataset_version"] = o.DatasetVersion.Get()
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["indexes"] = o.Indexes
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	toSerialize["search"] = o.Search
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationAnalyticsAggregate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute        *[]LLMObsExperimentationAnalyticsCompute `json:"compute"`
		DatasetVersion datadog.NullableInt64                    `json:"dataset_version,omitempty"`
		GroupBy        []LLMObsExperimentationAnalyticsGroupBy  `json:"group_by,omitempty"`
		Indexes        *[]string                                `json:"indexes"`
		Limit          datadog.NullableInt32                    `json:"limit,omitempty"`
		Search         *LLMObsExperimentationAnalyticsSearch    `json:"search"`
		Time           *LLMObsExperimentationAnalyticsTimeRange `json:"time,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.Indexes == nil {
		return fmt.Errorf("required field indexes missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "dataset_version", "group_by", "indexes", "limit", "search", "time"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Compute = *all.Compute
	o.DatasetVersion = all.DatasetVersion
	o.GroupBy = all.GroupBy
	o.Indexes = *all.Indexes
	o.Limit = all.Limit
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search
	if all.Time != nil && all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Time = all.Time

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
