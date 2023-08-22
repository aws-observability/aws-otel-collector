// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// MuteFindingRequestProperties Object containing the new mute properties of the finding.
type MuteFindingRequestProperties struct {
	// Additional information about the reason why this finding is muted or unmuted. This field has a maximum limit of 280 characters.
	Description *string `json:"description,omitempty"`
	// The expiration date of the mute or unmute action (Unix ms). It must be set to a value greater than the current timestamp.
	// If this field is not provided, the finding will be muted or unmuted indefinitely, which is equivalent to setting the expiration date to 9999999999999.
	//
	ExpirationDate *int64 `json:"expiration_date,omitempty"`
	// Whether this finding is muted or unmuted.
	Muted bool `json:"muted"`
	// The reason why this finding is muted or unmuted.
	Reason FindingMuteReason `json:"reason"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMuteFindingRequestProperties instantiates a new MuteFindingRequestProperties object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMuteFindingRequestProperties(muted bool, reason FindingMuteReason) *MuteFindingRequestProperties {
	this := MuteFindingRequestProperties{}
	this.Muted = muted
	this.Reason = reason
	return &this
}

// NewMuteFindingRequestPropertiesWithDefaults instantiates a new MuteFindingRequestProperties object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMuteFindingRequestPropertiesWithDefaults() *MuteFindingRequestProperties {
	this := MuteFindingRequestProperties{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MuteFindingRequestProperties) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingRequestProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MuteFindingRequestProperties) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MuteFindingRequestProperties) SetDescription(v string) {
	o.Description = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *MuteFindingRequestProperties) GetExpirationDate() int64 {
	if o == nil || o.ExpirationDate == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingRequestProperties) GetExpirationDateOk() (*int64, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *MuteFindingRequestProperties) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *MuteFindingRequestProperties) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetMuted returns the Muted field value.
func (o *MuteFindingRequestProperties) GetMuted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Muted
}

// GetMutedOk returns a tuple with the Muted field value
// and a boolean to check if the value has been set.
func (o *MuteFindingRequestProperties) GetMutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Muted, true
}

// SetMuted sets field value.
func (o *MuteFindingRequestProperties) SetMuted(v bool) {
	o.Muted = v
}

// GetReason returns the Reason field value.
func (o *MuteFindingRequestProperties) GetReason() FindingMuteReason {
	if o == nil {
		var ret FindingMuteReason
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *MuteFindingRequestProperties) GetReasonOk() (*FindingMuteReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *MuteFindingRequestProperties) SetReason(v FindingMuteReason) {
	o.Reason = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MuteFindingRequestProperties) MarshalJSON() ([]byte, error) {
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
	toSerialize["muted"] = o.Muted
	toSerialize["reason"] = o.Reason
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MuteFindingRequestProperties) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Description    *string            `json:"description,omitempty"`
		ExpirationDate *int64             `json:"expiration_date,omitempty"`
		Muted          *bool              `json:"muted"`
		Reason         *FindingMuteReason `json:"reason"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Muted == nil {
		return fmt.Errorf("required field muted missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if v := all.Reason; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Description = all.Description
	o.ExpirationDate = all.ExpirationDate
	o.Muted = *all.Muted
	o.Reason = *all.Reason

	return nil
}
