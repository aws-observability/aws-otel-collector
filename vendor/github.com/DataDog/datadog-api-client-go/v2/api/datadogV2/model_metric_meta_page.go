// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricMetaPage Paging attributes. Only present if pagination query parameters were provided.
type MetricMetaPage struct {
	// The cursor used to get the current results, if any.
	Cursor datadog.NullableString `json:"cursor,omitempty"`
	// Number of results returned
	Limit *int32 `json:"limit,omitempty"`
	// The cursor used to get the next results, if any.
	NextCursor datadog.NullableString `json:"next_cursor,omitempty"`
	// Type of metric pagination.
	Type *MetricMetaPageType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricMetaPage instantiates a new MetricMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricMetaPage() *MetricMetaPage {
	this := MetricMetaPage{}
	var typeVar MetricMetaPageType = METRICMETAPAGETYPE_CURSOR_LIMIT
	this.Type = &typeVar
	return &this
}

// NewMetricMetaPageWithDefaults instantiates a new MetricMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricMetaPageWithDefaults() *MetricMetaPage {
	this := MetricMetaPage{}
	var typeVar MetricMetaPageType = METRICMETAPAGETYPE_CURSOR_LIMIT
	this.Type = &typeVar
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricMetaPage) GetCursor() string {
	if o == nil || o.Cursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetaPage) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *MetricMetaPage) HasCursor() bool {
	return o != nil && o.Cursor.IsSet()
}

// SetCursor gets a reference to the given datadog.NullableString and assigns it to the Cursor field.
func (o *MetricMetaPage) SetCursor(v string) {
	o.Cursor.Set(&v)
}

// SetCursorNil sets the value for Cursor to be an explicit nil.
func (o *MetricMetaPage) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil.
func (o *MetricMetaPage) UnsetCursor() {
	o.Cursor.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *MetricMetaPage) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetaPage) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *MetricMetaPage) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *MetricMetaPage) SetLimit(v int32) {
	o.Limit = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricMetaPage) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetaPage) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *MetricMetaPage) HasNextCursor() bool {
	return o != nil && o.NextCursor.IsSet()
}

// SetNextCursor gets a reference to the given datadog.NullableString and assigns it to the NextCursor field.
func (o *MetricMetaPage) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// SetNextCursorNil sets the value for NextCursor to be an explicit nil.
func (o *MetricMetaPage) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil.
func (o *MetricMetaPage) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetricMetaPage) GetType() MetricMetaPageType {
	if o == nil || o.Type == nil {
		var ret MetricMetaPageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetaPage) GetTypeOk() (*MetricMetaPageType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetricMetaPage) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given MetricMetaPageType and assigns it to the Type field.
func (o *MetricMetaPage) SetType(v MetricMetaPageType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextCursor.IsSet() {
		toSerialize["next_cursor"] = o.NextCursor.Get()
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
func (o *MetricMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cursor     datadog.NullableString `json:"cursor,omitempty"`
		Limit      *int32                 `json:"limit,omitempty"`
		NextCursor datadog.NullableString `json:"next_cursor,omitempty"`
		Type       *MetricMetaPageType    `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cursor", "limit", "next_cursor", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Cursor = all.Cursor
	o.Limit = all.Limit
	o.NextCursor = all.NextCursor
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
