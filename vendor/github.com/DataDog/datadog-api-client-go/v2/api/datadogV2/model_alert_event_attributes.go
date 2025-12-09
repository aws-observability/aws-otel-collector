// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertEventAttributes Alert event attributes.
type AlertEventAttributes struct {
	// Aggregation key of the event.
	AggregationKey *string `json:"aggregation_key,omitempty"`
	// JSON object of custom attributes.
	Custom interface{} `json:"custom,omitempty"`
	// JSON object of event system attributes.
	Evt *EventSystemAttributes `json:"evt,omitempty"`
	// The links related to the event.
	Links []AlertEventAttributesLinksItem `json:"links,omitempty"`
	// The priority of the alert.
	Priority *AlertEventAttributesPriority `json:"priority,omitempty"`
	// Service that triggered the event.
	Service *string `json:"service,omitempty"`
	// The status of the alert.
	Status *AlertEventAttributesStatus `json:"status,omitempty"`
	// POSIX timestamp of the event.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// The title of the event.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertEventAttributes instantiates a new AlertEventAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertEventAttributes() *AlertEventAttributes {
	this := AlertEventAttributes{}
	return &this
}

// NewAlertEventAttributesWithDefaults instantiates a new AlertEventAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertEventAttributesWithDefaults() *AlertEventAttributes {
	this := AlertEventAttributes{}
	return &this
}

// GetAggregationKey returns the AggregationKey field value if set, zero value otherwise.
func (o *AlertEventAttributes) GetAggregationKey() string {
	if o == nil || o.AggregationKey == nil {
		var ret string
		return ret
	}
	return *o.AggregationKey
}

// GetAggregationKeyOk returns a tuple with the AggregationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributes) GetAggregationKeyOk() (*string, bool) {
	if o == nil || o.AggregationKey == nil {
		return nil, false
	}
	return o.AggregationKey, true
}

// HasAggregationKey returns a boolean if a field has been set.
func (o *AlertEventAttributes) HasAggregationKey() bool {
	return o != nil && o.AggregationKey != nil
}

// SetAggregationKey gets a reference to the given string and assigns it to the AggregationKey field.
func (o *AlertEventAttributes) SetAggregationKey(v string) {
	o.AggregationKey = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *AlertEventAttributes) GetCustom() interface{} {
	if o == nil || o.Custom == nil {
		var ret interface{}
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributes) GetCustomOk() (*interface{}, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return &o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *AlertEventAttributes) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given interface{} and assigns it to the Custom field.
func (o *AlertEventAttributes) SetCustom(v interface{}) {
	o.Custom = v
}

// GetEvt returns the Evt field value if set, zero value otherwise.
func (o *AlertEventAttributes) GetEvt() EventSystemAttributes {
	if o == nil || o.Evt == nil {
		var ret EventSystemAttributes
		return ret
	}
	return *o.Evt
}

// GetEvtOk returns a tuple with the Evt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributes) GetEvtOk() (*EventSystemAttributes, bool) {
	if o == nil || o.Evt == nil {
		return nil, false
	}
	return o.Evt, true
}

// HasEvt returns a boolean if a field has been set.
func (o *AlertEventAttributes) HasEvt() bool {
	return o != nil && o.Evt != nil
}

// SetEvt gets a reference to the given EventSystemAttributes and assigns it to the Evt field.
func (o *AlertEventAttributes) SetEvt(v EventSystemAttributes) {
	o.Evt = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlertEventAttributes) GetLinks() []AlertEventAttributesLinksItem {
	if o == nil || o.Links == nil {
		var ret []AlertEventAttributesLinksItem
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributes) GetLinksOk() (*[]AlertEventAttributesLinksItem, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlertEventAttributes) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given []AlertEventAttributesLinksItem and assigns it to the Links field.
func (o *AlertEventAttributes) SetLinks(v []AlertEventAttributesLinksItem) {
	o.Links = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AlertEventAttributes) GetPriority() AlertEventAttributesPriority {
	if o == nil || o.Priority == nil {
		var ret AlertEventAttributesPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributes) GetPriorityOk() (*AlertEventAttributesPriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AlertEventAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given AlertEventAttributesPriority and assigns it to the Priority field.
func (o *AlertEventAttributes) SetPriority(v AlertEventAttributesPriority) {
	o.Priority = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *AlertEventAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *AlertEventAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *AlertEventAttributes) SetService(v string) {
	o.Service = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlertEventAttributes) GetStatus() AlertEventAttributesStatus {
	if o == nil || o.Status == nil {
		var ret AlertEventAttributesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributes) GetStatusOk() (*AlertEventAttributesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlertEventAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AlertEventAttributesStatus and assigns it to the Status field.
func (o *AlertEventAttributes) SetStatus(v AlertEventAttributesStatus) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AlertEventAttributes) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributes) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AlertEventAttributes) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *AlertEventAttributes) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlertEventAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlertEventAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlertEventAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertEventAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AggregationKey != nil {
		toSerialize["aggregation_key"] = o.AggregationKey
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	if o.Evt != nil {
		toSerialize["evt"] = o.Evt
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertEventAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregationKey *string                         `json:"aggregation_key,omitempty"`
		Custom         interface{}                     `json:"custom,omitempty"`
		Evt            *EventSystemAttributes          `json:"evt,omitempty"`
		Links          []AlertEventAttributesLinksItem `json:"links,omitempty"`
		Priority       *AlertEventAttributesPriority   `json:"priority,omitempty"`
		Service        *string                         `json:"service,omitempty"`
		Status         *AlertEventAttributesStatus     `json:"status,omitempty"`
		Timestamp      *int64                          `json:"timestamp,omitempty"`
		Title          *string                         `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation_key", "custom", "evt", "links", "priority", "service", "status", "timestamp", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AggregationKey = all.AggregationKey
	o.Custom = all.Custom
	if all.Evt != nil && all.Evt.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Evt = all.Evt
	o.Links = all.Links
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
	o.Service = all.Service
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Timestamp = all.Timestamp
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
