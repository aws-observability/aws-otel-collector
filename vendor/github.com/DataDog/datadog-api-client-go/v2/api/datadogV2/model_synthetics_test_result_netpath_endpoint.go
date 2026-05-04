// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultNetpathEndpoint Source endpoint of a network path measurement.
type SyntheticsTestResultNetpathEndpoint struct {
	// Hostname of the endpoint.
	Hostname *string `json:"hostname,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultNetpathEndpoint instantiates a new SyntheticsTestResultNetpathEndpoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultNetpathEndpoint() *SyntheticsTestResultNetpathEndpoint {
	this := SyntheticsTestResultNetpathEndpoint{}
	return &this
}

// NewSyntheticsTestResultNetpathEndpointWithDefaults instantiates a new SyntheticsTestResultNetpathEndpoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultNetpathEndpointWithDefaults() *SyntheticsTestResultNetpathEndpoint {
	this := SyntheticsTestResultNetpathEndpoint{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpathEndpoint) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpathEndpoint) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpathEndpoint) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SyntheticsTestResultNetpathEndpoint) SetHostname(v string) {
	o.Hostname = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultNetpathEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultNetpathEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hostname *string `json:"hostname,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hostname"})
	} else {
		return err
	}
	o.Hostname = all.Hostname

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
