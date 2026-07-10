// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader Extract the token from a specific HTTP request header.
type ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader struct {
	// The name of the HTTP header that carries the token.
	Header string `json:"header"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader instantiates a new ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader(header string) *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader {
	this := ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader{}
	this.Header = header
	return &this
}

// NewObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeaderWithDefaults instantiates a new ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeaderWithDefaults() *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader {
	this := ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader{}
	return &this
}

// GetHeader returns the Header field value.
func (o *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader) GetHeader() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Header, true
}

// SetHeader sets field value.
func (o *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader) SetHeader(v string) {
	o.Header = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["header"] = o.Header

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Header *string `json:"header"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Header == nil {
		return fmt.Errorf("required field header missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"header"})
	} else {
		return err
	}
	o.Header = *all.Header

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
