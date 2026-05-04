// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultNetpathDestination Destination endpoint of a network path measurement.
type SyntheticsTestResultNetpathDestination struct {
	// Hostname of the destination.
	Hostname *string `json:"hostname,omitempty"`
	// IP address of the destination.
	IpAddress *string `json:"ip_address,omitempty"`
	// Port of the destination service.
	Port *int64 `json:"port,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultNetpathDestination instantiates a new SyntheticsTestResultNetpathDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultNetpathDestination() *SyntheticsTestResultNetpathDestination {
	this := SyntheticsTestResultNetpathDestination{}
	return &this
}

// NewSyntheticsTestResultNetpathDestinationWithDefaults instantiates a new SyntheticsTestResultNetpathDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultNetpathDestinationWithDefaults() *SyntheticsTestResultNetpathDestination {
	this := SyntheticsTestResultNetpathDestination{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpathDestination) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpathDestination) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpathDestination) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SyntheticsTestResultNetpathDestination) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpathDestination) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpathDestination) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpathDestination) HasIpAddress() bool {
	return o != nil && o.IpAddress != nil
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *SyntheticsTestResultNetpathDestination) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpathDestination) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpathDestination) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpathDestination) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *SyntheticsTestResultNetpathDestination) SetPort(v int64) {
	o.Port = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultNetpathDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.IpAddress != nil {
		toSerialize["ip_address"] = o.IpAddress
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultNetpathDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hostname  *string `json:"hostname,omitempty"`
		IpAddress *string `json:"ip_address,omitempty"`
		Port      *int64  `json:"port,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hostname", "ip_address", "port"})
	} else {
		return err
	}
	o.Hostname = all.Hostname
	o.IpAddress = all.IpAddress
	o.Port = all.Port

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
