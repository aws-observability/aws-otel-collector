// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamSyncAttributes Team sync attributes.
type TeamSyncAttributes struct {
	// How often the sync process should be run. Defaults to `once` when not provided.
	Frequency *TeamSyncAttributesFrequency `json:"frequency,omitempty"`
	// The external source platform for team synchronization. Only "github" is supported.
	Source TeamSyncAttributesSource `json:"source"`
	// Whether to sync members from the external team to the Datadog team. Defaults to `false` when not provided.
	SyncMembership *bool `json:"sync_membership,omitempty"`
	// The type of synchronization operation. "link" connects teams by matching names. "provision" creates new teams when no match is found.
	Type TeamSyncAttributesType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamSyncAttributes instantiates a new TeamSyncAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamSyncAttributes(source TeamSyncAttributesSource, typeVar TeamSyncAttributesType) *TeamSyncAttributes {
	this := TeamSyncAttributes{}
	this.Source = source
	this.Type = typeVar
	return &this
}

// NewTeamSyncAttributesWithDefaults instantiates a new TeamSyncAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamSyncAttributesWithDefaults() *TeamSyncAttributes {
	this := TeamSyncAttributes{}
	return &this
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *TeamSyncAttributes) GetFrequency() TeamSyncAttributesFrequency {
	if o == nil || o.Frequency == nil {
		var ret TeamSyncAttributesFrequency
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamSyncAttributes) GetFrequencyOk() (*TeamSyncAttributesFrequency, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *TeamSyncAttributes) HasFrequency() bool {
	return o != nil && o.Frequency != nil
}

// SetFrequency gets a reference to the given TeamSyncAttributesFrequency and assigns it to the Frequency field.
func (o *TeamSyncAttributes) SetFrequency(v TeamSyncAttributesFrequency) {
	o.Frequency = &v
}

// GetSource returns the Source field value.
func (o *TeamSyncAttributes) GetSource() TeamSyncAttributesSource {
	if o == nil {
		var ret TeamSyncAttributesSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TeamSyncAttributes) GetSourceOk() (*TeamSyncAttributesSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *TeamSyncAttributes) SetSource(v TeamSyncAttributesSource) {
	o.Source = v
}

// GetSyncMembership returns the SyncMembership field value if set, zero value otherwise.
func (o *TeamSyncAttributes) GetSyncMembership() bool {
	if o == nil || o.SyncMembership == nil {
		var ret bool
		return ret
	}
	return *o.SyncMembership
}

// GetSyncMembershipOk returns a tuple with the SyncMembership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamSyncAttributes) GetSyncMembershipOk() (*bool, bool) {
	if o == nil || o.SyncMembership == nil {
		return nil, false
	}
	return o.SyncMembership, true
}

// HasSyncMembership returns a boolean if a field has been set.
func (o *TeamSyncAttributes) HasSyncMembership() bool {
	return o != nil && o.SyncMembership != nil
}

// SetSyncMembership gets a reference to the given bool and assigns it to the SyncMembership field.
func (o *TeamSyncAttributes) SetSyncMembership(v bool) {
	o.SyncMembership = &v
}

// GetType returns the Type field value.
func (o *TeamSyncAttributes) GetType() TeamSyncAttributesType {
	if o == nil {
		var ret TeamSyncAttributesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamSyncAttributes) GetTypeOk() (*TeamSyncAttributesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamSyncAttributes) SetType(v TeamSyncAttributesType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamSyncAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	toSerialize["source"] = o.Source
	if o.SyncMembership != nil {
		toSerialize["sync_membership"] = o.SyncMembership
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamSyncAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Frequency      *TeamSyncAttributesFrequency `json:"frequency,omitempty"`
		Source         *TeamSyncAttributesSource    `json:"source"`
		SyncMembership *bool                        `json:"sync_membership,omitempty"`
		Type           *TeamSyncAttributesType      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"frequency", "source", "sync_membership", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Frequency != nil && !all.Frequency.IsValid() {
		hasInvalidField = true
	} else {
		o.Frequency = all.Frequency
	}
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.SyncMembership = all.SyncMembership
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
