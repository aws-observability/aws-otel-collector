// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerMetaPage Paging attributes.
type ContainerMetaPage struct {
	// The cursor used to get the current results, if any.
	Cursor *string `json:"cursor,omitempty"`
	// Number of results returned
	Limit *int32 `json:"limit,omitempty"`
	// The cursor used to get the next results, if any.
	NextCursor *string `json:"next_cursor,omitempty"`
	// The cursor used to get the previous results, if any.
	PrevCursor datadog.NullableString `json:"prev_cursor,omitempty"`
	// Total number of records that match the query.
	Total *int64 `json:"total,omitempty"`
	// Type of Container pagination.
	Type *ContainerMetaPageType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerMetaPage instantiates a new ContainerMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerMetaPage() *ContainerMetaPage {
	this := ContainerMetaPage{}
	var typeVar ContainerMetaPageType = CONTAINERMETAPAGETYPE_CURSOR_LIMIT
	this.Type = &typeVar
	return &this
}

// NewContainerMetaPageWithDefaults instantiates a new ContainerMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerMetaPageWithDefaults() *ContainerMetaPage {
	this := ContainerMetaPage{}
	var typeVar ContainerMetaPageType = CONTAINERMETAPAGETYPE_CURSOR_LIMIT
	this.Type = &typeVar
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ContainerMetaPage) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerMetaPage) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ContainerMetaPage) HasCursor() bool {
	return o != nil && o.Cursor != nil
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ContainerMetaPage) SetCursor(v string) {
	o.Cursor = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ContainerMetaPage) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerMetaPage) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ContainerMetaPage) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ContainerMetaPage) SetLimit(v int32) {
	o.Limit = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *ContainerMetaPage) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerMetaPage) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *ContainerMetaPage) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *ContainerMetaPage) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetPrevCursor returns the PrevCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerMetaPage) GetPrevCursor() string {
	if o == nil || o.PrevCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrevCursor.Get()
}

// GetPrevCursorOk returns a tuple with the PrevCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ContainerMetaPage) GetPrevCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevCursor.Get(), o.PrevCursor.IsSet()
}

// HasPrevCursor returns a boolean if a field has been set.
func (o *ContainerMetaPage) HasPrevCursor() bool {
	return o != nil && o.PrevCursor.IsSet()
}

// SetPrevCursor gets a reference to the given datadog.NullableString and assigns it to the PrevCursor field.
func (o *ContainerMetaPage) SetPrevCursor(v string) {
	o.PrevCursor.Set(&v)
}

// SetPrevCursorNil sets the value for PrevCursor to be an explicit nil.
func (o *ContainerMetaPage) SetPrevCursorNil() {
	o.PrevCursor.Set(nil)
}

// UnsetPrevCursor ensures that no value is present for PrevCursor, not even an explicit nil.
func (o *ContainerMetaPage) UnsetPrevCursor() {
	o.PrevCursor.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ContainerMetaPage) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerMetaPage) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ContainerMetaPage) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *ContainerMetaPage) SetTotal(v int64) {
	o.Total = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContainerMetaPage) GetType() ContainerMetaPageType {
	if o == nil || o.Type == nil {
		var ret ContainerMetaPageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerMetaPage) GetTypeOk() (*ContainerMetaPageType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContainerMetaPage) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ContainerMetaPageType and assigns it to the Type field.
func (o *ContainerMetaPage) SetType(v ContainerMetaPageType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerMetaPage) MarshalJSON() ([]byte, error) {
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
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if o.PrevCursor.IsSet() {
		toSerialize["prev_cursor"] = o.PrevCursor.Get()
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
func (o *ContainerMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cursor     *string                `json:"cursor,omitempty"`
		Limit      *int32                 `json:"limit,omitempty"`
		NextCursor *string                `json:"next_cursor,omitempty"`
		PrevCursor datadog.NullableString `json:"prev_cursor,omitempty"`
		Total      *int64                 `json:"total,omitempty"`
		Type       *ContainerMetaPageType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cursor", "limit", "next_cursor", "prev_cursor", "total", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Cursor = all.Cursor
	o.Limit = all.Limit
	o.NextCursor = all.NextCursor
	o.PrevCursor = all.PrevCursor
	o.Total = all.Total
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
