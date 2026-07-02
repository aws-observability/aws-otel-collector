// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipHistoryAttributes The attributes of an ownership history response.
type OwnershipHistoryAttributes struct {
	// The list of history entries returned for this page.
	Items []OwnershipHistoryItem `json:"items"`
	// Cursor-based pagination metadata for the history response.
	Pagination OwnershipHistoryPagination `json:"pagination"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipHistoryAttributes instantiates a new OwnershipHistoryAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipHistoryAttributes(items []OwnershipHistoryItem, pagination OwnershipHistoryPagination) *OwnershipHistoryAttributes {
	this := OwnershipHistoryAttributes{}
	this.Items = items
	this.Pagination = pagination
	return &this
}

// NewOwnershipHistoryAttributesWithDefaults instantiates a new OwnershipHistoryAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipHistoryAttributesWithDefaults() *OwnershipHistoryAttributes {
	this := OwnershipHistoryAttributes{}
	return &this
}

// GetItems returns the Items field value.
func (o *OwnershipHistoryAttributes) GetItems() []OwnershipHistoryItem {
	if o == nil {
		var ret []OwnershipHistoryItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryAttributes) GetItemsOk() (*[]OwnershipHistoryItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *OwnershipHistoryAttributes) SetItems(v []OwnershipHistoryItem) {
	o.Items = v
}

// GetPagination returns the Pagination field value.
func (o *OwnershipHistoryAttributes) GetPagination() OwnershipHistoryPagination {
	if o == nil {
		var ret OwnershipHistoryPagination
		return ret
	}
	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryAttributes) GetPaginationOk() (*OwnershipHistoryPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value.
func (o *OwnershipHistoryAttributes) SetPagination(v OwnershipHistoryPagination) {
	o.Pagination = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipHistoryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	toSerialize["pagination"] = o.Pagination

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipHistoryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      *[]OwnershipHistoryItem     `json:"items"`
		Pagination *OwnershipHistoryPagination `json:"pagination"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.Pagination == nil {
		return fmt.Errorf("required field pagination missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"items", "pagination"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Items = *all.Items
	if all.Pagination.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagination = *all.Pagination

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
