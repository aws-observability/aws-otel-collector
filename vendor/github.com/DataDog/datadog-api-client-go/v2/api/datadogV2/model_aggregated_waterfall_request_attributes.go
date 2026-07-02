// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedWaterfallRequestAttributes Attributes for an aggregated waterfall query.
type AggregatedWaterfallRequestAttributes struct {
	// The RUM application ID to analyze.
	ApplicationId string `json:"application_id"`
	// Performance criteria to filter view instances by a metric threshold.
	Criteria *AggregatedWaterfallPerformanceCriteria `json:"criteria,omitempty"`
	// RUM query string to filter events (for example, @session.type:user @geo.country:US).
	Filter *string `json:"filter,omitempty"`
	// Start of the time range as a Unix timestamp in seconds.
	From int64 `json:"from"`
	// When true, enriches each resource with cross-view appearance statistics.
	IncludeGlobalAppearance *bool `json:"include_global_appearance,omitempty"`
	// Number of view instances to sample, between 1 and 500.
	SampleSize int32 `json:"sample_size"`
	// End of the time range as a Unix timestamp in seconds.
	To int64 `json:"to"`
	// The RUM view name to analyze (for example, /account/login).
	ViewName string `json:"view_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedWaterfallRequestAttributes instantiates a new AggregatedWaterfallRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedWaterfallRequestAttributes(applicationId string, from int64, sampleSize int32, to int64, viewName string) *AggregatedWaterfallRequestAttributes {
	this := AggregatedWaterfallRequestAttributes{}
	this.ApplicationId = applicationId
	this.From = from
	this.SampleSize = sampleSize
	this.To = to
	this.ViewName = viewName
	return &this
}

// NewAggregatedWaterfallRequestAttributesWithDefaults instantiates a new AggregatedWaterfallRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedWaterfallRequestAttributesWithDefaults() *AggregatedWaterfallRequestAttributes {
	this := AggregatedWaterfallRequestAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *AggregatedWaterfallRequestAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallRequestAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *AggregatedWaterfallRequestAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *AggregatedWaterfallRequestAttributes) GetCriteria() AggregatedWaterfallPerformanceCriteria {
	if o == nil || o.Criteria == nil {
		var ret AggregatedWaterfallPerformanceCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallRequestAttributes) GetCriteriaOk() (*AggregatedWaterfallPerformanceCriteria, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *AggregatedWaterfallRequestAttributes) HasCriteria() bool {
	return o != nil && o.Criteria != nil
}

// SetCriteria gets a reference to the given AggregatedWaterfallPerformanceCriteria and assigns it to the Criteria field.
func (o *AggregatedWaterfallRequestAttributes) SetCriteria(v AggregatedWaterfallPerformanceCriteria) {
	o.Criteria = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AggregatedWaterfallRequestAttributes) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallRequestAttributes) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AggregatedWaterfallRequestAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *AggregatedWaterfallRequestAttributes) SetFilter(v string) {
	o.Filter = &v
}

// GetFrom returns the From field value.
func (o *AggregatedWaterfallRequestAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallRequestAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *AggregatedWaterfallRequestAttributes) SetFrom(v int64) {
	o.From = v
}

// GetIncludeGlobalAppearance returns the IncludeGlobalAppearance field value if set, zero value otherwise.
func (o *AggregatedWaterfallRequestAttributes) GetIncludeGlobalAppearance() bool {
	if o == nil || o.IncludeGlobalAppearance == nil {
		var ret bool
		return ret
	}
	return *o.IncludeGlobalAppearance
}

// GetIncludeGlobalAppearanceOk returns a tuple with the IncludeGlobalAppearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallRequestAttributes) GetIncludeGlobalAppearanceOk() (*bool, bool) {
	if o == nil || o.IncludeGlobalAppearance == nil {
		return nil, false
	}
	return o.IncludeGlobalAppearance, true
}

// HasIncludeGlobalAppearance returns a boolean if a field has been set.
func (o *AggregatedWaterfallRequestAttributes) HasIncludeGlobalAppearance() bool {
	return o != nil && o.IncludeGlobalAppearance != nil
}

// SetIncludeGlobalAppearance gets a reference to the given bool and assigns it to the IncludeGlobalAppearance field.
func (o *AggregatedWaterfallRequestAttributes) SetIncludeGlobalAppearance(v bool) {
	o.IncludeGlobalAppearance = &v
}

// GetSampleSize returns the SampleSize field value.
func (o *AggregatedWaterfallRequestAttributes) GetSampleSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.SampleSize
}

// GetSampleSizeOk returns a tuple with the SampleSize field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallRequestAttributes) GetSampleSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampleSize, true
}

// SetSampleSize sets field value.
func (o *AggregatedWaterfallRequestAttributes) SetSampleSize(v int32) {
	o.SampleSize = v
}

// GetTo returns the To field value.
func (o *AggregatedWaterfallRequestAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallRequestAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *AggregatedWaterfallRequestAttributes) SetTo(v int64) {
	o.To = v
}

// GetViewName returns the ViewName field value.
func (o *AggregatedWaterfallRequestAttributes) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallRequestAttributes) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value.
func (o *AggregatedWaterfallRequestAttributes) SetViewName(v string) {
	o.ViewName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedWaterfallRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["from"] = o.From
	if o.IncludeGlobalAppearance != nil {
		toSerialize["include_global_appearance"] = o.IncludeGlobalAppearance
	}
	toSerialize["sample_size"] = o.SampleSize
	toSerialize["to"] = o.To
	toSerialize["view_name"] = o.ViewName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedWaterfallRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId           *string                                 `json:"application_id"`
		Criteria                *AggregatedWaterfallPerformanceCriteria `json:"criteria,omitempty"`
		Filter                  *string                                 `json:"filter,omitempty"`
		From                    *int64                                  `json:"from"`
		IncludeGlobalAppearance *bool                                   `json:"include_global_appearance,omitempty"`
		SampleSize              *int32                                  `json:"sample_size"`
		To                      *int64                                  `json:"to"`
		ViewName                *string                                 `json:"view_name"`
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
	if all.SampleSize == nil {
		return fmt.Errorf("required field sample_size missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	if all.ViewName == nil {
		return fmt.Errorf("required field view_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "criteria", "filter", "from", "include_global_appearance", "sample_size", "to", "view_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = *all.ApplicationId
	if all.Criteria != nil && all.Criteria.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Criteria = all.Criteria
	o.Filter = all.Filter
	o.From = *all.From
	o.IncludeGlobalAppearance = all.IncludeGlobalAppearance
	o.SampleSize = *all.SampleSize
	o.To = *all.To
	o.ViewName = *all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
