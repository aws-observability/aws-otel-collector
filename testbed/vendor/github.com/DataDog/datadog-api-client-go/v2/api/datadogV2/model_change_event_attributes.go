// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventAttributes Change event attributes.
type ChangeEventAttributes struct {
	// Aggregation key of the event.
	AggregationKey *string `json:"aggregation_key,omitempty"`
	// The entity that made the change.
	Author *ChangeEventAttributesAuthor `json:"author,omitempty"`
	// JSON object of change metadata.
	ChangeMetadata interface{} `json:"change_metadata,omitempty"`
	// A uniquely identified resource.
	ChangedResource *ChangeEventAttributesChangedResource `json:"changed_resource,omitempty"`
	// JSON object of event system attributes.
	Evt *EventSystemAttributes `json:"evt,omitempty"`
	// A list of resources impacted by this change.
	ImpactedResources []ChangeEventAttributesImpactedResourcesItem `json:"impacted_resources,omitempty"`
	// The new state of the changed resource.
	NewValue interface{} `json:"new_value,omitempty"`
	// The previous state of the changed resource.
	PrevValue interface{} `json:"prev_value,omitempty"`
	// Service that triggered the event.
	Service *string `json:"service,omitempty"`
	// POSIX timestamp of the event.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// The title of the event.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeEventAttributes instantiates a new ChangeEventAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeEventAttributes() *ChangeEventAttributes {
	this := ChangeEventAttributes{}
	return &this
}

// NewChangeEventAttributesWithDefaults instantiates a new ChangeEventAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeEventAttributesWithDefaults() *ChangeEventAttributes {
	this := ChangeEventAttributes{}
	return &this
}

// GetAggregationKey returns the AggregationKey field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetAggregationKey() string {
	if o == nil || o.AggregationKey == nil {
		var ret string
		return ret
	}
	return *o.AggregationKey
}

// GetAggregationKeyOk returns a tuple with the AggregationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetAggregationKeyOk() (*string, bool) {
	if o == nil || o.AggregationKey == nil {
		return nil, false
	}
	return o.AggregationKey, true
}

// HasAggregationKey returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasAggregationKey() bool {
	return o != nil && o.AggregationKey != nil
}

// SetAggregationKey gets a reference to the given string and assigns it to the AggregationKey field.
func (o *ChangeEventAttributes) SetAggregationKey(v string) {
	o.AggregationKey = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetAuthor() ChangeEventAttributesAuthor {
	if o == nil || o.Author == nil {
		var ret ChangeEventAttributesAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetAuthorOk() (*ChangeEventAttributesAuthor, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given ChangeEventAttributesAuthor and assigns it to the Author field.
func (o *ChangeEventAttributes) SetAuthor(v ChangeEventAttributesAuthor) {
	o.Author = &v
}

// GetChangeMetadata returns the ChangeMetadata field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetChangeMetadata() interface{} {
	if o == nil || o.ChangeMetadata == nil {
		var ret interface{}
		return ret
	}
	return o.ChangeMetadata
}

// GetChangeMetadataOk returns a tuple with the ChangeMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetChangeMetadataOk() (*interface{}, bool) {
	if o == nil || o.ChangeMetadata == nil {
		return nil, false
	}
	return &o.ChangeMetadata, true
}

// HasChangeMetadata returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasChangeMetadata() bool {
	return o != nil && o.ChangeMetadata != nil
}

// SetChangeMetadata gets a reference to the given interface{} and assigns it to the ChangeMetadata field.
func (o *ChangeEventAttributes) SetChangeMetadata(v interface{}) {
	o.ChangeMetadata = v
}

// GetChangedResource returns the ChangedResource field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetChangedResource() ChangeEventAttributesChangedResource {
	if o == nil || o.ChangedResource == nil {
		var ret ChangeEventAttributesChangedResource
		return ret
	}
	return *o.ChangedResource
}

// GetChangedResourceOk returns a tuple with the ChangedResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetChangedResourceOk() (*ChangeEventAttributesChangedResource, bool) {
	if o == nil || o.ChangedResource == nil {
		return nil, false
	}
	return o.ChangedResource, true
}

// HasChangedResource returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasChangedResource() bool {
	return o != nil && o.ChangedResource != nil
}

// SetChangedResource gets a reference to the given ChangeEventAttributesChangedResource and assigns it to the ChangedResource field.
func (o *ChangeEventAttributes) SetChangedResource(v ChangeEventAttributesChangedResource) {
	o.ChangedResource = &v
}

// GetEvt returns the Evt field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetEvt() EventSystemAttributes {
	if o == nil || o.Evt == nil {
		var ret EventSystemAttributes
		return ret
	}
	return *o.Evt
}

// GetEvtOk returns a tuple with the Evt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetEvtOk() (*EventSystemAttributes, bool) {
	if o == nil || o.Evt == nil {
		return nil, false
	}
	return o.Evt, true
}

