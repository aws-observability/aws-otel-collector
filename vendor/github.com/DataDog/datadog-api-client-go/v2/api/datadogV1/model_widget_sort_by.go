// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetSortBy The controls for sorting the widget.
type WidgetSortBy struct {
	// The number of items to limit the widget to.
	Count *int64 `json:"count,omitempty"`
	// The array of items to sort the widget by in order.
	OrderBy []WidgetSortOrderBy `json:"order_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetSortBy instantiates a new WidgetSortBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetSortBy() *WidgetSortBy {
	this := WidgetSortBy{}
	return &this
}

// NewWidgetSortByWithDefaults instantiates a new WidgetSortBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetSortByWithDefaults() *WidgetSortBy {
	this := WidgetSortBy{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *WidgetSortBy) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetSortBy) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *WidgetSortBy) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *WidgetSortBy) SetCount(v int64) {
	o.Count = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *WidgetSortBy) GetOrderBy() []WidgetSortOrderBy {
	if o == nil || o.OrderBy == nil {
		var ret []WidgetSortOrderBy
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetSortBy) GetOrderByOk() (*[]WidgetSortOrderBy, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return &o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *WidgetSortBy) HasOrderBy() bool {
	return o != nil && o.OrderBy != nil
}

// SetOrderBy gets a reference to the given []WidgetSortOrderBy and assigns it to the OrderBy field.
func (o *WidgetSortBy) SetOrderBy(v []WidgetSortOrderBy) {
	o.OrderBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetSortBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.OrderBy != nil {
		toSerialize["order_by"] = o.OrderBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetSortBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count   *int64              `json:"count,omitempty"`
		OrderBy []WidgetSortOrderBy `json:"order_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "order_by"})
	} else {
		return err
	}
	o.Count = all.Count
	o.OrderBy = all.OrderBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
