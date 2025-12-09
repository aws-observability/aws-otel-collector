// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitSort Controls the order in which graphs appear in the split.
type SplitSort struct {
	// Defines the metric and aggregation used as the sort value.
	Compute *SplitConfigSortCompute `json:"compute,omitempty"`
	// Widget sorting methods.
	Order WidgetSort `json:"order"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSplitSort instantiates a new SplitSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitSort(order WidgetSort) *SplitSort {
	this := SplitSort{}
	this.Order = order
	return &this
}

// NewSplitSortWithDefaults instantiates a new SplitSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitSortWithDefaults() *SplitSort {
	this := SplitSort{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *SplitSort) GetCompute() SplitConfigSortCompute {
	if o == nil || o.Compute == nil {
		var ret SplitConfigSortCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitSort) GetComputeOk() (*SplitConfigSortCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *SplitSort) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given SplitConfigSortCompute and assigns it to the Compute field.
func (o *SplitSort) SetCompute(v SplitConfigSortCompute) {
	o.Compute = &v
}

// GetOrder returns the Order field value.
func (o *SplitSort) GetOrder() WidgetSort {
	if o == nil {
		var ret WidgetSort
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *SplitSort) GetOrderOk() (*WidgetSort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *SplitSort) SetOrder(v WidgetSort) {
	o.Order = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	toSerialize["order"] = o.Order

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute *SplitConfigSortCompute `json:"compute,omitempty"`
		Order   *WidgetSort             `json:"order"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "order"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = all.Compute
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
