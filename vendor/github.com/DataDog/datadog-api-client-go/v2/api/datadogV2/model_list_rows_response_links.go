// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListRowsResponseLinks Pagination links for the list rows response.
type ListRowsResponseLinks struct {
	// Link to the first page of results.
	First string `json:"first"`
	// Link to the next page of results. Only present when more rows are available.
	Next *string `json:"next,omitempty"`
	// Link to the current page of results.
	Self string `json:"self"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListRowsResponseLinks instantiates a new ListRowsResponseLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListRowsResponseLinks(first string, self string) *ListRowsResponseLinks {
	this := ListRowsResponseLinks{}
	this.First = first
	this.Self = self
	return &this
}

// NewListRowsResponseLinksWithDefaults instantiates a new ListRowsResponseLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListRowsResponseLinksWithDefaults() *ListRowsResponseLinks {
	this := ListRowsResponseLinks{}
	return &this
}

// GetFirst returns the First field value.
func (o *ListRowsResponseLinks) GetFirst() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.First
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
func (o *ListRowsResponseLinks) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.First, true
}

// SetFirst sets field value.
func (o *ListRowsResponseLinks) SetFirst(v string) {
	o.First = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListRowsResponseLinks) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRowsResponseLinks) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListRowsResponseLinks) HasNext() bool {
	return o != nil && o.Next != nil
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ListRowsResponseLinks) SetNext(v string) {
	o.Next = &v
}

// GetSelf returns the Self field value.
func (o *ListRowsResponseLinks) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ListRowsResponseLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value.
func (o *ListRowsResponseLinks) SetSelf(v string) {
	o.Self = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListRowsResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["first"] = o.First
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListRowsResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		First *string `json:"first"`
		Next  *string `json:"next,omitempty"`
		Self  *string `json:"self"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.First == nil {
		return fmt.Errorf("required field first missing")
	}
	if all.Self == nil {
		return fmt.Errorf("required field self missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first", "next", "self"})
	} else {
		return err
	}
	o.First = *all.First
	o.Next = all.Next
	o.Self = *all.Self

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
