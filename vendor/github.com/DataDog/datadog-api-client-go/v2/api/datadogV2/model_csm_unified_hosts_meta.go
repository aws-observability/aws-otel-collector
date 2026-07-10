// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmUnifiedHostsMeta Pagination metadata for a unified hosts list response.
type CsmUnifiedHostsMeta struct {
	// The current page index (zero-based).
	PageIndex int64 `json:"page_index"`
	// The number of hosts returned per page.
	PageSize int64 `json:"page_size"`
	// The total number of hosts matching the filter criteria.
	TotalFiltered int64 `json:"total_filtered"`
	// The total number of pages available.
	TotalPages int64 `json:"total_pages"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmUnifiedHostsMeta instantiates a new CsmUnifiedHostsMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmUnifiedHostsMeta(pageIndex int64, pageSize int64, totalFiltered int64, totalPages int64) *CsmUnifiedHostsMeta {
	this := CsmUnifiedHostsMeta{}
	this.PageIndex = pageIndex
	this.PageSize = pageSize
	this.TotalFiltered = totalFiltered
	this.TotalPages = totalPages
	return &this
}

// NewCsmUnifiedHostsMetaWithDefaults instantiates a new CsmUnifiedHostsMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmUnifiedHostsMetaWithDefaults() *CsmUnifiedHostsMeta {
	this := CsmUnifiedHostsMeta{}
	return &this
}

// GetPageIndex returns the PageIndex field value.
func (o *CsmUnifiedHostsMeta) GetPageIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value
// and a boolean to check if the value has been set.
func (o *CsmUnifiedHostsMeta) GetPageIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageIndex, true
}

// SetPageIndex sets field value.
func (o *CsmUnifiedHostsMeta) SetPageIndex(v int64) {
	o.PageIndex = v
}

// GetPageSize returns the PageSize field value.
func (o *CsmUnifiedHostsMeta) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *CsmUnifiedHostsMeta) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value.
func (o *CsmUnifiedHostsMeta) SetPageSize(v int64) {
	o.PageSize = v
}

// GetTotalFiltered returns the TotalFiltered field value.
func (o *CsmUnifiedHostsMeta) GetTotalFiltered() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalFiltered
}

// GetTotalFilteredOk returns a tuple with the TotalFiltered field value
// and a boolean to check if the value has been set.
func (o *CsmUnifiedHostsMeta) GetTotalFilteredOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFiltered, true
}

// SetTotalFiltered sets field value.
func (o *CsmUnifiedHostsMeta) SetTotalFiltered(v int64) {
	o.TotalFiltered = v
}

// GetTotalPages returns the TotalPages field value.
func (o *CsmUnifiedHostsMeta) GetTotalPages() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *CsmUnifiedHostsMeta) GetTotalPagesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value.
func (o *CsmUnifiedHostsMeta) SetTotalPages(v int64) {
	o.TotalPages = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmUnifiedHostsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["page_index"] = o.PageIndex
	toSerialize["page_size"] = o.PageSize
	toSerialize["total_filtered"] = o.TotalFiltered
	toSerialize["total_pages"] = o.TotalPages

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmUnifiedHostsMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PageIndex     *int64 `json:"page_index"`
		PageSize      *int64 `json:"page_size"`
		TotalFiltered *int64 `json:"total_filtered"`
		TotalPages    *int64 `json:"total_pages"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PageIndex == nil {
		return fmt.Errorf("required field page_index missing")
	}
	if all.PageSize == nil {
		return fmt.Errorf("required field page_size missing")
	}
	if all.TotalFiltered == nil {
		return fmt.Errorf("required field total_filtered missing")
	}
	if all.TotalPages == nil {
		return fmt.Errorf("required field total_pages missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"page_index", "page_size", "total_filtered", "total_pages"})
	} else {
		return err
	}
	o.PageIndex = *all.PageIndex
	o.PageSize = *all.PageSize
	o.TotalFiltered = *all.TotalFiltered
	o.TotalPages = *all.TotalPages

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
