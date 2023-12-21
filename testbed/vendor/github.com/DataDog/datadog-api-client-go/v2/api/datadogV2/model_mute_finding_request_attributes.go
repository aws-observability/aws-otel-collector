// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// MuteFindingRequestAttributes The mute properties to be updated.
type MuteFindingRequestAttributes struct {
	// Object containing the new mute properties of the finding.
	Mute MuteFindingRequestProperties `json:"mute"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMuteFindingRequestAttributes instantiates a new MuteFindingRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMuteFindingRequestAttributes(mute MuteFindingRequestProperties) *MuteFindingRequestAttributes {
	this := MuteFindingRequestAttributes{}
	this.Mute = mute
	return &this
}

// NewMuteFindingRequestAttributesWithDefaults instantiates a new MuteFindingRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMuteFindingRequestAttributesWithDefaults() *MuteFindingRequestAttributes {
	this := MuteFindingRequestAttributes{}
	return &this
}

// GetMute returns the Mute field value.
func (o *MuteFindingRequestAttributes) GetMute() MuteFindingRequestProperties {
	if o == nil {
		var ret MuteFindingRequestProperties
		return ret
	}
	return o.Mute
}

// GetMuteOk returns a tuple with the Mute field value
// and a boolean to check if the value has been set.
func (o *MuteFindingRequestAttributes) GetMuteOk() (*MuteFindingRequestProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mute, true
}

// SetMute sets field value.
func (o *MuteFindingRequestAttributes) SetMute(v MuteFindingRequestProperties) {
	o.Mute = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MuteFindingRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["mute"] = o.Mute
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MuteFindingRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mute *MuteFindingRequestProperties `json:"mute"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mute == nil {
		return fmt.Errorf("required field mute missing")
	}

	hasInvalidField := false
	if all.Mute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Mute = *all.Mute

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
