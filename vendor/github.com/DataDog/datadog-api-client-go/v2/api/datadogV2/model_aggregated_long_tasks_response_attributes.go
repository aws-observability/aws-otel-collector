// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedLongTasksResponseAttributes Attributes of an aggregated long tasks response.
type AggregatedLongTasksResponseAttributes struct {
	// The RUM application ID that was analyzed.
	ApplicationId string `json:"application_id"`
	// Performance criteria to filter view instances by a metric threshold.
	Criteria *AggregatedWaterfallPerformanceCriteria `json:"criteria,omitempty"`
	// Start of the analyzed time range as a Unix timestamp in seconds.
	From int64 `json:"from"`
	// Long task statistics grouped by invoker type, sorted by impact score descending.
	LongTasksByInvokerType []AggregatedLongTasksByInvokerType `json:"long_tasks_by_invoker_type"`
	// List of RUM view IDs sampled for this aggregation, capped at 50.
	SampledViewIds []string `json:"sampled_view_ids"`
	// End of the analyzed time range as a Unix timestamp in seconds.
	To int64 `json:"to"`
	// Number of view instances included in the analysis.
	ViewCount int32 `json:"view_count"`
	// The RUM view name that was analyzed.
	ViewName string `json:"view_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedLongTasksResponseAttributes instantiates a new AggregatedLongTasksResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedLongTasksResponseAttributes(applicationId string, from int64, longTasksByInvokerType []AggregatedLongTasksByInvokerType, sampledViewIds []string, to int64, viewCount int32, viewName string) *AggregatedLongTasksResponseAttributes {
	this := AggregatedLongTasksResponseAttributes{}
	this.ApplicationId = applicationId
	this.From = from
	this.LongTasksByInvokerType = longTasksByInvokerType
	this.SampledViewIds = sampledViewIds
	this.To = to
	this.ViewCount = viewCount
	this.ViewName = viewName
	return &this
}

// NewAggregatedLongTasksResponseAttributesWithDefaults instantiates a new AggregatedLongTasksResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedLongTasksResponseAttributesWithDefaults() *AggregatedLongTasksResponseAttributes {
	this := AggregatedLongTasksResponseAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *AggregatedLongTasksResponseAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksResponseAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *AggregatedLongTasksResponseAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *AggregatedLongTasksResponseAttributes) GetCriteria() AggregatedWaterfallPerformanceCriteria {
	if o == nil || o.Criteria == nil {
		var ret AggregatedWaterfallPerformanceCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksResponseAttributes) GetCriteriaOk() (*AggregatedWaterfallPerformanceCriteria, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *AggregatedLongTasksResponseAttributes) HasCriteria() bool {
	return o != nil && o.Criteria != nil
}

// SetCriteria gets a reference to the given AggregatedWaterfallPerformanceCriteria and assigns it to the Criteria field.
func (o *AggregatedLongTasksResponseAttributes) SetCriteria(v AggregatedWaterfallPerformanceCriteria) {
	o.Criteria = &v
}

// GetFrom returns the From field value.
func (o *AggregatedLongTasksResponseAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksResponseAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *AggregatedLongTasksResponseAttributes) SetFrom(v int64) {
	o.From = v
}

// GetLongTasksByInvokerType returns the LongTasksByInvokerType field value.
func (o *AggregatedLongTasksResponseAttributes) GetLongTasksByInvokerType() []AggregatedLongTasksByInvokerType {
	if o == nil {
		var ret []AggregatedLongTasksByInvokerType
		return ret
	}
	return o.LongTasksByInvokerType
}

// GetLongTasksByInvokerTypeOk returns a tuple with the LongTasksByInvokerType field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksResponseAttributes) GetLongTasksByInvokerTypeOk() (*[]AggregatedLongTasksByInvokerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LongTasksByInvokerType, true
}

// SetLongTasksByInvokerType sets field value.
func (o *AggregatedLongTasksResponseAttributes) SetLongTasksByInvokerType(v []AggregatedLongTasksByInvokerType) {
	o.LongTasksByInvokerType = v
}

// GetSampledViewIds returns the SampledViewIds field value.
func (o *AggregatedLongTasksResponseAttributes) GetSampledViewIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SampledViewIds
}

// GetSampledViewIdsOk returns a tuple with the SampledViewIds field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksResponseAttributes) GetSampledViewIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampledViewIds, true
}

// SetSampledViewIds sets field value.
func (o *AggregatedLongTasksResponseAttributes) SetSampledViewIds(v []string) {
	o.SampledViewIds = v
}

// GetTo returns the To field value.
func (o *AggregatedLongTasksResponseAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksResponseAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *AggregatedLongTasksResponseAttributes) SetTo(v int64) {
	o.To = v
}

// GetViewCount returns the ViewCount field value.
func (o *AggregatedLongTasksResponseAttributes) GetViewCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewCount
}

// GetViewCountOk returns a tuple with the ViewCount field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksResponseAttributes) GetViewCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewCount, true
}

// SetViewCount sets field value.
func (o *AggregatedLongTasksResponseAttributes) SetViewCount(v int32) {
	o.ViewCount = v
}

// GetViewName returns the ViewName field value.
func (o *AggregatedLongTasksResponseAttributes) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *AggregatedLongTasksResponseAttributes) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value.
func (o *AggregatedLongTasksResponseAttributes) SetViewName(v string) {
	o.ViewName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedLongTasksResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	toSerialize["from"] = o.From
	toSerialize["long_tasks_by_invoker_type"] = o.LongTasksByInvokerType
	toSerialize["sampled_view_ids"] = o.SampledViewIds
	toSerialize["to"] = o.To
	toSerialize["view_count"] = o.ViewCount
	toSerialize["view_name"] = o.ViewName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedLongTasksResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId          *string                                 `json:"application_id"`
		Criteria               *AggregatedWaterfallPerformanceCriteria `json:"criteria,omitempty"`
		From                   *int64                                  `json:"from"`
		LongTasksByInvokerType *[]AggregatedLongTasksByInvokerType     `json:"long_tasks_by_invoker_type"`
		SampledViewIds         *[]string                               `json:"sampled_view_ids"`
		To                     *int64                                  `json:"to"`
		ViewCount              *int32                                  `json:"view_count"`
		ViewName               *string                                 `json:"view_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.LongTasksByInvokerType == nil {
		return fmt.Errorf("required field long_tasks_by_invoker_type missing")
	}
	if all.SampledViewIds == nil {
		return fmt.Errorf("required field sampled_view_ids missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	if all.ViewCount == nil {
		return fmt.Errorf("required field view_count missing")
	}
	if all.ViewName == nil {
		return fmt.Errorf("required field view_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "criteria", "from", "long_tasks_by_invoker_type", "sampled_view_ids", "to", "view_count", "view_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = *all.ApplicationId
	if all.Criteria != nil && all.Criteria.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Criteria = all.Criteria
	o.From = *all.From
	o.LongTasksByInvokerType = *all.LongTasksByInvokerType
	o.SampledViewIds = *all.SampledViewIds
	o.To = *all.To
	o.ViewCount = *all.ViewCount
	o.ViewName = *all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
