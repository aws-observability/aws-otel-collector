// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MuteFindingsRequestDataAttributes Attributes of the mute request.
type MuteFindingsRequestDataAttributes struct {
	// Mute properties to apply to the findings.
	Mute MuteFindingsMuteAttributes `json:"mute"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMuteFindingsRequestDataAttributes instantiates a new MuteFindingsRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMuteFindingsRequestDataAttributes(mute MuteFindingsMuteAttributes) *MuteFindingsRequestDataAttributes {
	this := MuteFindingsRequestDataAttributes{}
	this.Mute = mute
	return &this
}

// NewMuteFindingsRequestDataAttributesWithDefaults instantiates a new MuteFindingsRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMuteFindingsRequestDataAttributesWithDefaults() *MuteFindingsRequestDataAttributes {
	this := MuteFindingsRequestDataAttributes{}
	return &this
}

// GetMute returns the Mute field value.
func (o *MuteFindingsRequestDataAttributes) GetMute() MuteFindingsMuteAttributes {
	if o == nil {
		var ret MuteFindingsMuteAttributes
		return ret
	}
	return o.Mute
}

// GetMuteOk returns a tuple with the Mute field value
// and a boolean to check if the value has been set.
func (o *MuteFindingsRequestDataAttributes) GetMuteOk() (*MuteFindingsMuteAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mute, true
}

// SetMute sets field value.
func (o *MuteFindingsRequestDataAttributes) SetMute(v MuteFindingsMuteAttributes) {
	o.Mute = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MuteFindingsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["mute"] = o.Mute

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MuteFindingsRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mute *MuteFindingsMuteAttributes `json:"mute"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mute == nil {
		return fmt.Errorf("required field mute missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"mute"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Mute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Mute = *all.Mute

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
