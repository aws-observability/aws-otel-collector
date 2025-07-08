// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Links The JSON:API links related to pagination.
type Links struct {
	// First page link.
	First string `json:"first"`
	// Last page link.
	Last string `json:"last"`
	// Next page link.
	Next *string `json:"next,omitempty"`
	// Previous page link.
	Previous *string `json:"previous,omitempty"`
	// Request link.
	Self string `json:"self"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLinks instantiates a new Links object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLinks(first string, last string, self string) *Links {
	this := Links{}
	this.First = first
	this.Last = last
	this.Self = self
	return &this
}

// NewLinksWithDefaults instantiates a new Links object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLinksWithDefaults() *Links {
	this := Links{}
	return &this
}

// GetFirst returns the First field value.
func (o *Links) GetFirst() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.First
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
func (o *Links) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.First, true
}

// SetFirst sets field value.
func (o *Links) SetFirst(v string) {
	o.First = v
}

// GetLast returns the Last field value.
func (o *Links) GetLast() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Last
}

// GetLastOk returns a tuple with the Last field value
// and a boolean to check if the value has been set.
func (o *Links) GetLastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last, true
}

// SetLast sets field value.
func (o *Links) SetLast(v string) {
	o.Last = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *Links) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *Links) HasNext() bool {
	return o != nil && o.Next != nil
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *Links) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *Links) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *Links) HasPrevious() bool {
	return o != nil && o.Previous != nil
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *Links) SetPrevious(v string) {
	o.Previous = &v
}

// GetSelf returns the Self field value.
func (o *Links) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *Links) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value.
func (o *Links) SetSelf(v string) {
	o.Self = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Links) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["first"] = o.First
	toSerialize["last"] = o.Last
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Links) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		First    *string `json:"first"`
		Last     *string `json:"last"`
		Next     *string `json:"next,omitempty"`
		Previous *string `json:"previous,omitempty"`
		Self     *string `json:"self"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.First == nil {
		return fmt.Errorf("required field first missing")
	}
	if all.Last == nil {
		return fmt.Errorf("required field last missing")
	}
	if all.Self == nil {
		return fmt.Errorf("required field self missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first", "last", "next", "previous", "self"})
	} else {
		return err
	}
	o.First = *all.First
	o.Last = *all.Last
	o.Next = all.Next
	o.Previous = all.Previous
	o.Self = *all.Self

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
