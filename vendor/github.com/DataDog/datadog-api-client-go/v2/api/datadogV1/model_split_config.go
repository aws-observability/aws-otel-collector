// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitConfig Encapsulates all user choices about how to split a graph.
type SplitConfig struct {
	// Maximum number of graphs to display in the widget.
	Limit int64 `json:"limit"`
	// Controls the order in which graphs appear in the split.
	Sort SplitSort `json:"sort"`
	// The dimension(s) on which to split the graph
	SplitDimensions []SplitDimension `json:"split_dimensions"`
	// Manual selection of tags making split graph widget static
	StaticSplits [][]SplitVectorEntryItem `json:"static_splits,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSplitConfig instantiates a new SplitConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitConfig(limit int64, sort SplitSort, splitDimensions []SplitDimension) *SplitConfig {
	this := SplitConfig{}
	this.Limit = limit
	this.Sort = sort
	this.SplitDimensions = splitDimensions
	return &this
}

// NewSplitConfigWithDefaults instantiates a new SplitConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitConfigWithDefaults() *SplitConfig {
	this := SplitConfig{}
	return &this
}

// GetLimit returns the Limit field value.
func (o *SplitConfig) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *SplitConfig) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *SplitConfig) SetLimit(v int64) {
	o.Limit = v
}

// GetSort returns the Sort field value.
func (o *SplitConfig) GetSort() SplitSort {
	if o == nil {
		var ret SplitSort
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *SplitConfig) GetSortOk() (*SplitSort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value.
func (o *SplitConfig) SetSort(v SplitSort) {
	o.Sort = v
}

// GetSplitDimensions returns the SplitDimensions field value.
func (o *SplitConfig) GetSplitDimensions() []SplitDimension {
	if o == nil {
		var ret []SplitDimension
		return ret
	}
	return o.SplitDimensions
}

// GetSplitDimensionsOk returns a tuple with the SplitDimensions field value
// and a boolean to check if the value has been set.
func (o *SplitConfig) GetSplitDimensionsOk() (*[]SplitDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplitDimensions, true
}

// SetSplitDimensions sets field value.
func (o *SplitConfig) SetSplitDimensions(v []SplitDimension) {
	o.SplitDimensions = v
}

// GetStaticSplits returns the StaticSplits field value if set, zero value otherwise.
func (o *SplitConfig) GetStaticSplits() [][]SplitVectorEntryItem {
	if o == nil || o.StaticSplits == nil {
		var ret [][]SplitVectorEntryItem
		return ret
	}
	return o.StaticSplits
}

// GetStaticSplitsOk returns a tuple with the StaticSplits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfig) GetStaticSplitsOk() (*[][]SplitVectorEntryItem, bool) {
	if o == nil || o.StaticSplits == nil {
		return nil, false
	}
	return &o.StaticSplits, true
}

// HasStaticSplits returns a boolean if a field has been set.
func (o *SplitConfig) HasStaticSplits() bool {
	return o != nil && o.StaticSplits != nil
}

// SetStaticSplits gets a reference to the given [][]SplitVectorEntryItem and assigns it to the StaticSplits field.
func (o *SplitConfig) SetStaticSplits(v [][]SplitVectorEntryItem) {
	o.StaticSplits = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["limit"] = o.Limit
	toSerialize["sort"] = o.Sort
	toSerialize["split_dimensions"] = o.SplitDimensions
	if o.StaticSplits != nil {
		toSerialize["static_splits"] = o.StaticSplits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit           *int64                   `json:"limit"`
		Sort            *SplitSort               `json:"sort"`
		SplitDimensions *[]SplitDimension        `json:"split_dimensions"`
		StaticSplits    [][]SplitVectorEntryItem `json:"static_splits,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.Sort == nil {
		return fmt.Errorf("required field sort missing")
	}
	if all.SplitDimensions == nil {
		return fmt.Errorf("required field split_dimensions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"limit", "sort", "split_dimensions", "static_splits"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Limit = *all.Limit
	if all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = *all.Sort
	o.SplitDimensions = *all.SplitDimensions
	o.StaticSplits = all.StaticSplits

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
