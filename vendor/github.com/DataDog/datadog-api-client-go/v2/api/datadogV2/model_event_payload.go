// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventPayload Event attributes.
type EventPayload struct {
	// A string used for aggregation when [correlating](https://docs.datadoghq.com/service_management/events/correlation/) events. If you specify a key, events are deduplicated to alerts based on this key. Limited to 100 characters.
	AggregationKey *string `json:"aggregation_key,omitempty"`
	// JSON object for category-specific attributes. Schema is different per event category.
	Attributes EventPayloadAttributes `json:"attributes"`
	// Event category identifying the type of event.
	Category EventCategory `json:"category"`
	// Integration ID sourced from integration manifests.
	IntegrationId *EventPayloadIntegrationId `json:"integration_id,omitempty"`
	// Free formed text associated with the event. It's suggested to use `data.attributes.attributes.custom` for well-structured attributes. Limited to 4000 characters.
	Message *string `json:"message,omitempty"`
	// A list of tags associated with the event. Maximum of 100 tags allowed.
	// Refer to [Tags docs](https://docs.datadoghq.com/getting_started/tagging/).
	Tags []string `json:"tags,omitempty"`
	// Timestamp when the event occurred. Must follow [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format.
	// For example `"2017-01-15T01:30:15.010000Z"`.
	// Defaults to the timestamp of receipt. Limited to values no older than 18 hours.
	Timestamp *string `json:"timestamp,omitempty"`
	// The title of the event. Limited to 500 characters.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventPayload instantiates a new EventPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventPayload(attributes EventPayloadAttributes, category EventCategory, title string) *EventPayload {
	this := EventPayload{}
	this.Attributes = attributes
	this.Category = category
	this.Title = title
	return &this
}

// NewEventPayloadWithDefaults instantiates a new EventPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventPayloadWithDefaults() *EventPayload {
	this := EventPayload{}
	return &this
}

// GetAggregationKey returns the AggregationKey field value if set, zero value otherwise.
func (o *EventPayload) GetAggregationKey() string {
	if o == nil || o.AggregationKey == nil {
		var ret string
		return ret
	}
	return *o.AggregationKey
}

// GetAggregationKeyOk returns a tuple with the AggregationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventPayload) GetAggregationKeyOk() (*string, bool) {
	if o == nil || o.AggregationKey == nil {
		return nil, false
	}
	return o.AggregationKey, true
}

// HasAggregationKey returns a boolean if a field has been set.
func (o *EventPayload) HasAggregationKey() bool {
	return o != nil && o.AggregationKey != nil
}

// SetAggregationKey gets a reference to the given string and assigns it to the AggregationKey field.
func (o *EventPayload) SetAggregationKey(v string) {
	o.AggregationKey = &v
}

// GetAttributes returns the Attributes field value.
func (o *EventPayload) GetAttributes() EventPayloadAttributes {
	if o == nil {
		var ret EventPayloadAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EventPayload) GetAttributesOk() (*EventPayloadAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *EventPayload) SetAttributes(v EventPayloadAttributes) {
	o.Attributes = v
}

// GetCategory returns the Category field value.
func (o *EventPayload) GetCategory() EventCategory {
	if o == nil {
		var ret EventCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *EventPayload) GetCategoryOk() (*EventCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *EventPayload) SetCategory(v EventCategory) {
	o.Category = v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *EventPayload) GetIntegrationId() EventPayloadIntegrationId {
	if o == nil || o.IntegrationId == nil {
		var ret EventPayloadIntegrationId
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventPayload) GetIntegrationIdOk() (*EventPayloadIntegrationId, bool) {
	if o == nil || o.IntegrationId == nil {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *EventPayload) HasIntegrationId() bool {
	return o != nil && o.IntegrationId != nil
}

// SetIntegrationId gets a reference to the given EventPayloadIntegrationId and assigns it to the IntegrationId field.
func (o *EventPayload) SetIntegrationId(v EventPayloadIntegrationId) {
	o.IntegrationId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EventPayload) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventPayload) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EventPayload) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EventPayload) SetMessage(v string) {
	o.Message = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EventPayload) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventPayload) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventPayload) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EventPayload) SetTags(v []string) {
	o.Tags = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *EventPayload) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventPayload) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *EventPayload) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *EventPayload) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetTitle returns the Title field value.
func (o *EventPayload) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *EventPayload) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *EventPayload) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AggregationKey != nil {
		toSerialize["aggregation_key"] = o.AggregationKey
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["category"] = o.Category
	if o.IntegrationId != nil {
		toSerialize["integration_id"] = o.IntegrationId
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	toSerialize["title"] = o.Title
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregationKey *string                    `json:"aggregation_key,omitempty"`
		Attributes     *EventPayloadAttributes    `json:"attributes"`
		Category       *EventCategory             `json:"category"`
		IntegrationId  *EventPayloadIntegrationId `json:"integration_id,omitempty"`
		Message        *string                    `json:"message,omitempty"`
		Tags           []string                   `json:"tags,omitempty"`
		Timestamp      *string                    `json:"timestamp,omitempty"`
		Title          *string                    `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}

	hasInvalidField := false
	o.AggregationKey = all.AggregationKey
	o.Attributes = *all.Attributes
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	if all.IntegrationId != nil && !all.IntegrationId.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationId = all.IntegrationId
	}
	o.Message = all.Message
	o.Tags = all.Tags
	o.Timestamp = all.Timestamp
	o.Title = *all.Title

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
