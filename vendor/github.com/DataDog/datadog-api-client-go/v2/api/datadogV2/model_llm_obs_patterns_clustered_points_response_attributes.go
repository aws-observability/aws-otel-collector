// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsClusteredPointsResponseAttributes Attributes of an LLM Observability patterns clustered points response.
type LLMObsPatternsClusteredPointsResponseAttributes struct {
	// Pagination token for the next page of points. Null if there are no more pages.
	NextPageToken datadog.NullableString `json:"next_page_token"`
	// List of clustered points.
	Points []LLMObsPatternsClusteredPoint `json:"points"`
	// Identifier of the topic the points belong to.
	TopicId string `json:"topic_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsClusteredPointsResponseAttributes instantiates a new LLMObsPatternsClusteredPointsResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsClusteredPointsResponseAttributes(nextPageToken datadog.NullableString, points []LLMObsPatternsClusteredPoint, topicId string) *LLMObsPatternsClusteredPointsResponseAttributes {
	this := LLMObsPatternsClusteredPointsResponseAttributes{}
	this.NextPageToken = nextPageToken
	this.Points = points
	this.TopicId = topicId
	return &this
}

// NewLLMObsPatternsClusteredPointsResponseAttributesWithDefaults instantiates a new LLMObsPatternsClusteredPointsResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsClusteredPointsResponseAttributesWithDefaults() *LLMObsPatternsClusteredPointsResponseAttributes {
	this := LLMObsPatternsClusteredPointsResponseAttributes{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *LLMObsPatternsClusteredPointsResponseAttributes) GetNextPageToken() string {
	if o == nil || o.NextPageToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken.Get()
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsPatternsClusteredPointsResponseAttributes) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageToken.Get(), o.NextPageToken.IsSet()
}

// SetNextPageToken sets field value.
func (o *LLMObsPatternsClusteredPointsResponseAttributes) SetNextPageToken(v string) {
	o.NextPageToken.Set(&v)
}

// GetPoints returns the Points field value.
func (o *LLMObsPatternsClusteredPointsResponseAttributes) GetPoints() []LLMObsPatternsClusteredPoint {
	if o == nil {
		var ret []LLMObsPatternsClusteredPoint
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPointsResponseAttributes) GetPointsOk() (*[]LLMObsPatternsClusteredPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Points, true
}

// SetPoints sets field value.
func (o *LLMObsPatternsClusteredPointsResponseAttributes) SetPoints(v []LLMObsPatternsClusteredPoint) {
	o.Points = v
}

// GetTopicId returns the TopicId field value.
func (o *LLMObsPatternsClusteredPointsResponseAttributes) GetTopicId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TopicId
}

// GetTopicIdOk returns a tuple with the TopicId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsClusteredPointsResponseAttributes) GetTopicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicId, true
}

// SetTopicId sets field value.
func (o *LLMObsPatternsClusteredPointsResponseAttributes) SetTopicId(v string) {
	o.TopicId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsClusteredPointsResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["next_page_token"] = o.NextPageToken.Get()
	toSerialize["points"] = o.Points
	toSerialize["topic_id"] = o.TopicId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsClusteredPointsResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextPageToken datadog.NullableString          `json:"next_page_token"`
		Points        *[]LLMObsPatternsClusteredPoint `json:"points"`
		TopicId       *string                         `json:"topic_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.NextPageToken.IsSet() {
		return fmt.Errorf("required field next_page_token missing")
	}
	if all.Points == nil {
		return fmt.Errorf("required field points missing")
	}
	if all.TopicId == nil {
		return fmt.Errorf("required field topic_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next_page_token", "points", "topic_id"})
	} else {
		return err
	}
	o.NextPageToken = all.NextPageToken
	o.Points = *all.Points
	o.TopicId = *all.TopicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
