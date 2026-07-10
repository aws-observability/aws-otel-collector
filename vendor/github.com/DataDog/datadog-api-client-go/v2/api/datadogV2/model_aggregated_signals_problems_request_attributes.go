// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedSignalsProblemsRequestAttributes Attributes for an aggregated signals and problems query.
type AggregatedSignalsProblemsRequestAttributes struct {
	// The RUM application ID to analyze.
	ApplicationId string `json:"application_id"`
	// Performance criteria to filter view instances by a metric threshold.
	Criteria *AggregatedWaterfallPerformanceCriteria `json:"criteria,omitempty"`
	// List of detection types to include in the response. When omitted, all types are returned.
	DetectionTypes []string `json:"detection_types,omitempty"`
	// RUM query string to filter events (for example, @session.type:user @geo.country:US).
	Filter *string `json:"filter,omitempty"`
	// Start of the time range as a Unix timestamp in seconds.
	From int64 `json:"from"`
	// Number of view instances to sample, between 1 and 50.
	SampleSize int32 `json:"sample_size"`
	// End of the time range as a Unix timestamp in seconds.
	To int64 `json:"to"`
	// The RUM view name to analyze (for example, /account/login).
	ViewName string `json:"view_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedSignalsProblemsRequestAttributes instantiates a new AggregatedSignalsProblemsRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedSignalsProblemsRequestAttributes(applicationId string, from int64, sampleSize int32, to int64, viewName string) *AggregatedSignalsProblemsRequestAttributes {
	this := AggregatedSignalsProblemsRequestAttributes{}
	this.ApplicationId = applicationId
	this.From = from
	this.SampleSize = sampleSize
	this.To = to
	this.ViewName = viewName
	return &this
}

// NewAggregatedSignalsProblemsRequestAttributesWithDefaults instantiates a new AggregatedSignalsProblemsRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedSignalsProblemsRequestAttributesWithDefaults() *AggregatedSignalsProblemsRequestAttributes {
	this := AggregatedSignalsProblemsRequestAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *AggregatedSignalsProblemsRequestAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *AggregatedSignalsProblemsRequestAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *AggregatedSignalsProblemsRequestAttributes) GetCriteria() AggregatedWaterfallPerformanceCriteria {
	if o == nil || o.Criteria == nil {
		var ret AggregatedWaterfallPerformanceCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) GetCriteriaOk() (*AggregatedWaterfallPerformanceCriteria, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) HasCriteria() bool {
	return o != nil && o.Criteria != nil
}

// SetCriteria gets a reference to the given AggregatedWaterfallPerformanceCriteria and assigns it to the Criteria field.
func (o *AggregatedSignalsProblemsRequestAttributes) SetCriteria(v AggregatedWaterfallPerformanceCriteria) {
	o.Criteria = &v
}

// GetDetectionTypes returns the DetectionTypes field value if set, zero value otherwise.
func (o *AggregatedSignalsProblemsRequestAttributes) GetDetectionTypes() []string {
	if o == nil || o.DetectionTypes == nil {
		var ret []string
		return ret
	}
	return o.DetectionTypes
}

// GetDetectionTypesOk returns a tuple with the DetectionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) GetDetectionTypesOk() (*[]string, bool) {
	if o == nil || o.DetectionTypes == nil {
		return nil, false
	}
	return &o.DetectionTypes, true
}

// HasDetectionTypes returns a boolean if a field has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) HasDetectionTypes() bool {
	return o != nil && o.DetectionTypes != nil
}

// SetDetectionTypes gets a reference to the given []string and assigns it to the DetectionTypes field.
func (o *AggregatedSignalsProblemsRequestAttributes) SetDetectionTypes(v []string) {
	o.DetectionTypes = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AggregatedSignalsProblemsRequestAttributes) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *AggregatedSignalsProblemsRequestAttributes) SetFilter(v string) {
	o.Filter = &v
}

// GetFrom returns the From field value.
func (o *AggregatedSignalsProblemsRequestAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *AggregatedSignalsProblemsRequestAttributes) SetFrom(v int64) {
	o.From = v
}

// GetSampleSize returns the SampleSize field value.
func (o *AggregatedSignalsProblemsRequestAttributes) GetSampleSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.SampleSize
}

// GetSampleSizeOk returns a tuple with the SampleSize field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) GetSampleSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampleSize, true
}

// SetSampleSize sets field value.
func (o *AggregatedSignalsProblemsRequestAttributes) SetSampleSize(v int32) {
	o.SampleSize = v
}

// GetTo returns the To field value.
func (o *AggregatedSignalsProblemsRequestAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *AggregatedSignalsProblemsRequestAttributes) SetTo(v int64) {
	o.To = v
}

// GetViewName returns the ViewName field value.
func (o *AggregatedSignalsProblemsRequestAttributes) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsRequestAttributes) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value.
func (o *AggregatedSignalsProblemsRequestAttributes) SetViewName(v string) {
	o.ViewName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedSignalsProblemsRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	if o.DetectionTypes != nil {
		toSerialize["detection_types"] = o.DetectionTypes
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["from"] = o.From
	toSerialize["sample_size"] = o.SampleSize
	toSerialize["to"] = o.To
	toSerialize["view_name"] = o.ViewName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedSignalsProblemsRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId  *string                                 `json:"application_id"`
		Criteria       *AggregatedWaterfallPerformanceCriteria `json:"criteria,omitempty"`
		DetectionTypes []string                                `json:"detection_types,omitempty"`
		Filter         *string                                 `json:"filter,omitempty"`
		From           *int64                                  `json:"from"`
		SampleSize     *int32                                  `json:"sample_size"`
		To             *int64                                  `json:"to"`
		ViewName       *string                                 `json:"view_name"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "criteria", "detection_types", "filter", "from", "sample_size", "to", "view_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = *all.ApplicationId
	if all.Criteria != nil && all.Criteria.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Criteria = all.Criteria
	o.DetectionTypes = all.DetectionTypes
	o.Filter = all.Filter
	o.From = *all.From
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
