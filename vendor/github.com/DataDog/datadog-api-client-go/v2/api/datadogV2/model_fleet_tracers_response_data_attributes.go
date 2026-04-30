// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetTracersResponseDataAttributes Attributes of the fleet tracers response containing the list of tracers.
type FleetTracersResponseDataAttributes struct {
	// Array of tracers matching the query criteria.
	Tracers []FleetTracerAttributes `json:"tracers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetTracersResponseDataAttributes instantiates a new FleetTracersResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetTracersResponseDataAttributes() *FleetTracersResponseDataAttributes {
	this := FleetTracersResponseDataAttributes{}
	return &this
}

// NewFleetTracersResponseDataAttributesWithDefaults instantiates a new FleetTracersResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetTracersResponseDataAttributesWithDefaults() *FleetTracersResponseDataAttributes {
	this := FleetTracersResponseDataAttributes{}
	return &this
}

// GetTracers returns the Tracers field value if set, zero value otherwise.
func (o *FleetTracersResponseDataAttributes) GetTracers() []FleetTracerAttributes {
	if o == nil || o.Tracers == nil {
		var ret []FleetTracerAttributes
		return ret
	}
	return o.Tracers
}

// GetTracersOk returns a tuple with the Tracers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracersResponseDataAttributes) GetTracersOk() (*[]FleetTracerAttributes, bool) {
	if o == nil || o.Tracers == nil {
		return nil, false
	}
	return &o.Tracers, true
}

// HasTracers returns a boolean if a field has been set.
func (o *FleetTracersResponseDataAttributes) HasTracers() bool {
	return o != nil && o.Tracers != nil
}

// SetTracers gets a reference to the given []FleetTracerAttributes and assigns it to the Tracers field.
func (o *FleetTracersResponseDataAttributes) SetTracers(v []FleetTracerAttributes) {
	o.Tracers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetTracersResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Tracers != nil {
		toSerialize["tracers"] = o.Tracers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetTracersResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tracers []FleetTracerAttributes `json:"tracers,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tracers"})
	} else {
		return err
	}
	o.Tracers = all.Tracers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
