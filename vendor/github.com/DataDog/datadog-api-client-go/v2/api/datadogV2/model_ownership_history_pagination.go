// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipHistoryPagination Cursor-based pagination metadata for the history response.
type OwnershipHistoryPagination struct {
	// Whether more history entries are available beyond this page.
	HasMore bool `json:"has_more"`
	// An opaque, base64-encoded cursor token. Pass it as the `cursor` query parameter to retrieve the next page. Absent or `null` when there are no further pages.
	NextCursor datadog.NullableString `json:"next_cursor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipHistoryPagination instantiates a new OwnershipHistoryPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipHistoryPagination(hasMore bool) *OwnershipHistoryPagination {
	this := OwnershipHistoryPagination{}
	this.HasMore = hasMore
	return &this
}

// NewOwnershipHistoryPaginationWithDefaults instantiates a new OwnershipHistoryPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipHistoryPaginationWithDefaults() *OwnershipHistoryPagination {
	this := OwnershipHistoryPagination{}
	return &this
}

// GetHasMore returns the HasMore field value.
func (o *OwnershipHistoryPagination) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryPagination) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value.
func (o *OwnershipHistoryPagination) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnershipHistoryPagination) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipHistoryPagination) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *OwnershipHistoryPagination) HasNextCursor() bool {
	return o != nil && o.NextCursor.IsSet()
}

// SetNextCursor gets a reference to the given datadog.NullableString and assigns it to the NextCursor field.
func (o *OwnershipHistoryPagination) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// SetNextCursorNil sets the value for NextCursor to be an explicit nil.
func (o *OwnershipHistoryPagination) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil.
func (o *OwnershipHistoryPagination) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipHistoryPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["has_more"] = o.HasMore
	if o.NextCursor.IsSet() {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipHistoryPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HasMore    *bool                  `json:"has_more"`
		NextCursor datadog.NullableString `json:"next_cursor,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HasMore == nil {
		return fmt.Errorf("required field has_more missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"has_more", "next_cursor"})
	} else {
		return err
	}
	o.HasMore = *all.HasMore
	o.NextCursor = all.NextCursor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
