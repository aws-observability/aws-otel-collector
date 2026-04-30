// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentEventsDataAttributesRequest Attributes for pushing experiment events including spans and metrics.
type LLMObsExperimentEventsDataAttributesRequest struct {
	// List of metrics to push for the experiment.
	Metrics []LLMObsExperimentMetric `json:"metrics,omitempty"`
	// List of spans to push for the experiment.
	Spans []LLMObsExperimentSpan `json:"spans,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentEventsDataAttributesRequest instantiates a new LLMObsExperimentEventsDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentEventsDataAttributesRequest() *LLMObsExperimentEventsDataAttributesRequest {
	this := LLMObsExperimentEventsDataAttributesRequest{}
	return &this
}

// NewLLMObsExperimentEventsDataAttributesRequestWithDefaults instantiates a new LLMObsExperimentEventsDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentEventsDataAttributesRequestWithDefaults() *LLMObsExperimentEventsDataAttributesRequest {
	this := LLMObsExperimentEventsDataAttributesRequest{}
	return &this
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *LLMObsExperimentEventsDataAttributesRequest) GetMetrics() []LLMObsExperimentMetric {
	if o == nil || o.Metrics == nil {
		var ret []LLMObsExperimentMetric
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEventsDataAttributesRequest) GetMetricsOk() (*[]LLMObsExperimentMetric, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *LLMObsExperimentEventsDataAttributesRequest) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []LLMObsExperimentMetric and assigns it to the Metrics field.
func (o *LLMObsExperimentEventsDataAttributesRequest) SetMetrics(v []LLMObsExperimentMetric) {
	o.Metrics = v
}

// GetSpans returns the Spans field value if set, zero value otherwise.
func (o *LLMObsExperimentEventsDataAttributesRequest) GetSpans() []LLMObsExperimentSpan {
	if o == nil || o.Spans == nil {
		var ret []LLMObsExperimentSpan
		return ret
	}
	return o.Spans
}

// GetSpansOk returns a tuple with the Spans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEventsDataAttributesRequest) GetSpansOk() (*[]LLMObsExperimentSpan, bool) {
	if o == nil || o.Spans == nil {
		return nil, false
	}
	return &o.Spans, true
}

// HasSpans returns a boolean if a field has been set.
func (o *LLMObsExperimentEventsDataAttributesRequest) HasSpans() bool {
	return o != nil && o.Spans != nil
}

// SetSpans gets a reference to the given []LLMObsExperimentSpan and assigns it to the Spans field.
func (o *LLMObsExperimentEventsDataAttributesRequest) SetSpans(v []LLMObsExperimentSpan) {
	o.Spans = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentEventsDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Spans != nil {
		toSerialize["spans"] = o.Spans
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentEventsDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metrics []LLMObsExperimentMetric `json:"metrics,omitempty"`
		Spans   []LLMObsExperimentSpan   `json:"spans,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metrics", "spans"})
	} else {
		return err
	}
	o.Metrics = all.Metrics
	o.Spans = all.Spans

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
