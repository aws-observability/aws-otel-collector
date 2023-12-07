// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// MuteFindingResponseProperties Information about the mute status of this finding.
type MuteFindingResponseProperties struct {
	// Additional information about the reason why this finding is muted or unmuted.
	// This attribute will not be included in the response if the description is not provided in the request body.
	//
	Description *string `json:"description,omitempty"`
	// The expiration date of the mute or unmute action.
	// If the expiration date is not provided in the request body, this attribute will not be included in the response and the finding will be muted or unmuted indefinitely.
	//
	ExpirationDate *int64 `json:"expiration_date,omitempty"`
	// Whether this finding is muted or unmuted.
	Muted *bool `json:"muted,omitempty"`
	// The reason why this finding is muted or unmuted.
	Reason *FindingMuteReason `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMuteFindingResponseProperties instantiates a new MuteFindingResponseProperties object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMuteFindingResponseProperties() *MuteFindingResponseProperties {
	this := MuteFindingResponseProperties{}
	return &this
}

// NewMuteFindingResponsePropertiesWithDefaults instantiates a new MuteFindingResponseProperties object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMuteFindingResponsePropertiesWithDefaults() *MuteFindingResponseProperties {
	this := MuteFindingResponseProperties{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MuteFindingResponseProperties) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MuteFindingResponseProperties) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MuteFindingResponseProperties) SetDescription(v string) {
	o.Description = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *MuteFindingResponseProperties) GetExpirationDate() int64 {
	if o == nil || o.ExpirationDate == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseProperties) GetExpirationDateOk() (*int64, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *MuteFindingResponseProperties) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *MuteFindingResponseProperties) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetMuted returns the Muted field value if set, zero value otherwise.
func (o *MuteFindingResponseProperties) GetMuted() bool {
	if o == nil || o.Muted == nil {
		var ret bool
		return ret
	}
	return *o.Muted
}

// GetMutedOk returns a tuple with the Muted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseProperties) GetMutedOk() (*bool, bool) {
	if o == nil || o.Muted == nil {
		return nil, false
	}
	return o.Muted, true
}

// HasMuted returns a boolean if a field has been set.
func (o *MuteFindingResponseProperties) HasMuted() bool {
	return o != nil && o.Muted != nil
}

// SetMuted gets a reference to the given bool and assigns it to the Muted field.
func (o *MuteFindingResponseProperties) SetMuted(v bool) {
	o.Muted = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MuteFindingResponseProperties) GetReason() FindingMuteReason {
	if o == nil || o.Reason == nil {
		var ret FindingMuteReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseProperties) GetReasonOk() (*FindingMuteReason, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MuteFindingResponseProperties) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given FindingMuteReason and assigns it to the Reason field.
func (o *MuteFindingResponseProperties) SetReason(v FindingMuteReason) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MuteFindingResponseProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
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
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MuteFindingResponseProperties) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description    *string            `json:"description,omitempty"`
		ExpirationDate *int64             `json:"expiration_date,omitempty"`
		Muted          *bool              `json:"muted,omitempty"`
		Reason         *FindingMuteReason `json:"reason,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
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

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
