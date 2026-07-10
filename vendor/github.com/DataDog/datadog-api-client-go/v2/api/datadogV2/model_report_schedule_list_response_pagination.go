// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleListResponsePagination Offset and limit pagination metadata for a report schedule list response.
type ReportScheduleListResponsePagination struct {
	// The first offset.
	FirstOffset *int64 `json:"first_offset,omitempty"`
	// The last offset when the total count is known, or `null` if it is unavailable.
	LastOffset datadog.NullableInt64 `json:"last_offset,omitempty"`
	// The maximum number of schedules returned.
	Limit *int64 `json:"limit,omitempty"`
	// The next offset.
	NextOffset *int64 `json:"next_offset,omitempty"`
	// The current offset.
	Offset *int64 `json:"offset,omitempty"`
	// The previous offset.
	PrevOffset *int64 `json:"prev_offset,omitempty"`
	// The total number of matching schedules.
	Total *int64 `json:"total,omitempty"`
	// The pagination type.
	Type *ReportScheduleListResponsePaginationType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleListResponsePagination instantiates a new ReportScheduleListResponsePagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleListResponsePagination() *ReportScheduleListResponsePagination {
	this := ReportScheduleListResponsePagination{}
	return &this
}

// NewReportScheduleListResponsePaginationWithDefaults instantiates a new ReportScheduleListResponsePagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleListResponsePaginationWithDefaults() *ReportScheduleListResponsePagination {
	this := ReportScheduleListResponsePagination{}
	return &this
}

// GetFirstOffset returns the FirstOffset field value if set, zero value otherwise.
func (o *ReportScheduleListResponsePagination) GetFirstOffset() int64 {
	if o == nil || o.FirstOffset == nil {
		var ret int64
		return ret
	}
	return *o.FirstOffset
}

// GetFirstOffsetOk returns a tuple with the FirstOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponsePagination) GetFirstOffsetOk() (*int64, bool) {
	if o == nil || o.FirstOffset == nil {
		return nil, false
	}
	return o.FirstOffset, true
}

// HasFirstOffset returns a boolean if a field has been set.
func (o *ReportScheduleListResponsePagination) HasFirstOffset() bool {
	return o != nil && o.FirstOffset != nil
}

// SetFirstOffset gets a reference to the given int64 and assigns it to the FirstOffset field.
func (o *ReportScheduleListResponsePagination) SetFirstOffset(v int64) {
	o.FirstOffset = &v
}

// GetLastOffset returns the LastOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleListResponsePagination) GetLastOffset() int64 {
	if o == nil || o.LastOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastOffset.Get()
}

// GetLastOffsetOk returns a tuple with the LastOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleListResponsePagination) GetLastOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastOffset.Get(), o.LastOffset.IsSet()
}

// HasLastOffset returns a boolean if a field has been set.
func (o *ReportScheduleListResponsePagination) HasLastOffset() bool {
	return o != nil && o.LastOffset.IsSet()
}

// SetLastOffset gets a reference to the given datadog.NullableInt64 and assigns it to the LastOffset field.
func (o *ReportScheduleListResponsePagination) SetLastOffset(v int64) {
	o.LastOffset.Set(&v)
}

// SetLastOffsetNil sets the value for LastOffset to be an explicit nil.
func (o *ReportScheduleListResponsePagination) SetLastOffsetNil() {
	o.LastOffset.Set(nil)
}

// UnsetLastOffset ensures that no value is present for LastOffset, not even an explicit nil.
func (o *ReportScheduleListResponsePagination) UnsetLastOffset() {
	o.LastOffset.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ReportScheduleListResponsePagination) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponsePagination) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ReportScheduleListResponsePagination) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ReportScheduleListResponsePagination) SetLimit(v int64) {
	o.Limit = &v
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise.
func (o *ReportScheduleListResponsePagination) GetNextOffset() int64 {
	if o == nil || o.NextOffset == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponsePagination) GetNextOffsetOk() (*int64, bool) {
	if o == nil || o.NextOffset == nil {
		return nil, false
	}
	return o.NextOffset, true
}

// HasNextOffset returns a boolean if a field has been set.
func (o *ReportScheduleListResponsePagination) HasNextOffset() bool {
	return o != nil && o.NextOffset != nil
}

// SetNextOffset gets a reference to the given int64 and assigns it to the NextOffset field.
func (o *ReportScheduleListResponsePagination) SetNextOffset(v int64) {
	o.NextOffset = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ReportScheduleListResponsePagination) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponsePagination) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ReportScheduleListResponsePagination) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ReportScheduleListResponsePagination) SetOffset(v int64) {
	o.Offset = &v
}

// GetPrevOffset returns the PrevOffset field value if set, zero value otherwise.
func (o *ReportScheduleListResponsePagination) GetPrevOffset() int64 {
	if o == nil || o.PrevOffset == nil {
		var ret int64
		return ret
	}
	return *o.PrevOffset
}

// GetPrevOffsetOk returns a tuple with the PrevOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponsePagination) GetPrevOffsetOk() (*int64, bool) {
	if o == nil || o.PrevOffset == nil {
		return nil, false
	}
	return o.PrevOffset, true
}

// HasPrevOffset returns a boolean if a field has been set.
func (o *ReportScheduleListResponsePagination) HasPrevOffset() bool {
	return o != nil && o.PrevOffset != nil
}

// SetPrevOffset gets a reference to the given int64 and assigns it to the PrevOffset field.
func (o *ReportScheduleListResponsePagination) SetPrevOffset(v int64) {
	o.PrevOffset = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ReportScheduleListResponsePagination) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponsePagination) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ReportScheduleListResponsePagination) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *ReportScheduleListResponsePagination) SetTotal(v int64) {
	o.Total = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReportScheduleListResponsePagination) GetType() ReportScheduleListResponsePaginationType {
	if o == nil || o.Type == nil {
		var ret ReportScheduleListResponsePaginationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponsePagination) GetTypeOk() (*ReportScheduleListResponsePaginationType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReportScheduleListResponsePagination) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ReportScheduleListResponsePaginationType and assigns it to the Type field.
func (o *ReportScheduleListResponsePagination) SetType(v ReportScheduleListResponsePaginationType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleListResponsePagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FirstOffset != nil {
		toSerialize["first_offset"] = o.FirstOffset
	}
	if o.LastOffset.IsSet() {
		toSerialize["last_offset"] = o.LastOffset.Get()
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextOffset != nil {
		toSerialize["next_offset"] = o.NextOffset
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.PrevOffset != nil {
		toSerialize["prev_offset"] = o.PrevOffset
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
func (o *ReportScheduleListResponsePagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstOffset *int64                                    `json:"first_offset,omitempty"`
		LastOffset  datadog.NullableInt64                     `json:"last_offset,omitempty"`
		Limit       *int64                                    `json:"limit,omitempty"`
		NextOffset  *int64                                    `json:"next_offset,omitempty"`
		Offset      *int64                                    `json:"offset,omitempty"`
		PrevOffset  *int64                                    `json:"prev_offset,omitempty"`
		Total       *int64                                    `json:"total,omitempty"`
		Type        *ReportScheduleListResponsePaginationType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_offset", "last_offset", "limit", "next_offset", "offset", "prev_offset", "total", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FirstOffset = all.FirstOffset
	o.LastOffset = all.LastOffset
	o.Limit = all.Limit
	o.NextOffset = all.NextOffset
	o.Offset = all.Offset
	o.PrevOffset = all.PrevOffset
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
