// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPaginationMetaPage Page-based pagination details.
type OrgGroupPaginationMetaPage struct {
	// The total number of items.
	TotalCount int64 `json:"total_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPaginationMetaPage instantiates a new OrgGroupPaginationMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPaginationMetaPage(totalCount int64) *OrgGroupPaginationMetaPage {
	this := OrgGroupPaginationMetaPage{}
	this.TotalCount = totalCount
	return &this
}

// NewOrgGroupPaginationMetaPageWithDefaults instantiates a new OrgGroupPaginationMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPaginationMetaPageWithDefaults() *OrgGroupPaginationMetaPage {
	this := OrgGroupPaginationMetaPage{}
	return &this
}

// GetTotalCount returns the TotalCount field value.
func (o *OrgGroupPaginationMetaPage) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPaginationMetaPage) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *OrgGroupPaginationMetaPage) SetTotalCount(v int64) {
	o.TotalCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPaginationMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPaginationMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalCount *int64 `json:"total_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field total_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"total_count"})
	} else {
		return err
	}
	o.TotalCount = *all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
