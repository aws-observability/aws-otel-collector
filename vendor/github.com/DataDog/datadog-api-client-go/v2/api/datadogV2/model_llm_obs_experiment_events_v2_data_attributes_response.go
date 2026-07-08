// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentEventsV2DataAttributesResponse Attributes of an experiment events response.
type LLMObsExperimentEventsV2DataAttributesResponse struct {
	// Experiment spans, each enriched with their associated evaluation metrics.
	Spans []LLMObsExperimentSpanWithEvals `json:"spans"`
	// Experiment-level summary evaluation metrics (not tied to individual spans).
	SummaryMetrics []LLMObsExperimentEvalMetricEvent `json:"summary_metrics"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentEventsV2DataAttributesResponse instantiates a new LLMObsExperimentEventsV2DataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentEventsV2DataAttributesResponse(spans []LLMObsExperimentSpanWithEvals, summaryMetrics []LLMObsExperimentEvalMetricEvent) *LLMObsExperimentEventsV2DataAttributesResponse {
	this := LLMObsExperimentEventsV2DataAttributesResponse{}
	this.Spans = spans
	this.SummaryMetrics = summaryMetrics
	return &this
}

// NewLLMObsExperimentEventsV2DataAttributesResponseWithDefaults instantiates a new LLMObsExperimentEventsV2DataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentEventsV2DataAttributesResponseWithDefaults() *LLMObsExperimentEventsV2DataAttributesResponse {
	this := LLMObsExperimentEventsV2DataAttributesResponse{}
	return &this
}

// GetSpans returns the Spans field value.
func (o *LLMObsExperimentEventsV2DataAttributesResponse) GetSpans() []LLMObsExperimentSpanWithEvals {
	if o == nil {
		var ret []LLMObsExperimentSpanWithEvals
		return ret
	}
	return o.Spans
}

// GetSpansOk returns a tuple with the Spans field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEventsV2DataAttributesResponse) GetSpansOk() (*[]LLMObsExperimentSpanWithEvals, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spans, true
}

// SetSpans sets field value.
func (o *LLMObsExperimentEventsV2DataAttributesResponse) SetSpans(v []LLMObsExperimentSpanWithEvals) {
	o.Spans = v
}

// GetSummaryMetrics returns the SummaryMetrics field value.
func (o *LLMObsExperimentEventsV2DataAttributesResponse) GetSummaryMetrics() []LLMObsExperimentEvalMetricEvent {
	if o == nil {
		var ret []LLMObsExperimentEvalMetricEvent
		return ret
	}
	return o.SummaryMetrics
}

// GetSummaryMetricsOk returns a tuple with the SummaryMetrics field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEventsV2DataAttributesResponse) GetSummaryMetricsOk() (*[]LLMObsExperimentEvalMetricEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SummaryMetrics, true
}

// SetSummaryMetrics sets field value.
func (o *LLMObsExperimentEventsV2DataAttributesResponse) SetSummaryMetrics(v []LLMObsExperimentEvalMetricEvent) {
	o.SummaryMetrics = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentEventsV2DataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["spans"] = o.Spans
	toSerialize["summary_metrics"] = o.SummaryMetrics

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentEventsV2DataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Spans          *[]LLMObsExperimentSpanWithEvals   `json:"spans"`
		SummaryMetrics *[]LLMObsExperimentEvalMetricEvent `json:"summary_metrics"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Spans == nil {
		return fmt.Errorf("required field spans missing")
	}
	if all.SummaryMetrics == nil {
		return fmt.Errorf("required field summary_metrics missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"spans", "summary_metrics"})
	} else {
		return err
	}
	o.Spans = *all.Spans
	o.SummaryMetrics = *all.SummaryMetrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
