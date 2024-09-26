// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetGroupSort The group to sort the widget by.
type WidgetGroupSort struct {
	// The name of the group.
	Name string `json:"name"`
	// Widget sorting methods.
	Order WidgetSort `json:"order"`
	// Set the sort type to group.
	Type GroupType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetGroupSort instantiates a new WidgetGroupSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetGroupSort(name string, order WidgetSort, typeVar GroupType) *WidgetGroupSort {
	this := WidgetGroupSort{}
	this.Name = name
	this.Order = order
	this.Type = typeVar
	return &this
}

// NewWidgetGroupSortWithDefaults instantiates a new WidgetGroupSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetGroupSortWithDefaults() *WidgetGroupSort {
	this := WidgetGroupSort{}
	return &this
}

// GetName returns the Name field value.
func (o *WidgetGroupSort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WidgetGroupSort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *WidgetGroupSort) SetName(v string) {
	o.Name = v
}

// GetOrder returns the Order field value.
func (o *WidgetGroupSort) GetOrder() WidgetSort {
	if o == nil {
		var ret WidgetSort
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *WidgetGroupSort) GetOrderOk() (*WidgetSort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *WidgetGroupSort) SetOrder(v WidgetSort) {
	o.Order = v
}

// GetType returns the Type field value.
func (o *WidgetGroupSort) GetType() GroupType {
	if o == nil {
		var ret GroupType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WidgetGroupSort) GetTypeOk() (*GroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WidgetGroupSort) SetType(v GroupType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetGroupSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["order"] = o.Order
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetGroupSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string     `json:"name"`
		Order *WidgetSort `json:"order"`
		Type  *GroupType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "order", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
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
