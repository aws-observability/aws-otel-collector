// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultNetpathHop A single hop along a network path.
type SyntheticsTestResultNetpathHop struct {
	// Resolved hostname of the hop.
	Hostname *string `json:"hostname,omitempty"`
	// IP address of the hop.
	IpAddress *string `json:"ip_address,omitempty"`
	// Whether this hop was reachable.
	Reachable *bool `json:"reachable,omitempty"`
	// Round-trip time to this hop in milliseconds.
	Rtt *float64 `json:"rtt,omitempty"`
	// Time-to-live value of the probe packet at this hop.
	Ttl *int64 `json:"ttl,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultNetpathHop instantiates a new SyntheticsTestResultNetpathHop object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultNetpathHop() *SyntheticsTestResultNetpathHop {
	this := SyntheticsTestResultNetpathHop{}
	return &this
}

// NewSyntheticsTestResultNetpathHopWithDefaults instantiates a new SyntheticsTestResultNetpathHop object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultNetpathHopWithDefaults() *SyntheticsTestResultNetpathHop {
	this := SyntheticsTestResultNetpathHop{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpathHop) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpathHop) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpathHop) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SyntheticsTestResultNetpathHop) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpathHop) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpathHop) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpathHop) HasIpAddress() bool {
	return o != nil && o.IpAddress != nil
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *SyntheticsTestResultNetpathHop) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetReachable returns the Reachable field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpathHop) GetReachable() bool {
	if o == nil || o.Reachable == nil {
		var ret bool
		return ret
	}
	return *o.Reachable
}

// GetReachableOk returns a tuple with the Reachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpathHop) GetReachableOk() (*bool, bool) {
	if o == nil || o.Reachable == nil {
		return nil, false
	}
	return o.Reachable, true
}

// HasReachable returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpathHop) HasReachable() bool {
	return o != nil && o.Reachable != nil
}

// SetReachable gets a reference to the given bool and assigns it to the Reachable field.
func (o *SyntheticsTestResultNetpathHop) SetReachable(v bool) {
	o.Reachable = &v
}

// GetRtt returns the Rtt field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpathHop) GetRtt() float64 {
	if o == nil || o.Rtt == nil {
		var ret float64
		return ret
	}
	return *o.Rtt
}

// GetRttOk returns a tuple with the Rtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpathHop) GetRttOk() (*float64, bool) {
	if o == nil || o.Rtt == nil {
		return nil, false
	}
	return o.Rtt, true
}

// HasRtt returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpathHop) HasRtt() bool {
	return o != nil && o.Rtt != nil
}

// SetRtt gets a reference to the given float64 and assigns it to the Rtt field.
func (o *SyntheticsTestResultNetpathHop) SetRtt(v float64) {
	o.Rtt = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpathHop) GetTtl() int64 {
	if o == nil || o.Ttl == nil {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpathHop) GetTtlOk() (*int64, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpathHop) HasTtl() bool {
	return o != nil && o.Ttl != nil
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *SyntheticsTestResultNetpathHop) SetTtl(v int64) {
	o.Ttl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultNetpathHop) MarshalJSON() ([]byte, error) {
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
	if o.Reachable != nil {
		toSerialize["reachable"] = o.Reachable
	}
	if o.Rtt != nil {
		toSerialize["rtt"] = o.Rtt
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultNetpathHop) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hostname  *string  `json:"hostname,omitempty"`
		IpAddress *string  `json:"ip_address,omitempty"`
		Reachable *bool    `json:"reachable,omitempty"`
		Rtt       *float64 `json:"rtt,omitempty"`
		Ttl       *int64   `json:"ttl,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hostname", "ip_address", "reachable", "rtt", "ttl"})
	} else {
		return err
	}
	o.Hostname = all.Hostname
	o.IpAddress = all.IpAddress
	o.Reachable = all.Reachable
	o.Rtt = all.Rtt
	o.Ttl = all.Ttl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
