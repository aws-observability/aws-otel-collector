// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamNotificationRulesResponseMetaPage Metadata related to paging information that is included in the response when querying the team notification rules
type TeamNotificationRulesResponseMetaPage struct {
	// The first offset.
	FirstOffset *int64 `json:"first_offset,omitempty"`
	// The last offset.
	LastOffset *int64 `json:"last_offset,omitempty"`
	// Pagination limit.
	Limit *int64 `json:"limit,omitempty"`
	// The next offset.
	NextOffset datadog.NullableInt64 `json:"next_offset,omitempty"`
	// The offset.
	Offset *int64 `json:"offset,omitempty"`
	// The previous offset.
	PrevOffset datadog.NullableInt64 `json:"prev_offset,omitempty"`
	// Total results.
	Total *int64 `json:"total,omitempty"`
	// Offset type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamNotificationRulesResponseMetaPage instantiates a new TeamNotificationRulesResponseMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamNotificationRulesResponseMetaPage() *TeamNotificationRulesResponseMetaPage {
	this := TeamNotificationRulesResponseMetaPage{}
	return &this
}

// NewTeamNotificationRulesResponseMetaPageWithDefaults instantiates a new TeamNotificationRulesResponseMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamNotificationRulesResponseMetaPageWithDefaults() *TeamNotificationRulesResponseMetaPage {
	this := TeamNotificationRulesResponseMetaPage{}
	return &this
}

// GetFirstOffset returns the FirstOffset field value if set, zero value otherwise.
func (o *TeamNotificationRulesResponseMetaPage) GetFirstOffset() int64 {
	if o == nil || o.FirstOffset == nil {
		var ret int64
		return ret
	}
	return *o.FirstOffset
}

// GetFirstOffsetOk returns a tuple with the FirstOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRulesResponseMetaPage) GetFirstOffsetOk() (*int64, bool) {
	if o == nil || o.FirstOffset == nil {
		return nil, false
	}
	return o.FirstOffset, true
}

// HasFirstOffset returns a boolean if a field has been set.
func (o *TeamNotificationRulesResponseMetaPage) HasFirstOffset() bool {
	return o != nil && o.FirstOffset != nil
}

// SetFirstOffset gets a reference to the given int64 and assigns it to the FirstOffset field.
func (o *TeamNotificationRulesResponseMetaPage) SetFirstOffset(v int64) {
	o.FirstOffset = &v
}

// GetLastOffset returns the LastOffset field value if set, zero value otherwise.
func (o *TeamNotificationRulesResponseMetaPage) GetLastOffset() int64 {
	if o == nil || o.LastOffset == nil {
		var ret int64
		return ret
	}
	return *o.LastOffset
}

// GetLastOffsetOk returns a tuple with the LastOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRulesResponseMetaPage) GetLastOffsetOk() (*int64, bool) {
	if o == nil || o.LastOffset == nil {
		return nil, false
	}
	return o.LastOffset, true
}

// HasLastOffset returns a boolean if a field has been set.
func (o *TeamNotificationRulesResponseMetaPage) HasLastOffset() bool {
	return o != nil && o.LastOffset != nil
}

// SetLastOffset gets a reference to the given int64 and assigns it to the LastOffset field.
func (o *TeamNotificationRulesResponseMetaPage) SetLastOffset(v int64) {
	o.LastOffset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TeamNotificationRulesResponseMetaPage) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRulesResponseMetaPage) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TeamNotificationRulesResponseMetaPage) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *TeamNotificationRulesResponseMetaPage) SetLimit(v int64) {
	o.Limit = &v
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamNotificationRulesResponseMetaPage) GetNextOffset() int64 {
	if o == nil || o.NextOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset.Get()
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamNotificationRulesResponseMetaPage) GetNextOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextOffset.Get(), o.NextOffset.IsSet()
}

// HasNextOffset returns a boolean if a field has been set.
func (o *TeamNotificationRulesResponseMetaPage) HasNextOffset() bool {
	return o != nil && o.NextOffset.IsSet()
}

