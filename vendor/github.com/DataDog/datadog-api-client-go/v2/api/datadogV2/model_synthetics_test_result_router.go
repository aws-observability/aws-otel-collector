// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultRouter A router along the traceroute path.
type SyntheticsTestResultRouter struct {
	// IP address of the router.
	Ip *string `json:"ip,omitempty"`
	// Resolved hostname of the router.
	ResolvedHost *string `json:"resolved_host,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultRouter instantiates a new SyntheticsTestResultRouter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultRouter() *SyntheticsTestResultRouter {
	this := SyntheticsTestResultRouter{}
	return &this
}

// NewSyntheticsTestResultRouterWithDefaults instantiates a new SyntheticsTestResultRouter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultRouterWithDefaults() *SyntheticsTestResultRouter {
	this := SyntheticsTestResultRouter{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *SyntheticsTestResultRouter) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRouter) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *SyntheticsTestResultRouter) HasIp() bool {
	return o != nil && o.Ip != nil
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *SyntheticsTestResultRouter) SetIp(v string) {
	o.Ip = &v
}

// GetResolvedHost returns the ResolvedHost field value if set, zero value otherwise.
func (o *SyntheticsTestResultRouter) GetResolvedHost() string {
	if o == nil || o.ResolvedHost == nil {
		var ret string
		return ret
	}
	return *o.ResolvedHost
}

// GetResolvedHostOk returns a tuple with the ResolvedHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRouter) GetResolvedHostOk() (*string, bool) {
	if o == nil || o.ResolvedHost == nil {
		return nil, false
	}
	return o.ResolvedHost, true
}

// HasResolvedHost returns a boolean if a field has been set.
func (o *SyntheticsTestResultRouter) HasResolvedHost() bool {
	return o != nil && o.ResolvedHost != nil
}

// SetResolvedHost gets a reference to the given string and assigns it to the ResolvedHost field.
func (o *SyntheticsTestResultRouter) SetResolvedHost(v string) {
	o.ResolvedHost = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultRouter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.ResolvedHost != nil {
		toSerialize["resolved_host"] = o.ResolvedHost
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultRouter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ip           *string `json:"ip,omitempty"`
		ResolvedHost *string `json:"resolved_host,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ip", "resolved_host"})
	} else {
		return err
	}
	o.Ip = all.Ip
	o.ResolvedHost = all.ResolvedHost

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
