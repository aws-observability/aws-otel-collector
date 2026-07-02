// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SignalEntitiesAttributes Attributes containing the entities related to the signal.
type SignalEntitiesAttributes struct {
	// The identity entities related to the signal. Each item is a free-form object describing an identity (for example, a user or principal).
	Identities []map[string]interface{} `json:"identities"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSignalEntitiesAttributes instantiates a new SignalEntitiesAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSignalEntitiesAttributes(identities []map[string]interface{}) *SignalEntitiesAttributes {
	this := SignalEntitiesAttributes{}
	this.Identities = identities
	return &this
}

// NewSignalEntitiesAttributesWithDefaults instantiates a new SignalEntitiesAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSignalEntitiesAttributesWithDefaults() *SignalEntitiesAttributes {
	this := SignalEntitiesAttributes{}
	return &this
}

// GetIdentities returns the Identities field value.
func (o *SignalEntitiesAttributes) GetIdentities() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value
// and a boolean to check if the value has been set.
func (o *SignalEntitiesAttributes) GetIdentitiesOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identities, true
}

// SetIdentities sets field value.
func (o *SignalEntitiesAttributes) SetIdentities(v []map[string]interface{}) {
	o.Identities = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SignalEntitiesAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["identities"] = o.Identities

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SignalEntitiesAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Identities *[]map[string]interface{} `json:"identities"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Identities == nil {
		return fmt.Errorf("required field identities missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"identities"})
	} else {
		return err
	}
	o.Identities = *all.Identities

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