// HasEvt returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasEvt() bool {
	return o != nil && o.Evt != nil
}

// SetEvt gets a reference to the given EventSystemAttributes and assigns it to the Evt field.
func (o *ChangeEventAttributes) SetEvt(v EventSystemAttributes) {
	o.Evt = &v
}

// GetImpactedResources returns the ImpactedResources field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetImpactedResources() []ChangeEventAttributesImpactedResourcesItem {
	if o == nil || o.ImpactedResources == nil {
		var ret []ChangeEventAttributesImpactedResourcesItem
		return ret
	}
	return o.ImpactedResources
}

// GetImpactedResourcesOk returns a tuple with the ImpactedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetImpactedResourcesOk() (*[]ChangeEventAttributesImpactedResourcesItem, bool) {
	if o == nil || o.ImpactedResources == nil {
		return nil, false
	}
	return &o.ImpactedResources, true
}

// HasImpactedResources returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasImpactedResources() bool {
	return o != nil && o.ImpactedResources != nil
}

// SetImpactedResources gets a reference to the given []ChangeEventAttributesImpactedResourcesItem and assigns it to the ImpactedResources field.
func (o *ChangeEventAttributes) SetImpactedResources(v []ChangeEventAttributesImpactedResourcesItem) {
	o.ImpactedResources = v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetNewValue() interface{} {
	if o == nil || o.NewValue == nil {
		var ret interface{}
		return ret
	}
	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetNewValueOk() (*interface{}, bool) {
	if o == nil || o.NewValue == nil {
		return nil, false
	}
	return &o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasNewValue() bool {
	return o != nil && o.NewValue != nil
}

// SetNewValue gets a reference to the given interface{} and assigns it to the NewValue field.
func (o *ChangeEventAttributes) SetNewValue(v interface{}) {
	o.NewValue = v
}

// GetPrevValue returns the PrevValue field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetPrevValue() interface{} {
	if o == nil || o.PrevValue == nil {
		var ret interface{}
		return ret
	}
	return o.PrevValue
}

// GetPrevValueOk returns a tuple with the PrevValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetPrevValueOk() (*interface{}, bool) {
	if o == nil || o.PrevValue == nil {
		return nil, false
	}
	return &o.PrevValue, true
}

// HasPrevValue returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasPrevValue() bool {
	return o != nil && o.PrevValue != nil
}

// SetPrevValue gets a reference to the given interface{} and assigns it to the PrevValue field.
func (o *ChangeEventAttributes) SetPrevValue(v interface{}) {
	o.PrevValue = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ChangeEventAttributes) SetService(v string) {
	o.Service = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ChangeEventAttributes) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ChangeEventAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ChangeEventAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ChangeEventAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeEventAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AggregationKey != nil {
		toSerialize["aggregation_key"] = o.AggregationKey
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.ChangeMetadata != nil {
		toSerialize["change_metadata"] = o.ChangeMetadata
	}
	if o.ChangedResource != nil {
		toSerialize["changed_resource"] = o.ChangedResource
	}
	if o.Evt != nil {
		toSerialize["evt"] = o.Evt
	}
	if o.ImpactedResources != nil {
		toSerialize["impacted_resources"] = o.ImpactedResources
	}
	if o.NewValue != nil {
		toSerialize["new_value"] = o.NewValue
	}
	if o.PrevValue != nil {
		toSerialize["prev_value"] = o.PrevValue
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
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
func (o *ChangeEventAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregationKey    *string                                      `json:"aggregation_key,omitempty"`
		Author            *ChangeEventAttributesAuthor                 `json:"author,omitempty"`
		ChangeMetadata    interface{}                                  `json:"change_metadata,omitempty"`
		ChangedResource   *ChangeEventAttributesChangedResource        `json:"changed_resource,omitempty"`
		Evt               *EventSystemAttributes                       `json:"evt,omitempty"`
		ImpactedResources []ChangeEventAttributesImpactedResourcesItem `json:"impacted_resources,omitempty"`
		NewValue          interface{}                                  `json:"new_value,omitempty"`
		PrevValue         interface{}                                  `json:"prev_value,omitempty"`
		Service           *string                                      `json:"service,omitempty"`
		Timestamp         *int64                                       `json:"timestamp,omitempty"`
		Title             *string                                      `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation_key", "author", "change_metadata", "changed_resource", "evt", "impacted_resources", "new_value", "prev_value", "service", "timestamp", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AggregationKey = all.AggregationKey
	if all.Author != nil && all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = all.Author
	o.ChangeMetadata = all.ChangeMetadata
	if all.ChangedResource != nil && all.ChangedResource.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ChangedResource = all.ChangedResource
	if all.Evt != nil && all.Evt.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Evt = all.Evt
	o.ImpactedResources = all.ImpactedResources
	o.NewValue = all.NewValue
	o.PrevValue = all.PrevValue
	o.Service = all.Service
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
