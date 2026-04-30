// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultDnsResolution DNS resolution details recorded during the test execution.
type SyntheticsTestResultDnsResolution struct {
	// DNS resolution attempts made during the test.
	Attempts []map[string]string `json:"attempts,omitempty"`
	// Resolved IP address for the target host.
	ResolvedIp *string `json:"resolved_ip,omitempty"`
	// Resolved port for the target service.
	ResolvedPort *string `json:"resolved_port,omitempty"`
	// DNS server used for the resolution.
	Server *string `json:"server,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultDnsResolution instantiates a new SyntheticsTestResultDnsResolution object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultDnsResolution() *SyntheticsTestResultDnsResolution {
	this := SyntheticsTestResultDnsResolution{}
	return &this
}

// NewSyntheticsTestResultDnsResolutionWithDefaults instantiates a new SyntheticsTestResultDnsResolution object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultDnsResolutionWithDefaults() *SyntheticsTestResultDnsResolution {
	this := SyntheticsTestResultDnsResolution{}
	return &this
}

// GetAttempts returns the Attempts field value if set, zero value otherwise.
func (o *SyntheticsTestResultDnsResolution) GetAttempts() []map[string]string {
	if o == nil || o.Attempts == nil {
		var ret []map[string]string
		return ret
	}
	return o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDnsResolution) GetAttemptsOk() (*[]map[string]string, bool) {
	if o == nil || o.Attempts == nil {
		return nil, false
	}
	return &o.Attempts, true
}

// HasAttempts returns a boolean if a field has been set.
func (o *SyntheticsTestResultDnsResolution) HasAttempts() bool {
	return o != nil && o.Attempts != nil
}

// SetAttempts gets a reference to the given []map[string]string and assigns it to the Attempts field.
func (o *SyntheticsTestResultDnsResolution) SetAttempts(v []map[string]string) {
	o.Attempts = v
}

// GetResolvedIp returns the ResolvedIp field value if set, zero value otherwise.
func (o *SyntheticsTestResultDnsResolution) GetResolvedIp() string {
	if o == nil || o.ResolvedIp == nil {
		var ret string
		return ret
	}
	return *o.ResolvedIp
}

// GetResolvedIpOk returns a tuple with the ResolvedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDnsResolution) GetResolvedIpOk() (*string, bool) {
	if o == nil || o.ResolvedIp == nil {
		return nil, false
	}
	return o.ResolvedIp, true
}

// HasResolvedIp returns a boolean if a field has been set.
func (o *SyntheticsTestResultDnsResolution) HasResolvedIp() bool {
	return o != nil && o.ResolvedIp != nil
}

// SetResolvedIp gets a reference to the given string and assigns it to the ResolvedIp field.
func (o *SyntheticsTestResultDnsResolution) SetResolvedIp(v string) {
	o.ResolvedIp = &v
}

// GetResolvedPort returns the ResolvedPort field value if set, zero value otherwise.
func (o *SyntheticsTestResultDnsResolution) GetResolvedPort() string {
	if o == nil || o.ResolvedPort == nil {
		var ret string
		return ret
	}
	return *o.ResolvedPort
}

// GetResolvedPortOk returns a tuple with the ResolvedPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDnsResolution) GetResolvedPortOk() (*string, bool) {
	if o == nil || o.ResolvedPort == nil {
		return nil, false
	}
	return o.ResolvedPort, true
}

// HasResolvedPort returns a boolean if a field has been set.
func (o *SyntheticsTestResultDnsResolution) HasResolvedPort() bool {
	return o != nil && o.ResolvedPort != nil
}

// SetResolvedPort gets a reference to the given string and assigns it to the ResolvedPort field.
func (o *SyntheticsTestResultDnsResolution) SetResolvedPort(v string) {
	o.ResolvedPort = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *SyntheticsTestResultDnsResolution) GetServer() string {
	if o == nil || o.Server == nil {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDnsResolution) GetServerOk() (*string, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *SyntheticsTestResultDnsResolution) HasServer() bool {
	return o != nil && o.Server != nil
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *SyntheticsTestResultDnsResolution) SetServer(v string) {
	o.Server = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultDnsResolution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attempts != nil {
		toSerialize["attempts"] = o.Attempts
	}
	if o.ResolvedIp != nil {
		toSerialize["resolved_ip"] = o.ResolvedIp
	}
	if o.ResolvedPort != nil {
		toSerialize["resolved_port"] = o.ResolvedPort
	}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultDnsResolution) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attempts     []map[string]string `json:"attempts,omitempty"`
		ResolvedIp   *string             `json:"resolved_ip,omitempty"`
		ResolvedPort *string             `json:"resolved_port,omitempty"`
		Server       *string             `json:"server,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attempts", "resolved_ip", "resolved_port", "server"})
	} else {
		return err
	}
	o.Attempts = all.Attempts
	o.ResolvedIp = all.ResolvedIp
	o.ResolvedPort = all.ResolvedPort
	o.Server = all.Server

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
