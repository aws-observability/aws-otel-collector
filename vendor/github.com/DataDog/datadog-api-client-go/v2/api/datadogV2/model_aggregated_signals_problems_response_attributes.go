// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedSignalsProblemsResponseAttributes Attributes of an aggregated signals and problems response.
type AggregatedSignalsProblemsResponseAttributes struct {
	// The RUM application ID that was analyzed.
	ApplicationId string `json:"application_id"`
	// Performance criteria to filter view instances by a metric threshold.
	Criteria *AggregatedWaterfallPerformanceCriteria `json:"criteria,omitempty"`
	// Start of the analyzed time range as a Unix timestamp in seconds.
	From int64 `json:"from"`
	// Grouped detection results by detection type.
	ProblemDetections SignalsProblemsDetections `json:"problem_detections"`
	// Metadata about the sampling quality for a signals and problems query.
	SampleMetadata SignalsProblemsSampleMetadata `json:"sample_metadata"`
	// End of the analyzed time range as a Unix timestamp in seconds.
	To int64 `json:"to"`
	// The RUM view name that was analyzed.
	ViewName string `json:"view_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedSignalsProblemsResponseAttributes instantiates a new AggregatedSignalsProblemsResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedSignalsProblemsResponseAttributes(applicationId string, from int64, problemDetections SignalsProblemsDetections, sampleMetadata SignalsProblemsSampleMetadata, to int64, viewName string) *AggregatedSignalsProblemsResponseAttributes {
	this := AggregatedSignalsProblemsResponseAttributes{}
	this.ApplicationId = applicationId
	this.From = from
	this.ProblemDetections = problemDetections
	this.SampleMetadata = sampleMetadata
	this.To = to
	this.ViewName = viewName
	return &this
}

// NewAggregatedSignalsProblemsResponseAttributesWithDefaults instantiates a new AggregatedSignalsProblemsResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedSignalsProblemsResponseAttributesWithDefaults() *AggregatedSignalsProblemsResponseAttributes {
	this := AggregatedSignalsProblemsResponseAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *AggregatedSignalsProblemsResponseAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsResponseAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *AggregatedSignalsProblemsResponseAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *AggregatedSignalsProblemsResponseAttributes) GetCriteria() AggregatedWaterfallPerformanceCriteria {
	if o == nil || o.Criteria == nil {
		var ret AggregatedWaterfallPerformanceCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsResponseAttributes) GetCriteriaOk() (*AggregatedWaterfallPerformanceCriteria, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *AggregatedSignalsProblemsResponseAttributes) HasCriteria() bool {
	return o != nil && o.Criteria != nil
}

// SetCriteria gets a reference to the given AggregatedWaterfallPerformanceCriteria and assigns it to the Criteria field.
func (o *AggregatedSignalsProblemsResponseAttributes) SetCriteria(v AggregatedWaterfallPerformanceCriteria) {
	o.Criteria = &v
}

// GetFrom returns the From field value.
func (o *AggregatedSignalsProblemsResponseAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsResponseAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *AggregatedSignalsProblemsResponseAttributes) SetFrom(v int64) {
	o.From = v
}

// GetProblemDetections returns the ProblemDetections field value.
func (o *AggregatedSignalsProblemsResponseAttributes) GetProblemDetections() SignalsProblemsDetections {
	if o == nil {
		var ret SignalsProblemsDetections
		return ret
	}
	return o.ProblemDetections
}

// GetProblemDetectionsOk returns a tuple with the ProblemDetections field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsResponseAttributes) GetProblemDetectionsOk() (*SignalsProblemsDetections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProblemDetections, true
}

// SetProblemDetections sets field value.
func (o *AggregatedSignalsProblemsResponseAttributes) SetProblemDetections(v SignalsProblemsDetections) {
	o.ProblemDetections = v
}

// GetSampleMetadata returns the SampleMetadata field value.
func (o *AggregatedSignalsProblemsResponseAttributes) GetSampleMetadata() SignalsProblemsSampleMetadata {
	if o == nil {
		var ret SignalsProblemsSampleMetadata
		return ret
	}
	return o.SampleMetadata
}

// GetSampleMetadataOk returns a tuple with the SampleMetadata field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsResponseAttributes) GetSampleMetadataOk() (*SignalsProblemsSampleMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampleMetadata, true
}

// SetSampleMetadata sets field value.
func (o *AggregatedSignalsProblemsResponseAttributes) SetSampleMetadata(v SignalsProblemsSampleMetadata) {
	o.SampleMetadata = v
}

// GetTo returns the To field value.
func (o *AggregatedSignalsProblemsResponseAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsResponseAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *AggregatedSignalsProblemsResponseAttributes) SetTo(v int64) {
	o.To = v
}

// GetViewName returns the ViewName field value.
func (o *AggregatedSignalsProblemsResponseAttributes) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *AggregatedSignalsProblemsResponseAttributes) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value.
func (o *AggregatedSignalsProblemsResponseAttributes) SetViewName(v string) {
	o.ViewName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedSignalsProblemsResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	toSerialize["from"] = o.From
	toSerialize["problem_detections"] = o.ProblemDetections
	toSerialize["sample_metadata"] = o.SampleMetadata
	toSerialize["to"] = o.To
	toSerialize["view_name"] = o.ViewName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedSignalsProblemsResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId     *string                                 `json:"application_id"`
		Criteria          *AggregatedWaterfallPerformanceCriteria `json:"criteria,omitempty"`
		From              *int64                                  `json:"from"`
		ProblemDetections *SignalsProblemsDetections              `json:"problem_detections"`
		SampleMetadata    *SignalsProblemsSampleMetadata          `json:"sample_metadata"`
		To                *int64                                  `json:"to"`
		ViewName          *string                                 `json:"view_name"`
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
	if all.ProblemDetections == nil {
		return fmt.Errorf("required field problem_detections missing")
	}
	if all.SampleMetadata == nil {
		return fmt.Errorf("required field sample_metadata missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	if all.ViewName == nil {
		return fmt.Errorf("required field view_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "criteria", "from", "problem_detections", "sample_metadata", "to", "view_name"})
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
	if all.ProblemDetections.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProblemDetections = *all.ProblemDetections
	if all.SampleMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SampleMetadata = *all.SampleMetadata
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
