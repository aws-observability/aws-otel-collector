// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsHierarchyLinksResponseMetaPage Metadata related to paging information that is included in the response when querying the team hierarchy links
type TeamsHierarchyLinksResponseMetaPage struct {
	// First page number.
	FirstNumber *int64 `json:"first_number,omitempty"`
	// Last page number.
	LastNumber *int64 `json:"last_number,omitempty"`
	// Next page number.
	NextNumber datadog.NullableInt64 `json:"next_number,omitempty"`
	// Page number.
	Number *int64 `json:"number,omitempty"`
	// Previous page number.
	PrevNumber datadog.NullableInt64 `json:"prev_number,omitempty"`
	// Page size.
	Size *int64 `json:"size,omitempty"`
	// Total number of results.
	Total *int64 `json:"total,omitempty"`
	// Pagination type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsHierarchyLinksResponseMetaPage instantiates a new TeamsHierarchyLinksResponseMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsHierarchyLinksResponseMetaPage() *TeamsHierarchyLinksResponseMetaPage {
	this := TeamsHierarchyLinksResponseMetaPage{}
	return &this
}

// NewTeamsHierarchyLinksResponseMetaPageWithDefaults instantiates a new TeamsHierarchyLinksResponseMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsHierarchyLinksResponseMetaPageWithDefaults() *TeamsHierarchyLinksResponseMetaPage {
	this := TeamsHierarchyLinksResponseMetaPage{}
	return &this
}

// GetFirstNumber returns the FirstNumber field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetFirstNumber() int64 {
	if o == nil || o.FirstNumber == nil {
		var ret int64
		return ret
	}
	return *o.FirstNumber
}

// GetFirstNumberOk returns a tuple with the FirstNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetFirstNumberOk() (*int64, bool) {
	if o == nil || o.FirstNumber == nil {
		return nil, false
	}
	return o.FirstNumber, true
}

// HasFirstNumber returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasFirstNumber() bool {
	return o != nil && o.FirstNumber != nil
}

// SetFirstNumber gets a reference to the given int64 and assigns it to the FirstNumber field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetFirstNumber(v int64) {
	o.FirstNumber = &v
}

// GetLastNumber returns the LastNumber field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetLastNumber() int64 {
	if o == nil || o.LastNumber == nil {
		var ret int64
		return ret
	}
	return *o.LastNumber
}

// GetLastNumberOk returns a tuple with the LastNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetLastNumberOk() (*int64, bool) {
	if o == nil || o.LastNumber == nil {
		return nil, false
	}
	return o.LastNumber, true
}

// HasLastNumber returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasLastNumber() bool {
	return o != nil && o.LastNumber != nil
}

// SetLastNumber gets a reference to the given int64 and assigns it to the LastNumber field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetLastNumber(v int64) {
	o.LastNumber = &v
}

// GetNextNumber returns the NextNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsHierarchyLinksResponseMetaPage) GetNextNumber() int64 {
	if o == nil || o.NextNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextNumber.Get()
}

// GetNextNumberOk returns a tuple with the NextNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamsHierarchyLinksResponseMetaPage) GetNextNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextNumber.Get(), o.NextNumber.IsSet()
}

// HasNextNumber returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasNextNumber() bool {
	return o != nil && o.NextNumber.IsSet()
}

// SetNextNumber gets a reference to the given datadog.NullableInt64 and assigns it to the NextNumber field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetNextNumber(v int64) {
	o.NextNumber.Set(&v)
}

// SetNextNumberNil sets the value for NextNumber to be an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) SetNextNumberNil() {
	o.NextNumber.Set(nil)
}

// UnsetNextNumber ensures that no value is present for NextNumber, not even an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) UnsetNextNumber() {
	o.NextNumber.Unset()
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetNumber() int64 {
	if o == nil || o.Number == nil {
		var ret int64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetNumberOk() (*int64, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasNumber() bool {
	return o != nil && o.Number != nil
}

// SetNumber gets a reference to the given int64 and assigns it to the Number field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetNumber(v int64) {
	o.Number = &v
}

// GetPrevNumber returns the PrevNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsHierarchyLinksResponseMetaPage) GetPrevNumber() int64 {
	if o == nil || o.PrevNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrevNumber.Get()
}

// GetPrevNumberOk returns a tuple with the PrevNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamsHierarchyLinksResponseMetaPage) GetPrevNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevNumber.Get(), o.PrevNumber.IsSet()
}

// HasPrevNumber returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasPrevNumber() bool {
	return o != nil && o.PrevNumber.IsSet()
}

// SetPrevNumber gets a reference to the given datadog.NullableInt64 and assigns it to the PrevNumber field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetPrevNumber(v int64) {
	o.PrevNumber.Set(&v)
}

// SetPrevNumberNil sets the value for PrevNumber to be an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) SetPrevNumberNil() {
	o.PrevNumber.Set(nil)
}

// UnsetPrevNumber ensures that no value is present for PrevNumber, not even an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) UnsetPrevNumber() {
	o.PrevNumber.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetSize(v int64) {
	o.Size = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetTotal(v int64) {
	o.Total = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsHierarchyLinksResponseMetaPage) MarshalJSON() ([]byte, error) {
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
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.PrevNumber.IsSet() {
		toSerialize["prev_number"] = o.PrevNumber.Get()
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamsHierarchyLinksResponseMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstNumber *int64                `json:"first_number,omitempty"`
		LastNumber  *int64                `json:"last_number,omitempty"`
		NextNumber  datadog.NullableInt64 `json:"next_number,omitempty"`
		Number      *int64                `json:"number,omitempty"`
		PrevNumber  datadog.NullableInt64 `json:"prev_number,omitempty"`
		Size        *int64                `json:"size,omitempty"`
		Total       *int64                `json:"total,omitempty"`
		Type        *string               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_number", "last_number", "next_number", "number", "prev_number", "size", "total", "type"})
	} else {
		return err
	}
	o.FirstNumber = all.FirstNumber
	o.LastNumber = all.LastNumber
	o.NextNumber = all.NextNumber
	o.Number = all.Number
	o.PrevNumber = all.PrevNumber
	o.Size = all.Size
	o.Total = all.Total
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
