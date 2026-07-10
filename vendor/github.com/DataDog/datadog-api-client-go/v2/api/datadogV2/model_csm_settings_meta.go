// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmSettingsMeta Pagination metadata for a CSM settings list response.
type CsmSettingsMeta struct {
	// The current page index (zero-based).
	PageIndex int64 `json:"page_index"`
	// The number of resources returned per page.
	PageSize int64 `json:"page_size"`
	// The total number of resources matching the filter criteria.
	TotalFiltered int64 `json:"total_filtered"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmSettingsMeta instantiates a new CsmSettingsMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmSettingsMeta(pageIndex int64, pageSize int64, totalFiltered int64) *CsmSettingsMeta {
	this := CsmSettingsMeta{}
	this.PageIndex = pageIndex
	this.PageSize = pageSize
	this.TotalFiltered = totalFiltered
	return &this
}

// NewCsmSettingsMetaWithDefaults instantiates a new CsmSettingsMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmSettingsMetaWithDefaults() *CsmSettingsMeta {
	this := CsmSettingsMeta{}
	return &this
}

// GetPageIndex returns the PageIndex field value.
func (o *CsmSettingsMeta) GetPageIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value
// and a boolean to check if the value has been set.
func (o *CsmSettingsMeta) GetPageIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageIndex, true
}

// SetPageIndex sets field value.
func (o *CsmSettingsMeta) SetPageIndex(v int64) {
	o.PageIndex = v
}

// GetPageSize returns the PageSize field value.
func (o *CsmSettingsMeta) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *CsmSettingsMeta) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value.
func (o *CsmSettingsMeta) SetPageSize(v int64) {
	o.PageSize = v
}

// GetTotalFiltered returns the TotalFiltered field value.
func (o *CsmSettingsMeta) GetTotalFiltered() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalFiltered
}

// GetTotalFilteredOk returns a tuple with the TotalFiltered field value
// and a boolean to check if the value has been set.
func (o *CsmSettingsMeta) GetTotalFilteredOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFiltered, true
}

// SetTotalFiltered sets field value.
func (o *CsmSettingsMeta) SetTotalFiltered(v int64) {
	o.TotalFiltered = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmSettingsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["page_index"] = o.PageIndex
	toSerialize["page_size"] = o.PageSize
	toSerialize["total_filtered"] = o.TotalFiltered

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmSettingsMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PageIndex     *int64 `json:"page_index"`
		PageSize      *int64 `json:"page_size"`
		TotalFiltered *int64 `json:"total_filtered"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"page_index", "page_size", "total_filtered"})
	} else {
		return err
	}
	o.PageIndex = *all.PageIndex
	o.PageSize = *all.PageSize
	o.TotalFiltered = *all.TotalFiltered

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
