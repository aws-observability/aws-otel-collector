// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventSystemAttributes JSON object of event system attributes.
type EventSystemAttributes struct {
	// Event category identifying the type of event.
	Category *EventSystemAttributesCategory `json:"category,omitempty"`
	// Event identifier. This field is deprecated and will be removed in a future version. Use the `uid` field instead.
	Id *string `json:"id,omitempty"`
	// Integration ID sourced from integration manifests.
	IntegrationId *EventSystemAttributesIntegrationId `json:"integration_id,omitempty"`
	// The source type ID of the event.
	SourceId *int64 `json:"source_id,omitempty"`
	// A unique identifier for the event. You can use this identifier to query or reference the event.
	Uid *string `json:"uid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventSystemAttributes instantiates a new EventSystemAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventSystemAttributes() *EventSystemAttributes {
	this := EventSystemAttributes{}
	return &this
}

// NewEventSystemAttributesWithDefaults instantiates a new EventSystemAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventSystemAttributesWithDefaults() *EventSystemAttributes {
	this := EventSystemAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *EventSystemAttributes) GetCategory() EventSystemAttributesCategory {
	if o == nil || o.Category == nil {
		var ret EventSystemAttributesCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSystemAttributes) GetCategoryOk() (*EventSystemAttributesCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *EventSystemAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given EventSystemAttributesCategory and assigns it to the Category field.
func (o *EventSystemAttributes) SetCategory(v EventSystemAttributesCategory) {
	o.Category = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventSystemAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSystemAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventSystemAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventSystemAttributes) SetId(v string) {
	o.Id = &v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *EventSystemAttributes) GetIntegrationId() EventSystemAttributesIntegrationId {
	if o == nil || o.IntegrationId == nil {
		var ret EventSystemAttributesIntegrationId
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSystemAttributes) GetIntegrationIdOk() (*EventSystemAttributesIntegrationId, bool) {
	if o == nil || o.IntegrationId == nil {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *EventSystemAttributes) HasIntegrationId() bool {
	return o != nil && o.IntegrationId != nil
}

// SetIntegrationId gets a reference to the given EventSystemAttributesIntegrationId and assigns it to the IntegrationId field.
func (o *EventSystemAttributes) SetIntegrationId(v EventSystemAttributesIntegrationId) {
	o.IntegrationId = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *EventSystemAttributes) GetSourceId() int64 {
	if o == nil || o.SourceId == nil {
		var ret int64
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSystemAttributes) GetSourceIdOk() (*int64, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *EventSystemAttributes) HasSourceId() bool {
	return o != nil && o.SourceId != nil
}

// SetSourceId gets a reference to the given int64 and assigns it to the SourceId field.
func (o *EventSystemAttributes) SetSourceId(v int64) {
	o.SourceId = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *EventSystemAttributes) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSystemAttributes) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *EventSystemAttributes) HasUid() bool {
	return o != nil && o.Uid != nil
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *EventSystemAttributes) SetUid(v string) {
	o.Uid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventSystemAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IntegrationId != nil {
		toSerialize["integration_id"] = o.IntegrationId
	}
	if o.SourceId != nil {
		toSerialize["source_id"] = o.SourceId
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventSystemAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category      *EventSystemAttributesCategory      `json:"category,omitempty"`
		Id            *string                             `json:"id,omitempty"`
		IntegrationId *EventSystemAttributesIntegrationId `json:"integration_id,omitempty"`
		SourceId      *int64                              `json:"source_id,omitempty"`
		Uid           *string                             `json:"uid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "id", "integration_id", "source_id", "uid"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Category != nil && !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = all.Category
	}
	o.Id = all.Id
	if all.IntegrationId != nil && !all.IntegrationId.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationId = all.IntegrationId
	}
	o.SourceId = all.SourceId
	o.Uid = all.Uid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
