// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VariantWeight Variant weight details.
type VariantWeight struct {
	// The timestamp when the variant weight was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Unique identifier of the variant weight assignment.
	Id *uuid.UUID `json:"id,omitempty"`
	// The timestamp when the variant weight was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The percentage weight for the variant.
	Value float64 `json:"value"`
	// A variant of a feature flag.
	Variant *Variant `json:"variant,omitempty"`
	// The variant ID.
	VariantId uuid.UUID `json:"variant_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVariantWeight instantiates a new VariantWeight object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVariantWeight(value float64, variantId uuid.UUID) *VariantWeight {
	this := VariantWeight{}
	this.Value = value
	this.VariantId = variantId
	return &this
}

// NewVariantWeightWithDefaults instantiates a new VariantWeight object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVariantWeightWithDefaults() *VariantWeight {
	this := VariantWeight{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VariantWeight) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantWeight) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VariantWeight) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VariantWeight) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VariantWeight) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantWeight) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VariantWeight) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *VariantWeight) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VariantWeight) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantWeight) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VariantWeight) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VariantWeight) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetValue returns the Value field value.
func (o *VariantWeight) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VariantWeight) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *VariantWeight) SetValue(v float64) {
	o.Value = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *VariantWeight) GetVariant() Variant {
	if o == nil || o.Variant == nil {
		var ret Variant
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantWeight) GetVariantOk() (*Variant, bool) {
	if o == nil || o.Variant == nil {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *VariantWeight) HasVariant() bool {
	return o != nil && o.Variant != nil
}

// SetVariant gets a reference to the given Variant and assigns it to the Variant field.
func (o *VariantWeight) SetVariant(v Variant) {
	o.Variant = &v
}

// GetVariantId returns the VariantId field value.
func (o *VariantWeight) GetVariantId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value
// and a boolean to check if the value has been set.
func (o *VariantWeight) GetVariantIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantId, true
}

// SetVariantId sets field value.
func (o *VariantWeight) SetVariantId(v uuid.UUID) {
	o.VariantId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VariantWeight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["value"] = o.Value
	if o.Variant != nil {
		toSerialize["variant"] = o.Variant
	}
	toSerialize["variant_id"] = o.VariantId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VariantWeight) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time `json:"created_at,omitempty"`
		Id        *uuid.UUID `json:"id,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
		Value     *float64   `json:"value"`
		Variant   *Variant   `json:"variant,omitempty"`
		VariantId *uuid.UUID `json:"variant_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	if all.VariantId == nil {
		return fmt.Errorf("required field variant_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "id", "updated_at", "value", "variant", "variant_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.Id = all.Id
	o.UpdatedAt = all.UpdatedAt
	o.Value = *all.Value
	if all.Variant != nil && all.Variant.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Variant = all.Variant
	o.VariantId = *all.VariantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
