// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleListResponseLinks Pagination links for navigating a report schedule list response.
type ReportScheduleListResponseLinks struct {
	// Link to the first page.
	First datadog.NullableString `json:"first,omitempty"`
	// Link to the last page, or `null` if it is unavailable.
	Last datadog.NullableString `json:"last,omitempty"`
	// Link to the next page, or `null` if it is unavailable.
	Next datadog.NullableString `json:"next,omitempty"`
	// Link to the previous page, or `null` if it is unavailable.
	Prev datadog.NullableString `json:"prev,omitempty"`
	// Link to the current page.
	Self datadog.NullableString `json:"self,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleListResponseLinks instantiates a new ReportScheduleListResponseLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleListResponseLinks() *ReportScheduleListResponseLinks {
	this := ReportScheduleListResponseLinks{}
	return &this
}

// NewReportScheduleListResponseLinksWithDefaults instantiates a new ReportScheduleListResponseLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleListResponseLinksWithDefaults() *ReportScheduleListResponseLinks {
	this := ReportScheduleListResponseLinks{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleListResponseLinks) GetFirst() string {
	if o == nil || o.First.Get() == nil {
		var ret string
		return ret
	}
	return *o.First.Get()
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleListResponseLinks) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.First.Get(), o.First.IsSet()
}

// HasFirst returns a boolean if a field has been set.
func (o *ReportScheduleListResponseLinks) HasFirst() bool {
	return o != nil && o.First.IsSet()
}

// SetFirst gets a reference to the given datadog.NullableString and assigns it to the First field.
func (o *ReportScheduleListResponseLinks) SetFirst(v string) {
	o.First.Set(&v)
}

// SetFirstNil sets the value for First to be an explicit nil.
func (o *ReportScheduleListResponseLinks) SetFirstNil() {
	o.First.Set(nil)
}

// UnsetFirst ensures that no value is present for First, not even an explicit nil.
func (o *ReportScheduleListResponseLinks) UnsetFirst() {
	o.First.Unset()
}

// GetLast returns the Last field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleListResponseLinks) GetLast() string {
	if o == nil || o.Last.Get() == nil {
		var ret string
		return ret
	}
	return *o.Last.Get()
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleListResponseLinks) GetLastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Last.Get(), o.Last.IsSet()
}

// HasLast returns a boolean if a field has been set.
func (o *ReportScheduleListResponseLinks) HasLast() bool {
	return o != nil && o.Last.IsSet()
}

// SetLast gets a reference to the given datadog.NullableString and assigns it to the Last field.
func (o *ReportScheduleListResponseLinks) SetLast(v string) {
	o.Last.Set(&v)
}

// SetLastNil sets the value for Last to be an explicit nil.
func (o *ReportScheduleListResponseLinks) SetLastNil() {
	o.Last.Set(nil)
}

// UnsetLast ensures that no value is present for Last, not even an explicit nil.
func (o *ReportScheduleListResponseLinks) UnsetLast() {
	o.Last.Unset()
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleListResponseLinks) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleListResponseLinks) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *ReportScheduleListResponseLinks) HasNext() bool {
	return o != nil && o.Next.IsSet()
}

// SetNext gets a reference to the given datadog.NullableString and assigns it to the Next field.
func (o *ReportScheduleListResponseLinks) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil.
func (o *ReportScheduleListResponseLinks) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil.
func (o *ReportScheduleListResponseLinks) UnsetNext() {
	o.Next.Unset()
}

// GetPrev returns the Prev field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleListResponseLinks) GetPrev() string {
	if o == nil || o.Prev.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prev.Get()
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleListResponseLinks) GetPrevOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prev.Get(), o.Prev.IsSet()
}

// HasPrev returns a boolean if a field has been set.
func (o *ReportScheduleListResponseLinks) HasPrev() bool {
	return o != nil && o.Prev.IsSet()
}

// SetPrev gets a reference to the given datadog.NullableString and assigns it to the Prev field.
func (o *ReportScheduleListResponseLinks) SetPrev(v string) {
	o.Prev.Set(&v)
}

// SetPrevNil sets the value for Prev to be an explicit nil.
func (o *ReportScheduleListResponseLinks) SetPrevNil() {
	o.Prev.Set(nil)
}

// UnsetPrev ensures that no value is present for Prev, not even an explicit nil.
func (o *ReportScheduleListResponseLinks) UnsetPrev() {
	o.Prev.Unset()
}

// GetSelf returns the Self field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleListResponseLinks) GetSelf() string {
	if o == nil || o.Self.Get() == nil {
		var ret string
		return ret
	}
	return *o.Self.Get()
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleListResponseLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Self.Get(), o.Self.IsSet()
}

// HasSelf returns a boolean if a field has been set.
func (o *ReportScheduleListResponseLinks) HasSelf() bool {
	return o != nil && o.Self.IsSet()
}

// SetSelf gets a reference to the given datadog.NullableString and assigns it to the Self field.
func (o *ReportScheduleListResponseLinks) SetSelf(v string) {
	o.Self.Set(&v)
}

// SetSelfNil sets the value for Self to be an explicit nil.
func (o *ReportScheduleListResponseLinks) SetSelfNil() {
	o.Self.Set(nil)
}

// UnsetSelf ensures that no value is present for Self, not even an explicit nil.
func (o *ReportScheduleListResponseLinks) UnsetSelf() {
	o.Self.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleListResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.First.IsSet() {
		toSerialize["first"] = o.First.Get()
	}
	if o.Last.IsSet() {
		toSerialize["last"] = o.Last.Get()
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Prev.IsSet() {
		toSerialize["prev"] = o.Prev.Get()
	}
	if o.Self.IsSet() {
		toSerialize["self"] = o.Self.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReportScheduleListResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		First datadog.NullableString `json:"first,omitempty"`
		Last  datadog.NullableString `json:"last,omitempty"`
		Next  datadog.NullableString `json:"next,omitempty"`
		Prev  datadog.NullableString `json:"prev,omitempty"`
		Self  datadog.NullableString `json:"self,omitempty"`
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
