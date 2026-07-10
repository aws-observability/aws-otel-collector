// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabPageMetaPage Pagination details for a list response.
type ModelLabPageMetaPage struct {
	// The first page number.
	FirstNumber *int64 `json:"first_number,omitempty"`
	// The last page number.
	LastNumber *int64 `json:"last_number,omitempty"`
	// The next page number.
	NextNumber datadog.NullableInt64 `json:"next_number,omitempty"`
	// The current page number.
	Number int64 `json:"number"`
	// The previous page number.
	PrevNumber datadog.NullableInt64 `json:"prev_number,omitempty"`
	// The number of items per page.
	Size int64 `json:"size"`
	// The total number of items.
	Total int64 `json:"total"`
	// The pagination type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabPageMetaPage instantiates a new ModelLabPageMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabPageMetaPage(number int64, size int64, total int64) *ModelLabPageMetaPage {
	this := ModelLabPageMetaPage{}
	this.Number = number
	this.Size = size
	this.Total = total
	return &this
}

// NewModelLabPageMetaPageWithDefaults instantiates a new ModelLabPageMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabPageMetaPageWithDefaults() *ModelLabPageMetaPage {
	this := ModelLabPageMetaPage{}
	return &this
}

// GetFirstNumber returns the FirstNumber field value if set, zero value otherwise.
func (o *ModelLabPageMetaPage) GetFirstNumber() int64 {
	if o == nil || o.FirstNumber == nil {
		var ret int64
		return ret
	}
	return *o.FirstNumber
}

// GetFirstNumberOk returns a tuple with the FirstNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLabPageMetaPage) GetFirstNumberOk() (*int64, bool) {
	if o == nil || o.FirstNumber == nil {
		return nil, false
	}
	return o.FirstNumber, true
}

// HasFirstNumber returns a boolean if a field has been set.
func (o *ModelLabPageMetaPage) HasFirstNumber() bool {
	return o != nil && o.FirstNumber != nil
}

// SetFirstNumber gets a reference to the given int64 and assigns it to the FirstNumber field.
func (o *ModelLabPageMetaPage) SetFirstNumber(v int64) {
	o.FirstNumber = &v
}

// GetLastNumber returns the LastNumber field value if set, zero value otherwise.
func (o *ModelLabPageMetaPage) GetLastNumber() int64 {
	if o == nil || o.LastNumber == nil {
		var ret int64
		return ret
	}
	return *o.LastNumber
}

// GetLastNumberOk returns a tuple with the LastNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLabPageMetaPage) GetLastNumberOk() (*int64, bool) {
	if o == nil || o.LastNumber == nil {
		return nil, false
	}
	return o.LastNumber, true
}

// HasLastNumber returns a boolean if a field has been set.
func (o *ModelLabPageMetaPage) HasLastNumber() bool {
	return o != nil && o.LastNumber != nil
}

// SetLastNumber gets a reference to the given int64 and assigns it to the LastNumber field.
func (o *ModelLabPageMetaPage) SetLastNumber(v int64) {
	o.LastNumber = &v
}

// GetNextNumber returns the NextNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabPageMetaPage) GetNextNumber() int64 {
	if o == nil || o.NextNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextNumber.Get()
}

// GetNextNumberOk returns a tuple with the NextNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabPageMetaPage) GetNextNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextNumber.Get(), o.NextNumber.IsSet()
}

// HasNextNumber returns a boolean if a field has been set.
func (o *ModelLabPageMetaPage) HasNextNumber() bool {
	return o != nil && o.NextNumber.IsSet()
}

// SetNextNumber gets a reference to the given datadog.NullableInt64 and assigns it to the NextNumber field.
func (o *ModelLabPageMetaPage) SetNextNumber(v int64) {
	o.NextNumber.Set(&v)
}

// SetNextNumberNil sets the value for NextNumber to be an explicit nil.
func (o *ModelLabPageMetaPage) SetNextNumberNil() {
	o.NextNumber.Set(nil)
}

// UnsetNextNumber ensures that no value is present for NextNumber, not even an explicit nil.
func (o *ModelLabPageMetaPage) UnsetNextNumber() {
	o.NextNumber.Unset()
}

// GetNumber returns the Number field value.
func (o *ModelLabPageMetaPage) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *ModelLabPageMetaPage) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value.
func (o *ModelLabPageMetaPage) SetNumber(v int64) {
	o.Number = v
}

// GetPrevNumber returns the PrevNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabPageMetaPage) GetPrevNumber() int64 {
	if o == nil || o.PrevNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrevNumber.Get()
}

// GetPrevNumberOk returns a tuple with the PrevNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabPageMetaPage) GetPrevNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevNumber.Get(), o.PrevNumber.IsSet()
}

// HasPrevNumber returns a boolean if a field has been set.
func (o *ModelLabPageMetaPage) HasPrevNumber() bool {
	return o != nil && o.PrevNumber.IsSet()
}

// SetPrevNumber gets a reference to the given datadog.NullableInt64 and assigns it to the PrevNumber field.
func (o *ModelLabPageMetaPage) SetPrevNumber(v int64) {
	o.PrevNumber.Set(&v)
}

// SetPrevNumberNil sets the value for PrevNumber to be an explicit nil.
func (o *ModelLabPageMetaPage) SetPrevNumberNil() {
	o.PrevNumber.Set(nil)
}

// UnsetPrevNumber ensures that no value is present for PrevNumber, not even an explicit nil.
func (o *ModelLabPageMetaPage) UnsetPrevNumber() {
	o.PrevNumber.Unset()
}

// GetSize returns the Size field value.
func (o *ModelLabPageMetaPage) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ModelLabPageMetaPage) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *ModelLabPageMetaPage) SetSize(v int64) {
	o.Size = v
}

// GetTotal returns the Total field value.
func (o *ModelLabPageMetaPage) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ModelLabPageMetaPage) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *ModelLabPageMetaPage) SetTotal(v int64) {
	o.Total = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelLabPageMetaPage) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLabPageMetaPage) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelLabPageMetaPage) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ModelLabPageMetaPage) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabPageMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FirstNumber != nil {
		toSerialize["first_number"] = o.FirstNumber
	}
	if o.LastNumber != nil {
		toSerialize["last_number"] = o.LastNumber
	}
	if o.NextNumber.IsSet() {
		toSerialize["next_number"] = o.NextNumber.Get()
	}
	toSerialize["number"] = o.Number
	if o.PrevNumber.IsSet() {
		toSerialize["prev_number"] = o.PrevNumber.Get()
	}
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabPageMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstNumber *int64                `json:"first_number,omitempty"`
		LastNumber  *int64                `json:"last_number,omitempty"`
		NextNumber  datadog.NullableInt64 `json:"next_number,omitempty"`
		Number      *int64                `json:"number"`
		PrevNumber  datadog.NullableInt64 `json:"prev_number,omitempty"`
		Size        *int64                `json:"size"`
		Total       *int64                `json:"total"`
		Type        *string               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Number == nil {
		return fmt.Errorf("required field number missing")
	}
	if all.Size == nil {
		return fmt.Errorf("required field size missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_number", "last_number", "next_number", "number", "prev_number", "size", "total", "type"})
	} else {
		return err
	}
	o.FirstNumber = all.FirstNumber
	o.LastNumber = all.LastNumber
	o.NextNumber = all.NextNumber
	o.Number = *all.Number
	o.PrevNumber = all.PrevNumber
	o.Size = *all.Size
	o.Total = *all.Total
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
