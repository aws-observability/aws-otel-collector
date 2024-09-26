// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsResponseMetaPagination Teams response metadata.
type TeamsResponseMetaPagination struct {
	// The first offset.
	FirstOffset *int64 `json:"first_offset,omitempty"`
	// The last offset.
	LastOffset *int64 `json:"last_offset,omitempty"`
	// Pagination limit.
	Limit *int64 `json:"limit,omitempty"`
	// The next offset.
	NextOffset *int64 `json:"next_offset,omitempty"`
	// The offset.
	Offset *int64 `json:"offset,omitempty"`
	// The previous offset.
	PrevOffset *int64 `json:"prev_offset,omitempty"`
	// Total results.
	Total *int64 `json:"total,omitempty"`
	// Offset type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsResponseMetaPagination instantiates a new TeamsResponseMetaPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsResponseMetaPagination() *TeamsResponseMetaPagination {
	this := TeamsResponseMetaPagination{}
	return &this
}

// NewTeamsResponseMetaPaginationWithDefaults instantiates a new TeamsResponseMetaPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsResponseMetaPaginationWithDefaults() *TeamsResponseMetaPagination {
	this := TeamsResponseMetaPagination{}
	return &this
}

// GetFirstOffset returns the FirstOffset field value if set, zero value otherwise.
func (o *TeamsResponseMetaPagination) GetFirstOffset() int64 {
	if o == nil || o.FirstOffset == nil {
		var ret int64
		return ret
	}
	return *o.FirstOffset
}

// GetFirstOffsetOk returns a tuple with the FirstOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponseMetaPagination) GetFirstOffsetOk() (*int64, bool) {
	if o == nil || o.FirstOffset == nil {
		return nil, false
	}
	return o.FirstOffset, true
}

// HasFirstOffset returns a boolean if a field has been set.
func (o *TeamsResponseMetaPagination) HasFirstOffset() bool {
	return o != nil && o.FirstOffset != nil
}

// SetFirstOffset gets a reference to the given int64 and assigns it to the FirstOffset field.
func (o *TeamsResponseMetaPagination) SetFirstOffset(v int64) {
	o.FirstOffset = &v
}

// GetLastOffset returns the LastOffset field value if set, zero value otherwise.
func (o *TeamsResponseMetaPagination) GetLastOffset() int64 {
	if o == nil || o.LastOffset == nil {
		var ret int64
		return ret
	}
	return *o.LastOffset
}

// GetLastOffsetOk returns a tuple with the LastOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponseMetaPagination) GetLastOffsetOk() (*int64, bool) {
	if o == nil || o.LastOffset == nil {
		return nil, false
	}
	return o.LastOffset, true
}

// HasLastOffset returns a boolean if a field has been set.
func (o *TeamsResponseMetaPagination) HasLastOffset() bool {
	return o != nil && o.LastOffset != nil
}

// SetLastOffset gets a reference to the given int64 and assigns it to the LastOffset field.
func (o *TeamsResponseMetaPagination) SetLastOffset(v int64) {
	o.LastOffset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TeamsResponseMetaPagination) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponseMetaPagination) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TeamsResponseMetaPagination) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *TeamsResponseMetaPagination) SetLimit(v int64) {
	o.Limit = &v
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise.
func (o *TeamsResponseMetaPagination) GetNextOffset() int64 {
	if o == nil || o.NextOffset == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponseMetaPagination) GetNextOffsetOk() (*int64, bool) {
	if o == nil || o.NextOffset == nil {
		return nil, false
	}
	return o.NextOffset, true
}

// HasNextOffset returns a boolean if a field has been set.
func (o *TeamsResponseMetaPagination) HasNextOffset() bool {
	return o != nil && o.NextOffset != nil
}

// SetNextOffset gets a reference to the given int64 and assigns it to the NextOffset field.
func (o *TeamsResponseMetaPagination) SetNextOffset(v int64) {
	o.NextOffset = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TeamsResponseMetaPagination) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponseMetaPagination) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TeamsResponseMetaPagination) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *TeamsResponseMetaPagination) SetOffset(v int64) {
	o.Offset = &v
}

// GetPrevOffset returns the PrevOffset field value if set, zero value otherwise.
func (o *TeamsResponseMetaPagination) GetPrevOffset() int64 {
	if o == nil || o.PrevOffset == nil {
		var ret int64
		return ret
	}
	return *o.PrevOffset
}

// GetPrevOffsetOk returns a tuple with the PrevOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponseMetaPagination) GetPrevOffsetOk() (*int64, bool) {
	if o == nil || o.PrevOffset == nil {
		return nil, false
	}
	return o.PrevOffset, true
}

// HasPrevOffset returns a boolean if a field has been set.
func (o *TeamsResponseMetaPagination) HasPrevOffset() bool {
	return o != nil && o.PrevOffset != nil
}

// SetPrevOffset gets a reference to the given int64 and assigns it to the PrevOffset field.
func (o *TeamsResponseMetaPagination) SetPrevOffset(v int64) {
	o.PrevOffset = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TeamsResponseMetaPagination) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponseMetaPagination) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TeamsResponseMetaPagination) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *TeamsResponseMetaPagination) SetTotal(v int64) {
	o.Total = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TeamsResponseMetaPagination) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponseMetaPagination) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TeamsResponseMetaPagination) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TeamsResponseMetaPagination) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsResponseMetaPagination) MarshalJSON() ([]byte, error) {
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
func (o *TeamsResponseMetaPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstOffset *int64  `json:"first_offset,omitempty"`
		LastOffset  *int64  `json:"last_offset,omitempty"`
		Limit       *int64  `json:"limit,omitempty"`
		NextOffset  *int64  `json:"next_offset,omitempty"`
		Offset      *int64  `json:"offset,omitempty"`
		PrevOffset  *int64  `json:"prev_offset,omitempty"`
		Total       *int64  `json:"total,omitempty"`
		Type        *string `json:"type,omitempty"`
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
