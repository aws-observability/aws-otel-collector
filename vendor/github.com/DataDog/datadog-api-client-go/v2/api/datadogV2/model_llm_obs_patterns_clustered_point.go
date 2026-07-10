// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsClusteredPoint A single data point grouped into a topic.
type LLMObsPatternsClusteredPoint struct {
	// Identifier of the source event.
	EventId string `json:"event_id"`
	// Unique identifier of the clustered point.
	Id string `json:"id"`
	// Input text of the source span.
	Input string `json:"input"`
	// Whether the point is included in the patterns dataset.
	IsIncluded bool `json:"is_included"`
	// Whether the point is suggested for inclusion in the patterns dataset.
	IsSuggested bool `json:"is_suggested"`
	// Identifier of the source session.
	SessionId string `json:"session_id"`
	// Identifier of the source span.
	SpanId string `json:"span_id"`
	// Identifier of the topic the point belongs to.
	TopicId string `json:"topic_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsClusteredPoint instantiates a new LLMObsPatternsClusteredPoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsClusteredPoint(eventId string, id string, input string, isIncluded bool, isSuggested bool, sessionId string, spanId string, topicId string) *LLMObsPatternsClusteredPoint {
	this := LLMObsPatternsClusteredPoint{}
	this.EventId = eventId
	this.Id = id
	this.Input = input
	this.IsIncluded = isIncluded
	this.IsSuggested = isSuggested
	this.SessionId = sessionId
	this.SpanId = spanId
	this.TopicId = topicId
	return &this
}

// NewLLMObsPatternsClusteredPointWithDefaults instantiates a new LLMObsPatternsClusteredPoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsClusteredPointWithDefaults() *LLMObsPatternsClusteredPoint {
	this := LLMObsPatternsClusteredPoint{}
	return &this
}

// GetEventId returns the EventId field value.
func (o *LLMObsPatternsClusteredPoint) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPoint) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value.
func (o *LLMObsPatternsClusteredPoint) SetEventId(v string) {
	o.EventId = v
}

// GetId returns the Id field value.
func (o *LLMObsPatternsClusteredPoint) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPoint) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsPatternsClusteredPoint) SetId(v string) {
	o.Id = v
}

// GetInput returns the Input field value.
func (o *LLMObsPatternsClusteredPoint) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPoint) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value.
func (o *LLMObsPatternsClusteredPoint) SetInput(v string) {
	o.Input = v
}

// GetIsIncluded returns the IsIncluded field value.
func (o *LLMObsPatternsClusteredPoint) GetIsIncluded() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsIncluded
}

// GetIsIncludedOk returns a tuple with the IsIncluded field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPoint) GetIsIncludedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsIncluded, true
}

// SetIsIncluded sets field value.
func (o *LLMObsPatternsClusteredPoint) SetIsIncluded(v bool) {
	o.IsIncluded = v
}

// GetIsSuggested returns the IsSuggested field value.
func (o *LLMObsPatternsClusteredPoint) GetIsSuggested() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsSuggested
}

// GetIsSuggestedOk returns a tuple with the IsSuggested field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPoint) GetIsSuggestedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSuggested, true
}

// SetIsSuggested sets field value.
func (o *LLMObsPatternsClusteredPoint) SetIsSuggested(v bool) {
	o.IsSuggested = v
}

// GetSessionId returns the SessionId field value.
func (o *LLMObsPatternsClusteredPoint) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPoint) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *LLMObsPatternsClusteredPoint) SetSessionId(v string) {
	o.SessionId = v
}

// GetSpanId returns the SpanId field value.
func (o *LLMObsPatternsClusteredPoint) GetSpanId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPoint) GetSpanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanId, true
}

// SetSpanId sets field value.
func (o *LLMObsPatternsClusteredPoint) SetSpanId(v string) {
	o.SpanId = v
}

// GetTopicId returns the TopicId field value.
func (o *LLMObsPatternsClusteredPoint) GetTopicId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TopicId
}

// GetTopicIdOk returns a tuple with the TopicId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPoint) GetTopicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicId, true
}

// SetTopicId sets field value.
func (o *LLMObsPatternsClusteredPoint) SetTopicId(v string) {
	o.TopicId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsClusteredPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["event_id"] = o.EventId
	toSerialize["id"] = o.Id
	toSerialize["input"] = o.Input
	toSerialize["is_included"] = o.IsIncluded
	toSerialize["is_suggested"] = o.IsSuggested
	toSerialize["session_id"] = o.SessionId
	toSerialize["span_id"] = o.SpanId
	toSerialize["topic_id"] = o.TopicId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsClusteredPoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventId     *string `json:"event_id"`
		Id          *string `json:"id"`
		Input       *string `json:"input"`
		IsIncluded  *bool   `json:"is_included"`
		IsSuggested *bool   `json:"is_suggested"`
		SessionId   *string `json:"session_id"`
		SpanId      *string `json:"span_id"`
		TopicId     *string `json:"topic_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EventId == nil {
		return fmt.Errorf("required field event_id missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Input == nil {
		return fmt.Errorf("required field input missing")
	}
	if all.IsIncluded == nil {
		return fmt.Errorf("required field is_included missing")
	}
	if all.IsSuggested == nil {
		return fmt.Errorf("required field is_suggested missing")
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field session_id missing")
	}
	if all.SpanId == nil {
		return fmt.Errorf("required field span_id missing")
	}
	if all.TopicId == nil {
		return fmt.Errorf("required field topic_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_id", "id", "input", "is_included", "is_suggested", "session_id", "span_id", "topic_id"})
	} else {
		return err
	}
	o.EventId = *all.EventId
	o.Id = *all.Id
	o.Input = *all.Input
	o.IsIncluded = *all.IsIncluded
	o.IsSuggested = *all.IsSuggested
	o.SessionId = *all.SessionId
	o.SpanId = *all.SpanId
	o.TopicId = *all.TopicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