// SetNextOffset gets a reference to the given datadog.NullableInt64 and assigns it to the NextOffset field.
func (o *TeamNotificationRulesResponseMetaPage) SetNextOffset(v int64) {
	o.NextOffset.Set(&v)
}

// SetNextOffsetNil sets the value for NextOffset to be an explicit nil.
func (o *TeamNotificationRulesResponseMetaPage) SetNextOffsetNil() {
	o.NextOffset.Set(nil)
}

// UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil.
func (o *TeamNotificationRulesResponseMetaPage) UnsetNextOffset() {
	o.NextOffset.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TeamNotificationRulesResponseMetaPage) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRulesResponseMetaPage) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TeamNotificationRulesResponseMetaPage) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *TeamNotificationRulesResponseMetaPage) SetOffset(v int64) {
	o.Offset = &v
}

// GetPrevOffset returns the PrevOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamNotificationRulesResponseMetaPage) GetPrevOffset() int64 {
	if o == nil || o.PrevOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrevOffset.Get()
}

// GetPrevOffsetOk returns a tuple with the PrevOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamNotificationRulesResponseMetaPage) GetPrevOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevOffset.Get(), o.PrevOffset.IsSet()
}

// HasPrevOffset returns a boolean if a field has been set.
func (o *TeamNotificationRulesResponseMetaPage) HasPrevOffset() bool {
	return o != nil && o.PrevOffset.IsSet()
}

// SetPrevOffset gets a reference to the given datadog.NullableInt64 and assigns it to the PrevOffset field.
func (o *TeamNotificationRulesResponseMetaPage) SetPrevOffset(v int64) {
	o.PrevOffset.Set(&v)
}

// SetPrevOffsetNil sets the value for PrevOffset to be an explicit nil.
func (o *TeamNotificationRulesResponseMetaPage) SetPrevOffsetNil() {
	o.PrevOffset.Set(nil)
}

// UnsetPrevOffset ensures that no value is present for PrevOffset, not even an explicit nil.
func (o *TeamNotificationRulesResponseMetaPage) UnsetPrevOffset() {
	o.PrevOffset.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TeamNotificationRulesResponseMetaPage) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRulesResponseMetaPage) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TeamNotificationRulesResponseMetaPage) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *TeamNotificationRulesResponseMetaPage) SetTotal(v int64) {
	o.Total = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TeamNotificationRulesResponseMetaPage) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRulesResponseMetaPage) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TeamNotificationRulesResponseMetaPage) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TeamNotificationRulesResponseMetaPage) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamNotificationRulesResponseMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FirstOffset != nil {
		toSerialize["first_offset"] = o.FirstOffset
	}
	if o.LastOffset != nil {
		toSerialize["last_offset"] = o.LastOffset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextOffset.IsSet() {
		toSerialize["next_offset"] = o.NextOffset.Get()
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.PrevOffset.IsSet() {
		toSerialize["prev_offset"] = o.PrevOffset.Get()
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
func (o *TeamNotificationRulesResponseMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstOffset *int64                `json:"first_offset,omitempty"`
		LastOffset  *int64                `json:"last_offset,omitempty"`
		Limit       *int64                `json:"limit,omitempty"`
		NextOffset  datadog.NullableInt64 `json:"next_offset,omitempty"`
		Offset      *int64                `json:"offset,omitempty"`
		PrevOffset  datadog.NullableInt64 `json:"prev_offset,omitempty"`
		Total       *int64                `json:"total,omitempty"`
		Type        *string               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_offset", "last_offset", "limit", "next_offset", "offset", "prev_offset", "total", "type"})
	} else {
		return err
	}
	o.FirstOffset = all.FirstOffset
	o.LastOffset = all.LastOffset
	o.Limit = all.Limit
	o.NextOffset = all.NextOffset
	o.Offset = all.Offset
	o.PrevOffset = all.PrevOffset
	o.Total = all.Total
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
