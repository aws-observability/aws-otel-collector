// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationBulkSubscriptionResultItem A single result entry returned by the bulk subscription endpoint.
type SampleLogGenerationBulkSubscriptionResultItem struct {
	// The attributes describing a sample log generation subscription.
	Attributes SampleLogGenerationSubscriptionAttributes `json:"attributes"`
	// The unique identifier of the subscription, when one was created.
	Id string `json:"id"`
	// Per-item status returned for a bulk subscription request.
	Meta SampleLogGenerationBulkSubscriptionItemMeta `json:"meta"`
	// The type of the resource. The value should always be `subscriptions`.
	Type SampleLogGenerationSubscriptionResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSampleLogGenerationBulkSubscriptionResultItem instantiates a new SampleLogGenerationBulkSubscriptionResultItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSampleLogGenerationBulkSubscriptionResultItem(attributes SampleLogGenerationSubscriptionAttributes, id string, meta SampleLogGenerationBulkSubscriptionItemMeta, typeVar SampleLogGenerationSubscriptionResourceType) *SampleLogGenerationBulkSubscriptionResultItem {
	this := SampleLogGenerationBulkSubscriptionResultItem{}
	this.Attributes = attributes
	this.Id = id
	this.Meta = meta
	this.Type = typeVar
	return &this
}

// NewSampleLogGenerationBulkSubscriptionResultItemWithDefaults instantiates a new SampleLogGenerationBulkSubscriptionResultItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSampleLogGenerationBulkSubscriptionResultItemWithDefaults() *SampleLogGenerationBulkSubscriptionResultItem {
	this := SampleLogGenerationBulkSubscriptionResultItem{}
	var typeVar SampleLogGenerationSubscriptionResourceType = SAMPLELOGGENERATIONSUBSCRIPTIONRESOURCETYPE_SUBSCRIPTIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SampleLogGenerationBulkSubscriptionResultItem) GetAttributes() SampleLogGenerationSubscriptionAttributes {
	if o == nil {
		var ret SampleLogGenerationSubscriptionAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationBulkSubscriptionResultItem) GetAttributesOk() (*SampleLogGenerationSubscriptionAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SampleLogGenerationBulkSubscriptionResultItem) SetAttributes(v SampleLogGenerationSubscriptionAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *SampleLogGenerationBulkSubscriptionResultItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationBulkSubscriptionResultItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SampleLogGenerationBulkSubscriptionResultItem) SetId(v string) {
	o.Id = v
}

// GetMeta returns the Meta field value.
func (o *SampleLogGenerationBulkSubscriptionResultItem) GetMeta() SampleLogGenerationBulkSubscriptionItemMeta {
	if o == nil {
		var ret SampleLogGenerationBulkSubscriptionItemMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationBulkSubscriptionResultItem) GetMetaOk() (*SampleLogGenerationBulkSubscriptionItemMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *SampleLogGenerationBulkSubscriptionResultItem) SetMeta(v SampleLogGenerationBulkSubscriptionItemMeta) {
	o.Meta = v
}

// GetType returns the Type field value.
func (o *SampleLogGenerationBulkSubscriptionResultItem) GetType() SampleLogGenerationSubscriptionResourceType {
	if o == nil {
		var ret SampleLogGenerationSubscriptionResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationBulkSubscriptionResultItem) GetTypeOk() (*SampleLogGenerationSubscriptionResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SampleLogGenerationBulkSubscriptionResultItem) SetType(v SampleLogGenerationSubscriptionResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SampleLogGenerationBulkSubscriptionResultItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["meta"] = o.Meta
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SampleLogGenerationBulkSubscriptionResultItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SampleLogGenerationSubscriptionAttributes   `json:"attributes"`
		Id         *string                                      `json:"id"`
		Meta       *SampleLogGenerationBulkSubscriptionItemMeta `json:"meta"`
		Type       *SampleLogGenerationSubscriptionResourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "meta", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
