// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSuppressionsPageMeta Pagination metadata.
type SecurityMonitoringSuppressionsPageMeta struct {
	// Current page number.
	PageNumber *int64 `json:"pageNumber,omitempty"`
	// Current page size.
	PageSize *int64 `json:"pageSize,omitempty"`
	// Total count of suppressions.
	TotalCount *int64 `json:"totalCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSuppressionsPageMeta instantiates a new SecurityMonitoringSuppressionsPageMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSuppressionsPageMeta() *SecurityMonitoringSuppressionsPageMeta {
	this := SecurityMonitoringSuppressionsPageMeta{}
	return &this
}

// NewSecurityMonitoringSuppressionsPageMetaWithDefaults instantiates a new SecurityMonitoringSuppressionsPageMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSuppressionsPageMetaWithDefaults() *SecurityMonitoringSuppressionsPageMeta {
	this := SecurityMonitoringSuppressionsPageMeta{}
	return &this
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionsPageMeta) GetPageNumber() int64 {
	if o == nil || o.PageNumber == nil {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionsPageMeta) GetPageNumberOk() (*int64, bool) {
	if o == nil || o.PageNumber == nil {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionsPageMeta) HasPageNumber() bool {
	return o != nil && o.PageNumber != nil
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *SecurityMonitoringSuppressionsPageMeta) SetPageNumber(v int64) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionsPageMeta) GetPageSize() int64 {
	if o == nil || o.PageSize == nil {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionsPageMeta) GetPageSizeOk() (*int64, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionsPageMeta) HasPageSize() bool {
	return o != nil && o.PageSize != nil
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *SecurityMonitoringSuppressionsPageMeta) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionsPageMeta) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionsPageMeta) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionsPageMeta) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *SecurityMonitoringSuppressionsPageMeta) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSuppressionsPageMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PageNumber != nil {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSuppressionsPageMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PageNumber *int64 `json:"pageNumber,omitempty"`
		PageSize   *int64 `json:"pageSize,omitempty"`
		TotalCount *int64 `json:"totalCount,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"pageNumber", "pageSize", "totalCount"})
	} else {
		return err
	}
	o.PageNumber = all.PageNumber
	o.PageSize = all.PageSize
	o.TotalCount = all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
