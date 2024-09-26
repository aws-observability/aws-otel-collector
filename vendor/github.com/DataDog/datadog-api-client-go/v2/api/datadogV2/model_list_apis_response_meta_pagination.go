// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAPIsResponseMetaPagination Pagination metadata information for `ListAPIsResponse`.
type ListAPIsResponseMetaPagination struct {
	// Number of items in the current page.
	Limit *int64 `json:"limit,omitempty"`
	// Offset for pagination.
	Offset *int64 `json:"offset,omitempty"`
	// Total number of items.
	TotalCount *int64 `json:"total_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListAPIsResponseMetaPagination instantiates a new ListAPIsResponseMetaPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListAPIsResponseMetaPagination() *ListAPIsResponseMetaPagination {
	this := ListAPIsResponseMetaPagination{}
	return &this
}

// NewListAPIsResponseMetaPaginationWithDefaults instantiates a new ListAPIsResponseMetaPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListAPIsResponseMetaPaginationWithDefaults() *ListAPIsResponseMetaPagination {
	this := ListAPIsResponseMetaPagination{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListAPIsResponseMetaPagination) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAPIsResponseMetaPagination) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListAPIsResponseMetaPagination) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ListAPIsResponseMetaPagination) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListAPIsResponseMetaPagination) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAPIsResponseMetaPagination) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListAPIsResponseMetaPagination) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ListAPIsResponseMetaPagination) SetOffset(v int64) {
	o.Offset = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListAPIsResponseMetaPagination) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAPIsResponseMetaPagination) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListAPIsResponseMetaPagination) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ListAPIsResponseMetaPagination) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListAPIsResponseMetaPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListAPIsResponseMetaPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit      *int64 `json:"limit,omitempty"`
		Offset     *int64 `json:"offset,omitempty"`
		TotalCount *int64 `json:"total_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"limit", "offset", "total_count"})
	} else {
		return err
	}
	o.Limit = all.Limit
	o.Offset = all.Offset
	o.TotalCount = all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
