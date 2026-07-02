// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsOpenAIMetadata OpenAI-specific metadata for an inference request.
type LLMObsOpenAIMetadata struct {
	// The reasoning effort level for OpenAI models that support it.
	ReasoningEffort NullableLLMObsOpenAIReasoningEffort `json:"reasoning_effort,omitempty"`
	// The verbosity of the reasoning summary.
	ReasoningSummary NullableLLMObsOpenAIReasoningSummary `json:"reasoning_summary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsOpenAIMetadata instantiates a new LLMObsOpenAIMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsOpenAIMetadata() *LLMObsOpenAIMetadata {
	this := LLMObsOpenAIMetadata{}
	return &this
}

// NewLLMObsOpenAIMetadataWithDefaults instantiates a new LLMObsOpenAIMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsOpenAIMetadataWithDefaults() *LLMObsOpenAIMetadata {
	this := LLMObsOpenAIMetadata{}
	return &this
}

// GetReasoningEffort returns the ReasoningEffort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsOpenAIMetadata) GetReasoningEffort() LLMObsOpenAIReasoningEffort {
	if o == nil || o.ReasoningEffort.Get() == nil {
		var ret LLMObsOpenAIReasoningEffort
		return ret
	}
	return *o.ReasoningEffort.Get()
}

// GetReasoningEffortOk returns a tuple with the ReasoningEffort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsOpenAIMetadata) GetReasoningEffortOk() (*LLMObsOpenAIReasoningEffort, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasoningEffort.Get(), o.ReasoningEffort.IsSet()
}

// HasReasoningEffort returns a boolean if a field has been set.
func (o *LLMObsOpenAIMetadata) HasReasoningEffort() bool {
	return o != nil && o.ReasoningEffort.IsSet()
}

// SetReasoningEffort gets a reference to the given NullableLLMObsOpenAIReasoningEffort and assigns it to the ReasoningEffort field.
func (o *LLMObsOpenAIMetadata) SetReasoningEffort(v LLMObsOpenAIReasoningEffort) {
	o.ReasoningEffort.Set(&v)
}

// SetReasoningEffortNil sets the value for ReasoningEffort to be an explicit nil.
func (o *LLMObsOpenAIMetadata) SetReasoningEffortNil() {
	o.ReasoningEffort.Set(nil)
}

// UnsetReasoningEffort ensures that no value is present for ReasoningEffort, not even an explicit nil.
func (o *LLMObsOpenAIMetadata) UnsetReasoningEffort() {
	o.ReasoningEffort.Unset()
}

// GetReasoningSummary returns the ReasoningSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsOpenAIMetadata) GetReasoningSummary() LLMObsOpenAIReasoningSummary {
	if o == nil || o.ReasoningSummary.Get() == nil {
		var ret LLMObsOpenAIReasoningSummary
		return ret
	}
	return *o.ReasoningSummary.Get()
}

// GetReasoningSummaryOk returns a tuple with the ReasoningSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsOpenAIMetadata) GetReasoningSummaryOk() (*LLMObsOpenAIReasoningSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasoningSummary.Get(), o.ReasoningSummary.IsSet()
}

// HasReasoningSummary returns a boolean if a field has been set.
func (o *LLMObsOpenAIMetadata) HasReasoningSummary() bool {
	return o != nil && o.ReasoningSummary.IsSet()
}

// SetReasoningSummary gets a reference to the given NullableLLMObsOpenAIReasoningSummary and assigns it to the ReasoningSummary field.
func (o *LLMObsOpenAIMetadata) SetReasoningSummary(v LLMObsOpenAIReasoningSummary) {
	o.ReasoningSummary.Set(&v)
}

// SetReasoningSummaryNil sets the value for ReasoningSummary to be an explicit nil.
func (o *LLMObsOpenAIMetadata) SetReasoningSummaryNil() {
	o.ReasoningSummary.Set(nil)
}

// UnsetReasoningSummary ensures that no value is present for ReasoningSummary, not even an explicit nil.
func (o *LLMObsOpenAIMetadata) UnsetReasoningSummary() {
	o.ReasoningSummary.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsOpenAIMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ReasoningEffort.IsSet() {
		toSerialize["reasoning_effort"] = o.ReasoningEffort.Get()
	}
	if o.ReasoningSummary.IsSet() {
		toSerialize["reasoning_summary"] = o.ReasoningSummary.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsOpenAIMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReasoningEffort  NullableLLMObsOpenAIReasoningEffort  `json:"reasoning_effort,omitempty"`
		ReasoningSummary NullableLLMObsOpenAIReasoningSummary `json:"reasoning_summary,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reasoning_effort", "reasoning_summary"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ReasoningEffort.Get() != nil && !all.ReasoningEffort.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.ReasoningEffort = all.ReasoningEffort
	}
	if all.ReasoningSummary.Get() != nil && !all.ReasoningSummary.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.ReasoningSummary = all.ReasoningSummary
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
