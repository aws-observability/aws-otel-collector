// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ItemApiPayloadMetaPage Pagination information for a collection of datastore items.
type ItemApiPayloadMetaPage struct {
	// Whether there are additional pages of items beyond the current page.
	HasMore *bool `json:"hasMore,omitempty"`
	// The total number of items in the datastore, ignoring any filters.
	TotalCount *int64 `json:"totalCount,omitempty"`
	// The total number of items that match the current filter criteria.
	TotalFilteredCount *int64 `json:"totalFilteredCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewItemApiPayloadMetaPage instantiates a new ItemApiPayloadMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewItemApiPayloadMetaPage() *ItemApiPayloadMetaPage {
	this := ItemApiPayloadMetaPage{}
	return &this
}

// NewItemApiPayloadMetaPageWithDefaults instantiates a new ItemApiPayloadMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewItemApiPayloadMetaPageWithDefaults() *ItemApiPayloadMetaPage {
	this := ItemApiPayloadMetaPage{}
	return &this
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ItemApiPayloadMetaPage) GetHasMore() bool {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadMetaPage) GetHasMoreOk() (*bool, bool) {
	if o == nil || o.HasMore == nil {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ItemApiPayloadMetaPage) HasHasMore() bool {
	return o != nil && o.HasMore != nil
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ItemApiPayloadMetaPage) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ItemApiPayloadMetaPage) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadMetaPage) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ItemApiPayloadMetaPage) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ItemApiPayloadMetaPage) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetTotalFilteredCount returns the TotalFilteredCount field value if set, zero value otherwise.
func (o *ItemApiPayloadMetaPage) GetTotalFilteredCount() int64 {
	if o == nil || o.TotalFilteredCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalFilteredCount
}

// GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadMetaPage) GetTotalFilteredCountOk() (*int64, bool) {
	if o == nil || o.TotalFilteredCount == nil {
		return nil, false
	}
	return o.TotalFilteredCount, true
}

// HasTotalFilteredCount returns a boolean if a field has been set.
func (o *ItemApiPayloadMetaPage) HasTotalFilteredCount() bool {
	return o != nil && o.TotalFilteredCount != nil
}

// SetTotalFilteredCount gets a reference to the given int64 and assigns it to the TotalFilteredCount field.
func (o *ItemApiPayloadMetaPage) SetTotalFilteredCount(v int64) {
	o.TotalFilteredCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ItemApiPayloadMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HasMore != nil {
		toSerialize["hasMore"] = o.HasMore
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.TotalFilteredCount != nil {
		toSerialize["totalFilteredCount"] = o.TotalFilteredCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ItemApiPayloadMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HasMore            *bool  `json:"hasMore,omitempty"`
		TotalCount         *int64 `json:"totalCount,omitempty"`
		TotalFilteredCount *int64 `json:"totalFilteredCount,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hasMore", "totalCount", "totalFilteredCount"})
	} else {
		return err
	}
	o.HasMore = all.HasMore
	o.TotalCount = all.TotalCount
	o.TotalFilteredCount = all.TotalFilteredCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
