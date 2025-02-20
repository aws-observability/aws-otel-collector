// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BulkMuteFindingsRequestAttributes The mute properties to be updated.
type BulkMuteFindingsRequestAttributes struct {
	// Object containing the new mute properties of the findings.
	Mute BulkMuteFindingsRequestProperties `json:"mute"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewBulkMuteFindingsRequestAttributes instantiates a new BulkMuteFindingsRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBulkMuteFindingsRequestAttributes(mute BulkMuteFindingsRequestProperties) *BulkMuteFindingsRequestAttributes {
	this := BulkMuteFindingsRequestAttributes{}
	this.Mute = mute
	return &this
}

// NewBulkMuteFindingsRequestAttributesWithDefaults instantiates a new BulkMuteFindingsRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBulkMuteFindingsRequestAttributesWithDefaults() *BulkMuteFindingsRequestAttributes {
	this := BulkMuteFindingsRequestAttributes{}
	return &this
}

// GetMute returns the Mute field value.
func (o *BulkMuteFindingsRequestAttributes) GetMute() BulkMuteFindingsRequestProperties {
	if o == nil {
		var ret BulkMuteFindingsRequestProperties
		return ret
	}
	return o.Mute
}

// GetMuteOk returns a tuple with the Mute field value
// and a boolean to check if the value has been set.
func (o *BulkMuteFindingsRequestAttributes) GetMuteOk() (*BulkMuteFindingsRequestProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mute, true
}

// SetMute sets field value.
func (o *BulkMuteFindingsRequestAttributes) SetMute(v BulkMuteFindingsRequestProperties) {
	o.Mute = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BulkMuteFindingsRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["mute"] = o.Mute
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BulkMuteFindingsRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mute *BulkMuteFindingsRequestProperties `json:"mute"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
