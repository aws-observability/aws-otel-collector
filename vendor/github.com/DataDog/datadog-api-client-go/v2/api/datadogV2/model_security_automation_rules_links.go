// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAutomationRulesLinks Pagination links for the list of automation rules.
type SecurityAutomationRulesLinks struct {
	// Link to the first page of results.
	First string `json:"first"`
	// Link to the last page of results.
	Last string `json:"last"`
	// Link to the next page of results.
	Next *string `json:"next,omitempty"`
	// Link to the previous page of results.
	Prev *string `json:"prev,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAutomationRulesLinks instantiates a new SecurityAutomationRulesLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAutomationRulesLinks(first string, last string) *SecurityAutomationRulesLinks {
	this := SecurityAutomationRulesLinks{}
	this.First = first
	this.Last = last
	return &this
}

// NewSecurityAutomationRulesLinksWithDefaults instantiates a new SecurityAutomationRulesLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAutomationRulesLinksWithDefaults() *SecurityAutomationRulesLinks {
	this := SecurityAutomationRulesLinks{}
	return &this
}

// GetFirst returns the First field value.
func (o *SecurityAutomationRulesLinks) GetFirst() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.First
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
func (o *SecurityAutomationRulesLinks) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.First, true
}

// SetFirst sets field value.
func (o *SecurityAutomationRulesLinks) SetFirst(v string) {
	o.First = v
}

// GetLast returns the Last field value.
func (o *SecurityAutomationRulesLinks) GetLast() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Last
}

// GetLastOk returns a tuple with the Last field value
// and a boolean to check if the value has been set.
func (o *SecurityAutomationRulesLinks) GetLastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last, true
}

// SetLast sets field value.
func (o *SecurityAutomationRulesLinks) SetLast(v string) {
	o.Last = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *SecurityAutomationRulesLinks) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAutomationRulesLinks) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *SecurityAutomationRulesLinks) HasNext() bool {
	return o != nil && o.Next != nil
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *SecurityAutomationRulesLinks) SetNext(v string) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *SecurityAutomationRulesLinks) GetPrev() string {
	if o == nil || o.Prev == nil {
		var ret string
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAutomationRulesLinks) GetPrevOk() (*string, bool) {
	if o == nil || o.Prev == nil {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *SecurityAutomationRulesLinks) HasPrev() bool {
	return o != nil && o.Prev != nil
}

// SetPrev gets a reference to the given string and assigns it to the Prev field.
func (o *SecurityAutomationRulesLinks) SetPrev(v string) {
	o.Prev = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAutomationRulesLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["first"] = o.First
	toSerialize["last"] = o.Last
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Prev != nil {
		toSerialize["prev"] = o.Prev
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAutomationRulesLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		First *string `json:"first"`
		Last  *string `json:"last"`
		Next  *string `json:"next,omitempty"`
		Prev  *string `json:"prev,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first", "last", "next", "prev"})
	} else {
		return err
	}
	o.First = *all.First
	o.Last = *all.Last
	o.Next = all.Next
	o.Prev = all.Prev

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
