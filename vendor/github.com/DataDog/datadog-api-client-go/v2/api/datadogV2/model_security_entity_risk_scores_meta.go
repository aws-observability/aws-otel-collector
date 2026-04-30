// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityEntityRiskScoresMeta Metadata for pagination
type SecurityEntityRiskScoresMeta struct {
	// Current page number (1-indexed)
	PageNumber int64 `json:"pageNumber"`
	// Number of items per page
	PageSize int64 `json:"pageSize"`
	// Query ID for pagination consistency
	QueryId string `json:"queryId"`
	// Total number of entities matching the query
	TotalRowCount int64 `json:"totalRowCount"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityEntityRiskScoresMeta instantiates a new SecurityEntityRiskScoresMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityEntityRiskScoresMeta(pageNumber int64, pageSize int64, queryId string, totalRowCount int64) *SecurityEntityRiskScoresMeta {
	this := SecurityEntityRiskScoresMeta{}
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	this.QueryId = queryId
	this.TotalRowCount = totalRowCount
	return &this
}

// NewSecurityEntityRiskScoresMetaWithDefaults instantiates a new SecurityEntityRiskScoresMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityEntityRiskScoresMetaWithDefaults() *SecurityEntityRiskScoresMeta {
	this := SecurityEntityRiskScoresMeta{}
	return &this
}

// GetPageNumber returns the PageNumber field value.
func (o *SecurityEntityRiskScoresMeta) GetPageNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoresMeta) GetPageNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value.
func (o *SecurityEntityRiskScoresMeta) SetPageNumber(v int64) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value.
func (o *SecurityEntityRiskScoresMeta) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoresMeta) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value.
func (o *SecurityEntityRiskScoresMeta) SetPageSize(v int64) {
	o.PageSize = v
}

// GetQueryId returns the QueryId field value.
func (o *SecurityEntityRiskScoresMeta) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoresMeta) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value.
func (o *SecurityEntityRiskScoresMeta) SetQueryId(v string) {
	o.QueryId = v
}

// GetTotalRowCount returns the TotalRowCount field value.
func (o *SecurityEntityRiskScoresMeta) GetTotalRowCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalRowCount
}

// GetTotalRowCountOk returns a tuple with the TotalRowCount field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoresMeta) GetTotalRowCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRowCount, true
}

// SetTotalRowCount sets field value.
func (o *SecurityEntityRiskScoresMeta) SetTotalRowCount(v int64) {
	o.TotalRowCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityEntityRiskScoresMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["pageNumber"] = o.PageNumber
	toSerialize["pageSize"] = o.PageSize
	toSerialize["queryId"] = o.QueryId
	toSerialize["totalRowCount"] = o.TotalRowCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityEntityRiskScoresMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PageNumber    *int64  `json:"pageNumber"`
		PageSize      *int64  `json:"pageSize"`
		QueryId       *string `json:"queryId"`
		TotalRowCount *int64  `json:"totalRowCount"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PageNumber == nil {
		return fmt.Errorf("required field pageNumber missing")
	}
	if all.PageSize == nil {
		return fmt.Errorf("required field pageSize missing")
	}
	if all.QueryId == nil {
		return fmt.Errorf("required field queryId missing")
	}
	if all.TotalRowCount == nil {
		return fmt.Errorf("required field totalRowCount missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"pageNumber", "pageSize", "queryId", "totalRowCount"})
	} else {
		return err
	}
	o.PageNumber = *all.PageNumber
	o.PageSize = *all.PageSize
	o.QueryId = *all.QueryId
	o.TotalRowCount = *all.TotalRowCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
