// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PaginatedResponseMeta Metadata for scores response.
type PaginatedResponseMeta struct {
	// Number of entities in this response.
	Count int64 `json:"count"`
	// Pagination limit.
	Limit int64 `json:"limit"`
	// Pagination offset.
	Offset int64 `json:"offset"`
	// Total number of entities available.
	Total int64 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPaginatedResponseMeta instantiates a new PaginatedResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPaginatedResponseMeta(count int64, limit int64, offset int64, total int64) *PaginatedResponseMeta {
	this := PaginatedResponseMeta{}
	this.Count = count
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	return &this
}

// NewPaginatedResponseMetaWithDefaults instantiates a new PaginatedResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPaginatedResponseMetaWithDefaults() *PaginatedResponseMeta {
	this := PaginatedResponseMeta{}
	return &this
}

// GetCount returns the Count field value.
func (o *PaginatedResponseMeta) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *PaginatedResponseMeta) SetCount(v int64) {
	o.Count = v
}

// GetLimit returns the Limit field value.
func (o *PaginatedResponseMeta) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *PaginatedResponseMeta) SetLimit(v int64) {
	o.Limit = v
}

// GetOffset returns the Offset field value.
func (o *PaginatedResponseMeta) GetOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value.
func (o *PaginatedResponseMeta) SetOffset(v int64) {
	o.Offset = v
}

// GetTotal returns the Total field value.
func (o *PaginatedResponseMeta) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *PaginatedResponseMeta) SetTotal(v int64) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PaginatedResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PaginatedResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count  *int64 `json:"count"`
		Limit  *int64 `json:"limit"`
		Offset *int64 `json:"offset"`
		Total  *int64 `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.Offset == nil {
		return fmt.Errorf("required field offset missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "limit", "offset", "total"})
	} else {
		return err
	}
	o.Count = *all.Count
	o.Limit = *all.Limit
	o.Offset = *all.Offset
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
