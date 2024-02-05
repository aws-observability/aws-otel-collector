// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingMute Information about the mute status of this finding.
type FindingMute struct {
	// Additional information about the reason why this finding is muted or unmuted.
	Description *string `json:"description,omitempty"`
	// The expiration date of the mute or unmute action (Unix ms).
	ExpirationDate *int64 `json:"expiration_date,omitempty"`
	// Whether this finding is muted or unmuted.
	Muted *bool `json:"muted,omitempty"`
	// The reason why this finding is muted or unmuted.
	Reason *FindingMuteReason `json:"reason,omitempty"`
	// The start of the mute period.
	StartDate *int64 `json:"start_date,omitempty"`
	// The ID of the user who muted or unmuted this finding.
	Uuid *string `json:"uuid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFindingMute instantiates a new FindingMute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFindingMute() *FindingMute {
	this := FindingMute{}
	return &this
}

// NewFindingMuteWithDefaults instantiates a new FindingMute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFindingMuteWithDefaults() *FindingMute {
	this := FindingMute{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FindingMute) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingMute) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FindingMute) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FindingMute) SetDescription(v string) {
	o.Description = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *FindingMute) GetExpirationDate() int64 {
	if o == nil || o.ExpirationDate == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingMute) GetExpirationDateOk() (*int64, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *FindingMute) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *FindingMute) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetMuted returns the Muted field value if set, zero value otherwise.
func (o *FindingMute) GetMuted() bool {
	if o == nil || o.Muted == nil {
		var ret bool
		return ret
	}
	return *o.Muted
}

// GetMutedOk returns a tuple with the Muted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingMute) GetMutedOk() (*bool, bool) {
	if o == nil || o.Muted == nil {
		return nil, false
	}
	return o.Muted, true
}

// HasMuted returns a boolean if a field has been set.
func (o *FindingMute) HasMuted() bool {
	return o != nil && o.Muted != nil
}

// SetMuted gets a reference to the given bool and assigns it to the Muted field.
func (o *FindingMute) SetMuted(v bool) {
	o.Muted = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *FindingMute) GetReason() FindingMuteReason {
	if o == nil || o.Reason == nil {
		var ret FindingMuteReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingMute) GetReasonOk() (*FindingMuteReason, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *FindingMute) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given FindingMuteReason and assigns it to the Reason field.
func (o *FindingMute) SetReason(v FindingMuteReason) {
	o.Reason = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *FindingMute) GetStartDate() int64 {
	if o == nil || o.StartDate == nil {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingMute) GetStartDateOk() (*int64, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *FindingMute) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *FindingMute) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *FindingMute) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingMute) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *FindingMute) HasUuid() bool {
	return o != nil && o.Uuid != nil
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *FindingMute) SetUuid(v string) {
	o.Uuid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FindingMute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if o.Muted != nil {
		toSerialize["muted"] = o.Muted
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FindingMute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description    *string            `json:"description,omitempty"`
		ExpirationDate *int64             `json:"expiration_date,omitempty"`
		Muted          *bool              `json:"muted,omitempty"`
		Reason         *FindingMuteReason `json:"reason,omitempty"`
		StartDate      *int64             `json:"start_date,omitempty"`
		Uuid           *string            `json:"uuid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.Description = all.Description
	o.ExpirationDate = all.ExpirationDate
	o.Muted = all.Muted
	if all.Reason != nil && !all.Reason.IsValid() {
		hasInvalidField = true
	} else {
		o.Reason = all.Reason
	}
	o.StartDate = all.StartDate
	o.Uuid = all.Uuid

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
