// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListInvestigationsResponseLinks Pagination links for the list investigations response.
type ListInvestigationsResponseLinks struct {
	// Link to the first page.
	First string `json:"first"`
	// Link to the last page.
	Last datadog.NullableString `json:"last,omitempty"`
	// Link to the next page.
	Next string `json:"next"`
	// Link to the previous page.
	Prev datadog.NullableString `json:"prev,omitempty"`
	// Link to the current page.
	Self string `json:"self"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListInvestigationsResponseLinks instantiates a new ListInvestigationsResponseLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListInvestigationsResponseLinks(first string, next string, self string) *ListInvestigationsResponseLinks {
	this := ListInvestigationsResponseLinks{}
	this.First = first
	this.Next = next
	this.Self = self
	return &this
}

// NewListInvestigationsResponseLinksWithDefaults instantiates a new ListInvestigationsResponseLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListInvestigationsResponseLinksWithDefaults() *ListInvestigationsResponseLinks {
	this := ListInvestigationsResponseLinks{}
	return &this
}

// GetFirst returns the First field value.
func (o *ListInvestigationsResponseLinks) GetFirst() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.First
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponseLinks) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.First, true
}

// SetFirst sets field value.
func (o *ListInvestigationsResponseLinks) SetFirst(v string) {
	o.First = v
}

// GetLast returns the Last field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListInvestigationsResponseLinks) GetLast() string {
	if o == nil || o.Last.Get() == nil {
		var ret string
		return ret
	}
	return *o.Last.Get()
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ListInvestigationsResponseLinks) GetLastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Last.Get(), o.Last.IsSet()
}

// HasLast returns a boolean if a field has been set.
func (o *ListInvestigationsResponseLinks) HasLast() bool {
	return o != nil && o.Last.IsSet()
}

// SetLast gets a reference to the given datadog.NullableString and assigns it to the Last field.
func (o *ListInvestigationsResponseLinks) SetLast(v string) {
	o.Last.Set(&v)
}

// SetLastNil sets the value for Last to be an explicit nil.
func (o *ListInvestigationsResponseLinks) SetLastNil() {
	o.Last.Set(nil)
}

// UnsetLast ensures that no value is present for Last, not even an explicit nil.
func (o *ListInvestigationsResponseLinks) UnsetLast() {
	o.Last.Unset()
}

// GetNext returns the Next field value.
func (o *ListInvestigationsResponseLinks) GetNext() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponseLinks) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value.
func (o *ListInvestigationsResponseLinks) SetNext(v string) {
	o.Next = v
}

// GetPrev returns the Prev field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListInvestigationsResponseLinks) GetPrev() string {
	if o == nil || o.Prev.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prev.Get()
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ListInvestigationsResponseLinks) GetPrevOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prev.Get(), o.Prev.IsSet()
}

// HasPrev returns a boolean if a field has been set.
func (o *ListInvestigationsResponseLinks) HasPrev() bool {
	return o != nil && o.Prev.IsSet()
}

// SetPrev gets a reference to the given datadog.NullableString and assigns it to the Prev field.
func (o *ListInvestigationsResponseLinks) SetPrev(v string) {
	o.Prev.Set(&v)
}

// SetPrevNil sets the value for Prev to be an explicit nil.
func (o *ListInvestigationsResponseLinks) SetPrevNil() {
	o.Prev.Set(nil)
}

// UnsetPrev ensures that no value is present for Prev, not even an explicit nil.
func (o *ListInvestigationsResponseLinks) UnsetPrev() {
	o.Prev.Unset()
}

// GetSelf returns the Self field value.
func (o *ListInvestigationsResponseLinks) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponseLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value.
func (o *ListInvestigationsResponseLinks) SetSelf(v string) {
	o.Self = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListInvestigationsResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["first"] = o.First
	if o.Last.IsSet() {
		toSerialize["last"] = o.Last.Get()
	}
	toSerialize["next"] = o.Next
	if o.Prev.IsSet() {
		toSerialize["prev"] = o.Prev.Get()
	}
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListInvestigationsResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		First *string                `json:"first"`
		Last  datadog.NullableString `json:"last,omitempty"`
		Next  *string                `json:"next"`
		Prev  datadog.NullableString `json:"prev,omitempty"`
		Self  *string                `json:"self"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.First == nil {
		return fmt.Errorf("required field first missing")
	}
	if all.Next == nil {
		return fmt.Errorf("required field next missing")
	}
	if all.Self == nil {
		return fmt.Errorf("required field self missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first", "last", "next", "prev", "self"})
	} else {
		return err
	}
	o.First = *all.First
	o.Last = all.Last
	o.Next = *all.Next
	o.Prev = all.Prev
	o.Self = *all.Self

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
