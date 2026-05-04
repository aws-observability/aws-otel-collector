// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetInstrumentedPodsResponseDataAttributes Attributes of the instrumented pods response containing the list of pod groups.
type FleetInstrumentedPodsResponseDataAttributes struct {
	// Array of instrumented pod groups in the cluster.
	Groups []FleetInstrumentedPodGroupAttributes `json:"groups,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetInstrumentedPodsResponseDataAttributes instantiates a new FleetInstrumentedPodsResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetInstrumentedPodsResponseDataAttributes() *FleetInstrumentedPodsResponseDataAttributes {
	this := FleetInstrumentedPodsResponseDataAttributes{}
	return &this
}

// NewFleetInstrumentedPodsResponseDataAttributesWithDefaults instantiates a new FleetInstrumentedPodsResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetInstrumentedPodsResponseDataAttributesWithDefaults() *FleetInstrumentedPodsResponseDataAttributes {
	this := FleetInstrumentedPodsResponseDataAttributes{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *FleetInstrumentedPodsResponseDataAttributes) GetGroups() []FleetInstrumentedPodGroupAttributes {
	if o == nil || o.Groups == nil {
		var ret []FleetInstrumentedPodGroupAttributes
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodsResponseDataAttributes) GetGroupsOk() (*[]FleetInstrumentedPodGroupAttributes, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *FleetInstrumentedPodsResponseDataAttributes) HasGroups() bool {
	return o != nil && o.Groups != nil
}

// SetGroups gets a reference to the given []FleetInstrumentedPodGroupAttributes and assigns it to the Groups field.
func (o *FleetInstrumentedPodsResponseDataAttributes) SetGroups(v []FleetInstrumentedPodGroupAttributes) {
	o.Groups = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetInstrumentedPodsResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetInstrumentedPodsResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Groups []FleetInstrumentedPodGroupAttributes `json:"groups,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"groups"})
	} else {
		return err
	}
	o.Groups = all.Groups

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
