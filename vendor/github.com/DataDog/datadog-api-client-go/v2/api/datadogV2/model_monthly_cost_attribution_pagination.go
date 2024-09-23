// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyCostAttributionPagination The metadata for the current pagination.
type MonthlyCostAttributionPagination struct {
	// The cursor to use to get the next results, if any. To make the next request, use the same parameters with the addition of the `next_record_id`.
	NextRecordId datadog.NullableString `json:"next_record_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonthlyCostAttributionPagination instantiates a new MonthlyCostAttributionPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonthlyCostAttributionPagination() *MonthlyCostAttributionPagination {
	this := MonthlyCostAttributionPagination{}
	return &this
}

// NewMonthlyCostAttributionPaginationWithDefaults instantiates a new MonthlyCostAttributionPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonthlyCostAttributionPaginationWithDefaults() *MonthlyCostAttributionPagination {
	this := MonthlyCostAttributionPagination{}
	return &this
}

// GetNextRecordId returns the NextRecordId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyCostAttributionPagination) GetNextRecordId() string {
	if o == nil || o.NextRecordId.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextRecordId.Get()
}

// GetNextRecordIdOk returns a tuple with the NextRecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyCostAttributionPagination) GetNextRecordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextRecordId.Get(), o.NextRecordId.IsSet()
}

// HasNextRecordId returns a boolean if a field has been set.
func (o *MonthlyCostAttributionPagination) HasNextRecordId() bool {
	return o != nil && o.NextRecordId.IsSet()
}

// SetNextRecordId gets a reference to the given datadog.NullableString and assigns it to the NextRecordId field.
func (o *MonthlyCostAttributionPagination) SetNextRecordId(v string) {
	o.NextRecordId.Set(&v)
}

// SetNextRecordIdNil sets the value for NextRecordId to be an explicit nil.
func (o *MonthlyCostAttributionPagination) SetNextRecordIdNil() {
	o.NextRecordId.Set(nil)
}

// UnsetNextRecordId ensures that no value is present for NextRecordId, not even an explicit nil.
func (o *MonthlyCostAttributionPagination) UnsetNextRecordId() {
	o.NextRecordId.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o MonthlyCostAttributionPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NextRecordId.IsSet() {
		toSerialize["next_record_id"] = o.NextRecordId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonthlyCostAttributionPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextRecordId datadog.NullableString `json:"next_record_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next_record_id"})
	} else {
		return err
	}
	o.NextRecordId = all.NextRecordId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
