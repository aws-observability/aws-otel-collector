// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VariantWeightRequest Variant weight request payload.
type VariantWeightRequest struct {
	// The percentage weight for this variant.
	Value float64 `json:"value"`
	// The variant ID to assign weight to.
	VariantId *uuid.UUID `json:"variant_id,omitempty"`
	// The variant key to assign weight to.
	VariantKey *string `json:"variant_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVariantWeightRequest instantiates a new VariantWeightRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVariantWeightRequest(value float64) *VariantWeightRequest {
	this := VariantWeightRequest{}
	this.Value = value
	return &this
}

// NewVariantWeightRequestWithDefaults instantiates a new VariantWeightRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVariantWeightRequestWithDefaults() *VariantWeightRequest {
	this := VariantWeightRequest{}
	return &this
}

// GetValue returns the Value field value.
func (o *VariantWeightRequest) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VariantWeightRequest) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *VariantWeightRequest) SetValue(v float64) {
	o.Value = v
}

// GetVariantId returns the VariantId field value if set, zero value otherwise.
func (o *VariantWeightRequest) GetVariantId() uuid.UUID {
	if o == nil || o.VariantId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantWeightRequest) GetVariantIdOk() (*uuid.UUID, bool) {
	if o == nil || o.VariantId == nil {
		return nil, false
	}
	return o.VariantId, true
}

// HasVariantId returns a boolean if a field has been set.
func (o *VariantWeightRequest) HasVariantId() bool {
	return o != nil && o.VariantId != nil
}

// SetVariantId gets a reference to the given uuid.UUID and assigns it to the VariantId field.
func (o *VariantWeightRequest) SetVariantId(v uuid.UUID) {
	o.VariantId = &v
}

// GetVariantKey returns the VariantKey field value if set, zero value otherwise.
func (o *VariantWeightRequest) GetVariantKey() string {
	if o == nil || o.VariantKey == nil {
		var ret string
		return ret
	}
	return *o.VariantKey
}

// GetVariantKeyOk returns a tuple with the VariantKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantWeightRequest) GetVariantKeyOk() (*string, bool) {
	if o == nil || o.VariantKey == nil {
		return nil, false
	}
	return o.VariantKey, true
}

// HasVariantKey returns a boolean if a field has been set.
func (o *VariantWeightRequest) HasVariantKey() bool {
	return o != nil && o.VariantKey != nil
}

// SetVariantKey gets a reference to the given string and assigns it to the VariantKey field.
func (o *VariantWeightRequest) SetVariantKey(v string) {
	o.VariantKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o VariantWeightRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["value"] = o.Value
	if o.VariantId != nil {
		toSerialize["variant_id"] = o.VariantId
	}
	if o.VariantKey != nil {
		toSerialize["variant_key"] = o.VariantKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VariantWeightRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value      *float64   `json:"value"`
		VariantId  *uuid.UUID `json:"variant_id,omitempty"`
		VariantKey *string    `json:"variant_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"value", "variant_id", "variant_key"})
	} else {
		return err
	}
	o.Value = *all.Value
	o.VariantId = all.VariantId
	o.VariantKey = all.VariantKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
