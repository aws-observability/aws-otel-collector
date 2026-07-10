// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPaginationLinks Pagination links for navigating between pages of an org group list response.
type OrgGroupPaginationLinks struct {
	// Link to the first page.
	First *string `json:"first,omitempty"`
	// Link to the last page.
	Last *string `json:"last,omitempty"`
	// Link to the next page.
	Next datadog.NullableString `json:"next,omitempty"`
	// Link to the previous page.
	Prev datadog.NullableString `json:"prev,omitempty"`
	// Link to the current page.
	Self *string `json:"self,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPaginationLinks instantiates a new OrgGroupPaginationLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPaginationLinks() *OrgGroupPaginationLinks {
	this := OrgGroupPaginationLinks{}
	return &this
}

// NewOrgGroupPaginationLinksWithDefaults instantiates a new OrgGroupPaginationLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPaginationLinksWithDefaults() *OrgGroupPaginationLinks {
	this := OrgGroupPaginationLinks{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *OrgGroupPaginationLinks) GetFirst() string {
	if o == nil || o.First == nil {
		var ret string
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPaginationLinks) GetFirstOk() (*string, bool) {
	if o == nil || o.First == nil {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *OrgGroupPaginationLinks) HasFirst() bool {
	return o != nil && o.First != nil
}

// SetFirst gets a reference to the given string and assigns it to the First field.
func (o *OrgGroupPaginationLinks) SetFirst(v string) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *OrgGroupPaginationLinks) GetLast() string {
	if o == nil || o.Last == nil {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPaginationLinks) GetLastOk() (*string, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *OrgGroupPaginationLinks) HasLast() bool {
	return o != nil && o.Last != nil
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *OrgGroupPaginationLinks) SetLast(v string) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgGroupPaginationLinks) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OrgGroupPaginationLinks) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *OrgGroupPaginationLinks) HasNext() bool {
	return o != nil && o.Next.IsSet()
}

// SetNext gets a reference to the given datadog.NullableString and assigns it to the Next field.
func (o *OrgGroupPaginationLinks) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil.
func (o *OrgGroupPaginationLinks) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil.
func (o *OrgGroupPaginationLinks) UnsetNext() {
	o.Next.Unset()
}

// GetPrev returns the Prev field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgGroupPaginationLinks) GetPrev() string {
	if o == nil || o.Prev.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prev.Get()
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OrgGroupPaginationLinks) GetPrevOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prev.Get(), o.Prev.IsSet()
}

// HasPrev returns a boolean if a field has been set.
func (o *OrgGroupPaginationLinks) HasPrev() bool {
	return o != nil && o.Prev.IsSet()
}

// SetPrev gets a reference to the given datadog.NullableString and assigns it to the Prev field.
func (o *OrgGroupPaginationLinks) SetPrev(v string) {
	o.Prev.Set(&v)
}

// SetPrevNil sets the value for Prev to be an explicit nil.
func (o *OrgGroupPaginationLinks) SetPrevNil() {
	o.Prev.Set(nil)
}

// UnsetPrev ensures that no value is present for Prev, not even an explicit nil.
func (o *OrgGroupPaginationLinks) UnsetPrev() {
	o.Prev.Unset()
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *OrgGroupPaginationLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPaginationLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *OrgGroupPaginationLinks) HasSelf() bool {
	return o != nil && o.Self != nil
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *OrgGroupPaginationLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPaginationLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.First != nil {
		toSerialize["first"] = o.First
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Prev.IsSet() {
		toSerialize["prev"] = o.Prev.Get()
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPaginationLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		First *string                `json:"first,omitempty"`
		Last  *string                `json:"last,omitempty"`
		Next  datadog.NullableString `json:"next,omitempty"`
		Prev  datadog.NullableString `json:"prev,omitempty"`
		Self  *string                `json:"self,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first", "last", "next", "prev", "self"})
	} else {
		return err
	}
	o.First = all.First
	o.Last = all.Last
	o.Next = all.Next
	o.Prev = all.Prev
	o.Self = all.Self

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
