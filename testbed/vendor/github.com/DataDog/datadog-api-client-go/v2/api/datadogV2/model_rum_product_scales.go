// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMProductScales Product Scales configuration for the RUM application.
type RUMProductScales struct {
	// Product Analytics retention scale configuration.
	ProductAnalyticsRetentionScale *RUMProductAnalyticsRetentionScale `json:"product_analytics_retention_scale,omitempty"`
	// RUM event processing scale configuration.
	RumEventProcessingScale *RUMEventProcessingScale `json:"rum_event_processing_scale,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMProductScales instantiates a new RUMProductScales object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMProductScales() *RUMProductScales {
	this := RUMProductScales{}
	return &this
}

// NewRUMProductScalesWithDefaults instantiates a new RUMProductScales object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMProductScalesWithDefaults() *RUMProductScales {
	this := RUMProductScales{}
	return &this
}

// GetProductAnalyticsRetentionScale returns the ProductAnalyticsRetentionScale field value if set, zero value otherwise.
func (o *RUMProductScales) GetProductAnalyticsRetentionScale() RUMProductAnalyticsRetentionScale {
	if o == nil || o.ProductAnalyticsRetentionScale == nil {
		var ret RUMProductAnalyticsRetentionScale
		return ret
	}
	return *o.ProductAnalyticsRetentionScale
}

// GetProductAnalyticsRetentionScaleOk returns a tuple with the ProductAnalyticsRetentionScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMProductScales) GetProductAnalyticsRetentionScaleOk() (*RUMProductAnalyticsRetentionScale, bool) {
	if o == nil || o.ProductAnalyticsRetentionScale == nil {
		return nil, false
	}
	return o.ProductAnalyticsRetentionScale, true
}

// HasProductAnalyticsRetentionScale returns a boolean if a field has been set.
func (o *RUMProductScales) HasProductAnalyticsRetentionScale() bool {
	return o != nil && o.ProductAnalyticsRetentionScale != nil
}

// SetProductAnalyticsRetentionScale gets a reference to the given RUMProductAnalyticsRetentionScale and assigns it to the ProductAnalyticsRetentionScale field.
func (o *RUMProductScales) SetProductAnalyticsRetentionScale(v RUMProductAnalyticsRetentionScale) {
	o.ProductAnalyticsRetentionScale = &v
}

// GetRumEventProcessingScale returns the RumEventProcessingScale field value if set, zero value otherwise.
func (o *RUMProductScales) GetRumEventProcessingScale() RUMEventProcessingScale {
	if o == nil || o.RumEventProcessingScale == nil {
		var ret RUMEventProcessingScale
		return ret
	}
	return *o.RumEventProcessingScale
}

// GetRumEventProcessingScaleOk returns a tuple with the RumEventProcessingScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMProductScales) GetRumEventProcessingScaleOk() (*RUMEventProcessingScale, bool) {
	if o == nil || o.RumEventProcessingScale == nil {
		return nil, false
	}
	return o.RumEventProcessingScale, true
}

// HasRumEventProcessingScale returns a boolean if a field has been set.
func (o *RUMProductScales) HasRumEventProcessingScale() bool {
	return o != nil && o.RumEventProcessingScale != nil
}

// SetRumEventProcessingScale gets a reference to the given RUMEventProcessingScale and assigns it to the RumEventProcessingScale field.
func (o *RUMProductScales) SetRumEventProcessingScale(v RUMEventProcessingScale) {
	o.RumEventProcessingScale = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMProductScales) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ProductAnalyticsRetentionScale != nil {
		toSerialize["product_analytics_retention_scale"] = o.ProductAnalyticsRetentionScale
	}
	if o.RumEventProcessingScale != nil {
		toSerialize["rum_event_processing_scale"] = o.RumEventProcessingScale
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMProductScales) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProductAnalyticsRetentionScale *RUMProductAnalyticsRetentionScale `json:"product_analytics_retention_scale,omitempty"`
		RumEventProcessingScale        *RUMEventProcessingScale           `json:"rum_event_processing_scale,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"product_analytics_retention_scale", "rum_event_processing_scale"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ProductAnalyticsRetentionScale != nil && all.ProductAnalyticsRetentionScale.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProductAnalyticsRetentionScale = all.ProductAnalyticsRetentionScale
	if all.RumEventProcessingScale != nil && all.RumEventProcessingScale.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RumEventProcessingScale = all.RumEventProcessingScale

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
