// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestsPagination Pagination metadata for flaky tests.
type FlakyTestsPagination struct {
	// Cursor for the next page of results.
	NextPage datadog.NullableString `json:"next_page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlakyTestsPagination instantiates a new FlakyTestsPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlakyTestsPagination() *FlakyTestsPagination {
	this := FlakyTestsPagination{}
	return &this
}

// NewFlakyTestsPaginationWithDefaults instantiates a new FlakyTestsPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlakyTestsPaginationWithDefaults() *FlakyTestsPagination {
	this := FlakyTestsPagination{}
	return &this
}

// GetNextPage returns the NextPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestsPagination) GetNextPage() string {
	if o == nil || o.NextPage.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextPage.Get()
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestsPagination) GetNextPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPage.Get(), o.NextPage.IsSet()
}

// HasNextPage returns a boolean if a field has been set.
func (o *FlakyTestsPagination) HasNextPage() bool {
	return o != nil && o.NextPage.IsSet()
}

// SetNextPage gets a reference to the given datadog.NullableString and assigns it to the NextPage field.
func (o *FlakyTestsPagination) SetNextPage(v string) {
	o.NextPage.Set(&v)
}

// SetNextPageNil sets the value for NextPage to be an explicit nil.
func (o *FlakyTestsPagination) SetNextPageNil() {
	o.NextPage.Set(nil)
}

// UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil.
func (o *FlakyTestsPagination) UnsetNextPage() {
	o.NextPage.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o FlakyTestsPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NextPage.IsSet() {
		toSerialize["next_page"] = o.NextPage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlakyTestsPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextPage datadog.NullableString `json:"next_page,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next_page"})
	} else {
		return err
	}
	o.NextPage = all.NextPage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
