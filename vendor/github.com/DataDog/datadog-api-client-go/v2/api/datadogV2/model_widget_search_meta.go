// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetSearchMeta Metadata about the search results.
type WidgetSearchMeta struct {
	// Total number of widgets created by anyone.
	CreatedByAnyoneTotal *int64 `json:"created_by_anyone_total,omitempty"`
	// Total number of widgets created by the current user.
	CreatedByYouTotal *int64 `json:"created_by_you_total,omitempty"`
	// Total number of widgets favorited by the current user.
	FavoritedByYouTotal *int64 `json:"favorited_by_you_total,omitempty"`
	// Total number of widgets matching the current filter criteria.
	FilteredTotal *int64 `json:"filtered_total,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetSearchMeta instantiates a new WidgetSearchMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetSearchMeta() *WidgetSearchMeta {
	this := WidgetSearchMeta{}
	return &this
}

// NewWidgetSearchMetaWithDefaults instantiates a new WidgetSearchMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetSearchMetaWithDefaults() *WidgetSearchMeta {
	this := WidgetSearchMeta{}
	return &this
}

// GetCreatedByAnyoneTotal returns the CreatedByAnyoneTotal field value if set, zero value otherwise.
func (o *WidgetSearchMeta) GetCreatedByAnyoneTotal() int64 {
	if o == nil || o.CreatedByAnyoneTotal == nil {
		var ret int64
		return ret
	}
	return *o.CreatedByAnyoneTotal
}

// GetCreatedByAnyoneTotalOk returns a tuple with the CreatedByAnyoneTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetSearchMeta) GetCreatedByAnyoneTotalOk() (*int64, bool) {
	if o == nil || o.CreatedByAnyoneTotal == nil {
		return nil, false
	}
	return o.CreatedByAnyoneTotal, true
}

// HasCreatedByAnyoneTotal returns a boolean if a field has been set.
func (o *WidgetSearchMeta) HasCreatedByAnyoneTotal() bool {
	return o != nil && o.CreatedByAnyoneTotal != nil
}

// SetCreatedByAnyoneTotal gets a reference to the given int64 and assigns it to the CreatedByAnyoneTotal field.
func (o *WidgetSearchMeta) SetCreatedByAnyoneTotal(v int64) {
	o.CreatedByAnyoneTotal = &v
}

// GetCreatedByYouTotal returns the CreatedByYouTotal field value if set, zero value otherwise.
func (o *WidgetSearchMeta) GetCreatedByYouTotal() int64 {
	if o == nil || o.CreatedByYouTotal == nil {
		var ret int64
		return ret
	}
	return *o.CreatedByYouTotal
}

// GetCreatedByYouTotalOk returns a tuple with the CreatedByYouTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetSearchMeta) GetCreatedByYouTotalOk() (*int64, bool) {
	if o == nil || o.CreatedByYouTotal == nil {
		return nil, false
	}
	return o.CreatedByYouTotal, true
}

// HasCreatedByYouTotal returns a boolean if a field has been set.
func (o *WidgetSearchMeta) HasCreatedByYouTotal() bool {
	return o != nil && o.CreatedByYouTotal != nil
}

// SetCreatedByYouTotal gets a reference to the given int64 and assigns it to the CreatedByYouTotal field.
func (o *WidgetSearchMeta) SetCreatedByYouTotal(v int64) {
	o.CreatedByYouTotal = &v
}

// GetFavoritedByYouTotal returns the FavoritedByYouTotal field value if set, zero value otherwise.
func (o *WidgetSearchMeta) GetFavoritedByYouTotal() int64 {
	if o == nil || o.FavoritedByYouTotal == nil {
		var ret int64
		return ret
	}
	return *o.FavoritedByYouTotal
}

// GetFavoritedByYouTotalOk returns a tuple with the FavoritedByYouTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetSearchMeta) GetFavoritedByYouTotalOk() (*int64, bool) {
	if o == nil || o.FavoritedByYouTotal == nil {
		return nil, false
	}
	return o.FavoritedByYouTotal, true
}

// HasFavoritedByYouTotal returns a boolean if a field has been set.
func (o *WidgetSearchMeta) HasFavoritedByYouTotal() bool {
	return o != nil && o.FavoritedByYouTotal != nil
}

// SetFavoritedByYouTotal gets a reference to the given int64 and assigns it to the FavoritedByYouTotal field.
func (o *WidgetSearchMeta) SetFavoritedByYouTotal(v int64) {
	o.FavoritedByYouTotal = &v
}

// GetFilteredTotal returns the FilteredTotal field value if set, zero value otherwise.
func (o *WidgetSearchMeta) GetFilteredTotal() int64 {
	if o == nil || o.FilteredTotal == nil {
		var ret int64
		return ret
	}
	return *o.FilteredTotal
}

// GetFilteredTotalOk returns a tuple with the FilteredTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetSearchMeta) GetFilteredTotalOk() (*int64, bool) {
	if o == nil || o.FilteredTotal == nil {
		return nil, false
	}
	return o.FilteredTotal, true
}

// HasFilteredTotal returns a boolean if a field has been set.
func (o *WidgetSearchMeta) HasFilteredTotal() bool {
	return o != nil && o.FilteredTotal != nil
}

// SetFilteredTotal gets a reference to the given int64 and assigns it to the FilteredTotal field.
func (o *WidgetSearchMeta) SetFilteredTotal(v int64) {
	o.FilteredTotal = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetSearchMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedByAnyoneTotal != nil {
		toSerialize["created_by_anyone_total"] = o.CreatedByAnyoneTotal
	}
	if o.CreatedByYouTotal != nil {
		toSerialize["created_by_you_total"] = o.CreatedByYouTotal
	}
	if o.FavoritedByYouTotal != nil {
		toSerialize["favorited_by_you_total"] = o.FavoritedByYouTotal
	}
	if o.FilteredTotal != nil {
		toSerialize["filtered_total"] = o.FilteredTotal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetSearchMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByAnyoneTotal *int64 `json:"created_by_anyone_total,omitempty"`
		CreatedByYouTotal    *int64 `json:"created_by_you_total,omitempty"`
		FavoritedByYouTotal  *int64 `json:"favorited_by_you_total,omitempty"`
		FilteredTotal        *int64 `json:"filtered_total,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by_anyone_total", "created_by_you_total", "favorited_by_you_total", "filtered_total"})
	} else {
		return err
	}
	o.CreatedByAnyoneTotal = all.CreatedByAnyoneTotal
	o.CreatedByYouTotal = all.CreatedByYouTotal
	o.FavoritedByYouTotal = all.FavoritedByYouTotal
	o.FilteredTotal = all.FilteredTotal

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
