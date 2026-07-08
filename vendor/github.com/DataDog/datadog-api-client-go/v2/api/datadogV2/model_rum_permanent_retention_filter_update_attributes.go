// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumPermanentRetentionFilterUpdateAttributes The configuration to update on a permanent RUM retention filter.
type RumPermanentRetentionFilterUpdateAttributes struct {
	// The configuration for cross-product retention filters. All fields are optional for partial updates.
	CrossProductSampling *RumCrossProductSamplingUpdate `json:"cross_product_sampling,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumPermanentRetentionFilterUpdateAttributes instantiates a new RumPermanentRetentionFilterUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumPermanentRetentionFilterUpdateAttributes() *RumPermanentRetentionFilterUpdateAttributes {
	this := RumPermanentRetentionFilterUpdateAttributes{}
	return &this
}

// NewRumPermanentRetentionFilterUpdateAttributesWithDefaults instantiates a new RumPermanentRetentionFilterUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumPermanentRetentionFilterUpdateAttributesWithDefaults() *RumPermanentRetentionFilterUpdateAttributes {
	this := RumPermanentRetentionFilterUpdateAttributes{}
	return &this
}

// GetCrossProductSampling returns the CrossProductSampling field value if set, zero value otherwise.
func (o *RumPermanentRetentionFilterUpdateAttributes) GetCrossProductSampling() RumCrossProductSamplingUpdate {
	if o == nil || o.CrossProductSampling == nil {
		var ret RumCrossProductSamplingUpdate
		return ret
	}
	return *o.CrossProductSampling
}

// GetCrossProductSamplingOk returns a tuple with the CrossProductSampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumPermanentRetentionFilterUpdateAttributes) GetCrossProductSamplingOk() (*RumCrossProductSamplingUpdate, bool) {
	if o == nil || o.CrossProductSampling == nil {
		return nil, false
	}
	return o.CrossProductSampling, true
}

// HasCrossProductSampling returns a boolean if a field has been set.
func (o *RumPermanentRetentionFilterUpdateAttributes) HasCrossProductSampling() bool {
	return o != nil && o.CrossProductSampling != nil
}

// SetCrossProductSampling gets a reference to the given RumCrossProductSamplingUpdate and assigns it to the CrossProductSampling field.
func (o *RumPermanentRetentionFilterUpdateAttributes) SetCrossProductSampling(v RumCrossProductSamplingUpdate) {
	o.CrossProductSampling = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumPermanentRetentionFilterUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrossProductSampling != nil {
		toSerialize["cross_product_sampling"] = o.CrossProductSampling
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumPermanentRetentionFilterUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossProductSampling *RumCrossProductSamplingUpdate `json:"cross_product_sampling,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cross_product_sampling"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CrossProductSampling != nil && all.CrossProductSampling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CrossProductSampling = all.CrossProductSampling

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
