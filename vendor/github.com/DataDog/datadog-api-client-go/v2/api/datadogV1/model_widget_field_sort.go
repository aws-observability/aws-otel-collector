// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetFieldSort Which column and order to sort by
type WidgetFieldSort struct {
	// Facet path for the column
	Column string `json:"column"`
	// Widget sorting methods.
	Order WidgetSort `json:"order"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetFieldSort instantiates a new WidgetFieldSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetFieldSort(column string, order WidgetSort) *WidgetFieldSort {
	this := WidgetFieldSort{}
	this.Column = column
	this.Order = order
	return &this
}

// NewWidgetFieldSortWithDefaults instantiates a new WidgetFieldSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetFieldSortWithDefaults() *WidgetFieldSort {
	this := WidgetFieldSort{}
	return &this
}

// GetColumn returns the Column field value.
func (o *WidgetFieldSort) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *WidgetFieldSort) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *WidgetFieldSort) SetColumn(v string) {
	o.Column = v
}

// GetOrder returns the Order field value.
func (o *WidgetFieldSort) GetOrder() WidgetSort {
	if o == nil {
		var ret WidgetSort
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *WidgetFieldSort) GetOrderOk() (*WidgetSort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *WidgetFieldSort) SetOrder(v WidgetSort) {
	o.Order = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetFieldSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["order"] = o.Order

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetFieldSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string     `json:"column"`
		Order  *WidgetSort `json:"order"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column", "order"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Column = *all.Column
	if !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = *all.Order
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
