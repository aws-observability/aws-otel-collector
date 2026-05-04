// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionGroupBySort Sort configuration for retention group by.
type RetentionGroupBySort struct {
	// Widget sorting methods.
	Order *WidgetSort `json:"order,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionGroupBySort instantiates a new RetentionGroupBySort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionGroupBySort() *RetentionGroupBySort {
	this := RetentionGroupBySort{}
	return &this
}

// NewRetentionGroupBySortWithDefaults instantiates a new RetentionGroupBySort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionGroupBySortWithDefaults() *RetentionGroupBySort {
	this := RetentionGroupBySort{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RetentionGroupBySort) GetOrder() WidgetSort {
	if o == nil || o.Order == nil {
		var ret WidgetSort
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionGroupBySort) GetOrderOk() (*WidgetSort, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RetentionGroupBySort) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given WidgetSort and assigns it to the Order field.
func (o *RetentionGroupBySort) SetOrder(v WidgetSort) {
	o.Order = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionGroupBySort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionGroupBySort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Order *WidgetSort `json:"order,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
