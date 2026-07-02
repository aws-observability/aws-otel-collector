// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SourcemapsListMetaPage Page information for the source maps list response.
type SourcemapsListMetaPage struct {
	// Whether there are more results available beyond the current page.
	HasMoreResults bool `json:"has_more_results"`
	// Total number of source maps matching the filter criteria.
	TotalFilteredCount int64 `json:"total_filtered_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourcemapsListMetaPage instantiates a new SourcemapsListMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourcemapsListMetaPage(hasMoreResults bool, totalFilteredCount int64) *SourcemapsListMetaPage {
	this := SourcemapsListMetaPage{}
	this.HasMoreResults = hasMoreResults
	this.TotalFilteredCount = totalFilteredCount
	return &this
}

// NewSourcemapsListMetaPageWithDefaults instantiates a new SourcemapsListMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourcemapsListMetaPageWithDefaults() *SourcemapsListMetaPage {
	this := SourcemapsListMetaPage{}
	return &this
}

// GetHasMoreResults returns the HasMoreResults field value.
func (o *SourcemapsListMetaPage) GetHasMoreResults() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasMoreResults
}

// GetHasMoreResultsOk returns a tuple with the HasMoreResults field value
// and a boolean to check if the value has been set.
func (o *SourcemapsListMetaPage) GetHasMoreResultsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMoreResults, true
}

// SetHasMoreResults sets field value.
func (o *SourcemapsListMetaPage) SetHasMoreResults(v bool) {
	o.HasMoreResults = v
}

// GetTotalFilteredCount returns the TotalFilteredCount field value.
func (o *SourcemapsListMetaPage) GetTotalFilteredCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalFilteredCount
}

// GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field value
// and a boolean to check if the value has been set.
func (o *SourcemapsListMetaPage) GetTotalFilteredCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFilteredCount, true
}

// SetTotalFilteredCount sets field value.
func (o *SourcemapsListMetaPage) SetTotalFilteredCount(v int64) {
	o.TotalFilteredCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourcemapsListMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["has_more_results"] = o.HasMoreResults
	toSerialize["total_filtered_count"] = o.TotalFilteredCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SourcemapsListMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HasMoreResults     *bool  `json:"has_more_results"`
		TotalFilteredCount *int64 `json:"total_filtered_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HasMoreResults == nil {
		return fmt.Errorf("required field has_more_results missing")
	}
	if all.TotalFilteredCount == nil {
		return fmt.Errorf("required field total_filtered_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"has_more_results", "total_filtered_count"})
	} else {
		return err
	}
	o.HasMoreResults = *all.HasMoreResults
	o.TotalFilteredCount = *all.TotalFilteredCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
