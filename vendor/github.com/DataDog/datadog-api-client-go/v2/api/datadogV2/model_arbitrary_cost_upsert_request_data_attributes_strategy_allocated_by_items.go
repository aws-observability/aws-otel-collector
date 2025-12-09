// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems The definition of `ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems` object.
type ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems struct {
	// The `items` `allocated_tags`.
	AllocatedTags []ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItemsAllocatedTagsItems `json:"allocated_tags"`
	// The `items` `percentage`. The numeric value format should be a 32bit float value.
	Percentage float64 `json:"percentage"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems instantiates a new ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems(allocatedTags []ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItemsAllocatedTagsItems, percentage float64) *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems {
	this := ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems{}
	this.AllocatedTags = allocatedTags
	this.Percentage = percentage
	return &this
}

// NewArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItemsWithDefaults instantiates a new ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItemsWithDefaults() *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems {
	this := ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems{}
	return &this
}

// GetAllocatedTags returns the AllocatedTags field value.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems) GetAllocatedTags() []ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItemsAllocatedTagsItems {
	if o == nil {
		var ret []ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItemsAllocatedTagsItems
		return ret
	}
	return o.AllocatedTags
}

// GetAllocatedTagsOk returns a tuple with the AllocatedTags field value
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems) GetAllocatedTagsOk() (*[]ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItemsAllocatedTagsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocatedTags, true
}

// SetAllocatedTags sets field value.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems) SetAllocatedTags(v []ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItemsAllocatedTagsItems) {
	o.AllocatedTags = v
}

// GetPercentage returns the Percentage field value.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems) GetPercentage() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems) GetPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems) SetPercentage(v float64) {
	o.Percentage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["allocated_tags"] = o.AllocatedTags
	toSerialize["percentage"] = o.Percentage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllocatedTags *[]ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByItemsAllocatedTagsItems `json:"allocated_tags"`
		Percentage    *float64                                                                              `json:"percentage"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AllocatedTags == nil {
		return fmt.Errorf("required field allocated_tags missing")
	}
	if all.Percentage == nil {
		return fmt.Errorf("required field percentage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allocated_tags", "percentage"})
	} else {
		return err
	}
	o.AllocatedTags = *all.AllocatedTags
	o.Percentage = *all.Percentage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
