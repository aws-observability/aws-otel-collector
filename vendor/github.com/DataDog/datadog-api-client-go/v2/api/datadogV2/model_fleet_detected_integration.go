// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDetectedIntegration An integration detected on the agent but not necessarily configured.
type FleetDetectedIntegration struct {
	// Escaped integration name.
	EscapedName *string `json:"escaped_name,omitempty"`
	// Integration prefix identifier.
	Prefix *string `json:"prefix,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDetectedIntegration instantiates a new FleetDetectedIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDetectedIntegration() *FleetDetectedIntegration {
	this := FleetDetectedIntegration{}
	return &this
}

// NewFleetDetectedIntegrationWithDefaults instantiates a new FleetDetectedIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDetectedIntegrationWithDefaults() *FleetDetectedIntegration {
	this := FleetDetectedIntegration{}
	return &this
}

// GetEscapedName returns the EscapedName field value if set, zero value otherwise.
func (o *FleetDetectedIntegration) GetEscapedName() string {
	if o == nil || o.EscapedName == nil {
		var ret string
		return ret
	}
	return *o.EscapedName
}

// GetEscapedNameOk returns a tuple with the EscapedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDetectedIntegration) GetEscapedNameOk() (*string, bool) {
	if o == nil || o.EscapedName == nil {
		return nil, false
	}
	return o.EscapedName, true
}

// HasEscapedName returns a boolean if a field has been set.
func (o *FleetDetectedIntegration) HasEscapedName() bool {
	return o != nil && o.EscapedName != nil
}

// SetEscapedName gets a reference to the given string and assigns it to the EscapedName field.
func (o *FleetDetectedIntegration) SetEscapedName(v string) {
	o.EscapedName = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *FleetDetectedIntegration) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDetectedIntegration) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *FleetDetectedIntegration) HasPrefix() bool {
	return o != nil && o.Prefix != nil
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *FleetDetectedIntegration) SetPrefix(v string) {
	o.Prefix = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDetectedIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EscapedName != nil {
		toSerialize["escaped_name"] = o.EscapedName
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDetectedIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EscapedName *string `json:"escaped_name,omitempty"`
		Prefix      *string `json:"prefix,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"escaped_name", "prefix"})
	} else {
		return err
	}
	o.EscapedName = all.EscapedName
	o.Prefix = all.Prefix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
