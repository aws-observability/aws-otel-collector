// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalOrgsLinks Pagination links.
type GlobalOrgsLinks struct {
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

// NewGlobalOrgsLinks instantiates a new GlobalOrgsLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGlobalOrgsLinks() *GlobalOrgsLinks {
	this := GlobalOrgsLinks{}
	return &this
}

// NewGlobalOrgsLinksWithDefaults instantiates a new GlobalOrgsLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGlobalOrgsLinksWithDefaults() *GlobalOrgsLinks {
	this := GlobalOrgsLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalOrgsLinks) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalOrgsLinks) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *GlobalOrgsLinks) HasNext() bool {
	return o != nil && o.Next.IsSet()
}

// SetNext gets a reference to the given datadog.NullableString and assigns it to the Next field.
func (o *GlobalOrgsLinks) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil.
func (o *GlobalOrgsLinks) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil.
func (o *GlobalOrgsLinks) UnsetNext() {
	o.Next.Unset()
}

// GetPrev returns the Prev field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalOrgsLinks) GetPrev() string {
	if o == nil || o.Prev.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prev.Get()
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalOrgsLinks) GetPrevOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prev.Get(), o.Prev.IsSet()
}

// HasPrev returns a boolean if a field has been set.
func (o *GlobalOrgsLinks) HasPrev() bool {
	return o != nil && o.Prev.IsSet()
}

// SetPrev gets a reference to the given datadog.NullableString and assigns it to the Prev field.
func (o *GlobalOrgsLinks) SetPrev(v string) {
	o.Prev.Set(&v)
}

// SetPrevNil sets the value for Prev to be an explicit nil.
func (o *GlobalOrgsLinks) SetPrevNil() {
	o.Prev.Set(nil)
}

// UnsetPrev ensures that no value is present for Prev, not even an explicit nil.
func (o *GlobalOrgsLinks) UnsetPrev() {
	o.Prev.Unset()
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *GlobalOrgsLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalOrgsLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *GlobalOrgsLinks) HasSelf() bool {
	return o != nil && o.Self != nil
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *GlobalOrgsLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GlobalOrgsLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *GlobalOrgsLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Next datadog.NullableString `json:"next,omitempty"`
		Prev datadog.NullableString `json:"prev,omitempty"`
		Self *string                `json:"self,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next", "prev", "self"})
	} else {
		return err
	}
	o.Next = all.Next
	o.Prev = all.Prev
	o.Self = all.Self

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
