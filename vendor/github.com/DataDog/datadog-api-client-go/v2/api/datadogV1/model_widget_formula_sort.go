// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetFormulaSort The formula to sort the widget by.
type WidgetFormulaSort struct {
	// The index of the formula to sort by.
	Index int64 `json:"index"`
	// Widget sorting methods.
	Order WidgetSort `json:"order"`
	// Set the sort type to formula.
	Type FormulaType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetFormulaSort instantiates a new WidgetFormulaSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetFormulaSort(index int64, order WidgetSort, typeVar FormulaType) *WidgetFormulaSort {
	this := WidgetFormulaSort{}
	this.Index = index
	this.Order = order
	this.Type = typeVar
	return &this
}

// NewWidgetFormulaSortWithDefaults instantiates a new WidgetFormulaSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetFormulaSortWithDefaults() *WidgetFormulaSort {
	this := WidgetFormulaSort{}
	return &this
}

// GetIndex returns the Index field value.
func (o *WidgetFormulaSort) GetIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *WidgetFormulaSort) GetIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value.
func (o *WidgetFormulaSort) SetIndex(v int64) {
	o.Index = v
}

// GetOrder returns the Order field value.
func (o *WidgetFormulaSort) GetOrder() WidgetSort {
	if o == nil {
		var ret WidgetSort
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *WidgetFormulaSort) GetOrderOk() (*WidgetSort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *WidgetFormulaSort) SetOrder(v WidgetSort) {
	o.Order = v
}

// GetType returns the Type field value.
func (o *WidgetFormulaSort) GetType() FormulaType {
	if o == nil {
		var ret FormulaType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WidgetFormulaSort) GetTypeOk() (*FormulaType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WidgetFormulaSort) SetType(v FormulaType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetFormulaSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["index"] = o.Index
	toSerialize["order"] = o.Order
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetFormulaSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Index *int64       `json:"index"`
		Order *WidgetSort  `json:"order"`
		Type  *FormulaType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Index == nil {
		return fmt.Errorf("required field index missing")
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"index", "order", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Index = *all.Index
	if !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = *all.Order
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
