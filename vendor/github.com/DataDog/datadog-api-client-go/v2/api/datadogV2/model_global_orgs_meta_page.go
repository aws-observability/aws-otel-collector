// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalOrgsMetaPage Paging attributes.
type GlobalOrgsMetaPage struct {
	// The cursor used to get the current results, if any.
	Cursor *string `json:"cursor,omitempty"`
	// Number of results returned.
	Limit *int32 `json:"limit,omitempty"`
	// The cursor used to get the next results, if any.
	NextCursor datadog.NullableString `json:"next_cursor,omitempty"`
	// The cursor used to get the previous results, if any.
	PrevCursor datadog.NullableString `json:"prev_cursor,omitempty"`
	// Type of global orgs pagination.
	Type *GlobalOrgsMetaPageType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGlobalOrgsMetaPage instantiates a new GlobalOrgsMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGlobalOrgsMetaPage() *GlobalOrgsMetaPage {
	this := GlobalOrgsMetaPage{}
	return &this
}

// NewGlobalOrgsMetaPageWithDefaults instantiates a new GlobalOrgsMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGlobalOrgsMetaPageWithDefaults() *GlobalOrgsMetaPage {
	this := GlobalOrgsMetaPage{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *GlobalOrgsMetaPage) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalOrgsMetaPage) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *GlobalOrgsMetaPage) HasCursor() bool {
	return o != nil && o.Cursor != nil
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *GlobalOrgsMetaPage) SetCursor(v string) {
	o.Cursor = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GlobalOrgsMetaPage) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalOrgsMetaPage) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GlobalOrgsMetaPage) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GlobalOrgsMetaPage) SetLimit(v int32) {
	o.Limit = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalOrgsMetaPage) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalOrgsMetaPage) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *GlobalOrgsMetaPage) HasNextCursor() bool {
	return o != nil && o.NextCursor.IsSet()
}

// SetNextCursor gets a reference to the given datadog.NullableString and assigns it to the NextCursor field.
func (o *GlobalOrgsMetaPage) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// SetNextCursorNil sets the value for NextCursor to be an explicit nil.
func (o *GlobalOrgsMetaPage) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil.
func (o *GlobalOrgsMetaPage) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// GetPrevCursor returns the PrevCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalOrgsMetaPage) GetPrevCursor() string {
	if o == nil || o.PrevCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrevCursor.Get()
}

// GetPrevCursorOk returns a tuple with the PrevCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalOrgsMetaPage) GetPrevCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevCursor.Get(), o.PrevCursor.IsSet()
}

// HasPrevCursor returns a boolean if a field has been set.
func (o *GlobalOrgsMetaPage) HasPrevCursor() bool {
	return o != nil && o.PrevCursor.IsSet()
}

// SetPrevCursor gets a reference to the given datadog.NullableString and assigns it to the PrevCursor field.
func (o *GlobalOrgsMetaPage) SetPrevCursor(v string) {
	o.PrevCursor.Set(&v)
}

// SetPrevCursorNil sets the value for PrevCursor to be an explicit nil.
func (o *GlobalOrgsMetaPage) SetPrevCursorNil() {
	o.PrevCursor.Set(nil)
}

// UnsetPrevCursor ensures that no value is present for PrevCursor, not even an explicit nil.
func (o *GlobalOrgsMetaPage) UnsetPrevCursor() {
	o.PrevCursor.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GlobalOrgsMetaPage) GetType() GlobalOrgsMetaPageType {
	if o == nil || o.Type == nil {
		var ret GlobalOrgsMetaPageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalOrgsMetaPage) GetTypeOk() (*GlobalOrgsMetaPageType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GlobalOrgsMetaPage) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given GlobalOrgsMetaPageType and assigns it to the Type field.
func (o *GlobalOrgsMetaPage) SetType(v GlobalOrgsMetaPageType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GlobalOrgsMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextCursor.IsSet() {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if o.PrevCursor.IsSet() {
		toSerialize["prev_cursor"] = o.PrevCursor.Get()
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
func (o *GlobalOrgsMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cursor     *string                 `json:"cursor,omitempty"`
		Limit      *int32                  `json:"limit,omitempty"`
		NextCursor datadog.NullableString  `json:"next_cursor,omitempty"`
		PrevCursor datadog.NullableString  `json:"prev_cursor,omitempty"`
		Type       *GlobalOrgsMetaPageType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cursor", "limit", "next_cursor", "prev_cursor", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Cursor = all.Cursor
	o.Limit = all.Limit
	o.NextCursor = all.NextCursor
	o.PrevCursor = all.PrevCursor
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
