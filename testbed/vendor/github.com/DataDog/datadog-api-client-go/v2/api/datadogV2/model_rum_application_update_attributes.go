// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApplicationUpdateAttributes RUM application update attributes.
type RUMApplicationUpdateAttributes struct {
	// Name of the RUM application.
	Name *string `json:"name,omitempty"`
	// Controls the retention policy for Product Analytics data derived from RUM events.
	ProductAnalyticsRetentionState *RUMProductAnalyticsRetentionState `json:"product_analytics_retention_state,omitempty"`
	// Configures which RUM events are processed and stored for the application.
	RumEventProcessingState *RUMEventProcessingState `json:"rum_event_processing_state,omitempty"`
	// Type of the RUM application. Supported values are `browser`, `ios`, `android`, `react-native`, `flutter`, `roku`, `electron`, `unity`, `kotlin-multiplatform`.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMApplicationUpdateAttributes instantiates a new RUMApplicationUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMApplicationUpdateAttributes() *RUMApplicationUpdateAttributes {
	this := RUMApplicationUpdateAttributes{}
	return &this
}

// NewRUMApplicationUpdateAttributesWithDefaults instantiates a new RUMApplicationUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMApplicationUpdateAttributesWithDefaults() *RUMApplicationUpdateAttributes {
	this := RUMApplicationUpdateAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RUMApplicationUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetProductAnalyticsRetentionState returns the ProductAnalyticsRetentionState field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetProductAnalyticsRetentionState() RUMProductAnalyticsRetentionState {
	if o == nil || o.ProductAnalyticsRetentionState == nil {
		var ret RUMProductAnalyticsRetentionState
		return ret
	}
	return *o.ProductAnalyticsRetentionState
}

// GetProductAnalyticsRetentionStateOk returns a tuple with the ProductAnalyticsRetentionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetProductAnalyticsRetentionStateOk() (*RUMProductAnalyticsRetentionState, bool) {
	if o == nil || o.ProductAnalyticsRetentionState == nil {
		return nil, false
	}
	return o.ProductAnalyticsRetentionState, true
}

// HasProductAnalyticsRetentionState returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasProductAnalyticsRetentionState() bool {
	return o != nil && o.ProductAnalyticsRetentionState != nil
}

// SetProductAnalyticsRetentionState gets a reference to the given RUMProductAnalyticsRetentionState and assigns it to the ProductAnalyticsRetentionState field.
func (o *RUMApplicationUpdateAttributes) SetProductAnalyticsRetentionState(v RUMProductAnalyticsRetentionState) {
	o.ProductAnalyticsRetentionState = &v
}

// GetRumEventProcessingState returns the RumEventProcessingState field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetRumEventProcessingState() RUMEventProcessingState {
	if o == nil || o.RumEventProcessingState == nil {
		var ret RUMEventProcessingState
		return ret
	}
	return *o.RumEventProcessingState
}

// GetRumEventProcessingStateOk returns a tuple with the RumEventProcessingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetRumEventProcessingStateOk() (*RUMEventProcessingState, bool) {
	if o == nil || o.RumEventProcessingState == nil {
		return nil, false
	}
	return o.RumEventProcessingState, true
}

// HasRumEventProcessingState returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasRumEventProcessingState() bool {
	return o != nil && o.RumEventProcessingState != nil
}

// SetRumEventProcessingState gets a reference to the given RUMEventProcessingState and assigns it to the RumEventProcessingState field.
func (o *RUMApplicationUpdateAttributes) SetRumEventProcessingState(v RUMEventProcessingState) {
	o.RumEventProcessingState = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RUMApplicationUpdateAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationUpdateAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RUMApplicationUpdateAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RUMApplicationUpdateAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMApplicationUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProductAnalyticsRetentionState != nil {
		toSerialize["product_analytics_retention_state"] = o.ProductAnalyticsRetentionState
	}
	if o.RumEventProcessingState != nil {
		toSerialize["rum_event_processing_state"] = o.RumEventProcessingState
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMApplicationUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                           *string                            `json:"name,omitempty"`
		ProductAnalyticsRetentionState *RUMProductAnalyticsRetentionState `json:"product_analytics_retention_state,omitempty"`
		RumEventProcessingState        *RUMEventProcessingState           `json:"rum_event_processing_state,omitempty"`
		Type                           *string                            `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "product_analytics_retention_state", "rum_event_processing_state", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.ProductAnalyticsRetentionState != nil && !all.ProductAnalyticsRetentionState.IsValid() {
		hasInvalidField = true
	} else {
		o.ProductAnalyticsRetentionState = all.ProductAnalyticsRetentionState
	}
	if all.RumEventProcessingState != nil && !all.RumEventProcessingState.IsValid() {
		hasInvalidField = true
	} else {
		o.RumEventProcessingState = all.RumEventProcessingState
	}
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
