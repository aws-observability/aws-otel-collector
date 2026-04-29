// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultHandshake Handshake request and response for protocol-level tests.
type SyntheticsTestResultHandshake struct {
	// Details of the outgoing request made during the test execution.
	Request *SyntheticsTestResultRequestInfo `json:"request,omitempty"`
	// Details of the response received during the test execution.
	Response *SyntheticsTestResultResponseInfo `json:"response,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultHandshake instantiates a new SyntheticsTestResultHandshake object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultHandshake() *SyntheticsTestResultHandshake {
	this := SyntheticsTestResultHandshake{}
	return &this
}

// NewSyntheticsTestResultHandshakeWithDefaults instantiates a new SyntheticsTestResultHandshake object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultHandshakeWithDefaults() *SyntheticsTestResultHandshake {
	this := SyntheticsTestResultHandshake{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SyntheticsTestResultHandshake) GetRequest() SyntheticsTestResultRequestInfo {
	if o == nil || o.Request == nil {
		var ret SyntheticsTestResultRequestInfo
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultHandshake) GetRequestOk() (*SyntheticsTestResultRequestInfo, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SyntheticsTestResultHandshake) HasRequest() bool {
	return o != nil && o.Request != nil
}

// SetRequest gets a reference to the given SyntheticsTestResultRequestInfo and assigns it to the Request field.
func (o *SyntheticsTestResultHandshake) SetRequest(v SyntheticsTestResultRequestInfo) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SyntheticsTestResultHandshake) GetResponse() SyntheticsTestResultResponseInfo {
	if o == nil || o.Response == nil {
		var ret SyntheticsTestResultResponseInfo
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultHandshake) GetResponseOk() (*SyntheticsTestResultResponseInfo, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SyntheticsTestResultHandshake) HasResponse() bool {
	return o != nil && o.Response != nil
}

// SetResponse gets a reference to the given SyntheticsTestResultResponseInfo and assigns it to the Response field.
func (o *SyntheticsTestResultHandshake) SetResponse(v SyntheticsTestResultResponseInfo) {
	o.Response = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultHandshake) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultHandshake) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Request  *SyntheticsTestResultRequestInfo  `json:"request,omitempty"`
		Response *SyntheticsTestResultResponseInfo `json:"response,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"request", "response"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Request != nil && all.Request.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Request = all.Request
	if all.Response != nil && all.Response.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Response = all.Response

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
