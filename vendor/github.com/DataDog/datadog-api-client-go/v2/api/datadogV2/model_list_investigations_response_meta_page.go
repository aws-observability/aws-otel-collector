// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListInvestigationsResponseMetaPage Pagination metadata.
type ListInvestigationsResponseMetaPage struct {
	// Maximum number of results per page.
	Limit int64 `json:"limit"`
	// Offset of the current page.
	Offset int64 `json:"offset"`
	// Total number of investigations.
	Total int64 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListInvestigationsResponseMetaPage instantiates a new ListInvestigationsResponseMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListInvestigationsResponseMetaPage(limit int64, offset int64, total int64) *ListInvestigationsResponseMetaPage {
	this := ListInvestigationsResponseMetaPage{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	return &this
}

// NewListInvestigationsResponseMetaPageWithDefaults instantiates a new ListInvestigationsResponseMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListInvestigationsResponseMetaPageWithDefaults() *ListInvestigationsResponseMetaPage {
	this := ListInvestigationsResponseMetaPage{}
	return &this
}

// GetLimit returns the Limit field value.
func (o *ListInvestigationsResponseMetaPage) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponseMetaPage) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *ListInvestigationsResponseMetaPage) SetLimit(v int64) {
	o.Limit = v
}

// GetOffset returns the Offset field value.
func (o *ListInvestigationsResponseMetaPage) GetOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponseMetaPage) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value.
func (o *ListInvestigationsResponseMetaPage) SetOffset(v int64) {
	o.Offset = v
}

// GetTotal returns the Total field value.
func (o *ListInvestigationsResponseMetaPage) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponseMetaPage) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *ListInvestigationsResponseMetaPage) SetTotal(v int64) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListInvestigationsResponseMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListInvestigationsResponseMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit  *int64 `json:"limit"`
		Offset *int64 `json:"offset"`
		Total  *int64 `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
		datadog.DeleteKeys(additionalProperties, &[]string{"limit", "offset", "total"})
	} else {
		return err
	}
	o.Limit = *all.Limit
	o.Offset = *all.Offset
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
