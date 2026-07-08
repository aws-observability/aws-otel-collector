// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MuteFindingsMuteAttributes Mute properties to apply to the findings.
type MuteFindingsMuteAttributes struct {
	// Additional information about the reason why the findings are muted or unmuted. This field has a limit of 280 characters.
	Description *string `json:"description,omitempty"`
	// The expiration date of the mute action (Unix ms). It must be set to a value greater than the current timestamp. If this field is not provided, the findings remain muted indefinitely.
	ExpireAt *int64 `json:"expire_at,omitempty"`
	// Whether the findings should be muted or unmuted.
	IsMuted bool `json:"is_muted"`
	// The reason why the findings are muted or unmuted.
	Reason MuteFindingsReason `json:"reason"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMuteFindingsMuteAttributes instantiates a new MuteFindingsMuteAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMuteFindingsMuteAttributes(isMuted bool, reason MuteFindingsReason) *MuteFindingsMuteAttributes {
	this := MuteFindingsMuteAttributes{}
	this.IsMuted = isMuted
	this.Reason = reason
	return &this
}

// NewMuteFindingsMuteAttributesWithDefaults instantiates a new MuteFindingsMuteAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMuteFindingsMuteAttributesWithDefaults() *MuteFindingsMuteAttributes {
	this := MuteFindingsMuteAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MuteFindingsMuteAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingsMuteAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MuteFindingsMuteAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MuteFindingsMuteAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetExpireAt returns the ExpireAt field value if set, zero value otherwise.
func (o *MuteFindingsMuteAttributes) GetExpireAt() int64 {
	if o == nil || o.ExpireAt == nil {
		var ret int64
		return ret
	}
	return *o.ExpireAt
}

// GetExpireAtOk returns a tuple with the ExpireAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingsMuteAttributes) GetExpireAtOk() (*int64, bool) {
	if o == nil || o.ExpireAt == nil {
		return nil, false
	}
	return o.ExpireAt, true
}

// HasExpireAt returns a boolean if a field has been set.
func (o *MuteFindingsMuteAttributes) HasExpireAt() bool {
	return o != nil && o.ExpireAt != nil
}

// SetExpireAt gets a reference to the given int64 and assigns it to the ExpireAt field.
func (o *MuteFindingsMuteAttributes) SetExpireAt(v int64) {
	o.ExpireAt = &v
}

// GetIsMuted returns the IsMuted field value.
func (o *MuteFindingsMuteAttributes) GetIsMuted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value
// and a boolean to check if the value has been set.
func (o *MuteFindingsMuteAttributes) GetIsMutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMuted, true
}

// SetIsMuted sets field value.
func (o *MuteFindingsMuteAttributes) SetIsMuted(v bool) {
	o.IsMuted = v
}

// GetReason returns the Reason field value.
func (o *MuteFindingsMuteAttributes) GetReason() MuteFindingsReason {
	if o == nil {
		var ret MuteFindingsReason
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *MuteFindingsMuteAttributes) GetReasonOk() (*MuteFindingsReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *MuteFindingsMuteAttributes) SetReason(v MuteFindingsReason) {
	o.Reason = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MuteFindingsMuteAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExpireAt != nil {
		toSerialize["expire_at"] = o.ExpireAt
	}
	toSerialize["is_muted"] = o.IsMuted
	toSerialize["reason"] = o.Reason

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MuteFindingsMuteAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string             `json:"description,omitempty"`
		ExpireAt    *int64              `json:"expire_at,omitempty"`
		IsMuted     *bool               `json:"is_muted"`
		Reason      *MuteFindingsReason `json:"reason"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsMuted == nil {
		return fmt.Errorf("required field is_muted missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "expire_at", "is_muted", "reason"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.ExpireAt = all.ExpireAt
	o.IsMuted = *all.IsMuted
	if !all.Reason.IsValid() {
		hasInvalidField = true
	} else {
		o.Reason = *all.Reason
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
